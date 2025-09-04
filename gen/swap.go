package main

import "text/template"

var templSwap = template.Must(tmpl.New("swap").
	Funcs(template.FuncMap{
		"get": get,
		"set": set,
	}).
	Parse(`
v := {{get .Name .Immediate}}
res := (v >> 4) | (v << 4)
var flags Flags
if res == 0 {
	flags |= FLAGC
}
cpu.F = flags
{{set .Name .Immediate "res"}}
cpu.Cycles += {{.CycleCount}}
`))

type templDataSwap struct {
	Name       string // name of register for what to add
	Immediate  bool   // is it an immediate value?
	CycleCount int    // number of cycles for this instruction
}

func (o Opcode) DataSwap() templDataSwap {
	return templDataSwap{
		Name:       o.Operands.First().Name,
		Immediate:  o.Operands.First().Immediate,
		CycleCount: o.CycleCount(),
	}
}
