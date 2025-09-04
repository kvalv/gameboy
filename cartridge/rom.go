package cartridge

import (
	"bytes"
	"fmt"
)

const ROM_BANK_SIZE = 16 * kB

type cartridgeROM []byte

func newROM(data []byte) cartridgeROM {
	return data
}

func (c cartridgeROM) Title() string {
	b := c[0x0134:0x0143]
	return string(bytes.TrimRightFunc(b, func(r rune) bool {
		return r == 0
	}))
}

// The Cartridge Type. Indicates what hardware is present on the cartridge.
// We only care about the MBC, so we'll return the MBC variant here.
func (c cartridgeROM) Type() MBCType {
	if c == nil {
		return Type0
	}
	switch code := c[0x0147]; code {
	case 0x00:
		return Type0
	case 0x01, 0x02, 0x03:
		return Type1
	case 0x0f, 0x10, 0x11, 0x12, 0x13:
		return Type3 // MBC3 + RAM + BATTERY
	default:
		panic(fmt.Sprintf("Unknown code %#2x", code))
	}
}

// How much ROM is present in the cartridge.
func (c cartridgeROM) ROMSize() uint {
	if c == nil {
		return 0
	}
	return 32 * kB * (1 << c[0x148])
}

// func (c Cartridge) ROMSize() uint { return 32 * kB * (1 << c[0x0148]) }
func (c cartridgeROM) RAMSize() uint {
	if c == nil {
		return 0
	}
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

// Returns the number of RAM banks for this ROM.
// Basically the ROM size divided by size per bank
func (c cartridgeROM) BankCount() int {
	return int(c.ROMSize()) / ROM_BANK_SIZE
}

func (rom cartridgeROM) Bank(i int) []byte {
	return rom[i*ROM_BANK_SIZE : (i+1)*ROM_BANK_SIZE]
}
