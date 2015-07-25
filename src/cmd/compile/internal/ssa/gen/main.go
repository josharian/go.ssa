// Copyright 2015 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// The gen command generates Go code (in the parent directory) for all
// the architecture-specific opcodes, blocks, and rewrites.

package main

import (
	"bytes"
	"fmt"
	"go/format"
	"io/ioutil"
	"log"
)

type arch struct {
	name     string
	ops      []opData
	blocks   []blockData
	regnames []string
}

type opData struct {
	name string
	reg  regInfo
	asm  string
}

type blockData struct {
	name string
}

type regInfo struct {
	inputs   []regMask
	clobbers regMask
	outputs  []regMask
	inplace  bool
}

type regMask uint64

func (a arch) regMaskComment(r regMask) string {
	var buf bytes.Buffer
	for i := uint64(0); r != 0; i++ {
		if r&1 != 0 {
			if buf.Len() == 0 {
				buf.WriteString(" //")
			}
			buf.WriteString(" ")
			buf.WriteString(a.regnames[i])
		}
		r >>= 1
	}
	return buf.String()
}

var archs []arch

func main() {
	genOp()
	genLower()
}

func genOp() {
	w := new(bytes.Buffer)
	fmt.Fprintf(w, "// autogenerated: do not edit!\n")
	fmt.Fprintf(w, "// generated from gen/*Ops.go\n")
	fmt.Fprintln(w, "package ssa")

	fmt.Fprintln(w, "import \"cmd/internal/obj/x86\"")

	// generate Block* declarations
	fmt.Fprintln(w, "const (")
	fmt.Fprintln(w, "blockInvalid BlockKind = iota")
	for _, a := range archs {
		fmt.Fprintln(w)
		for _, d := range a.blocks {
			fmt.Fprintf(w, "Block%s%s\n", a.Name(), d.name)
		}
	}
	fmt.Fprintln(w, ")")

	// generate block kind string method
	fmt.Fprintln(w, "var blockString = [...]string{")
	fmt.Fprintln(w, "blockInvalid:\"BlockInvalid\",")
	for _, a := range archs {
		fmt.Fprintln(w)
		for _, b := range a.blocks {
			fmt.Fprintf(w, "Block%s%s:\"%s\",\n", a.Name(), b.name, b.name)
		}
	}
	fmt.Fprintln(w, "}")
	fmt.Fprintln(w, "func (k BlockKind) String() string {return blockString[k]}")

	// generate Op* declarations
	fmt.Fprintln(w, "const (")
	fmt.Fprintln(w, "OpInvalid Op = iota")
	for _, a := range archs {
		fmt.Fprintln(w)
		for _, v := range a.ops {
			fmt.Fprintf(w, "Op%s%s\n", a.Name(), v.name)
		}
	}
	fmt.Fprintln(w, ")")

	// generate OpInfo table
	fmt.Fprintln(w, "var opcodeTable = [...]opInfo{")
	fmt.Fprintln(w, " { name: \"OpInvalid\" },")
	for _, a := range archs {
		fmt.Fprintln(w)
		for _, v := range a.ops {
			fmt.Fprintln(w, "{")
			fmt.Fprintf(w, "name:\"%s\",\n", v.name)
			if a.name == "generic" {
				fmt.Fprintln(w, "generic:true,")
				fmt.Fprintln(w, "},") // close op
				// generic ops have no reg info or asm
				continue
			}
			if v.asm != "" {
				fmt.Fprintf(w, "asm: x86.A%s,\n", v.asm)
			}
			fmt.Fprintln(w, "reg:regInfo{")
			// reg inputs
			if len(v.reg.inputs) > 0 {
				fmt.Fprintln(w, "inputs: []regMask{")
				for _, r := range v.reg.inputs {
					fmt.Fprintf(w, "%d,%s\n", r, a.regMaskComment(r))
				}
				fmt.Fprintln(w, "},")
			}
			if v.reg.clobbers > 0 {
				fmt.Fprintf(w, "clobbers: %d,%s\n", v.reg.clobbers, a.regMaskComment(v.reg.clobbers))
			}
			// reg outputs
			if len(v.reg.outputs) > 0 {
				fmt.Fprintln(w, "outputs: []regMask{")
				for _, r := range v.reg.outputs {
					fmt.Fprintf(w, "%d,%s\n", r, a.regMaskComment(r))
				}
				fmt.Fprintln(w, "},")
			}
			if v.reg.inplace {
				fmt.Fprintln(w, "inplace: true,")
			}
			fmt.Fprintln(w, "},") // close reg info
			fmt.Fprintln(w, "},") // close op
		}
	}
	fmt.Fprintln(w, "}")

	fmt.Fprintln(w, "func (o Op) Asm() int {return opcodeTable[o].asm}")

	// generate op string method
	fmt.Fprintln(w, "func (o Op) String() string {return opcodeTable[o].name }")

	// gofmt result
	b := w.Bytes()
	var err error
	b, err = format.Source(b)
	if err != nil {
		panic(err)
	}

	err = ioutil.WriteFile("../opGen.go", b, 0666)
	if err != nil {
		log.Fatalf("can't write output: %v\n", err)
	}
}

// Name returns the name of the architecture for use in Op* and Block* enumerations.
func (a arch) Name() string {
	s := a.name
	if s == "generic" {
		s = ""
	}
	return s
}

func genLower() {
	for _, a := range archs {
		genRules(a)
	}
}
