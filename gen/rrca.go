package main

import (
	"text/template"
)

var templRrca = template.Must(tmpl.New("rrca").
	Funcs(template.FuncMap{
		"get": get,
		"set": set,
	}).
	Parse(`
cpu.A, cpu.F = rotate(cpu.A, 1, cpu.F, true)
cpu.cycles += {{.CycleCount}}
`))

type templDataRrca struct {
	CycleCount int
}

func (o Opcode) DataRrca() templDataRrca {
	return templDataRrca{
		CycleCount: o.CycleCount(),
	}
}
