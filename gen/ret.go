package main

import "text/template"

var templRet = template.Must(tmpl.New("ret").
	Funcs(template.FuncMap{
		"get":         get,
		"set":         set,
		"cond":        cond,
		"indexOrLast": indexOrLast[int],
	}).
	Parse(`
if {{cond .Predicate}} {
	{{set "PC" true "cpu.PopStack()"}}
	cpu.Cycles += {{ indexOrLast .Cycles 0 }}
} else {
	cpu.Cycles += {{ indexOrLast .Cycles 1 }}
}
`))

type templDataRet struct {
	Predicate string
	Cycles    []int
}

func (o Opcode) DataRet() templDataRet {
	return templDataRet{
		Predicate: o.Operands.First().Name,
		Cycles:    o.Cycles,
	}
}
