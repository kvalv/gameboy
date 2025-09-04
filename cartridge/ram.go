package cartridge

type cartridgeRAM []byte

const RAM_BANK_SIZE = 8 * kB

func (ram cartridgeRAM) Bank(i int) []byte {
	return ram[i*RAM_BANK_SIZE : (i+1)*RAM_BANK_SIZE]
}
