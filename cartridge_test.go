package gameboy

import (
	_ "embed"
	"testing"

	"github.com/stretchr/testify/require"
)

//go:embed pokemon.gb
var POKEMON []byte

//go:embed tetris.gb
var TETRIS []byte

func TestTitle(t *testing.T) {
	got := NewCartridge(POKEMON).Title()
	want := "POKEMON BLUE"

	if got != want {
		t.Fatalf("title mismatch: want=%q, got=%q", want, got)
	}
}

func TestTetris(t *testing.T) {
	req := require.New(t)
	cart := NewCartridge(TETRIS)
	req.Equal(MemoryBankControllerNone, cart.Type())
	req.Equal(32*kB, cart.ROMSize())
	req.Equal(uint(0), cart.RAMSize())
}
