package main

import "text/template"

var templInc = template.Must(tmpl.New("inc").
	Funcs(template.FuncMap{
		"reg":    getRegister,
		"setreg": setRegister,
	}).
	Parse(`
{{if .Immediate}}
	res, flags := add({{reg .Name}}, 0x01)
	{{setreg .Name "res"}}
	cpu.F = flags
{{else}}
	// Increments data at the absolute address specified by the register
	var val uint8
	cpu.load({{reg .Name}}, &val)
	next, flags := cpu.Add(val, 0x01)
	cpu.write({{reg .Name}}, next)
	cpu.F = FlagRegister(flags)
{{end}}
`))

type templDataInc struct {
	Name      string // name of register for what to add
	Instr16   bool
	Immediate bool // if not true, we require a load
}

func (o Op) DataInc() templDataInc {
	return templDataInc{
		Name:      o.Operands.First().Name,
		Immediate: o.Operands.First().Immediate,
		Instr16:   o.Operands.First().Is16Bit(),
	}
}

