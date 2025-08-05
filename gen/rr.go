package main

import (
	"text/template"
)

// RL r: rotate left
// RLC: rotate left circular and RRC: rotate right circular
// RL (HL): rotate left, indirect
// RR r and RR (HL): rotate right, also

var templRr = template.Must(tmpl.New("rr").
	Funcs(template.FuncMap{
		"get": get,
		"set": set,
	}).
	Parse(`
res, flags := rotate({{get .Name .Immediate}}, 1, cpu.F, false)
{{set .Name .Immediate "res"}}
cpu.F = flags

cpu.cycles += {{.CycleCount}}
`))

type templDataRr struct {
	Name       string // what register to check, or location
	Immediate  bool
	CycleCount int
}

func (o Opcode) DataRr() templDataRr {
	return templDataRr{
		Name:       o.Operands.First().Name,
		Immediate:  o.Operands.First().Immediate,
		CycleCount: o.CycleCount(),
	}
}
