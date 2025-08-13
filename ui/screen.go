package ui

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/kvalv/gameboy"
)

// Represents the Screen screen for the game boy
type Screen struct {
	cpu *gameboy.CPU
}

func NewScreen(cpu *gameboy.CPU) *Screen {
	return &Screen{cpu: cpu}
}

func (s *Screen) Size() (int, int) {
	return 160, 144
}

func (s *Screen) Draw(img *ebiten.Image) {
	// W, H := s.Size()
	mem := s.cpu.Mem
	vram := mem.VRAM()

	// draw tile window 1
	// 32 x 32 view, and each tile is 8 x 8 -> 256 x 256
	// 32 x 32 = 1024 = 1kB
	// each tile is 16 bytes, and tile data1 has 4096 bytes -> 256 tiles
	// .. so each byte in tile view 1 is just an index into one of the 256 tiles

	topLeftX := int(mem.SCX())
	topLeftY := int(mem.SCY())
	// botRightX := uint8((int(mem.SCX()) + 159) % 256)
	// botRightY := uint8((int(mem.SCY()) + 143) % 256)

	for i, index := range vram.TileView1 {
		tile := vram.Tile(gameboy.TilesetBackground, index)

		x0 := (i % 32) * gameboy.TILE_WIDTH_PX
		y0 := (i / 32) * gameboy.TILE_WIDTH_PX

		for x := range 8 {
			for y := range 8 {
				color := tile.PixelAt(x, y, PALETTE)
				img.Set(topLeftX+x0+x, topLeftY+y0+y, color)
			}
		}
	}

	// TODO: foreground...

}
