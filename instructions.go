package gameboy

type Instruction func(cpu *CPU)

// LD A,(HL-)    code=0x3a
func LD_3A(cpu *CPU) {

	data := cpu.loadU8(cpu.HL())

	cpu.A = data

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
	cpu.IncProgramCounter("add")
	cpu.cycles += 8
}

// ADC A,C    code=0x89
func ADC_89(cpu *CPU) {
	// TODO: ADC_89
}

// LD D,C    code=0x51
func LD_51(cpu *CPU) {

	data := cpu.C

	cpu.D = data

	cpu.IncProgramCounter("ld")
	cpu.cycles += 4

}

// LD H,A    code=0x67
func LD_67(cpu *CPU) {

	data := cpu.A

	cpu.H = data

	cpu.IncProgramCounter("ld")
	cpu.cycles += 4

}

// SUB A,C    code=0x91
func SUB_91(cpu *CPU) {
	// TODO: SUB_91
}

// AND A,(HL)    code=0xa6
func AND_A6(cpu *CPU) {
	// TODO: AND_A6
}

// CP A,H    code=0xbc
func CP_BC(cpu *CPU) {
	// TODO: CP_BC
}

// RST $08    code=0xcf
func RST_CF(cpu *CPU) {
	// TODO: RST_CF
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

// LD E,C    code=0x59
func LD_59(cpu *CPU) {

	data := cpu.C

	cpu.E = data

	cpu.IncProgramCounter("ld")
	cpu.cycles += 4

}

// SBC A,E    code=0x9b
func SBC_9B(cpu *CPU) {
	// TODO: SBC_9B
}

// LD DE,n16    code=0x11
func LD_11(cpu *CPU) {

	data := cpu.readU16(cpu.PC)

	cpu.D, cpu.E = split(data)

	cpu.IncProgramCounter("ld")
	cpu.cycles += 12

}

// SCF     code=0x37
func SCF_37(cpu *CPU) {
	// TODO: SCF_37
}

// HALT     code=0x76
func HALT_76(cpu *CPU) {
	// TODO: HALT_76
}

// PREFIX     code=0xcb
func PREFIX_CB(cpu *CPU) {
	// TODO: PREFIX_CB
}

// DI     code=0xf3
func DI_F3(cpu *CPU) {
	// TODO: DI_F3
}

// RLCA     code=0x07
func RLCA_07(cpu *CPU) {
	// TODO: RLCA_07
}

// DEC A    code=0x3d
func DEC_3D(cpu *CPU) {
	res, flags := sub(cpu.A, 0x01)
	cpu.F = flags
	cpu.A = res
	cpu.IncProgramCounter("dec")
	cpu.cycles += 4
}

// LD (HL),L    code=0x75
func LD_75(cpu *CPU) {

	data := cpu.L

	cpu.WriteMemory(cpu.HL(), data)

	cpu.IncProgramCounter("ld")
	cpu.cycles += 8

}

// LD (HL),B    code=0x70
func LD_70(cpu *CPU) {

	data := cpu.B

	cpu.WriteMemory(cpu.HL(), data)

	cpu.IncProgramCounter("ld")
	cpu.cycles += 8

}

// SUB A,A    code=0x97
func SUB_97(cpu *CPU) {
	// TODO: SUB_97
}

// RST $30    code=0xf7
func RST_F7(cpu *CPU) {
	// TODO: RST_F7
}

// JR e8    code=0x18
func JR_18(cpu *CPU) {
	// TODO: JR_18
}

// ADD HL,HL    code=0x29
func ADD_29(cpu *CPU) {
	lhs := cpu.HL()
	rhs := cpu.HL()
	res, flags := add(lhs, rhs)
	cpu.H, cpu.L = split(res)
	cpu.F = flags
	cpu.IncProgramCounter("add")
	cpu.cycles += 8
}

// POP BC    code=0xc1
func POP_C1(cpu *CPU) {
	value := cpu.PopStack()
	cpu.B, cpu.C = split(value)
	cpu.IncProgramCounter("pop")
	cpu.cycles += 12
}

// LD D,n8    code=0x16
func LD_16(cpu *CPU) {

	data := cpu.readU8(cpu.PC)

	cpu.D = data

	cpu.IncProgramCounter("ld")
	cpu.cycles += 8

}

// ADD HL,DE    code=0x19
func ADD_19(cpu *CPU) {
	lhs := cpu.HL()
	rhs := cpu.DE()
	res, flags := add(lhs, rhs)
	cpu.H, cpu.L = split(res)
	cpu.F = flags
	cpu.IncProgramCounter("add")
	cpu.cycles += 8
}

// LD L,C    code=0x69
func LD_69(cpu *CPU) {

	data := cpu.C

	cpu.L = data

	cpu.IncProgramCounter("ld")
	cpu.cycles += 4

}

// LD A,L    code=0x7d
func LD_7D(cpu *CPU) {

	data := cpu.L

	cpu.A = data

	cpu.IncProgramCounter("ld")
	cpu.cycles += 4

}

// RET Z    code=0xc8
func RET_C8(cpu *CPU) {
	// TODO: RET_C8
}

// LD A,(HL)    code=0x7e
func LD_7E(cpu *CPU) {

	data := cpu.loadU8(cpu.HL())

	cpu.A = data

	cpu.IncProgramCounter("ld")
	cpu.cycles += 8

}

// ADC A,(HL)    code=0x8e
func ADC_8E(cpu *CPU) {
	// TODO: ADC_8E
}

// LD A,(a16)    code=0xfa
func LD_FA(cpu *CPU) {

	data := cpu.loadU8(cpu.readU16(cpu.PC))

	cpu.A = data

	cpu.IncProgramCounter("ld")
	cpu.cycles += 16

}

// DEC DE    code=0x1b
func DEC_1B(cpu *CPU) {
	res, flags := sub(cpu.DE(), 0x01)
	cpu.F = flags
	cpu.D, cpu.E = split(res)
	cpu.IncProgramCounter("dec")
	cpu.cycles += 8
}

// LD B,E    code=0x43
func LD_43(cpu *CPU) {

	data := cpu.E

	cpu.B = data

	cpu.IncProgramCounter("ld")
	cpu.cycles += 4

}

// ADD A,C    code=0x81
func ADD_81(cpu *CPU) {
	lhs := cpu.A
	rhs := cpu.C
	res, flags := add(lhs, rhs)
	cpu.A = res
	cpu.F = flags
	cpu.IncProgramCounter("add")
	cpu.cycles += 4
}

// LD C,A    code=0x4f
func LD_4F(cpu *CPU) {

	data := cpu.A

	cpu.C = data

	cpu.IncProgramCounter("ld")
	cpu.cycles += 4

}

// LD (HL),D    code=0x72
func LD_72(cpu *CPU) {

	data := cpu.D

	cpu.WriteMemory(cpu.HL(), data)

	cpu.IncProgramCounter("ld")
	cpu.cycles += 8

}

// CP A,B    code=0xb8
func CP_B8(cpu *CPU) {
	// TODO: CP_B8
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

// INC L    code=0x2c
func INC_2C(cpu *CPU) {
	res, flags := add(cpu.L, 0x01)
	cpu.F = flags
	cpu.L = res
	cpu.IncProgramCounter("inc")
	cpu.cycles += 4
}

// DEC L    code=0x2d
func DEC_2D(cpu *CPU) {
	res, flags := sub(cpu.L, 0x01)
	cpu.F = flags
	cpu.L = res
	cpu.IncProgramCounter("dec")
	cpu.cycles += 4
}

// LD H,D    code=0x62
func LD_62(cpu *CPU) {

	data := cpu.D

	cpu.H = data

	cpu.IncProgramCounter("ld")
	cpu.cycles += 4

}

// ADD A,L    code=0x85
func ADD_85(cpu *CPU) {
	lhs := cpu.A
	rhs := cpu.L
	res, flags := add(lhs, rhs)
	cpu.A = res
	cpu.F = flags
	cpu.IncProgramCounter("add")
	cpu.cycles += 4
}

// SUB A,E    code=0x93
func SUB_93(cpu *CPU) {
	// TODO: SUB_93
}

// LD L,A    code=0x6f
func LD_6F(cpu *CPU) {

	data := cpu.A

	cpu.L = data

	cpu.IncProgramCounter("ld")
	cpu.cycles += 4

}

// XOR A,D    code=0xaa
func XOR_AA(cpu *CPU) {
	// TODO: XOR_AA
}

// SUB A,n8    code=0xd6
func SUB_D6(cpu *CPU) {
	// TODO: SUB_D6
}

// INC DE    code=0x13
func INC_13(cpu *CPU) {
	res, flags := add(cpu.DE(), 0x01)
	cpu.F = flags
	cpu.D, cpu.E = split(res)
	cpu.IncProgramCounter("inc")
	cpu.cycles += 8
}

// INC (HL)    code=0x34
func INC_34(cpu *CPU) {
	res, flags := add(cpu.loadU8(cpu.HL()), 0x01)
	cpu.F = flags
	cpu.WriteMemory(cpu.HL(), res)
	cpu.IncProgramCounter("inc")
	cpu.cycles += 12
}

// LD C,(HL)    code=0x4e
func LD_4E(cpu *CPU) {

	data := cpu.loadU8(cpu.HL())

	cpu.C = data

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

// SUB A,D    code=0x92
func SUB_92(cpu *CPU) {
	// TODO: SUB_92
}

// JP Z,a16    code=0xca
func JP_CA(cpu *CPU) {
	// TODO: JP_CA
}

// ILLEGAL_EB     code=0xeb
func ILLEGAL_EB_EB(cpu *CPU) {
	// TODO: ILLEGAL_EB_EB
}

// SBC A,B    code=0x98
func SBC_98(cpu *CPU) {
	// TODO: SBC_98
}

// OR A,n8    code=0xf6
func OR_F6(cpu *CPU) {
	// TODO: OR_F6
}

// INC E    code=0x1c
func INC_1C(cpu *CPU) {
	res, flags := add(cpu.E, 0x01)
	cpu.F = flags
	cpu.E = res
	cpu.IncProgramCounter("inc")
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

// SBC A,A    code=0x9f
func SBC_9F(cpu *CPU) {
	// TODO: SBC_9F
}

// XOR A,E    code=0xab
func XOR_AB(cpu *CPU) {
	// TODO: XOR_AB
}

// ADC A,n8    code=0xce
func ADC_CE(cpu *CPU) {
	// TODO: ADC_CE
}

// RST $38    code=0xff
func RST_FF(cpu *CPU) {
	// TODO: RST_FF
}

// SBC A,D    code=0x9a
func SBC_9A(cpu *CPU) {
	// TODO: SBC_9A
}

// XOR A,L    code=0xad
func XOR_AD(cpu *CPU) {
	// TODO: XOR_AD
}

// CP A,E    code=0xbb
func CP_BB(cpu *CPU) {
	// TODO: CP_BB
}

// RET C    code=0xd8
func RET_D8(cpu *CPU) {
	// TODO: RET_D8
}

// LD HL,n16    code=0x21
func LD_21(cpu *CPU) {

	data := cpu.readU16(cpu.PC)

	cpu.H, cpu.L = split(data)

	cpu.IncProgramCounter("ld")
	cpu.cycles += 12

}

// LD B,C    code=0x41
func LD_41(cpu *CPU) {

	data := cpu.C

	cpu.B = data

	cpu.IncProgramCounter("ld")
	cpu.cycles += 4

}

// OR A,B    code=0xb0
func OR_B0(cpu *CPU) {
	// TODO: OR_B0
}

// LDH (a8),A    code=0xe0
func LDH_E0(cpu *CPU) {
	// TODO: LDH_E0
}

// ILLEGAL_E4     code=0xe4
func ILLEGAL_E4_E4(cpu *CPU) {
	// TODO: ILLEGAL_E4_E4
}

// LD D,D    code=0x52
func LD_52(cpu *CPU) {

	data := cpu.D

	cpu.D = data

	cpu.IncProgramCounter("ld")
	cpu.cycles += 4

}

// LD E,(HL)    code=0x5e
func LD_5E(cpu *CPU) {

	data := cpu.loadU8(cpu.HL())

	cpu.E = data

	cpu.IncProgramCounter("ld")
	cpu.cycles += 8

}

// RETI     code=0xd9
func RETI_D9(cpu *CPU) {
	// TODO: RETI_D9
}

// DEC C    code=0x0d
func DEC_0D(cpu *CPU) {
	res, flags := sub(cpu.C, 0x01)
	cpu.F = flags
	cpu.C = res
	cpu.IncProgramCounter("dec")
	cpu.cycles += 4
}

// LD L,E    code=0x6b
func LD_6B(cpu *CPU) {

	data := cpu.E

	cpu.L = data

	cpu.IncProgramCounter("ld")
	cpu.cycles += 4

}

// SBC A,(HL)    code=0x9e
func SBC_9E(cpu *CPU) {
	// TODO: SBC_9E
}

// POP AF    code=0xf1
func POP_F1(cpu *CPU) {
	value := cpu.PopStack()
	msb, lsb := split(value)
	cpu.A, cpu.F = msb, FlagRegister(lsb)
	cpu.IncProgramCounter("pop")
	cpu.cycles += 12
}

// LD C,H    code=0x4c
func LD_4C(cpu *CPU) {

	data := cpu.H

	cpu.C = data

	cpu.IncProgramCounter("ld")
	cpu.cycles += 4

}

// ADD A,A    code=0x87
func ADD_87(cpu *CPU) {
	lhs := cpu.A
	rhs := cpu.A
	res, flags := add(lhs, rhs)
	cpu.A = res
	cpu.F = flags
	cpu.IncProgramCounter("add")
	cpu.cycles += 4
}

// SUB A,L    code=0x95
func SUB_95(cpu *CPU) {
	// TODO: SUB_95
}

// SBC A,L    code=0x9d
func SBC_9D(cpu *CPU) {
	// TODO: SBC_9D
}

// OR A,D    code=0xb2
func OR_B2(cpu *CPU) {
	// TODO: OR_B2
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

// RRA     code=0x1f
func RRA_1F(cpu *CPU) {
	// TODO: RRA_1F
}

// LD L,n8    code=0x2e
func LD_2E(cpu *CPU) {

	data := cpu.readU8(cpu.PC)

	cpu.L = data

	cpu.IncProgramCounter("ld")
	cpu.cycles += 8

}

// LD C,C    code=0x49
func LD_49(cpu *CPU) {

	data := cpu.C

	cpu.C = data

	cpu.IncProgramCounter("ld")
	cpu.cycles += 4

}

// LD D,L    code=0x55
func LD_55(cpu *CPU) {

	data := cpu.L

	cpu.D = data

	cpu.IncProgramCounter("ld")
	cpu.cycles += 4

}

// SUB A,B    code=0x90
func SUB_90(cpu *CPU) {
	// TODO: SUB_90
}

// AND A,D    code=0xa2
func AND_A2(cpu *CPU) {
	// TODO: AND_A2
}

// XOR A,H    code=0xac
func XOR_AC(cpu *CPU) {
	// TODO: XOR_AC
}

// LDH A,(C)    code=0xf2
func LDH_F2(cpu *CPU) {
	// TODO: LDH_F2
}

// ILLEGAL_FD     code=0xfd
func ILLEGAL_FD_FD(cpu *CPU) {
	// TODO: ILLEGAL_FD_FD
}

// LD (HL),n8    code=0x36
func LD_36(cpu *CPU) {

	data := cpu.readU8(cpu.PC)

	cpu.WriteMemory(cpu.HL(), data)

	cpu.IncProgramCounter("ld")
	cpu.cycles += 12

}

// JR C,e8    code=0x38
func JR_38(cpu *CPU) {
	// TODO: JR_38
}

// LD (HL),E    code=0x73
func LD_73(cpu *CPU) {

	data := cpu.E

	cpu.WriteMemory(cpu.HL(), data)

	cpu.IncProgramCounter("ld")
	cpu.cycles += 8

}

// AND A,C    code=0xa1
func AND_A1(cpu *CPU) {
	// TODO: AND_A1
}

// XOR A,(HL)    code=0xae
func XOR_AE(cpu *CPU) {
	// TODO: XOR_AE
}

// DEC B    code=0x05
func DEC_05(cpu *CPU) {
	res, flags := sub(cpu.B, 0x01)
	cpu.F = flags
	cpu.B = res
	cpu.IncProgramCounter("dec")
	cpu.cycles += 4
}

// AND A,B    code=0xa0
func AND_A0(cpu *CPU) {
	// TODO: AND_A0
}

// PUSH AF    code=0xf5
func PUSH_F5(cpu *CPU) {
	cpu.PushStack(cpu.AF())
	cpu.cycles += 16
}

// CCF     code=0x3f
func CCF_3F(cpu *CPU) {
	// TODO: CCF_3F
}

// LD C,E    code=0x4b
func LD_4B(cpu *CPU) {

	data := cpu.E

	cpu.C = data

	cpu.IncProgramCounter("ld")
	cpu.cycles += 4

}

// CP A,C    code=0xb9
func CP_B9(cpu *CPU) {
	// TODO: CP_B9
}

// RST $00    code=0xc7
func RST_C7(cpu *CPU) {
	// TODO: RST_C7
}

// PUSH DE    code=0xd5
func PUSH_D5(cpu *CPU) {
	cpu.PushStack(cpu.DE())
	cpu.cycles += 16
}

// PUSH HL    code=0xe5
func PUSH_E5(cpu *CPU) {
	cpu.PushStack(cpu.HL())
	cpu.cycles += 16
}

// CP A,n8    code=0xfe
func CP_FE(cpu *CPU) {
	// TODO: CP_FE
}

// LD E,H    code=0x5c
func LD_5C(cpu *CPU) {

	data := cpu.H

	cpu.E = data

	cpu.IncProgramCounter("ld")
	cpu.cycles += 4

}

// LD L,H    code=0x6c
func LD_6C(cpu *CPU) {

	data := cpu.H

	cpu.L = data

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

// SBC A,C    code=0x99
func SBC_99(cpu *CPU) {
	// TODO: SBC_99
}

// LD E,n8    code=0x1e
func LD_1E(cpu *CPU) {

	data := cpu.readU8(cpu.PC)

	cpu.E = data

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

// LD C,L    code=0x4d
func LD_4D(cpu *CPU) {

	data := cpu.L

	cpu.C = data

	cpu.IncProgramCounter("ld")
	cpu.cycles += 4

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

// LD B,D    code=0x42
func LD_42(cpu *CPU) {

	data := cpu.D

	cpu.B = data

	cpu.IncProgramCounter("ld")
	cpu.cycles += 4

}

// ADC A,L    code=0x8d
func ADC_8D(cpu *CPU) {
	// TODO: ADC_8D
}

// ADC A,A    code=0x8f
func ADC_8F(cpu *CPU) {
	// TODO: ADC_8F
}

// PUSH BC    code=0xc5
func PUSH_C5(cpu *CPU) {
	cpu.PushStack(cpu.BC())
	cpu.cycles += 16
}

// CPL     code=0x2f
func CPL_2F(cpu *CPU) {
	// TODO: CPL_2F
}

// LD D,H    code=0x54
func LD_54(cpu *CPU) {

	data := cpu.H

	cpu.D = data

	cpu.IncProgramCounter("ld")
	cpu.cycles += 4

}

// LD H,H    code=0x64
func LD_64(cpu *CPU) {

	data := cpu.H

	cpu.H = data

	cpu.IncProgramCounter("ld")
	cpu.cycles += 4

}

// RST $20    code=0xe7
func RST_E7(cpu *CPU) {
	// TODO: RST_E7
}

// LD B,B    code=0x40
func LD_40(cpu *CPU) {

	data := cpu.B

	cpu.B = data

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

// DEC D    code=0x15
func DEC_15(cpu *CPU) {
	res, flags := sub(cpu.D, 0x01)
	cpu.F = flags
	cpu.D = res
	cpu.IncProgramCounter("dec")
	cpu.cycles += 4
}

// DEC E    code=0x1d
func DEC_1D(cpu *CPU) {
	res, flags := sub(cpu.E, 0x01)
	cpu.F = flags
	cpu.E = res
	cpu.IncProgramCounter("dec")
	cpu.cycles += 4
}

// LD L,D    code=0x6a
func LD_6A(cpu *CPU) {

	data := cpu.D

	cpu.L = data

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

// ADD A,D    code=0x82
func ADD_82(cpu *CPU) {
	lhs := cpu.A
	rhs := cpu.D
	res, flags := add(lhs, rhs)
	cpu.A = res
	cpu.F = flags
	cpu.IncProgramCounter("add")
	cpu.cycles += 4
}

// RLA     code=0x17
func RLA_17(cpu *CPU) {
	// TODO: RLA_17
}

// LD A,(HL+)    code=0x2a
func LD_2A(cpu *CPU) {

	data := cpu.loadU8(cpu.HL())

	cpu.A = data

	cpu.IncProgramCounter("ld")
	cpu.cycles += 8

}

// LD A,B    code=0x78
func LD_78(cpu *CPU) {

	data := cpu.B

	cpu.A = data

	cpu.IncProgramCounter("ld")
	cpu.cycles += 4

}

// AND A,L    code=0xa5
func AND_A5(cpu *CPU) {
	// TODO: AND_A5
}

// RST $18    code=0xdf
func RST_DF(cpu *CPU) {
	// TODO: RST_DF
}

// JP HL    code=0xe9
func JP_E9(cpu *CPU) {
	// TODO: JP_E9
}

// CP A,L    code=0xbd
func CP_BD(cpu *CPU) {
	// TODO: CP_BD
}

// POP HL    code=0xe1
func POP_E1(cpu *CPU) {
	value := cpu.PopStack()
	cpu.H, cpu.L = split(value)
	cpu.IncProgramCounter("pop")
	cpu.cycles += 12
}

// AND A,n8    code=0xe6
func AND_E6(cpu *CPU) {
	// TODO: AND_E6
}

// LDH A,(a8)    code=0xf0
func LDH_F0(cpu *CPU) {
	// TODO: LDH_F0
}

// DEC BC    code=0x0b
func DEC_0B(cpu *CPU) {
	res, flags := sub(cpu.BC(), 0x01)
	cpu.F = flags
	cpu.B, cpu.C = split(res)
	cpu.IncProgramCounter("dec")
	cpu.cycles += 8
}

// ADD HL,SP    code=0x39
func ADD_39(cpu *CPU) {
	lhs := cpu.HL()
	rhs := cpu.SP
	res, flags := add(lhs, rhs)
	cpu.H, cpu.L = split(res)
	cpu.F = flags
	cpu.IncProgramCounter("add")
	cpu.cycles += 8
}

// LD B,n8    code=0x06
func LD_06(cpu *CPU) {

	data := cpu.readU8(cpu.PC)

	cpu.B = data

	cpu.IncProgramCounter("ld")
	cpu.cycles += 8

}

// LD B,(HL)    code=0x46
func LD_46(cpu *CPU) {

	data := cpu.loadU8(cpu.HL())

	cpu.B = data

	cpu.IncProgramCounter("ld")
	cpu.cycles += 8

}

// LD H,(HL)    code=0x66
func LD_66(cpu *CPU) {

	data := cpu.loadU8(cpu.HL())

	cpu.H = data

	cpu.IncProgramCounter("ld")
	cpu.cycles += 8

}

// LD A,C    code=0x79
func LD_79(cpu *CPU) {

	data := cpu.C

	cpu.A = data

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

// ADD A,E    code=0x83
func ADD_83(cpu *CPU) {
	lhs := cpu.A
	rhs := cpu.E
	res, flags := add(lhs, rhs)
	cpu.A = res
	cpu.F = flags
	cpu.IncProgramCounter("add")
	cpu.cycles += 4
}

// SUB A,(HL)    code=0x96
func SUB_96(cpu *CPU) {
	// TODO: SUB_96
}

// INC B    code=0x04
func INC_04(cpu *CPU) {
	res, flags := add(cpu.B, 0x01)
	cpu.F = flags
	cpu.B = res
	cpu.IncProgramCounter("inc")
	cpu.cycles += 4
}

// DAA     code=0x27
func DAA_27(cpu *CPU) {
	// TODO: DAA_27
}

// LD A,n8    code=0x3e
func LD_3E(cpu *CPU) {

	data := cpu.readU8(cpu.PC)

	cpu.A = data

	cpu.IncProgramCounter("ld")
	cpu.cycles += 8

}

// LD (a16),A    code=0xea
func LD_EA(cpu *CPU) {

	data := cpu.A

	cpu.WriteMemory(cpu.readU16(cpu.PC), data)

	cpu.IncProgramCounter("ld")
	cpu.cycles += 16

}

// ADC A,E    code=0x8b
func ADC_8B(cpu *CPU) {
	// TODO: ADC_8B
}

// ADD A,n8    code=0xc6
func ADD_C6(cpu *CPU) {
	lhs := cpu.A
	rhs := cpu.readU8(cpu.PC)
	res, flags := add(lhs, rhs)
	cpu.A = res
	cpu.F = flags
	cpu.IncProgramCounter("add")
	cpu.cycles += 8
}

// LD A,(BC)    code=0x0a
func LD_0A(cpu *CPU) {

	data := cpu.loadU8(cpu.BC())

	cpu.A = data

	cpu.IncProgramCounter("ld")
	cpu.cycles += 8

}

// SUB A,H    code=0x94
func SUB_94(cpu *CPU) {
	// TODO: SUB_94
}

// RET NZ    code=0xc0
func RET_C0(cpu *CPU) {
	// TODO: RET_C0
}

// RET     code=0xc9
func RET_C9(cpu *CPU) {
	// TODO: RET_C9
}

// ADD SP,e8    code=0xe8
func ADD_E8(cpu *CPU) {
	lhs := cpu.SP
	rhs := cpu.readI8(cpu.PC)
	res, flags := add(lhs, rhs)
	cpu.SP = res
	cpu.F = flags
	cpu.IncProgramCounter("add")
	cpu.cycles += 16
}

// AND A,H    code=0xa4
func AND_A4(cpu *CPU) {
	// TODO: AND_A4
}

// XOR A,C    code=0xa9
func XOR_A9(cpu *CPU) {
	// TODO: XOR_A9
}

// NOP     code=0x00
func NOP_00(cpu *CPU) {
	// TODO: NOP_00
}

// LD D,(HL)    code=0x56
func LD_56(cpu *CPU) {

	data := cpu.loadU8(cpu.HL())

	cpu.D = data

	cpu.IncProgramCounter("ld")
	cpu.cycles += 8

}

// LD D,A    code=0x57
func LD_57(cpu *CPU) {

	data := cpu.A

	cpu.D = data

	cpu.IncProgramCounter("ld")
	cpu.cycles += 4

}

// ADC A,H    code=0x8c
func ADC_8C(cpu *CPU) {
	// TODO: ADC_8C
}

// LD C,n8    code=0x0e
func LD_0E(cpu *CPU) {

	data := cpu.readU8(cpu.PC)

	cpu.C = data

	cpu.IncProgramCounter("ld")
	cpu.cycles += 8

}

// INC SP    code=0x33
func INC_33(cpu *CPU) {
	res, flags := add(cpu.SP, 0x01)
	cpu.F = flags
	cpu.SP = res
	cpu.IncProgramCounter("inc")
	cpu.cycles += 8
}

// LD C,D    code=0x4a
func LD_4A(cpu *CPU) {

	data := cpu.D

	cpu.C = data

	cpu.IncProgramCounter("ld")
	cpu.cycles += 4

}

// JP C,a16    code=0xda
func JP_DA(cpu *CPU) {
	// TODO: JP_DA
}

// JR Z,e8    code=0x28
func JR_28(cpu *CPU) {
	// TODO: JR_28
}

// LD B,H    code=0x44
func LD_44(cpu *CPU) {

	data := cpu.H

	cpu.B = data

	cpu.IncProgramCounter("ld")
	cpu.cycles += 4

}

// SBC A,H    code=0x9c
func SBC_9C(cpu *CPU) {
	// TODO: SBC_9C
}

// CP A,(HL)    code=0xbe
func CP_BE(cpu *CPU) {
	// TODO: CP_BE
}

// ILLEGAL_D3     code=0xd3
func ILLEGAL_D3_D3(cpu *CPU) {
	// TODO: ILLEGAL_D3_D3
}

// ILLEGAL_EC     code=0xec
func ILLEGAL_EC_EC(cpu *CPU) {
	// TODO: ILLEGAL_EC_EC
}

// LD A,E    code=0x7b
func LD_7B(cpu *CPU) {

	data := cpu.E

	cpu.A = data

	cpu.IncProgramCounter("ld")
	cpu.cycles += 4

}

// POP DE    code=0xd1
func POP_D1(cpu *CPU) {
	value := cpu.PopStack()
	cpu.D, cpu.E = split(value)
	cpu.IncProgramCounter("pop")
	cpu.cycles += 12
}

// LD E,L    code=0x5d
func LD_5D(cpu *CPU) {

	data := cpu.L

	cpu.E = data

	cpu.IncProgramCounter("ld")
	cpu.cycles += 4

}

// OR A,E    code=0xb3
func OR_B3(cpu *CPU) {
	// TODO: OR_B3
}

// CP A,A    code=0xbf
func CP_BF(cpu *CPU) {
	// TODO: CP_BF
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

// ILLEGAL_F4     code=0xf4
func ILLEGAL_F4_F4(cpu *CPU) {
	// TODO: ILLEGAL_F4_F4
}

// INC BC    code=0x03
func INC_03(cpu *CPU) {
	res, flags := add(cpu.BC(), 0x01)
	cpu.F = flags
	cpu.B, cpu.C = split(res)
	cpu.IncProgramCounter("inc")
	cpu.cycles += 8
}

// ADD HL,BC    code=0x09
func ADD_09(cpu *CPU) {
	lhs := cpu.HL()
	rhs := cpu.BC()
	res, flags := add(lhs, rhs)
	cpu.H, cpu.L = split(res)
	cpu.F = flags
	cpu.IncProgramCounter("add")
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

// LD E,A    code=0x5f
func LD_5F(cpu *CPU) {

	data := cpu.A

	cpu.E = data

	cpu.IncProgramCounter("ld")
	cpu.cycles += 4

}

// ADC A,B    code=0x88
func ADC_88(cpu *CPU) {
	// TODO: ADC_88
}

// XOR A,B    code=0xa8
func XOR_A8(cpu *CPU) {
	// TODO: XOR_A8
}

// RST $28    code=0xef
func RST_EF(cpu *CPU) {
	// TODO: RST_EF
}

// LD SP,HL    code=0xf9
func LD_F9(cpu *CPU) {

	data := cpu.HL()

	cpu.SP = data

	cpu.IncProgramCounter("ld")
	cpu.cycles += 8

}

// JP NC,a16    code=0xd2
func JP_D2(cpu *CPU) {
	// TODO: JP_D2
}

// LD B,L    code=0x45
func LD_45(cpu *CPU) {

	data := cpu.L

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

// ILLEGAL_E3     code=0xe3
func ILLEGAL_E3_E3(cpu *CPU) {
	// TODO: ILLEGAL_E3_E3
}

// LD BC,n16    code=0x01
func LD_01(cpu *CPU) {

	data := cpu.readU16(cpu.PC)

	cpu.B, cpu.C = split(data)

	cpu.IncProgramCounter("ld")
	cpu.cycles += 12

}

// LD H,n8    code=0x26
func LD_26(cpu *CPU) {

	data := cpu.readU8(cpu.PC)

	cpu.H = data

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

// LD E,B    code=0x58
func LD_58(cpu *CPU) {

	data := cpu.B

	cpu.E = data

	cpu.IncProgramCounter("ld")
	cpu.cycles += 4

}

// ADC A,D    code=0x8a
func ADC_8A(cpu *CPU) {
	// TODO: ADC_8A
}

// AND A,E    code=0xa3
func AND_A3(cpu *CPU) {
	// TODO: AND_A3
}

// RST $10    code=0xd7
func RST_D7(cpu *CPU) {
	// TODO: RST_D7
}

// LD (BC),A    code=0x02
func LD_02(cpu *CPU) {

	data := cpu.A

	cpu.WriteMemory(cpu.BC(), data)

	cpu.IncProgramCounter("ld")
	cpu.cycles += 8

}

// LD D,E    code=0x53
func LD_53(cpu *CPU) {

	data := cpu.E

	cpu.D = data

	cpu.IncProgramCounter("ld")
	cpu.cycles += 4

}

// INC H    code=0x24
func INC_24(cpu *CPU) {
	res, flags := add(cpu.H, 0x01)
	cpu.F = flags
	cpu.H = res
	cpu.IncProgramCounter("inc")
	cpu.cycles += 4
}

// OR A,C    code=0xb1
func OR_B1(cpu *CPU) {
	// TODO: OR_B1
}

// DEC (HL)    code=0x35
func DEC_35(cpu *CPU) {
	res, flags := sub(cpu.loadU8(cpu.HL()), 0x01)
	cpu.F = flags
	cpu.WriteMemory(cpu.HL(), res)
	cpu.IncProgramCounter("dec")
	cpu.cycles += 12
}

// JR NZ,e8    code=0x20
func JR_20(cpu *CPU) {
	// TODO: JR_20
}

// INC HL    code=0x23
func INC_23(cpu *CPU) {
	res, flags := add(cpu.HL(), 0x01)
	cpu.F = flags
	cpu.H, cpu.L = split(res)
	cpu.IncProgramCounter("inc")
	cpu.cycles += 8
}

// LD A,(DE)    code=0x1a
func LD_1A(cpu *CPU) {

	data := cpu.loadU8(cpu.DE())

	cpu.A = data

	cpu.IncProgramCounter("ld")
	cpu.cycles += 8

}

// LD H,B    code=0x60
func LD_60(cpu *CPU) {

	data := cpu.B

	cpu.H = data

	cpu.IncProgramCounter("ld")
	cpu.cycles += 4

}

// XOR A,A    code=0xaf
func XOR_AF(cpu *CPU) {
	// TODO: XOR_AF
}

// LD (a16),SP    code=0x08
func LD_08(cpu *CPU) {

	data := cpu.SP

	cpu.WriteMemory(cpu.readU16(cpu.PC), data)

	cpu.IncProgramCounter("ld")
	cpu.cycles += 20

}

// JR NC,e8    code=0x30
func JR_30(cpu *CPU) {
	// TODO: JR_30
}

// LD H,C    code=0x61
func LD_61(cpu *CPU) {

	data := cpu.C

	cpu.H = data

	cpu.IncProgramCounter("ld")
	cpu.cycles += 4

}

// LD L,(HL)    code=0x6e
func LD_6E(cpu *CPU) {

	data := cpu.loadU8(cpu.HL())

	cpu.L = data

	cpu.IncProgramCounter("ld")
	cpu.cycles += 8

}

// LD (HL),C    code=0x71
func LD_71(cpu *CPU) {

	data := cpu.C

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
	cpu.IncProgramCounter("add")
	cpu.cycles += 4
}

// OR A,H    code=0xb4
func OR_B4(cpu *CPU) {
	// TODO: OR_B4
}

// CP A,D    code=0xba
func CP_BA(cpu *CPU) {
	// TODO: CP_BA
}

// INC C    code=0x0c
func INC_0C(cpu *CPU) {
	res, flags := add(cpu.C, 0x01)
	cpu.F = flags
	cpu.C = res
	cpu.IncProgramCounter("inc")
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

// LD L,B    code=0x68
func LD_68(cpu *CPU) {

	data := cpu.B

	cpu.L = data

	cpu.IncProgramCounter("ld")
	cpu.cycles += 4

}

// JP a16    code=0xc3
func JP_C3(cpu *CPU) {
	// TODO: JP_C3
}

// LD E,D    code=0x5a
func LD_5A(cpu *CPU) {

	data := cpu.D

	cpu.E = data

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

// ADD A,H    code=0x84
func ADD_84(cpu *CPU) {
	lhs := cpu.A
	rhs := cpu.H
	res, flags := add(lhs, rhs)
	cpu.A = res
	cpu.F = flags
	cpu.IncProgramCounter("add")
	cpu.cycles += 4
}

// ILLEGAL_DD     code=0xdd
func ILLEGAL_DD_DD(cpu *CPU) {
	// TODO: ILLEGAL_DD_DD
}

// LD H,L    code=0x65
func LD_65(cpu *CPU) {

	data := cpu.L

	cpu.H = data

	cpu.IncProgramCounter("ld")
	cpu.cycles += 4

}

// AND A,A    code=0xa7
func AND_A7(cpu *CPU) {
	// TODO: AND_A7
}

// LD B,A    code=0x47
func LD_47(cpu *CPU) {

	data := cpu.A

	cpu.B = data

	cpu.IncProgramCounter("ld")
	cpu.cycles += 4

}

// OR A,L    code=0xb5
func OR_B5(cpu *CPU) {
	// TODO: OR_B5
}

// LDH (C),A    code=0xe2
func LDH_E2(cpu *CPU) {
	// TODO: LDH_E2
}

// EI     code=0xfb
func EI_FB(cpu *CPU) {
	// TODO: EI_FB
}

// DEC H    code=0x25
func DEC_25(cpu *CPU) {
	res, flags := sub(cpu.H, 0x01)
	cpu.F = flags
	cpu.H = res
	cpu.IncProgramCounter("dec")
	cpu.cycles += 4
}

// LD C,B    code=0x48
func LD_48(cpu *CPU) {

	data := cpu.B

	cpu.C = data

	cpu.IncProgramCounter("ld")
	cpu.cycles += 4

}

// OR A,(HL)    code=0xb6
func OR_B6(cpu *CPU) {
	// TODO: OR_B6
}

// JP NZ,a16    code=0xc2
func JP_C2(cpu *CPU) {
	// TODO: JP_C2
}

// ILLEGAL_DB     code=0xdb
func ILLEGAL_DB_DB(cpu *CPU) {
	// TODO: ILLEGAL_DB_DB
}

// SBC A,n8    code=0xde
func SBC_DE(cpu *CPU) {
	// TODO: SBC_DE
}

// LD L,L    code=0x6d
func LD_6D(cpu *CPU) {

	data := cpu.L

	cpu.L = data

	cpu.IncProgramCounter("ld")
	cpu.cycles += 4

}

// OR A,A    code=0xb7
func OR_B7(cpu *CPU) {
	// TODO: OR_B7
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

// RET NC    code=0xd0
func RET_D0(cpu *CPU) {
	// TODO: RET_D0
}

// XOR A,n8    code=0xee
func XOR_EE(cpu *CPU) {
	// TODO: XOR_EE
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

// ILLEGAL_FC     code=0xfc
func ILLEGAL_FC_FC(cpu *CPU) {
	// TODO: ILLEGAL_FC_FC
}

// DEC HL    code=0x2b
func DEC_2B(cpu *CPU) {
	res, flags := sub(cpu.HL(), 0x01)
	cpu.F = flags
	cpu.H, cpu.L = split(res)
	cpu.IncProgramCounter("dec")
	cpu.cycles += 8
}

// LD D,B    code=0x50
func LD_50(cpu *CPU) {

	data := cpu.B

	cpu.D = data

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

// ILLEGAL_ED     code=0xed
func ILLEGAL_ED_ED(cpu *CPU) {
	// TODO: ILLEGAL_ED_ED
}

var ops = map[uint8]Instruction{
	0x3a: LD_3A,
	0x86: ADD_86,
	0x89: ADC_89,
	0x51: LD_51,
	0x67: LD_67,
	0x91: SUB_91,
	0xa6: AND_A6,
	0xbc: CP_BC,
	0xcf: RST_CF,
	0x32: LD_32,
	0x59: LD_59,
	0x9b: SBC_9B,
	0x11: LD_11,
	0x37: SCF_37,
	0x76: HALT_76,
	0xcb: PREFIX_CB,
	0xf3: DI_F3,
	0x7:  RLCA_07,
	0x3d: DEC_3D,
	0x75: LD_75,
	0x70: LD_70,
	0x97: SUB_97,
	0xf7: RST_F7,
	0x18: JR_18,
	0x29: ADD_29,
	0xc1: POP_C1,
	0x16: LD_16,
	0x19: ADD_19,
	0x69: LD_69,
	0x7d: LD_7D,
	0xc8: RET_C8,
	0x7e: LD_7E,
	0x8e: ADC_8E,
	0xfa: LD_FA,
	0x1b: DEC_1B,
	0x43: LD_43,
	0x81: ADD_81,
	0x4f: LD_4F,
	0x72: LD_72,
	0xb8: CP_B8,
	0xd4: CALL_D4,
	0x2c: INC_2C,
	0x2d: DEC_2D,
	0x62: LD_62,
	0x85: ADD_85,
	0x93: SUB_93,
	0x6f: LD_6F,
	0xaa: XOR_AA,
	0xd6: SUB_D6,
	0x13: INC_13,
	0x34: INC_34,
	0x4e: LD_4E,
	0x74: LD_74,
	0x92: SUB_92,
	0xca: JP_CA,
	0xeb: ILLEGAL_EB_EB,
	0x98: SBC_98,
	0xf6: OR_F6,
	0x1c: INC_1C,
	0x3b: DEC_3B,
	0x9f: SBC_9F,
	0xab: XOR_AB,
	0xce: ADC_CE,
	0xff: RST_FF,
	0x9a: SBC_9A,
	0xad: XOR_AD,
	0xbb: CP_BB,
	0xd8: RET_D8,
	0x21: LD_21,
	0x41: LD_41,
	0xb0: OR_B0,
	0xe0: LDH_E0,
	0xe4: ILLEGAL_E4_E4,
	0x52: LD_52,
	0x5e: LD_5E,
	0xd9: RETI_D9,
	0xd:  DEC_0D,
	0x6b: LD_6B,
	0x9e: SBC_9E,
	0xf1: POP_F1,
	0x4c: LD_4C,
	0x87: ADD_87,
	0x95: SUB_95,
	0x9d: SBC_9D,
	0xb2: OR_B2,
	0xc4: CALL_C4,
	0x1f: RRA_1F,
	0x2e: LD_2E,
	0x49: LD_49,
	0x55: LD_55,
	0x90: SUB_90,
	0xa2: AND_A2,
	0xac: XOR_AC,
	0xf2: LDH_F2,
	0xfd: ILLEGAL_FD_FD,
	0x36: LD_36,
	0x38: JR_38,
	0x73: LD_73,
	0xa1: AND_A1,
	0xae: XOR_AE,
	0x5:  DEC_05,
	0xa0: AND_A0,
	0xf5: PUSH_F5,
	0x3f: CCF_3F,
	0x4b: LD_4B,
	0xb9: CP_B9,
	0xc7: RST_C7,
	0xd5: PUSH_D5,
	0xe5: PUSH_E5,
	0xfe: CP_FE,
	0x5c: LD_5C,
	0x6c: LD_6C,
	0x77: LD_77,
	0x99: SBC_99,
	0x1e: LD_1E,
	0x22: LD_22,
	0x4d: LD_4D,
	0xdc: CALL_DC,
	0x42: LD_42,
	0x8d: ADC_8D,
	0x8f: ADC_8F,
	0xc5: PUSH_C5,
	0x2f: CPL_2F,
	0x54: LD_54,
	0x64: LD_64,
	0xe7: RST_E7,
	0x40: LD_40,
	0x12: LD_12,
	0x15: DEC_15,
	0x1d: DEC_1D,
	0x6a: LD_6A,
	0x7c: LD_7C,
	0x82: ADD_82,
	0x17: RLA_17,
	0x2a: LD_2A,
	0x78: LD_78,
	0xa5: AND_A5,
	0xdf: RST_DF,
	0xe9: JP_E9,
	0xbd: CP_BD,
	0xe1: POP_E1,
	0xe6: AND_E6,
	0xf0: LDH_F0,
	0xb:  DEC_0B,
	0x39: ADD_39,
	0x6:  LD_06,
	0x46: LD_46,
	0x66: LD_66,
	0x79: LD_79,
	0x7a: LD_7A,
	0x83: ADD_83,
	0x96: SUB_96,
	0x4:  INC_04,
	0x27: DAA_27,
	0x3e: LD_3E,
	0xea: LD_EA,
	0x8b: ADC_8B,
	0xc6: ADD_C6,
	0xa:  LD_0A,
	0x94: SUB_94,
	0xc0: RET_C0,
	0xc9: RET_C9,
	0xe8: ADD_E8,
	0xa4: AND_A4,
	0xa9: XOR_A9,
	0x0:  NOP_00,
	0x56: LD_56,
	0x57: LD_57,
	0x8c: ADC_8C,
	0xe:  LD_0E,
	0x33: INC_33,
	0x4a: LD_4A,
	0xda: JP_DA,
	0x28: JR_28,
	0x44: LD_44,
	0x9c: SBC_9C,
	0xbe: CP_BE,
	0xd3: ILLEGAL_D3_D3,
	0xec: ILLEGAL_EC_EC,
	0x7b: LD_7B,
	0xd1: POP_D1,
	0x5d: LD_5D,
	0xb3: OR_B3,
	0xbf: CP_BF,
	0xcd: CALL_CD,
	0xf4: ILLEGAL_F4_F4,
	0x3:  INC_03,
	0x9:  ADD_09,
	0x3c: INC_3C,
	0x5f: LD_5F,
	0x88: ADC_88,
	0xa8: XOR_A8,
	0xef: RST_EF,
	0xf9: LD_F9,
	0xd2: JP_D2,
	0x45: LD_45,
	0x63: LD_63,
	0xe3: ILLEGAL_E3_E3,
	0x1:  LD_01,
	0x26: LD_26,
	0x31: LD_31,
	0x58: LD_58,
	0x8a: ADC_8A,
	0xa3: AND_A3,
	0xd7: RST_D7,
	0x2:  LD_02,
	0x53: LD_53,
	0x24: INC_24,
	0xb1: OR_B1,
	0x35: DEC_35,
	0x20: JR_20,
	0x23: INC_23,
	0x1a: LD_1A,
	0x60: LD_60,
	0xaf: XOR_AF,
	0x8:  LD_08,
	0x30: JR_30,
	0x61: LD_61,
	0x6e: LD_6E,
	0x71: LD_71,
	0x80: ADD_80,
	0xb4: OR_B4,
	0xba: CP_BA,
	0xc:  INC_0C,
	0x14: INC_14,
	0x68: LD_68,
	0xc3: JP_C3,
	0x5a: LD_5A,
	0x7f: LD_7F,
	0x84: ADD_84,
	0xdd: ILLEGAL_DD_DD,
	0x65: LD_65,
	0xa7: AND_A7,
	0x47: LD_47,
	0xb5: OR_B5,
	0xe2: LDH_E2,
	0xfb: EI_FB,
	0x25: DEC_25,
	0x48: LD_48,
	0xb6: OR_B6,
	0xc2: JP_C2,
	0xdb: ILLEGAL_DB_DB,
	0xde: SBC_DE,
	0x6d: LD_6D,
	0xb7: OR_B7,
	0xcc: CALL_CC,
	0xd0: RET_D0,
	0xee: XOR_EE,
	0xf8: LD_F8,
	0xfc: ILLEGAL_FC_FC,
	0x2b: DEC_2B,
	0x50: LD_50,
	0x5b: LD_5B,
	0xed: ILLEGAL_ED_ED,
}
