// autogenerated from gen/generic.rules: do not edit!
// generated with: cd gen; go run *.go
package ssa

func rewriteValuegeneric(v *Value, config *Config) bool {
	switch v.Op {
	case OpAdd64:
		// match: (Add64 (Const64 [c]) (Const64 [d]))
		// cond:
		// result: (Const64 [c+d])
		{
			if v.Args[0].Op != OpConst64 {
				goto end8c46df6f85a11cb1d594076b0e467908
			}
			c := v.Args[0].AuxInt
			if v.Args[1].Op != OpConst64 {
				goto end8c46df6f85a11cb1d594076b0e467908
			}
			d := v.Args[1].AuxInt
			v.Op = OpConst64
			v.AuxInt = 0
			v.Aux = nil
			v.resetArgs()
			v.AuxInt = c + d
			return true
		}
		goto end8c46df6f85a11cb1d594076b0e467908
	end8c46df6f85a11cb1d594076b0e467908:
		;
	case OpAddPtr:
		// match: (AddPtr (ConstPtr [c]) (ConstPtr [d]))
		// cond:
		// result: (ConstPtr [c+d])
		{
			if v.Args[0].Op != OpConstPtr {
				goto end145c1aec793b2befff34bc8983b48a38
			}
			c := v.Args[0].AuxInt
			if v.Args[1].Op != OpConstPtr {
				goto end145c1aec793b2befff34bc8983b48a38
			}
			d := v.Args[1].AuxInt
			v.Op = OpConstPtr
			v.AuxInt = 0
			v.Aux = nil
			v.resetArgs()
			v.AuxInt = c + d
			return true
		}
		goto end145c1aec793b2befff34bc8983b48a38
	end145c1aec793b2befff34bc8983b48a38:
		;
	case OpArrayIndex:
		// match: (ArrayIndex (Load ptr mem) idx)
		// cond:
		// result: (Load (PtrIndex <v.Type().PtrTo()> ptr idx) mem)
		{
			if v.Args[0].Op != OpLoad {
				goto end2b1b8c303f8a75fe0f2923b1f283f672
			}
			ptr := v.Args[0].Args[0]
			mem := v.Args[0].Args[1]
			idx := v.Args[1]
			v.Op = OpLoad
			v.AuxInt = 0
			v.Aux = nil
			v.resetArgs()
			v0 := v.Block.NewValue0(v.Line, OpPtrIndex, TypeInvalid)
			v0.TypeIndex = v.Block.Func.typeIndex(v.Type().PtrTo())
			v0.AddArg(ptr)
			v0.AddArg(idx)
			v.AddArg(v0)
			v.AddArg(mem)
			return true
		}
		goto end2b1b8c303f8a75fe0f2923b1f283f672
	end2b1b8c303f8a75fe0f2923b1f283f672:
		;
	case OpConstString:
		// match: (ConstString {s})
		// cond:
		// result: (StringMake (Addr <TypeBytePtr> {config.fe.StringData(s.(string))} (SB <TypeUintptr>)) (ConstPtr <TypeUintptr> [int64(len(s.(string)))]))
		{
			s := v.Aux
			v.Op = OpStringMake
			v.AuxInt = 0
			v.Aux = nil
			v.resetArgs()
			v0 := v.Block.NewValue0(v.Line, OpAddr, TypeInvalid)
			v0.TypeIndex = v.Block.Func.typeIndex(v.Block.Func.Config.Frontend().TypeBytePtr())
			v0.Aux = config.fe.StringData(s.(string))
			v1 := v.Block.NewValue0(v.Line, OpSB, TypeInvalid)
			v1.TypeIndex = v.Block.Func.typeIndex(v.Block.Func.Config.Frontend().TypeUintptr())
			v0.AddArg(v1)
			v.AddArg(v0)
			v2 := v.Block.NewValue0(v.Line, OpConstPtr, TypeInvalid)
			v2.TypeIndex = v.Block.Func.typeIndex(v.Block.Func.Config.Frontend().TypeUintptr())
			v2.AuxInt = int64(len(s.(string)))
			v.AddArg(v2)
			return true
		}
		goto endfe7f015ae1a9e49ba6991229e1cfa803
	endfe7f015ae1a9e49ba6991229e1cfa803:
		;
	case OpEqFat:
		// match: (EqFat x y)
		// cond: x.Op == OpConstNil && y.Op != OpConstNil
		// result: (EqFat y x)
		{
			x := v.Args[0]
			y := v.Args[1]
			if !(x.Op == OpConstNil && y.Op != OpConstNil) {
				goto endcea7f7399afcff860c54d82230a9a934
			}
			v.Op = OpEqFat
			v.AuxInt = 0
			v.Aux = nil
			v.resetArgs()
			v.AddArg(y)
			v.AddArg(x)
			return true
		}
		goto endcea7f7399afcff860c54d82230a9a934
	endcea7f7399afcff860c54d82230a9a934:
		;
		// match: (EqFat (Load ptr mem) (ConstNil))
		// cond:
		// result: (EqPtr (Load <TypeUintptr> ptr mem) (ConstPtr <TypeUintptr> [0]))
		{
			if v.Args[0].Op != OpLoad {
				goto end63efc0917be74e47be687e35bbef06e2
			}
			ptr := v.Args[0].Args[0]
			mem := v.Args[0].Args[1]
			if v.Args[1].Op != OpConstNil {
				goto end63efc0917be74e47be687e35bbef06e2
			}
			v.Op = OpEqPtr
			v.AuxInt = 0
			v.Aux = nil
			v.resetArgs()
			v0 := v.Block.NewValue0(v.Line, OpLoad, TypeInvalid)
			v0.TypeIndex = v.Block.Func.typeIndex(v.Block.Func.Config.Frontend().TypeUintptr())
			v0.AddArg(ptr)
			v0.AddArg(mem)
			v.AddArg(v0)
			v1 := v.Block.NewValue0(v.Line, OpConstPtr, TypeInvalid)
			v1.TypeIndex = v.Block.Func.typeIndex(v.Block.Func.Config.Frontend().TypeUintptr())
			v1.AuxInt = 0
			v.AddArg(v1)
			return true
		}
		goto end63efc0917be74e47be687e35bbef06e2
	end63efc0917be74e47be687e35bbef06e2:
		;
	case OpIsInBounds:
		// match: (IsInBounds (ConstPtr [c]) (ConstPtr [d]))
		// cond:
		// result: (ConstPtr {inBounds(c,d)})
		{
			if v.Args[0].Op != OpConstPtr {
				goto enddfd340bc7103ca323354aec96b113c23
			}
			c := v.Args[0].AuxInt
			if v.Args[1].Op != OpConstPtr {
				goto enddfd340bc7103ca323354aec96b113c23
			}
			d := v.Args[1].AuxInt
			v.Op = OpConstPtr
			v.AuxInt = 0
			v.Aux = nil
			v.resetArgs()
			v.Aux = inBounds(c, d)
			return true
		}
		goto enddfd340bc7103ca323354aec96b113c23
	enddfd340bc7103ca323354aec96b113c23:
		;
	case OpLoad:
		// match: (Load <t> ptr mem)
		// cond: t.IsString()
		// result: (StringMake (Load <TypeBytePtr> ptr mem) (Load <TypeUintptr> (OffPtr <TypeBytePtr> [config.PtrSize] ptr) mem))
		{
			t := v.Type()
			ptr := v.Args[0]
			mem := v.Args[1]
			if !(t.IsString()) {
				goto endf9240d058613f1432e6225bba6c9404f
			}
			v.Op = OpStringMake
			v.AuxInt = 0
			v.Aux = nil
			v.resetArgs()
			v0 := v.Block.NewValue0(v.Line, OpLoad, TypeInvalid)
			v0.TypeIndex = v.Block.Func.typeIndex(v.Block.Func.Config.Frontend().TypeBytePtr())
			v0.AddArg(ptr)
			v0.AddArg(mem)
			v.AddArg(v0)
			v1 := v.Block.NewValue0(v.Line, OpLoad, TypeInvalid)
			v1.TypeIndex = v.Block.Func.typeIndex(v.Block.Func.Config.Frontend().TypeUintptr())
			v2 := v.Block.NewValue0(v.Line, OpOffPtr, TypeInvalid)
			v2.TypeIndex = v.Block.Func.typeIndex(v.Block.Func.Config.Frontend().TypeBytePtr())
			v2.AuxInt = config.PtrSize
			v2.AddArg(ptr)
			v1.AddArg(v2)
			v1.AddArg(mem)
			v.AddArg(v1)
			return true
		}
		goto endf9240d058613f1432e6225bba6c9404f
	endf9240d058613f1432e6225bba6c9404f:
		;
	case OpMul64:
		// match: (Mul64 (Const64 [c]) (Const64 [d]))
		// cond:
		// result: (Const64 [c*d])
		{
			if v.Args[0].Op != OpConst64 {
				goto end7aea1048b5d1230974b97f17238380ae
			}
			c := v.Args[0].AuxInt
			if v.Args[1].Op != OpConst64 {
				goto end7aea1048b5d1230974b97f17238380ae
			}
			d := v.Args[1].AuxInt
			v.Op = OpConst64
			v.AuxInt = 0
			v.Aux = nil
			v.resetArgs()
			v.AuxInt = c * d
			return true
		}
		goto end7aea1048b5d1230974b97f17238380ae
	end7aea1048b5d1230974b97f17238380ae:
		;
	case OpMulPtr:
		// match: (MulPtr (ConstPtr [c]) (ConstPtr [d]))
		// cond:
		// result: (ConstPtr [c*d])
		{
			if v.Args[0].Op != OpConstPtr {
				goto end808c190f346658bb1ad032bf37a1059f
			}
			c := v.Args[0].AuxInt
			if v.Args[1].Op != OpConstPtr {
				goto end808c190f346658bb1ad032bf37a1059f
			}
			d := v.Args[1].AuxInt
			v.Op = OpConstPtr
			v.AuxInt = 0
			v.Aux = nil
			v.resetArgs()
			v.AuxInt = c * d
			return true
		}
		goto end808c190f346658bb1ad032bf37a1059f
	end808c190f346658bb1ad032bf37a1059f:
		;
	case OpNeqFat:
		// match: (NeqFat x y)
		// cond: x.Op == OpConstNil && y.Op != OpConstNil
		// result: (NeqFat y x)
		{
			x := v.Args[0]
			y := v.Args[1]
			if !(x.Op == OpConstNil && y.Op != OpConstNil) {
				goto end94c68f7dc30c66ed42e507e01c4e5dc7
			}
			v.Op = OpNeqFat
			v.AuxInt = 0
			v.Aux = nil
			v.resetArgs()
			v.AddArg(y)
			v.AddArg(x)
			return true
		}
		goto end94c68f7dc30c66ed42e507e01c4e5dc7
	end94c68f7dc30c66ed42e507e01c4e5dc7:
		;
		// match: (NeqFat (Load ptr mem) (ConstNil))
		// cond:
		// result: (NeqPtr (Load <TypeUintptr> ptr mem) (ConstPtr <TypeUintptr> [0]))
		{
			if v.Args[0].Op != OpLoad {
				goto endec6168797e39f7543bfb991841a466a7
			}
			ptr := v.Args[0].Args[0]
			mem := v.Args[0].Args[1]
			if v.Args[1].Op != OpConstNil {
				goto endec6168797e39f7543bfb991841a466a7
			}
			v.Op = OpNeqPtr
			v.AuxInt = 0
			v.Aux = nil
			v.resetArgs()
			v0 := v.Block.NewValue0(v.Line, OpLoad, TypeInvalid)
			v0.TypeIndex = v.Block.Func.typeIndex(v.Block.Func.Config.Frontend().TypeUintptr())
			v0.AddArg(ptr)
			v0.AddArg(mem)
			v.AddArg(v0)
			v1 := v.Block.NewValue0(v.Line, OpConstPtr, TypeInvalid)
			v1.TypeIndex = v.Block.Func.typeIndex(v.Block.Func.Config.Frontend().TypeUintptr())
			v1.AuxInt = 0
			v.AddArg(v1)
			return true
		}
		goto endec6168797e39f7543bfb991841a466a7
	endec6168797e39f7543bfb991841a466a7:
		;
	case OpPtrIndex:
		// match: (PtrIndex <t> ptr idx)
		// cond:
		// result: (AddPtr ptr (MulPtr <TypeUintptr> idx (ConstPtr <TypeUintptr> [t.Elem().Size()])))
		{
			t := v.Type()
			ptr := v.Args[0]
			idx := v.Args[1]
			v.Op = OpAddPtr
			v.AuxInt = 0
			v.Aux = nil
			v.resetArgs()
			v.AddArg(ptr)
			v0 := v.Block.NewValue0(v.Line, OpMulPtr, TypeInvalid)
			v0.TypeIndex = v.Block.Func.typeIndex(v.Block.Func.Config.Frontend().TypeUintptr())
			v0.AddArg(idx)
			v1 := v.Block.NewValue0(v.Line, OpConstPtr, TypeInvalid)
			v1.TypeIndex = v.Block.Func.typeIndex(v.Block.Func.Config.Frontend().TypeUintptr())
			v1.AuxInt = t.Elem().Size()
			v0.AddArg(v1)
			v.AddArg(v0)
			return true
		}
		goto end8fe0b7cf8fe19c005286ad52fceef87e
	end8fe0b7cf8fe19c005286ad52fceef87e:
		;
	case OpSliceCap:
		// match: (SliceCap (Load ptr mem))
		// cond:
		// result: (Load (AddPtr <ptr.Type()> ptr (ConstPtr <TypeUintptr> [config.PtrSize*2])) mem)
		{
			if v.Args[0].Op != OpLoad {
				goto enddfc963c554800f2005aa63a441d2e8f0
			}
			ptr := v.Args[0].Args[0]
			mem := v.Args[0].Args[1]
			v.Op = OpLoad
			v.AuxInt = 0
			v.Aux = nil
			v.resetArgs()
			v0 := v.Block.NewValue0(v.Line, OpAddPtr, TypeInvalid)
			v0.TypeIndex = v.Block.Func.typeIndex(ptr.Type())
			v0.AddArg(ptr)
			v1 := v.Block.NewValue0(v.Line, OpConstPtr, TypeInvalid)
			v1.TypeIndex = v.Block.Func.typeIndex(v.Block.Func.Config.Frontend().TypeUintptr())
			v1.AuxInt = config.PtrSize * 2
			v0.AddArg(v1)
			v.AddArg(v0)
			v.AddArg(mem)
			return true
		}
		goto enddfc963c554800f2005aa63a441d2e8f0
	enddfc963c554800f2005aa63a441d2e8f0:
		;
	case OpSliceLen:
		// match: (SliceLen (Load ptr mem))
		// cond:
		// result: (Load (AddPtr <ptr.Type()> ptr (ConstPtr <TypeUintptr> [config.PtrSize])) mem)
		{
			if v.Args[0].Op != OpLoad {
				goto end3ed028ecd6469668e62563fe3efe9ac8
			}
			ptr := v.Args[0].Args[0]
			mem := v.Args[0].Args[1]
			v.Op = OpLoad
			v.AuxInt = 0
			v.Aux = nil
			v.resetArgs()
			v0 := v.Block.NewValue0(v.Line, OpAddPtr, TypeInvalid)
			v0.TypeIndex = v.Block.Func.typeIndex(ptr.Type())
			v0.AddArg(ptr)
			v1 := v.Block.NewValue0(v.Line, OpConstPtr, TypeInvalid)
			v1.TypeIndex = v.Block.Func.typeIndex(v.Block.Func.Config.Frontend().TypeUintptr())
			v1.AuxInt = config.PtrSize
			v0.AddArg(v1)
			v.AddArg(v0)
			v.AddArg(mem)
			return true
		}
		goto end3ed028ecd6469668e62563fe3efe9ac8
	end3ed028ecd6469668e62563fe3efe9ac8:
		;
	case OpSlicePtr:
		// match: (SlicePtr (Load ptr mem))
		// cond:
		// result: (Load ptr mem)
		{
			if v.Args[0].Op != OpLoad {
				goto end459613b83f95b65729d45c2ed663a153
			}
			ptr := v.Args[0].Args[0]
			mem := v.Args[0].Args[1]
			v.Op = OpLoad
			v.AuxInt = 0
			v.Aux = nil
			v.resetArgs()
			v.AddArg(ptr)
			v.AddArg(mem)
			return true
		}
		goto end459613b83f95b65729d45c2ed663a153
	end459613b83f95b65729d45c2ed663a153:
		;
	case OpStore:
		// match: (Store dst (Load <t> src mem) mem)
		// cond: t.Size() > 8
		// result: (Move [t.Size()] dst src mem)
		{
			dst := v.Args[0]
			if v.Args[1].Op != OpLoad {
				goto end324ffb6d2771808da4267f62c854e9c8
			}
			t := v.Args[1].Type()
			src := v.Args[1].Args[0]
			mem := v.Args[1].Args[1]
			if v.Args[2] != v.Args[1].Args[1] {
				goto end324ffb6d2771808da4267f62c854e9c8
			}
			if !(t.Size() > 8) {
				goto end324ffb6d2771808da4267f62c854e9c8
			}
			v.Op = OpMove
			v.AuxInt = 0
			v.Aux = nil
			v.resetArgs()
			v.AuxInt = t.Size()
			v.AddArg(dst)
			v.AddArg(src)
			v.AddArg(mem)
			return true
		}
		goto end324ffb6d2771808da4267f62c854e9c8
	end324ffb6d2771808da4267f62c854e9c8:
		;
		// match: (Store dst str mem)
		// cond: str.Type().IsString()
		// result: (Store (OffPtr <TypeBytePtr> [config.PtrSize] dst) (StringLen <TypeUintptr> str) (Store <TypeMem> dst (StringPtr <TypeBytePtr> str) mem))
		{
			dst := v.Args[0]
			str := v.Args[1]
			mem := v.Args[2]
			if !(str.Type().IsString()) {
				goto end74b8bd8e1b816dcbf80822caebe240b9
			}
			v.Op = OpStore
			v.AuxInt = 0
			v.Aux = nil
			v.resetArgs()
			v0 := v.Block.NewValue0(v.Line, OpOffPtr, TypeInvalid)
			v0.TypeIndex = v.Block.Func.typeIndex(v.Block.Func.Config.Frontend().TypeBytePtr())
			v0.AuxInt = config.PtrSize
			v0.AddArg(dst)
			v.AddArg(v0)
			v1 := v.Block.NewValue0(v.Line, OpStringLen, TypeInvalid)
			v1.TypeIndex = v.Block.Func.typeIndex(v.Block.Func.Config.Frontend().TypeUintptr())
			v1.AddArg(str)
			v.AddArg(v1)
			v2 := v.Block.NewValue0(v.Line, OpStore, TypeInvalid)
			v2.TypeIndex = typeMemIndex
			v2.AddArg(dst)
			v3 := v.Block.NewValue0(v.Line, OpStringPtr, TypeInvalid)
			v3.TypeIndex = v.Block.Func.typeIndex(v.Block.Func.Config.Frontend().TypeBytePtr())
			v3.AddArg(str)
			v2.AddArg(v3)
			v2.AddArg(mem)
			v.AddArg(v2)
			return true
		}
		goto end74b8bd8e1b816dcbf80822caebe240b9
	end74b8bd8e1b816dcbf80822caebe240b9:
		;
	case OpStringLen:
		// match: (StringLen (StringMake _ len))
		// cond:
		// result: len
		{
			if v.Args[0].Op != OpStringMake {
				goto end0d922460b7e5ca88324034f4bd6c027c
			}
			len := v.Args[0].Args[1]
			v.Op = len.Op
			v.AuxInt = len.AuxInt
			v.Aux = len.Aux
			v.resetArgs()
			v.AddArgs(len.Args...)
			return true
		}
		goto end0d922460b7e5ca88324034f4bd6c027c
	end0d922460b7e5ca88324034f4bd6c027c:
		;
	case OpStringPtr:
		// match: (StringPtr (StringMake ptr _))
		// cond:
		// result: ptr
		{
			if v.Args[0].Op != OpStringMake {
				goto end061edc5d85c73ad909089af2556d9380
			}
			ptr := v.Args[0].Args[0]
			v.Op = ptr.Op
			v.AuxInt = ptr.AuxInt
			v.Aux = ptr.Aux
			v.resetArgs()
			v.AddArgs(ptr.Args...)
			return true
		}
		goto end061edc5d85c73ad909089af2556d9380
	end061edc5d85c73ad909089af2556d9380:
		;
	case OpStructSelect:
		// match: (StructSelect [idx] (Load ptr mem))
		// cond:
		// result: (Load (OffPtr <v.Type().PtrTo()> [idx] ptr) mem)
		{
			idx := v.AuxInt
			if v.Args[0].Op != OpLoad {
				goto endd4c92247c08eb96b6c2e5c2023c815cb
			}
			ptr := v.Args[0].Args[0]
			mem := v.Args[0].Args[1]
			v.Op = OpLoad
			v.AuxInt = 0
			v.Aux = nil
			v.resetArgs()
			v0 := v.Block.NewValue0(v.Line, OpOffPtr, TypeInvalid)
			v0.TypeIndex = v.Block.Func.typeIndex(v.Type().PtrTo())
			v0.AuxInt = idx
			v0.AddArg(ptr)
			v.AddArg(v0)
			v.AddArg(mem)
			return true
		}
		goto endd4c92247c08eb96b6c2e5c2023c815cb
	endd4c92247c08eb96b6c2e5c2023c815cb:
	}
	return false
}
func rewriteBlockgeneric(b *Block) bool {
	switch b.Kind {
	case BlockIf:
		// match: (If (Not cond) yes no)
		// cond:
		// result: (If cond no yes)
		{
			v := b.Control
			if v.Op != OpNot {
				goto endebe19c1c3c3bec068cdb2dd29ef57f96
			}
			cond := v.Args[0]
			yes := b.Succs[0]
			no := b.Succs[1]
			b.Kind = BlockIf
			b.Control = cond
			b.Succs[0] = no
			b.Succs[1] = yes
			return true
		}
		goto endebe19c1c3c3bec068cdb2dd29ef57f96
	endebe19c1c3c3bec068cdb2dd29ef57f96:
		;
		// match: (If (ConstBool {c}) yes no)
		// cond: c.(bool)
		// result: (Plain nil yes)
		{
			v := b.Control
			if v.Op != OpConstBool {
				goto end9ff0273f9b1657f4afc287562ca889f0
			}
			c := v.Aux
			yes := b.Succs[0]
			no := b.Succs[1]
			if !(c.(bool)) {
				goto end9ff0273f9b1657f4afc287562ca889f0
			}
			v.Block.Func.removePredecessor(b, no)
			b.Kind = BlockPlain
			b.Control = nil
			b.Succs = b.Succs[:1]
			b.Succs[0] = yes
			return true
		}
		goto end9ff0273f9b1657f4afc287562ca889f0
	end9ff0273f9b1657f4afc287562ca889f0:
		;
		// match: (If (ConstBool {c}) yes no)
		// cond: !c.(bool)
		// result: (Plain nil no)
		{
			v := b.Control
			if v.Op != OpConstBool {
				goto endf401a4553c3c7c6bed64801da7bba076
			}
			c := v.Aux
			yes := b.Succs[0]
			no := b.Succs[1]
			if !(!c.(bool)) {
				goto endf401a4553c3c7c6bed64801da7bba076
			}
			v.Block.Func.removePredecessor(b, yes)
			b.Kind = BlockPlain
			b.Control = nil
			b.Succs = b.Succs[:1]
			b.Succs[0] = no
			return true
		}
		goto endf401a4553c3c7c6bed64801da7bba076
	endf401a4553c3c7c6bed64801da7bba076:
	}
	return false
}
