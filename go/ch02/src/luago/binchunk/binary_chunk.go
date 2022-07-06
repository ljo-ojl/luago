package binchunk

type binaryChunk struct {
	header
	sizeUpvalues byte
	mainFunc *Prototype
}

type header struct {
	signature [4]byte
	version byte
	format byte
	luacData [6]byte
	cintsize byte
	sizetSize byte
	instructionSize byte
	luaIntegerSize byte
	luaNumberSize byte
	luacInt int64
	luacNum float64
}

const {
	LUA_SIGNATURE = "\x1bLua"
	LUAC_VERSION = 0x54
	LUAC_FORMAT = 0
	LUAC_DATA = "\x19\x93\r\n\x1a\n"
	CINT_SIZE = 4
	CSIZET_SIZE = 8
	INSTRUCTION_SIZE = 8
	LUA_INTEGER_SIZE = 8
	LUNAR_NUMBER_SIZE = 8
	LUAC_INT = 0x5678
	LUAC_NUM = 370.5
}

type Prototype struct {
	Source string
	LineDefined uint32
	LastLineDefined uint32
	NumParams byte
	IsVararg byte
	MaxStackSize byte
	Code []uint32
	Constants []interface{}
	Upvalues []Upvalue
	Protos []*Prototype
	LineInfo []uint32
	LocVars []LocVar 
	UpvalueNames []string
}

const {
	TAG_NIL = 0x00
	TAG_BOOLEAN = 0x01
	TAG_NUMBER = 0x03
	TAG_INTEGER = 0x13
	TAG_SHORT_STR = 0X04
	TAG_LONG_STR = 0X14
}

type Upvalue struct {
	Instack byte
	Idx byte
}

type LocVar struct {
	VarName string
	StartPC uint32
	EndPC uint32
}

func Undump(data []byte)  {
	reader := &reader{data}
	reader.checkHeader()
	reader.readByte()
	return reader.readProto("")
}