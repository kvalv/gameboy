package main

import "text/template"

var templReti = template.Must(tmpl.New("reti").
	Funcs(template.FuncMap{
		"pc": pc,
	}).
	Parse(`
{{set "PC" true "cpu.PopStack()"}}
{{pc "a16"}}
cpu.ime = true
cpu.Cycles += {{.Cycles}}
`))

type templDataReti struct {
	Cycles int
}

func (o Opcode) DataReti() templDataReti {
	return templDataReti{
		Cycles: o.CycleCount(),
	}
}
