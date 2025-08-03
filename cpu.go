package gameboy

import (
	"fmt"
	"io"
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

	// Points to next instruction to run
	PC uint16
	// Points to the bottom of the stack (NOT the next free value - if you want the
	// next free value, then need to decrement first)
	SP uint16

	cycles int

	mem *Memory
	log *slog.Logger

	// last error from Step()
	err error
}

func (cpu *CPU) Err() error {
	return cpu.err
}

func (cpu *CPU) HL() uint16 { return concatU16(cpu.H, cpu.L) }
func (cpu *CPU) BC() uint16 { return concatU16(cpu.B, cpu.C) }
func (cpu *CPU) DE() uint16 { return concatU16(cpu.D, cpu.E) }
func (cpu *CPU) AF() uint16 { return concatU16(cpu.A, Register(cpu.F)) }

// loads and increments the progrm counter
func (cpu *CPU) load(addr uint16, dst any) {
	b, ok := cpu.mem.Access(addr)
	cpu.log.Debug("Accessing memory", "loc", hexstr(addr), "res", hexstr(b))
	if !ok {
		cpu.err = ErrNoMoreInstructions
	}
	cpu.IncProgramCounter()
	switch concr := dst.(type) {
	case *uint8:
		*concr = b
	case *int16:
		panic("TODO: int16")
	default:
		panic(fmt.Sprintf("cpu.load: not implemented for %T", dst))
	}
}
func (cpu *CPU) loadU8(addr uint16) uint8 {
	var dst uint8
	cpu.load(addr, &dst)
	return dst
}

func (cpu *CPU) readU8(addr uint16) uint8 {
	var value uint8
	cpu.load(addr, &value)
	return value
}
func (cpu *CPU) readU16(addr uint16) uint16 {
	var msb, lsb byte
	cpu.load(addr, &msb)
	cpu.load(addr+1, &lsb)
	return concatU16(msb, lsb)
}
func (cpu *CPU) readI8(addr uint16) int8 {
	var value byte
	cpu.load(addr, &value)
	return int8(value)
}

// Writing to stack: use PushStack instead of this.
func (cpu *CPU) WriteMemory(addr uint16, value any) {
	if cpu.mem == nil {
		panic("cpu.mem is nil")
	}
	switch value := value.(type) {
	case uint8:
		cpu.mem.WriteData(addr, []byte{value})
	case uint16:
		msb, lsb := split(value)
		cpu.mem.WriteData(addr, []byte{lsb, msb})
	}
}

func (cpu *CPU) PushStack(value any) {
	cpu.log.Debug("Writing to stack", "SP", hexstr(cpu.SP), "value", value)
	switch value := value.(type) {
	case uint16:
		msb, lsb := split(value)
		cpu.SP--
		cpu.WriteMemory(cpu.SP, msb)
		cpu.SP--
		cpu.WriteMemory(cpu.SP, lsb)
	default:
		panic(fmt.Sprintf("cpu.PushStack: not implemented for %T", value))
	}
	cpu.log.Debug("Stack after push", "SP", hexstr(cpu.SP))
}
func (cpu *CPU) PopStack() uint16 {
	cpu.log.Debug("Popping stack", "SP", hexstr(cpu.SP))
	if cpu.SP == 0xFFFF {
		cpu.err = ErrStackUnderflow
		return 0
	}
	lsb := cpu.readU8(cpu.SP)
	cpu.SP++
	msb := cpu.readU8(cpu.SP)
	cpu.SP++
	return concatU16(msb, lsb)
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
	log := cpu.log.With("PC", cpu.PC, "instr", hexstr(instr))
	// TODO: handle prefix...
	if instr == 0xCB {
		panic("TODO: handle prefix")
	}
	if instr == 0x00 {
		cpu.err = ErrNoMoreInstructions
		return false // NOP command
	}
	if instr == 0x10 {
		cpu.err = ErrNoMoreInstructions
		return false
	}
	op, ok := ops[instr]
	log.Debug("instruction start", "name", name(instr))
	if !ok {
		cpu.err = fmt.Errorf("unknown code: %#x", instr)
		return false
	}

	// Invariant: the PC is at the instruction at the start of the operation
	// and should end at the last instruction related to this command. For
	// very simple instructions (eg ADD A B), then the PC should not be changed
	// by op. However, for instructions that load data, the op should move the
	// PC towards the last instruction that is done by the op.

	op(cpu) // execute the operation
	log.Debug("instruction done", "curr", fmt.Sprintf("%#x", cpu.PC))

	return true
}

func (cpu *CPU) IncProgramCounter(src ...string) {
	cpu.PC++
	if cpu.log != nil {
		s := "N/A"
		if len(src) > 0 {
			s = src[0]
		}
		cpu.log.Debug("PC increment", "new", fmt.Sprintf("%#x", cpu.PC), "src", s)
	}
}

func (cpu *CPU) Dump(w io.Writer) {
	fmt.Fprintf(w, "Cycles: %d\n", cpu.cycles)
	fmt.Fprintf(w, "A:  %#02x     F:  %#02x     AF: %#04x\n", cpu.A, uint8(cpu.F), cpu.AF())
	fmt.Fprintf(w, "B:  %#02x     C:  %#02x     BC: %#04x\n", cpu.B, cpu.C, cpu.BC())
	fmt.Fprintf(w, "D:  %#02x     E:  %#02x	  DE: %#04x\n", cpu.D, cpu.E, cpu.DE())
	fmt.Fprintf(w, "H:  %#02x     L:  %#02x	  HL: %#04x\n", cpu.H, cpu.L, cpu.HL())
	fmt.Fprintf(w, "                          SP: %#04x\n", cpu.SP)
	fmt.Fprintf(w, "                          PC: %#04x\n", cpu.PC)
	// fmt.Fprintf(w, "HL: %#04x   BC: %#04x   DE: %#04x   AF: %#04x\n", cpu.HL(), cpu.BC(), cpu.DE(), cpu.AF())
}

type FlagRegister Register

func (f FlagRegister) HasCarry() bool { return (uint8(f) & uint8(FLAGC)) > 0 }
func (f FlagRegister) HasZero() bool  { return (uint8(f) & uint8(FLAGZ)) > 0 }

func Run(cpu *CPU, mem *Memory, log *slog.Logger) error {
	cpu.mem = mem
	cpu.log = log

	for cpu.Step() {
	}
	return cpu.Err()
}
