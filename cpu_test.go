package gameboy

import (
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
				m.WriteData(0x000a, []byte{0x44})
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
				m.WriteData(0x1122, []byte{0x02})
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
				m.WriteByteAt(0x1122, 0x33)
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
				m.WriteByteAt(0x1122, 0x33)
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
				m.WriteByteAt(0x1122, 0x00)
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
			desc: "LD BC,n16 0x01",
			cpu:  CPU{},
			initMem: func(m *Memory) {
				m.Write(0x01, 0x11, 0x22, INSTR_STOP)
			},
			check: func(t *testing.T, cpu *CPUHelper) {
				cpu.ExpectBC(0x1122)
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
			desc: "LD (a16),SP 0x08",
			cpu: CPU{
				SP: 0x3344,
			},
			initMem: func(m *Memory) {
				m.Write(0x08, 0x11, 0x22, INSTR_STOP)
			},
			check: func(t *testing.T, cpu *CPUHelper) {
				cpu.ExpectMem(0x1122, 0x44) // lsb
				cpu.ExpectMem(0x1123, 0x33) // msb
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
				m.WriteByteAt(0x1122, 0x33)
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
				m.WriteByteAt(0x1122, 0x01)
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
				m.WriteByteAt(0x1122, 0x01)
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
				m.WriteByteAt(0x1122, 0x33)
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
				m.WriteByteAt(0x1122, 0x01)
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
				m.WriteByteAt(0x1122, INSTR_STOP)
			},
			check: func(t *testing.T, cpu *CPUHelper) {
				cpu.ExpectPC(0x1123) // 0x1122 + 1, since we're reading another instruction before stopping
				cpu.ExpectSP(0xFFFD)

				// read instr   1
				// read a16     2
				// 3 instructions done, so the NEXT is at 0x03
				cpu.ExpectPeekStack(uint16(0x03))
			},
		},
		{
			desc: "CALL NZ,a16 0xC4 - flag zero",
			cpu: CPU{
				F: FlagRegister(FLAGZ),
			},
			initMem: func(m *Memory) {
				m.Write(0xC4, 0x22, 0x11) // stored as lsb, msb -- so bits reversed
				m.WriteByteAt(0x1122, INSTR_STOP)
			},
			check: func(t *testing.T, cpu *CPUHelper) {
				// instr + a16 + stop
				// 1     + 2   + 1
				// -> 4 instructions ran, so the next is at 0x04
				cpu.ExpectPC(0x04)
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
			initCPU(cpu)
			mem := NewMemory()
			tc.initMem(mem)
			if err := Run(cpu, mem, logger(tc.debug)); err != nil && !errors.Is(err, ErrNoMoreInstructions) {
				fmt.Println("")
				t.Fatalf("failed to run: %v", err)
			}

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
	if cpu.A != want {
		cpu.t.Fatalf("want=%#x, got=%#x", want, cpu.A)
	}
}
func (cpu *CPUHelper) ExpectB(want uint8) {
	if cpu.B != want {
		cpu.t.Fatalf("want=%#x, got=%#x", want, cpu.B)
	}
}
func (cpu *CPUHelper) ExpectC(want uint8) {
	if cpu.C != want {
		cpu.t.Fatalf("want=%#x, got=%#x", want, cpu.C)
	}
}
func (cpu *CPUHelper) ExpectD(want uint8) {
	if cpu.D != want {
		cpu.t.Fatalf("want=%#x, got=%#x", want, cpu.D)
	}
}
func (cpu *CPUHelper) ExpectE(want uint8) {
	if cpu.E != want {
		cpu.t.Fatalf("want=%#x, got=%#x", want, cpu.E)
	}
}
func (cpu *CPUHelper) ExpectH(want uint8) {
	if cpu.H != want {
		cpu.t.Fatalf("want=%#x, got=%#x", want, cpu.H)
	}
}
func (cpu *CPUHelper) ExpectL(want uint8) {
	if cpu.L != want {
		cpu.t.Fatalf("want=%#x, got=%#x", want, cpu.L)
	}
}
func (cpu *CPUHelper) ExpectPC(want uint16) {
	if cpu.PC != want {
		cpu.t.Fatalf("want=%#x, got=%#x", want, cpu.PC)
	}
}
func (cpu *CPUHelper) ExpectHL(want uint16) {
	if cpu.HL() != want {
		cpu.t.Fatalf("want=%#x, got=%#x", want, cpu.HL())
	}
}
func (cpu *CPUHelper) ExpectBC(want uint16) {
	if cpu.BC() != want {
		cpu.t.Fatalf("want=%#x, got=%#x", want, cpu.BC())
	}
}
func (cpu *CPUHelper) ExpectDE(want uint16) {
	if cpu.DE() != want {
		cpu.t.Fatalf("want=%#x, got=%#x", want, cpu.DE())
	}
}
func (cpu *CPUHelper) ExpectSP(want uint16) {
	if cpu.SP != want {
		cpu.t.Fatalf("want=%#x, got=%#x", want, cpu.SP)
	}
}
func (cpu *CPUHelper) ExpectFlagCarry() {
	if !cpu.F.HasCarry() {
		cpu.t.Fatalf("expected carry flag to be set, but it's not")
	}
}
func (cpu *CPUHelper) ExpectFlagCarryUnset() {
	if cpu.F.HasCarry() {
		cpu.t.Fatalf("carry flag is set, expected unset")
	}
}
func (cpu *CPUHelper) ExpectFlagZero() {
	if !cpu.F.HasZero() {
		cpu.t.Fatalf("expected zero flag to be set, but it's not")
	}
}
func (cpu *CPUHelper) ExpectMem(offset uint16, want byte) {
	got, ok := cpu.mem.Access(offset)
	if !ok {
		cpu.t.Fatalf("illegal offset %d", offset)
	}
	if got != want {
		cpu.t.Fatalf("ExpectByte: want=%#x, got=%#x", want, got)
	}
}
func (cpu *CPUHelper) ExpectPeekStack(want any) {
	t := cpu.t
	// in other words: MSB is the HIGH address, LSB is the LOW address
	msb, ok := cpu.mem.Access(cpu.SP + 1)
	if !ok {
		t.Fatalf("failed to read msb at %#v", cpu.SP)
	}
	lsb, ok := cpu.mem.Access(cpu.SP)
	if !ok {
		t.Fatalf("failed to read lsb at %#v", cpu.SP+1)
	}
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
		b, ok := cpu.mem.Access(i)
		if !ok {
			t.Fatalf("failed to read stack at %#x", i)
		}
		fmt.Fprintf(w, "  %#04X: %02x\n", i, b)
	}
	fmt.Fprintln(w, "=== End of stack dump")
}
