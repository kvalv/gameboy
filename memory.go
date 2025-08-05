package gameboy

import (
	"encoding/hex"
	"fmt"
	"io"
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

type VRAM struct {
	TileData1 []byte // 4kB: 2kB unique, 2kB overlap with TileData2
	TileData2 []byte // 4kB: 2kB overlap with TileData1 and 2kB unique
	TileView1 []byte // 1kB
	TileView2 []byte // 1kB
}

func (m *Memory) VRAM() VRAM {
	return VRAM{
		TileData1: m.data[0x8000:0x9000],
		TileData2: m.data[0x8800:0x9800],
		TileView1: m.data[0x9800:0x9c00],
		TileView2: m.data[0x9c00:0xa000],
	}
}

type Block struct {
	Offset uint16
	Data   []byte
}
