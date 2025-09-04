package cartridge

// A Memory Bank Controller (MBC) manages access to memory on
// the Cartridge.
// The Game Boy only has 16-bit addressing, meaning the maximum
// size of a game is 64kB without any tricks. The trick is to
// use a MBC that "swaps" out memory segments.
// Each bank is 16kB
type MemoryBankController interface {
	Write(rom cartridgeROM, ram cartridgeRAM, addr uint16, data byte)
	Read(rom cartridgeROM, ram cartridgeRAM, addr uint16) byte
}

const kB = 1024

func numRAMBanks(cart cartridgeROM) uint8 {
	nbanks := cart.RAMSize() / RAM_BANK_SIZE
	if nbanks > 256 {
		panic("expected limited to 256")
	}
	return uint8(nbanks)
}
func numROMBanks(cart cartridgeROM) uint8 {
	nbanks := cart.ROMSize() / ROM_BANK_SIZE
	if nbanks > 256 {
		panic("expected limited to 256")
	}
	return uint8(nbanks)
}
