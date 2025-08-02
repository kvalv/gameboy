package gameboy

type Instruction func(cpu *CPU)

// LD H,D    code=0x62
func LD_62(cpu *CPU) {

	data := cpu.D

	cpu.H = data

	cpu.cycles += 4
}

// LD D,D    code=0x52
func LD_52(cpu *CPU) {

	data := cpu.D

	cpu.D = data

	cpu.cycles += 4
}

// LD D,L    code=0x55
func LD_55(cpu *CPU) {

	data := cpu.L

	cpu.D = data

	cpu.cycles += 4
}

// LD (HL),L    code=0x75
func LD_75(cpu *CPU) {

	data := cpu.L

	cpu.WriteMemory(cpu.HL(), data)

	cpu.cycles += 8
}

// ADD A,n8    code=0xc6
func ADD_C6(cpu *CPU) {

	lhs := cpu.A
	rhs := cpu.readU8(cpu.PC)
	res, flags := add(lhs, rhs)
	cpu.A = res
	cpu.F = flags

	cpu.cycles += 8
}

// DEC B    code=0x05
func DEC_05(cpu *CPU) {

	res, flags := sub(cpu.B, 0x01)
	cpu.F = flags
	cpu.B = res

	cpu.cycles += 4
}

// ADD HL,BC    code=0x09
func ADD_09(cpu *CPU) {

	lhs := cpu.HL()
	rhs := cpu.BC()
	res, flags := add(lhs, rhs)
	cpu.H, cpu.L = split(res)
	cpu.F = flags

	cpu.cycles += 8
}

// INC E    code=0x1c
func INC_1C(cpu *CPU) {

	res, flags := add(cpu.E, 0x01)
	cpu.F = flags
	cpu.E = res

	cpu.cycles += 4
}

// LD DE,n16    code=0x11
func LD_11(cpu *CPU) {

	data := cpu.readU16(cpu.PC)

	cpu.D, cpu.E = split(data)

	cpu.cycles += 12
}

// LD H,L    code=0x65
func LD_65(cpu *CPU) {

	data := cpu.L

	cpu.H = data

	cpu.cycles += 4
}

// ADD A,B    code=0x80
func ADD_80(cpu *CPU) {

	lhs := cpu.A
	rhs := cpu.B
	res, flags := add(lhs, rhs)
	cpu.A = res
	cpu.F = flags

	cpu.cycles += 4
}

// ADD A,(HL)    code=0x86
func ADD_86(cpu *CPU) {

	lhs := cpu.A
	rhs := cpu.loadU8(cpu.HL())
	res, flags := add(lhs, rhs)
	cpu.A = res
	cpu.F = flags

	cpu.cycles += 8
}

// LD HL,SP+,e8    code=0xf8
func LD_F8(cpu *CPU) {

	e := cpu.readI8(cpu.PC)
	res, flags := add(cpu.SP, e)
	cpu.H, cpu.L = split(res)
	cpu.F = flags

	cpu.cycles += 12
}

// LD L,L    code=0x6d
func LD_6D(cpu *CPU) {

	data := cpu.L

	cpu.L = data

	cpu.cycles += 4
}

// LD A,(a16)    code=0xfa
func LD_FA(cpu *CPU) {

	data := cpu.loadU8(cpu.readU16(cpu.PC))

	cpu.A = data

	cpu.cycles += 16
}

// LD A,(BC)    code=0x0a
func LD_0A(cpu *CPU) {

	data := cpu.loadU8(cpu.BC())

	cpu.A = data

	cpu.cycles += 8
}

// INC HL    code=0x23
func INC_23(cpu *CPU) {

	res, flags := add(cpu.HL(), 0x01)
	cpu.F = flags
	cpu.H, cpu.L = split(res)

	cpu.cycles += 8
}

// LD C,C    code=0x49
func LD_49(cpu *CPU) {

	data := cpu.C

	cpu.C = data

	cpu.cycles += 4
}

// LD A,C    code=0x79
func LD_79(cpu *CPU) {

	data := cpu.C

	cpu.A = data

	cpu.cycles += 4
}

