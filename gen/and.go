package main

import "text/template"

var templAnd = template.Must(tmpl.New("and").
	Funcs(template.FuncMap{
		"get": get,
		"set": set,
	}).
	Parse(`
res := {{get "A" true}} & {{get .Name .Immediate}}
var flags Flags
if res == 0 {
	flags |= FLAGZ
}
cpu.F = FlagRegister(flags)
{{set "A" true "res"}}
cpu.Cycles += {{.CycleCount}}
`))

type templDataAnd struct {
	Name       string
	Immediate  bool
	CycleCount int // number of cycles for this instruction
}

func (o Opcode) DataAnd() templDataAnd {
	return templDataAnd{
		Name:       o.Operands.Second().Name,
		Immediate:  o.Operands.Second().Immediate,
		CycleCount: o.CycleCount(),
	}
}
