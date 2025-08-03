package main

import "testing"

func TestXor(t *testing.T) {
	a := 0b11110000
	b := 0b11000011
	w := 0b00110011
	if a^b != w {
		t.Fatalf("got %b", w)
	}
}
