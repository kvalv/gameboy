package gameboy

import "fmt"

func checkSerial(mem *Memory) (byte, bool) {
	b := mem.Read(0xFF01)
	sc := mem.Read(0xFF02)
	if b > 0 {
		fmt.Printf("yay b=%x sc=%x\n", b, sc)
	}
	return b, sc>>7 == 1
}
