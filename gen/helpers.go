package main

import (
	"fmt"
)

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
	case "HL", "BC", "DE", "AF":
		// via method
		return fmt.Sprintf("cpu.%s()", name)
	case "n16":
		return "cpu.readU16(cpu.PC)"
	case "n8":
		return "cpu.readU8(cpu.PC)"
	case "a16":
		// 16-bit address
		return "cpu.readU16(cpu.PC)"
	case "a8":
		// means 8 bit unsigned data, which are added to $FF00 in certain instructions
		return "uint32(cpu.readU8(cpu.PC)) + 0xFF00"
	case "e8":
		// e8: signed 8-bit data
		return "cpu.readI8(cpu.PC)"
	}
	return fmt.Sprintf("// TODO - get not implemented %s", name)
}

// generates the code that writes, either to a register, or to memory
// e.g. `cpu.WriteMemory(cpu.A, data)` or `cpu.A = cpu.B`
func set(name string, immediate bool, varname string) string {
	if !immediate {
		return fmt.Sprintf("cpu.WriteMemory(%s, %s)", get(name, true), varname)
	}
	switch name {
	case "A", "C", "E", "L", "B", "D", "H", "SP":
		return fmt.Sprintf("cpu.%s = %s", name, varname)
	case "BC", "DE", "HL":
		msb := string(name[0])
		lsb := string(name[1])
		return fmt.Sprintf("cpu.%s, cpu.%s = split(%s)", msb, lsb, varname)
	case "AF":
		return fmt.Sprintf("msb, lsb := split(%s)\ncpu.A, cpu.F = msb, FlagRegister(lsb)", varname)
	default:
		return fmt.Sprintf("// todo: set %s %s", name, varname)
	}
}

func cond(pred string) string {
	switch pred {
	case "NZ":
		return "!cpu.F.HasZero()"
	case "Z":
		return "cpu.F.HasZero()"
	case "NC":
		return "!cpu.F.HasCarry()"
	case "C":
		return "cpu.F.HasCarry()"
	default:
		return "true"
	}
}
