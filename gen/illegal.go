package main

import "text/template"

var templIllegal = template.Must(tmpl.New("illegal").
	Funcs(template.FuncMap{}).
	Parse(`
panic("illegal instruction {{.ID}}")
`))

type templDataIllegal struct {
	ID string
}

func (o Opcode) DataIllegal() templDataIllegal {
	return templDataIllegal{
		ID: o.ID,
	}
}
