package main

import "text/template"

var templStop = template.Must(tmpl.New("stop").
	Funcs(template.FuncMap{}).
	Parse(`
cpu.err = ErrNoMoreInstructions
cpu.Cycles += 4
`))

type templDataStop struct{}

func (o Opcode) DataStop() templDataStop {
	return templDataStop{}
}
