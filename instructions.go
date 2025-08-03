package gameboy

import "fmt"

type Instruction func(cpu *CPU)

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

// ADC A,C    code=0x89
func ADC_89(cpu *CPU) {
	panic("TODO ADC_89")
}

// CP A,H    code=0xbc
func CP_BC(cpu *CPU) {
	_, flags := sub(cpu.A, cpu.H)
	cpu.F = flags
	cpu.cycles += 4
}

// RST $18    code=0xdf
func RST_DF(cpu *CPU) {
	n := uint8(0x18)
	cpu.PushStack(cpu.PC)
	cpu.PC = concatU16(0x00, n)
}

// INC C    code=0x0c
func INC_0C(cpu *CPU) {
	res, flags := add(cpu.C, 0x01)
	cpu.F = flags
	cpu.C = res
	cpu.IncProgramCounter("inc")
	cpu.cycles += 4
}

// DEC DE    code=0x1b
func DEC_1B(cpu *CPU) {
	res, flags := sub(cpu.DE(), 0x01)
	cpu.F = flags
	cpu.D, cpu.E = split(res)
	cpu.IncProgramCounter("dec")
	cpu.cycles += 8
}

// LD D,L    code=0x55
func LD_55(cpu *CPU) {

	data := cpu.L

	cpu.D = data

	cpu.IncProgramCounter("ld")
	cpu.cycles += 4

}

// LD A,B    code=0x78
func LD_78(cpu *CPU) {

	data := cpu.B

	cpu.A = data

	cpu.IncProgramCounter("ld")
	cpu.cycles += 4

}

// LD (a16),A    code=0xea
func LD_EA(cpu *CPU) {

	data := cpu.A

	cpu.WriteMemory(cpu.readU16(cpu.PC), data)

	cpu.IncProgramCounter("ld")
	cpu.cycles += 16

}

// LD HL,SP+,e8    code=0xf8
func LD_F8(cpu *CPU) {

	e := cpu.readI8(cpu.PC)
	res, flags := add(cpu.SP, e)
	cpu.H, cpu.L = split(res)
	cpu.F = flags
	cpu.IncProgramCounter("ld")
	cpu.cycles += 12

}

// SCF     code=0x37
func SCF_37(cpu *CPU) {
	panic("TODO SCF_37")
}

// LD E,D    code=0x5a
func LD_5A(cpu *CPU) {

	data := cpu.D

	cpu.E = data

	cpu.IncProgramCounter("ld")
	cpu.cycles += 4

}

// ADC A,L    code=0x8d
func ADC_8D(cpu *CPU) {
	panic("TODO ADC_8D")
}

// SUB A,E    code=0x93
func SUB_93(cpu *CPU) {
	res, flags := sub(cpu.A, cpu.E)
	cpu.A = res
	cpu.F = flags
	cpu.cycles += 4
}

