package gameboy

import (
	"bytes"
	"fmt"
)

// Header: 0x0100 - 0x014F
// Entry point: 0x100 - 0x103 (typically NOP, JP
// Logo: 0x104 - 0x133
// Title: 0x134 - 0x143
type Cartridge []byte

// type Cartridge struct {
// 	data []byte
// }

// 16kB blocks of memory. Switch out each block depending on value
// in a register, somewhere.

func NewCartridge(data []byte) Cartridge {
	return Cartridge(data)
}

func (c Cartridge) Title() string {
	b := c[0x0134:0x0143]
	return string(bytes.TrimRightFunc(b, func(r rune) bool {
		return r == 0
	}))
}

type MemoryBankController int

const (
	// 32kB - ROM is directly mapped 0000-7FFF
	MemoryBankControllerNone MemoryBankController = iota
	// 32kB banked ROM, 512kB ROM
	MemoryBankController1
	MemoryBankController2
	MemoryBankController3
	// few others too
)

// The Cartridge Type. Indicates what hardware is present on the cartridge.
// We only care about the MBC, so we'll return the MBC variant here.
func (c Cartridge) Type() MemoryBankController {
	switch code := c[0x0147]; code {
	case 0x00:
		return MemoryBankControllerNone
	case 0x01, 0x02, 0x03:
		return MemoryBankController1
	case 0x0f, 0x10, 0x11, 0x12, 0x13:
		return MemoryBankController3 // MBC3 + RAM + BATTERY
	default:
		panic(fmt.Sprintf("Unknown code %#2x", code))
	}
}

// How much ROM is present in the cartridge.
func (c Cartridge) ROMSize() uint { return 32 * kB * (1 << c[0x148]) }

// func (c Cartridge) ROMSize() uint { return 32 * kB * (1 << c[0x0148]) }
func (c Cartridge) RAMSize() uint {
	switch code := c[0x0149]; code {
	case 0x00:
		return 0
	case 0x01:
		panic("unused")
	case 0x02:
		return 8 * kB
	case 0x03:
		return 32 * kB
	case 0x04:
		return 128 * kB
	case 0x05:
		return 64 * kB
	default:
		panic(fmt.Sprintf("Unknown code %#2x", code))
	}
}
