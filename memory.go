package gameboy

import (
	"encoding/hex"
	"fmt"
	"io"
)

type Memory struct {
	data []byte
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
	m.data = append(m.data, v)
	return m
}
func (m *Memory) Write(elems ...any) *Memory {
	for _, v := range elems {
		switch v := v.(type) {
		case []byte:
			m.data = append(m.data, v...)
		case uint8:
			m.data = append(m.data, v)
		case int:
			m.data = append(m.data, uint8(v))
		default:
			panic(fmt.Sprintf("not implemented for %T", v))
		}
	}
	return m
}

func (m *Memory) WriteByteAt(off uint16, value byte) *Memory {
	return m.WriteData(off, []byte{value})
}
func (m *Memory) Reserve(size uint16) *Memory {
	m.WriteByteAt(size, 0x00)
	return m
}
func (m *Memory) WriteData(off uint16, p []byte) *Memory {
	if len(m.data) < int(off)+len(p) {
		newData := make([]byte, off+uint16(len(p)))
		copy(newData, m.data)
		m.data = newData
	}
	for i := range len(p) {
		m.data[int(off)+i] = p[i]
	}
	return m
}

func (m *Memory) Dump(w io.Writer) {
	fmt.Fprintln(w, hex.Dump(m.data))
}
