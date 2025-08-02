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

		chec := func(wantVal any, wantFlag Flags, gotVal any, gotFlag FlagRegister) {
			if wantVal != gotVal {
				t.Fatalf("expected %d, got %d", wantVal, gotVal)
			}
			if uint8(gotFlag) != uint8(wantFlag) {
				t.Fatalf("flag mismatch")
			}
		}

		{
			val, fl := add(uint8(100), int8(-5))
			chec(uint8(95), 0, val, fl)
		}
		{
			val, fl := add(uint8(4), int8(-6))
			chec(uint8(254), FLAGC, val, fl)
		}
		{
			val, fl := add(uint8(4), int8(-4))
			chec(uint8(0), FLAGZ, val, fl)
		}
		{
			val, fl := add(uint16(0xffaf), int8(-0x01))
			chec(uint16(0xffae), 0, val, fl)
		}

	})

}