// LD E,L    code=0x5d
func LD_5D(cpu *CPU) {

	data := cpu.L

	cpu.E = data

	cpu.cycles += 4
}

// LD A,A    code=0x7f
func LD_7F(cpu *CPU) {

	data := cpu.A

	cpu.A = data

	cpu.cycles += 4
}

// INC DE    code=0x13
func INC_13(cpu *CPU) {

	res, flags := add(cpu.DE(), 0x01)
	cpu.F = flags
	cpu.D, cpu.E = split(res)

	cpu.cycles += 8
}

// INC C    code=0x0c
func INC_0C(cpu *CPU) {

	res, flags := add(cpu.C, 0x01)
	cpu.F = flags
	cpu.C = res

	cpu.cycles += 4
}

// INC BC    code=0x03
func INC_03(cpu *CPU) {

	res, flags := add(cpu.BC(), 0x01)
	cpu.F = flags
	cpu.B, cpu.C = split(res)

	cpu.cycles += 8
}

// LD L,E    code=0x6b
func LD_6B(cpu *CPU) {

	data := cpu.E

	cpu.L = data

	cpu.cycles += 4
}

// ADD A,D    code=0x82
func ADD_82(cpu *CPU) {

	lhs := cpu.A
	rhs := cpu.D
	res, flags := add(lhs, rhs)
	cpu.A = res
	cpu.F = flags

	cpu.cycles += 4
}

// LD A,n8    code=0x3e
func LD_3E(cpu *CPU) {

	data := cpu.readU8(cpu.PC)

	cpu.A = data

	cpu.cycles += 8
}

// INC SP    code=0x33
func INC_33(cpu *CPU) {

	res, flags := add(cpu.SP, 0x01)
	cpu.F = flags
	cpu.SP = res

	cpu.cycles += 8
}

// LD B,B    code=0x40
func LD_40(cpu *CPU) {

	data := cpu.B

	cpu.B = data

	cpu.cycles += 4
}

// LD (HL),B    code=0x70
func LD_70(cpu *CPU) {

	data := cpu.B

	cpu.WriteMemory(cpu.HL(), data)

	cpu.cycles += 8
}

// ADD SP,e8    code=0xe8
func ADD_E8(cpu *CPU) {

	lhs := cpu.SP
	rhs := cpu.readI8(cpu.PC)
	res, flags := add(lhs, rhs)
	cpu.SP = res
	cpu.F = flags

	cpu.cycles += 16
}

// LD E,(HL)    code=0x5e
func LD_5E(cpu *CPU) {

	data := cpu.loadU8(cpu.HL())

	cpu.E = data

	cpu.cycles += 8
}

// LD A,(HL-)    code=0x3a
func LD_3A(cpu *CPU) {

	data := cpu.loadU8(cpu.HL())

	cpu.A = data

	cpu.cycles += 8
}

// LD B,L    code=0x45
func LD_45(cpu *CPU) {

	data := cpu.L

	cpu.B = data

	cpu.cycles += 4
}

// LD B,A    code=0x47
func LD_47(cpu *CPU) {

	data := cpu.A

	cpu.B = data

	cpu.cycles += 4
}

// LD (HL+),A    code=0x22
func LD_22(cpu *CPU) {

	data := cpu.A

	cpu.WriteMemory(cpu.HL(), data)

	incr, flags := add(cpu.HL(), 0x01)
	cpu.H, cpu.L = split(incr)
	cpu.F = flags

	cpu.cycles += 8
}

// INC (HL)    code=0x34
func INC_34(cpu *CPU) {

	res, flags := add(cpu.loadU8(cpu.HL()), 0x01)
	cpu.F = flags
	cpu.WriteMemory(cpu.HL(), res)

	cpu.cycles += 12
}

// DEC E    code=0x1d
func DEC_1D(cpu *CPU) {

	res, flags := sub(cpu.E, 0x01)
	cpu.F = flags
	cpu.E = res

	cpu.cycles += 4
}

// LD A,E    code=0x7b
func LD_7B(cpu *CPU) {

	data := cpu.E

	cpu.A = data

	cpu.cycles += 4
}

