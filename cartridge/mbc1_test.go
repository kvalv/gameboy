package cartridge

import (
	"log/slog"
	"os"
	"testing"
)

// code RAM size  banks
// $00	0	      No RAM
// $01	–	      -
// $02	8 KiB	  1 bank
// $03	32 KiB	  4 banks of 8 KiB each
// $04	128 KiB	  16 banks of 8 KiB each
// $05	64 KiB	  8 banks of 8 KiB each
//
// code ROM size banks
// $00	32 KiB	 2 (no banking)
// $01	64 KiB	 4
// $02	128 KiB	 8
// $03	256 KiB	 16
// $04	512 KiB	 32
// $05	1 MiB	 64
// $06	2 MiB	 128
// $07	4 MiB	 256
// $08	8 MiB	 512
func newTestCart(t *testing.T, romCode, ramCode int) cartridgeHelper {
	data := make([]byte, 32*kB*(1<<romCode))
	data[0x147] = 0x03 // MBC1 + RAM + Battery
	data[0x148] = uint8(romCode)
	data[0x149] = uint8(ramCode)

	cart := New(data, logger())
	for i := range cart.rom.BankCount() {
		if i == 0 {
			continue
		}
		bank := cart.rom.Bank(i)
		for j := range bank {
			bank[j] = byte(i)
		}
	}
	return cartridgeHelper{cart, t}
}

func TestROMIndex(t *testing.T) {
	cart := newTestCart(t, 5, 3) // 1024
	cart.WriteROMBankLow(2)

	cart.ExpectRead(0x147, 0x03)
	cart.ExpectRead(0x4001, 0x02)

	cart.WriteROMBankLow(3)
	cart.Read(0x4000) // switch to bank 3
	cart.ExpectRead(0x4000, 0x03)

	// RAM
	cart.ExpectRead(0xa000, 0xFF) // RAM disabled

	cart.WriteROMBankLow(0xa3) // upper bit ignored, not bank 0
	cart.ExpectROMIndex(3)

	cart.WriteROMBankLow(0) // +1
	cart.ExpectROMIndex(1)

	cart.WriteROMBankLow(1)
	cart.ExpectROMIndex(1)

	cart.WriteROMBankHigh(1) // doesn't matter, RAM mode
	cart.ExpectROMIndex(1)

	cart.WriteBankMode(modeAdvanced)
	cart.WriteROMBankHigh(1)
	cart.ExpectROMIndex(0b_0010_0001)
}

func TestReadRAM(t *testing.T) {
	cart := newTestCart(t, 5, 3) // 1024
	cart.WriteRAM(0, func(i int, curr byte) byte {
		return byte(i + 1)
	})

	cart.ExpectRead(0xA000, 0xFF) // ram disabled
	cart.EnableRAM()
	cart.ExpectRead(0xA000, 0x01)
	cart.ExpectRead(0xA001, 0x02)

	cart.WriteRAMBank(100) // If neither ROM nor RAM is large enough, setting this register does nothing.
	cart.ExpectRead(0xA000, 0x01)

	cart.WriteRAMBank(1)
	cart.ExpectRead(0xA000, 0x00)

	cart.DisableRAM()
	cart.ExpectRead(0xA000, 0xFF)

}

func TestRAMTrick(t *testing.T) {
	// "Even with smaller ROMs that use less than 5 bits for bank selection, the full 5-bit register is still compared for the bank 00→01 translation logic. [...]"
	cart := newTestCart(t, 3, 4)
	cart.WriteROMBankLow(0x10)
	cart.ExpectROMIndex(0)
}

// cartridgeHelper is a Cartridge with various (test) helper functions
// connected to it.
type cartridgeHelper struct {
	Cartridge
	t *testing.T
}

func (c *cartridgeHelper) WriteROMBankLow(i int)  { c.Write(0x2000, uint8(i)) }
func (c *cartridgeHelper) WriteROMBankHigh(i int) { c.Write(0x4000, uint8(i)) }
func (c *cartridgeHelper) WriteRAMBank(i int)     { c.Write(0x4000, uint8(i)) }
func (c *cartridgeHelper) EnableRAM()             { c.Write(0x0000, 0x0a) }
func (c *cartridgeHelper) DisableRAM()            { c.Write(0x0000, 0x00) }
func (c *cartridgeHelper) WriteBankMode(m bmode)  { c.Write(0x6000, uint8(m)) }

func (c *cartridgeHelper) ExpectRead(addr uint16, want byte) {
	t := c.t
	t.Helper()
	got := c.Read(addr)
	if got != want {
		t.Fatalf("ExpectRead: want=%#2x got=%#2x", want, got)
	}
}
func (c *cartridgeHelper) ExpectROMIndex(want int) {
	t := c.t
	t.Helper()
	x, ok := c.mbc.(*MBC1)
	if !ok {
		t.Fatal("not mbc1")
	}
	if got := x.romIndex(); got != want {
		t.Fatalf("ROM index mismatch: want=%d, got=%d", want, got)
	}
}

// Accepts a callback to write to a particular RAM bank. Used during testing
func (c *cartridgeHelper) WriteRAM(i int, f func(i int, curr byte) byte) {
	bank := c.ram.Bank(i)
	for i, curr := range bank {
		bank[i] = f(i, curr)
	}
}

// Accepts a callback to write to a particular ROM bank. Used during testing
func (c *cartridgeHelper) WriteROM(i int, f func(i int, curr byte) byte) {
	bank := c.rom.Bank(i)
	for i, curr := range bank {
		bank[i] = f(i, curr)
	}
}

func logger(debug ...bool) *slog.Logger {
	lev := slog.LevelInfo
	if len(debug) > 0 && debug[0] {
		lev = slog.LevelDebug
	}
	return slog.New(slog.NewTextHandler(os.Stderr, &slog.HandlerOptions{
		Level: lev,
		ReplaceAttr: func(groups []string, a slog.Attr) slog.Attr {
			if a.Key == slog.TimeKey || a.Key == slog.LevelKey {
				return slog.Attr{} // remove time attribute
			}
			return a
		},
	}))
}
