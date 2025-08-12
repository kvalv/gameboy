package main

import (
	"fmt"
	"strconv"
	"strings"
)

// generates the code that fetches a value as an expression,
// e.g. `cpu.load(cpu.HL)` or `cpu.A`.
func get(name string, immediate bool) string {
	if !immediate {
		switch name {
		case "A", "F", "B", "C", "D", "E", "H", "L":
			return fmt.Sprintf("cpu.loadU8(concatU16(0xFF, cpu.%s))", name)
		}
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
		return "concatU16(0xFF, cpu.readU8(cpu.PC))"
	case "e8":
		// e8: signed 8-bit data
		return "cpu.readI8(cpu.PC)"
	}
	if strings.HasPrefix(name, "$") {
		n, err := strconv.ParseInt(name[1:], 16, 8)
		if err != nil {
			return fmt.Sprintf("// TODO: parse %s", name)
		}
		return fmt.Sprintf("uint8(%#x)", n)
	}

	return fmt.Sprintf("// TODO - get not implemented %s", name)
}

func pc(name string) string {
	switch name {
	case "a8", "n8", "e8":
		return "cpu.IncProgramCounter()"
	case "a16", "n16":
		return "cpu.IncProgramCounter(); cpu.IncProgramCounter()"
	default:
		return ""
	}
}

// generates the code that writes, either to a register, or to memory
// e.g. `cpu.WriteMemory(cpu.A, data)` or `cpu.A = cpu.B`
func set(name string, immediate bool, expr string) string {
	if !immediate {
		switch name {
		case "A", "F", "B", "C", "D", "E", "H", "L":
			return fmt.Sprintf("cpu.WriteMemory(concatU16(0xFF, cpu.%s), %s)", name, expr)
		}
		return fmt.Sprintf("cpu.WriteMemory(%s, %s)", get(name, true), expr)
	}
	switch name {
	case "A", "C", "E", "L", "B", "D", "H", "SP":
		return fmt.Sprintf("cpu.%s = %s", name, expr)
	case "BC", "DE", "HL":
		msb := string(name[0])
		lsb := string(name[1])
		return fmt.Sprintf("cpu.%s, cpu.%s = split(%s)", msb, lsb, expr)
	case "AF":
		return fmt.Sprintf("msb, lsb := split(%s)\ncpu.A, cpu.F = msb, FlagRegister(lsb)", expr)
	case "PC":
		return fmt.Sprintf("cpu.PC = %s", expr)
	default:
		return fmt.Sprintf("// todo: set %s %s", name, expr)
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

// retrieves value from an array. If index is too large, fetch the last entry.
// This is primarily related to cycles, where the cycle count may differ based
// on a condition.
func indexOrLast[T any](arr []T, i int) T {
	if i < len(arr) {
		return arr[i]
	}
	return arr[len(arr)-1]
}
