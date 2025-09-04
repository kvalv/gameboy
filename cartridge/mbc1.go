package cartridge

import (
	"fmt"
	"log/slog"
)

// memory banking mode, specific for MBC1 (and maybe others?)
type bmode int

const (
	modeSimple   bmode = iota // More RAM
	modeAdvanced              // More ROM, less RAM. 2-bit value in 4000-6000 register
)

func (m bmode) String() string {
	if m == modeSimple {
		return "simple"
	}
	return "advanced"
}

type MBC1 struct {
	// defaults to ROM mode
	// If ramMode ->  512kB ROM, 32kB RAM
	// else       -> 2048kB ROM,  8kB RAM
	// ramMode bool
	mode bmode

	romIdxLo int // first 5 bits of rom index
	romIdxHi int // last 2 bits of rom index, if needed
	ramIdx   int // index of RAM bank

	ramEnabled bool

	log *slog.Logger
}

func (mbc *MBC1) romIndex() int {
	return mbc.romIdxLo + mbc.romIdxHi<<5
}

func (mbc *MBC1) Write(rom cartridgeROM, ram cartridgeRAM, addr uint16, data byte) {
	log := mbc.log

	// https://gbdev.io/pandocs/MBC1.html#mbc1
	switch {
	case addr < 0x2000: // RAM enable
		mask := uint8(0x0a)
		mbc.ramEnabled = (data & mask) == mask
	case 0x2000 <= addr && addr < 0x4000: // ROM bank number
		// only care about the first 5 bits
		i := data & 0x1f

		// If this register is set to $00, it behaves as if it is set to $01.
		i = max(i, 1)

		// If the ROM Bank Number is set to a higher value than the number of banks in the cart, the bank number is masked to the required number of bits. e.g. a 256 KiB cart only needs a 4-bit bank number to address all of its 16 banks, so this register is masked to 4 bits. The upper bit would be ignored for bank selection.
		i = i % byte(rom.BankCount())

		mbc.romIdxLo = int(i)
		log.Debug("new ROM bank", slog.Int("romIdx", mbc.romIdxLo), slog.Int("romIdxHi", mbc.romIdxHi), slog.Int("ramIdx", mbc.ramIdx), slog.Bool("lowBits", false))

	case 0x4000 <= addr && addr < 0x6000: // RAM bank number - or - upper bits of ROM bank number
		// only care about the first 2 bits
		index := int(data) & 0x03
		switch mbc.mode {
		case modeSimple:
			if index > rom.BankCount() {
				return
			}
			mbc.ramIdx = index
			log.Info("new RAM bank", slog.Int("index", mbc.romIndex()))
			return
		case modeAdvanced:
			if (mbc.romIdxLo + (index << 5)) > rom.BankCount() {
				return
			}
			mbc.romIdxHi = index
			log.Info("new ROM bank", slog.Int("index", mbc.romIndex()))
		default:
			panic("Unknown mode")
		}

	case 0x6000 <= addr && addr < 0x8000: // Banking Mode Select
		if data != 0 && data != 1 {
			panic(fmt.Sprintf("write to memory - expected only 0 or 1, got %x", data))
		}

		// If the cart is not large enough to use the 2-bit register (≤ 8 KiB RAM and ≤ 512 KiB ROM) this mode select has no observable effect.
		if mbc.mode == modeSimple && rom.RAMSize() <= 8*kB {
			log.Info("Received write - Banking Mode Select - no effect as size is too smal")
		}
		if mbc.mode == modeAdvanced && rom.ROMSize() <= 512*kB {
			log.Info("Received write - Banking Mode Select - no effect as size is too smal")
		}

		if data == 1 {
			mbc.mode = modeAdvanced
		} else {
			mbc.mode = modeSimple
		}
		log.Info("Ram mode switch", slog.String("mode", mbc.mode.String()))
	}
}

func (mbc *MBC1) Read(rom cartridgeROM, ram cartridgeRAM, addr uint16) byte {
	switch {
	case addr < 0x4000: // ROM bank 0, always fixed and present
		return rom[addr]
	case 0x4000 <= addr && addr < 0x8000: // Switchable ROM Bank
		return rom.Bank(mbc.romIndex())[addr-0x4000]
	case 0xA000 <= addr && addr < 0xC000: // Cartridge RAM
		if !mbc.ramEnabled {
			return 0xFF // any random value, really
		}

		bank := ram.Bank(mbc.ramIdx)
		return bank[addr-0xA000]
	}

	panic(fmt.Sprintf("Not implemented: %#4x", addr))
}
