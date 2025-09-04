package cartridge

import "fmt"

type MBC0 struct {
}

func (mbc *MBC0) Write(rom cartridgeROM, ram cartridgeRAM, addr uint16, data byte) {
	switch {
	case addr < 0x8000: // rom banks, 32k size
		rom[addr] = data
	default:
		panic(fmt.Sprintf("MBC0: Cannot write to addr %#4x", addr))
	}
}

func (mbc *MBC0) Read(rom cartridgeROM, ram cartridgeRAM, addr uint16) byte {
	switch {
	case addr < 0x8000:
		return rom[addr]
	case 0xA000 <= addr && addr < 0xC000:
		return ram[addr-0xA000]
	}
	panic(fmt.Sprintf("MBC0: illegal memory access: %2x", addr))
}
