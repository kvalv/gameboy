package main

import "text/template"

var templLd = template.Must(tmpl.New("ld").
	Funcs(template.FuncMap{
		"get": get,
		"set": set,
	}).
	Parse(`
{{if eq .Code 0xF8}}
	// TOOD: load: special case
{{else}}
	data := {{get .Src .SrcImmediate}}
	
	{{set .Dst .DstImmediate "data"}}
{{end}}
`))

// either a register, or a constant, or indirect access (mem location)
type UnionThingy struct {
	Register string // name
	Memory   uint16 // location
	Direct   byte   // u8,
}

type templDataLd struct {
	Dst          string // name of register, where data is written
	DstImmediate bool
	Src          string // name of register, or a constant
	SrcImmediate bool   // if false, then we need to load from memory
	Code         uint8  // for the 0xF8 instruction, we'll just hardcode it, it's way too complicated
}

func (o Op) DataLd() templDataLd {
	return templDataLd{
		Dst:          o.Operands.First().Name,
		DstImmediate: o.Operands.First().Immediate,
		Src:          o.Operands.Second().Name,
		SrcImmediate: o.Operands.Second().Immediate,
		Code:         uint8(o.Code),
	}
}
