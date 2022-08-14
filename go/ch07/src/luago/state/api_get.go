package state

import . "luago/api"

func (self *luaState) CreateTable(nArr, nRec) {
	t := newLuaTable(nArr, nRec) 
	self.stack.push(t)
}

func (self *luaState) NewTable {
	self.CreateTable(0, 0)
}

func (self *luaState) GetTable(idx int) LuaType {
	t := self.stack.get(idx)
	k := self.stack.pop()
	return self.getTable(t, k)
}

func (self *luaState) getTable(t, k luaValue) LuaType {
	if tb1, ok := t.(*luaTable); ok {
		v := tb1.get(k)
		self.stack.push(v)
		return typeOf(v)
	}
	panic("not a table")
}

func (self *luaState) GetField(idx int, k string) LuaType {
	self.PushString(k)
	return self.GetTable(idx)
}

func (self *luaState) GetI(idx int, i int64) LuaType {
	t := self.stack.get(idx)
	return self.getTable(t, i)
}