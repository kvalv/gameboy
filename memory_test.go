package gameboy

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestDisableROM(t *testing.T) {
	req := require.New(t)
	mem := NewMemory(nil)
	req.True(mem.BootActive(), "boot should initially be active")
	mem.WriteAt(0xFF50, 0x01)
	req.False(mem.BootActive(), "boot should be disabled")
}
