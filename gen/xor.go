package main

import "text/template"

var templXor = template.Must(tmpl.New("xor").
	Funcs(template.FuncMap{
		"get": get,
		"set": set,
		"pc":  pc,
	}).
	Parse(`
res := {{get "A" true}} ^ {{get .Name .Immediate}}
{{pc .Name}}
var flags Flags
if res == 0 {
	flags |= FLAGZ
}
cpu.F = FlagRegister(flags)
{{set "A" true "res"}}
cpu.cycles += {{.CycleCount}}
`))

type templDataXor struct {
	Name       string
	Immediate  bool
	CycleCount int // number of cycles for this instruction
}

func (o Opcode) DataXor() templDataXor {
	return templDataXor{
		Name:       o.Operands.Second().Name,
		Immediate:  o.Operands.Second().Immediate,
		CycleCount: o.CycleCount(),
	}
}
