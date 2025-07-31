package gameboy

import (
	"fmt"
	"log/slog"
)

type Register = uint8

type Flags uint8

// https://gbdev.io/pandocs/CPU_Registers_and_Flags.html
var (
	// This bit is set if and only if the result of an operation is zero. Used by conditional jumps.
	FLAGZ Flags = 0x80

	// These flags are used by the DAA instruction only.
	// N indicates whether the previous instruction has been a subtraction, and
	// H indicates carry for the lower 4 bits of the result. DAA also uses the
	// C flag, which must indicate carry for the upper 4 bits. After
	// adding/subtracting two BCD numbers, DAA is used to convert the result to
	// BCD format. BCD numbers range from $00 to $99 rather than $00 to $FF.
	// Because only two flags (C and H) exist to indicate carry-outs of BCD
	// digits, DAA is ineffective for 16-bit operations (which have 4 digits),
	// and use for INC/DEC operations (which do not affect C-flag) has limits.
	FLAGN Flags = 0x40 // operation used subtraction
	FLAGH Flags = 0x20 // raised half-carry

	// Set in the following cases:
	// - When the result of an 8-bit addition is higher than $FF.
	// - When the result of a 16-bit addition is higher than $FFFF.
	// - When the result of a subtraction or comparison is lower than zero
	//   (like in Z80 and x86 CPUs, but unlike in 65XX and ARM CPUs).
	// - When a rotate/shift operation shifts out a “1” bit.
	FLAGC Flags = 0x10 // overflow (carry)
)

type CPU struct {
	A Register     // accumulator
	F FlagRegister // flag register

	B Register
	C Register

	D Register
	E Register

	H Register
	L Register

	PC uint16
	SP uint16 // stack pointer

	cycles int

	mem *Memory
	log *slog.Logger

	// last error from Step()
	err error
}

func (cpu *CPU) Err() error {
	return cpu.err
}

func NewCPU() *CPU {
	return &CPU{}
}

func (cpu *CPU) HL() uint16 { return uint16(cpu.H)<<8 | uint16(cpu.L) }

// bc, de,
func (cpu *CPU) BC() uint16 { return uint16(cpu.B)<<8 | uint16(cpu.C) }
func (cpu *CPU) DE() uint16 { return uint16(cpu.D)<<8 | uint16(cpu.E) }

func (cpu *CPU) Add(lhs, rhs uint8) (res uint8, flags FlagRegister) {
	var fl Flags
	res = lhs + rhs
	if res < lhs || res < rhs {
		fl |= FLAGC
	}
	if res == 0 {
		fl |= FLAGZ
	}
	return res, FlagRegister(fl)
}

func (cpu *CPU) Add16(lhs, rhs uint16) (res uint16, flags FlagRegister) {
	var fl Flags
	res = lhs + rhs
	if res < lhs || res < rhs {
		fl |= FLAGC
	}
	if res == 0 {
		fl |= FLAGZ
	}
	return res, FlagRegister(fl)
}

func (cpu *CPU) AddSigned16(lhs, rhs int16) (res int16, flags FlagRegister) {
	var fl Flags
	res = lhs + rhs
	// TODO: finish this
	return res, FlagRegister(fl)
}
func (cpu *CPU) AddSigned8(lhs, rhs int8) (res int8, flags FlagRegister) {
	var fl Flags
	res = lhs + rhs
	// TODO: finish this
	return res, FlagRegister(fl)
}

func (cpu *CPU) load(addr uint16, dst any) {
	b, ok := cpu.mem.Access(addr)
	if !ok {
		cpu.err = ErrNoMoreInstructions
	}
	if addr == 0x2233 {
		fmt.Printf("yay %d\n", b)
	}
	switch concr := dst.(type) {
	case *uint8:
		*concr = b
	case *int16:
		panic("TODO: int16")
	default:
		panic(fmt.Sprintf("cpu.load: not implemented for %T", dst))
	}
}

func (cpu *CPU) Step() bool {
	if cpu.err != nil {
		return false
	}

	var instr uint8
	cpu.load(cpu.PC, &instr)
	if cpu.err != nil { // if loading next instruction failed, we'll stop
		return false
	}
	log := cpu.log.With("PC", cpu.PC)
	cpu.PC++
	// TODO: handle prefix...
	if instr == 0xCB {
		panic("TODO: handle prefix")
	}
	if instr == 0x00 {
		return true // NOP command
	}
	if instr == 0x01 {
		cpu.err = ErrNoMoreInstructions
		return false
	}
	op, ok := ops[instr]
	if !ok {
		cpu.err = fmt.Errorf("unknown code: %#x", instr)
		return false
	}
	op(cpu) // execute the operation
	log.Info("ran a step", "curr", cpu.PC)

	return true
}

type FlagRegister Register

func (f FlagRegister) HasCarry() bool { return (uint8(f) & uint8(FLAGC)) > 0 }

func splitU16(v uint16) (msb, lsb uint8) {
	msb = uint8(v >> 8)
	lsb = uint8(v)
	return
}

func Run(cpu *CPU, mem *Memory, log *slog.Logger) error {
	cpu.mem = mem
	cpu.log = log

	for cpu.Step() {
	}
	return cpu.Err()
}