// LD A,(HL)    code=0x7e
func LD_7E(cpu *CPU) {

	data := cpu.loadU8(cpu.HL())

	cpu.A = data

	cpu.cycles += 8
}

// LD C,n8    code=0x0e
func LD_0E(cpu *CPU) {

	data := cpu.readU8(cpu.PC)

	cpu.C = data

	cpu.cycles += 8
}

// LD B,H    code=0x44
func LD_44(cpu *CPU) {

	data := cpu.H

	cpu.B = data

	cpu.cycles += 4
}

// LD D,H    code=0x54
func LD_54(cpu *CPU) {

	data := cpu.H

	cpu.D = data

	cpu.cycles += 4
}

// LD (DE),A    code=0x12
func LD_12(cpu *CPU) {

	data := cpu.A

	cpu.WriteMemory(cpu.DE(), data)

	cpu.cycles += 8
}

// LD L,n8    code=0x2e
func LD_2E(cpu *CPU) {

	data := cpu.readU8(cpu.PC)

	cpu.L = data

	cpu.cycles += 8
}

// LD D,(HL)    code=0x56
func LD_56(cpu *CPU) {

	data := cpu.loadU8(cpu.HL())

	cpu.D = data

	cpu.cycles += 8
}

// LD (BC),A    code=0x02
func LD_02(cpu *CPU) {

	data := cpu.A

	cpu.WriteMemory(cpu.BC(), data)

	cpu.cycles += 8
}

// INC D    code=0x14
func INC_14(cpu *CPU) {

	res, flags := add(cpu.D, 0x01)
	cpu.F = flags
	cpu.D = res

	cpu.cycles += 4
}

// LD BC,n16    code=0x01
func LD_01(cpu *CPU) {

	data := cpu.readU16(cpu.PC)

	cpu.B, cpu.C = split(data)

	cpu.cycles += 12
}

// LD (HL),D    code=0x72
func LD_72(cpu *CPU) {

	data := cpu.D

	cpu.WriteMemory(cpu.HL(), data)

	cpu.cycles += 8
}

// ADD A,C    code=0x81
func ADD_81(cpu *CPU) {

	lhs := cpu.A
	rhs := cpu.C
	res, flags := add(lhs, rhs)
	cpu.A = res
	cpu.F = flags

	cpu.cycles += 4
}

// LD (a16),A    code=0xea
func LD_EA(cpu *CPU) {

	data := cpu.A

	cpu.WriteMemory(cpu.readU16(cpu.PC), data)

	cpu.cycles += 16
}

// ADD HL,HL    code=0x29
func ADD_29(cpu *CPU) {

	lhs := cpu.HL()
	rhs := cpu.HL()
	res, flags := add(lhs, rhs)
	cpu.H, cpu.L = split(res)
	cpu.F = flags

	cpu.cycles += 8
}

// LD C,E    code=0x4b
func LD_4B(cpu *CPU) {

	data := cpu.E

	cpu.C = data

	cpu.cycles += 4
}

// LD H,B    code=0x60
func LD_60(cpu *CPU) {

	data := cpu.B

	cpu.H = data

	cpu.cycles += 4
}

// ADD A,L    code=0x85
func ADD_85(cpu *CPU) {

	lhs := cpu.A
	rhs := cpu.L
	res, flags := add(lhs, rhs)
	cpu.A = res
	cpu.F = flags

	cpu.cycles += 4
}

// LD E,A    code=0x5f
func LD_5F(cpu *CPU) {

	data := cpu.A

	cpu.E = data

	cpu.cycles += 4
}

// LD L,A    code=0x6f
func LD_6F(cpu *CPU) {

	data := cpu.A

	cpu.L = data

	cpu.cycles += 4
}

// LD (HL),n8    code=0x36
func LD_36(cpu *CPU) {

	data := cpu.readU8(cpu.PC)

	cpu.WriteMemory(cpu.HL(), data)

	cpu.cycles += 12
}

// LD E,H    code=0x5c
func LD_5C(cpu *CPU) {

	data := cpu.H

	cpu.E = data

	cpu.cycles += 4
}

