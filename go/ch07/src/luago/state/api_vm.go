package state

func (self *luaState) PC() int {
	return self.pc
}

func (self *luaState) AddPC(n int) {
	self.pc += 1
}

func (self *luaState) Fetch() uint32 {
	i := self.proto.Code[self.pc]
	self.pc++
	return i
}

func (self *luaState) GetConst(idx int) {
	c := self.proto.Constants[idx]
	self.stack.push(c)
}

func (self *luaState) GetRk(rk int) {
	if rk > 0xFF {
		self.GetConst(rk & 0xFF) 
	} else {
		self.PushValue(rk + 1)
	}
}

