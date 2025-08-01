package gameboy

type Instruction func(cpu *CPU)

// INC 0x03 BC
func INC_03(cpu *CPU) {

	res, flags := cpu.Add16(cpu.BC(), 0x01)

	cpu.B, cpu.C = splitU16(res)
	cpu.F = FlagRegister(flags)

	cpu.cycles += 8
}

// INC 0x04 B
func INC_04(cpu *CPU) {

	res, flags := cpu.Add(cpu.B, 0x01)

	cpu.B = res
	cpu.F = FlagRegister(flags)

	cpu.cycles += 4
}

// ADD 0x09 HL,BC
func ADD_09(cpu *CPU) {

	res, flags := cpu.Add16(cpu.HL(), cpu.BC())

	cpu.H, cpu.L = splitU16(res)

	cpu.F = FlagRegister(flags)

	cpu.cycles += 8
}

// INC 0x0C C
func INC_0C(cpu *CPU) {

	res, flags := cpu.Add(cpu.C, 0x01)

	cpu.C = res
	cpu.F = FlagRegister(flags)

	cpu.cycles += 4
}

// INC 0x13 DE
func INC_13(cpu *CPU) {

	res, flags := cpu.Add16(cpu.DE(), 0x01)

	cpu.D, cpu.E = splitU16(res)
	cpu.F = FlagRegister(flags)

	cpu.cycles += 8
}

// INC 0x14 D
func INC_14(cpu *CPU) {

	res, flags := cpu.Add(cpu.D, 0x01)

	cpu.D = res
	cpu.F = FlagRegister(flags)

	cpu.cycles += 4
}

// ADD 0x19 HL,DE
func ADD_19(cpu *CPU) {

	res, flags := cpu.Add16(cpu.HL(), cpu.DE())

	cpu.H, cpu.L = splitU16(res)

	cpu.F = FlagRegister(flags)

	cpu.cycles += 8
}

// INC 0x1C E
func INC_1C(cpu *CPU) {

	res, flags := cpu.Add(cpu.E, 0x01)

	cpu.E = res
	cpu.F = FlagRegister(flags)

	cpu.cycles += 4
}

// INC 0x23 HL
func INC_23(cpu *CPU) {

	res, flags := cpu.Add16(cpu.HL(), 0x01)

	cpu.H, cpu.L = splitU16(res)
	cpu.F = FlagRegister(flags)

	cpu.cycles += 8
}

// INC 0x24 H
func INC_24(cpu *CPU) {

	res, flags := cpu.Add(cpu.H, 0x01)

	cpu.H = res
	cpu.F = FlagRegister(flags)

	cpu.cycles += 4
}

// ADD 0x29 HL,HL
func ADD_29(cpu *CPU) {

	res, flags := cpu.Add16(cpu.HL(), cpu.HL())

	cpu.H, cpu.L = splitU16(res)

	cpu.F = FlagRegister(flags)

	cpu.cycles += 8
}

// INC 0x2C L
func INC_2C(cpu *CPU) {

	res, flags := cpu.Add(cpu.L, 0x01)

	cpu.L = res
	cpu.F = FlagRegister(flags)

	cpu.cycles += 4
}

// INC 0x33 SP
func INC_33(cpu *CPU) {

	res, flags := cpu.Add16(cpu.SP, 0x01)

	cpu.SP = res
	cpu.F = FlagRegister(flags)

	cpu.cycles += 8
}

// INC 0x34 HL
func INC_34(cpu *CPU) {

	// Increments data at the absolute address specified by the register
	var val uint8
	cpu.load(cpu.HL(), &val)
	next, flags := cpu.Add(val, 0x01)
	cpu.write(cpu.HL(), next)
	cpu.F = FlagRegister(flags)

	cpu.cycles += 12
}

// ADD 0x39 HL,SP
func ADD_39(cpu *CPU) {

	res, flags := cpu.Add16(cpu.HL(), cpu.SP)

	cpu.H, cpu.L = splitU16(res)

	cpu.F = FlagRegister(flags)

	cpu.cycles += 8
}

// INC 0x3C A
func INC_3C(cpu *CPU) {

	res, flags := cpu.Add(cpu.A, 0x01)

	cpu.A = res
	cpu.F = FlagRegister(flags)

	cpu.cycles += 4
}

// ADD 0x80 A,B
func ADD_80(cpu *CPU) {

	cpu.A, cpu.F = cpu.Add(cpu.A, cpu.B)

	cpu.cycles += 4
}

// ADD 0x81 A,C
func ADD_81(cpu *CPU) {

	cpu.A, cpu.F = cpu.Add(cpu.A, cpu.C)

	cpu.cycles += 4
}

// ADD 0x82 A,D
func ADD_82(cpu *CPU) {

	cpu.A, cpu.F = cpu.Add(cpu.A, cpu.D)

	cpu.cycles += 4
}

// ADD 0x83 A,E
func ADD_83(cpu *CPU) {

	cpu.A, cpu.F = cpu.Add(cpu.A, cpu.E)

	cpu.cycles += 4
}

// ADD 0x84 A,H
func ADD_84(cpu *CPU) {

	cpu.A, cpu.F = cpu.Add(cpu.A, cpu.H)

	cpu.cycles += 4
}

// ADD 0x85 A,L
func ADD_85(cpu *CPU) {

	cpu.A, cpu.F = cpu.Add(cpu.A, cpu.L)

	cpu.cycles += 4
}

// ADD 0x86 A,HL
func ADD_86(cpu *CPU) {

	// not immediate brah
	var n uint8
	cpu.load(cpu.HL(), &n)
	s8 := n // TODO: signed8(n)
	cpu.A, cpu.F = cpu.Add(cpu.A, s8)

	cpu.cycles += 8
}

// ADD 0x87 A,A
func ADD_87(cpu *CPU) {

	cpu.A, cpu.F = cpu.Add(cpu.A, cpu.A)

	cpu.cycles += 4
}

// ADD 0xC6 A,n8
func ADD_C6(cpu *CPU) {

	var n uint8
	cpu.load(cpu.PC, &n)
	cpu.A, cpu.F = cpu.Add(cpu.A, n)

	cpu.cycles += 8
}

// ADD 0xE8 SP,e8
func ADD_E8(cpu *CPU) {

	var n int16
	cpu.load(cpu.PC, &n)
	res, flags := cpu.AddSigned16(int16(cpu.A), n)
	cpu.PC = uint16(res)
	cpu.F = flags

	cpu.cycles += 16
}

var ops = map[uint8]Instruction{
	0x3:  INC_03,
	0x4:  INC_04,
	0x9:  ADD_09,
	0xc:  INC_0C,
	0x13: INC_13,
	0x14: INC_14,
	0x19: ADD_19,
	0x1c: INC_1C,
	0x23: INC_23,
	0x24: INC_24,
	0x29: ADD_29,
	0x2c: INC_2C,
	0x33: INC_33,
	0x34: INC_34,
	0x39: ADD_39,
	0x3c: INC_3C,
	0x80: ADD_80,
	0x81: ADD_81,
	0x82: ADD_82,
	0x83: ADD_83,
	0x84: ADD_84,
	0x85: ADD_85,
	0x86: ADD_86,
	0x87: ADD_87,
	0xc6: ADD_C6,
	0xe8: ADD_E8,
}
