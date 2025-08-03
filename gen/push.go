package main

import "text/template"

var templPush = template.Must(tmpl.New("push").
	Funcs(template.FuncMap{
		"get": get,
		"set": set,
	}).
	Parse(`
cpu.PushStack({{get .Name true}})
cpu.cycles += {{.CycleCount}}
`))

type templDataPush struct {
	Name       string // name of register
	CycleCount int
}

func (o Opcode) DataPush() templDataPush {
	return templDataPush{
		Name:       o.Operands.First().Name,
		CycleCount: o.CycleCount(),
	}
}
