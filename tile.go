package gameboy

import "image/color"

type Tile []byte

const (
	TILE_DATA_SIZE = 16
	TILE_WIDTH_PX  = 8
)

func (t Tile) PixelAt(x, y int, palette []color.RGBA) color.RGBA {
	// for each line, first byte is MSB, second byte is LSB in the palette

	// read bit at location n
	bit := func(b byte, n int) int { // 0 or 1
		if (b & (1 << n)) > 0 {
			return 1
		}
		return 0
	}

	firstByte := t[2*y]
	secondByte := t[2*y+1]

	// leftmost bit is bit 7, so we'll do 8 - x
	lsb := bit(firstByte, 8-x)
	msb := bit(secondByte, 8-x)
	color := palette[2*lsb+msb]
	return color
}
