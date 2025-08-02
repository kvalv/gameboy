package main

import "text/template"

var templAdd = template.Must(tmpl.New("add").
	Funcs(template.FuncMap{
		"get": get,
		"set": set,
	}).
	Parse(`
lhs := {{get .Dst true}}
rhs := {{get .Rhs .RhsImmediate}}
res, flags := add(lhs, rhs)
{{set .Dst true "res"}}
cpu.F = flags
`))

type templDataAdd struct {
	Dst          string
	Rhs          string
	RhsImmediate bool
}

func (o Opcode) DataAdd() templDataAdd {
	return templDataAdd{
		Dst:          o.Operands.First().Name,
		Rhs:          o.Operands.Second().Name,
		RhsImmediate: o.Operands.Second().Immediate,
	}
}
