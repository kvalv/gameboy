package main

import "text/template"

var templCpl = template.Must(tmpl.New("cpl").
	Funcs(template.FuncMap{}).
	Parse(`
cpu.A = ^cpu.A
cpu.F = FLAGH | FLAGN
cpu.Cycles += {{.Cycles}}
`))

type templDataCpl struct {
	Cycles int
}

func (o Opcode) DataCpl() templDataCpl {
	return templDataCpl{
		Cycles: o.CycleCount(),
	}
}
