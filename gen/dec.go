package main

import "text/template"

var templDec = template.Must(tmpl.New("dec").
	Funcs(template.FuncMap{
		"get": get,
		"set": set,
	}).
	Parse(`
res, flags := sub({{get .Name .Immediate}}, 0x01)
cpu.F = flags
{{set .Name .Immediate "res"}}
cpu.Cycles += {{.CycleCount}}
`))

type templDataDec struct {
	Name       string // name of register for what to add
	Immediate  bool   // if not true, we require a load
	CycleCount int    // number of cycles for this instruction
}

func (o Opcode) DataDec() templDataDec {
	return templDataDec{
		Name:       o.Operands.First().Name,
		Immediate:  o.Operands.First().Immediate,
		CycleCount: o.CycleCount(),
	}
}
