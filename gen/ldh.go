package main

import "text/template"

var templLdh = template.Must(tmpl.New("ldh").
	Funcs(template.FuncMap{
		"get": get,
		"set": set,
		"pc":  pc,
	}).
	Parse(`
pc0 := cpu.PC
value := {{get .Rhs .RhsImmediate}}
{{set .Lhs .LhsImmediate "value"}}
{{pc .Rhs}}
cpu.PC = pc0 + 1
cpu.cycles += {{.CycleCount}}
`))

type templDataLdh struct {
	Lhs          string
	LhsImmediate bool
	Rhs          string // name of register
	RhsImmediate bool

	CycleCount int // number of cycles for this instruction
}

func (o Opcode) DataLdh() templDataLdh {
	return templDataLdh{
		Lhs:          o.Operands.First().Name,
		LhsImmediate: o.Operands.First().Immediate,
		Rhs:          o.Operands.Second().Name,
		RhsImmediate: o.Operands.Second().Immediate,
		CycleCount:   o.CycleCount(),
	}
}
