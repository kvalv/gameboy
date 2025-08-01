package gameboy

import (
	"errors"
	"fmt"
	"log/slog"
	"os"
	"testing"
)

const (
	INSTR_STOP = 0x01
)

func TestInstructions(t *testing.T) {

	cases := []struct {
		desc    string
		cpu     *CPU                               // initial state of cpu
		initMem func(m *Memory)                    // instructions to run
		check   func(t *testing.T, cpu *CPUHelper) // post-check func
	}{
		{
			desc: "ADD A B",
			cpu: &CPU{
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
			cpu: &CPU{
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
			cpu: &CPU{
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
			cpu: &CPU{
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
			cpu: &CPU{
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
			cpu: &CPU{
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
			cpu: &CPU{
				H: 0x11,
				L: 0x22,
			},
			initMem: func(m *Memory) {
				m.WriteInstr(0x34)
				m.WriteInstr(INSTR_STOP)
				m.WriteByteAt(0x1122, 0x33)
			},
			check: func(t *testing.T, cpu *CPUHelper) {
				cpu.ExpectByte(0x1122, 0x34)
			},
		},
		{
			desc: "INC SP",
			cpu:  &CPU{SP: 0x0011},
			initMem: func(m *Memory) {
				m.WriteInstr(0x33)
				m.WriteInstr(INSTR_STOP)
			},
			check: func(t *testing.T, cpu *CPUHelper) { cpu.ExpectSP(0x0012) },
		},
		{
			desc: "DEC HL",
			cpu:  &CPU{},
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
			cpu:  &CPU{},
			initMem: func(m *Memory) {
				m.WriteInstr(0x06)
				m.Write(0xab) // the data
				m.WriteInstr(INSTR_STOP)
			},
			check: func(t *testing.T, cpu *CPUHelper) {
				cpu.ExpectB(0xab)
			},
		},
	}

	for _, tc := range cases {
		t.Run(tc.desc, func(t *testing.T) {
			cpu := tc.cpu
			var mem Memory
			tc.initMem(&mem)
			if err := Run(cpu, &mem, logger()); err != nil && !errors.Is(err, ErrNoMoreInstructions) {
				fmt.Println("")
				cpu.Dump(os.Stderr)
				t.Fatalf("failed to run: %v", err)
			}

			tc.check(t, &CPUHelper{
				t:   t,
				CPU: tc.cpu,
			})

		})
	}
}

func logger() *slog.Logger {
	return slog.New(slog.NewTextHandler(os.Stderr, &slog.HandlerOptions{
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
func (cpu *CPUHelper) ExpectByte(offset uint16, want byte) {
	got, ok := cpu.mem.Access(offset)
	if !ok {
		cpu.t.Fatalf("illegal offset %d", offset)
	}
	if got != want {
		cpu.t.Fatalf("ExpectByte: want=%#x, got=%#x", want, got)
	}
}
