package main

import (
	"fmt"
	"strings"
)

func getRegister(name string) (reg string) {
	defer func() {
		reg = "cpu." + reg // add accessor
	}()
	switch name {
	case "A", "F", "B", "C", "D", "E", "H", "L":
		// direct access
		reg = name
	case "PC", "SP":
		// direct access
		reg = name
	case "HL", "BC", "DE":
		// via method
		reg = name + "()"
	}
	return
}

// generates the code that fetches a value as an expression,
// e.g. `cpu.load(cpu.HL)` or `cpu.A`.
func get(name string, immediate bool) string {
	if !immediate {
		return fmt.Sprintf("cpu.loadU8(%s)", get(name, true))
	}
	switch name {
	case "A", "F", "B", "C", "D", "E", "H", "L":
		// direct access
		return fmt.Sprintf("cpu.%s", name)
	case "PC", "SP":
		// direct access
		return fmt.Sprintf("cpu.%s", name)
	case "HL", "BC", "DE":
		// via method
		return fmt.Sprintf("cpu.%s()", name)
	case "n16":
		// read the next 16-bit word from PC
		return "cpu.readStackU16()"
	case "n8":
		return "cpu.readStackU8()"
	case "a16":
		// TODO: should be signed integer here
		return "cpu.readStackU16()"
	case "a8":
		// TODO: should be signed integer here
		return "cpu.readStackU8()"

	}
	return fmt.Sprintf("// TODO - get not implemented %s", name)
}

// generates the code that writes, either to a register, or to memory
// e.g. `cpu.write(cpu.A, data)` or `cpu.A = cpu.B`
func set(name string, immediate bool, varname string) string {
	if !immediate {
		switch name {
		case "a16":
			var s strings.Builder
			// fmt.Fprintf(&s, "msb, lsb := split(0x1122)\n")
			fmt.Fprintf(&s, "cpu.write(%s, %s)\n", get(name, true), varname)
			return s.String()
		default:
			return fmt.Sprintf("cpu.write(%s, %s)", get(name, true), varname)
		}
	}
	switch name {
	case "A", "C", "E", "L", "B", "D", "H", "SP":
		return fmt.Sprintf("cpu.%s = %s", name, varname)
	case "BC", "DE", "HL":
		r1 := string(name[0])
		r2 := string(name[1])
		return fmt.Sprintf("cpu.%s, cpu.%s = splitU16(%s)", r1, r2, varname)
	default:
		return fmt.Sprintf("// todo: setreg... %s %s", name, varname)
	}
}

// inc bc, de, hl, sp, b, d, h, (hl), c, e, l, a
// writes a line of code that sets the register to variable name
// varname is the variable name holding the value
func setRegister(name string, varname string) string {
	if name == "a16" {
		return "// TODO: set signed integer..."
	}

	switch name {
	case "A", "C", "E", "L", "B", "D", "H", "SP":
		return fmt.Sprintf("cpu.%s = %s", name, varname)
	case "BC", "DE", "HL":
		r1 := string(name[0])
		r2 := string(name[1])
		return fmt.Sprintf("cpu.%s, cpu.%s = splitU16(%s)", r1, r2, varname)
	default:
		return fmt.Sprintf("// todo: setreg... %s %s", name, varname)
	}
}
