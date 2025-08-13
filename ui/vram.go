package ui

import (
	"image/color"
	"slices"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/kvalv/gameboy"
)

// - 0x8000 - 0x8fff - 4kB - Tile data 1
// - 0x8800 - 0x97ff - 4kB - tile data 2 (note overlap with above)
// - 0x9800 - 0x9bff - 1kB - Tile view 1
// - 0x9c00 - 0x9fff - 1kB - Tile view 2

type DisplayVRAM struct {
	mem          *gameboy.Memory
	nrows, ncols int
}

func NewDisplayVRAM(mem *gameboy.Memory) *DisplayVRAM {
	return &DisplayVRAM{
		mem:   mem,
		nrows: 16,
		ncols: 16,
	}
}

func (d *DisplayVRAM) Size() (width, height int) {
	// 256 tiles? :O
	// 8 rows each... 32
	pad := 0
	sz := 8

	width = d.ncols * (pad + sz)
	height = d.nrows * (pad + sz)

	height = 0
	height += d.nrows * 8 // tile data 1
	height += 1           // pad
	height += d.nrows * 9 // tile data 1
	height += 1           // pad

	return
}

func (d *DisplayVRAM) Draw(img *ebiten.Image) {
	vram := d.mem.VRAM()

	// W, H := d.Size()

	var palette = PALETTE

	img.Fill(color.RGBA{0x33, 0x33, 0x33, 0xFF})

	// something something render the tiles using the tiledata..

	var i int
	for data := range slices.Chunk(vram.TileData1, 16) {
		tile := gameboy.Tile(data)
		for x := range 8 {
			for y := range 8 {
				col := tile.PixelAt(x, y, palette)
				x0 := (i % d.ncols) * 8
				y0 := (i / d.ncols) * 8
				img.Set(x+x0, y+y0, col)
			}
		}
		i++
	}

	// 16 x 16
	// offsetY = 16 * 8 + 1

	{
		var i int
		offsetY := 8*d.nrows + 1
		for data := range slices.Chunk(vram.TileData2, 16) {
			tile := gameboy.Tile(data)
			for x := range 8 {
				for y := range 8 {
					col := tile.PixelAt(x, y, palette)
					x0 := (i % d.ncols) * 8
					y0 := (i/d.ncols)*8 + offsetY
					img.Set(x+x0, y+y0, col)
				}
			}
			i++
		}
	}

}
