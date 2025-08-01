package gameboy

import (
	"errors"
	"fmt"
	"log/slog"
	"os"
	"testing"
)

func TestCPUAdd(t *testing.T) {

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
				m.Append([]byte{0x80})
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
				m.WriteInstr(0x01)
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
func (cpu *CPUHelper) ExpectHL(want uint16) {
	if cpu.HL() != want {
		cpu.t.Fatalf("want=%#x, got=%#x", want, cpu.HL())
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
