package gameboy

import "testing"

func TestBootLoader(t *testing.T) {

	cpu := CPU{}
	mem := NewMemory().Write(BootCode)
	if err := Run(&cpu, mem, logger(true)); err != nil && err != ErrNoMoreInstructions {
		t.Fatal(err)
	}
}
