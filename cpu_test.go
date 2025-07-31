package gameboy

import (
	"errors"
	"log/slog"
	"os"
	"testing"
)

func TestCPUAdd(t *testing.T) {
	// cpu := NewCPU()
	// var mem Memory
	// cpu.A = 5
	// cpu.B = 2
	// want: 7
	// mem.Write([]byte{
	// 	0x80, // ADD A B
	// })

	cases := []struct {
		desc    string
		cpu     *CPU                               // initial state of cpu
		initMem func(m *Memory)                    // instructions to run
		check   func(t *testing.T, cpu *CPUHelper) // post-check func
	}{
		{
			desc: "ADD A B",
			cpu:  &CPU{A: 5, B: 2},
			initMem: func(m *Memory) {
				m.Write([]byte{0x80})
			},
			check: func(t *testing.T, cpu *CPUHelper) {
				cpu.ExpectA(7)
			},
		},
		{
			desc: "ADD A HL",
			cpu:  &CPU{A: 1, H: 0x00, L: 0x0a},
			initMem: func(m *Memory) {
				m.WriteU8(0x86) // ADD A,HL
				m.WriteU8(0x01) // STOP
				m.WriteAt(0x000a, []byte{0x44})
			},
			check: func(t *testing.T, cpu *CPUHelper) {
				cpu.ExpectA(0x45)
			},
		},
	}

	for _, tc := range cases {
		t.Run(tc.desc, func(t *testing.T) {
			var mem Memory
			tc.initMem(&mem)
			if err := Run(tc.cpu, &mem, logger()); err != nil && !errors.Is(err, ErrNoMoreInstructions) {
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
	return slog.New(slog.NewTextHandler(os.Stderr, nil))
}

type CPUHelper struct {
	t *testing.T
	*CPU
}

func (cpu *CPUHelper) ExpectA(want uint8) {
	if cpu.A != want {
		cpu.t.Fatalf("want=%d, got=%d", want, cpu.A)
	}
}
