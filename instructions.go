package gameboy

import "fmt"

type Instruction func(cpu *CPU)

// DEC HL    code=0x2b
func DEC_2B(cpu *CPU) {
	res, flags := sub(cpu.HL(), 0x01)
	cpu.F = flags
	cpu.H, cpu.L = split(res)
	cpu.cycles += 8
}

// LD C,E    code=0x4b
func LD_4B(cpu *CPU) {

	data := cpu.E

	cpu.C = data

	cpu.cycles += 4

}

// LD D,A    code=0x57
func LD_57(cpu *CPU) {

	data := cpu.A

	cpu.D = data

	cpu.cycles += 4

}

// LD A,D    code=0x7a
func LD_7A(cpu *CPU) {

	data := cpu.D

	cpu.A = data

	cpu.cycles += 4

}

// SBC A,D    code=0x9a
func SBC_9A(cpu *CPU) {
	panic("TODO SBC_9A")
}

// AND A,C    code=0xa1
func AND_A1(cpu *CPU) {
	panic("TODO AND_A1")
}

// AND A,E    code=0xa3
func AND_A3(cpu *CPU) {
	panic("TODO AND_A3")
}

// PUSH DE    code=0xd5
func PUSH_D5(cpu *CPU) {
	cpu.PushStack(cpu.DE())
	cpu.cycles += 16
}

// LD E,B    code=0x58
func LD_58(cpu *CPU) {

	data := cpu.B

	cpu.E = data

	cpu.cycles += 4

}

// LD (HL),L    code=0x75
func LD_75(cpu *CPU) {

	data := cpu.L

	cpu.WriteMemory(cpu.HL(), data)

	cpu.cycles += 8

}

// RET NZ    code=0xc0
func RET_C0(cpu *CPU) {
	if !cpu.F.HasZero() {
		cpu.PC = cpu.PopStack()
		cpu.cycles += 20
	} else {
		cpu.cycles += 8
	}
}

