package main

import "testing"

func TestGet(t *testing.T) {
	got := get("$18", true)
	want := "uint8(0x18)"
	if got != want {
		t.Fatalf("expected %q, got %q", want, got)
	}
}
