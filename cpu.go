package gameboy

import (
	"fmt"
	"io"
	"log/slog"
	"strconv"
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

const (
	CPU_FREQUENCY = 4_194_304 // 4.2MHz
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

	// whether previous command was the CB-prefix
	prefix bool

	Cycles     int
	InstrCount int
	limit      int
	hooks      []func(*CPU, int, Instruction, *slog.Logger)
	stopAtnop  bool

	Mem *Memory
	log *slog.Logger

	// peripherals
	ppu PPU

	// last error from Step()
	err error
}

func (cpu *CPU) CurrentInstr() Instruction {
	code := cpu.loadU8(cpu.PC)
	if cpu.prefix {
		return extOps[code]
	}
	return ops[code]
}

func (cpu *CPU) Err() error {
	return cpu.err
}
func (cpu *CPU) WithLog(log *slog.Logger) *CPU {
	cpu.log = log
	return cpu
}

func (cpu *CPU) HL() uint16 { return concatU16(cpu.H, cpu.L) }
func (cpu *CPU) BC() uint16 { return concatU16(cpu.B, cpu.C) }
func (cpu *CPU) DE() uint16 { return concatU16(cpu.D, cpu.E) }
func (cpu *CPU) AF() uint16 { return concatU16(cpu.A, Register(cpu.F)) }

// loads and increments the progrm counter
func (cpu *CPU) load(addr uint16, dst any) {
	b, ok := cpu.Mem.Access(addr)
	// cpu.log.Debug("Accessing memory", "loc", hexstr(addr), "res", hexstr(b))
	if !ok {
		cpu.err = ErrNoMoreInstructions
	}
	// cpu.IncProgramCounter()
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
	cpu.load(addr, &lsb)
	cpu.load(addr+1, &msb)
	return concatU16(msb, lsb)
}
func (cpu *CPU) readI8(addr uint16) int8 {
	var value byte
	cpu.load(addr, &value)
	return int8(value)
}

// Writing to stack: use PushStack instead of this.
func (cpu *CPU) WriteMemory(addr uint16, value any) {
	if cpu.Mem == nil {
		panic("cpu.mem is nil")
	}
	switch value := value.(type) {
	case uint8:
		cpu.Mem.WriteData(addr, []byte{value})
	case uint16:
		msb, lsb := split(value)
		cpu.Mem.WriteData(addr, []byte{lsb, msb})
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

// 0th entry is the ... of stack
func (cpu *CPU) Stack() []uint16 {
	var stack []uint16
	for ptr := cpu.SP; ptr < 0xFFFF; ptr = ptr - 2 {
		lsb := cpu.readU8(ptr)
		msb := cpu.readU8(ptr - 1)
		stack = append(stack, concatU16(msb, lsb))
	}
	return stack
}

func (cpu *CPU) WithHook(hook func(cpu *CPU, loc int, instr Instruction, log *slog.Logger)) *CPU {
	cpu.hooks = append(cpu.hooks, hook)
	return cpu
}

func (cpu *CPU) Step() bool {
	cpu.InstrCount++
	if cpu.err != nil {
		return false
	}
	if cpu.limit > 0 && cpu.InstrCount >= cpu.limit {
		cpu.err = fmt.Errorf("instruction limit reached: %d", cpu.InstrCount)
		cpu.log.Warn("Instruction limit reached", "limit", cpu.limit, "count", cpu.InstrCount)
		return false
	}
	code := cpu.loadU8(cpu.PC)
	cpu.IncProgramCounter()
	if cpu.err != nil { // if loading next instruction failed, we'll stop
		return false
	}

	log := cpu.log.With("loc", hexstr(cpu.PC-1, 4), "name", name(code, cpu.prefix), "instr", hexstr(code), "pre", cpu.prefix)
	if !cpu.prefix && code == 0x00 && cpu.stopAtnop {
		cpu.err = ErrNoMoreInstructions
		return false // NOP command
	}
	if !cpu.prefix && code == 0x10 {
		cpu.err = ErrNoMoreInstructions
		return false
	}

	// log.Debug("instr start")

	var (
		instr Instruction
		ok    bool
	)
	if cpu.prefix {
		instr, ok = extOps[code]
		cpu.prefix = false
	} else {
		instr, ok = ops[code]
	}

	for _, hook := range cpu.hooks {
		hook(cpu, int(cpu.PC)-1, instr, log)
	}

	if !ok {
		cpu.err = fmt.Errorf("unknown code: %#x", code)
		return false
	}

	// Invariant: the PC is at the instruction at the start of the operation
	// and should end at the last instruction related to this command. For
	// very simple instructions (eg ADD A B), then the PC should not be changed
	// by op. However, for instructions that load data, the op should move the
	// PC towards the last instruction that is done by the op.

	instr.Exec(cpu)
	// log.Debug("instr done", "curr", fmt.Sprintf("%#x", cpu.PC))

	cpu.ppu.Step(cpu)

	return true
}

func (cpu *CPU) IncProgramCounter(src ...string) {
	cpu.PC++
	if cpu.log != nil {
		s := "N/A"
		if len(src) > 0 {
			s = src[0]
		}
		if false {
			cpu.log.Debug("PC increment", "new", fmt.Sprintf("%#x", cpu.PC), "src", s)
		}
	}
}

func (cpu *CPU) Dump(w io.Writer) {
	formatCycle := func(n int) string {
		if n < 1000000 {
			return strconv.Itoa(n)
		}
		return fmt.Sprintf("%.1fM", float64(n)/1000000)
	}
	fmt.Fprintf(w, "Cycles: %s\n", formatCycle(cpu.Cycles))
	fmt.Fprintf(w, "A:  %#02x     F:  %#02x     AF: %#04x\n", cpu.A, uint8(cpu.F), cpu.AF())
	fmt.Fprintf(w, "B:  %#02x     C:  %#02x     BC: %#04x\n", cpu.B, cpu.C, cpu.BC())
	fmt.Fprintf(w, "D:  %#02x     E:  %#02x	  DE: %#04x\n", cpu.D, cpu.E, cpu.DE())
	fmt.Fprintf(w, "H:  %#02x     L:  %#02x	  HL: %#04x\n", cpu.H, cpu.L, cpu.HL())
	fmt.Fprintf(w, "                          SP: %#04x\n", cpu.SP)
	fmt.Fprintf(w, "                          PC: %#04x\n", cpu.PC)
	// fmt.Fprintf(w, "HL: %#04x   BC: %#04x   DE: %#04x   AF: %#04x\n", cpu.HL(), cpu.BC(), cpu.DE(), cpu.AF())
}

type FlagRegister = Flags

func (f Flags) HasCarry() bool { return (uint8(f) & uint8(FLAGC)) > 0 }
func (f Flags) HasZero() bool  { return (uint8(f) & uint8(FLAGZ)) > 0 }
func (f Flags) HasHigh() bool  { return (uint8(f) & uint8(FLAGH)) > 0 }

func Run(cpu *CPU, mem *Memory, log *slog.Logger) error {
	cpu.Mem = mem
	cpu.log = log

	for cpu.Step() {
	}
	return cpu.Err()
}