// LD L,C    code=0x69
func LD_69(cpu *CPU) {

	data := cpu.C

	cpu.L = data

	cpu.cycles += 4
}

// LD (a16),SP    code=0x08
func LD_08(cpu *CPU) {

	data := cpu.SP

	cpu.WriteMemory(cpu.readU16(cpu.PC), data)

	cpu.cycles += 20
}

// LD (HL),C    code=0x71
func LD_71(cpu *CPU) {

	data := cpu.C

	cpu.WriteMemory(cpu.HL(), data)

	cpu.cycles += 8
}

// DEC SP    code=0x3b
func DEC_3B(cpu *CPU) {

	res, flags := sub(cpu.SP, 0x01)
	cpu.F = flags
	cpu.SP = res

	cpu.cycles += 8
}

// ADD A,H    code=0x84
func ADD_84(cpu *CPU) {

	lhs := cpu.A
	rhs := cpu.H
	res, flags := add(lhs, rhs)
	cpu.A = res
	cpu.F = flags

	cpu.cycles += 4
}

// LD C,H    code=0x4c
func LD_4C(cpu *CPU) {

	data := cpu.H

	cpu.C = data

	cpu.cycles += 4
}

// LD L,D    code=0x6a
func LD_6A(cpu *CPU) {

	data := cpu.D

	cpu.L = data

	cpu.cycles += 4
}

// DEC (HL)    code=0x35
func DEC_35(cpu *CPU) {

	res, flags := sub(cpu.loadU8(cpu.HL()), 0x01)
	cpu.F = flags
	cpu.WriteMemory(cpu.HL(), res)

	cpu.cycles += 12
}

// LD D,n8    code=0x16
func LD_16(cpu *CPU) {

	data := cpu.readU8(cpu.PC)

	cpu.D = data

	cpu.cycles += 8
}

// LD HL,n16    code=0x21
func LD_21(cpu *CPU) {

	data := cpu.readU16(cpu.PC)

	cpu.H, cpu.L = split(data)

	cpu.cycles += 12
}

// LD D,E    code=0x53
func LD_53(cpu *CPU) {

	data := cpu.E

	cpu.D = data

	cpu.cycles += 4
}

// DEC D    code=0x15
func DEC_15(cpu *CPU) {

	res, flags := sub(cpu.D, 0x01)
	cpu.F = flags
	cpu.D = res

	cpu.cycles += 4
}

// LD A,(DE)    code=0x1a
func LD_1A(cpu *CPU) {

	data := cpu.loadU8(cpu.DE())

	cpu.A = data

	cpu.cycles += 8
}

// LD E,n8    code=0x1e
func LD_1E(cpu *CPU) {

	data := cpu.readU8(cpu.PC)

	cpu.E = data

	cpu.cycles += 8
}

// INC H    code=0x24
func INC_24(cpu *CPU) {

	res, flags := add(cpu.H, 0x01)
	cpu.F = flags
	cpu.H = res

	cpu.cycles += 4
}

// LD A,B    code=0x78
func LD_78(cpu *CPU) {

	data := cpu.B

	cpu.A = data

	cpu.cycles += 4
}

// LD A,H    code=0x7c
func LD_7C(cpu *CPU) {

	data := cpu.H

	cpu.A = data

	cpu.cycles += 4
}

// LD C,B    code=0x48
func LD_48(cpu *CPU) {

	data := cpu.B

	cpu.C = data

	cpu.cycles += 4
}

// LD D,B    code=0x50
func LD_50(cpu *CPU) {

	data := cpu.B

	cpu.D = data

	cpu.cycles += 4
}

// ADD HL,DE    code=0x19
func ADD_19(cpu *CPU) {

	lhs := cpu.HL()
	rhs := cpu.DE()
	res, flags := add(lhs, rhs)
	cpu.H, cpu.L = split(res)
	cpu.F = flags

	cpu.cycles += 8
}

// DEC A    code=0x3d
func DEC_3D(cpu *CPU) {

	res, flags := sub(cpu.A, 0x01)
	cpu.F = flags
	cpu.A = res

	cpu.cycles += 4
}

