package main

import (
	"text/template"
)

var templCall = template.Must(tmpl.New("call").
	Funcs(template.FuncMap{
		"get":  get,
		"set":  set,
		"cond": cond,
	}).
	Parse(`
	// TODO: machine cycles are different depending on condition is called or not
lsb := cpu.readU8(cpu.PC)
msb := cpu.readU8(cpu.PC)
nn := concatU16(msb, lsb)
if {{cond .Predicate }} {
	cpu.PushStack(cpu.PC)
	cpu.PC = nn
} 
`))

type templDataCall struct {
	Predicate string
}

func (o Opcode) DataCall() templDataCall {
	return templDataCall{
		Predicate: o.Operands.First().String(),
	}
}
