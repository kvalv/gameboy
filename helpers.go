package gameboy

type TwoBytes interface {
	uint16 | int16
}

func msb[V TwoBytes](v V) byte {
	return byte(v >> 8)
}
func lsb[V TwoBytes](v V) byte {
	return byte(v)
}
func split[V TwoBytes](v V) (byte, byte) {
	return msb(v), lsb(v)
}
func splitU16(v uint16) (msb, lsb uint8) {
	msb = uint8(v >> 8)
	lsb = uint8(v)
	return
}

func concatU16(msb, lsb uint8) uint16 { return uint16(msb)<<8 | uint16(lsb) }

type Value interface {
	uint8 | uint16
}

func sub[V Value](lhs, rhs V) (V, FlagRegister) {
	out := lhs - rhs
	var fl Flags
	if lhs < rhs {
		fl |= FLAGC
	}
	if out == 0 {
		fl |= FLAGZ
	}
	return out, FlagRegister(fl)
}

func add[V Value](lhs, rhs V) (V, FlagRegister) {
	out := lhs + rhs
	var fl Flags
	if out < lhs || out < rhs {
		fl |= FLAGC
	}
	if out == 0 {
		fl |= FLAGZ
	}
	return out, FlagRegister(fl)
}
