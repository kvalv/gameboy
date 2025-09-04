package main

import "text/template"

var templRst = template.Must(tmpl.New("rst").
	Funcs(template.FuncMap{
		"get": get,
	}).
	Parse(`
n := {{get .Name true}}
cpu.PushStack(cpu.PC)
cpu.PC = concatU16(0x00, n)
cpu.Cycles += {{.Cycles}}
`))

type templDataRst struct {
	Name   string
	Cycles int
}

func (o Opcode) DataRst() templDataRst {
	return templDataRst{
		Name:   o.Operands.First().Name,
		Cycles: o.CycleCount(),
	}
}
