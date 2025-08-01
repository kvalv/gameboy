package main

import "text/template"

var templDec = template.Must(tmpl.New("dec").
	Funcs(template.FuncMap{
		"reg":    getRegister,
		"setreg": setRegister,
	}).
	Parse(`
{{if .Immediate}}
	res, flags := sub({{reg .Name}}, 0x01)
	{{setreg .Name "res"}}
	cpu.F = flags
{{else}}
	var val uint8
	cpu.load({{reg .Name}}, &val)
	next, flags := sub(val, 0x01)
	cpu.WriteMemory({{reg .Name}}, next)
	cpu.F = flags
{{end}}
`))

type templDataDec struct {
	Name      string // name of register for what to add
	Immediate bool   // if not true, we require a load
}

func (o Op) DataDec() templDataDec {
	return templDataDec{
		Name:      o.Operands.First().Name,
		Immediate: o.Operands.First().Immediate,
	}
}
