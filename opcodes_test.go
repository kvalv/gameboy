package gameboy

import (
	"testing"
)

func TestAdd(t *testing.T) {
	cases := []struct {
		desc string
		op   Instruction
	}{
		{
			desc: "ADD_09",
			op:   ADD_09,
		},
	}

	for _, tc := range cases {
		t.Run(tc.desc, func(t *testing.T) {
		})
	}
}

func TestOpcodes(t *testing.T) {
	cases := []struct {
		desc      string
		lhs, rhs  uint8
		wantRes   uint8
		wantCarry bool
	}{
		{
			desc: "120+3",
			lhs:  120, rhs: 3,
			wantRes:   123,
			wantCarry: false,
		},
		{
			desc: "120+170",
			lhs:  120, rhs: 170,
			wantRes:   34,
			wantCarry: true,
		},
	}

	for _, tc := range cases {
		t.Run(tc.desc, func(t *testing.T) {
			cpu := &CPU{}
			cpu.A = tc.lhs
			cpu.B = tc.rhs

			ADD_80(cpu)
			if cpu.A != tc.wantRes {
				t.Fatalf("expected %d, got %d", tc.wantRes, cpu.A)
			}
			if cpu.F.HasCarry() != tc.wantCarry {
				t.Fatalf("expected carry %v, got %v", tc.wantCarry, cpu.F.HasCarry())
			}
		})
	}
}
