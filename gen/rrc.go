package main

import (
	"text/template"
)

// RL r: rotate left
// RLC: rotate left circular and RRC: rotate right circular
// RL (HL): rotate left, indirect
// RR r and RR (HL): rotate right, also

var templRrc = template.Must(tmpl.New("rrc").
	Funcs(template.FuncMap{
		"get": get,
		"set": set,
	}).
	Parse(`
res, flags := rotate({{get .Name .Immediate}}, 1, cpu.F, true)
{{set .Name .Immediate "res"}}
cpu.F = flags

cpu.cycles += {{.CycleCount}}
`))

type templDataRrc struct {
	Name       string // what register to check, or location
	Immediate  bool
	CycleCount int
}

func (o Opcode) DataRrc() templDataRrc {
	return templDataRrc{
		Name:       o.Operands.First().Name,
		Immediate:  o.Operands.First().Immediate,
		CycleCount: o.CycleCount(),
	}
}
