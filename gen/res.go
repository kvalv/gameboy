package main

import "text/template"

// Reset the bit b of the 8-bit register r to 0
var templRes = template.Must(tmpl.New("res").
	Funcs(template.FuncMap{
		"get": get,
		"set": set,
		"pc":  pc,
	}).
	Parse(`
res := {{get .Reg .Immediate}} & ^uint8(1 << {{get .Num true}})
{{set .Reg .Immediate "res"}}
cpu.Cycles += {{.CycleCount}}
`))

type templDataRes struct {
	Num        string
	Immediate  bool
	Reg        string
	CycleCount int // number of cycles for this instruction
}

func (o Opcode) DataRes() templDataRes {
	return templDataRes{
		Num:        o.Operands.First().Name,
		Immediate:  o.Operands.Second().Immediate,
		Reg:        o.Operands.Second().Name,
		CycleCount: o.CycleCount(),
	}
}
