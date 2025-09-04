package main

import "text/template"

var templJp = template.Must(tmpl.New("jp").
	Funcs(template.FuncMap{
		"get":         get,
		"cond":        cond,
		"indexOrLast": indexOrLast[int],
		"pc":          pc,
	}).
	Parse(`
nn := {{get "a16" true}}
{{pc "a16"}}
if {{cond .Predicate}} {
	cpu.PC = nn
	cpu.Cycles += {{ indexOrLast .Cycles 0 }}
} else {
	cpu.Cycles += {{ indexOrLast .Cycles 1 }}
}
`))

type templDataJp struct {
	Predicate string
	Cycles    []int
}

func (o Opcode) DataJp() templDataJp {
	return templDataJp{
		Predicate: o.Operands.First().Name,
		Cycles:    o.Cycles,
	}
}
