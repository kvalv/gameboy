package main

import "text/template"

var templPrefix = template.Must(tmpl.New("prefix").
	Funcs(template.FuncMap{}).
	Parse(`
cpu.prefix = true
cpu.Cycles += 4
`))

type templDataPrefix struct{}

func (o Opcode) DataPrefix() templDataPrefix {
	return templDataPrefix{}
}