// LD C,A    code=0x4f
func LD_4F(cpu *CPU) {

	data := cpu.A

	cpu.C = data

	cpu.cycles += 4
}

// LD B,n8    code=0x06
func LD_06(cpu *CPU) {

	data := cpu.readU8(cpu.PC)

	cpu.B = data

	cpu.cycles += 8
}

// DEC C    code=0x0d
func DEC_0D(cpu *CPU) {

	res, flags := sub(cpu.C, 0x01)
	cpu.F = flags
	cpu.C = res

	cpu.cycles += 4
}

// LD C,D    code=0x4a
func LD_4A(cpu *CPU) {

	data := cpu.D

	cpu.C = data

	cpu.cycles += 4
}

// LD D,A    code=0x57
func LD_57(cpu *CPU) {

	data := cpu.A

	cpu.D = data

	cpu.cycles += 4
}

// ADD A,A    code=0x87
func ADD_87(cpu *CPU) {

	lhs := cpu.A
	rhs := cpu.A
	res, flags := add(lhs, rhs)
	cpu.A = res
	cpu.F = flags

	cpu.cycles += 4
}

// LD L,(HL)    code=0x6e
func LD_6E(cpu *CPU) {

	data := cpu.loadU8(cpu.HL())

	cpu.L = data

	cpu.cycles += 8
}

// DEC H    code=0x25
func DEC_25(cpu *CPU) {

	res, flags := sub(cpu.H, 0x01)
	cpu.F = flags
	cpu.H = res

	cpu.cycles += 4
}

// LD D,C    code=0x51
func LD_51(cpu *CPU) {

	data := cpu.C

	cpu.D = data

	cpu.cycles += 4
}

// LD SP,HL    code=0xf9
func LD_F9(cpu *CPU) {

	data := cpu.HL()

	cpu.SP = data

	cpu.cycles += 8
}

// DEC DE    code=0x1b
func DEC_1B(cpu *CPU) {

	res, flags := sub(cpu.DE(), 0x01)
	cpu.F = flags
	cpu.D, cpu.E = split(res)

	cpu.cycles += 8
}

// DEC HL    code=0x2b
func DEC_2B(cpu *CPU) {

	res, flags := sub(cpu.HL(), 0x01)
	cpu.F = flags
	cpu.H, cpu.L = split(res)

	cpu.cycles += 8
}

// LD B,(HL)    code=0x46
func LD_46(cpu *CPU) {

	data := cpu.loadU8(cpu.HL())

	cpu.B = data

	cpu.cycles += 8
}

// LD L,B    code=0x68
func LD_68(cpu *CPU) {

	data := cpu.B

	cpu.L = data

	cpu.cycles += 4
}

// ADD A,E    code=0x83
func ADD_83(cpu *CPU) {

	lhs := cpu.A
	rhs := cpu.E
	res, flags := add(lhs, rhs)
	cpu.A = res
	cpu.F = flags

	cpu.cycles += 4
}

// LD H,n8    code=0x26
func LD_26(cpu *CPU) {

	data := cpu.readU8(cpu.PC)

	cpu.H = data

	cpu.cycles += 8
}

// LD SP,n16    code=0x31
func LD_31(cpu *CPU) {

	data := cpu.readU16(cpu.PC)

	cpu.SP = data

	cpu.cycles += 12
}

// INC A    code=0x3c
func INC_3C(cpu *CPU) {

	res, flags := add(cpu.A, 0x01)
	cpu.F = flags
	cpu.A = res

	cpu.cycles += 4
}

// LD H,H    code=0x64
func LD_64(cpu *CPU) {

	data := cpu.H

	cpu.H = data

	cpu.cycles += 4
}

// LD H,C    code=0x61
func LD_61(cpu *CPU) {

	data := cpu.C

	cpu.H = data

	cpu.cycles += 4
}

// LD H,A    code=0x67
func LD_67(cpu *CPU) {

	data := cpu.A

	cpu.H = data

	cpu.cycles += 4
}

// LD C,L    code=0x4d
func LD_4D(cpu *CPU) {

	data := cpu.L

	cpu.C = data

	cpu.cycles += 4
}

