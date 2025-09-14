package gameboy

import (
	"encoding/hex"
	"errors"
	"fmt"
	"io"
	"log/slog"
	"os"
	"testing"
)

const (
	INSTR_STOP = 0x10
)

func TestInstructions(t *testing.T) {

	cases := []struct {
		desc    string
		cpu     CPU                                // initial state of cpu
		initMem func(m *Memory)                    // instructions to run
		check   func(t *testing.T, cpu *CPUHelper) // post-check func
		debug   bool                               // if true, debug-level logging
	}{
		{
			desc: "ADD A B",
			cpu: CPU{
				A: 5, B: 2,
			},
			initMem: func(m *Memory) {
				m.Write([]byte{0x80})
			},
			check: func(t *testing.T, cpu *CPUHelper) {
				cpu.ExpectA(7)
			},
		},
		{
			desc: "ADD A HL",
			cpu: CPU{
				A: 1,
				H: 0x00,
				L: 0x0a,
			},
			initMem: func(m *Memory) {
				m.WriteInstr(0x86) // ADD A,HL
				m.WriteInstr(0x01) // STOP
				m.WriteAt(0x000a, 0x44)
			},
			check: func(t *testing.T, cpu *CPUHelper) {
				cpu.ExpectA(0x45)
			},
		},
		{
			desc: "ADD HL BC",
			cpu: CPU{
				H: 0x11,
				L: 0x11,
				B: 0x22,
				C: 0x22,
			},
			initMem: func(m *Memory) {
				m.WriteInstr(0x09)
			},
			check: func(t *testing.T, cpu *CPUHelper) {
				cpu.Dump(os.Stderr)
				cpu.ExpectHL(0x3333) // 0x1111 + 0x2222 = 0x3333
			},
		},
		{
			desc: "ADD HL HL",
			cpu: CPU{
				H: 0x11,
				L: 0x22,
			},
			initMem: func(m *Memory) {
				m.WriteInstr(0x29)
			},
			check: func(t *testing.T, cpu *CPUHelper) {
				cpu.ExpectHL(0x2244)
				cpu.ExpectFlagCarryUnset()
			},
		},
		{
			desc: "ADD A (HL)",
			cpu: CPU{
				A: 1,
				H: 0x11,
				L: 0x22,
			},
			initMem: func(m *Memory) {
				m.WriteInstr(0x86)
				m.WriteInstr(INSTR_STOP)
				m.WriteAt(0x1122, 0x02)
			},
			check: func(t *testing.T, cpu *CPUHelper) {
				cpu.ExpectA(0x03)
			},
		},
		{
			desc: "ADD A B with overflow",
			cpu: CPU{
				A: 56,
				B: 200,
			},
			initMem: func(m *Memory) {
				m.WriteInstr(0x80)
			},
			check: func(t *testing.T, cpu *CPUHelper) {
				cpu.ExpectA(0x00)
				cpu.ExpectFlagCarry()
				cpu.ExpectFlagZero()
			},
		},
		{
			desc: "INC (HL)",
			cpu: CPU{
				H: 0x11,
				L: 0x22,
			},
			initMem: func(m *Memory) {
				m.WriteInstr(0x34)
				m.WriteInstr(INSTR_STOP)
				m.WriteAt(0x1122, 0x33)
			},
			check: func(t *testing.T, cpu *CPUHelper) {
				cpu.ExpectMem(0x1122, 0x34)
			},
		},
		{
			desc: "INC BC 0x03",
			cpu: CPU{
				B: 0xFF,
				C: 0xFF,
			},
			initMem: func(m *Memory) {
				m.Write(0x03, INSTR_STOP)
			},
			check: func(t *testing.T, cpu *CPUHelper) {
				cpu.ExpectBC(0x00)
				cpu.ExpectFlagCarry()
				cpu.ExpectFlagZero()
			},
		},
		{
			desc: "INC B 0x04",
			cpu: CPU{
				B: 0x11,
			},
			initMem: func(m *Memory) {
				m.Write(0x04, INSTR_STOP)
			},
			check: func(t *testing.T, cpu *CPUHelper) {
				cpu.ExpectB(0x12)
			},
		},
		{
			desc: "DEC B 0x05",
			cpu: CPU{
				B: 0x00,
			},
			initMem: func(m *Memory) {
				m.Write(0x05, INSTR_STOP)
			},
			check: func(t *testing.T, cpu *CPUHelper) {
				cpu.ExpectB(0xFF)
				cpu.ExpectFlagCarry()
			},
		},
		{
			desc: "INC SP 0x33",
			cpu: CPU{
				SP: 0x0011,
			},
			initMem: func(m *Memory) {
				m.WriteInstr(0x33)
				m.WriteInstr(INSTR_STOP)
			},
			check: func(t *testing.T, cpu *CPUHelper) { cpu.ExpectSP(0x0012) },
		},
		{
			desc: "INC (HL) 0x34",
			cpu: CPU{
				H: 0x11,
				L: 0x22,
			},
			initMem: func(m *Memory) {
				m.Write(0x34)
				m.WriteAt(0x1122, 0x33)
			},
			check: func(t *testing.T, cpu *CPUHelper) {
				cpu.ExpectMem(0x1122, 0x34)
			},
		},
		{
			desc: "DEC (HL) 0x35",
			cpu: CPU{
				H: 0x11,
				L: 0x22,
			},
			initMem: func(m *Memory) {
				m.Write(0x35)
				m.WriteAt(0x1122, 0x00)
			},
			check: func(t *testing.T, cpu *CPUHelper) {
				cpu.ExpectMem(0x1122, 0xFF)
				cpu.ExpectFlagCarry()
			},
		},
		{
			desc: "DEC HL",
			cpu:  CPU{},
			initMem: func(m *Memory) {
				m.WriteInstr(0x2b)
			},
			check: func(t *testing.T, cpu *CPUHelper) {
				cpu.ExpectHL(0xffff)
				cpu.ExpectFlagCarry()
			},
		},
		{
			desc: "LD B,n8 0x06",
			cpu:  CPU{},
			initMem: func(m *Memory) {
				m.Write(0x06, 0xab, INSTR_STOP)
			},
			check: func(t *testing.T, cpu *CPUHelper) {
				cpu.ExpectB(0xab)
			},
		},
		{
			desc: "LD A,(BC) 0x0A",
			cpu: CPU{
				B: 0x11,
				C: 0x22,
			},
			initMem: func(m *Memory) {
				m.Write(0x0A, INSTR_STOP)
				m.WriteAt(0x1122, 0x33)
			},
			check: func(t *testing.T, cpu *CPUHelper) {
				cpu.ExpectA(0x33)
			},
		},
		{
			desc: "LD C,n8 0x0E",
			cpu:  CPU{},
			initMem: func(m *Memory) {
				m.Write(0x0E, 0x33, INSTR_STOP)
			},
			check: func(t *testing.T, cpu *CPUHelper) {
				cpu.ExpectC(0x33)
			},
		},
		{
			desc: "LD (HL+),A 0x22",
			cpu: CPU{
				H: 0x11,
				L: 0x22,
				A: 0x33,
			},
			initMem: func(m *Memory) {
				m.Write(0x22, INSTR_STOP)
				m.WriteAt(0x1122, 0x01)
			},
			check: func(t *testing.T, cpu *CPUHelper) {
				cpu.ExpectMem(0x1122, 0x33)
				cpu.ExpectHL(0x1123) // the post increment operator
			},
		},
		{
			desc: "LD (HL-),A 0x32",
			cpu: CPU{
				H: 0x11,
				L: 0x22,
				A: 0x33,
			},
			initMem: func(m *Memory) {
				m.Write(0x32, INSTR_STOP)
				m.WriteAt(0x1122, 0x01)
			},
			check: func(t *testing.T, cpu *CPUHelper) {
				cpu.ExpectMem(0x1122, 0x33)
				cpu.ExpectHL(0x1121) // the post decrement operator
			},
		},
		{
			// does this even make sense?? load into itself?
			desc: "LD B,B 0x40",
			cpu: CPU{
				B: 0x02,
			},
			initMem: func(m *Memory) {
				m.Write(0x40, INSTR_STOP)
			},
			check: func(t *testing.T, cpu *CPUHelper) {
				cpu.ExpectB(0x02)
			},
		},
		{
			desc: "LD B,A 0x47",
			cpu: CPU{
				A: 0x03,
			},
			initMem: func(m *Memory) {
				m.Write(0x47, INSTR_STOP)
			},
			check: func(t *testing.T, cpu *CPUHelper) {
				cpu.ExpectB(0x03)
			},
		},
		{
			desc: "LD C,(HL) 0x4E",
			cpu: CPU{
				H: 0x11,
				L: 0x22,
			},
			initMem: func(m *Memory) {
				m.Write(0x4E, INSTR_STOP)
				m.WriteAt(0x1122, 0x33)
			},
			check: func(t *testing.T, cpu *CPUHelper) {
				cpu.ExpectC(0x33)
			},
		},
		{
			desc: "LD (HL),B 0x70",
			cpu: CPU{
				H: 0x11,
				L: 0x22,
				B: 0x33,
			},
			initMem: func(m *Memory) {
				m.Write(0x70, INSTR_STOP)
				m.WriteAt(0x1122, 0x01)
			},
			check: func(t *testing.T, cpu *CPUHelper) {
				cpu.ExpectMem(0x1122, 0x33)
			},
		},
		{
			desc: "LD SP,HL 0xF9",
			cpu: CPU{
				H:  0x11,
				L:  0x22,
				SP: 0x5555,
			},
			initMem: func(m *Memory) {
				m.Write(0xF9, INSTR_STOP)
			},
			check: func(t *testing.T, cpu *CPUHelper) {
				cpu.ExpectSP(0x1122)
			},
		},
		{
			desc: "CALL NZ,a16 0xC4",
			cpu: CPU{
				F: 0x00,
			},
			initMem: func(m *Memory) {
				m.Write(0xC4, 0x22, 0x11) // stored as lsb, msb -- so bits reversed
				m.WriteAt(0x1122, INSTR_STOP)
			},
			check: func(t *testing.T, cpu *CPUHelper) {
				cpu.ExpectPC(0x1123) // 0x1122 + 1, since we're reading another instruction before stopping
				cpu.ExpectSP(0xFFFD)

				// read instr   1
				// read a16     2
				// 3 instructions done, so the NEXT is at 0x03
				cpu.ExpectPeekStack(uint16(0x03))
				cpu.ExpectCycleCount(24)
			},
		},
		{
			desc: "CALL NZ,a16 0xC4 - flag zero",
			cpu: CPU{
				F: FlagRegister(FLAGZ),
			},
			initMem: func(m *Memory) {
				m.Write(0xC4, 0x22, 0x11) // stored as lsb, msb -- so bits reversed
				m.WriteAt(0x1122, INSTR_STOP)
			},
			check: func(t *testing.T, cpu *CPUHelper) {
				// instr + a16 + stop
				// 1     + 2   + 1
				// -> 4 instructions ran, so the next is at 0x04
				cpu.ExpectPC(0x04)
				cpu.ExpectCycleCount(12)
			},
		},
		{
			desc: "CALL 0xCD",
			cpu:  CPU{},
			initMem: func(m *Memory) {
				m.Write(0xCD, 0x22, 0x11)
				m.WriteAt(0x1122, INSTR_STOP)
			},
			check: func(t *testing.T, cpu *CPUHelper) {
				cpu.ExpectPC(0x1123)
				cpu.ExpectCycleCount(24)
			},
		},
		{
			desc: "PUSH BC 0xC5",
			cpu: CPU{
				B: 0x11,
				C: 0x22,
			},
			initMem: func(m *Memory) {
				m.Write(0xC5, INSTR_STOP)
			},
			check: func(t *testing.T, cpu *CPUHelper) {
				cpu.ExpectPeekStack(uint16(0x1122)) // 0x22
			},
		},
		{
			desc: "POP BC 0xC1",
			cpu: CPU{
				D: 0x11,
				E: 0x22,
			},
			initMem: func(m *Memory) {
				m.Write(0xD5) // PUSH DE
				m.Write(0xC1) // POP BC
				m.Write(INSTR_STOP)
			},
			check: func(t *testing.T, cpu *CPUHelper) {
				cpu.DumpStack(os.Stderr)
				cpu.ExpectBC(0x1122)
			},
		},
		{
			desc: "POP BC 0xC1 - empty stack",
			cpu:  CPU{},
			initMem: func(m *Memory) {
				m.Write(0xC1)
			},
			check: func(t *testing.T, cpu *CPUHelper) {
				cpu.ExpectErr(t, ErrStackUnderflow)
			},
		},
		{
			// a subroutine that adds two numbers, then returns. We check whether the
			// summed value is in the accumulator and the PC is as expected
			desc: "RET 0xC9",
			cpu: CPU{
				A: 0x01,
				E: 0x02,
			},
			initMem: func(m *Memory) {
				m.Write("CALL a16", 0x22, 0x11)
				m.Write("STOP")
				m.CursorAt(0x1122)
				m.Write("ADD A,E")
				m.Write("RET")
			},
			check: func(t *testing.T, cpu *CPUHelper) {
				cpu.ExpectA(0x03)
				cpu.ExpectPC(0x04)
			},
		},
		{
			desc: "SUB A,B 0x90",
			cpu: CPU{
				A: 0x01,
				B: 0x03,
			},
			initMem: func(m *Memory) {
				m.Write("SUB A,B")
				m.Write("STOP")
			},
			check: func(t *testing.T, cpu *CPUHelper) {
				cpu.ExpectA(0xFE)
				cpu.ExpectFlagCarry()
			},
		},
		{
			desc: "SUB A,(HL) 0x96",
			cpu: CPU{
				A: 0x02,
				H: 0x11,
				L: 0x22,
			},
			initMem: func(m *Memory) {
				m.Write("SUB A,(HL)")
				m.Write("STOP")
				m.CursorAt(0x1122)
				m.Write(0x01)
			},
			check: func(t *testing.T, cpu *CPUHelper) {
				cpu.ExpectA(0x01)
			},
		},
		{
			desc: "SUB A,A 0x97",
			cpu: CPU{
				A: 0x01,
			},
			initMem: func(m *Memory) {
				m.Write("SUB A,A")
				m.Write("STOP")
			},
			check: func(t *testing.T, cpu *CPUHelper) {
				cpu.ExpectA(0x00)
				cpu.ExpectFlagZero()
			},
		},
		{
			desc: "OR A,B 0xB0",
			cpu: CPU{
				A: 0b11000011,
				B: 0b00001111,
			},
			initMem: func(m *Memory) {
				m.Write("OR A,B")
				m.Write("STOP")
			},
			check: func(t *testing.T, cpu *CPUHelper) {
				cpu.ExpectA(0b11001111)
			},
		},
		{
			desc: "OR A,(HL) 0xB6",
			cpu: CPU{
				A: 0b11000011,
				H: 0x11,
				L: 0x22,
			},
			initMem: func(m *Memory) {
				m.Write("OR A,(HL)")
				m.Write("STOP")
				m.CursorAt(0x1122)
				m.Write(0b00001111)
			},
			check: func(t *testing.T, cpu *CPUHelper) {
				cpu.ExpectA(0b11001111)
			},
		},
		{
			desc: "XOR A,B 0xA8",
			cpu: CPU{
				A: 0b11000011,
				B: 0b00001111,
			},
			initMem: func(m *Memory) {
				m.Write("XOR A,B")
				m.Write("STOP")
			},
			check: func(t *testing.T, cpu *CPUHelper) {
				cpu.ExpectA(0b11001100)
			},
		},
		{
			desc: "XOR A,(HL) 0xAE",
			cpu: CPU{
				A: 0b11000011,
				H: 0x11,
				L: 0x22,
			},
			initMem: func(m *Memory) {
				m.Write("XOR A,(HL)")
				m.Write("STOP")
				m.CursorAt(0x1122)
				m.Write(0b00001111)
			},
			check: func(t *testing.T, cpu *CPUHelper) {
				cpu.ExpectA(0b11001100)
			},
		},
		{
			desc: "CP A,B 0xB8",
			cpu: CPU{
				A: 0x01,
				B: 0x01,
			},
			initMem: func(m *Memory) {
				m.Write("CP A,B")
				m.Write("STOP")
			},
			check: func(t *testing.T, cpu *CPUHelper) {
				cpu.ExpectFlagZero()
				cpu.ExpectA(0x01)
			},
		},
		{
			desc: "CP A,(HL) 0xBE",
			cpu: CPU{
				A: 0x01,
				H: 0x11,
				L: 0x22,
			},
			initMem: func(m *Memory) {
				m.Write("CP A,(HL)")
				m.CursorAt(0x1122)
				m.Write(0x02)
			},
			check: func(t *testing.T, cpu *CPUHelper) {
				cpu.ExpectFlagCarry()
				cpu.ExpectA(0x01)
			},
		},
		{
			desc: "CP A,n8 0xFE",
			cpu: CPU{
				A: 0x01,
			},
			initMem: func(m *Memory) {
				m.Write("CP A,n8", 0x00)
				m.Write("STOP")
			},
			check: func(t *testing.T, cpu *CPUHelper) {
				cpu.ExpectFlagZeroUnset()
				cpu.ExpectFlagCarryUnset()
				cpu.ExpectA(0x01)
			},
		},
		{
			desc: "BIT 0,B 0x40",
			cpu: CPU{
				B: 0x00,
			},
			initMem: func(m *Memory) {
				m.Write("PREFIX", "BIT 0,B")
			},
			check: func(t *testing.T, cpu *CPUHelper) {
				cpu.ExpectFlagZero()
			},
		},
		{
			desc: "BIT 0,C 0x41",
			cpu: CPU{
				C: 0x01,
			},
			initMem: func(m *Memory) {
				m.Write("PREFIX", "BIT 0,C")
			},
			check: func(t *testing.T, cpu *CPUHelper) {
				cpu.ExpectFlagZeroUnset()
			},
		},
		{
			desc: "BIT 1,B 0x48",
			cpu: CPU{
				B: 0x01,
			},
			initMem: func(m *Memory) {
				m.Write("PREFIX", "BIT 1,B")
			},
			check: func(t *testing.T, cpu *CPUHelper) {
				cpu.ExpectFlagZero()
			},
		},
		{
			desc: "BIT 1,C 0x49",
			cpu: CPU{
				C: 0x02,
			},
			initMem: func(m *Memory) {
				m.Write("PREFIX", "BIT 1,C")
			},
			check: func(t *testing.T, cpu *CPUHelper) {
				cpu.ExpectFlagZeroUnset()
			},
		},
		{
			desc: "BIT 1,(HL) 0x4E",
			cpu: CPU{
				H: 0x11,
				L: 0x22,
			},
			initMem: func(m *Memory) {
				m.Write("PREFIX", "BIT 1,(HL)")
				m.CursorAt(0x1122)
				m.Write(0x02)
			},
			check: func(t *testing.T, cpu *CPUHelper) {
				cpu.ExpectFlagZeroUnset()
				cpu.ExpectFlagHigh()
			},
		},
		{
			desc: "JR e8 0x18",
			cpu:  CPU{},
			initMem: func(m *Memory) {
				m.Write("JR e8", 0x05)
			},
			check: func(t *testing.T, cpu *CPUHelper) {
				cpu.ExpectPC(0x08)
				cpu.ExpectCycleCount(12)
			},
		},
		{
			desc: "JR Z,e8 0x28",
			cpu: CPU{
				F: FlagRegister(FLAGZ),
			},
			initMem: func(m *Memory) {
				m.Write("JR Z,e8", 0x05)
			},
			check: func(t *testing.T, cpu *CPUHelper) {
				cpu.ExpectPC(0x08)
				cpu.ExpectCycleCount(12)
			},
		},
		{
			desc: "JR NZ,e8 0x20",
			cpu: CPU{
				F: FlagRegister(FLAGZ),
			},
			initMem: func(m *Memory) {
				m.Write("JR NZ,e8", 0x05)
			},
			check: func(t *testing.T, cpu *CPUHelper) {
				cpu.ExpectPC(0x03)
				cpu.ExpectCycleCount(8)
			},
		},
		{
			desc: "LDH (a8),A 0xE0",
			cpu: CPU{
				A: 0x01,
			},
			initMem: func(m *Memory) {
				m.Write("LDH (a8),A", 0x11)
			},
			check: func(t *testing.T, cpu *CPUHelper) {
				cpu.ExpectMem(0xFF11, 0x01)
			},
		},
		{
			desc: "LDH (C),A 0xE2",
			cpu: CPU{
				C: 0x11,
				A: 0x01,
			},
			initMem: func(m *Memory) {
				m.Write("LDH (C),A")
			},
			check: func(t *testing.T, cpu *CPUHelper) {
				cpu.ExpectMem(0xFF11, 0x01)
			},
		},
		{
			desc: "LDH A,(a8) 0xF0",
			cpu:  CPU{},
			initMem: func(m *Memory) {
				m.Write("LDH A,(a8)", 0x11)
				m.CursorAt(0xFF11)
				m.Write(0x01)
			},
			check: func(t *testing.T, cpu *CPUHelper) {
				cpu.ExpectA(0x01)
			},
		},
		{
			desc: "LDH A,(C) 0xF2",
			cpu: CPU{
				C: 0x11,
			},
			initMem: func(m *Memory) {
				m.Write("LDH A,(C)")
				m.CursorAt(0xFF11)
				m.Write(0x01)
			},
			check: func(t *testing.T, cpu *CPUHelper) {
				cpu.ExpectA(0x01)
			},
		},
		{
			desc:  "LD A,n8 0xF3",
			debug: true,
			cpu:   CPU{},
			initMem: func(m *Memory) {
				m.Write("LD A,n8", 0x01)
				m.Write("ADD A,A")
			},
			check: func(t *testing.T, cpu *CPUHelper) {
				cpu.DumpMem(0, 10)
				cpu.ExpectA(0x02)
			},
		},
		{
			desc: "RL B 0x01",
			cpu: CPU{
				B: 0b11000000,
			},
			initMem: func(m *Memory) {
				m.Write("PREFIX", "RL B", "STOP")
			},
			check: func(t *testing.T, cpu *CPUHelper) {
				cpu.ExpectB(0b10000000)
				cpu.ExpectFlagCarry()
			},
		},
		{
			desc: "RL (HL) 0x16",
			cpu: CPU{
				H: 0x11,
				L: 0x22,
			},
			initMem: func(m *Memory) {
				m.Write("PREFIX", "RL (HL)")
				m.CursorAt(0x1122)
				m.Write(0b11000000)
			},
			check: func(t *testing.T, cpu *CPUHelper) {
				cpu.ExpectMem(0x1122, 0b10000000)
				cpu.ExpectFlagCarry()
			},
		},
		{
			desc:  "LDH regression",
			cpu:   CPU{},
			debug: true,
			initMem: func(m *Memory) {
				m.Write("LDH A,(a8)", 0x44)
				m.Write("CP A,n8", 0x01)
				m.CursorAt(0xFF44)
				m.Write(0x01)
			},
			check: func(t *testing.T, cpu *CPUHelper) {
				cpu.ExpectFlagZero()
			},
		},
		{
			desc: "RLA with carry",
			cpu: CPU{
				F: FLAGC,
			},
			initMem: func(m *Memory) {
				m.Write("RLA")
			},
			check: func(t *testing.T, cpu *CPUHelper) {
				cpu.ExpectA(1)
			},
		},
		{
			desc: "RLA always zero flag",
			cpu:  CPU{},
			initMem: func(m *Memory) {
				m.Write("RLA")
			},
			check: func(t *testing.T, cpu *CPUHelper) {
				cpu.ExpectFlagZeroUnset()
			},
		},
		{
			desc: "RLA rotate around",
			cpu: CPU{
				A: 1,
			},
			initMem: func(m *Memory) {
				for range 8 {
					m.Write("RLA")
				}
			},
			check: func(t *testing.T, cpu *CPUHelper) {
				cpu.ExpectFlagCarry()
			},
		},
		{
			desc: "AND A,B",
			cpu:  CPU{A: 0b011, B: 0b101},
			initMem: func(m *Memory) {
				m.Write("AND A,B")
			},
			check: func(t *testing.T, cpu *CPUHelper) {
				cpu.ExpectA(0b001)
			},
		},
		{
			desc: "AND A,n8",
			cpu:  CPU{A: 0b011},
			initMem: func(m *Memory) {
				m.Write("AND A,n8", 0b101)
			},
			check: func(t *testing.T, cpu *CPUHelper) {
				cpu.ExpectA(0b001)
			},
		},
		{
			desc: "SWAP A",
			cpu:  CPU{A: 0xab},
			initMem: func(m *Memory) {
				m.Write("PREFIX", "SWAP A")
			},
			check: func(t *testing.T, cpu *CPUHelper) {
				cpu.ExpectA(0xba)
			},
		},
		{
			desc: "SWAP (HL)",
			cpu:  CPU{H: 0x11, L: 0x22},
			initMem: func(m *Memory) {
				m.Write("PREFIX", "SWAP (HL)")
				m.CursorAt(0x1122)
				m.Write(0xab)
			},
			check: func(t *testing.T, cpu *CPUHelper) {
				cpu.ExpectMem(0x1122, 0xba)
			},
		},
		{
			desc: "RES 1,A",
			cpu:  CPU{A: 0xff},
			initMem: func(m *Memory) {
				m.Write("PREFIX", "RES 1,A")
			},
			check: func(t *testing.T, cpu *CPUHelper) {
				cpu.ExpectA(0xfd)
			},
		},
		{
			desc: "RES 1,(HL)",
			cpu: CPU{
				H: 0x11,
				L: 0x22,
			},
			initMem: func(m *Memory) {
				m.Write("PREFIX", "RES 1,(HL)")
				m.CursorAt(0x1122)
				m.Write(0xff)
			},
			check: func(t *testing.T, cpu *CPUHelper) {
				cpu.ExpectMem(0x1122, 0xfd)
			},
		},
		{
			desc: "SET 1,A",
			cpu:  CPU{A: 0x00},
			initMem: func(m *Memory) {
				m.Write("PREFIX", "SET 1,A")
			},
			check: func(t *testing.T, cpu *CPUHelper) {
				cpu.ExpectA(0x02)
			},
		},
		{
			desc: "SET 1,(HL)",
			cpu: CPU{
				H: 0x11,
				L: 0x22,
			},
			initMem: func(m *Memory) {
				m.Write("PREFIX", "SET 1,(HL)")
			},
			check: func(t *testing.T, cpu *CPUHelper) {
				cpu.ExpectMem(0x1122, 0x02)
			},
		},
		{
			desc: "SRL A",
			cpu: CPU{
				A: 0xff,
			},
			initMem: func(m *Memory) {
				m.Write("PREFIX", "SRL A")
			},
			check: func(t *testing.T, cpu *CPUHelper) {
				cpu.ExpectA(0x7f)
				cpu.ExpectFlagCarry()
			},
		},
		{
			desc: "SRL (HL)",
			cpu: CPU{
				H: 0x11, L: 0x22,
			},
			initMem: func(m *Memory) {
				m.Write("PREFIX", "SRL (HL)")
				m.CursorAt(0x1122)
				m.Write(0xff)
			},
			check: func(t *testing.T, cpu *CPUHelper) {
				cpu.ExpectMem(0x1122, 0x7f)
				cpu.ExpectFlagCarry()
			},
		},
		{
			desc: "SRL A zero-flag",
			cpu:  CPU{A: 0x01},
			initMem: func(m *Memory) {
				m.Write("PREFIX", "SRL A")
			},
			check: func(t *testing.T, cpu *CPUHelper) {
				if cpu.err != nil {
					t.Logf("unexpected error: %v", cpu.err)
				}
				cpu.ExpectA(0x00)
				cpu.ExpectFlagCarry()
				cpu.ExpectFlagZero()
			},
		},
		{
			desc: "LD (a16),SP 0x08",
			cpu: CPU{
				SP: 0x3344,
			},
			initMem: func(m *Memory) {
				m.Write(0x08, 0x22, 0x11, INSTR_STOP)
			},
			check: func(t *testing.T, cpu *CPUHelper) {
				cpu.ExpectMem(0x1122, 0x44) // lsb
				cpu.ExpectMem(0x1123, 0x33) // msb
			},
		},
		{
			desc: "LD BC,n16 0x01",
			cpu:  CPU{},
			initMem: func(m *Memory) {
				m.Write(0x01, 0x22, 0x11, INSTR_STOP)
			},
			check: func(t *testing.T, cpu *CPUHelper) {
				cpu.ExpectBC(0x1122)
			},
		},
	}

	initCPU := func(cpu *CPU) {
		if cpu.SP == 0 {
			cpu.SP = 0xffff
		}
	}

	for _, tc := range cases {
		t.Run(tc.desc, func(t *testing.T) {
			cpu := &tc.cpu
			cpu.stopAtnop = true
			initCPU(cpu)
			mem := NewMemory(nil)
			mem.DisableBoot()
			tc.initMem(mem)
			mem.Write(INSTR_STOP) // ensure we have a stop instruction at the end

			Run(cpu, mem, logger(tc.debug))

			defer func() {
				if t.Failed() {
					cpu.Dump(os.Stderr)
				}
			}()
			tc.check(t, &CPUHelper{
				t:   t,
				CPU: cpu,
			})

		})
	}
}

