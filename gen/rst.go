package main

import "text/template"

var templRst = template.Must(tmpl.New("rst").
	Funcs(template.FuncMap{
		"get": get,
	}).
	Parse(`
n := {{get .Name true}}
cpu.PushStack(cpu.PC)
cpu.PC = concatU16(0x00, n)
`))

type templDataRst struct {
	Name string
}

func (o Opcode) DataRst() templDataRst {
	return templDataRst{
		Name: o.Operands.First().Name,
	}
}
