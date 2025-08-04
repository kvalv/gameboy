package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func loadOpcodes(file string) (main, extended []Opcode, err error) {

	b, err := os.ReadFile(file)
	if err != nil {
		err = fmt.Errorf("failed to read Opcodes.json: %w", err)
		return
	}

	var data struct {
		Unprefixed map[string]Opcode `json:"unprefixed"`
		Prefixed   map[string]Opcode `json:"cbprefixed"`
	}
	if err = json.Unmarshal(b, &data); err != nil {
		err = fmt.Errorf("failed to unmarshal Opcodes.json: %w", err)
		return
	}

	for code, v := range data.Unprefixed {
		v.ID = fmt.Sprintf("%s_%s", v.Mnemonic, code[2:])
		i64, _ := strconv.ParseInt(code, 0, 64)
		v.Code = int(i64)

		data.Unprefixed[code] = v
	}
	for code, v := range data.Prefixed {
		v.ID = fmt.Sprintf("%s_%s", v.Mnemonic, code[2:])
		i64, _ := strconv.ParseInt(code, 0, 64)
		v.Code = int(i64)

		data.Prefixed[code] = v
	}

	// data.Prefixed = map[string]Opcode{}

	for _, v := range data.Unprefixed {
		main = append(main, v)
	}
	for _, v := range data.Prefixed {
		extended = append(extended, v)
	}

	return main, extended, nil
}

type Operand struct {
	Name      string `json:"name"`
	Immediate bool   `json:"immediate"`
	Bytes     int    `json:"bytes"`

	Increment bool `json:"increment"` // if true, then this operand is incremented after use
	Decrement bool `json:"decrement"` // if true, then this operand is decremented after use
}

func (o Operand) String() string {
	var s strings.Builder
	if !o.Immediate {
		s.WriteString("(")
	}
	s.WriteString(o.Name)
	if o.Increment {
		s.WriteString("+")
	}
	if o.Decrement {
		s.WriteString("-")
	}
	if !o.Immediate {
		s.WriteString(")")
	}
	return s.String()
}

type Operands []Operand

func (ops Operands) First() Operand {
	if len(ops) == 0 {
		return Operand{}
	}
	return ops[0]
}
func (ops Operands) Second() Operand { return ops[1] }

type Opcode struct {
	Mnemonic  string            `json:"mnemonic"`
	Bytes     int               `json:"bytes"`
	Cycles    []int             `json:"cycles"`
	Operands  Operands          `json:"operands"`
	Immediate bool              `json:"immediate"`
	Flags     map[string]string `json:"flags"`
	ID        string
	Code      int
}

func (o Opcode) String() string {
	if o.Mnemonic == "STOP" {
		return "STOP"
	}
	var parts []string
	for _, arg := range o.Operands {
		parts = append(parts, arg.String())
	}
	return strings.TrimSpace(fmt.Sprintf("%s %s", o.Mnemonic, strings.Join(parts, ",")))
}
func (o Opcode) Desc() string {
	var parts []string
	for _, arg := range o.Operands {
		parts = append(parts, arg.String())
	}
	return fmt.Sprintf("%s %s    code=%#02x", o.Mnemonic, strings.Join(parts, ","), o.Code)
}

func (o Opcode) CycleCount() int {
	var count int
	for _, c := range o.Cycles {
		count += c
	}
	return count
}