// ADC A,n8    code=0xce
func ADC_CE(cpu *CPU) {
	panic("TODO ADC_CE")
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

// RST $38    code=0xff
func RST_FF(cpu *CPU) {
	n := uint8(0x38)
	cpu.PushStack(cpu.PC)
	cpu.PC = concatU16(0x00, n)
}

// LD C,n8    code=0x0e
func LD_0E(cpu *CPU) {

	data := cpu.readU8(cpu.PC)

	cpu.C = data

	cpu.cycles += 8

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

// LD H,B    code=0x60
func LD_60(cpu *CPU) {

	data := cpu.B

	cpu.H = data

	cpu.cycles += 4

}

// LD L,E    code=0x6b
func LD_6B(cpu *CPU) {

	data := cpu.E

	cpu.L = data

	cpu.cycles += 4

}

// DEC B    code=0x05
func DEC_05(cpu *CPU) {
	res, flags := sub(cpu.B, 0x01)
	cpu.F = flags
	cpu.B = res
	cpu.cycles += 4
}

// LD A,(BC)    code=0x0a
func LD_0A(cpu *CPU) {

	data := cpu.loadU8(cpu.BC())

	cpu.A = data

	cpu.cycles += 8

}

// DEC A    code=0x3d
func DEC_3D(cpu *CPU) {
	res, flags := sub(cpu.A, 0x01)
	cpu.F = flags
	cpu.A = res
	cpu.cycles += 4
}

// LD E,E    code=0x5b
func LD_5B(cpu *CPU) {

	data := cpu.E

	cpu.E = data

	cpu.cycles += 4

}

// LD E,H    code=0x5c
func LD_5C(cpu *CPU) {

	data := cpu.H

	cpu.E = data

	cpu.cycles += 4

}

// XOR A,C    code=0xa9
func XOR_A9(cpu *CPU) {
	res := cpu.A ^ cpu.C
	var flags Flags
	if res == 0 {
		flags |= FLAGZ
	}
	cpu.F = FlagRegister(flags)
	cpu.A = res
	cpu.cycles += 4
}

// CP A,B    code=0xb8
func CP_B8(cpu *CPU) {
	_, flags := sub(cpu.A, cpu.B)
	cpu.F = flags
	cpu.cycles += 4
}

// JP C,a16    code=0xda
func JP_DA(cpu *CPU) {
	panic("TODO JP_DA")
}

// RST $20    code=0xe7
func RST_E7(cpu *CPU) {
	n := uint8(0x20)
	cpu.PushStack(cpu.PC)
	cpu.PC = concatU16(0x00, n)
}

// DEC (HL)    code=0x35
func DEC_35(cpu *CPU) {
	res, flags := sub(cpu.loadU8(cpu.HL()), 0x01)
	cpu.F = flags
	cpu.WriteMemory(cpu.HL(), res)
	cpu.cycles += 12
}

// LD C,D    code=0x4a
func LD_4A(cpu *CPU) {

	data := cpu.D

	cpu.C = data

	cpu.cycles += 4

}

// LD L,C    code=0x69
func LD_69(cpu *CPU) {

	data := cpu.C

	cpu.L = data

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

// XOR A,L    code=0xad
func XOR_AD(cpu *CPU) {
	res := cpu.A ^ cpu.L
	var flags Flags
	if res == 0 {
		flags |= FLAGZ
	}
	cpu.F = FlagRegister(flags)
	cpu.A = res
	cpu.cycles += 4
}

// JP HL    code=0xe9
func JP_E9(cpu *CPU) {
	panic("TODO JP_E9")
}

// LD SP,n16    code=0x31
func LD_31(cpu *CPU) {

	data := cpu.readU16(cpu.PC)

	cpu.SP = data

	cpu.cycles += 12

}

// DEC H    code=0x25
func DEC_25(cpu *CPU) {
	res, flags := sub(cpu.H, 0x01)
	cpu.F = flags
	cpu.H = res
	cpu.cycles += 4
}

// LD C,H    code=0x4c
func LD_4C(cpu *CPU) {

	data := cpu.H

	cpu.C = data

	cpu.cycles += 4

}

// LD D,E    code=0x53
func LD_53(cpu *CPU) {

	data := cpu.E

	cpu.D = data

	cpu.cycles += 4

}

// LD L,B    code=0x68
func LD_68(cpu *CPU) {

	data := cpu.B

	cpu.L = data

	cpu.cycles += 4

}

// SUB A,H    code=0x94
func SUB_94(cpu *CPU) {
	res, flags := sub(cpu.A, cpu.H)
	cpu.A = res
	cpu.F = flags
	cpu.cycles += 4
}

// ILLEGAL_E4     code=0xe4
func ILLEGAL_E4_E4(cpu *CPU) {
	panic("TODO ILLEGAL_E4_E4")
}

// LD C,C    code=0x49
func LD_49(cpu *CPU) {

	data := cpu.C

	cpu.C = data

	cpu.cycles += 4

}

// LD E,(HL)    code=0x5e
func LD_5E(cpu *CPU) {

	data := cpu.loadU8(cpu.HL())

	cpu.E = data

	cpu.cycles += 8

}

// CP A,D    code=0xba
func CP_BA(cpu *CPU) {
	_, flags := sub(cpu.A, cpu.D)
	cpu.F = flags
	cpu.cycles += 4
}

// RST $10    code=0xd7
func RST_D7(cpu *CPU) {
	n := uint8(0x10)
	cpu.PushStack(cpu.PC)
	cpu.PC = concatU16(0x00, n)
}

// LD (BC),A    code=0x02
func LD_02(cpu *CPU) {

	data := cpu.A

	cpu.WriteMemory(cpu.BC(), data)

	cpu.cycles += 8

}

// INC SP    code=0x33
func INC_33(cpu *CPU) {
	res, flags := add(cpu.SP, 0x01)
	cpu.F = flags
	cpu.SP = res
	cpu.cycles += 8
}

// LD D,H    code=0x54
func LD_54(cpu *CPU) {

	data := cpu.H

	cpu.D = data

	cpu.cycles += 4

}

// ADC A,A    code=0x8f
func ADC_8F(cpu *CPU) {
	panic("TODO ADC_8F")
}

// OR A,L    code=0xb5
func OR_B5(cpu *CPU) {
	res := cpu.A | cpu.L
	var flags Flags
	if res == 0 {
		flags |= FLAGZ
	}
	cpu.F = FlagRegister(flags)
	cpu.A = res
	cpu.cycles += 4
}

// INC B    code=0x04
func INC_04(cpu *CPU) {
	res, flags := add(cpu.B, 0x01)
	cpu.F = flags
	cpu.B = res
	cpu.cycles += 4
}

// LD A,L    code=0x7d
func LD_7D(cpu *CPU) {

	data := cpu.L

	cpu.A = data

	cpu.cycles += 4

}

// PUSH HL    code=0xe5
func PUSH_E5(cpu *CPU) {
	cpu.PushStack(cpu.HL())
	cpu.cycles += 16
}

// AND A,n8    code=0xe6
func AND_E6(cpu *CPU) {
	panic("TODO AND_E6")
}

// ILLEGAL_F4     code=0xf4
func ILLEGAL_F4_F4(cpu *CPU) {
	panic("TODO ILLEGAL_F4_F4")
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

// ADD HL,HL    code=0x29
func ADD_29(cpu *CPU) {
	lhs := cpu.HL()
	rhs := cpu.HL()
	res, flags := add(lhs, rhs)
	cpu.H, cpu.L = split(res)
	cpu.F = flags
	cpu.cycles += 8
}

// HALT     code=0x76
func HALT_76(cpu *CPU) {
	panic("TODO HALT_76")
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

// XOR A,E    code=0xab
func XOR_AB(cpu *CPU) {
	res := cpu.A ^ cpu.E
	var flags Flags
	if res == 0 {
		flags |= FLAGZ
	}
	cpu.F = FlagRegister(flags)
	cpu.A = res
	cpu.cycles += 4
}

// CP A,E    code=0xbb
func CP_BB(cpu *CPU) {
	_, flags := sub(cpu.A, cpu.E)
	cpu.F = flags
	cpu.cycles += 4
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

// LD (DE),A    code=0x12
func LD_12(cpu *CPU) {

	data := cpu.A

	cpu.WriteMemory(cpu.DE(), data)

	cpu.cycles += 8

}

// ADC A,E    code=0x8b
func ADC_8B(cpu *CPU) {
	panic("TODO ADC_8B")
}

// RST $28    code=0xef
func RST_EF(cpu *CPU) {
	n := uint8(0x28)
	cpu.PushStack(cpu.PC)
	cpu.PC = concatU16(0x00, n)
}

// DEC L    code=0x2d
func DEC_2D(cpu *CPU) {
	res, flags := sub(cpu.L, 0x01)
	cpu.F = flags
	cpu.L = res
	cpu.cycles += 4
}

// LD B,(HL)    code=0x46
func LD_46(cpu *CPU) {

	data := cpu.loadU8(cpu.HL())

	cpu.B = data

	cpu.cycles += 8

}

// LD H,H    code=0x64
func LD_64(cpu *CPU) {

	data := cpu.H

	cpu.H = data

	cpu.cycles += 4

}

// AND A,H    code=0xa4
func AND_A4(cpu *CPU) {
	panic("TODO AND_A4")
}

// XOR A,A    code=0xaf
func XOR_AF(cpu *CPU) {
	res := cpu.A ^ cpu.A
	var flags Flags
	if res == 0 {
		flags |= FLAGZ
	}
	cpu.F = FlagRegister(flags)
	cpu.A = res
	cpu.cycles += 4
}

// CALL NC,a16    code=0xd4
func CALL_D4(cpu *CPU) {
	lsb := cpu.readU8(cpu.PC)
	msb := cpu.readU8(cpu.PC)
	nn := concatU16(msb, lsb)
	if !cpu.F.HasCarry() {
		cpu.PushStack(cpu.PC)
		cpu.PC = nn
		cpu.cycles += 24
	} else {
		cpu.cycles += 12
	}
}

// RLCA     code=0x07
func RLCA_07(cpu *CPU) {
	cpu.A, cpu.F = rotate(cpu.A, 0, cpu.F, true)
	cpu.cycles += 4
}

// LD (HL),B    code=0x70
func LD_70(cpu *CPU) {

	data := cpu.B

	cpu.WriteMemory(cpu.HL(), data)

	cpu.cycles += 8

}

// AND A,L    code=0xa5
func AND_A5(cpu *CPU) {
	panic("TODO AND_A5")
}

// LDH A,(C)    code=0xf2
func LDH_F2(cpu *CPU) {
	value := cpu.loadU8(concatU16(0xFF, cpu.C))
	cpu.A = value
	cpu.cycles += 8
}

// LD D,D    code=0x52
func LD_52(cpu *CPU) {

	data := cpu.D

	cpu.D = data

	cpu.cycles += 4

}

// JP a16    code=0xc3
func JP_C3(cpu *CPU) {
	panic("TODO JP_C3")
}

// LD D,n8    code=0x16
func LD_16(cpu *CPU) {

	data := cpu.readU8(cpu.PC)

	cpu.D = data

	cpu.cycles += 8

}

// RLA     code=0x17
func RLA_17(cpu *CPU) {
	cpu.A, cpu.F = rotate(cpu.A, 0, cpu.F, false)
	cpu.cycles += 4
}

// AND A,B    code=0xa0
func AND_A0(cpu *CPU) {
	panic("TODO AND_A0")
}

// AND A,D    code=0xa2
func AND_A2(cpu *CPU) {
	panic("TODO AND_A2")
}

// EI     code=0xfb
func EI_FB(cpu *CPU) {
	panic("TODO EI_FB")
}

// ILLEGAL_FC     code=0xfc
func ILLEGAL_FC_FC(cpu *CPU) {
	panic("TODO ILLEGAL_FC_FC")
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

// ADC A,(HL)    code=0x8e
func ADC_8E(cpu *CPU) {
	panic("TODO ADC_8E")
}

// CALL Z,a16    code=0xcc
func CALL_CC(cpu *CPU) {
	lsb := cpu.readU8(cpu.PC)
	msb := cpu.readU8(cpu.PC)
	nn := concatU16(msb, lsb)
	if cpu.F.HasZero() {
		cpu.PushStack(cpu.PC)
		cpu.PC = nn
		cpu.cycles += 24
	} else {
		cpu.cycles += 12
	}
}

// LD H,n8    code=0x26
func LD_26(cpu *CPU) {

	data := cpu.readU8(cpu.PC)

	cpu.H = data

	cpu.cycles += 8

}

// RST $00    code=0xc7
func RST_C7(cpu *CPU) {
	n := uint8(0x0)
	cpu.PushStack(cpu.PC)
	cpu.PC = concatU16(0x00, n)
}

// LDH (C),A    code=0xe2
func LDH_E2(cpu *CPU) {
	value := cpu.A
	cpu.WriteMemory(concatU16(0xFF, cpu.C), value)
	cpu.cycles += 8
}

// JR e8    code=0x18
func JR_18(cpu *CPU) {
	e := cpu.readI8(cpu.PC)
	if true {
		cpu.PC, cpu.F = add(cpu.PC, e)
		cpu.cycles += 12
	} else {
		cpu.cycles += 12
	}
}

// LD A,(HL-)    code=0x3a
func LD_3A(cpu *CPU) {

	data := cpu.loadU8(cpu.HL())

	cpu.A = data

	cpu.cycles += 8

}

// LD D,L    code=0x55
func LD_55(cpu *CPU) {

	data := cpu.L

	cpu.D = data

	cpu.cycles += 4

}

// XOR A,H    code=0xac
func XOR_AC(cpu *CPU) {
	res := cpu.A ^ cpu.H
	var flags Flags
	if res == 0 {
		flags |= FLAGZ
	}
	cpu.F = FlagRegister(flags)
	cpu.A = res
	cpu.cycles += 4
}

// LD (HL),C    code=0x71
func LD_71(cpu *CPU) {

	data := cpu.C

	cpu.WriteMemory(cpu.HL(), data)

	cpu.cycles += 8

}

// LD SP,HL    code=0xf9
func LD_F9(cpu *CPU) {

	data := cpu.HL()

	cpu.SP = data

	cpu.cycles += 8

}

// LD E,A    code=0x5f
func LD_5F(cpu *CPU) {

	data := cpu.A

	cpu.E = data

	cpu.cycles += 4

}

// OR A,A    code=0xb7
func OR_B7(cpu *CPU) {
	res := cpu.A | cpu.A
	var flags Flags
	if res == 0 {
		flags |= FLAGZ
	}
	cpu.F = FlagRegister(flags)
	cpu.A = res
	cpu.cycles += 4
}

// RET     code=0xc9
func RET_C9(cpu *CPU) {
	if true {
		cpu.PC = cpu.PopStack()
		cpu.cycles += 16
	} else {
		cpu.cycles += 16
	}
}

// XOR A,n8    code=0xee
func XOR_EE(cpu *CPU) {
	res := cpu.A ^ cpu.readU8(cpu.PC)
	var flags Flags
	if res == 0 {
		flags |= FLAGZ
	}
	cpu.F = FlagRegister(flags)
	cpu.A = res
	cpu.cycles += 8
}

// DI     code=0xf3
func DI_F3(cpu *CPU) {
	panic("TODO DI_F3")
}

// JR C,e8    code=0x38
func JR_38(cpu *CPU) {
	e := cpu.readI8(cpu.PC)
	if cpu.F.HasCarry() {
		cpu.PC, cpu.F = add(cpu.PC, e)
		cpu.cycles += 12
	} else {
		cpu.cycles += 8
	}
}

// JP NZ,a16    code=0xc2
func JP_C2(cpu *CPU) {
	panic("TODO JP_C2")
}

// PUSH BC    code=0xc5
func PUSH_C5(cpu *CPU) {
	cpu.PushStack(cpu.BC())
	cpu.cycles += 16
}

// LD DE,n16    code=0x11
func LD_11(cpu *CPU) {

	data := cpu.readU16(cpu.PC)

	cpu.D, cpu.E = split(data)

	cpu.cycles += 12

}

// INC DE    code=0x13
func INC_13(cpu *CPU) {
	res, flags := add(cpu.DE(), 0x01)
	cpu.F = flags
	cpu.D, cpu.E = split(res)
	cpu.cycles += 8
}

// INC L    code=0x2c
func INC_2C(cpu *CPU) {
	res, flags := add(cpu.L, 0x01)
	cpu.F = flags
	cpu.L = res
	cpu.cycles += 4
}

// LD L,n8    code=0x2e
func LD_2E(cpu *CPU) {

	data := cpu.readU8(cpu.PC)

	cpu.L = data

	cpu.cycles += 8

}

// LD A,C    code=0x79
func LD_79(cpu *CPU) {

	data := cpu.C

	cpu.A = data

	cpu.cycles += 4

}

// ADC A,L    code=0x8d
func ADC_8D(cpu *CPU) {
	panic("TODO ADC_8D")
}

// SUB A,A    code=0x97
func SUB_97(cpu *CPU) {
	res, flags := sub(cpu.A, cpu.A)
	cpu.A = res
	cpu.F = flags
	cpu.cycles += 4
}

// ILLEGAL_DD     code=0xdd
func ILLEGAL_DD_DD(cpu *CPU) {
	panic("TODO ILLEGAL_DD_DD")
}

// LD D,C    code=0x51
func LD_51(cpu *CPU) {

	data := cpu.C

	cpu.D = data

	cpu.cycles += 4

}

// LD H,L    code=0x65
func LD_65(cpu *CPU) {

	data := cpu.L

	cpu.H = data

	cpu.cycles += 4

}

// OR A,C    code=0xb1
func OR_B1(cpu *CPU) {
	res := cpu.A | cpu.C
	var flags Flags
	if res == 0 {
		flags |= FLAGZ
	}
	cpu.F = FlagRegister(flags)
	cpu.A = res
	cpu.cycles += 4
}

// LD (a16),A    code=0xea
func LD_EA(cpu *CPU) {

	data := cpu.A

	cpu.WriteMemory(cpu.readU16(cpu.PC), data)

	cpu.cycles += 16

}

// ADC A,B    code=0x88
func ADC_88(cpu *CPU) {
	panic("TODO ADC_88")
}

// RST $18    code=0xdf
func RST_DF(cpu *CPU) {
	n := uint8(0x18)
	cpu.PushStack(cpu.PC)
	cpu.PC = concatU16(0x00, n)
}

// JR NZ,e8    code=0x20
func JR_20(cpu *CPU) {
	e := cpu.readI8(cpu.PC)
	if !cpu.F.HasZero() {
		cpu.PC, cpu.F = add(cpu.PC, e)
		cpu.cycles += 12
	} else {
		cpu.cycles += 8
	}
}

// LD B,A    code=0x47
func LD_47(cpu *CPU) {

	data := cpu.A

	cpu.B = data

	cpu.cycles += 4

}

// LD A,E    code=0x7b
func LD_7B(cpu *CPU) {

	data := cpu.E

	cpu.A = data

	cpu.cycles += 4

}

// OR A,D    code=0xb2
func OR_B2(cpu *CPU) {
	res := cpu.A | cpu.D
	var flags Flags
	if res == 0 {
		flags |= FLAGZ
	}
	cpu.F = FlagRegister(flags)
	cpu.A = res
	cpu.cycles += 4
}

// CP A,A    code=0xbf
func CP_BF(cpu *CPU) {
	_, flags := sub(cpu.A, cpu.A)
	cpu.F = flags
	cpu.cycles += 4
}

// SBC A,n8    code=0xde
func SBC_DE(cpu *CPU) {
	panic("TODO SBC_DE")
}

// CP A,n8    code=0xfe
func CP_FE(cpu *CPU) {
	_, flags := sub(cpu.A, cpu.readU8(cpu.PC))
	cpu.F = flags
	cpu.cycles += 8
}

// DEC E    code=0x1d
func DEC_1D(cpu *CPU) {
	res, flags := sub(cpu.E, 0x01)
	cpu.F = flags
	cpu.E = res
	cpu.cycles += 4
}

// RRA     code=0x1f
func RRA_1F(cpu *CPU) {
	cpu.A, cpu.F = rotate(cpu.A, 1, cpu.F, false)
	cpu.cycles += 4
}

// LD (HL),D    code=0x72
func LD_72(cpu *CPU) {

	data := cpu.D

	cpu.WriteMemory(cpu.HL(), data)

	cpu.cycles += 8

}

// ADC A,H    code=0x8c
func ADC_8C(cpu *CPU) {
	panic("TODO ADC_8C")
}

// CP A,(HL)    code=0xbe
func CP_BE(cpu *CPU) {
	_, flags := sub(cpu.A, cpu.loadU8(cpu.HL()))
	cpu.F = flags
	cpu.cycles += 8
}

// JR Z,e8    code=0x28
func JR_28(cpu *CPU) {
	e := cpu.readI8(cpu.PC)
	if cpu.F.HasZero() {
		cpu.PC, cpu.F = add(cpu.PC, e)
		cpu.cycles += 12
	} else {
		cpu.cycles += 8
	}
}

// CP A,C    code=0xb9
func CP_B9(cpu *CPU) {
	_, flags := sub(cpu.A, cpu.C)
	cpu.F = flags
	cpu.cycles += 4
}

// INC A    code=0x3c
func INC_3C(cpu *CPU) {
	res, flags := add(cpu.A, 0x01)
	cpu.F = flags
	cpu.A = res
	cpu.cycles += 4
}

// LD D,B    code=0x50
func LD_50(cpu *CPU) {

	data := cpu.B

	cpu.D = data

	cpu.cycles += 4

}

// LD A,(HL)    code=0x7e
func LD_7E(cpu *CPU) {

	data := cpu.loadU8(cpu.HL())

	cpu.A = data

	cpu.cycles += 8

}

// RET Z    code=0xc8
func RET_C8(cpu *CPU) {
	if cpu.F.HasZero() {
		cpu.PC = cpu.PopStack()
		cpu.cycles += 20
	} else {
		cpu.cycles += 8
	}
}

// PREFIX     code=0xcb
func PREFIX_CB(cpu *CPU) {
	cpu.prefix = true
	cpu.cycles += 4
}

// RST $08    code=0xcf
func RST_CF(cpu *CPU) {
	n := uint8(0x8)
	cpu.PushStack(cpu.PC)
	cpu.PC = concatU16(0x00, n)
}

// RET C    code=0xd8
func RET_D8(cpu *CPU) {
	if cpu.F.HasCarry() {
		cpu.PC = cpu.PopStack()
		cpu.cycles += 20
	} else {
		cpu.cycles += 8
	}
}

// LD (a16),SP    code=0x08
func LD_08(cpu *CPU) {

	data := cpu.SP

	cpu.WriteMemory(cpu.readU16(cpu.PC), data)

	cpu.cycles += 20

}

// INC C    code=0x0c
func INC_0C(cpu *CPU) {
	res, flags := add(cpu.C, 0x01)
	cpu.F = flags
	cpu.C = res
	cpu.cycles += 4
}

// SCF     code=0x37
func SCF_37(cpu *CPU) {
	panic("TODO SCF_37")
}

// LD B,E    code=0x43
func LD_43(cpu *CPU) {

	data := cpu.E

	cpu.B = data

	cpu.cycles += 4

}

// CP A,L    code=0xbd
func CP_BD(cpu *CPU) {
	_, flags := sub(cpu.A, cpu.L)
	cpu.F = flags
	cpu.cycles += 4
}

// DEC DE    code=0x1b
func DEC_1B(cpu *CPU) {
	res, flags := sub(cpu.DE(), 0x01)
	cpu.F = flags
	cpu.D, cpu.E = split(res)
	cpu.cycles += 8
}

// DEC SP    code=0x3b
func DEC_3B(cpu *CPU) {
	res, flags := sub(cpu.SP, 0x01)
	cpu.F = flags
	cpu.SP = res
	cpu.cycles += 8
}

// LD (HL),E    code=0x73
func LD_73(cpu *CPU) {

	data := cpu.E

	cpu.WriteMemory(cpu.HL(), data)

	cpu.cycles += 8

}

// JP Z,a16    code=0xca
func JP_CA(cpu *CPU) {
	panic("TODO JP_CA")
}

// RET NC    code=0xd0
func RET_D0(cpu *CPU) {
	if !cpu.F.HasCarry() {
		cpu.PC = cpu.PopStack()
		cpu.cycles += 20
	} else {
		cpu.cycles += 8
	}
}

// LD HL,SP+,e8    code=0xf8
func LD_F8(cpu *CPU) {

	e := cpu.readI8(cpu.PC)
	res, flags := add(cpu.SP, e)
	cpu.H, cpu.L = split(res)
	cpu.F = flags
	cpu.cycles += 12

}

// LD L,D    code=0x6a
func LD_6A(cpu *CPU) {

	data := cpu.D

	cpu.L = data

	cpu.cycles += 4

}

// LD L,(HL)    code=0x6e
func LD_6E(cpu *CPU) {

	data := cpu.loadU8(cpu.HL())

	cpu.L = data

	cpu.cycles += 8

}

// ILLEGAL_EB     code=0xeb
func ILLEGAL_EB_EB(cpu *CPU) {
	panic("TODO ILLEGAL_EB_EB")
}

// POP AF    code=0xf1
func POP_F1(cpu *CPU) {
	value := cpu.PopStack()
	msb, lsb := split(value)
	cpu.A, cpu.F = msb, FlagRegister(lsb)
	cpu.cycles += 12
}

// LD E,C    code=0x59
func LD_59(cpu *CPU) {

	data := cpu.C

	cpu.E = data

	cpu.cycles += 4

}

// LD L,A    code=0x6f
func LD_6F(cpu *CPU) {

	data := cpu.A

	cpu.L = data

	cpu.cycles += 4

}

// SUB A,(HL)    code=0x96
func SUB_96(cpu *CPU) {
	res, flags := sub(cpu.A, cpu.loadU8(cpu.HL()))
	cpu.A = res
	cpu.F = flags
	cpu.cycles += 8
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

// CPL     code=0x2f
func CPL_2F(cpu *CPU) {
	panic("TODO CPL_2F")
}

// LD A,H    code=0x7c
func LD_7C(cpu *CPU) {

	data := cpu.H

	cpu.A = data

	cpu.cycles += 4

}

// INC BC    code=0x03
func INC_03(cpu *CPU) {
	res, flags := add(cpu.BC(), 0x01)
	cpu.F = flags
	cpu.B, cpu.C = split(res)
	cpu.cycles += 8
}

// JR NC,e8    code=0x30
func JR_30(cpu *CPU) {
	e := cpu.readI8(cpu.PC)
	if !cpu.F.HasCarry() {
		cpu.PC, cpu.F = add(cpu.PC, e)
		cpu.cycles += 12
	} else {
		cpu.cycles += 8
	}
}

// LD C,B    code=0x48
func LD_48(cpu *CPU) {

	data := cpu.B

	cpu.C = data

	cpu.cycles += 4

}

// LD H,(HL)    code=0x66
func LD_66(cpu *CPU) {

	data := cpu.loadU8(cpu.HL())

	cpu.H = data

	cpu.cycles += 8

}

// SBC A,E    code=0x9b
func SBC_9B(cpu *CPU) {
	panic("TODO SBC_9B")
}

// AND A,A    code=0xa7
func AND_A7(cpu *CPU) {
	panic("TODO AND_A7")
}

// LD B,C    code=0x41
func LD_41(cpu *CPU) {

	data := cpu.C

	cpu.B = data

	cpu.cycles += 4

}

// LD B,L    code=0x45
func LD_45(cpu *CPU) {

	data := cpu.L

	cpu.B = data

	cpu.cycles += 4

}

// SUB A,B    code=0x90
func SUB_90(cpu *CPU) {
	res, flags := sub(cpu.A, cpu.B)
	cpu.A = res
	cpu.F = flags
	cpu.cycles += 4
}

// AND A,(HL)    code=0xa6
func AND_A6(cpu *CPU) {
	panic("TODO AND_A6")
}

// POP DE    code=0xd1
func POP_D1(cpu *CPU) {
	value := cpu.PopStack()
	cpu.D, cpu.E = split(value)
	cpu.cycles += 12
}

// RETI     code=0xd9
func RETI_D9(cpu *CPU) {
	panic("TODO RETI_D9")
}

// LD C,L    code=0x4d
func LD_4D(cpu *CPU) {

	data := cpu.L

	cpu.C = data

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

// SBC A,H    code=0x9c
func SBC_9C(cpu *CPU) {
	panic("TODO SBC_9C")
}

// CALL C,a16    code=0xdc
func CALL_DC(cpu *CPU) {
	lsb := cpu.readU8(cpu.PC)
	msb := cpu.readU8(cpu.PC)
	nn := concatU16(msb, lsb)
	if cpu.F.HasCarry() {
		cpu.PushStack(cpu.PC)
		cpu.PC = nn
		cpu.cycles += 24
	} else {
		cpu.cycles += 12
	}
}

// ILLEGAL_EC     code=0xec
func ILLEGAL_EC_EC(cpu *CPU) {
	panic("TODO ILLEGAL_EC_EC")
}

// LD A,(HL+)    code=0x2a
func LD_2A(cpu *CPU) {

	data := cpu.loadU8(cpu.HL())

	cpu.A = data

	cpu.cycles += 8

}

// NOP     code=0x00
func NOP_00(cpu *CPU) {
	panic("TODO NOP_00")
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

// LD B,B    code=0x40
func LD_40(cpu *CPU) {

	data := cpu.B

	cpu.B = data

	cpu.cycles += 4

}

// LD A,A    code=0x7f
func LD_7F(cpu *CPU) {

	data := cpu.A

	cpu.A = data

	cpu.cycles += 4

}

// INC H    code=0x24
func INC_24(cpu *CPU) {
	res, flags := add(cpu.H, 0x01)
	cpu.F = flags
	cpu.H = res
	cpu.cycles += 4
}

// LDH (a8),A    code=0xe0
func LDH_E0(cpu *CPU) {
	value := cpu.A
	cpu.WriteMemory(concatU16(0xFF, cpu.readU8(cpu.PC)), value)
	cpu.cycles += 12
}

// LDH A,(a8)    code=0xf0
func LDH_F0(cpu *CPU) {
	value := cpu.loadU8(concatU16(0xFF, cpu.readU8(cpu.PC)))
	cpu.A = value
	cpu.cycles += 12
}

// CCF     code=0x3f
func CCF_3F(cpu *CPU) {
	panic("TODO CCF_3F")
}

// LD H,A    code=0x67
func LD_67(cpu *CPU) {

	data := cpu.A

	cpu.H = data

	cpu.cycles += 4

}

// OR A,B    code=0xb0
func OR_B0(cpu *CPU) {
	res := cpu.A | cpu.B
	var flags Flags
	if res == 0 {
		flags |= FLAGZ
	}
	cpu.F = FlagRegister(flags)
	cpu.A = res
	cpu.cycles += 4
}

// ILLEGAL_DB     code=0xdb
func ILLEGAL_DB_DB(cpu *CPU) {
	panic("TODO ILLEGAL_DB_DB")
}

// LD D,(HL)    code=0x56
func LD_56(cpu *CPU) {

	data := cpu.loadU8(cpu.HL())

	cpu.D = data

	cpu.cycles += 8

}

// ILLEGAL_FD     code=0xfd
func ILLEGAL_FD_FD(cpu *CPU) {
	panic("TODO ILLEGAL_FD_FD")
}

// LD A,(DE)    code=0x1a
func LD_1A(cpu *CPU) {

	data := cpu.loadU8(cpu.DE())

	cpu.A = data

	cpu.cycles += 8

}

// DEC C    code=0x0d
func DEC_0D(cpu *CPU) {
	res, flags := sub(cpu.C, 0x01)
	cpu.F = flags
	cpu.C = res
	cpu.cycles += 4
}

// SBC A,A    code=0x9f
func SBC_9F(cpu *CPU) {
	panic("TODO SBC_9F")
}

// INC D    code=0x14
func INC_14(cpu *CPU) {
	res, flags := add(cpu.D, 0x01)
	cpu.F = flags
	cpu.D = res
	cpu.cycles += 4
}

// LD (HL),n8    code=0x36
func LD_36(cpu *CPU) {

	data := cpu.readU8(cpu.PC)

	cpu.WriteMemory(cpu.HL(), data)

	cpu.cycles += 12

}

// LD B,H    code=0x44
func LD_44(cpu *CPU) {

	data := cpu.H

	cpu.B = data

	cpu.cycles += 4

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

// ADD A,D    code=0x82
func ADD_82(cpu *CPU) {
	lhs := cpu.A
	rhs := cpu.D
	res, flags := add(lhs, rhs)
	cpu.A = res
	cpu.F = flags
	cpu.cycles += 4
}

// LD L,L    code=0x6d
func LD_6D(cpu *CPU) {

	data := cpu.L

	cpu.L = data

	cpu.cycles += 4

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

// OR A,E    code=0xb3
func OR_B3(cpu *CPU) {
	res := cpu.A | cpu.E
	var flags Flags
	if res == 0 {
		flags |= FLAGZ
	}
	cpu.F = FlagRegister(flags)
	cpu.A = res
	cpu.cycles += 4
}

// DEC BC    code=0x0b
func DEC_0B(cpu *CPU) {
	res, flags := sub(cpu.BC(), 0x01)
	cpu.F = flags
	cpu.B, cpu.C = split(res)
	cpu.cycles += 8
}

// LD L,H    code=0x6c
func LD_6C(cpu *CPU) {

	data := cpu.H

	cpu.L = data

	cpu.cycles += 4

}

// OR A,H    code=0xb4
func OR_B4(cpu *CPU) {
	res := cpu.A | cpu.H
	var flags Flags
	if res == 0 {
		flags |= FLAGZ
	}
	cpu.F = FlagRegister(flags)
	cpu.A = res
	cpu.cycles += 4
}

// OR A,(HL)    code=0xb6
func OR_B6(cpu *CPU) {
	res := cpu.A | cpu.loadU8(cpu.HL())
	var flags Flags
	if res == 0 {
		flags |= FLAGZ
	}
	cpu.F = FlagRegister(flags)
	cpu.A = res
	cpu.cycles += 8
}

// POP BC    code=0xc1
func POP_C1(cpu *CPU) {
	value := cpu.PopStack()
	cpu.B, cpu.C = split(value)
	cpu.cycles += 12
}

// SBC A,L    code=0x9d
func SBC_9D(cpu *CPU) {
	panic("TODO SBC_9D")
}

// CP A,H    code=0xbc
func CP_BC(cpu *CPU) {
	_, flags := sub(cpu.A, cpu.H)
	cpu.F = flags
	cpu.cycles += 4
}

// OR A,n8    code=0xf6
func OR_F6(cpu *CPU) {
	res := cpu.A | cpu.readU8(cpu.PC)
	var flags Flags
	if res == 0 {
		flags |= FLAGZ
	}
	cpu.F = FlagRegister(flags)
	cpu.A = res
	cpu.cycles += 8
}

// DEC D    code=0x15
func DEC_15(cpu *CPU) {
	res, flags := sub(cpu.D, 0x01)
	cpu.F = flags
	cpu.D = res
	cpu.cycles += 4
}

// INC (HL)    code=0x34
func INC_34(cpu *CPU) {
	res, flags := add(cpu.loadU8(cpu.HL()), 0x01)
	cpu.F = flags
	cpu.WriteMemory(cpu.HL(), res)
	cpu.cycles += 12
}

// RST $30    code=0xf7
func RST_F7(cpu *CPU) {
	n := uint8(0x30)
	cpu.PushStack(cpu.PC)
	cpu.PC = concatU16(0x00, n)
}

// INC E    code=0x1c
func INC_1C(cpu *CPU) {
	res, flags := add(cpu.E, 0x01)
	cpu.F = flags
	cpu.E = res
	cpu.cycles += 4
}

// LD H,D    code=0x62
func LD_62(cpu *CPU) {

	data := cpu.D

	cpu.H = data

	cpu.cycles += 4

}

// ADC A,D    code=0x8a
func ADC_8A(cpu *CPU) {
	panic("TODO ADC_8A")
}

// XOR A,D    code=0xaa
func XOR_AA(cpu *CPU) {
	res := cpu.A ^ cpu.D
	var flags Flags
	if res == 0 {
		flags |= FLAGZ
	}
	cpu.F = FlagRegister(flags)
	cpu.A = res
	cpu.cycles += 4
}

// XOR A,(HL)    code=0xae
func XOR_AE(cpu *CPU) {
	res := cpu.A ^ cpu.loadU8(cpu.HL())
	var flags Flags
	if res == 0 {
		flags |= FLAGZ
	}
	cpu.F = FlagRegister(flags)
	cpu.A = res
	cpu.cycles += 8
}

// LD A,(a16)    code=0xfa
func LD_FA(cpu *CPU) {

	data := cpu.loadU8(cpu.readU16(cpu.PC))

	cpu.A = data

	cpu.cycles += 16

}

// LD (HL),H    code=0x74
func LD_74(cpu *CPU) {

	data := cpu.H

	cpu.WriteMemory(cpu.HL(), data)

	cpu.cycles += 8

}

// LD A,B    code=0x78
func LD_78(cpu *CPU) {

	data := cpu.B

	cpu.A = data

	cpu.cycles += 4

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

// ADC A,C    code=0x89
func ADC_89(cpu *CPU) {
	panic("TODO ADC_89")
}

// SUB A,E    code=0x93
func SUB_93(cpu *CPU) {
	res, flags := sub(cpu.A, cpu.E)
	cpu.A = res
	cpu.F = flags
	cpu.cycles += 4
}

// SBC A,B    code=0x98
func SBC_98(cpu *CPU) {
	panic("TODO SBC_98")
}

// LD E,n8    code=0x1e
func LD_1E(cpu *CPU) {

	data := cpu.readU8(cpu.PC)

	cpu.E = data

	cpu.cycles += 8

}

// LD HL,n16    code=0x21
func LD_21(cpu *CPU) {

	data := cpu.readU16(cpu.PC)

	cpu.H, cpu.L = split(data)

	cpu.cycles += 12

}

// LD B,D    code=0x42
func LD_42(cpu *CPU) {

	data := cpu.D

	cpu.B = data

	cpu.cycles += 4

}

// SUB A,L    code=0x95
func SUB_95(cpu *CPU) {
	res, flags := sub(cpu.A, cpu.L)
	cpu.A = res
	cpu.F = flags
	cpu.cycles += 4
}

// POP HL    code=0xe1
func POP_E1(cpu *CPU) {
	value := cpu.PopStack()
	cpu.H, cpu.L = split(value)
	cpu.cycles += 12
}

// ILLEGAL_ED     code=0xed
func ILLEGAL_ED_ED(cpu *CPU) {
	panic("TODO ILLEGAL_ED_ED")
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

// CALL NZ,a16    code=0xc4
func CALL_C4(cpu *CPU) {
	lsb := cpu.readU8(cpu.PC)
	msb := cpu.readU8(cpu.PC)
	nn := concatU16(msb, lsb)
	if !cpu.F.HasZero() {
		cpu.PushStack(cpu.PC)
		cpu.PC = nn
		cpu.cycles += 24
	} else {
		cpu.cycles += 12
	}
}

// LD BC,n16    code=0x01
func LD_01(cpu *CPU) {

	data := cpu.readU16(cpu.PC)

	cpu.B, cpu.C = split(data)

	cpu.cycles += 12

}

// STOP n8    code=0x10
func STOP_10(cpu *CPU) {
	cpu.err = ErrNoMoreInstructions
	cpu.cycles += 4
}

// INC HL    code=0x23
func INC_23(cpu *CPU) {
	res, flags := add(cpu.HL(), 0x01)
	cpu.F = flags
	cpu.H, cpu.L = split(res)
	cpu.cycles += 8
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

// SBC A,(HL)    code=0x9e
func SBC_9E(cpu *CPU) {
	panic("TODO SBC_9E")
}

// SBC A,C    code=0x99
func SBC_99(cpu *CPU) {
	panic("TODO SBC_99")
}

// JP NC,a16    code=0xd2
func JP_D2(cpu *CPU) {
	panic("TODO JP_D2")
}

// ILLEGAL_D3     code=0xd3
func ILLEGAL_D3_D3(cpu *CPU) {
	panic("TODO ILLEGAL_D3_D3")
}

// PUSH AF    code=0xf5
func PUSH_F5(cpu *CPU) {
	cpu.PushStack(cpu.AF())
	cpu.cycles += 16
}

// LD B,n8    code=0x06
func LD_06(cpu *CPU) {

	data := cpu.readU8(cpu.PC)

	cpu.B = data

	cpu.cycles += 8

}

// LD C,A    code=0x4f
func LD_4F(cpu *CPU) {

	data := cpu.A

	cpu.C = data

	cpu.cycles += 4

}

// LD E,L    code=0x5d
func LD_5D(cpu *CPU) {

	data := cpu.L

	cpu.E = data

	cpu.cycles += 4

}

// LD E,D    code=0x5a
func LD_5A(cpu *CPU) {

	data := cpu.D

	cpu.E = data

	cpu.cycles += 4

}

// SUB A,C    code=0x91
func SUB_91(cpu *CPU) {
	res, flags := sub(cpu.A, cpu.C)
	cpu.A = res
	cpu.F = flags
	cpu.cycles += 4
}

// DAA     code=0x27
func DAA_27(cpu *CPU) {
	panic("TODO DAA_27")
}

// LD H,C    code=0x61
func LD_61(cpu *CPU) {

	data := cpu.C

	cpu.H = data

	cpu.cycles += 4

}

// SUB A,D    code=0x92
func SUB_92(cpu *CPU) {
	res, flags := sub(cpu.A, cpu.D)
	cpu.A = res
	cpu.F = flags
	cpu.cycles += 4
}

// XOR A,B    code=0xa8
func XOR_A8(cpu *CPU) {
	res := cpu.A ^ cpu.B
	var flags Flags
	if res == 0 {
		flags |= FLAGZ
	}
	cpu.F = FlagRegister(flags)
	cpu.A = res
	cpu.cycles += 4
}

// CALL a16    code=0xcd
func CALL_CD(cpu *CPU) {
	lsb := cpu.readU8(cpu.PC)
	msb := cpu.readU8(cpu.PC)
	nn := concatU16(msb, lsb)
	if true {
		cpu.PushStack(cpu.PC)
		cpu.PC = nn
		cpu.cycles += 24
	} else {
		cpu.cycles += 12
	}
}

// LD H,E    code=0x63
func LD_63(cpu *CPU) {

	data := cpu.E

	cpu.H = data

	cpu.cycles += 4

}

// SUB A,n8    code=0xd6
func SUB_D6(cpu *CPU) {
	res, flags := sub(cpu.A, cpu.readU8(cpu.PC))
	cpu.A = res
	cpu.F = flags
	cpu.cycles += 8
}

// ILLEGAL_E3     code=0xe3
func ILLEGAL_E3_E3(cpu *CPU) {
	panic("TODO ILLEGAL_E3_E3")
}

// LD A,n8    code=0x3e
func LD_3E(cpu *CPU) {

	data := cpu.readU8(cpu.PC)

	cpu.A = data

	cpu.cycles += 8

}

// RES 0,E    code=0x83
func RES_83(cpu *CPU) {
	panic("TODO RES_83")
}

// SET 5,H    code=0xec
func SET_EC(cpu *CPU) {
	panic("TODO SET_EC")
}

// SET 6,L    code=0xf5
func SET_F5(cpu *CPU) {
	panic("TODO SET_F5")
}

// SET 7,E    code=0xfb
func SET_FB(cpu *CPU) {
	panic("TODO SET_FB")
}

// BIT 4,B    code=0x60
func BIT_60(cpu *CPU) {
	value := cpu.B
	var flags Flags
	if (value & (1 << 4)) == 0 {
		flags |= FLAGZ
	}
	flags |= FLAGH
	cpu.F = FlagRegister(flags)

	cpu.cycles += 8
}

// BIT 7,B    code=0x78
func BIT_78(cpu *CPU) {
	value := cpu.B
	var flags Flags
	if (value & (1 << 7)) == 0 {
		flags |= FLAGZ
	}
	flags |= FLAGH
	cpu.F = FlagRegister(flags)

	cpu.cycles += 8
}

// SET 0,E    code=0xc3
func SET_C3(cpu *CPU) {
	panic("TODO SET_C3")
}

// BIT 2,D    code=0x52
func BIT_52(cpu *CPU) {
	value := cpu.D
	var flags Flags
	if (value & (1 << 2)) == 0 {
		flags |= FLAGZ
	}
	flags |= FLAGH
	cpu.F = FlagRegister(flags)

	cpu.cycles += 8
}

// BIT 3,H    code=0x5c
func BIT_5C(cpu *CPU) {
	value := cpu.H
	var flags Flags
	if (value & (1 << 3)) == 0 {
		flags |= FLAGZ
	}
	flags |= FLAGH
	cpu.F = FlagRegister(flags)

	cpu.cycles += 8
}

// SLA H    code=0x24
func SLA_24(cpu *CPU) {
	panic("TODO SLA_24")
}

// BIT 0,C    code=0x41
func BIT_41(cpu *CPU) {
	value := cpu.C
	var flags Flags
	if (value & (1 << 0)) == 0 {
		flags |= FLAGZ
	}
	flags |= FLAGH
	cpu.F = FlagRegister(flags)

	cpu.cycles += 8
}

// BIT 1,E    code=0x4b
func BIT_4B(cpu *CPU) {
	value := cpu.E
	var flags Flags
	if (value & (1 << 1)) == 0 {
		flags |= FLAGZ
	}
	flags |= FLAGH
	cpu.F = FlagRegister(flags)

	cpu.cycles += 8
}

// RES 6,L    code=0xb5
func RES_B5(cpu *CPU) {
	panic("TODO RES_B5")
}

// RR H    code=0x1c
func RR_1C(cpu *CPU) {
	res, flags := rotate(cpu.H, 1, cpu.F, false)
	cpu.H = res
	cpu.F = flags

	cpu.cycles += 8
}

// BIT 2,A    code=0x57
func BIT_57(cpu *CPU) {
	value := cpu.A
	var flags Flags
	if (value & (1 << 2)) == 0 {
		flags |= FLAGZ
	}
	flags |= FLAGH
	cpu.F = FlagRegister(flags)

	cpu.cycles += 8
}

// BIT 6,C    code=0x71
func BIT_71(cpu *CPU) {
	value := cpu.C
	var flags Flags
	if (value & (1 << 6)) == 0 {
		flags |= FLAGZ
	}
	flags |= FLAGH
	cpu.F = FlagRegister(flags)

	cpu.cycles += 8
}

// RES 3,H    code=0x9c
func RES_9C(cpu *CPU) {
	panic("TODO RES_9C")
}

// SET 4,C    code=0xe1
func SET_E1(cpu *CPU) {
	panic("TODO SET_E1")
}

// RES 1,(HL)    code=0x8e
func RES_8E(cpu *CPU) {
	panic("TODO RES_8E")
}

// RES 4,D    code=0xa2
func RES_A2(cpu *CPU) {
	panic("TODO RES_A2")
}

// SET 6,H    code=0xf4
func SET_F4(cpu *CPU) {
	panic("TODO SET_F4")
}

// SET 3,D    code=0xda
func SET_DA(cpu *CPU) {
	panic("TODO SET_DA")
}

// RL E    code=0x13
func RL_13(cpu *CPU) {
	res, flags := rotate(cpu.E, 0, cpu.F, false)
	cpu.E = res
	cpu.F = flags

	cpu.cycles += 8
}

// SLA (HL)    code=0x26
func SLA_26(cpu *CPU) {
	panic("TODO SLA_26")
}

// SRA E    code=0x2b
func SRA_2B(cpu *CPU) {
	panic("TODO SRA_2B")
}

// BIT 1,H    code=0x4c
func BIT_4C(cpu *CPU) {
	value := cpu.H
	var flags Flags
	if (value & (1 << 1)) == 0 {
		flags |= FLAGZ
	}
	flags |= FLAGH
	cpu.F = FlagRegister(flags)

	cpu.cycles += 8
}

// RES 4,H    code=0xa4
func RES_A4(cpu *CPU) {
	panic("TODO RES_A4")
}

// SET 0,C    code=0xc1
func SET_C1(cpu *CPU) {
	panic("TODO SET_C1")
}

// SET 5,B    code=0xe8
func SET_E8(cpu *CPU) {
	panic("TODO SET_E8")
}

// RRC E    code=0x0b
func RRC_0B(cpu *CPU) {
	res, flags := rotate(cpu.E, 1, cpu.F, true)
	cpu.E = res
	cpu.F = flags

	cpu.cycles += 8
}

// BIT 1,D    code=0x4a
func BIT_4A(cpu *CPU) {
	value := cpu.D
	var flags Flags
	if (value & (1 << 1)) == 0 {
		flags |= FLAGZ
	}
	flags |= FLAGH
	cpu.F = FlagRegister(flags)

	cpu.cycles += 8
}

// RES 0,(HL)    code=0x86
func RES_86(cpu *CPU) {
	panic("TODO RES_86")
}

// SET 5,E    code=0xeb
func SET_EB(cpu *CPU) {
	panic("TODO SET_EB")
}

// RRC (HL)    code=0x0e
func RRC_0E(cpu *CPU) {
	res, flags := rotate(cpu.loadU8(cpu.HL()), 1, cpu.F, true)
	cpu.WriteMemory(cpu.HL(), res)
	cpu.F = flags

	cpu.cycles += 16
}

// SLA D    code=0x22
func SLA_22(cpu *CPU) {
	panic("TODO SLA_22")
}

// SRL B    code=0x38
func SRL_38(cpu *CPU) {
	panic("TODO SRL_38")
}

// SET 1,L    code=0xcd
func SET_CD(cpu *CPU) {
	panic("TODO SET_CD")
}

// BIT 1,B    code=0x48
func BIT_48(cpu *CPU) {
	value := cpu.B
	var flags Flags
	if (value & (1 << 1)) == 0 {
		flags |= FLAGZ
	}
	flags |= FLAGH
	cpu.F = FlagRegister(flags)

	cpu.cycles += 8
}

// BIT 6,(HL)    code=0x76
func BIT_76(cpu *CPU) {
	value := cpu.loadU8(cpu.HL())
	var flags Flags
	if (value & (1 << 6)) == 0 {
		flags |= FLAGZ
	}
	flags |= FLAGH
	cpu.F = FlagRegister(flags)

	cpu.cycles += 12
}

// RES 1,H    code=0x8c
func RES_8C(cpu *CPU) {
	panic("TODO RES_8C")
}

// RES 2,H    code=0x94
func RES_94(cpu *CPU) {
	panic("TODO RES_94")
}

// SET 0,H    code=0xc4
func SET_C4(cpu *CPU) {
	panic("TODO SET_C4")
}

// SWAP D    code=0x32
func SWAP_32(cpu *CPU) {
	panic("TODO SWAP_32")
}

// RES 6,C    code=0xb1
func RES_B1(cpu *CPU) {
	panic("TODO RES_B1")
}

// RES 7,E    code=0xbb
func RES_BB(cpu *CPU) {
	panic("TODO RES_BB")
}

// SET 0,D    code=0xc2
func SET_C2(cpu *CPU) {
	panic("TODO SET_C2")
}

// SET 2,L    code=0xd5
func SET_D5(cpu *CPU) {
	panic("TODO SET_D5")
}

// RRC C    code=0x09
func RRC_09(cpu *CPU) {
	res, flags := rotate(cpu.C, 1, cpu.F, true)
	cpu.C = res
	cpu.F = flags

	cpu.cycles += 8
}

// RL B    code=0x10
func RL_10(cpu *CPU) {
	res, flags := rotate(cpu.B, 0, cpu.F, false)
	cpu.B = res
	cpu.F = flags

	cpu.cycles += 8
}

// SRL D    code=0x3a
func SRL_3A(cpu *CPU) {
	panic("TODO SRL_3A")
}

// BIT 2,(HL)    code=0x56
func BIT_56(cpu *CPU) {
	value := cpu.loadU8(cpu.HL())
	var flags Flags
	if (value & (1 << 2)) == 0 {
		flags |= FLAGZ
	}
	flags |= FLAGH
	cpu.F = FlagRegister(flags)

	cpu.cycles += 12
}

// BIT 3,(HL)    code=0x5e
func BIT_5E(cpu *CPU) {
	value := cpu.loadU8(cpu.HL())
	var flags Flags
	if (value & (1 << 3)) == 0 {
		flags |= FLAGZ
	}
	flags |= FLAGH
	cpu.F = FlagRegister(flags)

	cpu.cycles += 12
}

// RES 6,A    code=0xb7
func RES_B7(cpu *CPU) {
	panic("TODO RES_B7")
}

// SET 3,E    code=0xdb
func SET_DB(cpu *CPU) {
	panic("TODO SET_DB")
}

// BIT 1,L    code=0x4d
func BIT_4D(cpu *CPU) {
	value := cpu.L
	var flags Flags
	if (value & (1 << 1)) == 0 {
		flags |= FLAGZ
	}
	flags |= FLAGH
	cpu.F = FlagRegister(flags)

	cpu.cycles += 8
}

// BIT 5,A    code=0x6f
func BIT_6F(cpu *CPU) {
	value := cpu.A
	var flags Flags
	if (value & (1 << 5)) == 0 {
		flags |= FLAGZ
	}
	flags |= FLAGH
	cpu.F = FlagRegister(flags)

	cpu.cycles += 8
}

// BIT 7,C    code=0x79
func BIT_79(cpu *CPU) {
	value := cpu.C
	var flags Flags
	if (value & (1 << 7)) == 0 {
		flags |= FLAGZ
	}
	flags |= FLAGH
	cpu.F = FlagRegister(flags)

	cpu.cycles += 8
}

// RES 0,L    code=0x85
func RES_85(cpu *CPU) {
	panic("TODO RES_85")
}

// RES 6,B    code=0xb0
func RES_B0(cpu *CPU) {
	panic("TODO RES_B0")
}

// SET 1,C    code=0xc9
func SET_C9(cpu *CPU) {
	panic("TODO SET_C9")
}

// SET 2,C    code=0xd1
func SET_D1(cpu *CPU) {
	panic("TODO SET_D1")
}

// RES 1,E    code=0x8b
func RES_8B(cpu *CPU) {
	panic("TODO RES_8B")
}

// RES 5,A    code=0xaf
func RES_AF(cpu *CPU) {
	panic("TODO RES_AF")
}

// SET 6,A    code=0xf7
func SET_F7(cpu *CPU) {
	panic("TODO SET_F7")
}

// RR B    code=0x18
func RR_18(cpu *CPU) {
	res, flags := rotate(cpu.B, 1, cpu.F, false)
	cpu.B = res
	cpu.F = flags

	cpu.cycles += 8
}

// BIT 4,D    code=0x62
func BIT_62(cpu *CPU) {
	value := cpu.D
	var flags Flags
	if (value & (1 << 4)) == 0 {
		flags |= FLAGZ
	}
	flags |= FLAGH
	cpu.F = FlagRegister(flags)

	cpu.cycles += 8
}

// RES 1,A    code=0x8f
func RES_8F(cpu *CPU) {
	panic("TODO RES_8F")
}

// RES 7,H    code=0xbc
func RES_BC(cpu *CPU) {
	panic("TODO RES_BC")
}

// SET 5,C    code=0xe9
func SET_E9(cpu *CPU) {
	panic("TODO SET_E9")
}

// SRA H    code=0x2c
func SRA_2C(cpu *CPU) {
	panic("TODO SRA_2C")
}

// BIT 0,A    code=0x47
func BIT_47(cpu *CPU) {
	value := cpu.A
	var flags Flags
	if (value & (1 << 0)) == 0 {
		flags |= FLAGZ
	}
	flags |= FLAGH
	cpu.F = FlagRegister(flags)

	cpu.cycles += 8
}

// BIT 6,A    code=0x77
func BIT_77(cpu *CPU) {
	value := cpu.A
	var flags Flags
	if (value & (1 << 6)) == 0 {
		flags |= FLAGZ
	}
	flags |= FLAGH
	cpu.F = FlagRegister(flags)

	cpu.cycles += 8
}

// RES 5,E    code=0xab
func RES_AB(cpu *CPU) {
	panic("TODO RES_AB")
}

// SET 7,C    code=0xf9
func SET_F9(cpu *CPU) {
	panic("TODO SET_F9")
}

// RR (HL)    code=0x1e
func RR_1E(cpu *CPU) {
	res, flags := rotate(cpu.loadU8(cpu.HL()), 1, cpu.F, false)
	cpu.WriteMemory(cpu.HL(), res)
	cpu.F = flags

	cpu.cycles += 16
}

// RL C    code=0x11
func RL_11(cpu *CPU) {
	res, flags := rotate(cpu.C, 0, cpu.F, false)
	cpu.C = res
	cpu.F = flags

	cpu.cycles += 8
}

// SRA A    code=0x2f
func SRA_2F(cpu *CPU) {
	panic("TODO SRA_2F")
}

// BIT 2,C    code=0x51
func BIT_51(cpu *CPU) {
	value := cpu.C
	var flags Flags
	if (value & (1 << 2)) == 0 {
		flags |= FLAGZ
	}
	flags |= FLAGH
	cpu.F = FlagRegister(flags)

	cpu.cycles += 8
}

// BIT 7,(HL)    code=0x7e
func BIT_7E(cpu *CPU) {
	value := cpu.loadU8(cpu.HL())
	var flags Flags
	if (value & (1 << 7)) == 0 {
		flags |= FLAGZ
	}
	flags |= FLAGH
	cpu.F = FlagRegister(flags)

	cpu.cycles += 12
}

// RES 1,B    code=0x88
func RES_88(cpu *CPU) {
	panic("TODO RES_88")
}

// RES 2,L    code=0x95
func RES_95(cpu *CPU) {
	panic("TODO RES_95")
}

// RES 3,(HL)    code=0x9e
func RES_9E(cpu *CPU) {
	panic("TODO RES_9E")
}

// RR C    code=0x19
func RR_19(cpu *CPU) {
	res, flags := rotate(cpu.C, 1, cpu.F, false)
	cpu.C = res
	cpu.F = flags

	cpu.cycles += 8
}

// RES 3,D    code=0x9a
func RES_9A(cpu *CPU) {
	panic("TODO RES_9A")
}

// RLC (HL)    code=0x06
func RLC_06(cpu *CPU) {
	res, flags := rotate(cpu.loadU8(cpu.HL()), 0, cpu.F, true)
	cpu.WriteMemory(cpu.HL(), res)
	cpu.F = flags

	cpu.cycles += 16
}

// SWAP C    code=0x31
func SWAP_31(cpu *CPU) {
	panic("TODO SWAP_31")
}

// BIT 1,C    code=0x49
func BIT_49(cpu *CPU) {
	value := cpu.C
	var flags Flags
	if (value & (1 << 1)) == 0 {
		flags |= FLAGZ
	}
	flags |= FLAGH
	cpu.F = FlagRegister(flags)

	cpu.cycles += 8
}

// SET 1,H    code=0xcc
func SET_CC(cpu *CPU) {
	panic("TODO SET_CC")
}

// SWAP E    code=0x33
func SWAP_33(cpu *CPU) {
	panic("TODO SWAP_33")
}

// BIT 3,B    code=0x58
func BIT_58(cpu *CPU) {
	value := cpu.B
	var flags Flags
	if (value & (1 << 3)) == 0 {
		flags |= FLAGZ
	}
	flags |= FLAGH
	cpu.F = FlagRegister(flags)

	cpu.cycles += 8
}

// RES 2,(HL)    code=0x96
func RES_96(cpu *CPU) {
	panic("TODO RES_96")
}

// RES 4,B    code=0xa0
func RES_A0(cpu *CPU) {
	panic("TODO RES_A0")
}

// RES 4,C    code=0xa1
func RES_A1(cpu *CPU) {
	panic("TODO RES_A1")
}

// SET 0,A    code=0xc7
func SET_C7(cpu *CPU) {
	panic("TODO SET_C7")
}

// SET 2,(HL)    code=0xd6
func SET_D6(cpu *CPU) {
	panic("TODO SET_D6")
}

// RLC H    code=0x04
func RLC_04(cpu *CPU) {
	res, flags := rotate(cpu.H, 0, cpu.F, true)
	cpu.H = res
	cpu.F = flags

	cpu.cycles += 8
}

// SWAP A    code=0x37
func SWAP_37(cpu *CPU) {
	panic("TODO SWAP_37")
}

// BIT 0,E    code=0x43
func BIT_43(cpu *CPU) {
	value := cpu.E
	var flags Flags
	if (value & (1 << 0)) == 0 {
		flags |= FLAGZ
	}
	flags |= FLAGH
	cpu.F = FlagRegister(flags)

	cpu.cycles += 8
}

// SRA B    code=0x28
func SRA_28(cpu *CPU) {
	panic("TODO SRA_28")
}

// RES 2,B    code=0x90
func RES_90(cpu *CPU) {
	panic("TODO RES_90")
}

// SET 1,(HL)    code=0xce
func SET_CE(cpu *CPU) {
	panic("TODO SET_CE")
}

// RES 0,C    code=0x81
func RES_81(cpu *CPU) {
	panic("TODO RES_81")
}

// RES 4,A    code=0xa7
func RES_A7(cpu *CPU) {
	panic("TODO RES_A7")
}

// SET 2,H    code=0xd4
func SET_D4(cpu *CPU) {
	panic("TODO SET_D4")
}

// SET 4,B    code=0xe0
func SET_E0(cpu *CPU) {
	panic("TODO SET_E0")
}

// RLC E    code=0x03
func RLC_03(cpu *CPU) {
	res, flags := rotate(cpu.E, 0, cpu.F, true)
	cpu.E = res
	cpu.F = flags

	cpu.cycles += 8
}

// RRC D    code=0x0a
func RRC_0A(cpu *CPU) {
	res, flags := rotate(cpu.D, 1, cpu.F, true)
	cpu.D = res
	cpu.F = flags

	cpu.cycles += 8
}

// SWAP (HL)    code=0x36
func SWAP_36(cpu *CPU) {
	panic("TODO SWAP_36")
}

// RES 4,L    code=0xa5
func RES_A5(cpu *CPU) {
	panic("TODO RES_A5")
}

// RLC B    code=0x00
func RLC_00(cpu *CPU) {
	res, flags := rotate(cpu.B, 0, cpu.F, true)
	cpu.B = res
	cpu.F = flags

	cpu.cycles += 8
}

// RL D    code=0x12
func RL_12(cpu *CPU) {
	res, flags := rotate(cpu.D, 0, cpu.F, false)
	cpu.D = res
	cpu.F = flags

	cpu.cycles += 8
}

// SWAP B    code=0x30
func SWAP_30(cpu *CPU) {
	panic("TODO SWAP_30")
}

// RES 1,L    code=0x8d
func RES_8D(cpu *CPU) {
	panic("TODO RES_8D")
}

// RES 3,B    code=0x98
func RES_98(cpu *CPU) {
	panic("TODO RES_98")
}

// RL H    code=0x14
func RL_14(cpu *CPU) {
	res, flags := rotate(cpu.H, 0, cpu.F, false)
	cpu.H = res
	cpu.F = flags

	cpu.cycles += 8
}

// RR L    code=0x1d
func RR_1D(cpu *CPU) {
	res, flags := rotate(cpu.L, 1, cpu.F, false)
	cpu.L = res
	cpu.F = flags

	cpu.cycles += 8
}

// SLA L    code=0x25
func SLA_25(cpu *CPU) {
	panic("TODO SLA_25")
}

// RES 3,A    code=0x9f
func RES_9F(cpu *CPU) {
	panic("TODO RES_9F")
}

// SET 1,D    code=0xca
func SET_CA(cpu *CPU) {
	panic("TODO SET_CA")
}

// SWAP L    code=0x35
func SWAP_35(cpu *CPU) {
	panic("TODO SWAP_35")
}

// BIT 4,(HL)    code=0x66
func BIT_66(cpu *CPU) {
	value := cpu.loadU8(cpu.HL())
	var flags Flags
	if (value & (1 << 4)) == 0 {
		flags |= FLAGZ
	}
	flags |= FLAGH
	cpu.F = FlagRegister(flags)

	cpu.cycles += 12
}

// BIT 5,L    code=0x6d
func BIT_6D(cpu *CPU) {
	value := cpu.L
	var flags Flags
	if (value & (1 << 5)) == 0 {
		flags |= FLAGZ
	}
	flags |= FLAGH
	cpu.F = FlagRegister(flags)

	cpu.cycles += 8
}

// RES 2,D    code=0x92
func RES_92(cpu *CPU) {
	panic("TODO RES_92")
}

// RES 5,B    code=0xa8
func RES_A8(cpu *CPU) {
	panic("TODO RES_A8")
}

// SET 2,B    code=0xd0
func SET_D0(cpu *CPU) {
	panic("TODO SET_D0")
}

// RR A    code=0x1f
func RR_1F(cpu *CPU) {
	res, flags := rotate(cpu.A, 1, cpu.F, false)
	cpu.A = res
	cpu.F = flags

	cpu.cycles += 8
}

// BIT 3,L    code=0x5d
func BIT_5D(cpu *CPU) {
	value := cpu.L
	var flags Flags
	if (value & (1 << 3)) == 0 {
		flags |= FLAGZ
	}
	flags |= FLAGH
	cpu.F = FlagRegister(flags)

	cpu.cycles += 8
}

// RES 0,B    code=0x80
func RES_80(cpu *CPU) {
	panic("TODO RES_80")
}

// RES 0,A    code=0x87
func RES_87(cpu *CPU) {
	panic("TODO RES_87")
}

// RES 2,A    code=0x97
func RES_97(cpu *CPU) {
	panic("TODO RES_97")
}

// SET 2,E    code=0xd3
func SET_D3(cpu *CPU) {
	panic("TODO SET_D3")
}

// SET 4,L    code=0xe5
func SET_E5(cpu *CPU) {
	panic("TODO SET_E5")
}

// SET 6,B    code=0xf0
func SET_F0(cpu *CPU) {
	panic("TODO SET_F0")
}

// BIT 0,B    code=0x40
func BIT_40(cpu *CPU) {
	value := cpu.B
	var flags Flags
	if (value & (1 << 0)) == 0 {
		flags |= FLAGZ
	}
	flags |= FLAGH
	cpu.F = FlagRegister(flags)

	cpu.cycles += 8
}

// RES 0,D    code=0x82
func RES_82(cpu *CPU) {
	panic("TODO RES_82")
}

// SET 0,B    code=0xc0
func SET_C0(cpu *CPU) {
	panic("TODO SET_C0")
}

// RR E    code=0x1b
func RR_1B(cpu *CPU) {
	res, flags := rotate(cpu.E, 1, cpu.F, false)
	cpu.E = res
	cpu.F = flags

	cpu.cycles += 8
}

// BIT 7,D    code=0x7a
func BIT_7A(cpu *CPU) {
	value := cpu.D
	var flags Flags
	if (value & (1 << 7)) == 0 {
		flags |= FLAGZ
	}
	flags |= FLAGH
	cpu.F = FlagRegister(flags)

	cpu.cycles += 8
}

// RES 7,C    code=0xb9
func RES_B9(cpu *CPU) {
	panic("TODO RES_B9")
}

// SET 3,C    code=0xd9
func SET_D9(cpu *CPU) {
	panic("TODO SET_D9")
}

// BIT 2,B    code=0x50
func BIT_50(cpu *CPU) {
	value := cpu.B
	var flags Flags
	if (value & (1 << 2)) == 0 {
		flags |= FLAGZ
	}
	flags |= FLAGH
	cpu.F = FlagRegister(flags)

	cpu.cycles += 8
}

// RES 6,D    code=0xb2
func RES_B2(cpu *CPU) {
	panic("TODO RES_B2")
}

// SLA B    code=0x20
func SLA_20(cpu *CPU) {
	panic("TODO SLA_20")
}

// SLA A    code=0x27
func SLA_27(cpu *CPU) {
	panic("TODO SLA_27")
}

// SWAP H    code=0x34
func SWAP_34(cpu *CPU) {
	panic("TODO SWAP_34")
}

// BIT 3,C    code=0x59
func BIT_59(cpu *CPU) {
	value := cpu.C
	var flags Flags
	if (value & (1 << 3)) == 0 {
		flags |= FLAGZ
	}
	flags |= FLAGH
	cpu.F = FlagRegister(flags)

	cpu.cycles += 8
}

// BIT 7,L    code=0x7d
func BIT_7D(cpu *CPU) {
	value := cpu.L
	var flags Flags
	if (value & (1 << 7)) == 0 {
		flags |= FLAGZ
	}
	flags |= FLAGH
	cpu.F = FlagRegister(flags)

	cpu.cycles += 8
}

// BIT 5,E    code=0x6b
func BIT_6B(cpu *CPU) {
	value := cpu.E
	var flags Flags
	if (value & (1 << 5)) == 0 {
		flags |= FLAGZ
	}
	flags |= FLAGH
	cpu.F = FlagRegister(flags)

	cpu.cycles += 8
}

// BIT 2,H    code=0x54
func BIT_54(cpu *CPU) {
	value := cpu.H
	var flags Flags
	if (value & (1 << 2)) == 0 {
		flags |= FLAGZ
	}
	flags |= FLAGH
	cpu.F = FlagRegister(flags)

	cpu.cycles += 8
}

// RES 5,H    code=0xac
func RES_AC(cpu *CPU) {
	panic("TODO RES_AC")
}

// RES 7,(HL)    code=0xbe
func RES_BE(cpu *CPU) {
	panic("TODO RES_BE")
}

// RLC L    code=0x05
func RLC_05(cpu *CPU) {
	res, flags := rotate(cpu.L, 0, cpu.F, true)
	cpu.L = res
	cpu.F = flags

	cpu.cycles += 8
}

// RES 0,H    code=0x84
func RES_84(cpu *CPU) {
	panic("TODO RES_84")
}

// RES 7,L    code=0xbd
func RES_BD(cpu *CPU) {
	panic("TODO RES_BD")
}

// BIT 5,H    code=0x6c
func BIT_6C(cpu *CPU) {
	value := cpu.H
	var flags Flags
	if (value & (1 << 5)) == 0 {
		flags |= FLAGZ
	}
	flags |= FLAGH
	cpu.F = FlagRegister(flags)

	cpu.cycles += 8
}

// SRA L    code=0x2d
func SRA_2D(cpu *CPU) {
	panic("TODO SRA_2D")
}

// SRL H    code=0x3c
func SRL_3C(cpu *CPU) {
	panic("TODO SRL_3C")
}

// BIT 6,B    code=0x70
func BIT_70(cpu *CPU) {
	value := cpu.B
	var flags Flags
	if (value & (1 << 6)) == 0 {
		flags |= FLAGZ
	}
	flags |= FLAGH
	cpu.F = FlagRegister(flags)

	cpu.cycles += 8
}

// RES 7,D    code=0xba
func RES_BA(cpu *CPU) {
	panic("TODO RES_BA")
}

// BIT 3,D    code=0x5a
func BIT_5A(cpu *CPU) {
	value := cpu.D
	var flags Flags
	if (value & (1 << 3)) == 0 {
		flags |= FLAGZ
	}
	flags |= FLAGH
	cpu.F = FlagRegister(flags)

	cpu.cycles += 8
}

// RES 1,D    code=0x8a
func RES_8A(cpu *CPU) {
	panic("TODO RES_8A")
}

// SET 6,D    code=0xf2
func SET_F2(cpu *CPU) {
	panic("TODO SET_F2")
}

// SLA E    code=0x23
func SLA_23(cpu *CPU) {
	panic("TODO SLA_23")
}

// SRL E    code=0x3b
func SRL_3B(cpu *CPU) {
	panic("TODO SRL_3B")
}

// RES 3,E    code=0x9b
func RES_9B(cpu *CPU) {
	panic("TODO RES_9B")
}

// RES 4,E    code=0xa3
func RES_A3(cpu *CPU) {
	panic("TODO RES_A3")
}

// RES 4,(HL)    code=0xa6
func RES_A6(cpu *CPU) {
	panic("TODO RES_A6")
}

// SET 0,L    code=0xc5
func SET_C5(cpu *CPU) {
	panic("TODO SET_C5")
}

// SET 1,A    code=0xcf
func SET_CF(cpu *CPU) {
	panic("TODO SET_CF")
}

// SET 4,(HL)    code=0xe6
func SET_E6(cpu *CPU) {
	panic("TODO SET_E6")
}

// RL L    code=0x15
func RL_15(cpu *CPU) {
	res, flags := rotate(cpu.L, 0, cpu.F, false)
	cpu.L = res
	cpu.F = flags

	cpu.cycles += 8
}

// SRL (HL)    code=0x3e
func SRL_3E(cpu *CPU) {
	panic("TODO SRL_3E")
}

// RES 6,E    code=0xb3
func RES_B3(cpu *CPU) {
	panic("TODO RES_B3")
}

// RES 6,(HL)    code=0xb6
func RES_B6(cpu *CPU) {
	panic("TODO RES_B6")
}

// RES 7,A    code=0xbf
func RES_BF(cpu *CPU) {
	panic("TODO RES_BF")
}

// BIT 7,H    code=0x7c
func BIT_7C(cpu *CPU) {
	value := cpu.H
	var flags Flags
	if (value & (1 << 7)) == 0 {
		flags |= FLAGZ
	}
	flags |= FLAGH
	cpu.F = FlagRegister(flags)

	cpu.cycles += 8
}

// RES 2,C    code=0x91
func RES_91(cpu *CPU) {
	panic("TODO RES_91")
}

// RL (HL)    code=0x16
func RL_16(cpu *CPU) {
	res, flags := rotate(cpu.loadU8(cpu.HL()), 0, cpu.F, false)
	cpu.WriteMemory(cpu.HL(), res)
	cpu.F = flags

	cpu.cycles += 16
}

// RES 3,L    code=0x9d
func RES_9D(cpu *CPU) {
	panic("TODO RES_9D")
}

// RES 5,C    code=0xa9
func RES_A9(cpu *CPU) {
	panic("TODO RES_A9")
}

// SET 4,A    code=0xe7
func SET_E7(cpu *CPU) {
	panic("TODO SET_E7")
}

// SRL A    code=0x3f
func SRL_3F(cpu *CPU) {
	panic("TODO SRL_3F")
}

// SET 3,H    code=0xdc
func SET_DC(cpu *CPU) {
	panic("TODO SET_DC")
}

// SET 7,A    code=0xff
func SET_FF(cpu *CPU) {
	panic("TODO SET_FF")
}

// RRC A    code=0x0f
func RRC_0F(cpu *CPU) {
	res, flags := rotate(cpu.A, 1, cpu.F, true)
	cpu.A = res
	cpu.F = flags

	cpu.cycles += 8
}

// BIT 0,H    code=0x44
func BIT_44(cpu *CPU) {
	value := cpu.H
	var flags Flags
	if (value & (1 << 0)) == 0 {
		flags |= FLAGZ
	}
	flags |= FLAGH
	cpu.F = FlagRegister(flags)

	cpu.cycles += 8
}

// BIT 4,A    code=0x67
func BIT_67(cpu *CPU) {
	value := cpu.A
	var flags Flags
	if (value & (1 << 4)) == 0 {
		flags |= FLAGZ
	}
	flags |= FLAGH
	cpu.F = FlagRegister(flags)

	cpu.cycles += 8
}

// BIT 7,A    code=0x7f
func BIT_7F(cpu *CPU) {
	value := cpu.A
	var flags Flags
	if (value & (1 << 7)) == 0 {
		flags |= FLAGZ
	}
	flags |= FLAGH
	cpu.F = FlagRegister(flags)

	cpu.cycles += 8
}

// SET 7,(HL)    code=0xfe
func SET_FE(cpu *CPU) {
	panic("TODO SET_FE")
}

// BIT 6,D    code=0x72
func BIT_72(cpu *CPU) {
	value := cpu.D
	var flags Flags
	if (value & (1 << 6)) == 0 {
		flags |= FLAGZ
	}
	flags |= FLAGH
	cpu.F = FlagRegister(flags)

	cpu.cycles += 8
}

// RES 6,H    code=0xb4
func RES_B4(cpu *CPU) {
	panic("TODO RES_B4")
}

// SET 5,D    code=0xea
func SET_EA(cpu *CPU) {
	panic("TODO SET_EA")
}

// BIT 5,(HL)    code=0x6e
func BIT_6E(cpu *CPU) {
	value := cpu.loadU8(cpu.HL())
	var flags Flags
	if (value & (1 << 5)) == 0 {
		flags |= FLAGZ
	}
	flags |= FLAGH
	cpu.F = FlagRegister(flags)

	cpu.cycles += 12
}

// RES 3,C    code=0x99
func RES_99(cpu *CPU) {
	panic("TODO RES_99")
}

// SET 3,(HL)    code=0xde
func SET_DE(cpu *CPU) {
	panic("TODO SET_DE")
}

// RL A    code=0x17
func RL_17(cpu *CPU) {
	res, flags := rotate(cpu.A, 0, cpu.F, false)
	cpu.A = res
	cpu.F = flags

	cpu.cycles += 8
}

// BIT 5,D    code=0x6a
func BIT_6A(cpu *CPU) {
	value := cpu.D
	var flags Flags
	if (value & (1 << 5)) == 0 {
		flags |= FLAGZ
	}
	flags |= FLAGH
	cpu.F = FlagRegister(flags)

	cpu.cycles += 8
}

// SET 2,D    code=0xd2
func SET_D2(cpu *CPU) {
	panic("TODO SET_D2")
}

// SET 3,L    code=0xdd
func SET_DD(cpu *CPU) {
	panic("TODO SET_DD")
}

// RRC B    code=0x08
func RRC_08(cpu *CPU) {
	res, flags := rotate(cpu.B, 1, cpu.F, true)
	cpu.B = res
	cpu.F = flags

	cpu.cycles += 8
}

// BIT 6,E    code=0x73
func BIT_73(cpu *CPU) {
	value := cpu.E
	var flags Flags
	if (value & (1 << 6)) == 0 {
		flags |= FLAGZ
	}
	flags |= FLAGH
	cpu.F = FlagRegister(flags)

	cpu.cycles += 8
}

// BIT 6,L    code=0x75
func BIT_75(cpu *CPU) {
	value := cpu.L
	var flags Flags
	if (value & (1 << 6)) == 0 {
		flags |= FLAGZ
	}
	flags |= FLAGH
	cpu.F = FlagRegister(flags)

	cpu.cycles += 8
}

// SET 1,B    code=0xc8
func SET_C8(cpu *CPU) {
	panic("TODO SET_C8")
}

// SET 5,L    code=0xed
func SET_ED(cpu *CPU) {
	panic("TODO SET_ED")
}

// BIT 5,C    code=0x69
func BIT_69(cpu *CPU) {
	value := cpu.C
	var flags Flags
	if (value & (1 << 5)) == 0 {
		flags |= FLAGZ
	}
	flags |= FLAGH
	cpu.F = FlagRegister(flags)

	cpu.cycles += 8
}

// RLC D    code=0x02
func RLC_02(cpu *CPU) {
	res, flags := rotate(cpu.D, 0, cpu.F, true)
	cpu.D = res
	cpu.F = flags

	cpu.cycles += 8
}

// RR D    code=0x1a
func RR_1A(cpu *CPU) {
	res, flags := rotate(cpu.D, 1, cpu.F, false)
	cpu.D = res
	cpu.F = flags

	cpu.cycles += 8
}

// SRA C    code=0x29
func SRA_29(cpu *CPU) {
	panic("TODO SRA_29")
}

// BIT 0,D    code=0x42
func BIT_42(cpu *CPU) {
	value := cpu.D
	var flags Flags
	if (value & (1 << 0)) == 0 {
		flags |= FLAGZ
	}
	flags |= FLAGH
	cpu.F = FlagRegister(flags)

	cpu.cycles += 8
}

// BIT 7,E    code=0x7b
func BIT_7B(cpu *CPU) {
	value := cpu.E
	var flags Flags
	if (value & (1 << 7)) == 0 {
		flags |= FLAGZ
	}
	flags |= FLAGH
	cpu.F = FlagRegister(flags)

	cpu.cycles += 8
}

// RES 2,E    code=0x93
func RES_93(cpu *CPU) {
	panic("TODO RES_93")
}

// SET 6,(HL)    code=0xf6
func SET_F6(cpu *CPU) {
	panic("TODO SET_F6")
}

// RRC H    code=0x0c
func RRC_0C(cpu *CPU) {
	res, flags := rotate(cpu.H, 1, cpu.F, true)
	cpu.H = res
	cpu.F = flags

	cpu.cycles += 8
}

// BIT 2,E    code=0x53
func BIT_53(cpu *CPU) {
	value := cpu.E
	var flags Flags
	if (value & (1 << 2)) == 0 {
		flags |= FLAGZ
	}
	flags |= FLAGH
	cpu.F = FlagRegister(flags)

	cpu.cycles += 8
}

// BIT 4,C    code=0x61
func BIT_61(cpu *CPU) {
	value := cpu.C
	var flags Flags
	if (value & (1 << 4)) == 0 {
		flags |= FLAGZ
	}
	flags |= FLAGH
	cpu.F = FlagRegister(flags)

	cpu.cycles += 8
}

// BIT 4,H    code=0x64
func BIT_64(cpu *CPU) {
	value := cpu.H
	var flags Flags
	if (value & (1 << 4)) == 0 {
		flags |= FLAGZ
	}
	flags |= FLAGH
	cpu.F = FlagRegister(flags)

	cpu.cycles += 8
}

// RES 5,(HL)    code=0xae
func RES_AE(cpu *CPU) {
	panic("TODO RES_AE")
}

// SET 5,(HL)    code=0xee
func SET_EE(cpu *CPU) {
	panic("TODO SET_EE")
}

// SET 6,E    code=0xf3
func SET_F3(cpu *CPU) {
	panic("TODO SET_F3")
}

// SRL C    code=0x39
func SRL_39(cpu *CPU) {
	panic("TODO SRL_39")
}

// RES 5,D    code=0xaa
func RES_AA(cpu *CPU) {
	panic("TODO RES_AA")
}

// SET 4,H    code=0xe4
func SET_E4(cpu *CPU) {
	panic("TODO SET_E4")
}

// BIT 0,(HL)    code=0x46
func BIT_46(cpu *CPU) {
	value := cpu.loadU8(cpu.HL())
	var flags Flags
	if (value & (1 << 0)) == 0 {
		flags |= FLAGZ
	}
	flags |= FLAGH
	cpu.F = FlagRegister(flags)

	cpu.cycles += 12
}

// SET 7,B    code=0xf8
func SET_F8(cpu *CPU) {
	panic("TODO SET_F8")
}

// RRC L    code=0x0d
func RRC_0D(cpu *CPU) {
	res, flags := rotate(cpu.L, 1, cpu.F, true)
	cpu.L = res
	cpu.F = flags

	cpu.cycles += 8
}

// BIT 3,A    code=0x5f
func BIT_5F(cpu *CPU) {
	value := cpu.A
	var flags Flags
	if (value & (1 << 3)) == 0 {
		flags |= FLAGZ
	}
	flags |= FLAGH
	cpu.F = FlagRegister(flags)

	cpu.cycles += 8
}

// BIT 4,E    code=0x63
func BIT_63(cpu *CPU) {
	value := cpu.E
	var flags Flags
	if (value & (1 << 4)) == 0 {
		flags |= FLAGZ
	}
	flags |= FLAGH
	cpu.F = FlagRegister(flags)

	cpu.cycles += 8
}

// SET 0,(HL)    code=0xc6
func SET_C6(cpu *CPU) {
	panic("TODO SET_C6")
}

// RLC C    code=0x01
func RLC_01(cpu *CPU) {
	res, flags := rotate(cpu.C, 0, cpu.F, true)
	cpu.C = res
	cpu.F = flags

	cpu.cycles += 8
}

// SRA (HL)    code=0x2e
func SRA_2E(cpu *CPU) {
	panic("TODO SRA_2E")
}

// BIT 0,L    code=0x45
func BIT_45(cpu *CPU) {
	value := cpu.L
	var flags Flags
	if (value & (1 << 0)) == 0 {
		flags |= FLAGZ
	}
	flags |= FLAGH
	cpu.F = FlagRegister(flags)

	cpu.cycles += 8
}

// BIT 2,L    code=0x55
func BIT_55(cpu *CPU) {
	value := cpu.L
	var flags Flags
	if (value & (1 << 2)) == 0 {
		flags |= FLAGZ
	}
	flags |= FLAGH
	cpu.F = FlagRegister(flags)

	cpu.cycles += 8
}

// BIT 1,(HL)    code=0x4e
func BIT_4E(cpu *CPU) {
	value := cpu.loadU8(cpu.HL())
	var flags Flags
	if (value & (1 << 1)) == 0 {
		flags |= FLAGZ
	}
	flags |= FLAGH
	cpu.F = FlagRegister(flags)

	cpu.cycles += 12
}

// BIT 4,L    code=0x65
func BIT_65(cpu *CPU) {
	value := cpu.L
	var flags Flags
	if (value & (1 << 4)) == 0 {
		flags |= FLAGZ
	}
	flags |= FLAGH
	cpu.F = FlagRegister(flags)

	cpu.cycles += 8
}

// SET 1,E    code=0xcb
func SET_CB(cpu *CPU) {
	panic("TODO SET_CB")
}

// SET 3,A    code=0xdf
func SET_DF(cpu *CPU) {
	panic("TODO SET_DF")
}

// SET 4,D    code=0xe2
func SET_E2(cpu *CPU) {
	panic("TODO SET_E2")
}

// SET 4,E    code=0xe3
func SET_E3(cpu *CPU) {
	panic("TODO SET_E3")
}

// RLC A    code=0x07
func RLC_07(cpu *CPU) {
	res, flags := rotate(cpu.A, 0, cpu.F, true)
	cpu.A = res
	cpu.F = flags

	cpu.cycles += 8
}

// SRA D    code=0x2a
func SRA_2A(cpu *CPU) {
	panic("TODO SRA_2A")
}

// BIT 5,B    code=0x68
func BIT_68(cpu *CPU) {
	value := cpu.B
	var flags Flags
	if (value & (1 << 5)) == 0 {
		flags |= FLAGZ
	}
	flags |= FLAGH
	cpu.F = FlagRegister(flags)

	cpu.cycles += 8
}

// RES 5,L    code=0xad
func RES_AD(cpu *CPU) {
	panic("TODO RES_AD")
}

// SET 3,B    code=0xd8
func SET_D8(cpu *CPU) {
	panic("TODO SET_D8")
}

// SET 7,D    code=0xfa
func SET_FA(cpu *CPU) {
	panic("TODO SET_FA")
}

// SET 7,L    code=0xfd
func SET_FD(cpu *CPU) {
	panic("TODO SET_FD")
}

// SLA C    code=0x21
func SLA_21(cpu *CPU) {
	panic("TODO SLA_21")
}

// RES 7,B    code=0xb8
func RES_B8(cpu *CPU) {
	panic("TODO RES_B8")
}

// BIT 3,E    code=0x5b
func BIT_5B(cpu *CPU) {
	value := cpu.E
	var flags Flags
	if (value & (1 << 3)) == 0 {
		flags |= FLAGZ
	}
	flags |= FLAGH
	cpu.F = FlagRegister(flags)

	cpu.cycles += 8
}

// SET 2,A    code=0xd7
func SET_D7(cpu *CPU) {
	panic("TODO SET_D7")
}

// SET 6,C    code=0xf1
func SET_F1(cpu *CPU) {
	panic("TODO SET_F1")
}

// SRL L    code=0x3d
func SRL_3D(cpu *CPU) {
	panic("TODO SRL_3D")
}

// BIT 6,H    code=0x74
func BIT_74(cpu *CPU) {
	value := cpu.H
	var flags Flags
	if (value & (1 << 6)) == 0 {
		flags |= FLAGZ
	}
	flags |= FLAGH
	cpu.F = FlagRegister(flags)

	cpu.cycles += 8
}

// SET 5,A    code=0xef
func SET_EF(cpu *CPU) {
	panic("TODO SET_EF")
}

// SET 7,H    code=0xfc
func SET_FC(cpu *CPU) {
	panic("TODO SET_FC")
}

// BIT 1,A    code=0x4f
func BIT_4F(cpu *CPU) {
	value := cpu.A
	var flags Flags
	if (value & (1 << 1)) == 0 {
		flags |= FLAGZ
	}
	flags |= FLAGH
	cpu.F = FlagRegister(flags)

	cpu.cycles += 8
}

// RES 1,C    code=0x89
func RES_89(cpu *CPU) {
	panic("TODO RES_89")
}

var ops = map[uint8]Instruction{
	0x2b: DEC_2B,
	0x4b: LD_4B,
	0x57: LD_57,
	0x7a: LD_7A,
	0x9a: SBC_9A,
	0xa1: AND_A1,
	0xa3: AND_A3,
	0xd5: PUSH_D5,
	0x58: LD_58,
	0x75: LD_75,
	0xc0: RET_C0,
	0xce: ADC_CE,
	0xe8: ADD_E8,
	0xff: RST_FF,
	0xe:  LD_0E,
	0x22: LD_22,
	0x60: LD_60,
	0x6b: LD_6B,
	0x5:  DEC_05,
	0xa:  LD_0A,
	0x3d: DEC_3D,
	0x5b: LD_5B,
	0x5c: LD_5C,
	0xa9: XOR_A9,
	0xb8: CP_B8,
	0xda: JP_DA,
	0xe7: RST_E7,
	0x35: DEC_35,
	0x4a: LD_4A,
	0x69: LD_69,
	0x80: ADD_80,
	0xad: XOR_AD,
	0xe9: JP_E9,
	0x31: LD_31,
	0x25: DEC_25,
	0x4c: LD_4C,
	0x53: LD_53,
	0x68: LD_68,
	0x94: SUB_94,
	0xe4: ILLEGAL_E4_E4,
	0x49: LD_49,
	0x5e: LD_5E,
	0xba: CP_BA,
	0xd7: RST_D7,
	0x2:  LD_02,
	0x33: INC_33,
	0x54: LD_54,
	0x8f: ADC_8F,
	0xb5: OR_B5,
	0x4:  INC_04,
	0x7d: LD_7D,
	0xe5: PUSH_E5,
	0xe6: AND_E6,
	0xf4: ILLEGAL_F4_F4,
	0x32: LD_32,
	0x29: ADD_29,
	0x76: HALT_76,
	0x83: ADD_83,
	0xab: XOR_AB,
	0xbb: CP_BB,
	0xc6: ADD_C6,
	0x12: LD_12,
	0x8b: ADC_8B,
	0xef: RST_EF,
	0x2d: DEC_2D,
	0x46: LD_46,
	0x64: LD_64,
	0xa4: AND_A4,
	0xaf: XOR_AF,
	0xd4: CALL_D4,
	0x7:  RLCA_07,
	0x70: LD_70,
	0xa5: AND_A5,
	0xf2: LDH_F2,
	0x52: LD_52,
	0xc3: JP_C3,
	0x16: LD_16,
	0x17: RLA_17,
	0xa0: AND_A0,
	0xa2: AND_A2,
	0xfb: EI_FB,
	0xfc: ILLEGAL_FC_FC,
	0x87: ADD_87,
	0x8e: ADC_8E,
	0xcc: CALL_CC,
	0x26: LD_26,
	0xc7: RST_C7,
	0xe2: LDH_E2,
	0x18: JR_18,
	0x3a: LD_3A,
	0x55: LD_55,
	0xac: XOR_AC,
	0x71: LD_71,
	0xf9: LD_F9,
	0x5f: LD_5F,
	0xb7: OR_B7,
	0xc9: RET_C9,
	0xee: XOR_EE,
	0xf3: DI_F3,
	0x38: JR_38,
	0xc2: JP_C2,
	0xc5: PUSH_C5,
	0x11: LD_11,
	0x13: INC_13,
	0x2c: INC_2C,
	0x2e: LD_2E,
	0x79: LD_79,
	0x8d: ADC_8D,
	0x97: SUB_97,
	0xdd: ILLEGAL_DD_DD,
	0x51: LD_51,
	0x65: LD_65,
	0xb1: OR_B1,
	0xea: LD_EA,
	0x88: ADC_88,
	0xdf: RST_DF,
	0x20: JR_20,
	0x47: LD_47,
	0x7b: LD_7B,
	0xb2: OR_B2,
	0xbf: CP_BF,
	0xde: SBC_DE,
	0xfe: CP_FE,
	0x1d: DEC_1D,
	0x1f: RRA_1F,
	0x72: LD_72,
	0x8c: ADC_8C,
	0xbe: CP_BE,
	0x28: JR_28,
	0xb9: CP_B9,
	0x3c: INC_3C,
	0x50: LD_50,
	0x7e: LD_7E,
	0xc8: RET_C8,
	0xcb: PREFIX_CB,
	0xcf: RST_CF,
	0xd8: RET_D8,
	0x8:  LD_08,
	0xc:  INC_0C,
	0x37: SCF_37,
	0x43: LD_43,
	0xbd: CP_BD,
	0x1b: DEC_1B,
	0x3b: DEC_3B,
	0x73: LD_73,
	0xca: JP_CA,
	0xd0: RET_D0,
	0xf8: LD_F8,
	0x6a: LD_6A,
	0x6e: LD_6E,
	0xeb: ILLEGAL_EB_EB,
	0xf1: POP_F1,
	0x59: LD_59,
	0x6f: LD_6F,
	0x96: SUB_96,
	0x19: ADD_19,
	0x2f: CPL_2F,
	0x7c: LD_7C,
	0x3:  INC_03,
	0x30: JR_30,
	0x48: LD_48,
	0x66: LD_66,
	0x9b: SBC_9B,
	0xa7: AND_A7,
	0x41: LD_41,
	0x45: LD_45,
	0x90: SUB_90,
	0xa6: AND_A6,
	0xd1: POP_D1,
	0xd9: RETI_D9,
	0x4d: LD_4D,
	0x86: ADD_86,
	0x9c: SBC_9C,
	0xdc: CALL_DC,
	0xec: ILLEGAL_EC_EC,
	0x2a: LD_2A,
	0x0:  NOP_00,
	0x9:  ADD_09,
	0x40: LD_40,
	0x7f: LD_7F,
	0x24: INC_24,
	0xe0: LDH_E0,
	0xf0: LDH_F0,
	0x3f: CCF_3F,
	0x67: LD_67,
	0xb0: OR_B0,
	0xdb: ILLEGAL_DB_DB,
	0x56: LD_56,
	0xfd: ILLEGAL_FD_FD,
	0x1a: LD_1A,
	0xd:  DEC_0D,
	0x9f: SBC_9F,
	0x14: INC_14,
	0x36: LD_36,
	0x44: LD_44,
	0x4e: LD_4E,
	0x77: LD_77,
	0x82: ADD_82,
	0x6d: LD_6D,
	0x81: ADD_81,
	0xb3: OR_B3,
	0xb:  DEC_0B,
	0x6c: LD_6C,
	0xb4: OR_B4,
	0xb6: OR_B6,
	0xc1: POP_C1,
	0x9d: SBC_9D,
	0xbc: CP_BC,
	0xf6: OR_F6,
	0x15: DEC_15,
	0x34: INC_34,
	0xf7: RST_F7,
	0x1c: INC_1C,
	0x62: LD_62,
	0x8a: ADC_8A,
	0xaa: XOR_AA,
	0xae: XOR_AE,
	0xfa: LD_FA,
	0x74: LD_74,
	0x78: LD_78,
	0x84: ADD_84,
	0x89: ADC_89,
	0x93: SUB_93,
	0x98: SBC_98,
	0x1e: LD_1E,
	0x21: LD_21,
	0x42: LD_42,
	0x95: SUB_95,
	0xe1: POP_E1,
	0xed: ILLEGAL_ED_ED,
	0x85: ADD_85,
	0xc4: CALL_C4,
	0x1:  LD_01,
	0x10: STOP_10,
	0x23: INC_23,
	0x39: ADD_39,
	0x9e: SBC_9E,
	0x99: SBC_99,
	0xd2: JP_D2,
	0xd3: ILLEGAL_D3_D3,
	0xf5: PUSH_F5,
	0x6:  LD_06,
	0x4f: LD_4F,
	0x5d: LD_5D,
	0x5a: LD_5A,
	0x91: SUB_91,
	0x27: DAA_27,
	0x61: LD_61,
	0x92: SUB_92,
	0xa8: XOR_A8,
	0xcd: CALL_CD,
	0x63: LD_63,
	0xd6: SUB_D6,
	0xe3: ILLEGAL_E3_E3,
	0x3e: LD_3E,
}
var extOps = map[uint8]Instruction{
	0x83: RES_83,
	0xec: SET_EC,
	0xf5: SET_F5,
	0xfb: SET_FB,
	0x60: BIT_60,
	0x78: BIT_78,
	0xc3: SET_C3,
	0x52: BIT_52,
	0x5c: BIT_5C,
	0x24: SLA_24,
	0x41: BIT_41,
	0x4b: BIT_4B,
	0xb5: RES_B5,
	0x1c: RR_1C,
	0x57: BIT_57,
	0x71: BIT_71,
	0x9c: RES_9C,
	0xe1: SET_E1,
	0x8e: RES_8E,
	0xa2: RES_A2,
	0xf4: SET_F4,
	0xda: SET_DA,
	0x13: RL_13,
	0x26: SLA_26,
	0x2b: SRA_2B,
	0x4c: BIT_4C,
	0xa4: RES_A4,
	0xc1: SET_C1,
	0xe8: SET_E8,
	0xb:  RRC_0B,
	0x4a: BIT_4A,
	0x86: RES_86,
	0xeb: SET_EB,
	0xe:  RRC_0E,
	0x22: SLA_22,
	0x38: SRL_38,
	0xcd: SET_CD,
	0x48: BIT_48,
	0x76: BIT_76,
	0x8c: RES_8C,
	0x94: RES_94,
	0xc4: SET_C4,
	0x32: SWAP_32,
	0xb1: RES_B1,
	0xbb: RES_BB,
	0xc2: SET_C2,
	0xd5: SET_D5,
	0x9:  RRC_09,
	0x10: RL_10,
	0x3a: SRL_3A,
	0x56: BIT_56,
	0x5e: BIT_5E,
	0xb7: RES_B7,
	0xdb: SET_DB,
	0x4d: BIT_4D,
	0x6f: BIT_6F,
	0x79: BIT_79,
	0x85: RES_85,
	0xb0: RES_B0,
	0xc9: SET_C9,
	0xd1: SET_D1,
	0x8b: RES_8B,
	0xaf: RES_AF,
	0xf7: SET_F7,
	0x18: RR_18,
	0x62: BIT_62,
	0x8f: RES_8F,
	0xbc: RES_BC,
	0xe9: SET_E9,
	0x2c: SRA_2C,
	0x47: BIT_47,
	0x77: BIT_77,
	0xab: RES_AB,
	0xf9: SET_F9,
	0x1e: RR_1E,
	0x11: RL_11,
	0x2f: SRA_2F,
	0x51: BIT_51,
	0x7e: BIT_7E,
	0x88: RES_88,
	0x95: RES_95,
	0x9e: RES_9E,
	0x19: RR_19,
	0x9a: RES_9A,
	0x6:  RLC_06,
	0x31: SWAP_31,
	0x49: BIT_49,
	0xcc: SET_CC,
	0x33: SWAP_33,
	0x58: BIT_58,
	0x96: RES_96,
	0xa0: RES_A0,
	0xa1: RES_A1,
	0xc7: SET_C7,
	0xd6: SET_D6,
	0x4:  RLC_04,
	0x37: SWAP_37,
	0x43: BIT_43,
	0x28: SRA_28,
	0x90: RES_90,
	0xce: SET_CE,
	0x81: RES_81,
	0xa7: RES_A7,
	0xd4: SET_D4,
	0xe0: SET_E0,
	0x3:  RLC_03,
	0xa:  RRC_0A,
	0x36: SWAP_36,
	0xa5: RES_A5,
	0x0:  RLC_00,
	0x12: RL_12,
	0x30: SWAP_30,
	0x8d: RES_8D,
	0x98: RES_98,
	0x14: RL_14,
	0x1d: RR_1D,
	0x25: SLA_25,
	0x9f: RES_9F,
	0xca: SET_CA,
	0x35: SWAP_35,
	0x66: BIT_66,
	0x6d: BIT_6D,
	0x92: RES_92,
	0xa8: RES_A8,
	0xd0: SET_D0,
	0x1f: RR_1F,
	0x5d: BIT_5D,
	0x80: RES_80,
	0x87: RES_87,
	0x97: RES_97,
	0xd3: SET_D3,
	0xe5: SET_E5,
	0xf0: SET_F0,
	0x40: BIT_40,
	0x82: RES_82,
	0xc0: SET_C0,
	0x1b: RR_1B,
	0x7a: BIT_7A,
	0xb9: RES_B9,
	0xd9: SET_D9,
	0x50: BIT_50,
	0xb2: RES_B2,
	0x20: SLA_20,
	0x27: SLA_27,
	0x34: SWAP_34,
	0x59: BIT_59,
	0x7d: BIT_7D,
	0x6b: BIT_6B,
	0x54: BIT_54,
	0xac: RES_AC,
	0xbe: RES_BE,
	0x5:  RLC_05,
	0x84: RES_84,
	0xbd: RES_BD,
	0x6c: BIT_6C,
	0x2d: SRA_2D,
	0x3c: SRL_3C,
	0x70: BIT_70,
	0xba: RES_BA,
	0x5a: BIT_5A,
	0x8a: RES_8A,
	0xf2: SET_F2,
	0x23: SLA_23,
	0x3b: SRL_3B,
	0x9b: RES_9B,
	0xa3: RES_A3,
	0xa6: RES_A6,
	0xc5: SET_C5,
	0xcf: SET_CF,
	0xe6: SET_E6,
	0x15: RL_15,
	0x3e: SRL_3E,
	0xb3: RES_B3,
	0xb6: RES_B6,
	0xbf: RES_BF,
	0x7c: BIT_7C,
	0x91: RES_91,
	0x16: RL_16,
	0x9d: RES_9D,
	0xa9: RES_A9,
	0xe7: SET_E7,
	0x3f: SRL_3F,
	0xdc: SET_DC,
	0xff: SET_FF,
	0xf:  RRC_0F,
	0x44: BIT_44,
	0x67: BIT_67,
	0x7f: BIT_7F,
	0xfe: SET_FE,
	0x72: BIT_72,
	0xb4: RES_B4,
	0xea: SET_EA,
	0x6e: BIT_6E,
	0x99: RES_99,
	0xde: SET_DE,
	0x17: RL_17,
	0x6a: BIT_6A,
	0xd2: SET_D2,
	0xdd: SET_DD,
	0x8:  RRC_08,
	0x73: BIT_73,
	0x75: BIT_75,
	0xc8: SET_C8,
	0xed: SET_ED,
	0x69: BIT_69,
	0x2:  RLC_02,
	0x1a: RR_1A,
	0x29: SRA_29,
	0x42: BIT_42,
	0x7b: BIT_7B,
	0x93: RES_93,
	0xf6: SET_F6,
	0xc:  RRC_0C,
	0x53: BIT_53,
	0x61: BIT_61,
	0x64: BIT_64,
	0xae: RES_AE,
	0xee: SET_EE,
	0xf3: SET_F3,
	0x39: SRL_39,
	0xaa: RES_AA,
	0xe4: SET_E4,
	0x46: BIT_46,
	0xf8: SET_F8,
	0xd:  RRC_0D,
	0x5f: BIT_5F,
	0x63: BIT_63,
	0xc6: SET_C6,
	0x1:  RLC_01,
	0x2e: SRA_2E,
	0x45: BIT_45,
	0x55: BIT_55,
	0x4e: BIT_4E,
	0x65: BIT_65,
	0xcb: SET_CB,
	0xdf: SET_DF,
	0xe2: SET_E2,
	0xe3: SET_E3,
	0x7:  RLC_07,
	0x2a: SRA_2A,
	0x68: BIT_68,
	0xad: RES_AD,
	0xd8: SET_D8,
	0xfa: SET_FA,
	0xfd: SET_FD,
	0x21: SLA_21,
	0xb8: RES_B8,
	0x5b: BIT_5B,
	0xd7: SET_D7,
	0xf1: SET_F1,
	0x3d: SRL_3D,
	0x74: BIT_74,
	0xef: SET_EF,
	0xfc: SET_FC,
	0x4f: BIT_4F,
	0x89: RES_89,
}

// returns code given a string. Useful during testing
func code(s string) uint8 {
	switch s {
	case "DEC HL":
		return 0x2B
	case "LD C,E":
		return 0x4B
	case "LD D,A":
		return 0x57
	case "LD A,D":
		return 0x7A
	case "SBC A,D":
		return 0x9A
	case "AND A,C":
		return 0xA1
	case "AND A,E":
		return 0xA3
	case "PUSH DE":
		return 0xD5
	case "LD E,B":
		return 0x58
	case "LD (HL),L":
		return 0x75
	case "RET NZ":
		return 0xC0
	case "ADC A,n8":
		return 0xCE
	case "ADD SP,e8":
		return 0xE8
	case "RST $38":
		return 0xFF
	case "LD C,n8":
		return 0xE
	case "LD (HL+),A":
		return 0x22
	case "LD H,B":
		return 0x60
	case "LD L,E":
		return 0x6B
	case "DEC B":
		return 0x5
	case "LD A,(BC)":
		return 0xA
	case "DEC A":
		return 0x3D
	case "LD E,E":
		return 0x5B
	case "LD E,H":
		return 0x5C
	case "XOR A,C":
		return 0xA9
	case "CP A,B":
		return 0xB8
	case "JP C,a16":
		return 0xDA
	case "RST $20":
		return 0xE7
	case "DEC (HL)":
		return 0x35
	case "LD C,D":
		return 0x4A
	case "LD L,C":
		return 0x69
	case "ADD A,B":
		return 0x80
	case "XOR A,L":
		return 0xAD
	case "JP HL":
		return 0xE9
	case "LD SP,n16":
		return 0x31
	case "DEC H":
		return 0x25
	case "LD C,H":
		return 0x4C
	case "LD D,E":
		return 0x53
	case "LD L,B":
		return 0x68
	case "SUB A,H":
		return 0x94
	case "ILLEGAL_E4":
		return 0xE4
	case "LD C,C":
		return 0x49
	case "LD E,(HL)":
		return 0x5E
	case "CP A,D":
		return 0xBA
	case "RST $10":
		return 0xD7
	case "LD (BC),A":
		return 0x2
	case "INC SP":
		return 0x33
	case "LD D,H":
		return 0x54
	case "ADC A,A":
		return 0x8F
	case "OR A,L":
		return 0xB5
	case "INC B":
		return 0x4
	case "LD A,L":
		return 0x7D
	case "PUSH HL":
		return 0xE5
	case "AND A,n8":
		return 0xE6
	case "ILLEGAL_F4":
		return 0xF4
	case "LD (HL-),A":
		return 0x32
	case "ADD HL,HL":
		return 0x29
	case "HALT":
		return 0x76
	case "ADD A,E":
		return 0x83
	case "XOR A,E":
		return 0xAB
	case "CP A,E":
		return 0xBB
	case "ADD A,n8":
		return 0xC6
	case "LD (DE),A":
		return 0x12
	case "ADC A,E":
		return 0x8B
	case "RST $28":
		return 0xEF
	case "DEC L":
		return 0x2D
	case "LD B,(HL)":
		return 0x46
	case "LD H,H":
		return 0x64
	case "AND A,H":
		return 0xA4
	case "XOR A,A":
		return 0xAF
	case "CALL NC,a16":
		return 0xD4
	case "RLCA":
		return 0x7
	case "LD (HL),B":
		return 0x70
	case "AND A,L":
		return 0xA5
	case "LDH A,(C)":
		return 0xF2
	case "LD D,D":
		return 0x52
	case "JP a16":
		return 0xC3
	case "LD D,n8":
		return 0x16
	case "RLA":
		return 0x17
	case "AND A,B":
		return 0xA0
	case "AND A,D":
		return 0xA2
	case "EI":
		return 0xFB
	case "ILLEGAL_FC":
		return 0xFC
	case "ADD A,A":
		return 0x87
	case "ADC A,(HL)":
		return 0x8E
	case "CALL Z,a16":
		return 0xCC
	case "LD H,n8":
		return 0x26
	case "RST $00":
		return 0xC7
	case "LDH (C),A":
		return 0xE2
	case "JR e8":
		return 0x18
	case "LD A,(HL-)":
		return 0x3A
	case "LD D,L":
		return 0x55
	case "XOR A,H":
		return 0xAC
	case "LD (HL),C":
		return 0x71
	case "LD SP,HL":
		return 0xF9
	case "LD E,A":
		return 0x5F
	case "OR A,A":
		return 0xB7
	case "RET":
		return 0xC9
	case "XOR A,n8":
		return 0xEE
	case "DI":
		return 0xF3
	case "JR C,e8":
		return 0x38
	case "JP NZ,a16":
		return 0xC2
	case "PUSH BC":
		return 0xC5
	case "LD DE,n16":
		return 0x11
	case "INC DE":
		return 0x13
	case "INC L":
		return 0x2C
	case "LD L,n8":
		return 0x2E
	case "LD A,C":
		return 0x79
	case "ADC A,L":
		return 0x8D
	case "SUB A,A":
		return 0x97
	case "ILLEGAL_DD":
		return 0xDD
	case "LD D,C":
		return 0x51
	case "LD H,L":
		return 0x65
	case "OR A,C":
		return 0xB1
	case "LD (a16),A":
		return 0xEA
	case "ADC A,B":
		return 0x88
	case "RST $18":
		return 0xDF
	case "JR NZ,e8":
		return 0x20
	case "LD B,A":
		return 0x47
	case "LD A,E":
		return 0x7B
	case "OR A,D":
		return 0xB2
	case "CP A,A":
		return 0xBF
	case "SBC A,n8":
		return 0xDE
	case "CP A,n8":
		return 0xFE
	case "DEC E":
		return 0x1D
	case "RRA":
		return 0x1F
	case "LD (HL),D":
		return 0x72
	case "ADC A,H":
		return 0x8C
	case "CP A,(HL)":
		return 0xBE
	case "JR Z,e8":
		return 0x28
	case "CP A,C":
		return 0xB9
	case "INC A":
		return 0x3C
	case "LD D,B":
		return 0x50
	case "LD A,(HL)":
		return 0x7E
	case "RET Z":
		return 0xC8
	case "PREFIX":
		return 0xCB
	case "RST $08":
		return 0xCF
	case "RET C":
		return 0xD8
	case "LD (a16),SP":
		return 0x8
	case "INC C":
		return 0xC
	case "SCF":
		return 0x37
	case "LD B,E":
		return 0x43
	case "CP A,L":
		return 0xBD
	case "DEC DE":
		return 0x1B
	case "DEC SP":
		return 0x3B
	case "LD (HL),E":
		return 0x73
	case "JP Z,a16":
		return 0xCA
	case "RET NC":
		return 0xD0
	case "LD HL,SP+,e8":
		return 0xF8
	case "LD L,D":
		return 0x6A
	case "LD L,(HL)":
		return 0x6E
	case "ILLEGAL_EB":
		return 0xEB
	case "POP AF":
		return 0xF1
	case "LD E,C":
		return 0x59
	case "LD L,A":
		return 0x6F
	case "SUB A,(HL)":
		return 0x96
	case "ADD HL,DE":
		return 0x19
	case "CPL":
		return 0x2F
	case "LD A,H":
		return 0x7C
	case "INC BC":
		return 0x3
	case "JR NC,e8":
		return 0x30
	case "LD C,B":
		return 0x48
	case "LD H,(HL)":
		return 0x66
	case "SBC A,E":
		return 0x9B
	case "AND A,A":
		return 0xA7
	case "LD B,C":
		return 0x41
	case "LD B,L":
		return 0x45
	case "SUB A,B":
		return 0x90
	case "AND A,(HL)":
		return 0xA6
	case "POP DE":
		return 0xD1
	case "RETI":
		return 0xD9
	case "LD C,L":
		return 0x4D
	case "ADD A,(HL)":
		return 0x86
	case "SBC A,H":
		return 0x9C
	case "CALL C,a16":
		return 0xDC
	case "ILLEGAL_EC":
		return 0xEC
	case "LD A,(HL+)":
		return 0x2A
	case "NOP":
		return 0x0
	case "ADD HL,BC":
		return 0x9
	case "LD B,B":
		return 0x40
	case "LD A,A":
		return 0x7F
	case "INC H":
		return 0x24
	case "LDH (a8),A":
		return 0xE0
	case "LDH A,(a8)":
		return 0xF0
	case "CCF":
		return 0x3F
	case "LD H,A":
		return 0x67
	case "OR A,B":
		return 0xB0
	case "ILLEGAL_DB":
		return 0xDB
	case "LD D,(HL)":
		return 0x56
	case "ILLEGAL_FD":
		return 0xFD
	case "LD A,(DE)":
		return 0x1A
	case "DEC C":
		return 0xD
	case "SBC A,A":
		return 0x9F
	case "INC D":
		return 0x14
	case "LD (HL),n8":
		return 0x36
	case "LD B,H":
		return 0x44
	case "LD C,(HL)":
		return 0x4E
	case "LD (HL),A":
		return 0x77
	case "ADD A,D":
		return 0x82
	case "LD L,L":
		return 0x6D
	case "ADD A,C":
		return 0x81
	case "OR A,E":
		return 0xB3
	case "DEC BC":
		return 0xB
	case "LD L,H":
		return 0x6C
	case "OR A,H":
		return 0xB4
	case "OR A,(HL)":
		return 0xB6
	case "POP BC":
		return 0xC1
	case "SBC A,L":
		return 0x9D
	case "CP A,H":
		return 0xBC
	case "OR A,n8":
		return 0xF6
	case "DEC D":
		return 0x15
	case "INC (HL)":
		return 0x34
	case "RST $30":
		return 0xF7
	case "INC E":
		return 0x1C
	case "LD H,D":
		return 0x62
	case "ADC A,D":
		return 0x8A
	case "XOR A,D":
		return 0xAA
	case "XOR A,(HL)":
		return 0xAE
	case "LD A,(a16)":
		return 0xFA
	case "LD (HL),H":
		return 0x74
	case "LD A,B":
		return 0x78
	case "ADD A,H":
		return 0x84
	case "ADC A,C":
		return 0x89
	case "SUB A,E":
		return 0x93
	case "SBC A,B":
		return 0x98
	case "LD E,n8":
		return 0x1E
	case "LD HL,n16":
		return 0x21
	case "LD B,D":
		return 0x42
	case "SUB A,L":
		return 0x95
	case "POP HL":
		return 0xE1
	case "ILLEGAL_ED":
		return 0xED
	case "ADD A,L":
		return 0x85
	case "CALL NZ,a16":
		return 0xC4
	case "LD BC,n16":
		return 0x1
	case "STOP":
		return 0x10
	case "INC HL":
		return 0x23
	case "ADD HL,SP":
		return 0x39
	case "SBC A,(HL)":
		return 0x9E
	case "SBC A,C":
		return 0x99
	case "JP NC,a16":
		return 0xD2
	case "ILLEGAL_D3":
		return 0xD3
	case "PUSH AF":
		return 0xF5
	case "LD B,n8":
		return 0x6
	case "LD C,A":
		return 0x4F
	case "LD E,L":
		return 0x5D
	case "LD E,D":
		return 0x5A
	case "SUB A,C":
		return 0x91
	case "DAA":
		return 0x27
	case "LD H,C":
		return 0x61
	case "SUB A,D":
		return 0x92
	case "XOR A,B":
		return 0xA8
	case "CALL a16":
		return 0xCD
	case "LD H,E":
		return 0x63
	case "SUB A,n8":
		return 0xD6
	case "ILLEGAL_E3":
		return 0xE3
	case "LD A,n8":
		return 0x3E
	case "RES 0,E":
		return 0x83
	case "SET 5,H":
		return 0xEC
	case "SET 6,L":
		return 0xF5
	case "SET 7,E":
		return 0xFB
	case "BIT 4,B":
		return 0x60
	case "BIT 7,B":
		return 0x78
	case "SET 0,E":
		return 0xC3
	case "BIT 2,D":
		return 0x52
	case "BIT 3,H":
		return 0x5C
	case "SLA H":
		return 0x24
	case "BIT 0,C":
		return 0x41
	case "BIT 1,E":
		return 0x4B
	case "RES 6,L":
		return 0xB5
	case "RR H":
		return 0x1C
	case "BIT 2,A":
		return 0x57
	case "BIT 6,C":
		return 0x71
	case "RES 3,H":
		return 0x9C
	case "SET 4,C":
		return 0xE1
	case "RES 1,(HL)":
		return 0x8E
	case "RES 4,D":
		return 0xA2
	case "SET 6,H":
		return 0xF4
	case "SET 3,D":
		return 0xDA
	case "RL E":
		return 0x13
	case "SLA (HL)":
		return 0x26
	case "SRA E":
		return 0x2B
	case "BIT 1,H":
		return 0x4C
	case "RES 4,H":
		return 0xA4
	case "SET 0,C":
		return 0xC1
	case "SET 5,B":
		return 0xE8
	case "RRC E":
		return 0xB
	case "BIT 1,D":
		return 0x4A
	case "RES 0,(HL)":
		return 0x86
	case "SET 5,E":
		return 0xEB
	case "RRC (HL)":
		return 0xE
	case "SLA D":
		return 0x22
	case "SRL B":
		return 0x38
	case "SET 1,L":
		return 0xCD
	case "BIT 1,B":
		return 0x48
	case "BIT 6,(HL)":
		return 0x76
	case "RES 1,H":
		return 0x8C
	case "RES 2,H":
		return 0x94
	case "SET 0,H":
		return 0xC4
	case "SWAP D":
		return 0x32
	case "RES 6,C":
		return 0xB1
	case "RES 7,E":
		return 0xBB
	case "SET 0,D":
		return 0xC2
	case "SET 2,L":
		return 0xD5
	case "RRC C":
		return 0x9
	case "RL B":
		return 0x10
	case "SRL D":
		return 0x3A
	case "BIT 2,(HL)":
		return 0x56
	case "BIT 3,(HL)":
		return 0x5E
	case "RES 6,A":
		return 0xB7
	case "SET 3,E":
		return 0xDB
	case "BIT 1,L":
		return 0x4D
	case "BIT 5,A":
		return 0x6F
	case "BIT 7,C":
		return 0x79
	case "RES 0,L":
		return 0x85
	case "RES 6,B":
		return 0xB0
	case "SET 1,C":
		return 0xC9
	case "SET 2,C":
		return 0xD1
	case "RES 1,E":
		return 0x8B
	case "RES 5,A":
		return 0xAF
	case "SET 6,A":
		return 0xF7
	case "RR B":
		return 0x18
	case "BIT 4,D":
		return 0x62
	case "RES 1,A":
		return 0x8F
	case "RES 7,H":
		return 0xBC
	case "SET 5,C":
		return 0xE9
	case "SRA H":
		return 0x2C
	case "BIT 0,A":
		return 0x47
	case "BIT 6,A":
		return 0x77
	case "RES 5,E":
		return 0xAB
	case "SET 7,C":
		return 0xF9
	case "RR (HL)":
		return 0x1E
	case "RL C":
		return 0x11
	case "SRA A":
		return 0x2F
	case "BIT 2,C":
		return 0x51
	case "BIT 7,(HL)":
		return 0x7E
	case "RES 1,B":
		return 0x88
	case "RES 2,L":
		return 0x95
	case "RES 3,(HL)":
		return 0x9E
	case "RR C":
		return 0x19
	case "RES 3,D":
		return 0x9A
	case "RLC (HL)":
		return 0x6
	case "SWAP C":
		return 0x31
	case "BIT 1,C":
		return 0x49
	case "SET 1,H":
		return 0xCC
	case "SWAP E":
		return 0x33
	case "BIT 3,B":
		return 0x58
	case "RES 2,(HL)":
		return 0x96
	case "RES 4,B":
		return 0xA0
	case "RES 4,C":
		return 0xA1
	case "SET 0,A":
		return 0xC7
	case "SET 2,(HL)":
		return 0xD6
	case "RLC H":
		return 0x4
	case "SWAP A":
		return 0x37
	case "BIT 0,E":
		return 0x43
	case "SRA B":
		return 0x28
	case "RES 2,B":
		return 0x90
	case "SET 1,(HL)":
		return 0xCE
	case "RES 0,C":
		return 0x81
	case "RES 4,A":
		return 0xA7
	case "SET 2,H":
		return 0xD4
	case "SET 4,B":
		return 0xE0
	case "RLC E":
		return 0x3
	case "RRC D":
		return 0xA
	case "SWAP (HL)":
		return 0x36
	case "RES 4,L":
		return 0xA5
	case "RLC B":
		return 0x0
	case "RL D":
		return 0x12
	case "SWAP B":
		return 0x30
	case "RES 1,L":
		return 0x8D
	case "RES 3,B":
		return 0x98
	case "RL H":
		return 0x14
	case "RR L":
		return 0x1D
	case "SLA L":
		return 0x25
	case "RES 3,A":
		return 0x9F
	case "SET 1,D":
		return 0xCA
	case "SWAP L":
		return 0x35
	case "BIT 4,(HL)":
		return 0x66
	case "BIT 5,L":
		return 0x6D
	case "RES 2,D":
		return 0x92
	case "RES 5,B":
		return 0xA8
	case "SET 2,B":
		return 0xD0
	case "RR A":
		return 0x1F
	case "BIT 3,L":
		return 0x5D
	case "RES 0,B":
		return 0x80
	case "RES 0,A":
		return 0x87
	case "RES 2,A":
		return 0x97
	case "SET 2,E":
		return 0xD3
	case "SET 4,L":
		return 0xE5
	case "SET 6,B":
		return 0xF0
	case "BIT 0,B":
		return 0x40
	case "RES 0,D":
		return 0x82
	case "SET 0,B":
		return 0xC0
	case "RR E":
		return 0x1B
	case "BIT 7,D":
		return 0x7A
	case "RES 7,C":
		return 0xB9
	case "SET 3,C":
		return 0xD9
	case "BIT 2,B":
		return 0x50
	case "RES 6,D":
		return 0xB2
	case "SLA B":
		return 0x20
	case "SLA A":
		return 0x27
	case "SWAP H":
		return 0x34
	case "BIT 3,C":
		return 0x59
	case "BIT 7,L":
		return 0x7D
	case "BIT 5,E":
		return 0x6B
	case "BIT 2,H":
		return 0x54
	case "RES 5,H":
		return 0xAC
	case "RES 7,(HL)":
		return 0xBE
	case "RLC L":
		return 0x5
	case "RES 0,H":
		return 0x84
	case "RES 7,L":
		return 0xBD
	case "BIT 5,H":
		return 0x6C
	case "SRA L":
		return 0x2D
	case "SRL H":
		return 0x3C
	case "BIT 6,B":
		return 0x70
	case "RES 7,D":
		return 0xBA
	case "BIT 3,D":
		return 0x5A
	case "RES 1,D":
		return 0x8A
	case "SET 6,D":
		return 0xF2
	case "SLA E":
		return 0x23
	case "SRL E":
		return 0x3B
	case "RES 3,E":
		return 0x9B
	case "RES 4,E":
		return 0xA3
	case "RES 4,(HL)":
		return 0xA6
	case "SET 0,L":
		return 0xC5
	case "SET 1,A":
		return 0xCF
	case "SET 4,(HL)":
		return 0xE6
	case "RL L":
		return 0x15
	case "SRL (HL)":
		return 0x3E
	case "RES 6,E":
		return 0xB3
	case "RES 6,(HL)":
		return 0xB6
	case "RES 7,A":
		return 0xBF
	case "BIT 7,H":
		return 0x7C
	case "RES 2,C":
		return 0x91
	case "RL (HL)":
		return 0x16
	case "RES 3,L":
		return 0x9D
	case "RES 5,C":
		return 0xA9
	case "SET 4,A":
		return 0xE7
	case "SRL A":
		return 0x3F
	case "SET 3,H":
		return 0xDC
	case "SET 7,A":
		return 0xFF
	case "RRC A":
		return 0xF
	case "BIT 0,H":
		return 0x44
	case "BIT 4,A":
		return 0x67
	case "BIT 7,A":
		return 0x7F
	case "SET 7,(HL)":
		return 0xFE
	case "BIT 6,D":
		return 0x72
	case "RES 6,H":
		return 0xB4
	case "SET 5,D":
		return 0xEA
	case "BIT 5,(HL)":
		return 0x6E
	case "RES 3,C":
		return 0x99
	case "SET 3,(HL)":
		return 0xDE
	case "RL A":
		return 0x17
	case "BIT 5,D":
		return 0x6A
	case "SET 2,D":
		return 0xD2
	case "SET 3,L":
		return 0xDD
	case "RRC B":
		return 0x8
	case "BIT 6,E":
		return 0x73
	case "BIT 6,L":
		return 0x75
	case "SET 1,B":
		return 0xC8
	case "SET 5,L":
		return 0xED
	case "BIT 5,C":
		return 0x69
	case "RLC D":
		return 0x2
	case "RR D":
		return 0x1A
	case "SRA C":
		return 0x29
	case "BIT 0,D":
		return 0x42
	case "BIT 7,E":
		return 0x7B
	case "RES 2,E":
		return 0x93
	case "SET 6,(HL)":
		return 0xF6
	case "RRC H":
		return 0xC
	case "BIT 2,E":
		return 0x53
	case "BIT 4,C":
		return 0x61
	case "BIT 4,H":
		return 0x64
	case "RES 5,(HL)":
		return 0xAE
	case "SET 5,(HL)":
		return 0xEE
	case "SET 6,E":
		return 0xF3
	case "SRL C":
		return 0x39
	case "RES 5,D":
		return 0xAA
	case "SET 4,H":
		return 0xE4
	case "BIT 0,(HL)":
		return 0x46
	case "SET 7,B":
		return 0xF8
	case "RRC L":
		return 0xD
	case "BIT 3,A":
		return 0x5F
	case "BIT 4,E":
		return 0x63
	case "SET 0,(HL)":
		return 0xC6
	case "RLC C":
		return 0x1
	case "SRA (HL)":
		return 0x2E
	case "BIT 0,L":
		return 0x45
	case "BIT 2,L":
		return 0x55
	case "BIT 1,(HL)":
		return 0x4E
	case "BIT 4,L":
		return 0x65
	case "SET 1,E":
		return 0xCB
	case "SET 3,A":
		return 0xDF
	case "SET 4,D":
		return 0xE2
	case "SET 4,E":
		return 0xE3
	case "RLC A":
		return 0x7
	case "SRA D":
		return 0x2A
	case "BIT 5,B":
		return 0x68
	case "RES 5,L":
		return 0xAD
	case "SET 3,B":
		return 0xD8
	case "SET 7,D":
		return 0xFA
	case "SET 7,L":
		return 0xFD
	case "SLA C":
		return 0x21
	case "RES 7,B":
		return 0xB8
	case "BIT 3,E":
		return 0x5B
	case "SET 2,A":
		return 0xD7
	case "SET 6,C":
		return 0xF1
	case "SRL L":
		return 0x3D
	case "BIT 6,H":
		return 0x74
	case "SET 5,A":
		return 0xEF
	case "SET 7,H":
		return 0xFC
	case "BIT 1,A":
		return 0x4F
	case "RES 1,C":
		return 0x89

	default:
		panic(fmt.Sprintf("Unknown code for %q", s))
	}
}
func name(code uint8, prefix bool) string {
	if prefix {
		switch code {
		case 131:
			return "RES 0,E"
		case 236:
			return "SET 5,H"
		case 245:
			return "SET 6,L"
		case 251:
			return "SET 7,E"
		case 96:
			return "BIT 4,B"
		case 120:
			return "BIT 7,B"
		case 195:
			return "SET 0,E"
		case 82:
			return "BIT 2,D"
		case 92:
			return "BIT 3,H"
		case 36:
			return "SLA H"
		case 65:
			return "BIT 0,C"
		case 75:
			return "BIT 1,E"
		case 181:
			return "RES 6,L"
		case 28:
			return "RR H"
		case 87:
			return "BIT 2,A"
		case 113:
			return "BIT 6,C"
		case 156:
			return "RES 3,H"
		case 225:
			return "SET 4,C"
		case 142:
			return "RES 1,(HL)"
		case 162:
			return "RES 4,D"
		case 244:
			return "SET 6,H"
		case 218:
			return "SET 3,D"
		case 19:
			return "RL E"
		case 38:
			return "SLA (HL)"
		case 43:
			return "SRA E"
		case 76:
			return "BIT 1,H"
		case 164:
			return "RES 4,H"
		case 193:
			return "SET 0,C"
		case 232:
			return "SET 5,B"
		case 11:
			return "RRC E"
		case 74:
			return "BIT 1,D"
		case 134:
			return "RES 0,(HL)"
		case 235:
			return "SET 5,E"
		case 14:
			return "RRC (HL)"
		case 34:
			return "SLA D"
		case 56:
			return "SRL B"
		case 205:
			return "SET 1,L"
		case 72:
			return "BIT 1,B"
		case 118:
			return "BIT 6,(HL)"
		case 140:
			return "RES 1,H"
		case 148:
			return "RES 2,H"
		case 196:
			return "SET 0,H"
		case 50:
			return "SWAP D"
		case 177:
			return "RES 6,C"
		case 187:
			return "RES 7,E"
		case 194:
			return "SET 0,D"
		case 213:
			return "SET 2,L"
		case 9:
			return "RRC C"
		case 16:
			return "RL B"
		case 58:
			return "SRL D"
		case 86:
			return "BIT 2,(HL)"
		case 94:
			return "BIT 3,(HL)"
		case 183:
			return "RES 6,A"
		case 219:
			return "SET 3,E"
		case 77:
			return "BIT 1,L"
		case 111:
			return "BIT 5,A"
		case 121:
			return "BIT 7,C"
		case 133:
			return "RES 0,L"
		case 176:
			return "RES 6,B"
		case 201:
			return "SET 1,C"
		case 209:
			return "SET 2,C"
		case 139:
			return "RES 1,E"
		case 175:
			return "RES 5,A"
		case 247:
			return "SET 6,A"
		case 24:
			return "RR B"
		case 98:
			return "BIT 4,D"
		case 143:
			return "RES 1,A"
		case 188:
			return "RES 7,H"
		case 233:
			return "SET 5,C"
		case 44:
			return "SRA H"
		case 71:
			return "BIT 0,A"
		case 119:
			return "BIT 6,A"
		case 171:
			return "RES 5,E"
		case 249:
			return "SET 7,C"
		case 30:
			return "RR (HL)"
		case 17:
			return "RL C"
		case 47:
			return "SRA A"
		case 81:
			return "BIT 2,C"
		case 126:
			return "BIT 7,(HL)"
		case 136:
			return "RES 1,B"
		case 149:
			return "RES 2,L"
		case 158:
			return "RES 3,(HL)"
		case 25:
			return "RR C"
		case 154:
			return "RES 3,D"
		case 6:
			return "RLC (HL)"
		case 49:
			return "SWAP C"
		case 73:
			return "BIT 1,C"
		case 204:
			return "SET 1,H"
		case 51:
			return "SWAP E"
		case 88:
			return "BIT 3,B"
		case 150:
			return "RES 2,(HL)"
		case 160:
			return "RES 4,B"
		case 161:
			return "RES 4,C"
		case 199:
			return "SET 0,A"
		case 214:
			return "SET 2,(HL)"
		case 4:
			return "RLC H"
		case 55:
			return "SWAP A"
		case 67:
			return "BIT 0,E"
		case 40:
			return "SRA B"
		case 144:
			return "RES 2,B"
		case 206:
			return "SET 1,(HL)"
		case 129:
			return "RES 0,C"
		case 167:
			return "RES 4,A"
		case 212:
			return "SET 2,H"
		case 224:
			return "SET 4,B"
		case 3:
			return "RLC E"
		case 10:
			return "RRC D"
		case 54:
			return "SWAP (HL)"
		case 165:
			return "RES 4,L"
		case 0:
			return "RLC B"
		case 18:
			return "RL D"
		case 48:
			return "SWAP B"
		case 141:
			return "RES 1,L"
		case 152:
			return "RES 3,B"
		case 20:
			return "RL H"
		case 29:
			return "RR L"
		case 37:
			return "SLA L"
		case 159:
			return "RES 3,A"
		case 202:
			return "SET 1,D"
		case 53:
			return "SWAP L"
		case 102:
			return "BIT 4,(HL)"
		case 109:
			return "BIT 5,L"
		case 146:
			return "RES 2,D"
		case 168:
			return "RES 5,B"
		case 208:
			return "SET 2,B"
		case 31:
			return "RR A"
		case 93:
			return "BIT 3,L"
		case 128:
			return "RES 0,B"
		case 135:
			return "RES 0,A"
		case 151:
			return "RES 2,A"
		case 211:
			return "SET 2,E"
		case 229:
			return "SET 4,L"
		case 240:
			return "SET 6,B"
		case 64:
			return "BIT 0,B"
		case 130:
			return "RES 0,D"
		case 192:
			return "SET 0,B"
		case 27:
			return "RR E"
		case 122:
			return "BIT 7,D"
		case 185:
			return "RES 7,C"
		case 217:
			return "SET 3,C"
		case 80:
			return "BIT 2,B"
		case 178:
			return "RES 6,D"
		case 32:
			return "SLA B"
		case 39:
			return "SLA A"
		case 52:
			return "SWAP H"
		case 89:
			return "BIT 3,C"
		case 125:
			return "BIT 7,L"
		case 107:
			return "BIT 5,E"
		case 84:
			return "BIT 2,H"
		case 172:
			return "RES 5,H"
		case 190:
			return "RES 7,(HL)"
		case 5:
			return "RLC L"
		case 132:
			return "RES 0,H"
		case 189:
			return "RES 7,L"
		case 108:
			return "BIT 5,H"
		case 45:
			return "SRA L"
		case 60:
			return "SRL H"
		case 112:
			return "BIT 6,B"
		case 186:
			return "RES 7,D"
		case 90:
			return "BIT 3,D"
		case 138:
			return "RES 1,D"
		case 242:
			return "SET 6,D"
		case 35:
			return "SLA E"
		case 59:
			return "SRL E"
		case 155:
			return "RES 3,E"
		case 163:
			return "RES 4,E"
		case 166:
			return "RES 4,(HL)"
		case 197:
			return "SET 0,L"
		case 207:
			return "SET 1,A"
		case 230:
			return "SET 4,(HL)"
		case 21:
			return "RL L"
		case 62:
			return "SRL (HL)"
		case 179:
			return "RES 6,E"
		case 182:
			return "RES 6,(HL)"
		case 191:
			return "RES 7,A"
		case 124:
			return "BIT 7,H"
		case 145:
			return "RES 2,C"
		case 22:
			return "RL (HL)"
		case 157:
			return "RES 3,L"
		case 169:
			return "RES 5,C"
		case 231:
			return "SET 4,A"
		case 63:
			return "SRL A"
		case 220:
			return "SET 3,H"
		case 255:
			return "SET 7,A"
		case 15:
			return "RRC A"
		case 68:
			return "BIT 0,H"
		case 103:
			return "BIT 4,A"
		case 127:
			return "BIT 7,A"
		case 254:
			return "SET 7,(HL)"
		case 114:
			return "BIT 6,D"
		case 180:
			return "RES 6,H"
		case 234:
			return "SET 5,D"
		case 110:
			return "BIT 5,(HL)"
		case 153:
			return "RES 3,C"
		case 222:
			return "SET 3,(HL)"
		case 23:
			return "RL A"
		case 106:
			return "BIT 5,D"
		case 210:
			return "SET 2,D"
		case 221:
			return "SET 3,L"
		case 8:
			return "RRC B"
		case 115:
			return "BIT 6,E"
		case 117:
			return "BIT 6,L"
		case 200:
			return "SET 1,B"
		case 237:
			return "SET 5,L"
		case 105:
			return "BIT 5,C"
		case 2:
			return "RLC D"
		case 26:
			return "RR D"
		case 41:
			return "SRA C"
		case 66:
			return "BIT 0,D"
		case 123:
			return "BIT 7,E"
		case 147:
			return "RES 2,E"
		case 246:
			return "SET 6,(HL)"
		case 12:
			return "RRC H"
		case 83:
			return "BIT 2,E"
		case 97:
			return "BIT 4,C"
		case 100:
			return "BIT 4,H"
		case 174:
			return "RES 5,(HL)"
		case 238:
			return "SET 5,(HL)"
		case 243:
			return "SET 6,E"
		case 57:
			return "SRL C"
		case 170:
			return "RES 5,D"
		case 228:
			return "SET 4,H"
		case 70:
			return "BIT 0,(HL)"
		case 248:
			return "SET 7,B"
		case 13:
			return "RRC L"
		case 95:
			return "BIT 3,A"
		case 99:
			return "BIT 4,E"
		case 198:
			return "SET 0,(HL)"
		case 1:
			return "RLC C"
		case 46:
			return "SRA (HL)"
		case 69:
			return "BIT 0,L"
		case 85:
			return "BIT 2,L"
		case 78:
			return "BIT 1,(HL)"
		case 101:
			return "BIT 4,L"
		case 203:
			return "SET 1,E"
		case 223:
			return "SET 3,A"
		case 226:
			return "SET 4,D"
		case 227:
			return "SET 4,E"
		case 7:
			return "RLC A"
		case 42:
			return "SRA D"
		case 104:
			return "BIT 5,B"
		case 173:
			return "RES 5,L"
		case 216:
			return "SET 3,B"
		case 250:
			return "SET 7,D"
		case 253:
			return "SET 7,L"
		case 33:
			return "SLA C"
		case 184:
			return "RES 7,B"
		case 91:
			return "BIT 3,E"
		case 215:
			return "SET 2,A"
		case 241:
			return "SET 6,C"
		case 61:
			return "SRL L"
		case 116:
			return "BIT 6,H"
		case 239:
			return "SET 5,A"
		case 252:
			return "SET 7,H"
		case 79:
			return "BIT 1,A"
		case 137:
			return "RES 1,C"

		}
	}
	switch code {
	case 43:
		return "DEC HL"
	case 75:
		return "LD C,E"
	case 87:
		return "LD D,A"
	case 122:
		return "LD A,D"
	case 154:
		return "SBC A,D"
	case 161:
		return "AND A,C"
	case 163:
		return "AND A,E"
	case 213:
		return "PUSH DE"
	case 88:
		return "LD E,B"
	case 117:
		return "LD (HL),L"
	case 192:
		return "RET NZ"
	case 206:
		return "ADC A,n8"
	case 232:
		return "ADD SP,e8"
	case 255:
		return "RST $38"
	case 14:
		return "LD C,n8"
	case 34:
		return "LD (HL+),A"
	case 96:
		return "LD H,B"
	case 107:
		return "LD L,E"
	case 5:
		return "DEC B"
	case 10:
		return "LD A,(BC)"
	case 61:
		return "DEC A"
	case 91:
		return "LD E,E"
	case 92:
		return "LD E,H"
	case 169:
		return "XOR A,C"
	case 184:
		return "CP A,B"
	case 218:
		return "JP C,a16"
	case 231:
		return "RST $20"
	case 53:
		return "DEC (HL)"
	case 74:
		return "LD C,D"
	case 105:
		return "LD L,C"
	case 128:
		return "ADD A,B"
	case 173:
		return "XOR A,L"
	case 233:
		return "JP HL"
	case 49:
		return "LD SP,n16"
	case 37:
		return "DEC H"
	case 76:
		return "LD C,H"
	case 83:
		return "LD D,E"
	case 104:
		return "LD L,B"
	case 148:
		return "SUB A,H"
	case 228:
		return "ILLEGAL_E4"
	case 73:
		return "LD C,C"
	case 94:
		return "LD E,(HL)"
	case 186:
		return "CP A,D"
	case 215:
		return "RST $10"
	case 2:
		return "LD (BC),A"
	case 51:
		return "INC SP"
	case 84:
		return "LD D,H"
	case 143:
		return "ADC A,A"
	case 181:
		return "OR A,L"
	case 4:
		return "INC B"
	case 125:
		return "LD A,L"
	case 229:
		return "PUSH HL"
	case 230:
		return "AND A,n8"
	case 244:
		return "ILLEGAL_F4"
	case 50:
		return "LD (HL-),A"
	case 41:
		return "ADD HL,HL"
	case 118:
		return "HALT"
	case 131:
		return "ADD A,E"
	case 171:
		return "XOR A,E"
	case 187:
		return "CP A,E"
	case 198:
		return "ADD A,n8"
	case 18:
		return "LD (DE),A"
	case 139:
		return "ADC A,E"
	case 239:
		return "RST $28"
	case 45:
		return "DEC L"
	case 70:
		return "LD B,(HL)"
	case 100:
		return "LD H,H"
	case 164:
		return "AND A,H"
	case 175:
		return "XOR A,A"
	case 212:
		return "CALL NC,a16"
	case 7:
		return "RLCA"
	case 112:
		return "LD (HL),B"
	case 165:
		return "AND A,L"
	case 242:
		return "LDH A,(C)"
	case 82:
		return "LD D,D"
	case 195:
		return "JP a16"
	case 22:
		return "LD D,n8"
	case 23:
		return "RLA"
	case 160:
		return "AND A,B"
	case 162:
		return "AND A,D"
	case 251:
		return "EI"
	case 252:
		return "ILLEGAL_FC"
	case 135:
		return "ADD A,A"
	case 142:
		return "ADC A,(HL)"
	case 204:
		return "CALL Z,a16"
	case 38:
		return "LD H,n8"
	case 199:
		return "RST $00"
	case 226:
		return "LDH (C),A"
	case 24:
		return "JR e8"
	case 58:
		return "LD A,(HL-)"
	case 85:
		return "LD D,L"
	case 172:
		return "XOR A,H"
	case 113:
		return "LD (HL),C"
	case 249:
		return "LD SP,HL"
	case 95:
		return "LD E,A"
	case 183:
		return "OR A,A"
	case 201:
		return "RET"
	case 238:
		return "XOR A,n8"
	case 243:
		return "DI"
	case 56:
		return "JR C,e8"
	case 194:
		return "JP NZ,a16"
	case 197:
		return "PUSH BC"
	case 17:
		return "LD DE,n16"
	case 19:
		return "INC DE"
	case 44:
		return "INC L"
	case 46:
		return "LD L,n8"
	case 121:
		return "LD A,C"
	case 141:
		return "ADC A,L"
	case 151:
		return "SUB A,A"
	case 221:
		return "ILLEGAL_DD"
	case 81:
		return "LD D,C"
	case 101:
		return "LD H,L"
	case 177:
		return "OR A,C"
	case 234:
		return "LD (a16),A"
	case 136:
		return "ADC A,B"
	case 223:
		return "RST $18"
	case 32:
		return "JR NZ,e8"
	case 71:
		return "LD B,A"
	case 123:
		return "LD A,E"
	case 178:
		return "OR A,D"
	case 191:
		return "CP A,A"
	case 222:
		return "SBC A,n8"
	case 254:
		return "CP A,n8"
	case 29:
		return "DEC E"
	case 31:
		return "RRA"
	case 114:
		return "LD (HL),D"
	case 140:
		return "ADC A,H"
	case 190:
		return "CP A,(HL)"
	case 40:
		return "JR Z,e8"
	case 185:
		return "CP A,C"
	case 60:
		return "INC A"
	case 80:
		return "LD D,B"
	case 126:
		return "LD A,(HL)"
	case 200:
		return "RET Z"
	case 203:
		return "PREFIX"
	case 207:
		return "RST $08"
	case 216:
		return "RET C"
	case 8:
		return "LD (a16),SP"
	case 12:
		return "INC C"
	case 55:
		return "SCF"
	case 67:
		return "LD B,E"
	case 189:
		return "CP A,L"
	case 27:
		return "DEC DE"
	case 59:
		return "DEC SP"
	case 115:
		return "LD (HL),E"
	case 202:
		return "JP Z,a16"
	case 208:
		return "RET NC"
	case 248:
		return "LD HL,SP+,e8"
	case 106:
		return "LD L,D"
	case 110:
		return "LD L,(HL)"
	case 235:
		return "ILLEGAL_EB"
	case 241:
		return "POP AF"
	case 89:
		return "LD E,C"
	case 111:
		return "LD L,A"
	case 150:
		return "SUB A,(HL)"
	case 25:
		return "ADD HL,DE"
	case 47:
		return "CPL"
	case 124:
		return "LD A,H"
	case 3:
		return "INC BC"
	case 48:
		return "JR NC,e8"
	case 72:
		return "LD C,B"
	case 102:
		return "LD H,(HL)"
	case 155:
		return "SBC A,E"
	case 167:
		return "AND A,A"
	case 65:
		return "LD B,C"
	case 69:
		return "LD B,L"
	case 144:
		return "SUB A,B"
	case 166:
		return "AND A,(HL)"
	case 209:
		return "POP DE"
	case 217:
		return "RETI"
	case 77:
		return "LD C,L"
	case 134:
		return "ADD A,(HL)"
	case 156:
		return "SBC A,H"
	case 220:
		return "CALL C,a16"
	case 236:
		return "ILLEGAL_EC"
	case 42:
		return "LD A,(HL+)"
	case 0:
		return "NOP"
	case 9:
		return "ADD HL,BC"
	case 64:
		return "LD B,B"
	case 127:
		return "LD A,A"
	case 36:
		return "INC H"
	case 224:
		return "LDH (a8),A"
	case 240:
		return "LDH A,(a8)"
	case 63:
		return "CCF"
	case 103:
		return "LD H,A"
	case 176:
		return "OR A,B"
	case 219:
		return "ILLEGAL_DB"
	case 86:
		return "LD D,(HL)"
	case 253:
		return "ILLEGAL_FD"
	case 26:
		return "LD A,(DE)"
	case 13:
		return "DEC C"
	case 159:
		return "SBC A,A"
	case 20:
		return "INC D"
	case 54:
		return "LD (HL),n8"
	case 68:
		return "LD B,H"
	case 78:
		return "LD C,(HL)"
	case 119:
		return "LD (HL),A"
	case 130:
		return "ADD A,D"
	case 109:
		return "LD L,L"
	case 129:
		return "ADD A,C"
	case 179:
		return "OR A,E"
	case 11:
		return "DEC BC"
	case 108:
		return "LD L,H"
	case 180:
		return "OR A,H"
	case 182:
		return "OR A,(HL)"
	case 193:
		return "POP BC"
	case 157:
		return "SBC A,L"
	case 188:
		return "CP A,H"
	case 246:
		return "OR A,n8"
	case 21:
		return "DEC D"
	case 52:
		return "INC (HL)"
	case 247:
		return "RST $30"
	case 28:
		return "INC E"
	case 98:
		return "LD H,D"
	case 138:
		return "ADC A,D"
	case 170:
		return "XOR A,D"
	case 174:
		return "XOR A,(HL)"
	case 250:
		return "LD A,(a16)"
	case 116:
		return "LD (HL),H"
	case 120:
		return "LD A,B"
	case 132:
		return "ADD A,H"
	case 137:
		return "ADC A,C"
	case 147:
		return "SUB A,E"
	case 152:
		return "SBC A,B"
	case 30:
		return "LD E,n8"
	case 33:
		return "LD HL,n16"
	case 66:
		return "LD B,D"
	case 149:
		return "SUB A,L"
	case 225:
		return "POP HL"
	case 237:
		return "ILLEGAL_ED"
	case 133:
		return "ADD A,L"
	case 196:
		return "CALL NZ,a16"
	case 1:
		return "LD BC,n16"
	case 16:
		return "STOP"
	case 35:
		return "INC HL"
	case 57:
		return "ADD HL,SP"
	case 158:
		return "SBC A,(HL)"
	case 153:
		return "SBC A,C"
	case 210:
		return "JP NC,a16"
	case 211:
		return "ILLEGAL_D3"
	case 245:
		return "PUSH AF"
	case 6:
		return "LD B,n8"
	case 79:
		return "LD C,A"
	case 93:
		return "LD E,L"
	case 90:
		return "LD E,D"
	case 145:
		return "SUB A,C"
	case 39:
		return "DAA"
	case 97:
		return "LD H,C"
	case 146:
		return "SUB A,D"
	case 168:
		return "XOR A,B"
	case 205:
		return "CALL a16"
	case 99:
		return "LD H,E"
	case 214:
		return "SUB A,n8"
	case 227:
		return "ILLEGAL_E3"
	case 62:
		return "LD A,n8"

	default:
		panic(fmt.Sprintf("Unknown code for %d", code))
	}
}
