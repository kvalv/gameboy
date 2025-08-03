package main

import "text/template"

var templPop = template.Must(tmpl.New("pop").
	Funcs(template.FuncMap{
		"get": get,
		"set": set,
	}).
	Parse(`
value := cpu.PopStack()
{{set .Name true "value"}}
cpu.IncProgramCounter("pop")
cpu.cycles += {{.CycleCount}}
`))

type templDataPop struct {
	Name       string // name of register
	CycleCount int
}

func (o Opcode) DataPop() templDataPop {
	return templDataPop{
		Name:       o.Operands.First().Name,
		CycleCount: o.CycleCount(),
	}
}