// LD B,E    code=0x43
func LD_43(cpu *CPU) {

	data := cpu.E

	cpu.B = data

	cpu.cycles += 4
}

// LD B,D    code=0x42
func LD_42(cpu *CPU) {

	data := cpu.D

	cpu.B = data

	cpu.cycles += 4
}

// LD E,B    code=0x58
func LD_58(cpu *CPU) {

	data := cpu.B

	cpu.E = data

	cpu.cycles += 4
}

// LD L,H    code=0x6c
func LD_6C(cpu *CPU) {

	data := cpu.H

	cpu.L = data

	cpu.cycles += 4
}

// LD (HL),E    code=0x73
func LD_73(cpu *CPU) {

	data := cpu.E

	cpu.WriteMemory(cpu.HL(), data)

	cpu.cycles += 8
}

// LD H,(HL)    code=0x66
func LD_66(cpu *CPU) {

	data := cpu.loadU8(cpu.HL())

	cpu.H = data

	cpu.cycles += 8
}

// LD H,E    code=0x63
func LD_63(cpu *CPU) {

	data := cpu.E

	cpu.H = data

	cpu.cycles += 4
}

// DEC L    code=0x2d
func DEC_2D(cpu *CPU) {

	res, flags := sub(cpu.L, 0x01)
	cpu.F = flags
	cpu.L = res

	cpu.cycles += 4
}

// LD B,C    code=0x41
func LD_41(cpu *CPU) {

	data := cpu.C

	cpu.B = data

	cpu.cycles += 4
}

// LD E,D    code=0x5a
func LD_5A(cpu *CPU) {

	data := cpu.D

	cpu.E = data

	cpu.cycles += 4
}

// LD A,(HL+)    code=0x2a
func LD_2A(cpu *CPU) {

	data := cpu.loadU8(cpu.HL())

	cpu.A = data

	cpu.cycles += 8
}

// LD E,C    code=0x59
func LD_59(cpu *CPU) {

	data := cpu.C

	cpu.E = data

	cpu.cycles += 4
}

// LD E,E    code=0x5b
func LD_5B(cpu *CPU) {

	data := cpu.E

	cpu.E = data

	cpu.cycles += 4
}

// LD (HL),H    code=0x74
func LD_74(cpu *CPU) {

	data := cpu.H

	cpu.WriteMemory(cpu.HL(), data)

	cpu.cycles += 8
}

// LD A,L    code=0x7d
func LD_7D(cpu *CPU) {

	data := cpu.L

	cpu.A = data

	cpu.cycles += 4
}

// LD (HL-),A    code=0x32
func LD_32(cpu *CPU) {

	data := cpu.A

	cpu.WriteMemory(cpu.HL(), data)

	decr, flags := sub(cpu.HL(), 0x01)
	cpu.H, cpu.L = split(decr)
	cpu.F = flags

	cpu.cycles += 8
}

// LD C,(HL)    code=0x4e
func LD_4E(cpu *CPU) {

	data := cpu.loadU8(cpu.HL())

	cpu.C = data

	cpu.cycles += 8
}

// LD (HL),A    code=0x77
func LD_77(cpu *CPU) {

	data := cpu.A

	cpu.WriteMemory(cpu.HL(), data)

	cpu.cycles += 8
}

// INC B    code=0x04
func INC_04(cpu *CPU) {

	res, flags := add(cpu.B, 0x01)
	cpu.F = flags
	cpu.B = res

	cpu.cycles += 4
}

// INC L    code=0x2c
func INC_2C(cpu *CPU) {

	res, flags := add(cpu.L, 0x01)
	cpu.F = flags
	cpu.L = res

	cpu.cycles += 4
}

// ADD HL,SP    code=0x39
func ADD_39(cpu *CPU) {

	lhs := cpu.HL()
	rhs := cpu.SP
	res, flags := add(lhs, rhs)
	cpu.H, cpu.L = split(res)
	cpu.F = flags

	cpu.cycles += 8
}

// LD A,D    code=0x7a
func LD_7A(cpu *CPU) {

	data := cpu.D

	cpu.A = data

	cpu.cycles += 4
}