func logger(debug ...bool) *slog.Logger {
	lev := slog.LevelInfo
	if len(debug) > 0 && debug[0] {
		lev = slog.LevelDebug
	}
	return slog.New(slog.NewTextHandler(os.Stderr, &slog.HandlerOptions{
		Level: lev,
		ReplaceAttr: func(groups []string, a slog.Attr) slog.Attr {
			if a.Key == slog.TimeKey || a.Key == slog.LevelKey {
				return slog.Attr{} // remove time attribute
			}
			return a
		},
	}))
}

// === Helper structs ===
type CPUHelper struct {
	t *testing.T
	*CPU
}

func (cpu *CPUHelper) ExpectA(want uint8) {
	cpu.t.Helper()
	if cpu.A != want {
		cpu.t.Fatalf("A: want=%#x, got=%#x", want, cpu.A)
	}
}
func (cpu *CPUHelper) ExpectB(want uint8) {
	cpu.t.Helper()
	if cpu.B != want {
		cpu.t.Fatalf("B: want=%#x, got=%#x", want, cpu.B)
	}
}
func (cpu *CPUHelper) ExpectC(want uint8) {
	cpu.t.Helper()
	if cpu.C != want {
		cpu.t.Fatalf("C: want=%#x, got=%#x", want, cpu.C)
	}
}
func (cpu *CPUHelper) ExpectD(want uint8) {
	cpu.t.Helper()
	if cpu.D != want {
		cpu.t.Fatalf("D: want=%#x, got=%#x", want, cpu.D)
	}
}
func (cpu *CPUHelper) ExpectE(want uint8) {
	cpu.t.Helper()
	if cpu.E != want {
		cpu.t.Fatalf("E: want=%#x, got=%#x", want, cpu.E)
	}
}
func (cpu *CPUHelper) ExpectH(want uint8) {
	cpu.t.Helper()
	if cpu.H != want {
		cpu.t.Fatalf("H: want=%#x, got=%#x", want, cpu.H)
	}
}
func (cpu *CPUHelper) ExpectL(want uint8) {
	cpu.t.Helper()
	if cpu.L != want {
		cpu.t.Fatalf("L: want=%#x, got=%#x", want, cpu.L)
	}
}
func (cpu *CPUHelper) ExpectPC(want uint16) {
	cpu.t.Helper()
	if cpu.PC != want {
		cpu.t.Fatalf("PC: want=%#x, got=%#x", want, cpu.PC)
	}
}
func (cpu *CPUHelper) ExpectHL(want uint16) {
	cpu.t.Helper()
	if cpu.HL() != want {
		cpu.t.Fatalf("HL: want=%#x, got=%#x", want, cpu.HL())
	}
}
func (cpu *CPUHelper) ExpectBC(want uint16) {
	cpu.t.Helper()
	if cpu.BC() != want {
		cpu.t.Fatalf("BC: want=%#x, got=%#x", want, cpu.BC())
	}
}
func (cpu *CPUHelper) ExpectDE(want uint16) {
	cpu.t.Helper()
	if cpu.DE() != want {
		cpu.t.Fatalf("DE: want=%#x, got=%#x", want, cpu.DE())
	}
}
func (cpu *CPUHelper) ExpectSP(want uint16) {
	cpu.t.Helper()
	if cpu.SP != want {
		cpu.t.Fatalf("SP: want=%#x, got=%#x", want, cpu.SP)
	}
}
func (cpu *CPUHelper) ExpectFlagCarry() {
	cpu.t.Helper()
	if !cpu.F.HasCarry() {
		cpu.t.Fatalf("expected carry flag to be set, but it's not")
	}
}
func (cpu *CPUHelper) ExpectFlagCarryUnset() {
	cpu.t.Helper()
	if cpu.F.HasCarry() {
		cpu.t.Fatalf("carry flag is set, expected unset")
	}
}
func (cpu *CPUHelper) ExpectFlagZero() {
	cpu.t.Helper()
	if !cpu.F.HasZero() {
		cpu.t.Fatalf("expected zero flag to be set, but it's not")
	}
}
func (cpu *CPUHelper) ExpectFlagZeroUnset() {
	cpu.t.Helper()
	if cpu.F.HasZero() {
		cpu.t.Fatalf("zero flag is set, expected unset")
	}
}
func (cpu *CPUHelper) ExpectFlagHighUnset() {
	cpu.t.Helper()
	if cpu.F.HasHigh() {
		cpu.t.Fatalf("high flag is set, expected unset")
	}
}
func (cpu *CPUHelper) ExpectFlagHigh() {
	cpu.t.Helper()
	if !cpu.F.HasHigh() {
		cpu.t.Fatalf("expected high flag to be set, but it's not")
	}
}
func (cpu *CPUHelper) ExpectMem(offset uint16, want byte) {
	cpu.t.Helper()
	got := cpu.Mem.Read(offset)
	if got != want {
		cpu.t.Fatalf("ExpectByte: want=%#x, got=%#x", want, got)
	}
}
func (cpu *CPUHelper) ExpectPeekStack(want any) {
	t := cpu.t
	// in other words: MSB is the HIGH address, LSB is the LOW address
	msb := cpu.Mem.Read(cpu.SP + 1)
	lsb := cpu.Mem.Read(cpu.SP)
	switch want := want.(type) {
	case uint16:
		got := concatU16(msb, lsb)
		if got != want {
			t.Fatalf("ExpectPeekStack: want=%#v, got=%#v", want, got)
		}
	default:
		t.Fatalf("ExpectPeekStack: not implemented for %T", want)
	}
}
func (cpu *CPUHelper) DumpStack(w io.Writer) {
	t := cpu.t
	t.Helper()
	fmt.Fprintln(w, "=== Dumping stack:")
	for i := uint16(0xffff); i >= cpu.SP; i-- {
		b := cpu.Mem.Read(i)
		fmt.Fprintf(w, "  %#04X: %02x\n", i, b)
	}
	fmt.Fprintln(w, "=== End of stack dump")
}
func (cpu *CPUHelper) DumpMem(lower, upper int) {
	// 	for i := lower; i <= upper; i++ {
	// 		b, _ := cpu.mem.Access(uint16(i))
	// 		cpu.t.Logf("mem[%04X] = %02X", i, b)
	// 	}
	fmt.Printf("%s\n", hex.Dump(cpu.Mem.data[lower:upper]))
}
func (cpu *CPUHelper) ExpectCycleCount(want int) {
	if cpu.Cycles != want {
		cpu.t.Fatalf("want=%d, got=%d", want, cpu.Cycles)
	}
}
func (cpu *CPUHelper) ExpectErr(t *testing.T, want error) {
	t.Helper()
	if cpu.err == nil {
		t.Fatalf("expected error %v, got nil", want)
	}
	if !errors.Is(cpu.err, want) {
		t.Fatalf("expected error %v, got %v", want, cpu.err)
	}
}
