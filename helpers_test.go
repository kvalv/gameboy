package gameboy

import "testing"

func TestParts(t *testing.T) {

	t.Run("signed", func(t *testing.T) {
		i1 := int16(0x2233)
		a, b := msb(i1), lsb(i1)
		retr := int16(a)<<8 | int16(b)
		if retr != i1 {
			t.Fatalf("mismatch - got %#x, want %#x", retr, i1)
		}
	})
	t.Run("unsigned", func(t *testing.T) {
		i1 := uint16(0x2233)
		a, b := msb(i1), lsb(i1)
		retr := uint16(a)<<8 | uint16(b)
		if retr != i1 {
			t.Fatalf("mismatch - got %#x, want %#x", retr, i1)
		}
	})

	t.Run("add", func(t *testing.T) {
		check := func(wantVal any, wantFlag Flags, gotVal any, gotFlag FlagRegister) {
			if wantVal != gotVal {
				t.Fatalf("expected %d, got %d", wantVal, gotVal)
			}
			if uint8(gotFlag) != uint8(wantFlag) {
				t.Fatalf("flag mismatch")
			}
		}

		{
			val, fl := add(uint8(100), int8(-5))
			check(uint8(95), 0, val, fl)
		}
		{
			val, fl := add(uint8(4), int8(-6))
			check(uint8(254), FLAGC, val, fl)
		}
		{
			val, fl := add(uint8(4), int8(-4))
			check(uint8(0), FLAGZ, val, fl)
		}
		{
			val, fl := add(uint16(0xffaf), int8(-0x01))
			check(uint16(0xffae), 0, val, fl)
		}
	})

	t.Run("rotate", func(t *testing.T) {
		cases := []struct {
			n         uint8
			dir       int
			circular  bool
			currFlags Flags
			want      uint8
			wantCarry bool
		}{
			{
				n:         uint8(0b11100000),
				dir:       0,
				currFlags: 0,
				circular:  true,
				want:      0b11000001,
				wantCarry: true,
			},
			{
				n:         uint8(0b11100000),
				dir:       1,
				currFlags: 0,
				circular:  true,
				want:      0b01110000,
				wantCarry: false,
			},
			{
				n:         uint8(0b11100000),
				dir:       0,
				currFlags: FLAGC,
				circular:  false,
				want:      0b11000001,
				wantCarry: true,
			},
			{
				n:         uint8(0b11100000),
				dir:       0,
				currFlags: 0,
				circular:  false,
				want:      0b11000000,
				wantCarry: true,
			},
			{
				n:         uint8(0b11100000),
				dir:       1,
				currFlags: 0,
				circular:  false,
				want:      0b01110000,
				wantCarry: false,
			},
			{
				n:         uint8(0b11000000),
				dir:       0,
				currFlags: 0,
				circular:  false,
				want:      0b10000000,
				wantCarry: true,
			},
		}
		for i, tc := range cases {
			got, fl := rotate(tc.n, tc.dir, tc.currFlags, tc.circular)
			if got != tc.want {
				t.Fatalf("%d: rotate(%b, %d, %#b, %t) expect=%b, got=%b", i, tc.n, tc.dir, uint8(tc.currFlags), tc.circular, tc.want, got)
			}
			gotCarry := (Flags(fl) & FLAGC) > 0
			if gotCarry != tc.wantCarry {
				t.Fatalf("carry mismatch - want %t, got %t", tc.wantCarry, gotCarry)
			}
		}

	})

}
