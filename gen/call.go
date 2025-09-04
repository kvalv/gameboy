package main

import (
	"text/template"
)

var templCall = template.Must(tmpl.New("call").
	Funcs(template.FuncMap{
		"get":  get,
		"set":  set,
		"cond": cond,
		"pc":   pc,
	}).
	Parse(`
lsb := cpu.readU8(cpu.PC)
cpu.IncProgramCounter()
msb := cpu.readU8(cpu.PC)
cpu.IncProgramCounter()
nn := concatU16(msb, lsb)
if {{cond .Predicate }} {
	cpu.PushStack(cpu.PC)
	cpu.PC = nn
	cpu.Cycles += 24
}  else {
	cpu.Cycles += 12
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
