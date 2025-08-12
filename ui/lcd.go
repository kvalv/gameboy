package ui

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/kvalv/gameboy"
)

func NewLCD() *LCD {
	return &LCD{
		width:  160,
		height: 144,
	}
}

// Represents the LCD screen for the game boy
type LCD struct {
	width, height int

	cpu *gameboy.CPU
}

func (lcd *LCD) Draw(img *ebiten.Image) {
	buf := make([]byte, lcd.width*lcd.height*4)
	for x := range lcd.width {
		for y := range lcd.height {
			i := 4 * (x*lcd.height + y)
			buf[i] = uint8(x)
			buf[i+1] = uint8(y)
			buf[i+3] = 0xff

		}
	}
	img.WritePixels(buf)
}

func (lcd *LCD) Size() (int, int) {
	return lcd.width, lcd.height
}

func (lcd *LCD) pixelAt(x, y int) {
	// vram := lcd.cpu.Mem.VRAM()
}
