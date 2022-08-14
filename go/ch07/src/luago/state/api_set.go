package state

func (self *luaState) SetTable(idx int) {
	t := self.stack.get(idx)
	v := self.stack.pop()
	k := self.stack.pop()
	self.setTable(t, k, v)
}

func (self *luaState) setTable(t, k, v luaValue) {
	if tb1, ok := t.(*luaTable); ok {
		tb1.put(k, v)
		return
	}
	panic("not a table")
}

func (self *luaState) SetField(idx int, k string) {
	t := self.stack.get(idx)
	v := self.stack.pop()
	self.setTable(t, k, v)
}

func (self *luaState) setI(idx int, i int64)  {
	t := self.get(idx)
	v := self.stack.pop()
	self.setTable(t, i, v)
}