// DEC BC    code=0x0b
func DEC_0B(cpu *CPU) {

	res, flags := sub(cpu.BC(), 0x01)
	cpu.F = flags
	cpu.B, cpu.C = split(res)

	cpu.cycles += 8
}

var ops = map[uint8]Instruction{
	0x62: LD_62,
	0x52: LD_52,
	0x55: LD_55,
	0x75: LD_75,
	0xc6: ADD_C6,
	0x5:  DEC_05,
	0x9:  ADD_09,
	0x1c: INC_1C,
	0x11: LD_11,
	0x65: LD_65,
	0x80: ADD_80,
	0x86: ADD_86,
	0xf8: LD_F8,
	0x6d: LD_6D,
	0xfa: LD_FA,
	0xa:  LD_0A,
	0x23: INC_23,
	0x49: LD_49,
	0x79: LD_79,
	0x5d: LD_5D,
	0x7f: LD_7F,
	0x13: INC_13,
	0xc:  INC_0C,
	0x3:  INC_03,
	0x6b: LD_6B,
	0x82: ADD_82,
	0x3e: LD_3E,
	0x33: INC_33,
	0x40: LD_40,
	0x70: LD_70,
	0xe8: ADD_E8,
	0x5e: LD_5E,
	0x3a: LD_3A,
	0x45: LD_45,
	0x47: LD_47,
	0x22: LD_22,
	0x34: INC_34,
	0x1d: DEC_1D,
	0x7b: LD_7B,
	0x7e: LD_7E,
	0xe:  LD_0E,
	0x44: LD_44,
	0x54: LD_54,
	0x12: LD_12,
	0x2e: LD_2E,
	0x56: LD_56,
	0x2:  LD_02,
	0x14: INC_14,
	0x1:  LD_01,
	0x72: LD_72,
	0x81: ADD_81,
	0xea: LD_EA,
	0x29: ADD_29,
	0x4b: LD_4B,
	0x60: LD_60,
	0x85: ADD_85,
	0x5f: LD_5F,
	0x6f: LD_6F,
	0x36: LD_36,
	0x5c: LD_5C,
	0x69: LD_69,
	0x8:  LD_08,
	0x71: LD_71,
	0x3b: DEC_3B,
	0x84: ADD_84,
	0x4c: LD_4C,
	0x6a: LD_6A,
	0x35: DEC_35,
	0x16: LD_16,
	0x21: LD_21,
	0x53: LD_53,
	0x15: DEC_15,
	0x1a: LD_1A,
	0x1e: LD_1E,
	0x24: INC_24,
	0x78: LD_78,
	0x7c: LD_7C,
	0x48: LD_48,
	0x50: LD_50,
	0x19: ADD_19,
	0x3d: DEC_3D,
	0x4f: LD_4F,
	0x6:  LD_06,
	0xd:  DEC_0D,
	0x4a: LD_4A,
	0x57: LD_57,
	0x87: ADD_87,
	0x6e: LD_6E,
	0x25: DEC_25,
	0x51: LD_51,
	0xf9: LD_F9,
	0x1b: DEC_1B,
	0x2b: DEC_2B,
	0x46: LD_46,
	0x68: LD_68,
	0x83: ADD_83,
	0x26: LD_26,
	0x31: LD_31,
	0x3c: INC_3C,
	0x64: LD_64,
	0x61: LD_61,
	0x67: LD_67,
	0x4d: LD_4D,
	0x43: LD_43,
	0x42: LD_42,
	0x58: LD_58,
	0x6c: LD_6C,
	0x73: LD_73,
	0x66: LD_66,
	0x63: LD_63,
	0x2d: DEC_2D,
	0x41: LD_41,
	0x5a: LD_5A,
	0x2a: LD_2A,
	0x59: LD_59,
	0x5b: LD_5B,
	0x74: LD_74,
	0x7d: LD_7D,
	0x32: LD_32,
	0x4e: LD_4E,
	0x77: LD_77,
	0x4:  INC_04,
	0x2c: INC_2C,
	0x39: ADD_39,
	0x7a: LD_7A,
	0xb:  DEC_0B,
}