// ILLEGAL_D3     code=0xd3
func ILLEGAL_D3_D3(cpu *CPU) {
	panic("TODO ILLEGAL_D3_D3")
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

// STOP n8    code=0x10
func STOP_10(cpu *CPU) {
	cpu.err = ErrNoMoreInstructions
	cpu.cycles += 4
}

// LD C,A    code=0x4f
func LD_4F(cpu *CPU) {

	data := cpu.A

	cpu.C = data

	cpu.IncProgramCounter("ld")
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

// PUSH AF    code=0xf5
func PUSH_F5(cpu *CPU) {
	cpu.PushStack(cpu.AF())
	cpu.cycles += 16
}

// LD C,E    code=0x4b
func LD_4B(cpu *CPU) {

	data := cpu.E

	cpu.C = data

	cpu.IncProgramCounter("ld")
	cpu.cycles += 4

}

// LD E,A    code=0x5f
func LD_5F(cpu *CPU) {

	data := cpu.A

	cpu.E = data

	cpu.IncProgramCounter("ld")
	cpu.cycles += 4

}

// LD A,D    code=0x7a
func LD_7A(cpu *CPU) {

	data := cpu.D

	cpu.A = data

	cpu.IncProgramCounter("ld")
	cpu.cycles += 4

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

// LD SP,HL    code=0xf9
func LD_F9(cpu *CPU) {

	data := cpu.HL()

	cpu.SP = data

	cpu.IncProgramCounter("ld")
	cpu.cycles += 8

}

// INC B    code=0x04
func INC_04(cpu *CPU) {
	res, flags := add(cpu.B, 0x01)
	cpu.F = flags
	cpu.B = res
	cpu.IncProgramCounter("inc")
	cpu.cycles += 4
}

// LD D,n8    code=0x16
func LD_16(cpu *CPU) {

	data := cpu.readU8(cpu.PC)

	cpu.D = data

	cpu.IncProgramCounter("ld")
	cpu.cycles += 8

}

// LD (HL+),A    code=0x22
func LD_22(cpu *CPU) {

	data := cpu.A

	cpu.WriteMemory(cpu.HL(), data)

	incr, flags := add(cpu.HL(), 0x01)
	cpu.H, cpu.L = split(incr)
	cpu.F = flags

	cpu.IncProgramCounter("ld")
	cpu.cycles += 8

}

// RETI     code=0xd9
func RETI_D9(cpu *CPU) {
	panic("TODO RETI_D9")
}

// ILLEGAL_DB     code=0xdb
func ILLEGAL_DB_DB(cpu *CPU) {
	panic("TODO ILLEGAL_DB_DB")
}

// LD BC,n16    code=0x01
func LD_01(cpu *CPU) {

	data := cpu.readU16(cpu.PC)

	cpu.B, cpu.C = split(data)

	cpu.IncProgramCounter("ld")
	cpu.cycles += 12

}

// SBC A,(HL)    code=0x9e
func SBC_9E(cpu *CPU) {
	panic("TODO SBC_9E")
}

// INC DE    code=0x13
func INC_13(cpu *CPU) {
	res, flags := add(cpu.DE(), 0x01)
	cpu.F = flags
	cpu.D, cpu.E = split(res)
	cpu.IncProgramCounter("inc")
	cpu.cycles += 8
}

// LD H,A    code=0x67
func LD_67(cpu *CPU) {

	data := cpu.A

	cpu.H = data

	cpu.IncProgramCounter("ld")
	cpu.cycles += 4

}

// JP C,a16    code=0xda
func JP_DA(cpu *CPU) {
	panic("TODO JP_DA")
}

// LD B,D    code=0x42
func LD_42(cpu *CPU) {

	data := cpu.D

	cpu.B = data

	cpu.IncProgramCounter("ld")
	cpu.cycles += 4

}

// LD D,(HL)    code=0x56
func LD_56(cpu *CPU) {

	data := cpu.loadU8(cpu.HL())

	cpu.D = data

	cpu.IncProgramCounter("ld")
	cpu.cycles += 8

}

// LD E,C    code=0x59
func LD_59(cpu *CPU) {

	data := cpu.C

	cpu.E = data

	cpu.IncProgramCounter("ld")
	cpu.cycles += 4

}

// HALT     code=0x76
func HALT_76(cpu *CPU) {
	panic("TODO HALT_76")
}

// AND A,C    code=0xa1
func AND_A1(cpu *CPU) {
	panic("TODO AND_A1")
}

// JP NC,a16    code=0xd2
func JP_D2(cpu *CPU) {
	panic("TODO JP_D2")
}

// JP HL    code=0xe9
func JP_E9(cpu *CPU) {
	panic("TODO JP_E9")
}

// DEC C    code=0x0d
func DEC_0D(cpu *CPU) {
	res, flags := sub(cpu.C, 0x01)
	cpu.F = flags
	cpu.C = res
	cpu.IncProgramCounter("dec")
	cpu.cycles += 4
}

// LD H,C    code=0x61
func LD_61(cpu *CPU) {

	data := cpu.C

	cpu.H = data

	cpu.IncProgramCounter("ld")
	cpu.cycles += 4

}

// LD A,H    code=0x7c
func LD_7C(cpu *CPU) {

	data := cpu.H

	cpu.A = data

	cpu.IncProgramCounter("ld")
	cpu.cycles += 4

}

// SUB A,D    code=0x92
func SUB_92(cpu *CPU) {
	res, flags := sub(cpu.A, cpu.D)
	cpu.A = res
	cpu.F = flags
	cpu.cycles += 4
}

// AND A,B    code=0xa0
func AND_A0(cpu *CPU) {
	panic("TODO AND_A0")
}

// LD E,n8    code=0x1e
func LD_1E(cpu *CPU) {

	data := cpu.readU8(cpu.PC)

	cpu.E = data

	cpu.IncProgramCounter("ld")
	cpu.cycles += 8

}

// CP A,n8    code=0xfe
func CP_FE(cpu *CPU) {
	_, flags := sub(cpu.A, cpu.readU8(cpu.PC))
	cpu.F = flags
	cpu.cycles += 8
}

// DEC A    code=0x3d
func DEC_3D(cpu *CPU) {
	res, flags := sub(cpu.A, 0x01)
	cpu.F = flags
	cpu.A = res
	cpu.IncProgramCounter("dec")
	cpu.cycles += 4
}

// LD A,n8    code=0x3e
func LD_3E(cpu *CPU) {

	data := cpu.readU8(cpu.PC)

	cpu.A = data

	cpu.IncProgramCounter("ld")
	cpu.cycles += 8

}

// LD L,D    code=0x6a
func LD_6A(cpu *CPU) {

	data := cpu.D

	cpu.L = data

	cpu.IncProgramCounter("ld")
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

// LDH (a8),A    code=0xe0
func LDH_E0(cpu *CPU) {
	panic("TODO LDH_E0")
}

// ILLEGAL_F4     code=0xf4
func ILLEGAL_F4_F4(cpu *CPU) {
	panic("TODO ILLEGAL_F4_F4")
}

// LD B,C    code=0x41
func LD_41(cpu *CPU) {

	data := cpu.C

	cpu.B = data

	cpu.IncProgramCounter("ld")
	cpu.cycles += 4

}

// LD D,C    code=0x51
func LD_51(cpu *CPU) {

	data := cpu.C

	cpu.D = data

	cpu.IncProgramCounter("ld")
	cpu.cycles += 4

}

// LD (HL),L    code=0x75
func LD_75(cpu *CPU) {

	data := cpu.L

	cpu.WriteMemory(cpu.HL(), data)

	cpu.IncProgramCounter("ld")
	cpu.cycles += 8

}

// PREFIX     code=0xcb
func PREFIX_CB(cpu *CPU) {
	panic("TODO PREFIX_CB")
}

// LD L,H    code=0x6c
func LD_6C(cpu *CPU) {

	data := cpu.H

	cpu.L = data

	cpu.IncProgramCounter("ld")
	cpu.cycles += 4

}

// LD (DE),A    code=0x12
func LD_12(cpu *CPU) {

	data := cpu.A

	cpu.WriteMemory(cpu.DE(), data)

	cpu.IncProgramCounter("ld")
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

// LDH A,(C)    code=0xf2
func LDH_F2(cpu *CPU) {
	panic("TODO LDH_F2")
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

// LD A,(HL+)    code=0x2a
func LD_2A(cpu *CPU) {

	data := cpu.loadU8(cpu.HL())

	cpu.A = data

	cpu.IncProgramCounter("ld")
	cpu.cycles += 8

}

// LD L,L    code=0x6d
func LD_6D(cpu *CPU) {

	data := cpu.L

	cpu.L = data

	cpu.IncProgramCounter("ld")
	cpu.cycles += 4

}

// RLCA     code=0x07
func RLCA_07(cpu *CPU) {
	panic("TODO RLCA_07")
}

// ILLEGAL_FC     code=0xfc
func ILLEGAL_FC_FC(cpu *CPU) {
	panic("TODO ILLEGAL_FC_FC")
}

// LD B,H    code=0x44
func LD_44(cpu *CPU) {

	data := cpu.H

	cpu.B = data

	cpu.IncProgramCounter("ld")
	cpu.cycles += 4

}

// LD C,B    code=0x48
func LD_48(cpu *CPU) {

	data := cpu.B

	cpu.C = data

	cpu.IncProgramCounter("ld")
	cpu.cycles += 4

}

// LD (HL),A    code=0x77
func LD_77(cpu *CPU) {

	data := cpu.A

	cpu.WriteMemory(cpu.HL(), data)

	cpu.IncProgramCounter("ld")
	cpu.cycles += 8

}

// AND A,E    code=0xa3
func AND_A3(cpu *CPU) {
	panic("TODO AND_A3")
}

// LD C,C    code=0x49
func LD_49(cpu *CPU) {

	data := cpu.C

	cpu.C = data

	cpu.IncProgramCounter("ld")
	cpu.cycles += 4

}

// LD D,A    code=0x57
func LD_57(cpu *CPU) {

	data := cpu.A

	cpu.D = data

	cpu.IncProgramCounter("ld")
	cpu.cycles += 4

}

// LD A,C    code=0x79
func LD_79(cpu *CPU) {

	data := cpu.C

	cpu.A = data

	cpu.IncProgramCounter("ld")
	cpu.cycles += 4

}

// LD C,n8    code=0x0e
func LD_0E(cpu *CPU) {

	data := cpu.readU8(cpu.PC)

	cpu.C = data

	cpu.IncProgramCounter("ld")
	cpu.cycles += 8

}

// JR NZ,e8    code=0x20
func JR_20(cpu *CPU) {
	panic("TODO JR_20")
}

// ADC A,B    code=0x88
func ADC_88(cpu *CPU) {
	panic("TODO ADC_88")
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

// JP Z,a16    code=0xca
func JP_CA(cpu *CPU) {
	panic("TODO JP_CA")
}

// LD A,E    code=0x7b
func LD_7B(cpu *CPU) {

	data := cpu.E

	cpu.A = data

	cpu.IncProgramCounter("ld")
	cpu.cycles += 4

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

// LD D,B    code=0x50
func LD_50(cpu *CPU) {

	data := cpu.B

	cpu.D = data

	cpu.IncProgramCounter("ld")
	cpu.cycles += 4

}

// LD L,C    code=0x69
func LD_69(cpu *CPU) {

	data := cpu.C

	cpu.L = data

	cpu.IncProgramCounter("ld")
	cpu.cycles += 4

}

// RST $10    code=0xd7
func RST_D7(cpu *CPU) {
	n := uint8(0x10)
	cpu.PushStack(cpu.PC)
	cpu.PC = concatU16(0x00, n)
}

// DI     code=0xf3
func DI_F3(cpu *CPU) {
	panic("TODO DI_F3")
}

// LD (a16),SP    code=0x08
func LD_08(cpu *CPU) {

	data := cpu.SP

	cpu.WriteMemory(cpu.readU16(cpu.PC), data)

	cpu.IncProgramCounter("ld")
	cpu.cycles += 20

}

// SUB A,C    code=0x91
func SUB_91(cpu *CPU) {
	res, flags := sub(cpu.A, cpu.C)
	cpu.A = res
	cpu.F = flags
	cpu.cycles += 4
}

// SUB A,n8    code=0xd6
func SUB_D6(cpu *CPU) {
	res, flags := sub(cpu.A, cpu.readU8(cpu.PC))
	cpu.A = res
	cpu.F = flags
	cpu.cycles += 8
}

// SBC A,n8    code=0xde
func SBC_DE(cpu *CPU) {
	panic("TODO SBC_DE")
}

// INC H    code=0x24
func INC_24(cpu *CPU) {
	res, flags := add(cpu.H, 0x01)
	cpu.F = flags
	cpu.H = res
	cpu.IncProgramCounter("inc")
	cpu.cycles += 4
}

// LD (HL),n8    code=0x36
func LD_36(cpu *CPU) {

	data := cpu.readU8(cpu.PC)

	cpu.WriteMemory(cpu.HL(), data)

	cpu.IncProgramCounter("ld")
	cpu.cycles += 12

}

// LD E,L    code=0x5d
func LD_5D(cpu *CPU) {

	data := cpu.L

	cpu.E = data

	cpu.IncProgramCounter("ld")
	cpu.cycles += 4

}

// LD A,(HL)    code=0x7e
func LD_7E(cpu *CPU) {

	data := cpu.loadU8(cpu.HL())

	cpu.A = data

	cpu.IncProgramCounter("ld")
	cpu.cycles += 8

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

// CP A,D    code=0xba
func CP_BA(cpu *CPU) {
	_, flags := sub(cpu.A, cpu.D)
	cpu.F = flags
	cpu.cycles += 4
}

// DEC SP    code=0x3b
func DEC_3B(cpu *CPU) {
	res, flags := sub(cpu.SP, 0x01)
	cpu.F = flags
	cpu.SP = res
	cpu.IncProgramCounter("dec")
	cpu.cycles += 8
}

// LD C,L    code=0x4d
func LD_4D(cpu *CPU) {

	data := cpu.L

	cpu.C = data

	cpu.IncProgramCounter("ld")
	cpu.cycles += 4

}

// LD H,D    code=0x62
func LD_62(cpu *CPU) {

	data := cpu.D

	cpu.H = data

	cpu.IncProgramCounter("ld")
	cpu.cycles += 4

}

// SBC A,B    code=0x98
func SBC_98(cpu *CPU) {
	panic("TODO SBC_98")
}

// ILLEGAL_FD     code=0xfd
func ILLEGAL_FD_FD(cpu *CPU) {
	panic("TODO ILLEGAL_FD_FD")
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

// ADD HL,SP    code=0x39
func ADD_39(cpu *CPU) {
	lhs := cpu.HL()
	rhs := cpu.SP
	res, flags := add(lhs, rhs)
	cpu.H, cpu.L = split(res)
	cpu.F = flags
	cpu.cycles += 8
}

// CP A,A    code=0xbf
func CP_BF(cpu *CPU) {
	_, flags := sub(cpu.A, cpu.A)
	cpu.F = flags
	cpu.cycles += 4
}

// DAA     code=0x27
func DAA_27(cpu *CPU) {
	panic("TODO DAA_27")
}

// LD (HL),D    code=0x72
func LD_72(cpu *CPU) {

	data := cpu.D

	cpu.WriteMemory(cpu.HL(), data)

	cpu.IncProgramCounter("ld")
	cpu.cycles += 8

}

// SBC A,E    code=0x9b
func SBC_9B(cpu *CPU) {
	panic("TODO SBC_9B")
}

// INC L    code=0x2c
func INC_2C(cpu *CPU) {
	res, flags := add(cpu.L, 0x01)
	cpu.F = flags
	cpu.L = res
	cpu.IncProgramCounter("inc")
	cpu.cycles += 4
}

// POP HL    code=0xe1
func POP_E1(cpu *CPU) {
	value := cpu.PopStack()
	cpu.H, cpu.L = split(value)
	cpu.IncProgramCounter("pop")
	cpu.cycles += 12
}

// PUSH HL    code=0xe5
func PUSH_E5(cpu *CPU) {
	cpu.PushStack(cpu.HL())
	cpu.cycles += 16
}

// LD A,(a16)    code=0xfa
func LD_FA(cpu *CPU) {

	data := cpu.loadU8(cpu.readU16(cpu.PC))

	cpu.A = data

	cpu.IncProgramCounter("ld")
	cpu.cycles += 16

}

// DEC (HL)    code=0x35
func DEC_35(cpu *CPU) {
	res, flags := sub(cpu.loadU8(cpu.HL()), 0x01)
	cpu.F = flags
	cpu.WriteMemory(cpu.HL(), res)
	cpu.IncProgramCounter("dec")
	cpu.cycles += 12
}

// INC BC    code=0x03
func INC_03(cpu *CPU) {
	res, flags := add(cpu.BC(), 0x01)
	cpu.F = flags
	cpu.B, cpu.C = split(res)
	cpu.IncProgramCounter("inc")
	cpu.cycles += 8
}

// LD (HL),E    code=0x73
func LD_73(cpu *CPU) {

	data := cpu.E

	cpu.WriteMemory(cpu.HL(), data)

	cpu.IncProgramCounter("ld")
	cpu.cycles += 8

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

// AND A,n8    code=0xe6
func AND_E6(cpu *CPU) {
	panic("TODO AND_E6")
}

// ADC A,A    code=0x8f
func ADC_8F(cpu *CPU) {
	panic("TODO ADC_8F")
}

// CP A,L    code=0xbd
func CP_BD(cpu *CPU) {
	_, flags := sub(cpu.A, cpu.L)
	cpu.F = flags
	cpu.cycles += 4
}

// LDH (C),A    code=0xe2
func LDH_E2(cpu *CPU) {
	panic("TODO LDH_E2")
}

// LDH A,(a8)    code=0xf0
func LDH_F0(cpu *CPU) {
	panic("TODO LDH_F0")
}

// AND A,(HL)    code=0xa6
func AND_A6(cpu *CPU) {
	panic("TODO AND_A6")
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

// RET NC    code=0xd0
func RET_D0(cpu *CPU) {
	if !cpu.F.HasCarry() {
		cpu.PC = cpu.PopStack()
		cpu.cycles += 20
	} else {
		cpu.cycles += 8
	}
}

// LD A,(DE)    code=0x1a
func LD_1A(cpu *CPU) {

	data := cpu.loadU8(cpu.DE())

	cpu.A = data

	cpu.IncProgramCounter("ld")
	cpu.cycles += 8

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

// DEC D    code=0x15
func DEC_15(cpu *CPU) {
	res, flags := sub(cpu.D, 0x01)
	cpu.F = flags
	cpu.D = res
	cpu.IncProgramCounter("dec")
	cpu.cycles += 4
}

// JR e8    code=0x18
func JR_18(cpu *CPU) {
	panic("TODO JR_18")
}

// LD H,L    code=0x65
func LD_65(cpu *CPU) {

	data := cpu.L

	cpu.H = data

	cpu.IncProgramCounter("ld")
	cpu.cycles += 4

}

// LD L,B    code=0x68
func LD_68(cpu *CPU) {

	data := cpu.B

	cpu.L = data

	cpu.IncProgramCounter("ld")
	cpu.cycles += 4

}

// RST $30    code=0xf7
func RST_F7(cpu *CPU) {
	n := uint8(0x30)
	cpu.PushStack(cpu.PC)
	cpu.PC = concatU16(0x00, n)
}

// JR Z,e8    code=0x28
func JR_28(cpu *CPU) {
	panic("TODO JR_28")
}

// LD B,E    code=0x43
func LD_43(cpu *CPU) {

	data := cpu.E

	cpu.B = data

	cpu.IncProgramCounter("ld")
	cpu.cycles += 4

}

// LD E,E    code=0x5b
func LD_5B(cpu *CPU) {

	data := cpu.E

	cpu.E = data

	cpu.IncProgramCounter("ld")
	cpu.cycles += 4

}

// LD E,H    code=0x5c
func LD_5C(cpu *CPU) {

	data := cpu.H

	cpu.E = data

	cpu.IncProgramCounter("ld")
	cpu.cycles += 4

}

// SBC A,L    code=0x9d
func SBC_9D(cpu *CPU) {
	panic("TODO SBC_9D")
}

// AND A,A    code=0xa7
func AND_A7(cpu *CPU) {
	panic("TODO AND_A7")
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

// PUSH DE    code=0xd5
func PUSH_D5(cpu *CPU) {
	cpu.PushStack(cpu.DE())
	cpu.cycles += 16
}

// LD L,(HL)    code=0x6e
func LD_6E(cpu *CPU) {

	data := cpu.loadU8(cpu.HL())

	cpu.L = data

	cpu.IncProgramCounter("ld")
	cpu.cycles += 8

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

// SBC A,D    code=0x9a
func SBC_9A(cpu *CPU) {
	panic("TODO SBC_9A")
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

// CP A,C    code=0xb9
func CP_B9(cpu *CPU) {
	_, flags := sub(cpu.A, cpu.C)
	cpu.F = flags
	cpu.cycles += 4
}

// RLA     code=0x17
func RLA_17(cpu *CPU) {
	panic("TODO RLA_17")
}

// DEC HL    code=0x2b
func DEC_2B(cpu *CPU) {
	res, flags := sub(cpu.HL(), 0x01)
	cpu.F = flags
	cpu.H, cpu.L = split(res)
	cpu.IncProgramCounter("dec")
	cpu.cycles += 8
}

// SUB A,L    code=0x95
func SUB_95(cpu *CPU) {
	res, flags := sub(cpu.A, cpu.L)
	cpu.A = res
	cpu.F = flags
	cpu.cycles += 4
}

// NOP     code=0x00
func NOP_00(cpu *CPU) {
	panic("TODO NOP_00")
}

// AND A,L    code=0xa5
func AND_A5(cpu *CPU) {
	panic("TODO AND_A5")
}

// ILLEGAL_EB     code=0xeb
func ILLEGAL_EB_EB(cpu *CPU) {
	panic("TODO ILLEGAL_EB_EB")
}

// LD C,H    code=0x4c
func LD_4C(cpu *CPU) {

	data := cpu.H

	cpu.C = data

	cpu.IncProgramCounter("ld")
	cpu.cycles += 4

}

// CP A,B    code=0xb8
func CP_B8(cpu *CPU) {
	_, flags := sub(cpu.A, cpu.B)
	cpu.F = flags
	cpu.cycles += 4
}

// ILLEGAL_ED     code=0xed
func ILLEGAL_ED_ED(cpu *CPU) {
	panic("TODO ILLEGAL_ED_ED")
}

// LD L,n8    code=0x2e
func LD_2E(cpu *CPU) {

	data := cpu.readU8(cpu.PC)

	cpu.L = data

	cpu.IncProgramCounter("ld")
	cpu.cycles += 8

}

// LD D,H    code=0x54
func LD_54(cpu *CPU) {

	data := cpu.H

	cpu.D = data

	cpu.IncProgramCounter("ld")
	cpu.cycles += 4

}

// AND A,D    code=0xa2
func AND_A2(cpu *CPU) {
	panic("TODO AND_A2")
}

// DEC L    code=0x2d
func DEC_2D(cpu *CPU) {
	res, flags := sub(cpu.L, 0x01)
	cpu.F = flags
	cpu.L = res
	cpu.IncProgramCounter("dec")
	cpu.cycles += 4
}

// LD H,(HL)    code=0x66
func LD_66(cpu *CPU) {

	data := cpu.loadU8(cpu.HL())

	cpu.H = data

	cpu.IncProgramCounter("ld")
	cpu.cycles += 8

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

// RST $28    code=0xef
func RST_EF(cpu *CPU) {
	n := uint8(0x28)
	cpu.PushStack(cpu.PC)
	cpu.PC = concatU16(0x00, n)
}

// RST $38    code=0xff
func RST_FF(cpu *CPU) {
	n := uint8(0x38)
	cpu.PushStack(cpu.PC)
	cpu.PC = concatU16(0x00, n)
}

// DEC BC    code=0x0b
func DEC_0B(cpu *CPU) {
	res, flags := sub(cpu.BC(), 0x01)
	cpu.F = flags
	cpu.B, cpu.C = split(res)
	cpu.IncProgramCounter("dec")
	cpu.cycles += 8
}

// INC E    code=0x1c
func INC_1C(cpu *CPU) {
	res, flags := add(cpu.E, 0x01)
	cpu.F = flags
	cpu.E = res
	cpu.IncProgramCounter("inc")
	cpu.cycles += 4
}

// ADC A,D    code=0x8a
func ADC_8A(cpu *CPU) {
	panic("TODO ADC_8A")
}

// SBC A,A    code=0x9f
func SBC_9F(cpu *CPU) {
	panic("TODO SBC_9F")
}

// ADC A,n8    code=0xce
func ADC_CE(cpu *CPU) {
	panic("TODO ADC_CE")
}

// RST $20    code=0xe7
func RST_E7(cpu *CPU) {
	n := uint8(0x20)
	cpu.PushStack(cpu.PC)
	cpu.PC = concatU16(0x00, n)
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

// CCF     code=0x3f
func CCF_3F(cpu *CPU) {
	panic("TODO CCF_3F")
}

// DEC H    code=0x25
func DEC_25(cpu *CPU) {
	res, flags := sub(cpu.H, 0x01)
	cpu.F = flags
	cpu.H = res
	cpu.IncProgramCounter("dec")
	cpu.cycles += 4
}

// LD C,D    code=0x4a
func LD_4A(cpu *CPU) {

	data := cpu.D

	cpu.C = data

	cpu.IncProgramCounter("ld")
	cpu.cycles += 4

}

// LD C,(HL)    code=0x4e
func LD_4E(cpu *CPU) {

	data := cpu.loadU8(cpu.HL())

	cpu.C = data

	cpu.IncProgramCounter("ld")
	cpu.cycles += 8

}

// LD H,H    code=0x64
func LD_64(cpu *CPU) {

	data := cpu.H

	cpu.H = data

	cpu.IncProgramCounter("ld")
	cpu.cycles += 4

}

// SUB A,H    code=0x94
func SUB_94(cpu *CPU) {
	res, flags := sub(cpu.A, cpu.H)
	cpu.A = res
	cpu.F = flags
	cpu.cycles += 4
}

// SUB A,A    code=0x97
func SUB_97(cpu *CPU) {
	res, flags := sub(cpu.A, cpu.A)
	cpu.A = res
	cpu.F = flags
	cpu.cycles += 4
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

// DEC B    code=0x05
func DEC_05(cpu *CPU) {
	res, flags := sub(cpu.B, 0x01)
	cpu.F = flags
	cpu.B = res
	cpu.IncProgramCounter("dec")
	cpu.cycles += 4
}

// LD A,(BC)    code=0x0a
func LD_0A(cpu *CPU) {

	data := cpu.loadU8(cpu.BC())

	cpu.A = data

	cpu.IncProgramCounter("ld")
	cpu.cycles += 8

}

// LD SP,n16    code=0x31
func LD_31(cpu *CPU) {

	data := cpu.readU16(cpu.PC)

	cpu.SP = data

	cpu.IncProgramCounter("ld")
	cpu.cycles += 12

}

// INC SP    code=0x33
func INC_33(cpu *CPU) {
	res, flags := add(cpu.SP, 0x01)
	cpu.F = flags
	cpu.SP = res
	cpu.IncProgramCounter("inc")
	cpu.cycles += 8
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

// ILLEGAL_DD     code=0xdd
func ILLEGAL_DD_DD(cpu *CPU) {
	panic("TODO ILLEGAL_DD_DD")
}

// LD D,E    code=0x53
func LD_53(cpu *CPU) {

	data := cpu.E

	cpu.D = data

	cpu.IncProgramCounter("ld")
	cpu.cycles += 4

}

// ADC A,E    code=0x8b
func ADC_8B(cpu *CPU) {
	panic("TODO ADC_8B")
}

// ILLEGAL_E3     code=0xe3
func ILLEGAL_E3_E3(cpu *CPU) {
	panic("TODO ILLEGAL_E3_E3")
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

// RRA     code=0x1f
func RRA_1F(cpu *CPU) {
	panic("TODO RRA_1F")
}

// LD HL,n16    code=0x21
func LD_21(cpu *CPU) {

	data := cpu.readU16(cpu.PC)

	cpu.H, cpu.L = split(data)

	cpu.IncProgramCounter("ld")
	cpu.cycles += 12

}

// LD (HL-),A    code=0x32
func LD_32(cpu *CPU) {

	data := cpu.A

	cpu.WriteMemory(cpu.HL(), data)

	decr, flags := sub(cpu.HL(), 0x01)
	cpu.H, cpu.L = split(decr)
	cpu.F = flags

	cpu.IncProgramCounter("ld")
	cpu.cycles += 8

}

// LD B,A    code=0x47
func LD_47(cpu *CPU) {

	data := cpu.A

	cpu.B = data

	cpu.IncProgramCounter("ld")
	cpu.cycles += 4

}

// LD H,E    code=0x63
func LD_63(cpu *CPU) {

	data := cpu.E

	cpu.H = data

	cpu.IncProgramCounter("ld")
	cpu.cycles += 4

}

// LD A,A    code=0x7f
func LD_7F(cpu *CPU) {

	data := cpu.A

	cpu.A = data

	cpu.IncProgramCounter("ld")
	cpu.cycles += 4

}

// SBC A,H    code=0x9c
func SBC_9C(cpu *CPU) {
	panic("TODO SBC_9C")
}

// POP BC    code=0xc1
func POP_C1(cpu *CPU) {
	value := cpu.PopStack()
	cpu.B, cpu.C = split(value)
	cpu.IncProgramCounter("pop")
	cpu.cycles += 12
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

// JP a16    code=0xc3
func JP_C3(cpu *CPU) {
	panic("TODO JP_C3")
}

// LD DE,n16    code=0x11
func LD_11(cpu *CPU) {

	data := cpu.readU16(cpu.PC)

	cpu.D, cpu.E = split(data)

	cpu.IncProgramCounter("ld")
	cpu.cycles += 12

}

// DEC E    code=0x1d
func DEC_1D(cpu *CPU) {
	res, flags := sub(cpu.E, 0x01)
	cpu.F = flags
	cpu.E = res
	cpu.IncProgramCounter("dec")
	cpu.cycles += 4
}

// JR C,e8    code=0x38
func JR_38(cpu *CPU) {
	panic("TODO JR_38")
}

// LD E,(HL)    code=0x5e
func LD_5E(cpu *CPU) {

	data := cpu.loadU8(cpu.HL())

	cpu.E = data

	cpu.IncProgramCounter("ld")
	cpu.cycles += 8

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

// INC D    code=0x14
func INC_14(cpu *CPU) {
	res, flags := add(cpu.D, 0x01)
	cpu.F = flags
	cpu.D = res
	cpu.IncProgramCounter("inc")
	cpu.cycles += 4
}

// LD B,L    code=0x45
func LD_45(cpu *CPU) {

	data := cpu.L

	cpu.B = data

	cpu.IncProgramCounter("ld")
	cpu.cycles += 4

}

// LD L,E    code=0x6b
func LD_6B(cpu *CPU) {

	data := cpu.E

	cpu.L = data

	cpu.IncProgramCounter("ld")
	cpu.cycles += 4

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

// JR NC,e8    code=0x30
func JR_30(cpu *CPU) {
	panic("TODO JR_30")
}

// LD D,D    code=0x52
func LD_52(cpu *CPU) {

	data := cpu.D

	cpu.D = data

	cpu.IncProgramCounter("ld")
	cpu.cycles += 4

}

// ADC A,(HL)    code=0x8e
func ADC_8E(cpu *CPU) {
	panic("TODO ADC_8E")
}

// SUB A,(HL)    code=0x96
func SUB_96(cpu *CPU) {
	res, flags := sub(cpu.A, cpu.loadU8(cpu.HL()))
	cpu.A = res
	cpu.F = flags
	cpu.cycles += 8
}

// ILLEGAL_E4     code=0xe4
func ILLEGAL_E4_E4(cpu *CPU) {
	panic("TODO ILLEGAL_E4_E4")
}

// POP AF    code=0xf1
func POP_F1(cpu *CPU) {
	value := cpu.PopStack()
	msb, lsb := split(value)
	cpu.A, cpu.F = msb, FlagRegister(lsb)
	cpu.IncProgramCounter("pop")
	cpu.cycles += 12
}

// INC (HL)    code=0x34
func INC_34(cpu *CPU) {
	res, flags := add(cpu.loadU8(cpu.HL()), 0x01)
	cpu.F = flags
	cpu.WriteMemory(cpu.HL(), res)
	cpu.IncProgramCounter("inc")
	cpu.cycles += 12
}

// LD (HL),B    code=0x70
func LD_70(cpu *CPU) {

	data := cpu.B

	cpu.WriteMemory(cpu.HL(), data)

	cpu.IncProgramCounter("ld")
	cpu.cycles += 8

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

// ADD A,D    code=0x82
func ADD_82(cpu *CPU) {
	lhs := cpu.A
	rhs := cpu.D
	res, flags := add(lhs, rhs)
	cpu.A = res
	cpu.F = flags
	cpu.cycles += 4
}

// CP A,E    code=0xbb
func CP_BB(cpu *CPU) {
	_, flags := sub(cpu.A, cpu.E)
	cpu.F = flags
	cpu.cycles += 4
}

// RST $08    code=0xcf
func RST_CF(cpu *CPU) {
	n := uint8(0x8)
	cpu.PushStack(cpu.PC)
	cpu.PC = concatU16(0x00, n)
}

// LD A,(HL-)    code=0x3a
func LD_3A(cpu *CPU) {

	data := cpu.loadU8(cpu.HL())

	cpu.A = data

	cpu.IncProgramCounter("ld")
	cpu.cycles += 8

}

// INC A    code=0x3c
func INC_3C(cpu *CPU) {
	res, flags := add(cpu.A, 0x01)
	cpu.F = flags
	cpu.A = res
	cpu.IncProgramCounter("inc")
	cpu.cycles += 4
}

// LD B,B    code=0x40
func LD_40(cpu *CPU) {

	data := cpu.B

	cpu.B = data

	cpu.IncProgramCounter("ld")
	cpu.cycles += 4

}

// LD B,(HL)    code=0x46
func LD_46(cpu *CPU) {

	data := cpu.loadU8(cpu.HL())

	cpu.B = data

	cpu.IncProgramCounter("ld")
	cpu.cycles += 8

}

// LD (HL),H    code=0x74
func LD_74(cpu *CPU) {

	data := cpu.H

	cpu.WriteMemory(cpu.HL(), data)

	cpu.IncProgramCounter("ld")
	cpu.cycles += 8

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

// SUB A,B    code=0x90
func SUB_90(cpu *CPU) {
	res, flags := sub(cpu.A, cpu.B)
	cpu.A = res
	cpu.F = flags
	cpu.cycles += 4
}

// CP A,(HL)    code=0xbe
func CP_BE(cpu *CPU) {
	_, flags := sub(cpu.A, cpu.loadU8(cpu.HL()))
	cpu.F = flags
	cpu.cycles += 8
}

// INC HL    code=0x23
func INC_23(cpu *CPU) {
	res, flags := add(cpu.HL(), 0x01)
	cpu.F = flags
	cpu.H, cpu.L = split(res)
	cpu.IncProgramCounter("inc")
	cpu.cycles += 8
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

// ILLEGAL_EC     code=0xec
func ILLEGAL_EC_EC(cpu *CPU) {
	panic("TODO ILLEGAL_EC_EC")
}

// EI     code=0xfb
func EI_FB(cpu *CPU) {
	panic("TODO EI_FB")
}

// LD L,A    code=0x6f
func LD_6F(cpu *CPU) {

	data := cpu.A

	cpu.L = data

	cpu.IncProgramCounter("ld")
	cpu.cycles += 4

}

// SBC A,C    code=0x99
func SBC_99(cpu *CPU) {
	panic("TODO SBC_99")
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

// POP DE    code=0xd1
func POP_D1(cpu *CPU) {
	value := cpu.PopStack()
	cpu.D, cpu.E = split(value)
	cpu.IncProgramCounter("pop")
	cpu.cycles += 12
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

// LD (HL),C    code=0x71
func LD_71(cpu *CPU) {

	data := cpu.C

	cpu.WriteMemory(cpu.HL(), data)

	cpu.IncProgramCounter("ld")
	cpu.cycles += 8

}

// RST $00    code=0xc7
func RST_C7(cpu *CPU) {
	n := uint8(0x0)
	cpu.PushStack(cpu.PC)
	cpu.PC = concatU16(0x00, n)
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

// LD E,B    code=0x58
func LD_58(cpu *CPU) {

	data := cpu.B

	cpu.E = data

	cpu.IncProgramCounter("ld")
	cpu.cycles += 4

}

// LD H,B    code=0x60
func LD_60(cpu *CPU) {

	data := cpu.B

	cpu.H = data

	cpu.IncProgramCounter("ld")
	cpu.cycles += 4

}

// PUSH BC    code=0xc5
func PUSH_C5(cpu *CPU) {
	cpu.PushStack(cpu.BC())
	cpu.cycles += 16
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

// LD H,n8    code=0x26
func LD_26(cpu *CPU) {

	data := cpu.readU8(cpu.PC)

	cpu.H = data

	cpu.IncProgramCounter("ld")
	cpu.cycles += 8

}

// LD A,L    code=0x7d
func LD_7D(cpu *CPU) {

	data := cpu.L

	cpu.A = data

	cpu.IncProgramCounter("ld")
	cpu.cycles += 4

}

// LD B,n8    code=0x06
func LD_06(cpu *CPU) {

	data := cpu.readU8(cpu.PC)

	cpu.B = data

	cpu.IncProgramCounter("ld")
	cpu.cycles += 8

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

// JP NZ,a16    code=0xc2
func JP_C2(cpu *CPU) {
	panic("TODO JP_C2")
}

// LD (BC),A    code=0x02
func LD_02(cpu *CPU) {

	data := cpu.A

	cpu.WriteMemory(cpu.BC(), data)

	cpu.IncProgramCounter("ld")
	cpu.cycles += 8

}

// AND A,H    code=0xa4
func AND_A4(cpu *CPU) {
	panic("TODO AND_A4")
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

// ADC A,H    code=0x8c
func ADC_8C(cpu *CPU) {
	panic("TODO ADC_8C")
}

var ops = map[uint8]Instruction{
	0xaa: XOR_AA,
	0xcc: CALL_CC,
	0x89: ADC_89,
	0xbc: CP_BC,
	0xdf: RST_DF,
	0xc:  INC_0C,
	0x1b: DEC_1B,
	0x55: LD_55,
	0x78: LD_78,
	0xea: LD_EA,
	0xf8: LD_F8,
	0x37: SCF_37,
	0x5a: LD_5A,
	0x8d: ADC_8D,
	0x93: SUB_93,
	0xd3: ILLEGAL_D3_D3,
	0xdc: CALL_DC,
	0x10: STOP_10,
	0x4f: LD_4F,
	0xb4: OR_B4,
	0xf5: PUSH_F5,
	0x4b: LD_4B,
	0x5f: LD_5F,
	0x7a: LD_7A,
	0xc8: RET_C8,
	0xf9: LD_F9,
	0x4:  INC_04,
	0x16: LD_16,
	0x22: LD_22,
	0xd9: RETI_D9,
	0xdb: ILLEGAL_DB_DB,
	0x1:  LD_01,
	0x9e: SBC_9E,
	0x13: INC_13,
	0x67: LD_67,
	0xda: JP_DA,
	0x42: LD_42,
	0x56: LD_56,
	0x59: LD_59,
	0x76: HALT_76,
	0xa1: AND_A1,
	0xd2: JP_D2,
	0xe9: JP_E9,
	0xd:  DEC_0D,
	0x61: LD_61,
	0x7c: LD_7C,
	0x92: SUB_92,
	0xa0: AND_A0,
	0x1e: LD_1E,
	0xfe: CP_FE,
	0x3d: DEC_3D,
	0x3e: LD_3E,
	0x6a: LD_6A,
	0xd4: CALL_D4,
	0xe0: LDH_E0,
	0xf4: ILLEGAL_F4_F4,
	0x41: LD_41,
	0x51: LD_51,
	0x75: LD_75,
	0xcb: PREFIX_CB,
	0x6c: LD_6C,
	0x12: LD_12,
	0xc6: ADD_C6,
	0xf2: LDH_F2,
	0x9:  ADD_09,
	0x2a: LD_2A,
	0x6d: LD_6D,
	0x7:  RLCA_07,
	0xfc: ILLEGAL_FC_FC,
	0x44: LD_44,
	0x48: LD_48,
	0x77: LD_77,
	0xa3: AND_A3,
	0x49: LD_49,
	0x57: LD_57,
	0x79: LD_79,
	0xe:  LD_0E,
	0x20: JR_20,
	0x88: ADC_88,
	0xb1: OR_B1,
	0xca: JP_CA,
	0x7b: LD_7B,
	0xd8: RET_D8,
	0x50: LD_50,
	0x69: LD_69,
	0xd7: RST_D7,
	0xf3: DI_F3,
	0x8:  LD_08,
	0x91: SUB_91,
	0xd6: SUB_D6,
	0xde: SBC_DE,
	0x24: INC_24,
	0x36: LD_36,
	0x5d: LD_5D,
	0x7e: LD_7E,
	0xad: XOR_AD,
	0xba: CP_BA,
	0x3b: DEC_3B,
	0x4d: LD_4D,
	0x62: LD_62,
	0x98: SBC_98,
	0xfd: ILLEGAL_FD_FD,
	0x87: ADD_87,
	0xb6: OR_B6,
	0x39: ADD_39,
	0xbf: CP_BF,
	0x27: DAA_27,
	0x72: LD_72,
	0x9b: SBC_9B,
	0x2c: INC_2C,
	0xe1: POP_E1,
	0xe5: PUSH_E5,
	0xfa: LD_FA,
	0x35: DEC_35,
	0x3:  INC_03,
	0x73: LD_73,
	0xb0: OR_B0,
	0xe6: AND_E6,
	0x8f: ADC_8F,
	0xbd: CP_BD,
	0xe2: LDH_E2,
	0xf0: LDH_F0,
	0xa6: AND_A6,
	0xc4: CALL_C4,
	0xd0: RET_D0,
	0x1a: LD_1A,
	0x85: ADD_85,
	0x15: DEC_15,
	0x18: JR_18,
	0x65: LD_65,
	0x68: LD_68,
	0xf7: RST_F7,
	0x28: JR_28,
	0x43: LD_43,
	0x5b: LD_5B,
	0x5c: LD_5C,
	0x9d: SBC_9D,
	0xa7: AND_A7,
	0x84: ADD_84,
	0xd5: PUSH_D5,
	0x6e: LD_6E,
	0x86: ADD_86,
	0x9a: SBC_9A,
	0xc9: RET_C9,
	0xb9: CP_B9,
	0x17: RLA_17,
	0x2b: DEC_2B,
	0x95: SUB_95,
	0x0:  NOP_00,
	0xa5: AND_A5,
	0xeb: ILLEGAL_EB_EB,
	0x4c: LD_4C,
	0xb8: CP_B8,
	0xed: ILLEGAL_ED_ED,
	0x2e: LD_2E,
	0x54: LD_54,
	0xa2: AND_A2,
	0x2d: DEC_2D,
	0x66: LD_66,
	0xa8: XOR_A8,
	0xef: RST_EF,
	0xff: RST_FF,
	0xb:  DEC_0B,
	0x1c: INC_1C,
	0x8a: ADC_8A,
	0x9f: SBC_9F,
	0xce: ADC_CE,
	0xe7: RST_E7,
	0x19: ADD_19,
	0x2f: CPL_2F,
	0x3f: CCF_3F,
	0x25: DEC_25,
	0x4a: LD_4A,
	0x4e: LD_4E,
	0x64: LD_64,
	0x94: SUB_94,
	0x97: SUB_97,
	0xaf: XOR_AF,
	0xb2: OR_B2,
	0x5:  DEC_05,
	0xa:  LD_0A,
	0x31: LD_31,
	0x33: INC_33,
	0xab: XOR_AB,
	0xb5: OR_B5,
	0xdd: ILLEGAL_DD_DD,
	0x53: LD_53,
	0x8b: ADC_8B,
	0xe3: ILLEGAL_E3_E3,
	0xe8: ADD_E8,
	0x1f: RRA_1F,
	0x21: LD_21,
	0x32: LD_32,
	0x47: LD_47,
	0x63: LD_63,
	0x7f: LD_7F,
	0x9c: SBC_9C,
	0xc1: POP_C1,
	0xc0: RET_C0,
	0xc3: JP_C3,
	0x11: LD_11,
	0x1d: DEC_1D,
	0x38: JR_38,
	0x5e: LD_5E,
	0xb7: OR_B7,
	0x14: INC_14,
	0x45: LD_45,
	0x6b: LD_6B,
	0xee: XOR_EE,
	0x30: JR_30,
	0x52: LD_52,
	0x8e: ADC_8E,
	0x96: SUB_96,
	0xe4: ILLEGAL_E4_E4,
	0xf1: POP_F1,
	0x34: INC_34,
	0x70: LD_70,
	0x80: ADD_80,
	0x82: ADD_82,
	0xbb: CP_BB,
	0xcf: RST_CF,
	0x3a: LD_3A,
	0x3c: INC_3C,
	0x40: LD_40,
	0x46: LD_46,
	0x74: LD_74,
	0x83: ADD_83,
	0x90: SUB_90,
	0xbe: CP_BE,
	0x23: INC_23,
	0xae: XOR_AE,
	0xec: ILLEGAL_EC_EC,
	0xfb: EI_FB,
	0x6f: LD_6F,
	0x99: SBC_99,
	0xa9: XOR_A9,
	0xd1: POP_D1,
	0x29: ADD_29,
	0x71: LD_71,
	0xc7: RST_C7,
	0xcd: CALL_CD,
	0x58: LD_58,
	0x60: LD_60,
	0xc5: PUSH_C5,
	0xf6: OR_F6,
	0x26: LD_26,
	0x7d: LD_7D,
	0x6:  LD_06,
	0xac: XOR_AC,
	0xb3: OR_B3,
	0xc2: JP_C2,
	0x2:  LD_02,
	0xa4: AND_A4,
	0x81: ADD_81,
	0x8c: ADC_8C,
}

// returns code given a string. Useful during testing
func code(s string) uint8 {
	switch s {
	case "XOR A,D":
		return 0xAA
	case "CALL Z,a16":
		return 0xCC
	case "ADC A,C":
		return 0x89
	case "CP A,H":
		return 0xBC
	case "RST $18":
		return 0xDF
	case "INC C":
		return 0xC
	case "DEC DE":
		return 0x1B
	case "LD D,L":
		return 0x55
	case "LD A,B":
		return 0x78
	case "LD (a16),A":
		return 0xEA
	case "LD HL,SP+,e8":
		return 0xF8
	case "SCF":
		return 0x37
	case "LD E,D":
		return 0x5A
	case "ADC A,L":
		return 0x8D
	case "SUB A,E":
		return 0x93
	case "ILLEGAL_D3":
		return 0xD3
	case "CALL C,a16":
		return 0xDC
	case "STOP":
		return 0x10
	case "LD C,A":
		return 0x4F
	case "OR A,H":
		return 0xB4
	case "PUSH AF":
		return 0xF5
	case "LD C,E":
		return 0x4B
	case "LD E,A":
		return 0x5F
	case "LD A,D":
		return 0x7A
	case "RET Z":
		return 0xC8
	case "LD SP,HL":
		return 0xF9
	case "INC B":
		return 0x4
	case "LD D,n8":
		return 0x16
	case "LD (HL+),A":
		return 0x22
	case "RETI":
		return 0xD9
	case "ILLEGAL_DB":
		return 0xDB
	case "LD BC,n16":
		return 0x1
	case "SBC A,(HL)":
		return 0x9E
	case "INC DE":
		return 0x13
	case "LD H,A":
		return 0x67
	case "JP C,a16":
		return 0xDA
	case "LD B,D":
		return 0x42
	case "LD D,(HL)":
		return 0x56
	case "LD E,C":
		return 0x59
	case "HALT":
		return 0x76
	case "AND A,C":
		return 0xA1
	case "JP NC,a16":
		return 0xD2
	case "JP HL":
		return 0xE9
	case "DEC C":
		return 0xD
	case "LD H,C":
		return 0x61
	case "LD A,H":
		return 0x7C
	case "SUB A,D":
		return 0x92
	case "AND A,B":
		return 0xA0
	case "LD E,n8":
		return 0x1E
	case "CP A,n8":
		return 0xFE
	case "DEC A":
		return 0x3D
	case "LD A,n8":
		return 0x3E
	case "LD L,D":
		return 0x6A
	case "CALL NC,a16":
		return 0xD4
	case "LDH (a8),A":
		return 0xE0
	case "ILLEGAL_F4":
		return 0xF4
	case "LD B,C":
		return 0x41
	case "LD D,C":
		return 0x51
	case "LD (HL),L":
		return 0x75
	case "PREFIX":
		return 0xCB
	case "LD L,H":
		return 0x6C
	case "LD (DE),A":
		return 0x12
	case "ADD A,n8":
		return 0xC6
	case "LDH A,(C)":
		return 0xF2
	case "ADD HL,BC":
		return 0x9
	case "LD A,(HL+)":
		return 0x2A
	case "LD L,L":
		return 0x6D
	case "RLCA":
		return 0x7
	case "ILLEGAL_FC":
		return 0xFC
	case "LD B,H":
		return 0x44
	case "LD C,B":
		return 0x48
	case "LD (HL),A":
		return 0x77
	case "AND A,E":
		return 0xA3
	case "LD C,C":
		return 0x49
	case "LD D,A":
		return 0x57
	case "LD A,C":
		return 0x79
	case "LD C,n8":
		return 0xE
	case "JR NZ,e8":
		return 0x20
	case "ADC A,B":
		return 0x88
	case "OR A,C":
		return 0xB1
	case "JP Z,a16":
		return 0xCA
	case "LD A,E":
		return 0x7B
	case "RET C":
		return 0xD8
	case "LD D,B":
		return 0x50
	case "LD L,C":
		return 0x69
	case "RST $10":
		return 0xD7
	case "DI":
		return 0xF3
	case "LD (a16),SP":
		return 0x8
	case "SUB A,C":
		return 0x91
	case "SUB A,n8":
		return 0xD6
	case "SBC A,n8":
		return 0xDE
	case "INC H":
		return 0x24
	case "LD (HL),n8":
		return 0x36
	case "LD E,L":
		return 0x5D
	case "LD A,(HL)":
		return 0x7E
	case "XOR A,L":
		return 0xAD
	case "CP A,D":
		return 0xBA
	case "DEC SP":
		return 0x3B
	case "LD C,L":
		return 0x4D
	case "LD H,D":
		return 0x62
	case "SBC A,B":
		return 0x98
	case "ILLEGAL_FD":
		return 0xFD
	case "ADD A,A":
		return 0x87
	case "OR A,(HL)":
		return 0xB6
	case "ADD HL,SP":
		return 0x39
	case "CP A,A":
		return 0xBF
	case "DAA":
		return 0x27
	case "LD (HL),D":
		return 0x72
	case "SBC A,E":
		return 0x9B
	case "INC L":
		return 0x2C
	case "POP HL":
		return 0xE1
	case "PUSH HL":
		return 0xE5
	case "LD A,(a16)":
		return 0xFA
	case "DEC (HL)":
		return 0x35
	case "INC BC":
		return 0x3
	case "LD (HL),E":
		return 0x73
	case "OR A,B":
		return 0xB0
	case "AND A,n8":
		return 0xE6
	case "ADC A,A":
		return 0x8F
	case "CP A,L":
		return 0xBD
	case "LDH (C),A":
		return 0xE2
	case "LDH A,(a8)":
		return 0xF0
	case "AND A,(HL)":
		return 0xA6
	case "CALL NZ,a16":
		return 0xC4
	case "RET NC":
		return 0xD0
	case "LD A,(DE)":
		return 0x1A
	case "ADD A,L":
		return 0x85
	case "DEC D":
		return 0x15
	case "JR e8":
		return 0x18
	case "LD H,L":
		return 0x65
	case "LD L,B":
		return 0x68
	case "RST $30":
		return 0xF7
	case "JR Z,e8":
		return 0x28
	case "LD B,E":
		return 0x43
	case "LD E,E":
		return 0x5B
	case "LD E,H":
		return 0x5C
	case "SBC A,L":
		return 0x9D
	case "AND A,A":
		return 0xA7
	case "ADD A,H":
		return 0x84
	case "PUSH DE":
		return 0xD5
	case "LD L,(HL)":
		return 0x6E
	case "ADD A,(HL)":
		return 0x86
	case "SBC A,D":
		return 0x9A
	case "RET":
		return 0xC9
	case "CP A,C":
		return 0xB9
	case "RLA":
		return 0x17
	case "DEC HL":
		return 0x2B
	case "SUB A,L":
		return 0x95
	case "NOP":
		return 0x0
	case "AND A,L":
		return 0xA5
	case "ILLEGAL_EB":
		return 0xEB
	case "LD C,H":
		return 0x4C
	case "CP A,B":
		return 0xB8
	case "ILLEGAL_ED":
		return 0xED
	case "LD L,n8":
		return 0x2E
	case "LD D,H":
		return 0x54
	case "AND A,D":
		return 0xA2
	case "DEC L":
		return 0x2D
	case "LD H,(HL)":
		return 0x66
	case "XOR A,B":
		return 0xA8
	case "RST $28":
		return 0xEF
	case "RST $38":
		return 0xFF
	case "DEC BC":
		return 0xB
	case "INC E":
		return 0x1C
	case "ADC A,D":
		return 0x8A
	case "SBC A,A":
		return 0x9F
	case "ADC A,n8":
		return 0xCE
	case "RST $20":
		return 0xE7
	case "ADD HL,DE":
		return 0x19
	case "CPL":
		return 0x2F
	case "CCF":
		return 0x3F
	case "DEC H":
		return 0x25
	case "LD C,D":
		return 0x4A
	case "LD C,(HL)":
		return 0x4E
	case "LD H,H":
		return 0x64
	case "SUB A,H":
		return 0x94
	case "SUB A,A":
		return 0x97
	case "XOR A,A":
		return 0xAF
	case "OR A,D":
		return 0xB2
	case "DEC B":
		return 0x5
	case "LD A,(BC)":
		return 0xA
	case "LD SP,n16":
		return 0x31
	case "INC SP":
		return 0x33
	case "XOR A,E":
		return 0xAB
	case "OR A,L":
		return 0xB5
	case "ILLEGAL_DD":
		return 0xDD
	case "LD D,E":
		return 0x53
	case "ADC A,E":
		return 0x8B
	case "ILLEGAL_E3":
		return 0xE3
	case "ADD SP,e8":
		return 0xE8
	case "RRA":
		return 0x1F
	case "LD HL,n16":
		return 0x21
	case "LD (HL-),A":
		return 0x32
	case "LD B,A":
		return 0x47
	case "LD H,E":
		return 0x63
	case "LD A,A":
		return 0x7F
	case "SBC A,H":
		return 0x9C
	case "POP BC":
		return 0xC1
	case "RET NZ":
		return 0xC0
	case "JP a16":
		return 0xC3
	case "LD DE,n16":
		return 0x11
	case "DEC E":
		return 0x1D
	case "JR C,e8":
		return 0x38
	case "LD E,(HL)":
		return 0x5E
	case "OR A,A":
		return 0xB7
	case "INC D":
		return 0x14
	case "LD B,L":
		return 0x45
	case "LD L,E":
		return 0x6B
	case "XOR A,n8":
		return 0xEE
	case "JR NC,e8":
		return 0x30
	case "LD D,D":
		return 0x52
	case "ADC A,(HL)":
		return 0x8E
	case "SUB A,(HL)":
		return 0x96
	case "ILLEGAL_E4":
		return 0xE4
	case "POP AF":
		return 0xF1
	case "INC (HL)":
		return 0x34
	case "LD (HL),B":
		return 0x70
	case "ADD A,B":
		return 0x80
	case "ADD A,D":
		return 0x82
	case "CP A,E":
		return 0xBB
	case "RST $08":
		return 0xCF
	case "LD A,(HL-)":
		return 0x3A
	case "INC A":
		return 0x3C
	case "LD B,B":
		return 0x40
	case "LD B,(HL)":
		return 0x46
	case "LD (HL),H":
		return 0x74
	case "ADD A,E":
		return 0x83
	case "SUB A,B":
		return 0x90
	case "CP A,(HL)":
		return 0xBE
	case "INC HL":
		return 0x23
	case "XOR A,(HL)":
		return 0xAE
	case "ILLEGAL_EC":
		return 0xEC
	case "EI":
		return 0xFB
	case "LD L,A":
		return 0x6F
	case "SBC A,C":
		return 0x99
	case "XOR A,C":
		return 0xA9
	case "POP DE":
		return 0xD1
	case "ADD HL,HL":
		return 0x29
	case "LD (HL),C":
		return 0x71
	case "RST $00":
		return 0xC7
	case "CALL a16":
		return 0xCD
	case "LD E,B":
		return 0x58
	case "LD H,B":
		return 0x60
	case "PUSH BC":
		return 0xC5
	case "OR A,n8":
		return 0xF6
	case "LD H,n8":
		return 0x26
	case "LD A,L":
		return 0x7D
	case "LD B,n8":
		return 0x6
	case "XOR A,H":
		return 0xAC
	case "OR A,E":
		return 0xB3
	case "JP NZ,a16":
		return 0xC2
	case "LD (BC),A":
		return 0x2
	case "AND A,H":
		return 0xA4
	case "ADD A,C":
		return 0x81
	case "ADC A,H":
		return 0x8C

	default:
		panic(fmt.Sprintf("Unknown code for %q", s))
	}
}
func name(code uint8) string {
	switch code {
	case 170:
		return "XOR A,D"
	case 204:
		return "CALL Z,a16"
	case 137:
		return "ADC A,C"
	case 188:
		return "CP A,H"
	case 223:
		return "RST $18"
	case 12:
		return "INC C"
	case 27:
		return "DEC DE"
	case 85:
		return "LD D,L"
	case 120:
		return "LD A,B"
	case 234:
		return "LD (a16),A"
	case 248:
		return "LD HL,SP+,e8"
	case 55:
		return "SCF"
	case 90:
		return "LD E,D"
	case 141:
		return "ADC A,L"
	case 147:
		return "SUB A,E"
	case 211:
		return "ILLEGAL_D3"
	case 220:
		return "CALL C,a16"
	case 16:
		return "STOP"
	case 79:
		return "LD C,A"
	case 180:
		return "OR A,H"
	case 245:
		return "PUSH AF"
	case 75:
		return "LD C,E"
	case 95:
		return "LD E,A"
	case 122:
		return "LD A,D"
	case 200:
		return "RET Z"
	case 249:
		return "LD SP,HL"
	case 4:
		return "INC B"
	case 22:
		return "LD D,n8"
	case 34:
		return "LD (HL+),A"
	case 217:
		return "RETI"
	case 219:
		return "ILLEGAL_DB"
	case 1:
		return "LD BC,n16"
	case 158:
		return "SBC A,(HL)"
	case 19:
		return "INC DE"
	case 103:
		return "LD H,A"
	case 218:
		return "JP C,a16"
	case 66:
		return "LD B,D"
	case 86:
		return "LD D,(HL)"
	case 89:
		return "LD E,C"
	case 118:
		return "HALT"
	case 161:
		return "AND A,C"
	case 210:
		return "JP NC,a16"
	case 233:
		return "JP HL"
	case 13:
		return "DEC C"
	case 97:
		return "LD H,C"
	case 124:
		return "LD A,H"
	case 146:
		return "SUB A,D"
	case 160:
		return "AND A,B"
	case 30:
		return "LD E,n8"
	case 254:
		return "CP A,n8"
	case 61:
		return "DEC A"
	case 62:
		return "LD A,n8"
	case 106:
		return "LD L,D"
	case 212:
		return "CALL NC,a16"
	case 224:
		return "LDH (a8),A"
	case 244:
		return "ILLEGAL_F4"
	case 65:
		return "LD B,C"
	case 81:
		return "LD D,C"
	case 117:
		return "LD (HL),L"
	case 203:
		return "PREFIX"
	case 108:
		return "LD L,H"
	case 18:
		return "LD (DE),A"
	case 198:
		return "ADD A,n8"
	case 242:
		return "LDH A,(C)"
	case 9:
		return "ADD HL,BC"
	case 42:
		return "LD A,(HL+)"
	case 109:
		return "LD L,L"
	case 7:
		return "RLCA"
	case 252:
		return "ILLEGAL_FC"
	case 68:
		return "LD B,H"
	case 72:
		return "LD C,B"
	case 119:
		return "LD (HL),A"
	case 163:
		return "AND A,E"
	case 73:
		return "LD C,C"
	case 87:
		return "LD D,A"
	case 121:
		return "LD A,C"
	case 14:
		return "LD C,n8"
	case 32:
		return "JR NZ,e8"
	case 136:
		return "ADC A,B"
	case 177:
		return "OR A,C"
	case 202:
		return "JP Z,a16"
	case 123:
		return "LD A,E"
	case 216:
		return "RET C"
	case 80:
		return "LD D,B"
	case 105:
		return "LD L,C"
	case 215:
		return "RST $10"
	case 243:
		return "DI"
	case 8:
		return "LD (a16),SP"
	case 145:
		return "SUB A,C"
	case 214:
		return "SUB A,n8"
	case 222:
		return "SBC A,n8"
	case 36:
		return "INC H"
	case 54:
		return "LD (HL),n8"
	case 93:
		return "LD E,L"
	case 126:
		return "LD A,(HL)"
	case 173:
		return "XOR A,L"
	case 186:
		return "CP A,D"
	case 59:
		return "DEC SP"
	case 77:
		return "LD C,L"
	case 98:
		return "LD H,D"
	case 152:
		return "SBC A,B"
	case 253:
		return "ILLEGAL_FD"
	case 135:
		return "ADD A,A"
	case 182:
		return "OR A,(HL)"
	case 57:
		return "ADD HL,SP"
	case 191:
		return "CP A,A"
	case 39:
		return "DAA"
	case 114:
		return "LD (HL),D"
	case 155:
		return "SBC A,E"
	case 44:
		return "INC L"
	case 225:
		return "POP HL"
	case 229:
		return "PUSH HL"
	case 250:
		return "LD A,(a16)"
	case 53:
		return "DEC (HL)"
	case 3:
		return "INC BC"
	case 115:
		return "LD (HL),E"
	case 176:
		return "OR A,B"
	case 230:
		return "AND A,n8"
	case 143:
		return "ADC A,A"
	case 189:
		return "CP A,L"
	case 226:
		return "LDH (C),A"
	case 240:
		return "LDH A,(a8)"
	case 166:
		return "AND A,(HL)"
	case 196:
		return "CALL NZ,a16"
	case 208:
		return "RET NC"
	case 26:
		return "LD A,(DE)"
	case 133:
		return "ADD A,L"
	case 21:
		return "DEC D"
	case 24:
		return "JR e8"
	case 101:
		return "LD H,L"
	case 104:
		return "LD L,B"
	case 247:
		return "RST $30"
	case 40:
		return "JR Z,e8"
	case 67:
		return "LD B,E"
	case 91:
		return "LD E,E"
	case 92:
		return "LD E,H"
	case 157:
		return "SBC A,L"
	case 167:
		return "AND A,A"
	case 132:
		return "ADD A,H"
	case 213:
		return "PUSH DE"
	case 110:
		return "LD L,(HL)"
	case 134:
		return "ADD A,(HL)"
	case 154:
		return "SBC A,D"
	case 201:
		return "RET"
	case 185:
		return "CP A,C"
	case 23:
		return "RLA"
	case 43:
		return "DEC HL"
	case 149:
		return "SUB A,L"
	case 0:
		return "NOP"
	case 165:
		return "AND A,L"
	case 235:
		return "ILLEGAL_EB"
	case 76:
		return "LD C,H"
	case 184:
		return "CP A,B"
	case 237:
		return "ILLEGAL_ED"
	case 46:
		return "LD L,n8"
	case 84:
		return "LD D,H"
	case 162:
		return "AND A,D"
	case 45:
		return "DEC L"
	case 102:
		return "LD H,(HL)"
	case 168:
		return "XOR A,B"
	case 239:
		return "RST $28"
	case 255:
		return "RST $38"
	case 11:
		return "DEC BC"
	case 28:
		return "INC E"
	case 138:
		return "ADC A,D"
	case 159:
		return "SBC A,A"
	case 206:
		return "ADC A,n8"
	case 231:
		return "RST $20"
	case 25:
		return "ADD HL,DE"
	case 47:
		return "CPL"
	case 63:
		return "CCF"
	case 37:
		return "DEC H"
	case 74:
		return "LD C,D"
	case 78:
		return "LD C,(HL)"
	case 100:
		return "LD H,H"
	case 148:
		return "SUB A,H"
	case 151:
		return "SUB A,A"
	case 175:
		return "XOR A,A"
	case 178:
		return "OR A,D"
	case 5:
		return "DEC B"
	case 10:
		return "LD A,(BC)"
	case 49:
		return "LD SP,n16"
	case 51:
		return "INC SP"
	case 171:
		return "XOR A,E"
	case 181:
		return "OR A,L"
	case 221:
		return "ILLEGAL_DD"
	case 83:
		return "LD D,E"
	case 139:
		return "ADC A,E"
	case 227:
		return "ILLEGAL_E3"
	case 232:
		return "ADD SP,e8"
	case 31:
		return "RRA"
	case 33:
		return "LD HL,n16"
	case 50:
		return "LD (HL-),A"
	case 71:
		return "LD B,A"
	case 99:
		return "LD H,E"
	case 127:
		return "LD A,A"
	case 156:
		return "SBC A,H"
	case 193:
		return "POP BC"
	case 192:
		return "RET NZ"
	case 195:
		return "JP a16"
	case 17:
		return "LD DE,n16"
	case 29:
		return "DEC E"
	case 56:
		return "JR C,e8"
	case 94:
		return "LD E,(HL)"
	case 183:
		return "OR A,A"
	case 20:
		return "INC D"
	case 69:
		return "LD B,L"
	case 107:
		return "LD L,E"
	case 238:
		return "XOR A,n8"
	case 48:
		return "JR NC,e8"
	case 82:
		return "LD D,D"
	case 142:
		return "ADC A,(HL)"
	case 150:
		return "SUB A,(HL)"
	case 228:
		return "ILLEGAL_E4"
	case 241:
		return "POP AF"
	case 52:
		return "INC (HL)"
	case 112:
		return "LD (HL),B"
	case 128:
		return "ADD A,B"
	case 130:
		return "ADD A,D"
	case 187:
		return "CP A,E"
	case 207:
		return "RST $08"
	case 58:
		return "LD A,(HL-)"
	case 60:
		return "INC A"
	case 64:
		return "LD B,B"
	case 70:
		return "LD B,(HL)"
	case 116:
		return "LD (HL),H"
	case 131:
		return "ADD A,E"
	case 144:
		return "SUB A,B"
	case 190:
		return "CP A,(HL)"
	case 35:
		return "INC HL"
	case 174:
		return "XOR A,(HL)"
	case 236:
		return "ILLEGAL_EC"
	case 251:
		return "EI"
	case 111:
		return "LD L,A"
	case 153:
		return "SBC A,C"
	case 169:
		return "XOR A,C"
	case 209:
		return "POP DE"
	case 41:
		return "ADD HL,HL"
	case 113:
		return "LD (HL),C"
	case 199:
		return "RST $00"
	case 205:
		return "CALL a16"
	case 88:
		return "LD E,B"
	case 96:
		return "LD H,B"
	case 197:
		return "PUSH BC"
	case 246:
		return "OR A,n8"
	case 38:
		return "LD H,n8"
	case 125:
		return "LD A,L"
	case 6:
		return "LD B,n8"
	case 172:
		return "XOR A,H"
	case 179:
		return "OR A,E"
	case 194:
		return "JP NZ,a16"
	case 2:
		return "LD (BC),A"
	case 164:
		return "AND A,H"
	case 129:
		return "ADD A,C"
	case 140:
		return "ADC A,H"

	default:
		panic(fmt.Sprintf("Unknown code for %d", code))
	}
}
