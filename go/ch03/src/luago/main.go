package main

import "fmt"
import "io/ioutil"
import "os"
import "luago/binchunk"
import . "luago/vm"

func main()  {
	if len(os.Args) > 1 {
		data, err := ioutil.ReadFile(os.Args[1])
		if err != nil { panic(err) }
		proto := binchunk.Undump(data)
		list(proto)
	}
}

func list(f *binchunk.Prototype)  {
	printHeader(f)
	printCode(f)
	printDetail(f)
	for _, p := range f.Protos {
		list(p)
	}
}

func printHeader(f *binchunk.Prototype) {
	funcType := "main"
	if f.LineDefined > 0 { funcType = "function" }

	varargFlag := ""
	if f.IsVararg > 0 { varargFlag = "+" }

	fmt.Printf("\n%s <%s:%d,%d> (%d instructions)\n", funcType, f.Source, f.LineDefined, f.LastLineDefined, len(f.Code))
	fmt.Printf("%d%s", f.NumParams, varargFlag)
	fmt.Printf("%d locals \n", len(f.LocVars))
}


func printCode(f *binchunk.Prototype) {
	for pc, c := range f.Code {
		line := "-"
		if len(f.LineInfo) > 0 {
			line = fmt.Sprintf("%d", f.LineInfo[pc])
		}
		i := Instruction(c)
		fmt.Printf("\t%d\t[%s]\t0x%08X\n", pc+1, line, c)
		printOperand(i)
		fmt.Printf("\n")
	}
}

func printDetail(f* binchunk.Prototype) {
	fmt.Printf("constants (%d): \n", len(f.Constants))
}

func constantToString(k interface{}) string {
	switch k.(type) {
	case nil: 		return "nil"
	case bool: 		return fmt.Sprintf("%t", k)
	case float64: 	return fmt.Sprintf("%g", k)
	case int64: 	return fmt.Sprintf("%d", k)
	case string: 	return fmt.Sprintf("%q", k)
	default: return "?"
	}
}

func UpvalueName(f *binchunk.Prototype, idx int) string {
	if len(f.UpvalueNames) > 0 {
		return f.UpvalueNames[idx]
	}
	return "-"
}  

