package main

import (
	"text/template"
)

// RL r: rotate left
// RLC: rotate left circular and RRC: rotate right circular
// RL (HL): rotate left, indirect
// RR r and RR (HL): rotate right, also

var templSrl = template.Must(tmpl.New("srl").
	Funcs(template.FuncMap{
		"get": get,
		"set": set,
	}).
	Parse(`
data := {{get .Name .Immediate}}
b0 := bit(data, 0)
result := data >> 1
var flags Flags
if result == 0 {
	flags |= FLAGZ
}
if b0 == 1 {
	flags |= FLAGC
}
cpu.F = flags
{{set .Name .Immediate "result"}}
cpu.Cycles += {{.CycleCount}}
`))

type templDataSrl struct {
	Name       string // what register to check, or location
	Immediate  bool
	CycleCount int
}

func (o Opcode) DataSrl() templDataSrl {
	return templDataSrl{
		Name:       o.Operands.First().Name,
		Immediate:  o.Operands.First().Immediate,
		CycleCount: o.CycleCount(),
	}
}
