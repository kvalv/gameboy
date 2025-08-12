package main

import "text/template"

var templNop = template.Must(tmpl.New("nop").
	Funcs(template.FuncMap{}).
	Parse(`
cpu.cycles += 4
`))

type templDataNop struct{}

func (o Opcode) DataNop() templDataNop {
	return templDataNop{}
}
