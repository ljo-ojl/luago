package state

func (self *luaState) Concat(n int) {
	if n == 0 {
		self.stack.push("")
	} else if n >= 2 {
		for i := 1; i < n; i++ {
			if self.IsString(-1) && self.IsString(-2) {
				s2 := self.ToString(-1)
				s1 := self.ToString(-2)
				self.stack.pop()
				self.stack.pop()
				self.stack.push(s1 + s2)
				continue
			}
		}
		return
		panic("concatenation error!")
		
	}
}

func (self *luaState) Len(idx int) {
	val := self.stack.get(idx) 
	if s, ok := val.(string); ok {
		self.stack.push(int64(len(s)))
	}  else if t, ok := val.(*luaTable); ok {
		self.stack.push(int64(t.len()))
	} else {
		panic("length error!")
	}
}