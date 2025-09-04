package main

import "text/template"

var templDi = template.Must(tmpl.New("di").
	Funcs(template.FuncMap{}).
	Parse(`
cpu.ime = false
cpu.Cycles += {{.Cycles}}
`))

type templDataDi struct {
	Cycles int
}

func (o Opcode) DataDi() templDataDi {
	return templDataDi{
		Cycles: o.CycleCount(),
	}
}
