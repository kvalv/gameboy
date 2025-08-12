package main

import "text/template"

var templCp = template.Must(tmpl.New("cp").
	Funcs(template.FuncMap{
		"get": get,
		"set": set,
		"pc":  pc,
	}).
	Parse(`
_, flags := sub({{get "A" true}}, {{get .Name .Immediate}})
{{pc .Name}}
cpu.F = flags
cpu.cycles += {{.CycleCount}}
`))

type templDataCp struct {
	Name       string
	Immediate  bool
	CycleCount int // number of cycles for this instruction
}

func (o Opcode) DataCp() templDataCp {
	return templDataCp{
		Name:       o.Operands.Second().Name,
		Immediate:  o.Operands.Second().Immediate,
		CycleCount: o.CycleCount(),
	}
}
