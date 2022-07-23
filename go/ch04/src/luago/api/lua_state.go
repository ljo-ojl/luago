package state

type LuaType = int

type LuaState interface {
	GetTop() int
}