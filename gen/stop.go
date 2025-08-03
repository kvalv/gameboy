package main

import "text/template"

var templStop = template.Must(tmpl.New("stop").
	Funcs(template.FuncMap{}).
	Parse(`
cpu.err = ErrNoMoreInstructions
cpu.cycles += 4
`))

type templDataStop struct{}

func (o Opcode) DataStop() templDataStop {
	return templDataStop{}
}
