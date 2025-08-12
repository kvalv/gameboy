package main

import (
	"text/template"
)

var templRla = template.Must(tmpl.New("rla").
	Funcs(template.FuncMap{
		"get": get,
		"set": set,
	}).
	Parse(`
cpu.A, cpu.F = rotate(cpu.A, 0, cpu.F, false)
// RLA always sets the zero flag to 0 without looking at the resulting value of the calculation.
cpu.F &= ^FLAGZ
cpu.cycles += {{.CycleCount}}
`))

type templDataRla struct {
	CycleCount int
}

func (o Opcode) DataRla() templDataRla {
	return templDataRla{
		CycleCount: o.CycleCount(),
	}
}
