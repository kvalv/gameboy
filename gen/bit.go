package main

import (
	"fmt"
	"strconv"
	"text/template"
)

var templBit = template.Must(tmpl.New("bit").
	Funcs(template.FuncMap{
		"get": get,
		"set": set,
	}).
	Parse(`
value := {{get .Name .Immediate}}
var flags Flags
if (value & (1 << {{.Num}})) == 0 {
	flags |= FLAGZ
}
flags |= FLAGH
cpu.F = FlagRegister(flags)

cpu.cycles += {{.CycleCount}}
`))

type templDataBit struct {
	Num        int    // 0, 1, ...
	Name       string // what register to check, or location
	Immediate  bool
	CycleCount int
}

func (o Opcode) DataBit() templDataBit {
	num, err := strconv.Atoi(o.Operands.First().Name)
	if err != nil {
		panic(fmt.Sprintf("failed to convert %q to number", o.Operands.First().Name))
	}

	return templDataBit{
		Num:        num,
		Name:       o.Operands.Second().Name,
		Immediate:  o.Operands.Second().Immediate,
		CycleCount: o.CycleCount(),
	}
}
