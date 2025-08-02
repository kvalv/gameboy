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

// func sub[V Value](lhs, rhs V) (V, FlagRegister) {
func sub[L Value, R int | uint8 | uint16 | int8](lhs L, rhs R) (L, FlagRegister) {
	v := (int(lhs) - int(rhs))
	out := L(v)
	var fl Flags
	if v < 0 {
		fl |= FLAGC
	} else if v > int(out) {
		// this means we had an overflow, e.g. 0x00 - 0x01 = 0xFF
		fl |= FLAGC
	}
	if out == 0 {
		fl |= FLAGZ
	}
	return out, FlagRegister(fl)
}

func add[L uint8 | uint16, R int | uint8 | uint16 | int8](lhs L, rhs R) (L, FlagRegister) {
	v := (int(lhs) + int(rhs))
	out := L(v)
	var fl Flags
	if v != int(out) {
		fl |= FLAGC
	}
	if out == 0 {
		fl |= FLAGZ
	}
	return out, FlagRegister(fl)
}
