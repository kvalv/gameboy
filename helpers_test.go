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

}
