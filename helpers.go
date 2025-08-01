package gameboy

// for the lack of a better name...
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
func concatI16(msb, lsb byte) int16   { return int16(msb)<<8 | int16(lsb) }

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

func addSigned[V uint8 | uint16](lhs V, rhs int8) (V, FlagRegister) {
	out := V(int16(lhs) + int16(rhs))
	var fl Flags
	if out < lhs || (out > 0 && out < V(rhs)) {
		fl |= FLAGC
	}
	if out == 0 {
		fl |= FLAGZ
	}
	return out, FlagRegister(fl)
}
