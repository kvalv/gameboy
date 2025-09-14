package cartridge

import (
	_ "embed"
	"testing"

	"github.com/stretchr/testify/require"
)

//go:embed data/pokemon.gb
var POKEMON []byte

//go:embed data/tetris.gb
var TETRIS []byte

//go:embed data/ld.gb
var LD []byte

func TestTitle(t *testing.T) {
	got := New(POKEMON).rom.Title()
	want := "POKEMON BLUE"

	if got != want {
		t.Fatalf("title mismatch: want=%q, got=%q", want, got)
	}
}

func TestTetris(t *testing.T) {
	req := require.New(t)
	cart := New(TETRIS)
	req.Equal("TETRIS", cart.rom.Title())
}
