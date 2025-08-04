package main

import "text/template"

var templJr = template.Must(tmpl.New("jr").
	Funcs(template.FuncMap{
		"get":         get,
		"cond":        cond,
		"indexOrLast": indexOrLast[int],
	}).
	Parse(`
e := {{get "e8" true}}
if {{cond .Predicate}} {
	cpu.PC, cpu.F = add(cpu.PC, e)
	cpu.cycles += {{ indexOrLast .Cycles 0 }}
} else {
	cpu.cycles += {{ indexOrLast .Cycles 1 }}
}
`))

type templDataJr struct {
	Predicate string
	Cycles    []int
}

func (o Opcode) DataJr() templDataJr {
	return templDataJr{
		Predicate: o.Operands.First().Name,
		Cycles:    o.Cycles,
	}
}
