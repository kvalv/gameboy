package main

import (
	"text/template"
)

var templRlca = template.Must(tmpl.New("rlca").
	Funcs(template.FuncMap{
		"get": get,
		"set": set,
	}).
	Parse(`
cpu.A, cpu.F = rotate(cpu.A, 0, cpu.F, true)
cpu.cycles += {{.CycleCount}}
`))

type templDataRlca struct {
	CycleCount int
}

func (o Opcode) DataRlca() templDataRlca {
	return templDataRlca{
		CycleCount: o.CycleCount(),
	}
}
