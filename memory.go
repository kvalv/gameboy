package gameboy

import (
	"encoding/hex"
	"fmt"
	"io"
	"slices"
)

type Memory struct {
	i    int
	data []byte
	cart Cartridge
	boot []byte
}

func NewMemory(cart []byte) *Memory {
	mem := &Memory{
		data: make([]byte, 64*1024),
		cart: Cartridge(cart),
		boot: BootROM,
	}
	if len(cart) == 32*1024 { // tetris
		copy(mem.data, mem.cart)
	}
	return mem
}

func (m *Memory) Size() int {
	return len(m.data)
}

func (m *Memory) Access(p uint16) (byte, bool) {
	if m.BootActive() && p <= 0xff { // intercept
		return m.boot[p], true
	}
	if int(p) >= len(m.data) {
		return 0, false
	}
	return m.data[p], true
}
func (m *Memory) AccessU16(p uint16) uint16 {
	// var msb, lsb byte
	lsb, _ := m.Access(p)
	msb, _ := m.Access(p + 1)
	return concatU16(msb, lsb)
}

func (m *Memory) WriteInstr(v uint8) *Memory {
	return m.Write(v)
}
func (m *Memory) Write(elems ...any) *Memory {
	for _, v := range elems {
		switch v := v.(type) {
		case []byte:
			for _, b := range v {
				m.data[m.i] = b
				m.i++
			}
		case []Block:
			for _, block := range v {
				m.Write(block)
			}
		case Block:
			m.i = int(v.Offset)
			for _, b := range v.Data {
				m.data[m.i] = b
				m.i++
			}
		case uint8:
			m.data[m.i] = v
			m.i++
		case int:
			m.data[m.i] = uint8(v)
			m.i++
		case string:
			// treat as code..
			m.data[m.i] = code(v)
			m.i++
		default:
			panic(fmt.Sprintf("not implemented for %T", v))
		}
	}
	return m
}
func (m *Memory) CursorAt(p int) *Memory {
	m.i = p
	return m
}

func (m *Memory) WriteByteAt(off uint16, value byte) *Memory {
	return m.WriteData(off, []byte{value})
}
func (m *Memory) WriteData(off uint16, p []byte) *Memory {
	if len(m.data) < int(off)+len(p) {
		panic("Outside bound")
	}
	for i := range len(p) {
		m.data[int(off)+i] = p[i]
	}
	return m
}

func (m *Memory) Dump(w io.Writer) {
	fmt.Fprintln(w, hex.Dump(m.data))
}

// These two registers specify the on-screen coordinates of the Window’s
// top-left pixel.
// WX=0..166 and WY=0..143
// Putting WX=7 and WY=0 places the Window at top left corner
func (m *Memory) WY() uint8 { return m.data[0xFF4a] }
func (m *Memory) WX() uint8 { return m.data[0xff4b] }

type tileset int

const (
	TilesetBackground tileset = iota
	TilesetWindow
)

type VRAM struct {
	TileData1 []byte // 4kB: 2kB unique, 2kB overlap with TileData2
	TileData2 []byte // 4kB: 2kB overlap with TileData1 and 2kB unique
	TileView1 []byte // 1kB
	TileView2 []byte // 1kB
}

func (v VRAM) HasData() bool {
	someSet := func(inp []byte) bool {
		for _, b := range inp {
			if b > 0 {
				return true
			}
		}
		return false
	}

	return slices.ContainsFunc([][]byte{
		v.TileData1, v.TileData2,
		v.TileView1, v.TileView2,
	}, someSet)
}

func (v VRAM) Tile(s tileset, index uint8) Tile {
	var buf []byte
	switch s {
	case TilesetWindow:
		buf = v.TileData2
	case TilesetBackground:
		buf = v.TileData1
	default:
		panic("Unknown tileset")
	}
	offset := int(index) * TILE_DATA_SIZE
	data := buf[offset : offset+TILE_DATA_SIZE]
	return Tile(data)
}

func (m *Memory) VRAM() VRAM {
	return VRAM{
		TileData1: m.data[0x8000:0x9000],
		TileData2: m.data[0x8800:0x9800],
		TileView1: m.data[0x9800:0x9c00],
		TileView2: m.data[0x9c00:0xa000],
	}
}

const (
	ADDR_LCDC = 0xff40
	ADDR_STAT = 0xff41
	ADDR_SCY  = 0xff42
	ADDR_SCK  = 0xff43
	ADDR_LY   = 0xff44
	ADDR_LYC  = 0xff45
)

type ControlRegisterPPU byte

func (r ControlRegisterPPU) bitb(n int) bool         { return bit(byte(r), n) > 0 }
func (r ControlRegisterPPU) BackgroundDisplay() bool { return r.bitb(0) }
func (r ControlRegisterPPU) SpriteDisplay() bool     { return r.bitb(1) }

// ... bunch of others

func (m *Memory) LCDC() ControlRegisterPPU { return ControlRegisterPPU(m.data[ADDR_LCDC]) }
func (m *Memory) STAT() uint8              { return m.data[ADDR_STAT] }

// These two registers specify the top-left coordinates of the visible 160×144
// pixel area within the 256×256 pixels BG map. Values in the range 0–255 may
// be used.
func (m *Memory) SCY() uint8 { return m.data[ADDR_SCY] } // Vertical Scroll Register
func (m *Memory) SCX() uint8 { return m.data[ADDR_SCK] } // Horizontal Scroll Register

// LY indicates the current horizontal line, which might be about to be drawn,
// being drawn, or just been drawn. LY can hold any value from 0 to 153, with
// values from 144 to 153 indicating the VBlank period.

func (m *Memory) LY() uint8  { return m.data[ADDR_LY] }  // Scanline Register
func (m *Memory) LYC() uint8 { return m.data[ADDR_LYC] } // Scanline Compare Register

// Read BOOT - Boot ROM lock register.
// If it's active, then it intercepts accesses to 0x0000 - 0x00FF and executes the
// boot rom code. Otherwise, it the address range 0x0000 - 0x00FF works normally.
func (m *Memory) BootActive() bool {
	return bit(m.data[0xFF50], 0) == 0
}

// For testing, we may disable boot and go straight to the cart
func (m *Memory) DisableBoot() {
	m.data[0xFF50] = 1
}

type Block struct {
	Offset uint16
	Data   []byte
}

// read bit at location n
func bit(b byte, n int) int { // 0 or 1
	if (b & (1 << n)) > 0 {
		return 1
	}
	return 0
}
