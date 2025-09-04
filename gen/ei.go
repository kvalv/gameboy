package main

import "text/template"

var templEi = template.Must(tmpl.New("ei").
	Funcs(template.FuncMap{}).
	Parse(`
// technically on the next cycle
cpu.ime = true
cpu.Cycles += {{ .Cycles }}
`))

type templDataEi struct {
	Cycles int
}

func (o Opcode) DataEi() templDataEi {
	return templDataEi{
		Cycles: o.CycleCount(),
	}
}
