package main

import "text/template"

var templInc = template.Must(tmpl.New("inc").
	Funcs(template.FuncMap{
		"get": get,
		"set": set,
	}).
	Parse(`
	res, flags := add({{get .Name .Immediate}}, 0x01)
	cpu.F = flags
	{{set .Name .Immediate "res"}}
	cpu.IncProgramCounter()
`))

type templDataInc struct {
	Name      string // name of register for what to add
	Immediate bool   // if not true, we require a load
}

func (o Opcode) DataInc() templDataInc {
	return templDataInc{
		Name:      o.Operands.First().Name,
		Immediate: o.Operands.First().Immediate,
	}
}
