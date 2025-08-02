package main

import "text/template"

var templLd = template.Must(tmpl.New("ld").
	Funcs(template.FuncMap{
		"get": get,
		"set": set,
	}).
	Parse(`
{{if eq .Code 0xF8}}
	e := {{get "e8" true}}
	res, flags := add({{get "SP" true}}, e)
	{{set "HL" true "res"}}
	cpu.F = flags
{{else}}
	data := {{get .Src .SrcImmediate}}
	
	{{set .Dst .DstImmediate "data"}}

	{{if .PostIncrement -}}
	// post increment
	incr, flags := add({{get .Dst true}}, 0x01)
	{{set .Dst true "incr"}}
	cpu.F = flags
	{{end}}

	{{if .PostDecrement -}}
	// post decrement
	decr, flags := sub({{get .Dst true}}, 0x01)
	{{set .Dst true "decr"}}
	cpu.F = flags
	{{end}}

{{end}}
`))

type templDataLd struct {
	Dst           string // name of register, where data is written
	DstImmediate  bool
	Src           string // name of register, or a constant
	SrcImmediate  bool   // if false, then we need to load from memory
	Code          uint8  // for the 0xF8 instruction, we'll just hardcode it, it's way too complicated
	PostIncrement bool
	PostDecrement bool
}

func (o Op) DataLd() templDataLd {
	return templDataLd{
		Dst:           o.Operands.First().Name,
		DstImmediate:  o.Operands.First().Immediate,
		Src:           o.Operands.Second().Name,
		SrcImmediate:  o.Operands.Second().Immediate,
		Code:          uint8(o.Code),
		PostIncrement: o.Operands.First().Increment,
		PostDecrement: o.Operands.First().Decrement,
	}
}
