package main

import (
	"text/template"
)

var templRra = template.Must(tmpl.New("rra").
	Funcs(template.FuncMap{
		"get": get,
		"set": set,
	}).
	Parse(`
cpu.A, cpu.F = rotate(cpu.A, 1, cpu.F, false)
cpu.cycles += {{.CycleCount}}
`))

type templDataRra struct {
	CycleCount int
}

func (o Opcode) DataRra() templDataRra {
	return templDataRra{
		CycleCount: o.CycleCount(),
	}
}
