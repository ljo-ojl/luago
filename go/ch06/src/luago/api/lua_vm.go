package api

type LuaVM interface {
	LuaState
	PC() int
	AddPC(n int) 
	Fetch() uint32
	GetConst(idx int)
	GetRk(rk int)
}