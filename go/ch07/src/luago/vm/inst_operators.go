package vm

import . "luago/api"

func _binaryArith(i Instruction, vm LuaVM, op ArithOp) {
	a, b, c := i.ABC()
	a += 1

	vm.GetRk(b)
	vm.GetRk(c)
	vm.Arith(op)
	vm.Replace(a)
}

func _unaryArith(i Instruction, vm LuaVM, op ArithOp) {
	a, b, _ := i.ABC()
	a += 1; b += 1

	vm.PushValue(b)
	vm.Arith(op)
	vm.Replace(a)
}

func add(i Instruction, vm LuaVM) { _binaryArith(i, vm, LUA_OPADD) }
func sub(i Instruction, vm LuaVM) { _binaryArith(i, vm, LUA_OPSUB) }
func mul(i Instruction, vm LuaVM) { _binaryArith(i, vm, LUA_OPMUL) }
func mod(i Instruction, vm LuaVM) { _binaryArith(i, vm, LUA_OPMOD) }



func _length(i Instruction, vm LuaVM)  {
	a, b, _ := i.ABC()
	a += 1; b += 1

	vm.Len(b)
	vm.Replace(a)
}

func concat(i Instruction, vm LuaVM) {
	a, b, c := i.ABC()
	a += 1; b += 1; c += 1; 

	n := c - b - 1;
	vm.CheckStack(n)
	for i :=b; i <= c; i++ {
		vm.PushValue(i)
	}
	vm.Concat(n)
	vm.Replace(a)
}

func _compare(i Instruction, vm LuaVM, op CompareOp) {
	a, b, c := i.ABC()

	vm.GetRk(b)
	vm.GetRk(c)
	if vm.Compare(-2, -1, op) != (a != 0) {
		vm.AddPC(1)
	}
	vm.Pop(2)
}

func eq(i Instruction, vm LuaVM) {
	_compare(i, vm, LUA_OPEQ)
}

func lt(i Instruction, vm LuaVM) {
	_compare(i, vm, LUA_OPLT)
}

func le(i Instruction, vm LuaVM) {
	_compare(i, vm LUA_OPLE)
}


func not(i Instruction, vm LuaVM) {
	a, b, _ := i.ABC()
	a += 1; b += 1
	vm.PushBoolean(!vm.ToBoolean(b)) 
	vm.Replace(a) 
}

func testSet(i Instruction, vm LuaVM) {
	a, b, c := i.ABC()
	a += 1; b += 1

	if vm.ToBoolean(b) == (c != 0) {
		vm.Copy(b, a) 
	} else {
		vm.AddPC(1)
	}
}

func test(i Instruction, vm LuaVM) {
	a, _, c := i.ABC()
	a += 1

	if vm.ToBoolean(a) != (c != 0) {
		vm.AddPC(1)
	}
}

func forPrep(i Instruction, vm LuaVM) {
	a, sBx := i.AsBx()
	a += 1

	vm.PushValue(a)
	vm.PushValue(a + 2)
	vm.Arith(LUA_OPSUB)
	vm.Replace(a)
	vm.AddPC(sBx)
}

func forLoop(i Instruction, vm LuaVM) {
	a, sBx := i.AsBx()
	a += 1

	vm.PushValue(a)
	vm.PushValue(a + 2)
	vm.Arith(LUA_OPADD)
	vm.Replace(a)
	
	
	isPositiveStep := vm.ToNumber(a + 2) >= 0 
	if isPositiveStep && vm.Compare(a, a + 1, LUA_OPLE) ||
	!isPositiveStep && vm.Compare(a + 1, a, LUA_OPLE) {
		vm.AddPC(sBx)
		vm.Copy(a, a + 3)
	}
}



