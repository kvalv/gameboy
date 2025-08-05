package gameboy

import "testing"

func TestBootLoader(t *testing.T) {
	// TODO: LD  ...

	cpu := CPU{}
	mem := NewMemory().Write(GetBootCode())
	if err := Run(&cpu, mem, logger(true)); err != nil && err != ErrNoMoreInstructions {
		t.Fatal(err)
	}
}
