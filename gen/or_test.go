package main

import "testing"

func TestOr(t *testing.T) {
	a := 0b11110000
	b := 0b11000011
	w := 0b11110011
	if a|b != w {
		t.Fatalf("got %b", w)
	}
}
