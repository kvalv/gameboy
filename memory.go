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

func (m *Memory) WriteU8(v uint8) *Memory {
	m.data = append(m.data, v)
	return m
}
func (m *Memory) Write(p []byte) *Memory {
	m.data = append(m.data, p...)
	return m
}
func (m *Memory) WriteAt(off uint16, p []byte) *Memory {
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
