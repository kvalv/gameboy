package gameboy

// the following instructions are hard-coded. We only generate instructions for
// e.g. ADD and LD that have a billion variants. The others we'll just
// implement here.

// NOP     code=0x01
func NOP_01(cpu *CPU) {
	cpu.cycles += 4
}

// RRCA   code=0x0F
func RRCA_0F(cpu *CPU) {
	panic("TODO")
}

// STOP    code=0x10
func STOP_10(cpu *CPU) {
	// this is really documented "TODO" so I don't know what heppens here..
	cpu.err = ErrNoMoreInstructions
	cpu.cycles += 4
}
