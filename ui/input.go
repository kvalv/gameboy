package ui

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type Input struct {
	KeyN bool
	KeyQ bool
}

func NewInput() *Input {
	return &Input{}
}

func (i *Input) Update() {
	i.KeyQ = inpututil.IsKeyJustPressed(ebiten.KeyQ)
	i.KeyN = inpututil.IsKeyJustPressed(ebiten.KeyN)
}
