package gameboy

// Pixel-Processing Unit -- the thing that is responsible for drawing on the screen
// https://gbdev.io/pandocs/Rendering.html
type PPU struct {
	prev int
}

func (p *PPU) Step(cpu *CPU) {

	// ppuFreq := 59.7 // fps

	// period := CPU_FREQUENCY / ppuFreq
	// n := cpu.Cycles / int(period)

	// increments every 456 cycles
	if cpu.Cycles-p.prev < 456 {
		return
	}

	p.prev = cpu.Cycles
	// otherwise increase

	// 0 - 143 -> drawing lines
	// 144 - 153 -> vblank
	LY := (cpu.Mem.LY() + 1) % (154)
	cpu.Mem.WriteAt(ADDR_LY, LY)
}
