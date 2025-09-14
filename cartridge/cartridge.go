package cartridge

import (
	"fmt"
	"log/slog"
)

// Specifically, this is the CartridgeROM ROM

type Cartridge struct {
	mbc MemoryBankController
	rom cartridgeROM
	ram cartridgeRAM
}

func New(data []byte, log ...*slog.Logger) Cartridge {

	var lg *slog.Logger
	if len(log) == 0 {
		lg = slog.New(slog.DiscardHandler)
	} else {
		lg = log[0]
	}

	if len(data) == 0 {
		data = make([]byte, 64*kB) // empty ROM
	}
	rom := newROM(data)

	// To consider: random data
	ram := make([]byte, rom.RAMSize())
	var mbc MemoryBankController
	switch tp := rom.Type(); tp {
	case Type0:
		mbc = &MBC0{}
	case Type1:
		mbc = &MBC1{log: lg}
	case Type3:
		fmt.Printf("warning: MBC3 not implemented, using MBC1 instead\n")
		mbc = &MBC1{log: lg}
	}

	if len(rom) == 0 {
		panic("empty ROM")
	}
	// if len(ram) == 0 {
	// 	panic("empty RAM")
	// }

	return Cartridge{
		mbc: mbc,
		rom: rom,
		ram: ram,
	}
}

func (cart Cartridge) Write(addr uint16, data byte) {
	cart.mbc.Write(cart.rom, cart.ram, addr, data)
}
func (cart Cartridge) Read(addr uint16) byte {
	return cart.mbc.Read(cart.rom, cart.ram, addr)
}

// Mode of the memory bank controller
type MBCType int

const (
	// 32kB - ROM is directly mapped 0000-7FFF
	Type0 MBCType = iota
	// Has two modes:
	// - 2MB ROM, 8kB RAM [mode=0]
	// - 512kB ROM, 32kB RAM [mode=1]
	// Writing 0 or 1 into 0x6000-0x7FFF switches between the two modes
	// respectively.
	Type1
	Type2
	Type3
	// few others too
)
