package main

import "text/template"

var templSub = template.Must(tmpl.New("sub").
	Funcs(template.FuncMap{
		"get": get,
		"set": set,
		"pc":  pc,
	}).
	Parse(`
res, flags := sub({{get "A" true}}, {{get .Name .Immediate}})
{{set "A" true "res"}}
{{pc .Name}}
cpu.F = flags
cpu.cycles += {{.CycleCount}}
`))

type templDataSub struct {
	Name       string
	Immediate  bool
	CycleCount int // number of cycles for this instruction
}

func (o Opcode) DataSub() templDataSub {
	return templDataSub{
		Name:       o.Operands.Second().Name,
		Immediate:  o.Operands.Second().Immediate,
		CycleCount: o.CycleCount(),
	}
}
