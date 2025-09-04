package main

import (
	"text/template"
)

// RL r: rotate left
// RLC: rotate left circular and RRC: rotate right circular
// RL (HL): rotate left, indirect
// RR r and RR (HL): rotate right, also

var templRl = template.Must(tmpl.New("rl").
	Funcs(template.FuncMap{
		"get": get,
		"set": set,
	}).
	Parse(`
res, flags := rotate({{get .Name .Immediate}}, 0, cpu.F, false)
{{set .Name .Immediate "res"}}
cpu.F = flags

cpu.Cycles += {{.CycleCount}}
`))

type templDataRl struct {
	Name       string // what register to check, or location
	Immediate  bool
	CycleCount int
}

func (o Opcode) DataRl() templDataRl {
	return templDataRl{
		Name:       o.Operands.First().Name,
		Immediate:  o.Operands.First().Immediate,
		CycleCount: o.CycleCount(),
	}
}
