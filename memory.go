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
}

func NewMemory() *Memory {
	kB := 1024
	return &Memory{
		data: make([]byte, 64*kB),
	}

}

func (m *Memory) Size() int {
	return len(m.data)
}

func (m *Memory) Access(p uint16) (byte, bool) {
	if int(p) >= len(m.data) {
		return 0, false
	}
	return m.data[p], true
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

// LY indicates the current horizontal line, which might be about to be drawn,
// being drawn, or just been drawn. LY can hold any value from 0 to 153, with
// values from 144 to 153 indicating the VBlank period.
func (m *Memory) LY() uint8 {
	return m.data[0xFF44]
}
func (m *Memory) SetLY() {
	m.data[0xFF44] = 0x90
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

// These two registers specify the top-left coordinates of the visible 160×144
// pixel area within the 256×256 pixels BG map. Values in the range 0–255 may
// be used.
func (m *Memory) SCY() uint8 { return m.data[0xFF42] }
func (m *Memory) SCX() uint8 { return m.data[0xFF43] }

type Block struct {
	Offset uint16
	Data   []byte
}
