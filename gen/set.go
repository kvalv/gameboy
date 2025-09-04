package main

import "text/template"

// Setet the bit b of the 8-bit register r to 0
var templSet = template.Must(tmpl.New("set").
	Funcs(template.FuncMap{
		"get": get,
		"set": set,
		"pc":  pc,
	}).
	Parse(`
v := {{get .Reg .Immediate}} | (1 << {{get .Num true}})
{{set .Reg .Immediate "v"}}
cpu.Cycles += {{.CycleCount}}
`))

type templDataSet struct {
	Num        string
	Immediate  bool
	Reg        string
	CycleCount int // number of cycles for this instruction
}

func (o Opcode) DataSet() templDataSet {
	return templDataSet{
		Num:        o.Operands.First().Name,
		Immediate:  o.Operands.Second().Immediate,
		Reg:        o.Operands.Second().Name,
		CycleCount: o.CycleCount(),
	}
}
