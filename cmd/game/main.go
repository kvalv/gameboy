package main

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/kvalv/gameboy/ui"
)

func main() {
	g := ui.NewGame()
	// ebiten.SetWindowSize(200, 200)
	ebiten.SetWindowTitle("Game Boy")
	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}
