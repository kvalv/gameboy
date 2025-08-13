package gameboy

// the following instructions are hard-coded. We only generate instructions for
// e.g. ADD and LD that have a billion variants. The others we'll just
// implement here.

// NOP     code=0x01
func NOP_01(cpu *CPU) {
	cpu.Cycles += 4
}

// RRCA   code=0x0F
func RRCA_0F(cpu *CPU) {
	panic("TODO")
}
