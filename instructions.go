package gameboy

import "fmt"

type Instruction interface {
	Exec(cpu *CPU)
	Code() uint8
	String() string
}

// POP DE    code=0xd1
type POP_D1 struct{}

func (POP_D1) Exec(cpu *CPU) {
	value := cpu.PopStack()
	cpu.D, cpu.E = split(value)
	cpu.Cycles += 12
}
func (POP_D1) Code() uint8 {
	return 0xD1
}
func (POP_D1) String() string {
	return "POP DE"
}

// JP C,a16    code=0xda
type JP_DA struct{}

func (JP_DA) Exec(cpu *CPU) {
	panic("TODO JP_DA")
}
func (JP_DA) Code() uint8 {
	return 0xDA
}
func (JP_DA) String() string {
	return "JP C,a16"
}

// LD (BC),A    code=0x02
type LD_02 struct{}

func (LD_02) Exec(cpu *CPU) {

	data := cpu.A

	cpu.WriteMemory(cpu.BC(), data)

	cpu.Cycles += 8

}
func (LD_02) Code() uint8 {
	return 0x2
}
func (LD_02) String() string {
	return "LD (BC),A"
}

// SCF     code=0x37
type SCF_37 struct{}

func (SCF_37) Exec(cpu *CPU) {
	panic("TODO SCF_37")
}
func (SCF_37) Code() uint8 {
	return 0x37
}
func (SCF_37) String() string {
	return "SCF"
}

// RST $30    code=0xf7
type RST_F7 struct{}

func (RST_F7) Exec(cpu *CPU) {
	n := uint8(0x30)
	cpu.PushStack(cpu.PC)
	cpu.PC = concatU16(0x00, n)
}
func (RST_F7) Code() uint8 {
	return 0xF7
}
func (RST_F7) String() string {
	return "RST $30"
}

// AND A,n8    code=0xe6
type AND_E6 struct{}

func (AND_E6) Exec(cpu *CPU) {
	panic("TODO AND_E6")
}
func (AND_E6) Code() uint8 {
	return 0xE6
}
func (AND_E6) String() string {
	return "AND A,n8"
}

// ILLEGAL_EB     code=0xeb
type ILLEGAL_EB_EB struct{}

func (ILLEGAL_EB_EB) Exec(cpu *CPU) {
	panic("TODO ILLEGAL_EB_EB")
}
func (ILLEGAL_EB_EB) Code() uint8 {
	return 0xEB
}
func (ILLEGAL_EB_EB) String() string {
	return "ILLEGAL_EB"
}

// INC BC    code=0x03
type INC_03 struct{}

func (INC_03) Exec(cpu *CPU) {
	res, flags := add(cpu.BC(), 0x01)
	cpu.F = flags
	cpu.B, cpu.C = split(res)
	cpu.Cycles += 8
}
func (INC_03) Code() uint8 {
	return 0x3
}
func (INC_03) String() string {
	return "INC BC"
}

// LD B,C    code=0x41
type LD_41 struct{}

func (LD_41) Exec(cpu *CPU) {

	data := cpu.C

	cpu.B = data

	cpu.Cycles += 4

}
func (LD_41) Code() uint8 {
	return 0x41
}
func (LD_41) String() string {
	return "LD B,C"
}

// LD C,B    code=0x48
type LD_48 struct{}

func (LD_48) Exec(cpu *CPU) {

	data := cpu.B

	cpu.C = data

	cpu.Cycles += 4

}
func (LD_48) Code() uint8 {
	return 0x48
}
func (LD_48) String() string {
	return "LD C,B"
}

// LD A,L    code=0x7d
type LD_7D struct{}

func (LD_7D) Exec(cpu *CPU) {

	data := cpu.L

	cpu.A = data

	cpu.Cycles += 4

}
func (LD_7D) Code() uint8 {
	return 0x7D
}
func (LD_7D) String() string {
	return "LD A,L"
}

// XOR A,C    code=0xa9
type XOR_A9 struct{}

func (XOR_A9) Exec(cpu *CPU) {
	res := cpu.A ^ cpu.C

	var flags Flags
	if res == 0 {
		flags |= FLAGZ
	}
	cpu.F = FlagRegister(flags)
	cpu.A = res
	cpu.Cycles += 4
}
func (XOR_A9) Code() uint8 {
	return 0xA9
}
func (XOR_A9) String() string {
	return "XOR A,C"
}

// CP A,E    code=0xbb
type CP_BB struct{}

func (CP_BB) Exec(cpu *CPU) {
	_, flags := sub(cpu.A, cpu.E)

	cpu.F = flags
	cpu.Cycles += 4
}
func (CP_BB) Code() uint8 {
	return 0xBB
}
func (CP_BB) String() string {
	return "CP A,E"
}

// ADC A,B    code=0x88
type ADC_88 struct{}

func (ADC_88) Exec(cpu *CPU) {
	panic("TODO ADC_88")
}
func (ADC_88) Code() uint8 {
	return 0x88
}
func (ADC_88) String() string {
	return "ADC A,B"
}

// LD A,(DE)    code=0x1a
type LD_1A struct{}

func (LD_1A) Exec(cpu *CPU) {

	data := cpu.loadU8(cpu.DE())

	cpu.A = data

	cpu.Cycles += 8

}
func (LD_1A) Code() uint8 {
	return 0x1A
}
func (LD_1A) String() string {
	return "LD A,(DE)"
}

// SUB A,D    code=0x92
type SUB_92 struct{}

func (SUB_92) Exec(cpu *CPU) {
	res, flags := sub(cpu.A, cpu.D)
	cpu.A = res

	cpu.F = flags
	cpu.Cycles += 4
}
func (SUB_92) Code() uint8 {
	return 0x92
}
func (SUB_92) String() string {
	return "SUB A,D"
}

// CALL C,a16    code=0xdc
type CALL_DC struct{}

func (CALL_DC) Exec(cpu *CPU) {
	lsb := cpu.readU8(cpu.PC)
	cpu.IncProgramCounter()
	msb := cpu.readU8(cpu.PC)
	cpu.IncProgramCounter()
	nn := concatU16(msb, lsb)
	if cpu.F.HasCarry() {
		cpu.PushStack(cpu.PC)
		cpu.PC = nn
		cpu.Cycles += 24
	} else {
		cpu.Cycles += 12
	}
}
func (CALL_DC) Code() uint8 {
	return 0xDC
}
func (CALL_DC) String() string {
	return "CALL C,a16"
}

// AND A,A    code=0xa7
type AND_A7 struct{}

func (AND_A7) Exec(cpu *CPU) {
	panic("TODO AND_A7")
}
func (AND_A7) Code() uint8 {
	return 0xA7
}
func (AND_A7) String() string {
	return "AND A,A"
}

// RST $18    code=0xdf
type RST_DF struct{}

func (RST_DF) Exec(cpu *CPU) {
	n := uint8(0x18)
	cpu.PushStack(cpu.PC)
	cpu.PC = concatU16(0x00, n)
}
func (RST_DF) Code() uint8 {
	return 0xDF
}
func (RST_DF) String() string {
	return "RST $18"
}

// LD (HL),B    code=0x70
type LD_70 struct{}

func (LD_70) Exec(cpu *CPU) {

	data := cpu.B

	cpu.WriteMemory(cpu.HL(), data)

	cpu.Cycles += 8

}
func (LD_70) Code() uint8 {
	return 0x70
}
func (LD_70) String() string {
	return "LD (HL),B"
}

// XOR A,(HL)    code=0xae
type XOR_AE struct{}

func (XOR_AE) Exec(cpu *CPU) {
	res := cpu.A ^ cpu.loadU8(cpu.HL())

	var flags Flags
	if res == 0 {
		flags |= FLAGZ
	}
	cpu.F = FlagRegister(flags)
	cpu.A = res
	cpu.Cycles += 8
}
func (XOR_AE) Code() uint8 {
	return 0xAE
}
func (XOR_AE) String() string {
	return "XOR A,(HL)"
}

// RET Z    code=0xc8
type RET_C8 struct{}

func (RET_C8) Exec(cpu *CPU) {
	if cpu.F.HasZero() {
		cpu.PC = cpu.PopStack()
		cpu.Cycles += 20
	} else {
		cpu.Cycles += 8
	}
}
func (RET_C8) Code() uint8 {
	return 0xC8
}
func (RET_C8) String() string {
	return "RET Z"
}

// RLCA     code=0x07
type RLCA_07 struct{}

func (RLCA_07) Exec(cpu *CPU) {
	cpu.A, cpu.F = rotate(cpu.A, 0, cpu.F, true)
	cpu.Cycles += 4
}
func (RLCA_07) Code() uint8 {
	return 0x7
}
func (RLCA_07) String() string {
	return "RLCA"
}

// RLA     code=0x17
type RLA_17 struct{}

func (RLA_17) Exec(cpu *CPU) {
	cpu.A, cpu.F = rotate(cpu.A, 0, cpu.F, false)
	// RLA always sets the zero flag to 0 without looking at the resulting value of the calculation.
	cpu.F &= ^FLAGZ
	cpu.Cycles += 4
}
func (RLA_17) Code() uint8 {
	return 0x17
}
func (RLA_17) String() string {
	return "RLA"
}

// LD B,B    code=0x40
type LD_40 struct{}

func (LD_40) Exec(cpu *CPU) {

	data := cpu.B

	cpu.B = data

	cpu.Cycles += 4

}
func (LD_40) Code() uint8 {
	return 0x40
}
func (LD_40) String() string {
	return "LD B,B"
}

// CP A,n8    code=0xfe
type CP_FE struct{}

func (CP_FE) Exec(cpu *CPU) {
	_, flags := sub(cpu.A, cpu.readU8(cpu.PC))
	cpu.IncProgramCounter()
	cpu.F = flags
	cpu.Cycles += 8
}
func (CP_FE) Code() uint8 {
	return 0xFE
}
func (CP_FE) String() string {
	return "CP A,n8"
}

// ADD A,(HL)    code=0x86
type ADD_86 struct{}

func (ADD_86) Exec(cpu *CPU) {
	lhs := cpu.A
	rhs := cpu.loadU8(cpu.HL())

	res, flags := add(lhs, rhs)
	cpu.A = res
	cpu.F = flags
	cpu.Cycles += 8
}
func (ADD_86) Code() uint8 {
	return 0x86
}
func (ADD_86) String() string {
	return "ADD A,(HL)"
}

// SUB A,(HL)    code=0x96
type SUB_96 struct{}

func (SUB_96) Exec(cpu *CPU) {
	res, flags := sub(cpu.A, cpu.loadU8(cpu.HL()))
	cpu.A = res

	cpu.F = flags
	cpu.Cycles += 8
}
func (SUB_96) Code() uint8 {
	return 0x96
}
func (SUB_96) String() string {
	return "SUB A,(HL)"
}

// AND A,B    code=0xa0
type AND_A0 struct{}

func (AND_A0) Exec(cpu *CPU) {
	panic("TODO AND_A0")
}
func (AND_A0) Code() uint8 {
	return 0xA0
}
func (AND_A0) String() string {
	return "AND A,B"
}

// CALL NC,a16    code=0xd4
type CALL_D4 struct{}

func (CALL_D4) Exec(cpu *CPU) {
	lsb := cpu.readU8(cpu.PC)
	cpu.IncProgramCounter()
	msb := cpu.readU8(cpu.PC)
	cpu.IncProgramCounter()
	nn := concatU16(msb, lsb)
	if !cpu.F.HasCarry() {
		cpu.PushStack(cpu.PC)
		cpu.PC = nn
		cpu.Cycles += 24
	} else {
		cpu.Cycles += 12
	}
}
func (CALL_D4) Code() uint8 {
	return 0xD4
}
func (CALL_D4) String() string {
	return "CALL NC,a16"
}

// RETI     code=0xd9
type RETI_D9 struct{}

func (RETI_D9) Exec(cpu *CPU) {
	panic("TODO RETI_D9")
}
func (RETI_D9) Code() uint8 {
	return 0xD9
}
func (RETI_D9) String() string {
	return "RETI"
}

// ADD HL,BC    code=0x09
type ADD_09 struct{}

func (ADD_09) Exec(cpu *CPU) {
	lhs := cpu.HL()
	rhs := cpu.BC()

	res, flags := add(lhs, rhs)
	cpu.H, cpu.L = split(res)
	cpu.F = flags
	cpu.Cycles += 8
}
func (ADD_09) Code() uint8 {
	return 0x9
}
func (ADD_09) String() string {
	return "ADD HL,BC"
}

// INC C    code=0x0c
type INC_0C struct{}

func (INC_0C) Exec(cpu *CPU) {
	res, flags := add(cpu.C, 0x01)
	cpu.F = flags
	cpu.C = res
	cpu.Cycles += 4
}
func (INC_0C) Code() uint8 {
	return 0xC
}
func (INC_0C) String() string {
	return "INC C"
}

// INC A    code=0x3c
type INC_3C struct{}

func (INC_3C) Exec(cpu *CPU) {
	res, flags := add(cpu.A, 0x01)
	cpu.F = flags
	cpu.A = res
	cpu.Cycles += 4
}
func (INC_3C) Code() uint8 {
	return 0x3C
}
func (INC_3C) String() string {
	return "INC A"
}

// LD H,H    code=0x64
type LD_64 struct{}

func (LD_64) Exec(cpu *CPU) {

	data := cpu.H

	cpu.H = data

	cpu.Cycles += 4

}
func (LD_64) Code() uint8 {
	return 0x64
}
func (LD_64) String() string {
	return "LD H,H"
}

// SBC A,n8    code=0xde
type SBC_DE struct{}

func (SBC_DE) Exec(cpu *CPU) {
	panic("TODO SBC_DE")
}
func (SBC_DE) Code() uint8 {
	return 0xDE
}
func (SBC_DE) String() string {
	return "SBC A,n8"
}

// EI     code=0xfb
type EI_FB struct{}

func (EI_FB) Exec(cpu *CPU) {
	panic("TODO EI_FB")
}
func (EI_FB) Code() uint8 {
	return 0xFB
}
func (EI_FB) String() string {
	return "EI"
}

// ILLEGAL_FC     code=0xfc
type ILLEGAL_FC_FC struct{}

func (ILLEGAL_FC_FC) Exec(cpu *CPU) {
	panic("TODO ILLEGAL_FC_FC")
}
func (ILLEGAL_FC_FC) Code() uint8 {
	return 0xFC
}
func (ILLEGAL_FC_FC) String() string {
	return "ILLEGAL_FC"
}

// LD D,D    code=0x52
type LD_52 struct{}

func (LD_52) Exec(cpu *CPU) {

	data := cpu.D

	cpu.D = data

	cpu.Cycles += 4

}
func (LD_52) Code() uint8 {
	return 0x52
}
func (LD_52) String() string {
	return "LD D,D"
}

// CALL NZ,a16    code=0xc4
type CALL_C4 struct{}

func (CALL_C4) Exec(cpu *CPU) {
	lsb := cpu.readU8(cpu.PC)
	cpu.IncProgramCounter()
	msb := cpu.readU8(cpu.PC)
	cpu.IncProgramCounter()
	nn := concatU16(msb, lsb)
	if !cpu.F.HasZero() {
		cpu.PushStack(cpu.PC)
		cpu.PC = nn
		cpu.Cycles += 24
	} else {
		cpu.Cycles += 12
	}
}
func (CALL_C4) Code() uint8 {
	return 0xC4
}
func (CALL_C4) String() string {
	return "CALL NZ,a16"
}

// RET C    code=0xd8
type RET_D8 struct{}

func (RET_D8) Exec(cpu *CPU) {
	if cpu.F.HasCarry() {
		cpu.PC = cpu.PopStack()
		cpu.Cycles += 20
	} else {
		cpu.Cycles += 8
	}
}
func (RET_D8) Code() uint8 {
	return 0xD8
}
func (RET_D8) String() string {
	return "RET C"
}

// ILLEGAL_DD     code=0xdd
type ILLEGAL_DD_DD struct{}

func (ILLEGAL_DD_DD) Exec(cpu *CPU) {
	panic("TODO ILLEGAL_DD_DD")
}
func (ILLEGAL_DD_DD) Code() uint8 {
	return 0xDD
}
func (ILLEGAL_DD_DD) String() string {
	return "ILLEGAL_DD"
}

// LD BC,n16    code=0x01
type LD_01 struct{}

func (LD_01) Exec(cpu *CPU) {

	data := cpu.readU16(cpu.PC)
	cpu.IncProgramCounter()
	cpu.IncProgramCounter()

	cpu.B, cpu.C = split(data)

	cpu.Cycles += 12

}
func (LD_01) Code() uint8 {
	return 0x1
}
func (LD_01) String() string {
	return "LD BC,n16"
}

// INC H    code=0x24
type INC_24 struct{}

func (INC_24) Exec(cpu *CPU) {
	res, flags := add(cpu.H, 0x01)
	cpu.F = flags
	cpu.H = res
	cpu.Cycles += 4
}
func (INC_24) Code() uint8 {
	return 0x24
}
func (INC_24) String() string {
	return "INC H"
}

// SUB A,B    code=0x90
type SUB_90 struct{}

func (SUB_90) Exec(cpu *CPU) {
	res, flags := sub(cpu.A, cpu.B)
	cpu.A = res

	cpu.F = flags
	cpu.Cycles += 4
}
func (SUB_90) Code() uint8 {
	return 0x90
}
func (SUB_90) String() string {
	return "SUB A,B"
}

// POP BC    code=0xc1
type POP_C1 struct{}

func (POP_C1) Exec(cpu *CPU) {
	value := cpu.PopStack()
	cpu.B, cpu.C = split(value)
	cpu.Cycles += 12
}
func (POP_C1) Code() uint8 {
	return 0xC1
}
func (POP_C1) String() string {
	return "POP BC"
}

// RST $08    code=0xcf
type RST_CF struct{}

func (RST_CF) Exec(cpu *CPU) {
	n := uint8(0x8)
	cpu.PushStack(cpu.PC)
	cpu.PC = concatU16(0x00, n)
}
func (RST_CF) Code() uint8 {
	return 0xCF
}
func (RST_CF) String() string {
	return "RST $08"
}

// ADD SP,e8    code=0xe8
type ADD_E8 struct{}

func (ADD_E8) Exec(cpu *CPU) {
	lhs := cpu.SP
	rhs := cpu.readI8(cpu.PC)

	cpu.IncProgramCounter()
	res, flags := add(lhs, rhs)
	cpu.SP = res
	cpu.F = flags
	cpu.Cycles += 16
}
func (ADD_E8) Code() uint8 {
	return 0xE8
}
func (ADD_E8) String() string {
	return "ADD SP,e8"
}

// ADC A,A    code=0x8f
type ADC_8F struct{}

func (ADC_8F) Exec(cpu *CPU) {
	panic("TODO ADC_8F")
}
func (ADC_8F) Code() uint8 {
	return 0x8F
}
func (ADC_8F) String() string {
	return "ADC A,A"
}

// LDH A,(a8)    code=0xf0
type LDH_F0 struct{}

func (LDH_F0) Exec(cpu *CPU) {
	pc0 := cpu.PC
	value := cpu.loadU8(concatU16(0xFF, cpu.readU8(cpu.PC)))
	cpu.A = value
	cpu.IncProgramCounter()
	cpu.PC = pc0 + 1
	cpu.Cycles += 12
}
func (LDH_F0) Code() uint8 {
	return 0xF0
}
func (LDH_F0) String() string {
	return "LDH A,(a8)"
}

// DEC (HL)    code=0x35
type DEC_35 struct{}

func (DEC_35) Exec(cpu *CPU) {
	res, flags := sub(cpu.loadU8(cpu.HL()), 0x01)
	cpu.F = flags
	cpu.WriteMemory(cpu.HL(), res)
	cpu.Cycles += 12
}
func (DEC_35) Code() uint8 {
	return 0x35
}
func (DEC_35) String() string {
	return "DEC (HL)"
}

// LD (HL),D    code=0x72
type LD_72 struct{}

func (LD_72) Exec(cpu *CPU) {

	data := cpu.D

	cpu.WriteMemory(cpu.HL(), data)

	cpu.Cycles += 8

}
func (LD_72) Code() uint8 {
	return 0x72
}
func (LD_72) String() string {
	return "LD (HL),D"
}

// SBC A,D    code=0x9a
type SBC_9A struct{}

func (SBC_9A) Exec(cpu *CPU) {
	panic("TODO SBC_9A")
}
func (SBC_9A) Code() uint8 {
	return 0x9A
}
func (SBC_9A) String() string {
	return "SBC A,D"
}

// AND A,(HL)    code=0xa6
type AND_A6 struct{}

func (AND_A6) Exec(cpu *CPU) {
	panic("TODO AND_A6")
}
func (AND_A6) Code() uint8 {
	return 0xA6
}
func (AND_A6) String() string {
	return "AND A,(HL)"
}

// ILLEGAL_ED     code=0xed
type ILLEGAL_ED_ED struct{}

func (ILLEGAL_ED_ED) Exec(cpu *CPU) {
	panic("TODO ILLEGAL_ED_ED")
}
func (ILLEGAL_ED_ED) Code() uint8 {
	return 0xED
}
func (ILLEGAL_ED_ED) String() string {
	return "ILLEGAL_ED"
}

// ILLEGAL_F4     code=0xf4
type ILLEGAL_F4_F4 struct{}

func (ILLEGAL_F4_F4) Exec(cpu *CPU) {
	panic("TODO ILLEGAL_F4_F4")
}
func (ILLEGAL_F4_F4) Code() uint8 {
	return 0xF4
}
func (ILLEGAL_F4_F4) String() string {
	return "ILLEGAL_F4"
}

// LD (a16),SP    code=0x08
type LD_08 struct{}

func (LD_08) Exec(cpu *CPU) {

	data := cpu.SP

	cpu.WriteMemory(cpu.readU16(cpu.PC), data)
	cpu.IncProgramCounter()
	cpu.IncProgramCounter()

	cpu.Cycles += 20

}
func (LD_08) Code() uint8 {
	return 0x8
}
func (LD_08) String() string {
	return "LD (a16),SP"
}

// INC DE    code=0x13
type INC_13 struct{}

func (INC_13) Exec(cpu *CPU) {
	res, flags := add(cpu.DE(), 0x01)
	cpu.F = flags
	cpu.D, cpu.E = split(res)
	cpu.Cycles += 8
}
func (INC_13) Code() uint8 {
	return 0x13
}
func (INC_13) String() string {
	return "INC DE"
}

// LD A,(BC)    code=0x0a
type LD_0A struct{}

func (LD_0A) Exec(cpu *CPU) {

	data := cpu.loadU8(cpu.BC())

	cpu.A = data

	cpu.Cycles += 8

}
func (LD_0A) Code() uint8 {
	return 0xA
}
func (LD_0A) String() string {
	return "LD A,(BC)"
}

// LD C,n8    code=0x0e
type LD_0E struct{}

func (LD_0E) Exec(cpu *CPU) {

	data := cpu.readU8(cpu.PC)
	cpu.IncProgramCounter()

	cpu.C = data

	cpu.Cycles += 8

}
func (LD_0E) Code() uint8 {
	return 0xE
}
func (LD_0E) String() string {
	return "LD C,n8"
}

// DAA     code=0x27
type DAA_27 struct{}

func (DAA_27) Exec(cpu *CPU) {
	panic("TODO DAA_27")
}
func (DAA_27) Code() uint8 {
	return 0x27
}
func (DAA_27) String() string {
	return "DAA"
}

// DEC HL    code=0x2b
type DEC_2B struct{}

func (DEC_2B) Exec(cpu *CPU) {
	res, flags := sub(cpu.HL(), 0x01)
	cpu.F = flags
	cpu.H, cpu.L = split(res)
	cpu.Cycles += 8
}
func (DEC_2B) Code() uint8 {
	return 0x2B
}
func (DEC_2B) String() string {
	return "DEC HL"
}

// LD A,(HL-)    code=0x3a
type LD_3A struct{}

func (LD_3A) Exec(cpu *CPU) {

	data := cpu.loadU8(cpu.HL())

	cpu.A = data

	cpu.Cycles += 8

}
func (LD_3A) Code() uint8 {
	return 0x3A
}
func (LD_3A) String() string {
	return "LD A,(HL-)"
}

// LD C,L    code=0x4d
type LD_4D struct{}

func (LD_4D) Exec(cpu *CPU) {

	data := cpu.L

	cpu.C = data

	cpu.Cycles += 4

}
func (LD_4D) Code() uint8 {
	return 0x4D
}
func (LD_4D) String() string {
	return "LD C,L"
}

// ILLEGAL_FD     code=0xfd
type ILLEGAL_FD_FD struct{}

func (ILLEGAL_FD_FD) Exec(cpu *CPU) {
	panic("TODO ILLEGAL_FD_FD")
}
func (ILLEGAL_FD_FD) Code() uint8 {
	return 0xFD
}
func (ILLEGAL_FD_FD) String() string {
	return "ILLEGAL_FD"
}

// LD HL,n16    code=0x21
type LD_21 struct{}

func (LD_21) Exec(cpu *CPU) {

	data := cpu.readU16(cpu.PC)
	cpu.IncProgramCounter()
	cpu.IncProgramCounter()

	cpu.H, cpu.L = split(data)

	cpu.Cycles += 12

}
func (LD_21) Code() uint8 {
	return 0x21
}
func (LD_21) String() string {
	return "LD HL,n16"
}

// LD H,n8    code=0x26
type LD_26 struct{}

func (LD_26) Exec(cpu *CPU) {

	data := cpu.readU8(cpu.PC)
	cpu.IncProgramCounter()

	cpu.H = data

	cpu.Cycles += 8

}
func (LD_26) Code() uint8 {
	return 0x26
}
func (LD_26) String() string {
	return "LD H,n8"
}

// ADD A,A    code=0x87
type ADD_87 struct{}

func (ADD_87) Exec(cpu *CPU) {
	lhs := cpu.A
	rhs := cpu.A

	res, flags := add(lhs, rhs)
	cpu.A = res
	cpu.F = flags
	cpu.Cycles += 4
}
func (ADD_87) Code() uint8 {
	return 0x87
}
func (ADD_87) String() string {
	return "ADD A,A"
}

// CP A,C    code=0xb9
type CP_B9 struct{}

func (CP_B9) Exec(cpu *CPU) {
	_, flags := sub(cpu.A, cpu.C)

	cpu.F = flags
	cpu.Cycles += 4
}
func (CP_B9) Code() uint8 {
	return 0xB9
}
func (CP_B9) String() string {
	return "CP A,C"
}

// CP A,(HL)    code=0xbe
type CP_BE struct{}

func (CP_BE) Exec(cpu *CPU) {
	_, flags := sub(cpu.A, cpu.loadU8(cpu.HL()))

	cpu.F = flags
	cpu.Cycles += 8
}
func (CP_BE) Code() uint8 {
	return 0xBE
}
func (CP_BE) String() string {
	return "CP A,(HL)"
}

// RST $20    code=0xe7
type RST_E7 struct{}

func (RST_E7) Exec(cpu *CPU) {
	n := uint8(0x20)
	cpu.PushStack(cpu.PC)
	cpu.PC = concatU16(0x00, n)
}
func (RST_E7) Code() uint8 {
	return 0xE7
}
func (RST_E7) String() string {
	return "RST $20"
}

// LD E,D    code=0x5a
type LD_5A struct{}

func (LD_5A) Exec(cpu *CPU) {

	data := cpu.D

	cpu.E = data

	cpu.Cycles += 4

}
func (LD_5A) Code() uint8 {
	return 0x5A
}
func (LD_5A) String() string {
	return "LD E,D"
}

// ILLEGAL_E3     code=0xe3
type ILLEGAL_E3_E3 struct{}

func (ILLEGAL_E3_E3) Exec(cpu *CPU) {
	panic("TODO ILLEGAL_E3_E3")
}
func (ILLEGAL_E3_E3) Code() uint8 {
	return 0xE3
}
func (ILLEGAL_E3_E3) String() string {
	return "ILLEGAL_E3"
}

// JR NZ,e8    code=0x20
type JR_20 struct{}

func (JR_20) Exec(cpu *CPU) {
	e := cpu.readI8(cpu.PC)
	cpu.IncProgramCounter()
	if !cpu.F.HasZero() {
		cpu.PC, cpu.F = add(cpu.PC, e)
		cpu.Cycles += 12
	} else {
		cpu.Cycles += 8
	}
}
func (JR_20) Code() uint8 {
	return 0x20
}
func (JR_20) String() string {
	return "JR NZ,e8"
}

// LD C,H    code=0x4c
type LD_4C struct{}

func (LD_4C) Exec(cpu *CPU) {

	data := cpu.H

	cpu.C = data

	cpu.Cycles += 4

}
func (LD_4C) Code() uint8 {
	return 0x4C
}
func (LD_4C) String() string {
	return "LD C,H"
}

// LD (HL-),A    code=0x32
type LD_32 struct{}

func (LD_32) Exec(cpu *CPU) {

	data := cpu.A

	cpu.WriteMemory(cpu.HL(), data)

	decr, flags := sub(cpu.HL(), 0x01)
	cpu.H, cpu.L = split(decr)
	cpu.F = flags

	cpu.Cycles += 8

}
func (LD_32) Code() uint8 {
	return 0x32
}
func (LD_32) String() string {
	return "LD (HL-),A"
}

// INC (HL)    code=0x34
type INC_34 struct{}

func (INC_34) Exec(cpu *CPU) {
	res, flags := add(cpu.loadU8(cpu.HL()), 0x01)
	cpu.F = flags
	cpu.WriteMemory(cpu.HL(), res)
	cpu.Cycles += 12
}
func (INC_34) Code() uint8 {
	return 0x34
}
func (INC_34) String() string {
	return "INC (HL)"
}

// LDH (a8),A    code=0xe0
type LDH_E0 struct{}

func (LDH_E0) Exec(cpu *CPU) {
	pc0 := cpu.PC
	value := cpu.A
	cpu.WriteMemory(concatU16(0xFF, cpu.readU8(cpu.PC)), value)

	cpu.PC = pc0 + 1
	cpu.Cycles += 12
}
func (LDH_E0) Code() uint8 {
	return 0xE0
}
func (LDH_E0) String() string {
	return "LDH (a8),A"
}

// LDH A,(C)    code=0xf2
type LDH_F2 struct{}

func (LDH_F2) Exec(cpu *CPU) {
	pc0 := cpu.PC
	value := cpu.loadU8(concatU16(0xFF, cpu.C))
	cpu.A = value

	cpu.PC = pc0 + 1
	cpu.Cycles += 8
}
func (LDH_F2) Code() uint8 {
	return 0xF2
}
func (LDH_F2) String() string {
	return "LDH A,(C)"
}

// JR NC,e8    code=0x30
type JR_30 struct{}

func (JR_30) Exec(cpu *CPU) {
	e := cpu.readI8(cpu.PC)
	cpu.IncProgramCounter()
	if !cpu.F.HasCarry() {
		cpu.PC, cpu.F = add(cpu.PC, e)
		cpu.Cycles += 12
	} else {
		cpu.Cycles += 8
	}
}
func (JR_30) Code() uint8 {
	return 0x30
}
func (JR_30) String() string {
	return "JR NC,e8"
}

// LD H,E    code=0x63
type LD_63 struct{}

func (LD_63) Exec(cpu *CPU) {

	data := cpu.E

	cpu.H = data

	cpu.Cycles += 4

}
func (LD_63) Code() uint8 {
	return 0x63
}
func (LD_63) String() string {
	return "LD H,E"
}

// ADC A,H    code=0x8c
type ADC_8C struct{}

func (ADC_8C) Exec(cpu *CPU) {
	panic("TODO ADC_8C")
}
func (ADC_8C) Code() uint8 {
	return 0x8C
}
func (ADC_8C) String() string {
	return "ADC A,H"
}

// AND A,C    code=0xa1
type AND_A1 struct{}

func (AND_A1) Exec(cpu *CPU) {
	panic("TODO AND_A1")
}
func (AND_A1) Code() uint8 {
	return 0xA1
}
func (AND_A1) String() string {
	return "AND A,C"
}

// LD L,E    code=0x6b
type LD_6B struct{}

func (LD_6B) Exec(cpu *CPU) {

	data := cpu.E

	cpu.L = data

	cpu.Cycles += 4

}
func (LD_6B) Code() uint8 {
	return 0x6B
}
func (LD_6B) String() string {
	return "LD L,E"
}

// INC SP    code=0x33
type INC_33 struct{}

func (INC_33) Exec(cpu *CPU) {
	res, flags := add(cpu.SP, 0x01)
	cpu.F = flags
	cpu.SP = res
	cpu.Cycles += 8
}
func (INC_33) Code() uint8 {
	return 0x33
}
func (INC_33) String() string {
	return "INC SP"
}

// LD D,C    code=0x51
type LD_51 struct{}

func (LD_51) Exec(cpu *CPU) {

	data := cpu.C

	cpu.D = data

	cpu.Cycles += 4

}
func (LD_51) Code() uint8 {
	return 0x51
}
func (LD_51) String() string {
	return "LD D,C"
}

// HALT     code=0x76
type HALT_76 struct{}

func (HALT_76) Exec(cpu *CPU) {
	panic("TODO HALT_76")
}
func (HALT_76) Code() uint8 {
	return 0x76
}
func (HALT_76) String() string {
	return "HALT"
}

// ADD A,C    code=0x81
type ADD_81 struct{}

func (ADD_81) Exec(cpu *CPU) {
	lhs := cpu.A
	rhs := cpu.C

	res, flags := add(lhs, rhs)
	cpu.A = res
	cpu.F = flags
	cpu.Cycles += 4
}
func (ADD_81) Code() uint8 {
	return 0x81
}
func (ADD_81) String() string {
	return "ADD A,C"
}

// ADD A,D    code=0x82
type ADD_82 struct{}

func (ADD_82) Exec(cpu *CPU) {
	lhs := cpu.A
	rhs := cpu.D

	res, flags := add(lhs, rhs)
	cpu.A = res
	cpu.F = flags
	cpu.Cycles += 4
}
func (ADD_82) Code() uint8 {
	return 0x82
}
func (ADD_82) String() string {
	return "ADD A,D"
}

// LD B,D    code=0x42
type LD_42 struct{}

func (LD_42) Exec(cpu *CPU) {

	data := cpu.D

	cpu.B = data

	cpu.Cycles += 4

}
func (LD_42) Code() uint8 {
	return 0x42
}
func (LD_42) String() string {
	return "LD B,D"
}

// LD C,(HL)    code=0x4e
type LD_4E struct{}

func (LD_4E) Exec(cpu *CPU) {

	data := cpu.loadU8(cpu.HL())

	cpu.C = data

	cpu.Cycles += 8

}
func (LD_4E) Code() uint8 {
	return 0x4E
}
func (LD_4E) String() string {
	return "LD C,(HL)"
}

// LD (HL),H    code=0x74
type LD_74 struct{}

func (LD_74) Exec(cpu *CPU) {

	data := cpu.H

	cpu.WriteMemory(cpu.HL(), data)

	cpu.Cycles += 8

}
func (LD_74) Code() uint8 {
	return 0x74
}
func (LD_74) String() string {
	return "LD (HL),H"
}

// ADC A,D    code=0x8a
type ADC_8A struct{}

func (ADC_8A) Exec(cpu *CPU) {
	panic("TODO ADC_8A")
}
func (ADC_8A) Code() uint8 {
	return 0x8A
}
func (ADC_8A) String() string {
	return "ADC A,D"
}

// INC B    code=0x04
type INC_04 struct{}

func (INC_04) Exec(cpu *CPU) {
	res, flags := add(cpu.B, 0x01)
	cpu.F = flags
	cpu.B = res
	cpu.Cycles += 4
}
func (INC_04) Code() uint8 {
	return 0x4
}
func (INC_04) String() string {
	return "INC B"
}

// LD (HL),A    code=0x77
type LD_77 struct{}

func (LD_77) Exec(cpu *CPU) {

	data := cpu.A

	cpu.WriteMemory(cpu.HL(), data)

	cpu.Cycles += 8

}
func (LD_77) Code() uint8 {
	return 0x77
}
func (LD_77) String() string {
	return "LD (HL),A"
}

// XOR A,E    code=0xab
type XOR_AB struct{}

func (XOR_AB) Exec(cpu *CPU) {
	res := cpu.A ^ cpu.E

	var flags Flags
	if res == 0 {
		flags |= FLAGZ
	}
	cpu.F = FlagRegister(flags)
	cpu.A = res
	cpu.Cycles += 4
}
func (XOR_AB) Code() uint8 {
	return 0xAB
}
func (XOR_AB) String() string {
	return "XOR A,E"
}

// CP A,B    code=0xb8
type CP_B8 struct{}

func (CP_B8) Exec(cpu *CPU) {
	_, flags := sub(cpu.A, cpu.B)

	cpu.F = flags
	cpu.Cycles += 4
}
func (CP_B8) Code() uint8 {
	return 0xB8
}
func (CP_B8) String() string {
	return "CP A,B"
}

// RET NC    code=0xd0
type RET_D0 struct{}

func (RET_D0) Exec(cpu *CPU) {
	if !cpu.F.HasCarry() {
		cpu.PC = cpu.PopStack()
		cpu.Cycles += 20
	} else {
		cpu.Cycles += 8
	}
}
func (RET_D0) Code() uint8 {
	return 0xD0
}
func (RET_D0) String() string {
	return "RET NC"
}

// OR A,C    code=0xb1
type OR_B1 struct{}

func (OR_B1) Exec(cpu *CPU) {
	res := cpu.A | cpu.C
	var flags Flags
	if res == 0 {
		flags |= FLAGZ
	}
	cpu.F = FlagRegister(flags)
	cpu.A = res
	cpu.Cycles += 4
}
func (OR_B1) Code() uint8 {
	return 0xB1
}
func (OR_B1) String() string {
	return "OR A,C"
}

// LD L,H    code=0x6c
type LD_6C struct{}

func (LD_6C) Exec(cpu *CPU) {

	data := cpu.H

	cpu.L = data

	cpu.Cycles += 4

}
func (LD_6C) Code() uint8 {
	return 0x6C
}
func (LD_6C) String() string {
	return "LD L,H"
}

// DEC L    code=0x2d
type DEC_2D struct{}

func (DEC_2D) Exec(cpu *CPU) {
	res, flags := sub(cpu.L, 0x01)
	cpu.F = flags
	cpu.L = res
	cpu.Cycles += 4
}
func (DEC_2D) Code() uint8 {
	return 0x2D
}
func (DEC_2D) String() string {
	return "DEC L"
}

// LD D,B    code=0x50
type LD_50 struct{}

func (LD_50) Exec(cpu *CPU) {

	data := cpu.B

	cpu.D = data

	cpu.Cycles += 4

}
func (LD_50) Code() uint8 {
	return 0x50
}
func (LD_50) String() string {
	return "LD D,B"
}

// LD A,A    code=0x7f
type LD_7F struct{}

func (LD_7F) Exec(cpu *CPU) {

	data := cpu.A

	cpu.A = data

	cpu.Cycles += 4

}
func (LD_7F) Code() uint8 {
	return 0x7F
}
func (LD_7F) String() string {
	return "LD A,A"
}

// SUB A,L    code=0x95
type SUB_95 struct{}

func (SUB_95) Exec(cpu *CPU) {
	res, flags := sub(cpu.A, cpu.L)
	cpu.A = res

	cpu.F = flags
	cpu.Cycles += 4
}
func (SUB_95) Code() uint8 {
	return 0x95
}
func (SUB_95) String() string {
	return "SUB A,L"
}

// SBC A,L    code=0x9d
type SBC_9D struct{}

func (SBC_9D) Exec(cpu *CPU) {
	panic("TODO SBC_9D")
}
func (SBC_9D) Code() uint8 {
	return 0x9D
}
func (SBC_9D) String() string {
	return "SBC A,L"
}

// SBC A,A    code=0x9f
type SBC_9F struct{}

func (SBC_9F) Exec(cpu *CPU) {
	panic("TODO SBC_9F")
}
func (SBC_9F) Code() uint8 {
	return 0x9F
}
func (SBC_9F) String() string {
	return "SBC A,A"
}

// XOR A,D    code=0xaa
type XOR_AA struct{}

func (XOR_AA) Exec(cpu *CPU) {
	res := cpu.A ^ cpu.D

	var flags Flags
	if res == 0 {
		flags |= FLAGZ
	}
	cpu.F = FlagRegister(flags)
	cpu.A = res
	cpu.Cycles += 4
}
func (XOR_AA) Code() uint8 {
	return 0xAA
}
func (XOR_AA) String() string {
	return "XOR A,D"
}

// LD E,n8    code=0x1e
type LD_1E struct{}

func (LD_1E) Exec(cpu *CPU) {

	data := cpu.readU8(cpu.PC)
	cpu.IncProgramCounter()

	cpu.E = data

	cpu.Cycles += 8

}
func (LD_1E) Code() uint8 {
	return 0x1E
}
func (LD_1E) String() string {
	return "LD E,n8"
}

// JR C,e8    code=0x38
type JR_38 struct{}

func (JR_38) Exec(cpu *CPU) {
	e := cpu.readI8(cpu.PC)
	cpu.IncProgramCounter()
	if cpu.F.HasCarry() {
		cpu.PC, cpu.F = add(cpu.PC, e)
		cpu.Cycles += 12
	} else {
		cpu.Cycles += 8
	}
}
func (JR_38) Code() uint8 {
	return 0x38
}
func (JR_38) String() string {
	return "JR C,e8"
}

// LD H,B    code=0x60
type LD_60 struct{}

func (LD_60) Exec(cpu *CPU) {

	data := cpu.B

	cpu.H = data

	cpu.Cycles += 4

}
func (LD_60) Code() uint8 {
	return 0x60
}
func (LD_60) String() string {
	return "LD H,B"
}

// ADD A,L    code=0x85
type ADD_85 struct{}

func (ADD_85) Exec(cpu *CPU) {
	lhs := cpu.A
	rhs := cpu.L

	res, flags := add(lhs, rhs)
	cpu.A = res
	cpu.F = flags
	cpu.Cycles += 4
}
func (ADD_85) Code() uint8 {
	return 0x85
}
func (ADD_85) String() string {
	return "ADD A,L"
}

// PUSH BC    code=0xc5
type PUSH_C5 struct{}

func (PUSH_C5) Exec(cpu *CPU) {
	cpu.PushStack(cpu.BC())
	cpu.Cycles += 16
}
func (PUSH_C5) Code() uint8 {
	return 0xC5
}
func (PUSH_C5) String() string {
	return "PUSH BC"
}

// PUSH DE    code=0xd5
type PUSH_D5 struct{}

func (PUSH_D5) Exec(cpu *CPU) {
	cpu.PushStack(cpu.DE())
	cpu.Cycles += 16
}
func (PUSH_D5) Code() uint8 {
	return 0xD5
}
func (PUSH_D5) String() string {
	return "PUSH DE"
}

// RST $38    code=0xff
type RST_FF struct{}

func (RST_FF) Exec(cpu *CPU) {
	n := uint8(0x38)
	cpu.PushStack(cpu.PC)
	cpu.PC = concatU16(0x00, n)
}
func (RST_FF) Code() uint8 {
	return 0xFF
}
func (RST_FF) String() string {
	return "RST $38"
}

// LD B,n8    code=0x06
type LD_06 struct{}

func (LD_06) Exec(cpu *CPU) {

	data := cpu.readU8(cpu.PC)
	cpu.IncProgramCounter()

	cpu.B = data

	cpu.Cycles += 8

}
func (LD_06) Code() uint8 {
	return 0x6
}
func (LD_06) String() string {
	return "LD B,n8"
}

// DEC E    code=0x1d
type DEC_1D struct{}

func (DEC_1D) Exec(cpu *CPU) {
	res, flags := sub(cpu.E, 0x01)
	cpu.F = flags
	cpu.E = res
	cpu.Cycles += 4
}
func (DEC_1D) Code() uint8 {
	return 0x1D
}
func (DEC_1D) String() string {
	return "DEC E"
}

// DEC SP    code=0x3b
type DEC_3B struct{}

func (DEC_3B) Exec(cpu *CPU) {
	res, flags := sub(cpu.SP, 0x01)
	cpu.F = flags
	cpu.SP = res
	cpu.Cycles += 8
}
func (DEC_3B) Code() uint8 {
	return 0x3B
}
func (DEC_3B) String() string {
	return "DEC SP"
}

// LD L,L    code=0x6d
type LD_6D struct{}

func (LD_6D) Exec(cpu *CPU) {

	data := cpu.L

	cpu.L = data

	cpu.Cycles += 4

}
func (LD_6D) Code() uint8 {
	return 0x6D
}
func (LD_6D) String() string {
	return "LD L,L"
}

// SBC A,H    code=0x9c
type SBC_9C struct{}

func (SBC_9C) Exec(cpu *CPU) {
	panic("TODO SBC_9C")
}
func (SBC_9C) Code() uint8 {
	return 0x9C
}
func (SBC_9C) String() string {
	return "SBC A,H"
}

// JP NC,a16    code=0xd2
type JP_D2 struct{}

func (JP_D2) Exec(cpu *CPU) {
	panic("TODO JP_D2")
}
func (JP_D2) Code() uint8 {
	return 0xD2
}
func (JP_D2) String() string {
	return "JP NC,a16"
}

// CCF     code=0x3f
type CCF_3F struct{}

func (CCF_3F) Exec(cpu *CPU) {
	panic("TODO CCF_3F")
}
func (CCF_3F) Code() uint8 {
	return 0x3F
}
func (CCF_3F) String() string {
	return "CCF"
}

// INC HL    code=0x23
type INC_23 struct{}

func (INC_23) Exec(cpu *CPU) {
	res, flags := add(cpu.HL(), 0x01)
	cpu.F = flags
	cpu.H, cpu.L = split(res)
	cpu.Cycles += 8
}
func (INC_23) Code() uint8 {
	return 0x23
}
func (INC_23) String() string {
	return "INC HL"
}

// LD SP,n16    code=0x31
type LD_31 struct{}

func (LD_31) Exec(cpu *CPU) {

	data := cpu.readU16(cpu.PC)
	cpu.IncProgramCounter()
	cpu.IncProgramCounter()

	cpu.SP = data

	cpu.Cycles += 12

}
func (LD_31) Code() uint8 {
	return 0x31
}
func (LD_31) String() string {
	return "LD SP,n16"
}

// ADD HL,SP    code=0x39
type ADD_39 struct{}

func (ADD_39) Exec(cpu *CPU) {
	lhs := cpu.HL()
	rhs := cpu.SP

	res, flags := add(lhs, rhs)
	cpu.H, cpu.L = split(res)
	cpu.F = flags
	cpu.Cycles += 8
}
func (ADD_39) Code() uint8 {
	return 0x39
}
func (ADD_39) String() string {
	return "ADD HL,SP"
}

// LD C,D    code=0x4a
type LD_4A struct{}

func (LD_4A) Exec(cpu *CPU) {

	data := cpu.D

	cpu.C = data

	cpu.Cycles += 4

}
func (LD_4A) Code() uint8 {
	return 0x4A
}
func (LD_4A) String() string {
	return "LD C,D"
}

// LD D,L    code=0x55
type LD_55 struct{}

func (LD_55) Exec(cpu *CPU) {

	data := cpu.L

	cpu.D = data

	cpu.Cycles += 4

}
func (LD_55) Code() uint8 {
	return 0x55
}
func (LD_55) String() string {
	return "LD D,L"
}

// LD E,E    code=0x5b
type LD_5B struct{}

func (LD_5B) Exec(cpu *CPU) {

	data := cpu.E

	cpu.E = data

	cpu.Cycles += 4

}
func (LD_5B) Code() uint8 {
	return 0x5B
}
func (LD_5B) String() string {
	return "LD E,E"
}

// SUB A,C    code=0x91
type SUB_91 struct{}

func (SUB_91) Exec(cpu *CPU) {
	res, flags := sub(cpu.A, cpu.C)
	cpu.A = res

	cpu.F = flags
	cpu.Cycles += 4
}
func (SUB_91) Code() uint8 {
	return 0x91
}
func (SUB_91) String() string {
	return "SUB A,C"
}

// SUB A,H    code=0x94
type SUB_94 struct{}

func (SUB_94) Exec(cpu *CPU) {
	res, flags := sub(cpu.A, cpu.H)
	cpu.A = res

	cpu.F = flags
	cpu.Cycles += 4
}
func (SUB_94) Code() uint8 {
	return 0x94
}
func (SUB_94) String() string {
	return "SUB A,H"
}

// DEC A    code=0x3d
type DEC_3D struct{}

func (DEC_3D) Exec(cpu *CPU) {
	res, flags := sub(cpu.A, 0x01)
	cpu.F = flags
	cpu.A = res
	cpu.Cycles += 4
}
func (DEC_3D) Code() uint8 {
	return 0x3D
}
func (DEC_3D) String() string {
	return "DEC A"
}

// LD D,A    code=0x57
type LD_57 struct{}

func (LD_57) Exec(cpu *CPU) {

	data := cpu.A

	cpu.D = data

	cpu.Cycles += 4

}
func (LD_57) Code() uint8 {
	return 0x57
}
func (LD_57) String() string {
	return "LD D,A"
}

// LD L,D    code=0x6a
type LD_6A struct{}

func (LD_6A) Exec(cpu *CPU) {

	data := cpu.D

	cpu.L = data

	cpu.Cycles += 4

}
func (LD_6A) Code() uint8 {
	return 0x6A
}
func (LD_6A) String() string {
	return "LD L,D"
}

// AND A,H    code=0xa4
type AND_A4 struct{}

func (AND_A4) Exec(cpu *CPU) {
	panic("TODO AND_A4")
}
func (AND_A4) Code() uint8 {
	return 0xA4
}
func (AND_A4) String() string {
	return "AND A,H"
}

// CP A,A    code=0xbf
type CP_BF struct{}

func (CP_BF) Exec(cpu *CPU) {
	_, flags := sub(cpu.A, cpu.A)

	cpu.F = flags
	cpu.Cycles += 4
}
func (CP_BF) Code() uint8 {
	return 0xBF
}
func (CP_BF) String() string {
	return "CP A,A"
}

// CALL a16    code=0xcd
type CALL_CD struct{}

func (CALL_CD) Exec(cpu *CPU) {
	lsb := cpu.readU8(cpu.PC)
	cpu.IncProgramCounter()
	msb := cpu.readU8(cpu.PC)
	cpu.IncProgramCounter()
	nn := concatU16(msb, lsb)
	if true {
		cpu.PushStack(cpu.PC)
		cpu.PC = nn
		cpu.Cycles += 24
	} else {
		cpu.Cycles += 12
	}
}
func (CALL_CD) Code() uint8 {
	return 0xCD
}
func (CALL_CD) String() string {
	return "CALL a16"
}

// DI     code=0xf3
type DI_F3 struct{}

func (DI_F3) Exec(cpu *CPU) {
	panic("TODO DI_F3")
}
func (DI_F3) Code() uint8 {
	return 0xF3
}
func (DI_F3) String() string {
	return "DI"
}

// LD B,A    code=0x47
type LD_47 struct{}

func (LD_47) Exec(cpu *CPU) {

	data := cpu.A

	cpu.B = data

	cpu.Cycles += 4

}
func (LD_47) Code() uint8 {
	return 0x47
}
func (LD_47) String() string {
	return "LD B,A"
}

// SBC A,E    code=0x9b
type SBC_9B struct{}

func (SBC_9B) Exec(cpu *CPU) {
	panic("TODO SBC_9B")
}
func (SBC_9B) Code() uint8 {
	return 0x9B
}
func (SBC_9B) String() string {
	return "SBC A,E"
}

// RST $10    code=0xd7
type RST_D7 struct{}

func (RST_D7) Exec(cpu *CPU) {
	n := uint8(0x10)
	cpu.PushStack(cpu.PC)
	cpu.PC = concatU16(0x00, n)
}
func (RST_D7) Code() uint8 {
	return 0xD7
}
func (RST_D7) String() string {
	return "RST $10"
}

// OR A,n8    code=0xf6
type OR_F6 struct{}

func (OR_F6) Exec(cpu *CPU) {
	res := cpu.A | cpu.readU8(cpu.PC)
	var flags Flags
	if res == 0 {
		flags |= FLAGZ
	}
	cpu.F = FlagRegister(flags)
	cpu.A = res
	cpu.Cycles += 8
}
func (OR_F6) Code() uint8 {
	return 0xF6
}
func (OR_F6) String() string {
	return "OR A,n8"
}

// LD A,(a16)    code=0xfa
type LD_FA struct{}

func (LD_FA) Exec(cpu *CPU) {

	data := cpu.loadU8(cpu.readU16(cpu.PC))
	cpu.IncProgramCounter()
	cpu.IncProgramCounter()

	cpu.A = data

	cpu.Cycles += 16

}
func (LD_FA) Code() uint8 {
	return 0xFA
}
func (LD_FA) String() string {
	return "LD A,(a16)"
}

// LD A,C    code=0x79
type LD_79 struct{}

func (LD_79) Exec(cpu *CPU) {

	data := cpu.C

	cpu.A = data

	cpu.Cycles += 4

}
func (LD_79) Code() uint8 {
	return 0x79
}
func (LD_79) String() string {
	return "LD A,C"
}

// SBC A,B    code=0x98
type SBC_98 struct{}

func (SBC_98) Exec(cpu *CPU) {
	panic("TODO SBC_98")
}
func (SBC_98) Code() uint8 {
	return 0x98
}
func (SBC_98) String() string {
	return "SBC A,B"
}

// LD A,E    code=0x7b
type LD_7B struct{}

func (LD_7B) Exec(cpu *CPU) {

	data := cpu.E

	cpu.A = data

	cpu.Cycles += 4

}
func (LD_7B) Code() uint8 {
	return 0x7B
}
func (LD_7B) String() string {
	return "LD A,E"
}

// ADC A,C    code=0x89
type ADC_89 struct{}

func (ADC_89) Exec(cpu *CPU) {
	panic("TODO ADC_89")
}
func (ADC_89) Code() uint8 {
	return 0x89
}
func (ADC_89) String() string {
	return "ADC A,C"
}

// LD HL,SP+,e8    code=0xf8
type LD_F8 struct{}

func (LD_F8) Exec(cpu *CPU) {

	e := cpu.readI8(cpu.PC)
	cpu.IncProgramCounter()
	res, flags := add(cpu.SP, e)
	cpu.H, cpu.L = split(res)
	cpu.F = flags
	cpu.Cycles += 12

}
func (LD_F8) Code() uint8 {
	return 0xF8
}
func (LD_F8) String() string {
	return "LD HL,SP+,e8"
}

// INC D    code=0x14
type INC_14 struct{}

func (INC_14) Exec(cpu *CPU) {
	res, flags := add(cpu.D, 0x01)
	cpu.F = flags
	cpu.D = res
	cpu.Cycles += 4
}
func (INC_14) Code() uint8 {
	return 0x14
}
func (INC_14) String() string {
	return "INC D"
}

// LD B,H    code=0x44
type LD_44 struct{}

func (LD_44) Exec(cpu *CPU) {

	data := cpu.H

	cpu.B = data

	cpu.Cycles += 4

}
func (LD_44) Code() uint8 {
	return 0x44
}
func (LD_44) String() string {
	return "LD B,H"
}

// LD A,(HL)    code=0x7e
type LD_7E struct{}

func (LD_7E) Exec(cpu *CPU) {

	data := cpu.loadU8(cpu.HL())

	cpu.A = data

	cpu.Cycles += 8

}
func (LD_7E) Code() uint8 {
	return 0x7E
}
func (LD_7E) String() string {
	return "LD A,(HL)"
}

// PREFIX     code=0xcb
type PREFIX_CB struct{}

func (PREFIX_CB) Exec(cpu *CPU) {
	cpu.prefix = true
	cpu.Cycles += 4
}
func (PREFIX_CB) Code() uint8 {
	return 0xCB
}
func (PREFIX_CB) String() string {
	return "PREFIX"
}

// NOP     code=0x00
type NOP_00 struct{}

func (NOP_00) Exec(cpu *CPU) {
	cpu.Cycles += 4
}
func (NOP_00) Code() uint8 {
	return 0x0
}
func (NOP_00) String() string {
	return "NOP"
}

// LD D,H    code=0x54
type LD_54 struct{}

func (LD_54) Exec(cpu *CPU) {

	data := cpu.H

	cpu.D = data

	cpu.Cycles += 4

}
func (LD_54) Code() uint8 {
	return 0x54
}
func (LD_54) String() string {
	return "LD D,H"
}

// LD E,C    code=0x59
type LD_59 struct{}

func (LD_59) Exec(cpu *CPU) {

	data := cpu.C

	cpu.E = data

	cpu.Cycles += 4

}
func (LD_59) Code() uint8 {
	return 0x59
}
func (LD_59) String() string {
	return "LD E,C"
}

// JP NZ,a16    code=0xc2
type JP_C2 struct{}

func (JP_C2) Exec(cpu *CPU) {
	panic("TODO JP_C2")
}
func (JP_C2) Code() uint8 {
	return 0xC2
}
func (JP_C2) String() string {
	return "JP NZ,a16"
}

// LD (HL),n8    code=0x36
type LD_36 struct{}

func (LD_36) Exec(cpu *CPU) {

	data := cpu.readU8(cpu.PC)
	cpu.IncProgramCounter()

	cpu.WriteMemory(cpu.HL(), data)

	cpu.Cycles += 12

}
func (LD_36) Code() uint8 {
	return 0x36
}
func (LD_36) String() string {
	return "LD (HL),n8"
}

// LD (HL),L    code=0x75
type LD_75 struct{}

func (LD_75) Exec(cpu *CPU) {

	data := cpu.L

	cpu.WriteMemory(cpu.HL(), data)

	cpu.Cycles += 8

}
func (LD_75) Code() uint8 {
	return 0x75
}
func (LD_75) String() string {
	return "LD (HL),L"
}

// XOR A,H    code=0xac
type XOR_AC struct{}

func (XOR_AC) Exec(cpu *CPU) {
	res := cpu.A ^ cpu.H

	var flags Flags
	if res == 0 {
		flags |= FLAGZ
	}
	cpu.F = FlagRegister(flags)
	cpu.A = res
	cpu.Cycles += 4
}
func (XOR_AC) Code() uint8 {
	return 0xAC
}
func (XOR_AC) String() string {
	return "XOR A,H"
}

// SUB A,n8    code=0xd6
type SUB_D6 struct{}

func (SUB_D6) Exec(cpu *CPU) {
	res, flags := sub(cpu.A, cpu.readU8(cpu.PC))
	cpu.A = res
	cpu.IncProgramCounter()
	cpu.F = flags
	cpu.Cycles += 8
}
func (SUB_D6) Code() uint8 {
	return 0xD6
}
func (SUB_D6) String() string {
	return "SUB A,n8"
}

// INC L    code=0x2c
type INC_2C struct{}

func (INC_2C) Exec(cpu *CPU) {
	res, flags := add(cpu.L, 0x01)
	cpu.F = flags
	cpu.L = res
	cpu.Cycles += 4
}
func (INC_2C) Code() uint8 {
	return 0x2C
}
func (INC_2C) String() string {
	return "INC L"
}

// XOR A,L    code=0xad
type XOR_AD struct{}

func (XOR_AD) Exec(cpu *CPU) {
	res := cpu.A ^ cpu.L

	var flags Flags
	if res == 0 {
		flags |= FLAGZ
	}
	cpu.F = FlagRegister(flags)
	cpu.A = res
	cpu.Cycles += 4
}
func (XOR_AD) Code() uint8 {
	return 0xAD
}
func (XOR_AD) String() string {
	return "XOR A,L"
}

// JP HL    code=0xe9
type JP_E9 struct{}

func (JP_E9) Exec(cpu *CPU) {
	panic("TODO JP_E9")
}
func (JP_E9) Code() uint8 {
	return 0xE9
}
func (JP_E9) String() string {
	return "JP HL"
}

// JR Z,e8    code=0x28
type JR_28 struct{}

func (JR_28) Exec(cpu *CPU) {
	e := cpu.readI8(cpu.PC)
	cpu.IncProgramCounter()
	if cpu.F.HasZero() {
		cpu.PC, cpu.F = add(cpu.PC, e)
		cpu.Cycles += 12
	} else {
		cpu.Cycles += 8
	}
}
func (JR_28) Code() uint8 {
	return 0x28
}
func (JR_28) String() string {
	return "JR Z,e8"
}

// LD E,B    code=0x58
type LD_58 struct{}

func (LD_58) Exec(cpu *CPU) {

	data := cpu.B

	cpu.E = data

	cpu.Cycles += 4

}
func (LD_58) Code() uint8 {
	return 0x58
}
func (LD_58) String() string {
	return "LD E,B"
}

// ADC A,L    code=0x8d
type ADC_8D struct{}

func (ADC_8D) Exec(cpu *CPU) {
	panic("TODO ADC_8D")
}
func (ADC_8D) Code() uint8 {
	return 0x8D
}
func (ADC_8D) String() string {
	return "ADC A,L"
}

// ILLEGAL_D3     code=0xd3
type ILLEGAL_D3_D3 struct{}

func (ILLEGAL_D3_D3) Exec(cpu *CPU) {
	panic("TODO ILLEGAL_D3_D3")
}
func (ILLEGAL_D3_D3) Code() uint8 {
	return 0xD3
}
func (ILLEGAL_D3_D3) String() string {
	return "ILLEGAL_D3"
}

// LD E,L    code=0x5d
type LD_5D struct{}

func (LD_5D) Exec(cpu *CPU) {

	data := cpu.L

	cpu.E = data

	cpu.Cycles += 4

}
func (LD_5D) Code() uint8 {
	return 0x5D
}
func (LD_5D) String() string {
	return "LD E,L"
}

// LD H,A    code=0x67
type LD_67 struct{}

func (LD_67) Exec(cpu *CPU) {

	data := cpu.A

	cpu.H = data

	cpu.Cycles += 4

}
func (LD_67) Code() uint8 {
	return 0x67
}
func (LD_67) String() string {
	return "LD H,A"
}

// ADD A,B    code=0x80
type ADD_80 struct{}

func (ADD_80) Exec(cpu *CPU) {
	lhs := cpu.A
	rhs := cpu.B

	res, flags := add(lhs, rhs)
	cpu.A = res
	cpu.F = flags
	cpu.Cycles += 4
}
func (ADD_80) Code() uint8 {
	return 0x80
}
func (ADD_80) String() string {
	return "ADD A,B"
}

// LDH (C),A    code=0xe2
type LDH_E2 struct{}

func (LDH_E2) Exec(cpu *CPU) {
	pc0 := cpu.PC
	value := cpu.A
	cpu.WriteMemory(concatU16(0xFF, cpu.C), value)

	cpu.PC = pc0 + 1
	cpu.Cycles += 8
}
func (LDH_E2) Code() uint8 {
	return 0xE2
}
func (LDH_E2) String() string {
	return "LDH (C),A"
}

// RST $28    code=0xef
type RST_EF struct{}

func (RST_EF) Exec(cpu *CPU) {
	n := uint8(0x28)
	cpu.PushStack(cpu.PC)
	cpu.PC = concatU16(0x00, n)
}
func (RST_EF) Code() uint8 {
	return 0xEF
}
func (RST_EF) String() string {
	return "RST $28"
}

// LD C,A    code=0x4f
type LD_4F struct{}

func (LD_4F) Exec(cpu *CPU) {

	data := cpu.A

	cpu.C = data

	cpu.Cycles += 4

}
func (LD_4F) Code() uint8 {
	return 0x4F
}
func (LD_4F) String() string {
	return "LD C,A"
}

// POP AF    code=0xf1
type POP_F1 struct{}

func (POP_F1) Exec(cpu *CPU) {
	value := cpu.PopStack()
	msb, lsb := split(value)
	cpu.A, cpu.F = msb, FlagRegister(lsb)
	cpu.Cycles += 12
}
func (POP_F1) Code() uint8 {
	return 0xF1
}
func (POP_F1) String() string {
	return "POP AF"
}

// DEC D    code=0x15
type DEC_15 struct{}

func (DEC_15) Exec(cpu *CPU) {
	res, flags := sub(cpu.D, 0x01)
	cpu.F = flags
	cpu.D = res
	cpu.Cycles += 4
}
func (DEC_15) Code() uint8 {
	return 0x15
}
func (DEC_15) String() string {
	return "DEC D"
}

// LD H,L    code=0x65
type LD_65 struct{}

func (LD_65) Exec(cpu *CPU) {

	data := cpu.L

	cpu.H = data

	cpu.Cycles += 4

}
func (LD_65) Code() uint8 {
	return 0x65
}
func (LD_65) String() string {
	return "LD H,L"
}

// LD L,B    code=0x68
type LD_68 struct{}

func (LD_68) Exec(cpu *CPU) {

	data := cpu.B

	cpu.L = data

	cpu.Cycles += 4

}
func (LD_68) Code() uint8 {
	return 0x68
}
func (LD_68) String() string {
	return "LD L,B"
}

// LD E,H    code=0x5c
type LD_5C struct{}

func (LD_5C) Exec(cpu *CPU) {

	data := cpu.H

	cpu.E = data

	cpu.Cycles += 4

}
func (LD_5C) Code() uint8 {
	return 0x5C
}
func (LD_5C) String() string {
	return "LD E,H"
}

// ADD A,H    code=0x84
type ADD_84 struct{}

func (ADD_84) Exec(cpu *CPU) {
	lhs := cpu.A
	rhs := cpu.H

	res, flags := add(lhs, rhs)
	cpu.A = res
	cpu.F = flags
	cpu.Cycles += 4
}
func (ADD_84) Code() uint8 {
	return 0x84
}
func (ADD_84) String() string {
	return "ADD A,H"
}

// SBC A,C    code=0x99
type SBC_99 struct{}

func (SBC_99) Exec(cpu *CPU) {
	panic("TODO SBC_99")
}
func (SBC_99) Code() uint8 {
	return 0x99
}
func (SBC_99) String() string {
	return "SBC A,C"
}

// OR A,A    code=0xb7
type OR_B7 struct{}

func (OR_B7) Exec(cpu *CPU) {
	res := cpu.A | cpu.A
	var flags Flags
	if res == 0 {
		flags |= FLAGZ
	}
	cpu.F = FlagRegister(flags)
	cpu.A = res
	cpu.Cycles += 4
}
func (OR_B7) Code() uint8 {
	return 0xB7
}
func (OR_B7) String() string {
	return "OR A,A"
}

// RET     code=0xc9
type RET_C9 struct{}

func (RET_C9) Exec(cpu *CPU) {
	if true {
		cpu.PC = cpu.PopStack()
		cpu.Cycles += 16
	} else {
		cpu.Cycles += 16
	}
}
func (RET_C9) Code() uint8 {
	return 0xC9
}
func (RET_C9) String() string {
	return "RET"
}

// PUSH HL    code=0xe5
type PUSH_E5 struct{}

func (PUSH_E5) Exec(cpu *CPU) {
	cpu.PushStack(cpu.HL())
	cpu.Cycles += 16
}
func (PUSH_E5) Code() uint8 {
	return 0xE5
}
func (PUSH_E5) String() string {
	return "PUSH HL"
}

// LD SP,HL    code=0xf9
type LD_F9 struct{}

func (LD_F9) Exec(cpu *CPU) {

	data := cpu.HL()

	cpu.SP = data

	cpu.Cycles += 8

}
func (LD_F9) Code() uint8 {
	return 0xF9
}
func (LD_F9) String() string {
	return "LD SP,HL"
}

// INC E    code=0x1c
type INC_1C struct{}

func (INC_1C) Exec(cpu *CPU) {
	res, flags := add(cpu.E, 0x01)
	cpu.F = flags
	cpu.E = res
	cpu.Cycles += 4
}
func (INC_1C) Code() uint8 {
	return 0x1C
}
func (INC_1C) String() string {
	return "INC E"
}

// LD (HL+),A    code=0x22
type LD_22 struct{}

func (LD_22) Exec(cpu *CPU) {

	data := cpu.A

	cpu.WriteMemory(cpu.HL(), data)

	incr, flags := add(cpu.HL(), 0x01)
	cpu.H, cpu.L = split(incr)
	cpu.F = flags

	cpu.Cycles += 8

}
func (LD_22) Code() uint8 {
	return 0x22
}
func (LD_22) String() string {
	return "LD (HL+),A"
}

// OR A,(HL)    code=0xb6
type OR_B6 struct{}

func (OR_B6) Exec(cpu *CPU) {
	res := cpu.A | cpu.loadU8(cpu.HL())
	var flags Flags
	if res == 0 {
		flags |= FLAGZ
	}
	cpu.F = FlagRegister(flags)
	cpu.A = res
	cpu.Cycles += 8
}
func (OR_B6) Code() uint8 {
	return 0xB6
}
func (OR_B6) String() string {
	return "OR A,(HL)"
}

// ILLEGAL_EC     code=0xec
type ILLEGAL_EC_EC struct{}

func (ILLEGAL_EC_EC) Exec(cpu *CPU) {
	panic("TODO ILLEGAL_EC_EC")
}
func (ILLEGAL_EC_EC) Code() uint8 {
	return 0xEC
}
func (ILLEGAL_EC_EC) String() string {
	return "ILLEGAL_EC"
}

// PUSH AF    code=0xf5
type PUSH_F5 struct{}

func (PUSH_F5) Exec(cpu *CPU) {
	cpu.PushStack(cpu.AF())
	cpu.Cycles += 16
}
func (PUSH_F5) Code() uint8 {
	return 0xF5
}
func (PUSH_F5) String() string {
	return "PUSH AF"
}

// ADD A,E    code=0x83
type ADD_83 struct{}

func (ADD_83) Exec(cpu *CPU) {
	lhs := cpu.A
	rhs := cpu.E

	res, flags := add(lhs, rhs)
	cpu.A = res
	cpu.F = flags
	cpu.Cycles += 4
}
func (ADD_83) Code() uint8 {
	return 0x83
}
func (ADD_83) String() string {
	return "ADD A,E"
}

// POP HL    code=0xe1
type POP_E1 struct{}

func (POP_E1) Exec(cpu *CPU) {
	value := cpu.PopStack()
	cpu.H, cpu.L = split(value)
	cpu.Cycles += 12
}
func (POP_E1) Code() uint8 {
	return 0xE1
}
func (POP_E1) String() string {
	return "POP HL"
}

// LD D,E    code=0x53
type LD_53 struct{}

func (LD_53) Exec(cpu *CPU) {

	data := cpu.E

	cpu.D = data

	cpu.Cycles += 4

}
func (LD_53) Code() uint8 {
	return 0x53
}
func (LD_53) String() string {
	return "LD D,E"
}

// SUB A,E    code=0x93
type SUB_93 struct{}

func (SUB_93) Exec(cpu *CPU) {
	res, flags := sub(cpu.A, cpu.E)
	cpu.A = res

	cpu.F = flags
	cpu.Cycles += 4
}
func (SUB_93) Code() uint8 {
	return 0x93
}
func (SUB_93) String() string {
	return "SUB A,E"
}

// OR A,D    code=0xb2
type OR_B2 struct{}

func (OR_B2) Exec(cpu *CPU) {
	res := cpu.A | cpu.D
	var flags Flags
	if res == 0 {
		flags |= FLAGZ
	}
	cpu.F = FlagRegister(flags)
	cpu.A = res
	cpu.Cycles += 4
}
func (OR_B2) Code() uint8 {
	return 0xB2
}
func (OR_B2) String() string {
	return "OR A,D"
}

// LD D,n8    code=0x16
type LD_16 struct{}

func (LD_16) Exec(cpu *CPU) {

	data := cpu.readU8(cpu.PC)
	cpu.IncProgramCounter()

	cpu.D = data

	cpu.Cycles += 8

}
func (LD_16) Code() uint8 {
	return 0x16
}
func (LD_16) String() string {
	return "LD D,n8"
}

// DEC DE    code=0x1b
type DEC_1B struct{}

func (DEC_1B) Exec(cpu *CPU) {
	res, flags := sub(cpu.DE(), 0x01)
	cpu.F = flags
	cpu.D, cpu.E = split(res)
	cpu.Cycles += 8
}
func (DEC_1B) Code() uint8 {
	return 0x1B
}
func (DEC_1B) String() string {
	return "DEC DE"
}

// ADD HL,HL    code=0x29
type ADD_29 struct{}

func (ADD_29) Exec(cpu *CPU) {
	lhs := cpu.HL()
	rhs := cpu.HL()

	res, flags := add(lhs, rhs)
	cpu.H, cpu.L = split(res)
	cpu.F = flags
	cpu.Cycles += 8
}
func (ADD_29) Code() uint8 {
	return 0x29
}
func (ADD_29) String() string {
	return "ADD HL,HL"
}

// LD A,(HL+)    code=0x2a
type LD_2A struct{}

func (LD_2A) Exec(cpu *CPU) {

	data := cpu.loadU8(cpu.HL())

	cpu.A = data

	cpu.Cycles += 8

}
func (LD_2A) Code() uint8 {
	return 0x2A
}
func (LD_2A) String() string {
	return "LD A,(HL+)"
}

// LD C,C    code=0x49
type LD_49 struct{}

func (LD_49) Exec(cpu *CPU) {

	data := cpu.C

	cpu.C = data

	cpu.Cycles += 4

}
func (LD_49) Code() uint8 {
	return 0x49
}
func (LD_49) String() string {
	return "LD C,C"
}

// LD H,C    code=0x61
type LD_61 struct{}

func (LD_61) Exec(cpu *CPU) {

	data := cpu.C

	cpu.H = data

	cpu.Cycles += 4

}
func (LD_61) Code() uint8 {
	return 0x61
}
func (LD_61) String() string {
	return "LD H,C"
}

// ADC A,E    code=0x8b
type ADC_8B struct{}

func (ADC_8B) Exec(cpu *CPU) {
	panic("TODO ADC_8B")
}
func (ADC_8B) Code() uint8 {
	return 0x8B
}
func (ADC_8B) String() string {
	return "ADC A,E"
}

// SBC A,(HL)    code=0x9e
type SBC_9E struct{}

func (SBC_9E) Exec(cpu *CPU) {
	panic("TODO SBC_9E")
}
func (SBC_9E) Code() uint8 {
	return 0x9E
}
func (SBC_9E) String() string {
	return "SBC A,(HL)"
}

// XOR A,A    code=0xaf
type XOR_AF struct{}

func (XOR_AF) Exec(cpu *CPU) {
	res := cpu.A ^ cpu.A

	var flags Flags
	if res == 0 {
		flags |= FLAGZ
	}
	cpu.F = FlagRegister(flags)
	cpu.A = res
	cpu.Cycles += 4
}
func (XOR_AF) Code() uint8 {
	return 0xAF
}
func (XOR_AF) String() string {
	return "XOR A,A"
}

// RST $00    code=0xc7
type RST_C7 struct{}

func (RST_C7) Exec(cpu *CPU) {
	n := uint8(0x0)
	cpu.PushStack(cpu.PC)
	cpu.PC = concatU16(0x00, n)
}
func (RST_C7) Code() uint8 {
	return 0xC7
}
func (RST_C7) String() string {
	return "RST $00"
}

// LD (DE),A    code=0x12
type LD_12 struct{}

func (LD_12) Exec(cpu *CPU) {

	data := cpu.A

	cpu.WriteMemory(cpu.DE(), data)

	cpu.Cycles += 8

}
func (LD_12) Code() uint8 {
	return 0x12
}
func (LD_12) String() string {
	return "LD (DE),A"
}

// DEC H    code=0x25
type DEC_25 struct{}

func (DEC_25) Exec(cpu *CPU) {
	res, flags := sub(cpu.H, 0x01)
	cpu.F = flags
	cpu.H = res
	cpu.Cycles += 4
}
func (DEC_25) Code() uint8 {
	return 0x25
}
func (DEC_25) String() string {
	return "DEC H"
}

// XOR A,B    code=0xa8
type XOR_A8 struct{}

func (XOR_A8) Exec(cpu *CPU) {
	res := cpu.A ^ cpu.B

	var flags Flags
	if res == 0 {
		flags |= FLAGZ
	}
	cpu.F = FlagRegister(flags)
	cpu.A = res
	cpu.Cycles += 4
}
func (XOR_A8) Code() uint8 {
	return 0xA8
}
func (XOR_A8) String() string {
	return "XOR A,B"
}

// ILLEGAL_DB     code=0xdb
type ILLEGAL_DB_DB struct{}

func (ILLEGAL_DB_DB) Exec(cpu *CPU) {
	panic("TODO ILLEGAL_DB_DB")
}
func (ILLEGAL_DB_DB) Code() uint8 {
	return 0xDB
}
func (ILLEGAL_DB_DB) String() string {
	return "ILLEGAL_DB"
}

// LD B,L    code=0x45
type LD_45 struct{}

func (LD_45) Exec(cpu *CPU) {

	data := cpu.L

	cpu.B = data

	cpu.Cycles += 4

}
func (LD_45) Code() uint8 {
	return 0x45
}
func (LD_45) String() string {
	return "LD B,L"
}

// LD (HL),C    code=0x71
type LD_71 struct{}

func (LD_71) Exec(cpu *CPU) {

	data := cpu.C

	cpu.WriteMemory(cpu.HL(), data)

	cpu.Cycles += 8

}
func (LD_71) Code() uint8 {
	return 0x71
}
func (LD_71) String() string {
	return "LD (HL),C"
}

// LD A,D    code=0x7a
type LD_7A struct{}

func (LD_7A) Exec(cpu *CPU) {

	data := cpu.D

	cpu.A = data

	cpu.Cycles += 4

}
func (LD_7A) Code() uint8 {
	return 0x7A
}
func (LD_7A) String() string {
	return "LD A,D"
}

// ADD A,n8    code=0xc6
type ADD_C6 struct{}

func (ADD_C6) Exec(cpu *CPU) {
	lhs := cpu.A
	rhs := cpu.readU8(cpu.PC)

	cpu.IncProgramCounter()
	res, flags := add(lhs, rhs)
	cpu.A = res
	cpu.F = flags
	cpu.Cycles += 8
}
func (ADD_C6) Code() uint8 {
	return 0xC6
}
func (ADD_C6) String() string {
	return "ADD A,n8"
}

// CP A,D    code=0xba
type CP_BA struct{}

func (CP_BA) Exec(cpu *CPU) {
	_, flags := sub(cpu.A, cpu.D)

	cpu.F = flags
	cpu.Cycles += 4
}
func (CP_BA) Code() uint8 {
	return 0xBA
}
func (CP_BA) String() string {
	return "CP A,D"
}

// JP Z,a16    code=0xca
type JP_CA struct{}

func (JP_CA) Exec(cpu *CPU) {
	panic("TODO JP_CA")
}
func (JP_CA) Code() uint8 {
	return 0xCA
}
func (JP_CA) String() string {
	return "JP Z,a16"
}

// ILLEGAL_E4     code=0xe4
type ILLEGAL_E4_E4 struct{}

func (ILLEGAL_E4_E4) Exec(cpu *CPU) {
	panic("TODO ILLEGAL_E4_E4")
}
func (ILLEGAL_E4_E4) Code() uint8 {
	return 0xE4
}
func (ILLEGAL_E4_E4) String() string {
	return "ILLEGAL_E4"
}

// LD A,n8    code=0x3e
type LD_3E struct{}

func (LD_3E) Exec(cpu *CPU) {

	data := cpu.readU8(cpu.PC)
	cpu.IncProgramCounter()

	cpu.A = data

	cpu.Cycles += 8

}
func (LD_3E) Code() uint8 {
	return 0x3E
}
func (LD_3E) String() string {
	return "LD A,n8"
}

// LD B,E    code=0x43
type LD_43 struct{}

func (LD_43) Exec(cpu *CPU) {

	data := cpu.E

	cpu.B = data

	cpu.Cycles += 4

}
func (LD_43) Code() uint8 {
	return 0x43
}
func (LD_43) String() string {
	return "LD B,E"
}

// LD E,A    code=0x5f
type LD_5F struct{}

func (LD_5F) Exec(cpu *CPU) {

	data := cpu.A

	cpu.E = data

	cpu.Cycles += 4

}
func (LD_5F) Code() uint8 {
	return 0x5F
}
func (LD_5F) String() string {
	return "LD E,A"
}

// LD L,C    code=0x69
type LD_69 struct{}

func (LD_69) Exec(cpu *CPU) {

	data := cpu.C

	cpu.L = data

	cpu.Cycles += 4

}
func (LD_69) Code() uint8 {
	return 0x69
}
func (LD_69) String() string {
	return "LD L,C"
}

// LD L,A    code=0x6f
type LD_6F struct{}

func (LD_6F) Exec(cpu *CPU) {

	data := cpu.A

	cpu.L = data

	cpu.Cycles += 4

}
func (LD_6F) Code() uint8 {
	return 0x6F
}
func (LD_6F) String() string {
	return "LD L,A"
}

// SUB A,A    code=0x97
type SUB_97 struct{}

func (SUB_97) Exec(cpu *CPU) {
	res, flags := sub(cpu.A, cpu.A)
	cpu.A = res

	cpu.F = flags
	cpu.Cycles += 4
}
func (SUB_97) Code() uint8 {
	return 0x97
}
func (SUB_97) String() string {
	return "SUB A,A"
}

// CP A,H    code=0xbc
type CP_BC struct{}

func (CP_BC) Exec(cpu *CPU) {
	_, flags := sub(cpu.A, cpu.H)

	cpu.F = flags
	cpu.Cycles += 4
}
func (CP_BC) Code() uint8 {
	return 0xBC
}
func (CP_BC) String() string {
	return "CP A,H"
}

// JP a16    code=0xc3
type JP_C3 struct{}

func (JP_C3) Exec(cpu *CPU) {
	panic("TODO JP_C3")
}
func (JP_C3) Code() uint8 {
	return 0xC3
}
func (JP_C3) String() string {
	return "JP a16"
}

// RRA     code=0x1f
type RRA_1F struct{}

func (RRA_1F) Exec(cpu *CPU) {
	cpu.A, cpu.F = rotate(cpu.A, 1, cpu.F, false)
	cpu.Cycles += 4
}
func (RRA_1F) Code() uint8 {
	return 0x1F
}
func (RRA_1F) String() string {
	return "RRA"
}

// LD C,E    code=0x4b
type LD_4B struct{}

func (LD_4B) Exec(cpu *CPU) {

	data := cpu.E

	cpu.C = data

	cpu.Cycles += 4

}
func (LD_4B) Code() uint8 {
	return 0x4B
}
func (LD_4B) String() string {
	return "LD C,E"
}

// DEC BC    code=0x0b
type DEC_0B struct{}

func (DEC_0B) Exec(cpu *CPU) {
	res, flags := sub(cpu.BC(), 0x01)
	cpu.F = flags
	cpu.B, cpu.C = split(res)
	cpu.Cycles += 8
}
func (DEC_0B) Code() uint8 {
	return 0xB
}
func (DEC_0B) String() string {
	return "DEC BC"
}

// STOP n8    code=0x10
type STOP_10 struct{}

func (STOP_10) Exec(cpu *CPU) {
	cpu.err = ErrNoMoreInstructions
	cpu.Cycles += 4
}
func (STOP_10) Code() uint8 {
	return 0x10
}
func (STOP_10) String() string {
	return "STOP"
}

// LD DE,n16    code=0x11
type LD_11 struct{}

func (LD_11) Exec(cpu *CPU) {

	data := cpu.readU16(cpu.PC)
	cpu.IncProgramCounter()
	cpu.IncProgramCounter()

	cpu.D, cpu.E = split(data)

	cpu.Cycles += 12

}
func (LD_11) Code() uint8 {
	return 0x11
}
func (LD_11) String() string {
	return "LD DE,n16"
}

// LD B,(HL)    code=0x46
type LD_46 struct{}

func (LD_46) Exec(cpu *CPU) {

	data := cpu.loadU8(cpu.HL())

	cpu.B = data

	cpu.Cycles += 8

}
func (LD_46) Code() uint8 {
	return 0x46
}
func (LD_46) String() string {
	return "LD B,(HL)"
}

// LD (a16),A    code=0xea
type LD_EA struct{}

func (LD_EA) Exec(cpu *CPU) {

	data := cpu.A

	cpu.WriteMemory(cpu.readU16(cpu.PC), data)
	cpu.IncProgramCounter()
	cpu.IncProgramCounter()

	cpu.Cycles += 16

}
func (LD_EA) Code() uint8 {
	return 0xEA
}
func (LD_EA) String() string {
	return "LD (a16),A"
}

// AND A,D    code=0xa2
type AND_A2 struct{}

func (AND_A2) Exec(cpu *CPU) {
	panic("TODO AND_A2")
}
func (AND_A2) Code() uint8 {
	return 0xA2
}
func (AND_A2) String() string {
	return "AND A,D"
}

// OR A,L    code=0xb5
type OR_B5 struct{}

func (OR_B5) Exec(cpu *CPU) {
	res := cpu.A | cpu.L
	var flags Flags
	if res == 0 {
		flags |= FLAGZ
	}
	cpu.F = FlagRegister(flags)
	cpu.A = res
	cpu.Cycles += 4
}
func (OR_B5) Code() uint8 {
	return 0xB5
}
func (OR_B5) String() string {
	return "OR A,L"
}

// JR e8    code=0x18
type JR_18 struct{}

func (JR_18) Exec(cpu *CPU) {
	e := cpu.readI8(cpu.PC)
	cpu.IncProgramCounter()
	if true {
		cpu.PC, cpu.F = add(cpu.PC, e)
		cpu.Cycles += 12
	} else {
		cpu.Cycles += 12
	}
}
func (JR_18) Code() uint8 {
	return 0x18
}
func (JR_18) String() string {
	return "JR e8"
}

// ADD HL,DE    code=0x19
type ADD_19 struct{}

func (ADD_19) Exec(cpu *CPU) {
	lhs := cpu.HL()
	rhs := cpu.DE()

	res, flags := add(lhs, rhs)
	cpu.H, cpu.L = split(res)
	cpu.F = flags
	cpu.Cycles += 8
}
func (ADD_19) Code() uint8 {
	return 0x19
}
func (ADD_19) String() string {
	return "ADD HL,DE"
}

// CPL     code=0x2f
type CPL_2F struct{}

func (CPL_2F) Exec(cpu *CPU) {
	panic("TODO CPL_2F")
}
func (CPL_2F) Code() uint8 {
	return 0x2F
}
func (CPL_2F) String() string {
	return "CPL"
}

// LD E,(HL)    code=0x5e
type LD_5E struct{}

func (LD_5E) Exec(cpu *CPU) {

	data := cpu.loadU8(cpu.HL())

	cpu.E = data

	cpu.Cycles += 8

}
func (LD_5E) Code() uint8 {
	return 0x5E
}
func (LD_5E) String() string {
	return "LD E,(HL)"
}

// CALL Z,a16    code=0xcc
type CALL_CC struct{}

func (CALL_CC) Exec(cpu *CPU) {
	lsb := cpu.readU8(cpu.PC)
	cpu.IncProgramCounter()
	msb := cpu.readU8(cpu.PC)
	cpu.IncProgramCounter()
	nn := concatU16(msb, lsb)
	if cpu.F.HasZero() {
		cpu.PushStack(cpu.PC)
		cpu.PC = nn
		cpu.Cycles += 24
	} else {
		cpu.Cycles += 12
	}
}
func (CALL_CC) Code() uint8 {
	return 0xCC
}
func (CALL_CC) String() string {
	return "CALL Z,a16"
}

// XOR A,n8    code=0xee
type XOR_EE struct{}

func (XOR_EE) Exec(cpu *CPU) {
	res := cpu.A ^ cpu.readU8(cpu.PC)
	cpu.IncProgramCounter()
	var flags Flags
	if res == 0 {
		flags |= FLAGZ
	}
	cpu.F = FlagRegister(flags)
	cpu.A = res
	cpu.Cycles += 8
}
func (XOR_EE) Code() uint8 {
	return 0xEE
}
func (XOR_EE) String() string {
	return "XOR A,n8"
}

// LD H,D    code=0x62
type LD_62 struct{}

func (LD_62) Exec(cpu *CPU) {

	data := cpu.D

	cpu.H = data

	cpu.Cycles += 4

}
func (LD_62) Code() uint8 {
	return 0x62
}
func (LD_62) String() string {
	return "LD H,D"
}

// AND A,E    code=0xa3
type AND_A3 struct{}

func (AND_A3) Exec(cpu *CPU) {
	panic("TODO AND_A3")
}
func (AND_A3) Code() uint8 {
	return 0xA3
}
func (AND_A3) String() string {
	return "AND A,E"
}

// CP A,L    code=0xbd
type CP_BD struct{}

func (CP_BD) Exec(cpu *CPU) {
	_, flags := sub(cpu.A, cpu.L)

	cpu.F = flags
	cpu.Cycles += 4
}
func (CP_BD) Code() uint8 {
	return 0xBD
}
func (CP_BD) String() string {
	return "CP A,L"
}

// DEC B    code=0x05
type DEC_05 struct{}

func (DEC_05) Exec(cpu *CPU) {
	res, flags := sub(cpu.B, 0x01)
	cpu.F = flags
	cpu.B = res
	cpu.Cycles += 4
}
func (DEC_05) Code() uint8 {
	return 0x5
}
func (DEC_05) String() string {
	return "DEC B"
}

// OR A,E    code=0xb3
type OR_B3 struct{}

func (OR_B3) Exec(cpu *CPU) {
	res := cpu.A | cpu.E
	var flags Flags
	if res == 0 {
		flags |= FLAGZ
	}
	cpu.F = FlagRegister(flags)
	cpu.A = res
	cpu.Cycles += 4
}
func (OR_B3) Code() uint8 {
	return 0xB3
}
func (OR_B3) String() string {
	return "OR A,E"
}

// DEC C    code=0x0d
type DEC_0D struct{}

func (DEC_0D) Exec(cpu *CPU) {
	res, flags := sub(cpu.C, 0x01)
	cpu.F = flags
	cpu.C = res
	cpu.Cycles += 4
}
func (DEC_0D) Code() uint8 {
	return 0xD
}
func (DEC_0D) String() string {
	return "DEC C"
}

// LD D,(HL)    code=0x56
type LD_56 struct{}

func (LD_56) Exec(cpu *CPU) {

	data := cpu.loadU8(cpu.HL())

	cpu.D = data

	cpu.Cycles += 8

}
func (LD_56) Code() uint8 {
	return 0x56
}
func (LD_56) String() string {
	return "LD D,(HL)"
}

// OR A,B    code=0xb0
type OR_B0 struct{}

func (OR_B0) Exec(cpu *CPU) {
	res := cpu.A | cpu.B
	var flags Flags
	if res == 0 {
		flags |= FLAGZ
	}
	cpu.F = FlagRegister(flags)
	cpu.A = res
	cpu.Cycles += 4
}
func (OR_B0) Code() uint8 {
	return 0xB0
}
func (OR_B0) String() string {
	return "OR A,B"
}

// ADC A,n8    code=0xce
type ADC_CE struct{}

func (ADC_CE) Exec(cpu *CPU) {
	panic("TODO ADC_CE")
}
func (ADC_CE) Code() uint8 {
	return 0xCE
}
func (ADC_CE) String() string {
	return "ADC A,n8"
}

// LD L,n8    code=0x2e
type LD_2E struct{}

func (LD_2E) Exec(cpu *CPU) {

	data := cpu.readU8(cpu.PC)
	cpu.IncProgramCounter()

	cpu.L = data

	cpu.Cycles += 8

}
func (LD_2E) Code() uint8 {
	return 0x2E
}
func (LD_2E) String() string {
	return "LD L,n8"
}

// LD L,(HL)    code=0x6e
type LD_6E struct{}

func (LD_6E) Exec(cpu *CPU) {

	data := cpu.loadU8(cpu.HL())

	cpu.L = data

	cpu.Cycles += 8

}
func (LD_6E) Code() uint8 {
	return 0x6E
}
func (LD_6E) String() string {
	return "LD L,(HL)"
}

// LD (HL),E    code=0x73
type LD_73 struct{}

func (LD_73) Exec(cpu *CPU) {

	data := cpu.E

	cpu.WriteMemory(cpu.HL(), data)

	cpu.Cycles += 8

}
func (LD_73) Code() uint8 {
	return 0x73
}
func (LD_73) String() string {
	return "LD (HL),E"
}

// LD A,B    code=0x78
type LD_78 struct{}

func (LD_78) Exec(cpu *CPU) {

	data := cpu.B

	cpu.A = data

	cpu.Cycles += 4

}
func (LD_78) Code() uint8 {
	return 0x78
}
func (LD_78) String() string {
	return "LD A,B"
}

// LD A,H    code=0x7c
type LD_7C struct{}

func (LD_7C) Exec(cpu *CPU) {

	data := cpu.H

	cpu.A = data

	cpu.Cycles += 4

}
func (LD_7C) Code() uint8 {
	return 0x7C
}
func (LD_7C) String() string {
	return "LD A,H"
}

// AND A,L    code=0xa5
type AND_A5 struct{}

func (AND_A5) Exec(cpu *CPU) {
	panic("TODO AND_A5")
}
func (AND_A5) Code() uint8 {
	return 0xA5
}
func (AND_A5) String() string {
	return "AND A,L"
}

// OR A,H    code=0xb4
type OR_B4 struct{}

func (OR_B4) Exec(cpu *CPU) {
	res := cpu.A | cpu.H
	var flags Flags
	if res == 0 {
		flags |= FLAGZ
	}
	cpu.F = FlagRegister(flags)
	cpu.A = res
	cpu.Cycles += 4
}
func (OR_B4) Code() uint8 {
	return 0xB4
}
func (OR_B4) String() string {
	return "OR A,H"
}

// RET NZ    code=0xc0
type RET_C0 struct{}

func (RET_C0) Exec(cpu *CPU) {
	if !cpu.F.HasZero() {
		cpu.PC = cpu.PopStack()
		cpu.Cycles += 20
	} else {
		cpu.Cycles += 8
	}
}
func (RET_C0) Code() uint8 {
	return 0xC0
}
func (RET_C0) String() string {
	return "RET NZ"
}

// LD H,(HL)    code=0x66
type LD_66 struct{}

func (LD_66) Exec(cpu *CPU) {

	data := cpu.loadU8(cpu.HL())

	cpu.H = data

	cpu.Cycles += 8

}
func (LD_66) Code() uint8 {
	return 0x66
}
func (LD_66) String() string {
	return "LD H,(HL)"
}

// ADC A,(HL)    code=0x8e
type ADC_8E struct{}

func (ADC_8E) Exec(cpu *CPU) {
	panic("TODO ADC_8E")
}
func (ADC_8E) Code() uint8 {
	return 0x8E
}
func (ADC_8E) String() string {
	return "ADC A,(HL)"
}

// RES 6,E    code=0xb3
type RES_B3 struct{}

func (RES_B3) Exec(cpu *CPU) {
	panic("TODO RES_B3")
}
func (RES_B3) Code() uint8 {
	return 0xB3
}
func (RES_B3) String() string {
	return "RES 6,E"
}

// SET 1,(HL)    code=0xce
type SET_CE struct{}

func (SET_CE) Exec(cpu *CPU) {
	panic("TODO SET_CE")
}
func (SET_CE) Code() uint8 {
	return 0xCE
}
func (SET_CE) String() string {
	return "SET 1,(HL)"
}

// SET 7,(HL)    code=0xfe
type SET_FE struct{}

func (SET_FE) Exec(cpu *CPU) {
	panic("TODO SET_FE")
}
func (SET_FE) Code() uint8 {
	return 0xFE
}
func (SET_FE) String() string {
	return "SET 7,(HL)"
}

// SWAP H    code=0x34
type SWAP_34 struct{}

func (SWAP_34) Exec(cpu *CPU) {
	panic("TODO SWAP_34")
}
func (SWAP_34) Code() uint8 {
	return 0x34
}
func (SWAP_34) String() string {
	return "SWAP H"
}

// RES 1,H    code=0x8c
type RES_8C struct{}

func (RES_8C) Exec(cpu *CPU) {
	panic("TODO RES_8C")
}
func (RES_8C) Code() uint8 {
	return 0x8C
}
func (RES_8C) String() string {
	return "RES 1,H"
}

// BIT 3,D    code=0x5a
type BIT_5A struct{}

func (BIT_5A) Exec(cpu *CPU) {
	value := cpu.D
	var flags Flags
	if (value & (1 << 3)) == 0 {
		flags |= FLAGZ
	}
	flags |= FLAGH
	cpu.F = FlagRegister(flags)

	cpu.Cycles += 8
}
func (BIT_5A) Code() uint8 {
	return 0x5A
}
func (BIT_5A) String() string {
	return "BIT 3,D"
}

// RES 3,L    code=0x9d
type RES_9D struct{}

func (RES_9D) Exec(cpu *CPU) {
	panic("TODO RES_9D")
}
func (RES_9D) Code() uint8 {
	return 0x9D
}
func (RES_9D) String() string {
	return "RES 3,L"
}

// SET 3,H    code=0xdc
type SET_DC struct{}

func (SET_DC) Exec(cpu *CPU) {
	panic("TODO SET_DC")
}
func (SET_DC) Code() uint8 {
	return 0xDC
}
func (SET_DC) String() string {
	return "SET 3,H"
}

// RR A    code=0x1f
type RR_1F struct{}

func (RR_1F) Exec(cpu *CPU) {
	res, flags := rotate(cpu.A, 1, cpu.F, false)
	cpu.A = res
	cpu.F = flags

	cpu.Cycles += 8
}
func (RR_1F) Code() uint8 {
	return 0x1F
}
func (RR_1F) String() string {
	return "RR A"
}

// BIT 2,(HL)    code=0x56
type BIT_56 struct{}

func (BIT_56) Exec(cpu *CPU) {
	value := cpu.loadU8(cpu.HL())
	var flags Flags
	if (value & (1 << 2)) == 0 {
		flags |= FLAGZ
	}
	flags |= FLAGH
	cpu.F = FlagRegister(flags)

	cpu.Cycles += 12
}
func (BIT_56) Code() uint8 {
	return 0x56
}
func (BIT_56) String() string {
	return "BIT 2,(HL)"
}

// RES 6,L    code=0xb5
type RES_B5 struct{}

func (RES_B5) Exec(cpu *CPU) {
	panic("TODO RES_B5")
}
func (RES_B5) Code() uint8 {
	return 0xB5
}
func (RES_B5) String() string {
	return "RES 6,L"
}

// SET 4,C    code=0xe1
type SET_E1 struct{}

func (SET_E1) Exec(cpu *CPU) {
	panic("TODO SET_E1")
}
func (SET_E1) Code() uint8 {
	return 0xE1
}
func (SET_E1) String() string {
	return "SET 4,C"
}

// SET 7,D    code=0xfa
type SET_FA struct{}

func (SET_FA) Exec(cpu *CPU) {
	panic("TODO SET_FA")
}
func (SET_FA) Code() uint8 {
	return 0xFA
}
func (SET_FA) String() string {
	return "SET 7,D"
}

// SRA (HL)    code=0x2e
type SRA_2E struct{}

func (SRA_2E) Exec(cpu *CPU) {
	panic("TODO SRA_2E")
}
func (SRA_2E) Code() uint8 {
	return 0x2E
}
func (SRA_2E) String() string {
	return "SRA (HL)"
}

// RES 0,C    code=0x81
type RES_81 struct{}

func (RES_81) Exec(cpu *CPU) {
	panic("TODO RES_81")
}
func (RES_81) Code() uint8 {
	return 0x81
}
func (RES_81) String() string {
	return "RES 0,C"
}

// RES 7,C    code=0xb9
type RES_B9 struct{}

func (RES_B9) Exec(cpu *CPU) {
	panic("TODO RES_B9")
}
func (RES_B9) Code() uint8 {
	return 0xB9
}
func (RES_B9) String() string {
	return "RES 7,C"
}

// SET 3,L    code=0xdd
type SET_DD struct{}

func (SET_DD) Exec(cpu *CPU) {
	panic("TODO SET_DD")
}
func (SET_DD) Code() uint8 {
	return 0xDD
}
func (SET_DD) String() string {
	return "SET 3,L"
}

// RL C    code=0x11
type RL_11 struct{}

func (RL_11) Exec(cpu *CPU) {
	res, flags := rotate(cpu.C, 0, cpu.F, false)
	cpu.C = res
	cpu.F = flags

	cpu.Cycles += 8
}
func (RL_11) Code() uint8 {
	return 0x11
}
func (RL_11) String() string {
	return "RL C"
}

// SRA E    code=0x2b
type SRA_2B struct{}

func (SRA_2B) Exec(cpu *CPU) {
	panic("TODO SRA_2B")
}
func (SRA_2B) Code() uint8 {
	return 0x2B
}
func (SRA_2B) String() string {
	return "SRA E"
}

// SRL A    code=0x3f
type SRL_3F struct{}

func (SRL_3F) Exec(cpu *CPU) {
	panic("TODO SRL_3F")
}
func (SRL_3F) Code() uint8 {
	return 0x3F
}
func (SRL_3F) String() string {
	return "SRL A"
}

// BIT 4,H    code=0x64
type BIT_64 struct{}

func (BIT_64) Exec(cpu *CPU) {
	value := cpu.H
	var flags Flags
	if (value & (1 << 4)) == 0 {
		flags |= FLAGZ
	}
	flags |= FLAGH
	cpu.F = FlagRegister(flags)

	cpu.Cycles += 8
}
func (BIT_64) Code() uint8 {
	return 0x64
}
func (BIT_64) String() string {
	return "BIT 4,H"
}

// RES 1,E    code=0x8b
type RES_8B struct{}

func (RES_8B) Exec(cpu *CPU) {
	panic("TODO RES_8B")
}
func (RES_8B) Code() uint8 {
	return 0x8B
}
func (RES_8B) String() string {
	return "RES 1,E"
}

// RES 4,A    code=0xa7
type RES_A7 struct{}

func (RES_A7) Exec(cpu *CPU) {
	panic("TODO RES_A7")
}
func (RES_A7) Code() uint8 {
	return 0xA7
}
func (RES_A7) String() string {
	return "RES 4,A"
}

// RLC (HL)    code=0x06
type RLC_06 struct{}

func (RLC_06) Exec(cpu *CPU) {
	res, flags := rotate(cpu.loadU8(cpu.HL()), 0, cpu.F, true)
	cpu.WriteMemory(cpu.HL(), res)
	cpu.F = flags

	cpu.Cycles += 16
}
func (RLC_06) Code() uint8 {
	return 0x6
}
func (RLC_06) String() string {
	return "RLC (HL)"
}

// BIT 0,C    code=0x41
type BIT_41 struct{}

func (BIT_41) Exec(cpu *CPU) {
	value := cpu.C
	var flags Flags
	if (value & (1 << 0)) == 0 {
		flags |= FLAGZ
	}
	flags |= FLAGH
	cpu.F = FlagRegister(flags)

	cpu.Cycles += 8
}
func (BIT_41) Code() uint8 {
	return 0x41
}
func (BIT_41) String() string {
	return "BIT 0,C"
}

// BIT 2,E    code=0x53
type BIT_53 struct{}

func (BIT_53) Exec(cpu *CPU) {
	value := cpu.E
	var flags Flags
	if (value & (1 << 2)) == 0 {
		flags |= FLAGZ
	}
	flags |= FLAGH
	cpu.F = FlagRegister(flags)

	cpu.Cycles += 8
}
func (BIT_53) Code() uint8 {
	return 0x53
}
func (BIT_53) String() string {
	return "BIT 2,E"
}

// BIT 4,A    code=0x67
type BIT_67 struct{}

func (BIT_67) Exec(cpu *CPU) {
	value := cpu.A
	var flags Flags
	if (value & (1 << 4)) == 0 {
		flags |= FLAGZ
	}
	flags |= FLAGH
	cpu.F = FlagRegister(flags)

	cpu.Cycles += 8
}
func (BIT_67) Code() uint8 {
	return 0x67
}
func (BIT_67) String() string {
	return "BIT 4,A"
}

// BIT 6,A    code=0x77
type BIT_77 struct{}

func (BIT_77) Exec(cpu *CPU) {
	value := cpu.A
	var flags Flags
	if (value & (1 << 6)) == 0 {
		flags |= FLAGZ
	}
	flags |= FLAGH
	cpu.F = FlagRegister(flags)

	cpu.Cycles += 8
}
func (BIT_77) Code() uint8 {
	return 0x77
}
func (BIT_77) String() string {
	return "BIT 6,A"
}

// RES 6,H    code=0xb4
type RES_B4 struct{}

func (RES_B4) Exec(cpu *CPU) {
	panic("TODO RES_B4")
}
func (RES_B4) Code() uint8 {
	return 0xB4
}
func (RES_B4) String() string {
	return "RES 6,H"
}

// SET 5,B    code=0xe8
type SET_E8 struct{}

func (SET_E8) Exec(cpu *CPU) {
	panic("TODO SET_E8")
}
func (SET_E8) Code() uint8 {
	return 0xE8
}
func (SET_E8) String() string {
	return "SET 5,B"
}

// SET 7,B    code=0xf8
type SET_F8 struct{}

func (SET_F8) Exec(cpu *CPU) {
	panic("TODO SET_F8")
}
func (SET_F8) Code() uint8 {
	return 0xF8
}
func (SET_F8) String() string {
	return "SET 7,B"
}

// RL D    code=0x12
type RL_12 struct{}

func (RL_12) Exec(cpu *CPU) {
	res, flags := rotate(cpu.D, 0, cpu.F, false)
	cpu.D = res
	cpu.F = flags

	cpu.Cycles += 8
}
func (RL_12) Code() uint8 {
	return 0x12
}
func (RL_12) String() string {
	return "RL D"
}

// BIT 2,B    code=0x50
type BIT_50 struct{}

func (BIT_50) Exec(cpu *CPU) {
	value := cpu.B
	var flags Flags
	if (value & (1 << 2)) == 0 {
		flags |= FLAGZ
	}
	flags |= FLAGH
	cpu.F = FlagRegister(flags)

	cpu.Cycles += 8
}
func (BIT_50) Code() uint8 {
	return 0x50
}
func (BIT_50) String() string {
	return "BIT 2,B"
}

// RES 1,B    code=0x88
type RES_88 struct{}

func (RES_88) Exec(cpu *CPU) {
	panic("TODO RES_88")
}
func (RES_88) Code() uint8 {
	return 0x88
}
func (RES_88) String() string {
	return "RES 1,B"
}

// RES 7,A    code=0xbf
type RES_BF struct{}

func (RES_BF) Exec(cpu *CPU) {
	panic("TODO RES_BF")
}
func (RES_BF) Code() uint8 {
	return 0xBF
}
func (RES_BF) String() string {
	return "RES 7,A"
}

// SET 1,C    code=0xc9
type SET_C9 struct{}

func (SET_C9) Exec(cpu *CPU) {
	panic("TODO SET_C9")
}
func (SET_C9) Code() uint8 {
	return 0xC9
}
func (SET_C9) String() string {
	return "SET 1,C"
}

// RES 2,H    code=0x94
type RES_94 struct{}

func (RES_94) Exec(cpu *CPU) {
	panic("TODO RES_94")
}
func (RES_94) Code() uint8 {
	return 0x94
}
func (RES_94) String() string {
	return "RES 2,H"
}

// SET 1,A    code=0xcf
type SET_CF struct{}

func (SET_CF) Exec(cpu *CPU) {
	panic("TODO SET_CF")
}
func (SET_CF) Code() uint8 {
	return 0xCF
}
func (SET_CF) String() string {
	return "SET 1,A"
}

// SLA L    code=0x25
type SLA_25 struct{}

func (SLA_25) Exec(cpu *CPU) {
	panic("TODO SLA_25")
}
func (SLA_25) Code() uint8 {
	return 0x25
}
func (SLA_25) String() string {
	return "SLA L"
}

// SRL L    code=0x3d
type SRL_3D struct{}

func (SRL_3D) Exec(cpu *CPU) {
	panic("TODO SRL_3D")
}
func (SRL_3D) Code() uint8 {
	return 0x3D
}
func (SRL_3D) String() string {
	return "SRL L"
}

// BIT 2,L    code=0x55
type BIT_55 struct{}

func (BIT_55) Exec(cpu *CPU) {
	value := cpu.L
	var flags Flags
	if (value & (1 << 2)) == 0 {
		flags |= FLAGZ
	}
	flags |= FLAGH
	cpu.F = FlagRegister(flags)

	cpu.Cycles += 8
}
func (BIT_55) Code() uint8 {
	return 0x55
}
func (BIT_55) String() string {
	return "BIT 2,L"
}

// BIT 4,L    code=0x65
type BIT_65 struct{}

func (BIT_65) Exec(cpu *CPU) {
	value := cpu.L
	var flags Flags
	if (value & (1 << 4)) == 0 {
		flags |= FLAGZ
	}
	flags |= FLAGH
	cpu.F = FlagRegister(flags)

	cpu.Cycles += 8
}
func (BIT_65) Code() uint8 {
	return 0x65
}
func (BIT_65) String() string {
	return "BIT 4,L"
}

// BIT 5,H    code=0x6c
type BIT_6C struct{}

func (BIT_6C) Exec(cpu *CPU) {
	value := cpu.H
	var flags Flags
	if (value & (1 << 5)) == 0 {
		flags |= FLAGZ
	}
	flags |= FLAGH
	cpu.F = FlagRegister(flags)

	cpu.Cycles += 8
}
func (BIT_6C) Code() uint8 {
	return 0x6C
}
func (BIT_6C) String() string {
	return "BIT 5,H"
}

// BIT 7,H    code=0x7c
type BIT_7C struct{}

func (BIT_7C) Exec(cpu *CPU) {
	value := cpu.H
	var flags Flags
	if (value & (1 << 7)) == 0 {
		flags |= FLAGZ
	}
	flags |= FLAGH
	cpu.F = FlagRegister(flags)

	cpu.Cycles += 8
}
func (BIT_7C) Code() uint8 {
	return 0x7C
}
func (BIT_7C) String() string {
	return "BIT 7,H"
}

// RES 4,H    code=0xa4
type RES_A4 struct{}

func (RES_A4) Exec(cpu *CPU) {
	panic("TODO RES_A4")
}
func (RES_A4) Code() uint8 {
	return 0xA4
}
func (RES_A4) String() string {
	return "RES 4,H"
}

// RES 4,L    code=0xa5
type RES_A5 struct{}

func (RES_A5) Exec(cpu *CPU) {
	panic("TODO RES_A5")
}
func (RES_A5) Code() uint8 {
	return 0xA5
}
func (RES_A5) String() string {
	return "RES 4,L"
}

// RRC H    code=0x0c
type RRC_0C struct{}

func (RRC_0C) Exec(cpu *CPU) {
	res, flags := rotate(cpu.H, 1, cpu.F, true)
	cpu.H = res
	cpu.F = flags

	cpu.Cycles += 8
}
func (RRC_0C) Code() uint8 {
	return 0xC
}
func (RRC_0C) String() string {
	return "RRC H"
}

// RR H    code=0x1c
type RR_1C struct{}

func (RR_1C) Exec(cpu *CPU) {
	res, flags := rotate(cpu.H, 1, cpu.F, false)
	cpu.H = res
	cpu.F = flags

	cpu.Cycles += 8
}
func (RR_1C) Code() uint8 {
	return 0x1C
}
func (RR_1C) String() string {
	return "RR H"
}

// BIT 3,H    code=0x5c
type BIT_5C struct{}

func (BIT_5C) Exec(cpu *CPU) {
	value := cpu.H
	var flags Flags
	if (value & (1 << 3)) == 0 {
		flags |= FLAGZ
	}
	flags |= FLAGH
	cpu.F = FlagRegister(flags)

	cpu.Cycles += 8
}
func (BIT_5C) Code() uint8 {
	return 0x5C
}
func (BIT_5C) String() string {
	return "BIT 3,H"
}

// RES 0,B    code=0x80
type RES_80 struct{}

func (RES_80) Exec(cpu *CPU) {
	panic("TODO RES_80")
}
func (RES_80) Code() uint8 {
	return 0x80
}
func (RES_80) String() string {
	return "RES 0,B"
}

// SET 1,H    code=0xcc
type SET_CC struct{}

func (SET_CC) Exec(cpu *CPU) {
	panic("TODO SET_CC")
}
func (SET_CC) Code() uint8 {
	return 0xCC
}
func (SET_CC) String() string {
	return "SET 1,H"
}

// SET 0,E    code=0xc3
type SET_C3 struct{}

func (SET_C3) Exec(cpu *CPU) {
	panic("TODO SET_C3")
}
func (SET_C3) Code() uint8 {
	return 0xC3
}
func (SET_C3) String() string {
	return "SET 0,E"
}

// SET 2,C    code=0xd1
type SET_D1 struct{}

func (SET_D1) Exec(cpu *CPU) {
	panic("TODO SET_D1")
}
func (SET_D1) Code() uint8 {
	return 0xD1
}
func (SET_D1) String() string {
	return "SET 2,C"
}

// BIT 5,L    code=0x6d
type BIT_6D struct{}

func (BIT_6D) Exec(cpu *CPU) {
	value := cpu.L
	var flags Flags
	if (value & (1 << 5)) == 0 {
		flags |= FLAGZ
	}
	flags |= FLAGH
	cpu.F = FlagRegister(flags)

	cpu.Cycles += 8
}
func (BIT_6D) Code() uint8 {
	return 0x6D
}
func (BIT_6D) String() string {
	return "BIT 5,L"
}

// SET 7,A    code=0xff
type SET_FF struct{}

func (SET_FF) Exec(cpu *CPU) {
	panic("TODO SET_FF")
}
func (SET_FF) Code() uint8 {
	return 0xFF
}
func (SET_FF) String() string {
	return "SET 7,A"
}

// SET 0,D    code=0xc2
type SET_C2 struct{}

func (SET_C2) Exec(cpu *CPU) {
	panic("TODO SET_C2")
}
func (SET_C2) Code() uint8 {
	return 0xC2
}
func (SET_C2) String() string {
	return "SET 0,D"
}

// RES 7,L    code=0xbd
type RES_BD struct{}

func (RES_BD) Exec(cpu *CPU) {
	panic("TODO RES_BD")
}
func (RES_BD) Code() uint8 {
	return 0xBD
}
func (RES_BD) String() string {
	return "RES 7,L"
}

// SRA D    code=0x2a
type SRA_2A struct{}

func (SRA_2A) Exec(cpu *CPU) {
	panic("TODO SRA_2A")
}
func (SRA_2A) Code() uint8 {
	return 0x2A
}
func (SRA_2A) String() string {
	return "SRA D"
}

// SRL E    code=0x3b
type SRL_3B struct{}

func (SRL_3B) Exec(cpu *CPU) {
	panic("TODO SRL_3B")
}
func (SRL_3B) Code() uint8 {
	return 0x3B
}
func (SRL_3B) String() string {
	return "SRL E"
}

// BIT 1,D    code=0x4a
type BIT_4A struct{}

func (BIT_4A) Exec(cpu *CPU) {
	value := cpu.D
	var flags Flags
	if (value & (1 << 1)) == 0 {
		flags |= FLAGZ
	}
	flags |= FLAGH
	cpu.F = FlagRegister(flags)

	cpu.Cycles += 8
}
func (BIT_4A) Code() uint8 {
	return 0x4A
}
func (BIT_4A) String() string {
	return "BIT 1,D"
}

// RES 6,(HL)    code=0xb6
type RES_B6 struct{}

func (RES_B6) Exec(cpu *CPU) {
	panic("TODO RES_B6")
}
func (RES_B6) Code() uint8 {
	return 0xB6
}
func (RES_B6) String() string {
	return "RES 6,(HL)"
}

// SET 3,E    code=0xdb
type SET_DB struct{}

func (SET_DB) Exec(cpu *CPU) {
	panic("TODO SET_DB")
}
func (SET_DB) Code() uint8 {
	return 0xDB
}
func (SET_DB) String() string {
	return "SET 3,E"
}

// RRC C    code=0x09
type RRC_09 struct{}

func (RRC_09) Exec(cpu *CPU) {
	res, flags := rotate(cpu.C, 1, cpu.F, true)
	cpu.C = res
	cpu.F = flags

	cpu.Cycles += 8
}
func (RRC_09) Code() uint8 {
	return 0x9
}
func (RRC_09) String() string {
	return "RRC C"
}

// SRL C    code=0x39
type SRL_39 struct{}

func (SRL_39) Exec(cpu *CPU) {
	panic("TODO SRL_39")
}
func (SRL_39) Code() uint8 {
	return 0x39
}
func (SRL_39) String() string {
	return "SRL C"
}

// BIT 3,(HL)    code=0x5e
type BIT_5E struct{}

func (BIT_5E) Exec(cpu *CPU) {
	value := cpu.loadU8(cpu.HL())
	var flags Flags
	if (value & (1 << 3)) == 0 {
		flags |= FLAGZ
	}
	flags |= FLAGH
	cpu.F = FlagRegister(flags)

	cpu.Cycles += 12
}
func (BIT_5E) Code() uint8 {
	return 0x5E
}
func (BIT_5E) String() string {
	return "BIT 3,(HL)"
}

// BIT 4,D    code=0x62
type BIT_62 struct{}

func (BIT_62) Exec(cpu *CPU) {
	value := cpu.D
	var flags Flags
	if (value & (1 << 4)) == 0 {
		flags |= FLAGZ
	}
	flags |= FLAGH
	cpu.F = FlagRegister(flags)

	cpu.Cycles += 8
}
func (BIT_62) Code() uint8 {
	return 0x62
}
func (BIT_62) String() string {
	return "BIT 4,D"
}

// SET 0,L    code=0xc5
type SET_C5 struct{}

func (SET_C5) Exec(cpu *CPU) {
	panic("TODO SET_C5")
}
func (SET_C5) Code() uint8 {
	return 0xC5
}
func (SET_C5) String() string {
	return "SET 0,L"
}

// SLA D    code=0x22
type SLA_22 struct{}

func (SLA_22) Exec(cpu *CPU) {
	panic("TODO SLA_22")
}
func (SLA_22) Code() uint8 {
	return 0x22
}
func (SLA_22) String() string {
	return "SLA D"
}

// BIT 4,B    code=0x60
type BIT_60 struct{}

func (BIT_60) Exec(cpu *CPU) {
	value := cpu.B
	var flags Flags
	if (value & (1 << 4)) == 0 {
		flags |= FLAGZ
	}
	flags |= FLAGH
	cpu.F = FlagRegister(flags)

	cpu.Cycles += 8
}
func (BIT_60) Code() uint8 {
	return 0x60
}
func (BIT_60) String() string {
	return "BIT 4,B"
}

// BIT 7,A    code=0x7f
type BIT_7F struct{}

func (BIT_7F) Exec(cpu *CPU) {
	value := cpu.A
	var flags Flags
	if (value & (1 << 7)) == 0 {
		flags |= FLAGZ
	}
	flags |= FLAGH
	cpu.F = FlagRegister(flags)

	cpu.Cycles += 8
}
func (BIT_7F) Code() uint8 {
	return 0x7F
}
func (BIT_7F) String() string {
	return "BIT 7,A"
}

// RES 0,L    code=0x85
type RES_85 struct{}

func (RES_85) Exec(cpu *CPU) {
	panic("TODO RES_85")
}
func (RES_85) Code() uint8 {
	return 0x85
}
func (RES_85) String() string {
	return "RES 0,L"
}

// SET 1,L    code=0xcd
type SET_CD struct{}

func (SET_CD) Exec(cpu *CPU) {
	panic("TODO SET_CD")
}
func (SET_CD) Code() uint8 {
	return 0xCD
}
func (SET_CD) String() string {
	return "SET 1,L"
}

// SET 5,E    code=0xeb
type SET_EB struct{}

func (SET_EB) Exec(cpu *CPU) {
	panic("TODO SET_EB")
}
func (SET_EB) Code() uint8 {
	return 0xEB
}
func (SET_EB) String() string {
	return "SET 5,E"
}

// RR E    code=0x1b
type RR_1B struct{}

func (RR_1B) Exec(cpu *CPU) {
	res, flags := rotate(cpu.E, 1, cpu.F, false)
	cpu.E = res
	cpu.F = flags

	cpu.Cycles += 8
}
func (RR_1B) Code() uint8 {
	return 0x1B
}
func (RR_1B) String() string {
	return "RR E"
}

// SRL H    code=0x3c
type SRL_3C struct{}

func (SRL_3C) Exec(cpu *CPU) {
	panic("TODO SRL_3C")
}
func (SRL_3C) Code() uint8 {
	return 0x3C
}
func (SRL_3C) String() string {
	return "SRL H"
}

// SET 1,D    code=0xca
type SET_CA struct{}

func (SET_CA) Exec(cpu *CPU) {
	panic("TODO SET_CA")
}
func (SET_CA) Code() uint8 {
	return 0xCA
}
func (SET_CA) String() string {
	return "SET 1,D"
}

// SET 2,D    code=0xd2
type SET_D2 struct{}

func (SET_D2) Exec(cpu *CPU) {
	panic("TODO SET_D2")
}
func (SET_D2) Code() uint8 {
	return 0xD2
}
func (SET_D2) String() string {
	return "SET 2,D"
}

// BIT 5,D    code=0x6a
type BIT_6A struct{}

func (BIT_6A) Exec(cpu *CPU) {
	value := cpu.D
	var flags Flags
	if (value & (1 << 5)) == 0 {
		flags |= FLAGZ
	}
	flags |= FLAGH
	cpu.F = FlagRegister(flags)

	cpu.Cycles += 8
}
func (BIT_6A) Code() uint8 {
	return 0x6A
}
func (BIT_6A) String() string {
	return "BIT 5,D"
}

// RR L    code=0x1d
type RR_1D struct{}

func (RR_1D) Exec(cpu *CPU) {
	res, flags := rotate(cpu.L, 1, cpu.F, false)
	cpu.L = res
	cpu.F = flags

	cpu.Cycles += 8
}
func (RR_1D) Code() uint8 {
	return 0x1D
}
func (RR_1D) String() string {
	return "RR L"
}

// BIT 3,C    code=0x59
type BIT_59 struct{}

func (BIT_59) Exec(cpu *CPU) {
	value := cpu.C
	var flags Flags
	if (value & (1 << 3)) == 0 {
		flags |= FLAGZ
	}
	flags |= FLAGH
	cpu.F = FlagRegister(flags)

	cpu.Cycles += 8
}
func (BIT_59) Code() uint8 {
	return 0x59
}
func (BIT_59) String() string {
	return "BIT 3,C"
}

// BIT 3,E    code=0x5b
type BIT_5B struct{}

func (BIT_5B) Exec(cpu *CPU) {
	value := cpu.E
	var flags Flags
	if (value & (1 << 3)) == 0 {
		flags |= FLAGZ
	}
	flags |= FLAGH
	cpu.F = FlagRegister(flags)

	cpu.Cycles += 8
}
func (BIT_5B) Code() uint8 {
	return 0x5B
}
func (BIT_5B) String() string {
	return "BIT 3,E"
}

// BIT 6,L    code=0x75
type BIT_75 struct{}

func (BIT_75) Exec(cpu *CPU) {
	value := cpu.L
	var flags Flags
	if (value & (1 << 6)) == 0 {
		flags |= FLAGZ
	}
	flags |= FLAGH
	cpu.F = FlagRegister(flags)

	cpu.Cycles += 8
}
func (BIT_75) Code() uint8 {
	return 0x75
}
func (BIT_75) String() string {
	return "BIT 6,L"
}

// RES 0,A    code=0x87
type RES_87 struct{}

func (RES_87) Exec(cpu *CPU) {
	panic("TODO RES_87")
}
func (RES_87) Code() uint8 {
	return 0x87
}
func (RES_87) String() string {
	return "RES 0,A"
}

// RES 2,B    code=0x90
type RES_90 struct{}

func (RES_90) Exec(cpu *CPU) {
	panic("TODO RES_90")
}
func (RES_90) Code() uint8 {
	return 0x90
}
func (RES_90) String() string {
	return "RES 2,B"
}

// RLC L    code=0x05
type RLC_05 struct{}

func (RLC_05) Exec(cpu *CPU) {
	res, flags := rotate(cpu.L, 0, cpu.F, true)
	cpu.L = res
	cpu.F = flags

	cpu.Cycles += 8
}
func (RLC_05) Code() uint8 {
	return 0x5
}
func (RLC_05) String() string {
	return "RLC L"
}

// RES 3,A    code=0x9f
type RES_9F struct{}

func (RES_9F) Exec(cpu *CPU) {
	panic("TODO RES_9F")
}
func (RES_9F) Code() uint8 {
	return 0x9F
}
func (RES_9F) String() string {
	return "RES 3,A"
}

// RES 5,L    code=0xad
type RES_AD struct{}

func (RES_AD) Exec(cpu *CPU) {
	panic("TODO RES_AD")
}
func (RES_AD) Code() uint8 {
	return 0xAD
}
func (RES_AD) String() string {
	return "RES 5,L"
}

// RES 3,E    code=0x9b
type RES_9B struct{}

func (RES_9B) Exec(cpu *CPU) {
	panic("TODO RES_9B")
}
func (RES_9B) Code() uint8 {
	return 0x9B
}
func (RES_9B) String() string {
	return "RES 3,E"
}

// SET 4,L    code=0xe5
type SET_E5 struct{}

func (SET_E5) Exec(cpu *CPU) {
	panic("TODO SET_E5")
}
func (SET_E5) Code() uint8 {
	return 0xE5
}
func (SET_E5) String() string {
	return "SET 4,L"
}

// SET 5,C    code=0xe9
type SET_E9 struct{}

func (SET_E9) Exec(cpu *CPU) {
	panic("TODO SET_E9")
}
func (SET_E9) Code() uint8 {
	return 0xE9
}
func (SET_E9) String() string {
	return "SET 5,C"
}

// SET 7,C    code=0xf9
type SET_F9 struct{}

func (SET_F9) Exec(cpu *CPU) {
	panic("TODO SET_F9")
}
func (SET_F9) Code() uint8 {
	return 0xF9
}
func (SET_F9) String() string {
	return "SET 7,C"
}

// RES 2,C    code=0x91
type RES_91 struct{}

func (RES_91) Exec(cpu *CPU) {
	panic("TODO RES_91")
}
func (RES_91) Code() uint8 {
	return 0x91
}
func (RES_91) String() string {
	return "RES 2,C"
}

// RES 6,D    code=0xb2
type RES_B2 struct{}

func (RES_B2) Exec(cpu *CPU) {
	panic("TODO RES_B2")
}
func (RES_B2) Code() uint8 {
	return 0xB2
}
func (RES_B2) String() string {
	return "RES 6,D"
}

// SET 7,H    code=0xfc
type SET_FC struct{}

func (SET_FC) Exec(cpu *CPU) {
	panic("TODO SET_FC")
}
func (SET_FC) Code() uint8 {
	return 0xFC
}
func (SET_FC) String() string {
	return "SET 7,H"
}

// RRC L    code=0x0d
type RRC_0D struct{}

func (RRC_0D) Exec(cpu *CPU) {
	res, flags := rotate(cpu.L, 1, cpu.F, true)
	cpu.L = res
	cpu.F = flags

	cpu.Cycles += 8
}
func (RRC_0D) Code() uint8 {
	return 0xD
}
func (RRC_0D) String() string {
	return "RRC L"
}

// BIT 0,B    code=0x40
type BIT_40 struct{}

func (BIT_40) Exec(cpu *CPU) {
	value := cpu.B
	var flags Flags
	if (value & (1 << 0)) == 0 {
		flags |= FLAGZ
	}
	flags |= FLAGH
	cpu.F = FlagRegister(flags)

	cpu.Cycles += 8
}
func (BIT_40) Code() uint8 {
	return 0x40
}
func (BIT_40) String() string {
	return "BIT 0,B"
}

// BIT 0,H    code=0x44
type BIT_44 struct{}

func (BIT_44) Exec(cpu *CPU) {
	value := cpu.H
	var flags Flags
	if (value & (1 << 0)) == 0 {
		flags |= FLAGZ
	}
	flags |= FLAGH
	cpu.F = FlagRegister(flags)

	cpu.Cycles += 8
}
func (BIT_44) Code() uint8 {
	return 0x44
}
func (BIT_44) String() string {
	return "BIT 0,H"
}

// RES 5,H    code=0xac
type RES_AC struct{}

func (RES_AC) Exec(cpu *CPU) {
	panic("TODO RES_AC")
}
func (RES_AC) Code() uint8 {
	return 0xAC
}
func (RES_AC) String() string {
	return "RES 5,H"
}

// RRC B    code=0x08
type RRC_08 struct{}

func (RRC_08) Exec(cpu *CPU) {
	res, flags := rotate(cpu.B, 1, cpu.F, true)
	cpu.B = res
	cpu.F = flags

	cpu.Cycles += 8
}
func (RRC_08) Code() uint8 {
	return 0x8
}
func (RRC_08) String() string {
	return "RRC B"
}

// BIT 4,(HL)    code=0x66
type BIT_66 struct{}

func (BIT_66) Exec(cpu *CPU) {
	value := cpu.loadU8(cpu.HL())
	var flags Flags
	if (value & (1 << 4)) == 0 {
		flags |= FLAGZ
	}
	flags |= FLAGH
	cpu.F = FlagRegister(flags)

	cpu.Cycles += 12
}
func (BIT_66) Code() uint8 {
	return 0x66
}
func (BIT_66) String() string {
	return "BIT 4,(HL)"
}

// BIT 5,A    code=0x6f
type BIT_6F struct{}

func (BIT_6F) Exec(cpu *CPU) {
	value := cpu.A
	var flags Flags
	if (value & (1 << 5)) == 0 {
		flags |= FLAGZ
	}
	flags |= FLAGH
	cpu.F = FlagRegister(flags)

	cpu.Cycles += 8
}
func (BIT_6F) Code() uint8 {
	return 0x6F
}
func (BIT_6F) String() string {
	return "BIT 5,A"
}

// SET 4,B    code=0xe0
type SET_E0 struct{}

func (SET_E0) Exec(cpu *CPU) {
	panic("TODO SET_E0")
}
func (SET_E0) Code() uint8 {
	return 0xE0
}
func (SET_E0) String() string {
	return "SET 4,B"
}

// RRC E    code=0x0b
type RRC_0B struct{}

func (RRC_0B) Exec(cpu *CPU) {
	res, flags := rotate(cpu.E, 1, cpu.F, true)
	cpu.E = res
	cpu.F = flags

	cpu.Cycles += 8
}
func (RRC_0B) Code() uint8 {
	return 0xB
}
func (RRC_0B) String() string {
	return "RRC E"
}

// BIT 4,E    code=0x63
type BIT_63 struct{}

func (BIT_63) Exec(cpu *CPU) {
	value := cpu.E
	var flags Flags
	if (value & (1 << 4)) == 0 {
		flags |= FLAGZ
	}
	flags |= FLAGH
	cpu.F = FlagRegister(flags)

	cpu.Cycles += 8
}
func (BIT_63) Code() uint8 {
	return 0x63
}
func (BIT_63) String() string {
	return "BIT 4,E"
}

// BIT 2,H    code=0x54
type BIT_54 struct{}

func (BIT_54) Exec(cpu *CPU) {
	value := cpu.H
	var flags Flags
	if (value & (1 << 2)) == 0 {
		flags |= FLAGZ
	}
	flags |= FLAGH
	cpu.F = FlagRegister(flags)

	cpu.Cycles += 8
}
func (BIT_54) Code() uint8 {
	return 0x54
}
func (BIT_54) String() string {
	return "BIT 2,H"
}

// RES 0,H    code=0x84
type RES_84 struct{}

func (RES_84) Exec(cpu *CPU) {
	panic("TODO RES_84")
}
func (RES_84) Code() uint8 {
	return 0x84
}
func (RES_84) String() string {
	return "RES 0,H"
}

// RES 1,D    code=0x8a
type RES_8A struct{}

func (RES_8A) Exec(cpu *CPU) {
	panic("TODO RES_8A")
}
func (RES_8A) Code() uint8 {
	return 0x8A
}
func (RES_8A) String() string {
	return "RES 1,D"
}

// RES 5,A    code=0xaf
type RES_AF struct{}

func (RES_AF) Exec(cpu *CPU) {
	panic("TODO RES_AF")
}
func (RES_AF) Code() uint8 {
	return 0xAF
}
func (RES_AF) String() string {
	return "RES 5,A"
}

// SET 6,C    code=0xf1
type SET_F1 struct{}

func (SET_F1) Exec(cpu *CPU) {
	panic("TODO SET_F1")
}
func (SET_F1) Code() uint8 {
	return 0xF1
}
func (SET_F1) String() string {
	return "SET 6,C"
}

// RLC D    code=0x02
type RLC_02 struct{}

func (RLC_02) Exec(cpu *CPU) {
	res, flags := rotate(cpu.D, 0, cpu.F, true)
	cpu.D = res
	cpu.F = flags

	cpu.Cycles += 8
}
func (RLC_02) Code() uint8 {
	return 0x2
}
func (RLC_02) String() string {
	return "RLC D"
}

// SRA L    code=0x2d
type SRA_2D struct{}

func (SRA_2D) Exec(cpu *CPU) {
	panic("TODO SRA_2D")
}
func (SRA_2D) Code() uint8 {
	return 0x2D
}
func (SRA_2D) String() string {
	return "SRA L"
}

// BIT 5,(HL)    code=0x6e
type BIT_6E struct{}

func (BIT_6E) Exec(cpu *CPU) {
	value := cpu.loadU8(cpu.HL())
	var flags Flags
	if (value & (1 << 5)) == 0 {
		flags |= FLAGZ
	}
	flags |= FLAGH
	cpu.F = FlagRegister(flags)

	cpu.Cycles += 12
}
func (BIT_6E) Code() uint8 {
	return 0x6E
}
func (BIT_6E) String() string {
	return "BIT 5,(HL)"
}

// RES 4,E    code=0xa3
type RES_A3 struct{}

func (RES_A3) Exec(cpu *CPU) {
	panic("TODO RES_A3")
}
func (RES_A3) Code() uint8 {
	return 0xA3
}
func (RES_A3) String() string {
	return "RES 4,E"
}

// SET 3,A    code=0xdf
type SET_DF struct{}

func (SET_DF) Exec(cpu *CPU) {
	panic("TODO SET_DF")
}
func (SET_DF) Code() uint8 {
	return 0xDF
}
func (SET_DF) String() string {
	return "SET 3,A"
}

// SET 4,E    code=0xe3
type SET_E3 struct{}

func (SET_E3) Exec(cpu *CPU) {
	panic("TODO SET_E3")
}
func (SET_E3) Code() uint8 {
	return 0xE3
}
func (SET_E3) String() string {
	return "SET 4,E"
}

// RES 3,D    code=0x9a
type RES_9A struct{}

func (RES_9A) Exec(cpu *CPU) {
	panic("TODO RES_9A")
}
func (RES_9A) Code() uint8 {
	return 0x9A
}
func (RES_9A) String() string {
	return "RES 3,D"
}

// RLC B    code=0x00
type RLC_00 struct{}

func (RLC_00) Exec(cpu *CPU) {
	res, flags := rotate(cpu.B, 0, cpu.F, true)
	cpu.B = res
	cpu.F = flags

	cpu.Cycles += 8
}
func (RLC_00) Code() uint8 {
	return 0x0
}
func (RLC_00) String() string {
	return "RLC B"
}

// RES 2,L    code=0x95
type RES_95 struct{}

func (RES_95) Exec(cpu *CPU) {
	panic("TODO RES_95")
}
func (RES_95) Code() uint8 {
	return 0x95
}
func (RES_95) String() string {
	return "RES 2,L"
}

// SLA B    code=0x20
type SLA_20 struct{}

func (SLA_20) Exec(cpu *CPU) {
	panic("TODO SLA_20")
}
func (SLA_20) Code() uint8 {
	return 0x20
}
func (SLA_20) String() string {
	return "SLA B"
}

// SRA H    code=0x2c
type SRA_2C struct{}

func (SRA_2C) Exec(cpu *CPU) {
	panic("TODO SRA_2C")
}
func (SRA_2C) Code() uint8 {
	return 0x2C
}
func (SRA_2C) String() string {
	return "SRA H"
}

// RES 0,(HL)    code=0x86
type RES_86 struct{}

func (RES_86) Exec(cpu *CPU) {
	panic("TODO RES_86")
}
func (RES_86) Code() uint8 {
	return 0x86
}
func (RES_86) String() string {
	return "RES 0,(HL)"
}

// RES 7,H    code=0xbc
type RES_BC struct{}

func (RES_BC) Exec(cpu *CPU) {
	panic("TODO RES_BC")
}
func (RES_BC) Code() uint8 {
	return 0xBC
}
func (RES_BC) String() string {
	return "RES 7,H"
}

// BIT 3,L    code=0x5d
type BIT_5D struct{}

func (BIT_5D) Exec(cpu *CPU) {
	value := cpu.L
	var flags Flags
	if (value & (1 << 3)) == 0 {
		flags |= FLAGZ
	}
	flags |= FLAGH
	cpu.F = FlagRegister(flags)

	cpu.Cycles += 8
}
func (BIT_5D) Code() uint8 {
	return 0x5D
}
func (BIT_5D) String() string {
	return "BIT 3,L"
}

// BIT 6,E    code=0x73
type BIT_73 struct{}

func (BIT_73) Exec(cpu *CPU) {
	value := cpu.E
	var flags Flags
	if (value & (1 << 6)) == 0 {
		flags |= FLAGZ
	}
	flags |= FLAGH
	cpu.F = FlagRegister(flags)

	cpu.Cycles += 8
}
func (BIT_73) Code() uint8 {
	return 0x73
}
func (BIT_73) String() string {
	return "BIT 6,E"
}

// BIT 7,E    code=0x7b
type BIT_7B struct{}

func (BIT_7B) Exec(cpu *CPU) {
	value := cpu.E
	var flags Flags
	if (value & (1 << 7)) == 0 {
		flags |= FLAGZ
	}
	flags |= FLAGH
	cpu.F = FlagRegister(flags)

	cpu.Cycles += 8
}
func (BIT_7B) Code() uint8 {
	return 0x7B
}
func (BIT_7B) String() string {
	return "BIT 7,E"
}

// RL H    code=0x14
type RL_14 struct{}

func (RL_14) Exec(cpu *CPU) {
	res, flags := rotate(cpu.H, 0, cpu.F, false)
	cpu.H = res
	cpu.F = flags

	cpu.Cycles += 8
}
func (RL_14) Code() uint8 {
	return 0x14
}
func (RL_14) String() string {
	return "RL H"
}

// BIT 0,A    code=0x47
type BIT_47 struct{}

func (BIT_47) Exec(cpu *CPU) {
	value := cpu.A
	var flags Flags
	if (value & (1 << 0)) == 0 {
		flags |= FLAGZ
	}
	flags |= FLAGH
	cpu.F = FlagRegister(flags)

	cpu.Cycles += 8
}
func (BIT_47) Code() uint8 {
	return 0x47
}
func (BIT_47) String() string {
	return "BIT 0,A"
}

// BIT 7,L    code=0x7d
type BIT_7D struct{}

func (BIT_7D) Exec(cpu *CPU) {
	value := cpu.L
	var flags Flags
	if (value & (1 << 7)) == 0 {
		flags |= FLAGZ
	}
	flags |= FLAGH
	cpu.F = FlagRegister(flags)

	cpu.Cycles += 8
}
func (BIT_7D) Code() uint8 {
	return 0x7D
}
func (BIT_7D) String() string {
	return "BIT 7,L"
}

// RLC H    code=0x04
type RLC_04 struct{}

func (RLC_04) Exec(cpu *CPU) {
	res, flags := rotate(cpu.H, 0, cpu.F, true)
	cpu.H = res
	cpu.F = flags

	cpu.Cycles += 8
}
func (RLC_04) Code() uint8 {
	return 0x4
}
func (RLC_04) String() string {
	return "RLC H"
}

// SLA A    code=0x27
type SLA_27 struct{}

func (SLA_27) Exec(cpu *CPU) {
	panic("TODO SLA_27")
}
func (SLA_27) Code() uint8 {
	return 0x27
}
func (SLA_27) String() string {
	return "SLA A"
}

// SWAP L    code=0x35
type SWAP_35 struct{}

func (SWAP_35) Exec(cpu *CPU) {
	panic("TODO SWAP_35")
}
func (SWAP_35) Code() uint8 {
	return 0x35
}
func (SWAP_35) String() string {
	return "SWAP L"
}

// BIT 7,B    code=0x78
type BIT_78 struct{}

func (BIT_78) Exec(cpu *CPU) {
	value := cpu.B
	var flags Flags
	if (value & (1 << 7)) == 0 {
		flags |= FLAGZ
	}
	flags |= FLAGH
	cpu.F = FlagRegister(flags)

	cpu.Cycles += 8
}
func (BIT_78) Code() uint8 {
	return 0x78
}
func (BIT_78) String() string {
	return "BIT 7,B"
}

// RES 1,C    code=0x89
type RES_89 struct{}

func (RES_89) Exec(cpu *CPU) {
	panic("TODO RES_89")
}
func (RES_89) Code() uint8 {
	return 0x89
}
func (RES_89) String() string {
	return "RES 1,C"
}

// RES 5,B    code=0xa8
type RES_A8 struct{}

func (RES_A8) Exec(cpu *CPU) {
	panic("TODO RES_A8")
}
func (RES_A8) Code() uint8 {
	return 0xA8
}
func (RES_A8) String() string {
	return "RES 5,B"
}

// RRC A    code=0x0f
type RRC_0F struct{}

func (RRC_0F) Exec(cpu *CPU) {
	res, flags := rotate(cpu.A, 1, cpu.F, true)
	cpu.A = res
	cpu.F = flags

	cpu.Cycles += 8
}
func (RRC_0F) Code() uint8 {
	return 0xF
}
func (RRC_0F) String() string {
	return "RRC A"
}

// RL B    code=0x10
type RL_10 struct{}

func (RL_10) Exec(cpu *CPU) {
	res, flags := rotate(cpu.B, 0, cpu.F, false)
	cpu.B = res
	cpu.F = flags

	cpu.Cycles += 8
}
func (RL_10) Code() uint8 {
	return 0x10
}
func (RL_10) String() string {
	return "RL B"
}

// RR (HL)    code=0x1e
type RR_1E struct{}

func (RR_1E) Exec(cpu *CPU) {
	res, flags := rotate(cpu.loadU8(cpu.HL()), 1, cpu.F, false)
	cpu.WriteMemory(cpu.HL(), res)
	cpu.F = flags

	cpu.Cycles += 16
}
func (RR_1E) Code() uint8 {
	return 0x1E
}
func (RR_1E) String() string {
	return "RR (HL)"
}

// SLA H    code=0x24
type SLA_24 struct{}

func (SLA_24) Exec(cpu *CPU) {
	panic("TODO SLA_24")
}
func (SLA_24) Code() uint8 {
	return 0x24
}
func (SLA_24) String() string {
	return "SLA H"
}

// SRA B    code=0x28
type SRA_28 struct{}

func (SRA_28) Exec(cpu *CPU) {
	panic("TODO SRA_28")
}
func (SRA_28) Code() uint8 {
	return 0x28
}
func (SRA_28) String() string {
	return "SRA B"
}

// SET 4,D    code=0xe2
type SET_E2 struct{}

func (SET_E2) Exec(cpu *CPU) {
	panic("TODO SET_E2")
}
func (SET_E2) Code() uint8 {
	return 0xE2
}
func (SET_E2) String() string {
	return "SET 4,D"
}

// SET 4,(HL)    code=0xe6
type SET_E6 struct{}

func (SET_E6) Exec(cpu *CPU) {
	panic("TODO SET_E6")
}
func (SET_E6) Code() uint8 {
	return 0xE6
}
func (SET_E6) String() string {
	return "SET 4,(HL)"
}

// SET 6,E    code=0xf3
type SET_F3 struct{}

func (SET_F3) Exec(cpu *CPU) {
	panic("TODO SET_F3")
}
func (SET_F3) Code() uint8 {
	return 0xF3
}
func (SET_F3) String() string {
	return "SET 6,E"
}

// BIT 5,C    code=0x69
type BIT_69 struct{}

func (BIT_69) Exec(cpu *CPU) {
	value := cpu.C
	var flags Flags
	if (value & (1 << 5)) == 0 {
		flags |= FLAGZ
	}
	flags |= FLAGH
	cpu.F = FlagRegister(flags)

	cpu.Cycles += 8
}
func (BIT_69) Code() uint8 {
	return 0x69
}
func (BIT_69) String() string {
	return "BIT 5,C"
}

// SET 2,A    code=0xd7
type SET_D7 struct{}

func (SET_D7) Exec(cpu *CPU) {
	panic("TODO SET_D7")
}
func (SET_D7) Code() uint8 {
	return 0xD7
}
func (SET_D7) String() string {
	return "SET 2,A"
}

// SLA C    code=0x21
type SLA_21 struct{}

func (SLA_21) Exec(cpu *CPU) {
	panic("TODO SLA_21")
}
func (SLA_21) Code() uint8 {
	return 0x21
}
func (SLA_21) String() string {
	return "SLA C"
}

// BIT 0,L    code=0x45
type BIT_45 struct{}

func (BIT_45) Exec(cpu *CPU) {
	value := cpu.L
	var flags Flags
	if (value & (1 << 0)) == 0 {
		flags |= FLAGZ
	}
	flags |= FLAGH
	cpu.F = FlagRegister(flags)

	cpu.Cycles += 8
}
func (BIT_45) Code() uint8 {
	return 0x45
}
func (BIT_45) String() string {
	return "BIT 0,L"
}

// SET 1,B    code=0xc8
type SET_C8 struct{}

func (SET_C8) Exec(cpu *CPU) {
	panic("TODO SET_C8")
}
func (SET_C8) Code() uint8 {
	return 0xC8
}
func (SET_C8) String() string {
	return "SET 1,B"
}

// SET 2,B    code=0xd0
type SET_D0 struct{}

func (SET_D0) Exec(cpu *CPU) {
	panic("TODO SET_D0")
}
func (SET_D0) Code() uint8 {
	return 0xD0
}
func (SET_D0) String() string {
	return "SET 2,B"
}

// SWAP (HL)    code=0x36
type SWAP_36 struct{}

func (SWAP_36) Exec(cpu *CPU) {
	panic("TODO SWAP_36")
}
func (SWAP_36) Code() uint8 {
	return 0x36
}
func (SWAP_36) String() string {
	return "SWAP (HL)"
}

// BIT 7,C    code=0x79
type BIT_79 struct{}

func (BIT_79) Exec(cpu *CPU) {
	value := cpu.C
	var flags Flags
	if (value & (1 << 7)) == 0 {
		flags |= FLAGZ
	}
	flags |= FLAGH
	cpu.F = FlagRegister(flags)

	cpu.Cycles += 8
}
func (BIT_79) Code() uint8 {
	return 0x79
}
func (BIT_79) String() string {
	return "BIT 7,C"
}

// RES 1,A    code=0x8f
type RES_8F struct{}

func (RES_8F) Exec(cpu *CPU) {
	panic("TODO RES_8F")
}
func (RES_8F) Code() uint8 {
	return 0x8F
}
func (RES_8F) String() string {
	return "RES 1,A"
}

// RES 3,B    code=0x98
type RES_98 struct{}

func (RES_98) Exec(cpu *CPU) {
	panic("TODO RES_98")
}
func (RES_98) Code() uint8 {
	return 0x98
}
func (RES_98) String() string {
	return "RES 3,B"
}

// SET 6,(HL)    code=0xf6
type SET_F6 struct{}

func (SET_F6) Exec(cpu *CPU) {
	panic("TODO SET_F6")
}
func (SET_F6) Code() uint8 {
	return 0xF6
}
func (SET_F6) String() string {
	return "SET 6,(HL)"
}

// BIT 0,E    code=0x43
type BIT_43 struct{}

func (BIT_43) Exec(cpu *CPU) {
	value := cpu.E
	var flags Flags
	if (value & (1 << 0)) == 0 {
		flags |= FLAGZ
	}
	flags |= FLAGH
	cpu.F = FlagRegister(flags)

	cpu.Cycles += 8
}
func (BIT_43) Code() uint8 {
	return 0x43
}
func (BIT_43) String() string {
	return "BIT 0,E"
}

// BIT 1,C    code=0x49
type BIT_49 struct{}

func (BIT_49) Exec(cpu *CPU) {
	value := cpu.C
	var flags Flags
	if (value & (1 << 1)) == 0 {
		flags |= FLAGZ
	}
	flags |= FLAGH
	cpu.F = FlagRegister(flags)

	cpu.Cycles += 8
}
func (BIT_49) Code() uint8 {
	return 0x49
}
func (BIT_49) String() string {
	return "BIT 1,C"
}

// BIT 5,E    code=0x6b
type BIT_6B struct{}

func (BIT_6B) Exec(cpu *CPU) {
	value := cpu.E
	var flags Flags
	if (value & (1 << 5)) == 0 {
		flags |= FLAGZ
	}
	flags |= FLAGH
	cpu.F = FlagRegister(flags)

	cpu.Cycles += 8
}
func (BIT_6B) Code() uint8 {
	return 0x6B
}
func (BIT_6B) String() string {
	return "BIT 5,E"
}

// RES 4,C    code=0xa1
type RES_A1 struct{}

func (RES_A1) Exec(cpu *CPU) {
	panic("TODO RES_A1")
}
func (RES_A1) Code() uint8 {
	return 0xA1
}
func (RES_A1) String() string {
	return "RES 4,C"
}

// RES 6,C    code=0xb1
type RES_B1 struct{}

func (RES_B1) Exec(cpu *CPU) {
	panic("TODO RES_B1")
}
func (RES_B1) Code() uint8 {
	return 0xB1
}
func (RES_B1) String() string {
	return "RES 6,C"
}

// SET 1,E    code=0xcb
type SET_CB struct{}

func (SET_CB) Exec(cpu *CPU) {
	panic("TODO SET_CB")
}
func (SET_CB) Code() uint8 {
	return 0xCB
}
func (SET_CB) String() string {
	return "SET 1,E"
}

// SET 5,A    code=0xef
type SET_EF struct{}

func (SET_EF) Exec(cpu *CPU) {
	panic("TODO SET_EF")
}
func (SET_EF) Code() uint8 {
	return 0xEF
}
func (SET_EF) String() string {
	return "SET 5,A"
}

// RR C    code=0x19
type RR_19 struct{}

func (RR_19) Exec(cpu *CPU) {
	res, flags := rotate(cpu.C, 1, cpu.F, false)
	cpu.C = res
	cpu.F = flags

	cpu.Cycles += 8
}
func (RR_19) Code() uint8 {
	return 0x19
}
func (RR_19) String() string {
	return "RR C"
}

// BIT 6,B    code=0x70
type BIT_70 struct{}

func (BIT_70) Exec(cpu *CPU) {
	value := cpu.B
	var flags Flags
	if (value & (1 << 6)) == 0 {
		flags |= FLAGZ
	}
	flags |= FLAGH
	cpu.F = FlagRegister(flags)

	cpu.Cycles += 8
}
func (BIT_70) Code() uint8 {
	return 0x70
}
func (BIT_70) String() string {
	return "BIT 6,B"
}

// RES 7,B    code=0xb8
type RES_B8 struct{}

func (RES_B8) Exec(cpu *CPU) {
	panic("TODO RES_B8")
}
func (RES_B8) Code() uint8 {
	return 0xB8
}
func (RES_B8) String() string {
	return "RES 7,B"
}

// SET 4,A    code=0xe7
type SET_E7 struct{}

func (SET_E7) Exec(cpu *CPU) {
	panic("TODO SET_E7")
}
func (SET_E7) Code() uint8 {
	return 0xE7
}
func (SET_E7) String() string {
	return "SET 4,A"
}

// SET 6,B    code=0xf0
type SET_F0 struct{}

func (SET_F0) Exec(cpu *CPU) {
	panic("TODO SET_F0")
}
func (SET_F0) Code() uint8 {
	return 0xF0
}
func (SET_F0) String() string {
	return "SET 6,B"
}

// SET 6,H    code=0xf4
type SET_F4 struct{}

func (SET_F4) Exec(cpu *CPU) {
	panic("TODO SET_F4")
}
func (SET_F4) Code() uint8 {
	return 0xF4
}
func (SET_F4) String() string {
	return "SET 6,H"
}

// SWAP B    code=0x30
type SWAP_30 struct{}

func (SWAP_30) Exec(cpu *CPU) {
	panic("TODO SWAP_30")
}
func (SWAP_30) Code() uint8 {
	return 0x30
}
func (SWAP_30) String() string {
	return "SWAP B"
}

// SWAP E    code=0x33
type SWAP_33 struct{}

func (SWAP_33) Exec(cpu *CPU) {
	panic("TODO SWAP_33")
}
func (SWAP_33) Code() uint8 {
	return 0x33
}
func (SWAP_33) String() string {
	return "SWAP E"
}

// SRL D    code=0x3a
type SRL_3A struct{}

func (SRL_3A) Exec(cpu *CPU) {
	panic("TODO SRL_3A")
}
func (SRL_3A) Code() uint8 {
	return 0x3A
}
func (SRL_3A) String() string {
	return "SRL D"
}

// BIT 6,(HL)    code=0x76
type BIT_76 struct{}

func (BIT_76) Exec(cpu *CPU) {
	value := cpu.loadU8(cpu.HL())
	var flags Flags
	if (value & (1 << 6)) == 0 {
		flags |= FLAGZ
	}
	flags |= FLAGH
	cpu.F = FlagRegister(flags)

	cpu.Cycles += 12
}
func (BIT_76) Code() uint8 {
	return 0x76
}
func (BIT_76) String() string {
	return "BIT 6,(HL)"
}

// SET 2,H    code=0xd4
type SET_D4 struct{}

func (SET_D4) Exec(cpu *CPU) {
	panic("TODO SET_D4")
}
func (SET_D4) Code() uint8 {
	return 0xD4
}
func (SET_D4) String() string {
	return "SET 2,H"
}

// SET 7,L    code=0xfd
type SET_FD struct{}

func (SET_FD) Exec(cpu *CPU) {
	panic("TODO SET_FD")
}
func (SET_FD) Code() uint8 {
	return 0xFD
}
func (SET_FD) String() string {
	return "SET 7,L"
}

// BIT 4,C    code=0x61
type BIT_61 struct{}

func (BIT_61) Exec(cpu *CPU) {
	value := cpu.C
	var flags Flags
	if (value & (1 << 4)) == 0 {
		flags |= FLAGZ
	}
	flags |= FLAGH
	cpu.F = FlagRegister(flags)

	cpu.Cycles += 8
}
func (BIT_61) Code() uint8 {
	return 0x61
}
func (BIT_61) String() string {
	return "BIT 4,C"
}

// BIT 5,B    code=0x68
type BIT_68 struct{}

func (BIT_68) Exec(cpu *CPU) {
	value := cpu.B
	var flags Flags
	if (value & (1 << 5)) == 0 {
		flags |= FLAGZ
	}
	flags |= FLAGH
	cpu.F = FlagRegister(flags)

	cpu.Cycles += 8
}
func (BIT_68) Code() uint8 {
	return 0x68
}
func (BIT_68) String() string {
	return "BIT 5,B"
}

// RES 3,C    code=0x99
type RES_99 struct{}

func (RES_99) Exec(cpu *CPU) {
	panic("TODO RES_99")
}
func (RES_99) Code() uint8 {
	return 0x99
}
func (RES_99) String() string {
	return "RES 3,C"
}

// SET 5,D    code=0xea
type SET_EA struct{}

func (SET_EA) Exec(cpu *CPU) {
	panic("TODO SET_EA")
}
func (SET_EA) Code() uint8 {
	return 0xEA
}
func (SET_EA) String() string {
	return "SET 5,D"
}

// SET 5,H    code=0xec
type SET_EC struct{}

func (SET_EC) Exec(cpu *CPU) {
	panic("TODO SET_EC")
}
func (SET_EC) Code() uint8 {
	return 0xEC
}
func (SET_EC) String() string {
	return "SET 5,H"
}

// SET 6,D    code=0xf2
type SET_F2 struct{}

func (SET_F2) Exec(cpu *CPU) {
	panic("TODO SET_F2")
}
func (SET_F2) Code() uint8 {
	return 0xF2
}
func (SET_F2) String() string {
	return "SET 6,D"
}

// SET 6,L    code=0xf5
type SET_F5 struct{}

func (SET_F5) Exec(cpu *CPU) {
	panic("TODO SET_F5")
}
func (SET_F5) Code() uint8 {
	return 0xF5
}
func (SET_F5) String() string {
	return "SET 6,L"
}

// RLC A    code=0x07
type RLC_07 struct{}

func (RLC_07) Exec(cpu *CPU) {
	res, flags := rotate(cpu.A, 0, cpu.F, true)
	cpu.A = res
	cpu.F = flags

	cpu.Cycles += 8
}
func (RLC_07) Code() uint8 {
	return 0x7
}
func (RLC_07) String() string {
	return "RLC A"
}

// RES 5,(HL)    code=0xae
type RES_AE struct{}

func (RES_AE) Exec(cpu *CPU) {
	panic("TODO RES_AE")
}
func (RES_AE) Code() uint8 {
	return 0xAE
}
func (RES_AE) String() string {
	return "RES 5,(HL)"
}

// RES 7,D    code=0xba
type RES_BA struct{}

func (RES_BA) Exec(cpu *CPU) {
	panic("TODO RES_BA")
}
func (RES_BA) Code() uint8 {
	return 0xBA
}
func (RES_BA) String() string {
	return "RES 7,D"
}

// RRC D    code=0x0a
type RRC_0A struct{}

func (RRC_0A) Exec(cpu *CPU) {
	res, flags := rotate(cpu.D, 1, cpu.F, true)
	cpu.D = res
	cpu.F = flags

	cpu.Cycles += 8
}
func (RRC_0A) Code() uint8 {
	return 0xA
}
func (RRC_0A) String() string {
	return "RRC D"
}

// RRC (HL)    code=0x0e
type RRC_0E struct{}

func (RRC_0E) Exec(cpu *CPU) {
	res, flags := rotate(cpu.loadU8(cpu.HL()), 1, cpu.F, true)
	cpu.WriteMemory(cpu.HL(), res)
	cpu.F = flags

	cpu.Cycles += 16
}
func (RRC_0E) Code() uint8 {
	return 0xE
}
func (RRC_0E) String() string {
	return "RRC (HL)"
}

// RL L    code=0x15
type RL_15 struct{}

func (RL_15) Exec(cpu *CPU) {
	res, flags := rotate(cpu.L, 0, cpu.F, false)
	cpu.L = res
	cpu.F = flags

	cpu.Cycles += 8
}
func (RL_15) Code() uint8 {
	return 0x15
}
func (RL_15) String() string {
	return "RL L"
}

// RES 1,L    code=0x8d
type RES_8D struct{}

func (RES_8D) Exec(cpu *CPU) {
	panic("TODO RES_8D")
}
func (RES_8D) Code() uint8 {
	return 0x8D
}
func (RES_8D) String() string {
	return "RES 1,L"
}

// RES 1,(HL)    code=0x8e
type RES_8E struct{}

func (RES_8E) Exec(cpu *CPU) {
	panic("TODO RES_8E")
}
func (RES_8E) Code() uint8 {
	return 0x8E
}
func (RES_8E) String() string {
	return "RES 1,(HL)"
}

// RES 5,E    code=0xab
type RES_AB struct{}

func (RES_AB) Exec(cpu *CPU) {
	panic("TODO RES_AB")
}
func (RES_AB) Code() uint8 {
	return 0xAB
}
func (RES_AB) String() string {
	return "RES 5,E"
}

// SET 2,E    code=0xd3
type SET_D3 struct{}

func (SET_D3) Exec(cpu *CPU) {
	panic("TODO SET_D3")
}
func (SET_D3) Code() uint8 {
	return 0xD3
}
func (SET_D3) String() string {
	return "SET 2,E"
}

// RR B    code=0x18
type RR_18 struct{}

func (RR_18) Exec(cpu *CPU) {
	res, flags := rotate(cpu.B, 1, cpu.F, false)
	cpu.B = res
	cpu.F = flags

	cpu.Cycles += 8
}
func (RR_18) Code() uint8 {
	return 0x18
}
func (RR_18) String() string {
	return "RR B"
}

// SLA E    code=0x23
type SLA_23 struct{}

func (SLA_23) Exec(cpu *CPU) {
	panic("TODO SLA_23")
}
func (SLA_23) Code() uint8 {
	return 0x23
}
func (SLA_23) String() string {
	return "SLA E"
}

// RES 3,H    code=0x9c
type RES_9C struct{}

func (RES_9C) Exec(cpu *CPU) {
	panic("TODO RES_9C")
}
func (RES_9C) Code() uint8 {
	return 0x9C
}
func (RES_9C) String() string {
	return "RES 3,H"
}

// RES 4,(HL)    code=0xa6
type RES_A6 struct{}

func (RES_A6) Exec(cpu *CPU) {
	panic("TODO RES_A6")
}
func (RES_A6) Code() uint8 {
	return 0xA6
}
func (RES_A6) String() string {
	return "RES 4,(HL)"
}

// SET 3,D    code=0xda
type SET_DA struct{}

func (SET_DA) Exec(cpu *CPU) {
	panic("TODO SET_DA")
}
func (SET_DA) Code() uint8 {
	return 0xDA
}
func (SET_DA) String() string {
	return "SET 3,D"
}

// BIT 1,A    code=0x4f
type BIT_4F struct{}

func (BIT_4F) Exec(cpu *CPU) {
	value := cpu.A
	var flags Flags
	if (value & (1 << 1)) == 0 {
		flags |= FLAGZ
	}
	flags |= FLAGH
	cpu.F = FlagRegister(flags)

	cpu.Cycles += 8
}
func (BIT_4F) Code() uint8 {
	return 0x4F
}
func (BIT_4F) String() string {
	return "BIT 1,A"
}

// RR D    code=0x1a
type RR_1A struct{}

func (RR_1A) Exec(cpu *CPU) {
	res, flags := rotate(cpu.D, 1, cpu.F, false)
	cpu.D = res
	cpu.F = flags

	cpu.Cycles += 8
}
func (RR_1A) Code() uint8 {
	return 0x1A
}
func (RR_1A) String() string {
	return "RR D"
}

// RES 2,E    code=0x93
type RES_93 struct{}

func (RES_93) Exec(cpu *CPU) {
	panic("TODO RES_93")
}
func (RES_93) Code() uint8 {
	return 0x93
}
func (RES_93) String() string {
	return "RES 2,E"
}

// SET 7,E    code=0xfb
type SET_FB struct{}

func (SET_FB) Exec(cpu *CPU) {
	panic("TODO SET_FB")
}
func (SET_FB) Code() uint8 {
	return 0xFB
}
func (SET_FB) String() string {
	return "SET 7,E"
}

// BIT 1,(HL)    code=0x4e
type BIT_4E struct{}

func (BIT_4E) Exec(cpu *CPU) {
	value := cpu.loadU8(cpu.HL())
	var flags Flags
	if (value & (1 << 1)) == 0 {
		flags |= FLAGZ
	}
	flags |= FLAGH
	cpu.F = FlagRegister(flags)

	cpu.Cycles += 12
}
func (BIT_4E) Code() uint8 {
	return 0x4E
}
func (BIT_4E) String() string {
	return "BIT 1,(HL)"
}

// RES 3,(HL)    code=0x9e
type RES_9E struct{}

func (RES_9E) Exec(cpu *CPU) {
	panic("TODO RES_9E")
}
func (RES_9E) Code() uint8 {
	return 0x9E
}
func (RES_9E) String() string {
	return "RES 3,(HL)"
}

// SET 0,A    code=0xc7
type SET_C7 struct{}

func (SET_C7) Exec(cpu *CPU) {
	panic("TODO SET_C7")
}
func (SET_C7) Code() uint8 {
	return 0xC7
}
func (SET_C7) String() string {
	return "SET 0,A"
}

// SET 5,(HL)    code=0xee
type SET_EE struct{}

func (SET_EE) Exec(cpu *CPU) {
	panic("TODO SET_EE")
}
func (SET_EE) Code() uint8 {
	return 0xEE
}
func (SET_EE) String() string {
	return "SET 5,(HL)"
}

// RL E    code=0x13
type RL_13 struct{}

func (RL_13) Exec(cpu *CPU) {
	res, flags := rotate(cpu.E, 0, cpu.F, false)
	cpu.E = res
	cpu.F = flags

	cpu.Cycles += 8
}
func (RL_13) Code() uint8 {
	return 0x13
}
func (RL_13) String() string {
	return "RL E"
}

// SWAP C    code=0x31
type SWAP_31 struct{}

func (SWAP_31) Exec(cpu *CPU) {
	panic("TODO SWAP_31")
}
func (SWAP_31) Code() uint8 {
	return 0x31
}
func (SWAP_31) String() string {
	return "SWAP C"
}

// BIT 1,H    code=0x4c
type BIT_4C struct{}

func (BIT_4C) Exec(cpu *CPU) {
	value := cpu.H
	var flags Flags
	if (value & (1 << 1)) == 0 {
		flags |= FLAGZ
	}
	flags |= FLAGH
	cpu.F = FlagRegister(flags)

	cpu.Cycles += 8
}
func (BIT_4C) Code() uint8 {
	return 0x4C
}
func (BIT_4C) String() string {
	return "BIT 1,H"
}

// RES 2,D    code=0x92
type RES_92 struct{}

func (RES_92) Exec(cpu *CPU) {
	panic("TODO RES_92")
}
func (RES_92) Code() uint8 {
	return 0x92
}
func (RES_92) String() string {
	return "RES 2,D"
}

// SET 4,H    code=0xe4
type SET_E4 struct{}

func (SET_E4) Exec(cpu *CPU) {
	panic("TODO SET_E4")
}
func (SET_E4) Code() uint8 {
	return 0xE4
}
func (SET_E4) String() string {
	return "SET 4,H"
}

// BIT 2,D    code=0x52
type BIT_52 struct{}

func (BIT_52) Exec(cpu *CPU) {
	value := cpu.D
	var flags Flags
	if (value & (1 << 2)) == 0 {
		flags |= FLAGZ
	}
	flags |= FLAGH
	cpu.F = FlagRegister(flags)

	cpu.Cycles += 8
}
func (BIT_52) Code() uint8 {
	return 0x52
}
func (BIT_52) String() string {
	return "BIT 2,D"
}

// SET 0,(HL)    code=0xc6
type SET_C6 struct{}

func (SET_C6) Exec(cpu *CPU) {
	panic("TODO SET_C6")
}
func (SET_C6) Code() uint8 {
	return 0xC6
}
func (SET_C6) String() string {
	return "SET 0,(HL)"
}

// SET 5,L    code=0xed
type SET_ED struct{}

func (SET_ED) Exec(cpu *CPU) {
	panic("TODO SET_ED")
}
func (SET_ED) Code() uint8 {
	return 0xED
}
func (SET_ED) String() string {
	return "SET 5,L"
}

// BIT 1,B    code=0x48
type BIT_48 struct{}

func (BIT_48) Exec(cpu *CPU) {
	value := cpu.B
	var flags Flags
	if (value & (1 << 1)) == 0 {
		flags |= FLAGZ
	}
	flags |= FLAGH
	cpu.F = FlagRegister(flags)

	cpu.Cycles += 8
}
func (BIT_48) Code() uint8 {
	return 0x48
}
func (BIT_48) String() string {
	return "BIT 1,B"
}

// RES 7,E    code=0xbb
type RES_BB struct{}

func (RES_BB) Exec(cpu *CPU) {
	panic("TODO RES_BB")
}
func (RES_BB) Code() uint8 {
	return 0xBB
}
func (RES_BB) String() string {
	return "RES 7,E"
}

// SRL (HL)    code=0x3e
type SRL_3E struct{}

func (SRL_3E) Exec(cpu *CPU) {
	panic("TODO SRL_3E")
}
func (SRL_3E) Code() uint8 {
	return 0x3E
}
func (SRL_3E) String() string {
	return "SRL (HL)"
}

// BIT 1,L    code=0x4d
type BIT_4D struct{}

func (BIT_4D) Exec(cpu *CPU) {
	value := cpu.L
	var flags Flags
	if (value & (1 << 1)) == 0 {
		flags |= FLAGZ
	}
	flags |= FLAGH
	cpu.F = FlagRegister(flags)

	cpu.Cycles += 8
}
func (BIT_4D) Code() uint8 {
	return 0x4D
}
func (BIT_4D) String() string {
	return "BIT 1,L"
}

// BIT 6,D    code=0x72
type BIT_72 struct{}

func (BIT_72) Exec(cpu *CPU) {
	value := cpu.D
	var flags Flags
	if (value & (1 << 6)) == 0 {
		flags |= FLAGZ
	}
	flags |= FLAGH
	cpu.F = FlagRegister(flags)

	cpu.Cycles += 8
}
func (BIT_72) Code() uint8 {
	return 0x72
}
func (BIT_72) String() string {
	return "BIT 6,D"
}

// RES 7,(HL)    code=0xbe
type RES_BE struct{}

func (RES_BE) Exec(cpu *CPU) {
	panic("TODO RES_BE")
}
func (RES_BE) Code() uint8 {
	return 0xBE
}
func (RES_BE) String() string {
	return "RES 7,(HL)"
}

// BIT 2,C    code=0x51
type BIT_51 struct{}

func (BIT_51) Exec(cpu *CPU) {
	value := cpu.C
	var flags Flags
	if (value & (1 << 2)) == 0 {
		flags |= FLAGZ
	}
	flags |= FLAGH
	cpu.F = FlagRegister(flags)

	cpu.Cycles += 8
}
func (BIT_51) Code() uint8 {
	return 0x51
}
func (BIT_51) String() string {
	return "BIT 2,C"
}

// BIT 6,C    code=0x71
type BIT_71 struct{}

func (BIT_71) Exec(cpu *CPU) {
	value := cpu.C
	var flags Flags
	if (value & (1 << 6)) == 0 {
		flags |= FLAGZ
	}
	flags |= FLAGH
	cpu.F = FlagRegister(flags)

	cpu.Cycles += 8
}
func (BIT_71) Code() uint8 {
	return 0x71
}
func (BIT_71) String() string {
	return "BIT 6,C"
}

// RL A    code=0x17
type RL_17 struct{}

func (RL_17) Exec(cpu *CPU) {
	res, flags := rotate(cpu.A, 0, cpu.F, false)
	cpu.A = res
	cpu.F = flags

	cpu.Cycles += 8
}
func (RL_17) Code() uint8 {
	return 0x17
}
func (RL_17) String() string {
	return "RL A"
}

// BIT 0,D    code=0x42
type BIT_42 struct{}

func (BIT_42) Exec(cpu *CPU) {
	value := cpu.D
	var flags Flags
	if (value & (1 << 0)) == 0 {
		flags |= FLAGZ
	}
	flags |= FLAGH
	cpu.F = FlagRegister(flags)

	cpu.Cycles += 8
}
func (BIT_42) Code() uint8 {
	return 0x42
}
func (BIT_42) String() string {
	return "BIT 0,D"
}

// SET 2,(HL)    code=0xd6
type SET_D6 struct{}

func (SET_D6) Exec(cpu *CPU) {
	panic("TODO SET_D6")
}
func (SET_D6) Code() uint8 {
	return 0xD6
}
func (SET_D6) String() string {
	return "SET 2,(HL)"
}

// RLC C    code=0x01
type RLC_01 struct{}

func (RLC_01) Exec(cpu *CPU) {
	res, flags := rotate(cpu.C, 0, cpu.F, true)
	cpu.C = res
	cpu.F = flags

	cpu.Cycles += 8
}
func (RLC_01) Code() uint8 {
	return 0x1
}
func (RLC_01) String() string {
	return "RLC C"
}

// BIT 1,E    code=0x4b
type BIT_4B struct{}

func (BIT_4B) Exec(cpu *CPU) {
	value := cpu.E
	var flags Flags
	if (value & (1 << 1)) == 0 {
		flags |= FLAGZ
	}
	flags |= FLAGH
	cpu.F = FlagRegister(flags)

	cpu.Cycles += 8
}
func (BIT_4B) Code() uint8 {
	return 0x4B
}
func (BIT_4B) String() string {
	return "BIT 1,E"
}

// SET 3,B    code=0xd8
type SET_D8 struct{}

func (SET_D8) Exec(cpu *CPU) {
	panic("TODO SET_D8")
}
func (SET_D8) Code() uint8 {
	return 0xD8
}
func (SET_D8) String() string {
	return "SET 3,B"
}

// SET 3,C    code=0xd9
type SET_D9 struct{}

func (SET_D9) Exec(cpu *CPU) {
	panic("TODO SET_D9")
}
func (SET_D9) Code() uint8 {
	return 0xD9
}
func (SET_D9) String() string {
	return "SET 3,C"
}

// SLA (HL)    code=0x26
type SLA_26 struct{}

func (SLA_26) Exec(cpu *CPU) {
	panic("TODO SLA_26")
}
func (SLA_26) Code() uint8 {
	return 0x26
}
func (SLA_26) String() string {
	return "SLA (HL)"
}

// SWAP A    code=0x37
type SWAP_37 struct{}

func (SWAP_37) Exec(cpu *CPU) {
	panic("TODO SWAP_37")
}
func (SWAP_37) Code() uint8 {
	return 0x37
}
func (SWAP_37) String() string {
	return "SWAP A"
}

// SET 0,B    code=0xc0
type SET_C0 struct{}

func (SET_C0) Exec(cpu *CPU) {
	panic("TODO SET_C0")
}
func (SET_C0) Code() uint8 {
	return 0xC0
}
func (SET_C0) String() string {
	return "SET 0,B"
}

// RLC E    code=0x03
type RLC_03 struct{}

func (RLC_03) Exec(cpu *CPU) {
	res, flags := rotate(cpu.E, 0, cpu.F, true)
	cpu.E = res
	cpu.F = flags

	cpu.Cycles += 8
}
func (RLC_03) Code() uint8 {
	return 0x3
}
func (RLC_03) String() string {
	return "RLC E"
}

// RES 2,(HL)    code=0x96
type RES_96 struct{}

func (RES_96) Exec(cpu *CPU) {
	panic("TODO RES_96")
}
func (RES_96) Code() uint8 {
	return 0x96
}
func (RES_96) String() string {
	return "RES 2,(HL)"
}

// RES 2,A    code=0x97
type RES_97 struct{}

func (RES_97) Exec(cpu *CPU) {
	panic("TODO RES_97")
}
func (RES_97) Code() uint8 {
	return 0x97
}
func (RES_97) String() string {
	return "RES 2,A"
}

// SRL B    code=0x38
type SRL_38 struct{}

func (SRL_38) Exec(cpu *CPU) {
	panic("TODO SRL_38")
}
func (SRL_38) Code() uint8 {
	return 0x38
}
func (SRL_38) String() string {
	return "SRL B"
}

// BIT 6,H    code=0x74
type BIT_74 struct{}

func (BIT_74) Exec(cpu *CPU) {
	value := cpu.H
	var flags Flags
	if (value & (1 << 6)) == 0 {
		flags |= FLAGZ
	}
	flags |= FLAGH
	cpu.F = FlagRegister(flags)

	cpu.Cycles += 8
}
func (BIT_74) Code() uint8 {
	return 0x74
}
func (BIT_74) String() string {
	return "BIT 6,H"
}

// RES 0,E    code=0x83
type RES_83 struct{}

func (RES_83) Exec(cpu *CPU) {
	panic("TODO RES_83")
}
func (RES_83) Code() uint8 {
	return 0x83
}
func (RES_83) String() string {
	return "RES 0,E"
}

// SET 0,H    code=0xc4
type SET_C4 struct{}

func (SET_C4) Exec(cpu *CPU) {
	panic("TODO SET_C4")
}
func (SET_C4) Code() uint8 {
	return 0xC4
}
func (SET_C4) String() string {
	return "SET 0,H"
}

// SET 0,C    code=0xc1
type SET_C1 struct{}

func (SET_C1) Exec(cpu *CPU) {
	panic("TODO SET_C1")
}
func (SET_C1) Code() uint8 {
	return 0xC1
}
func (SET_C1) String() string {
	return "SET 0,C"
}

// SWAP D    code=0x32
type SWAP_32 struct{}

func (SWAP_32) Exec(cpu *CPU) {
	panic("TODO SWAP_32")
}
func (SWAP_32) Code() uint8 {
	return 0x32
}
func (SWAP_32) String() string {
	return "SWAP D"
}

// BIT 0,(HL)    code=0x46
type BIT_46 struct{}

func (BIT_46) Exec(cpu *CPU) {
	value := cpu.loadU8(cpu.HL())
	var flags Flags
	if (value & (1 << 0)) == 0 {
		flags |= FLAGZ
	}
	flags |= FLAGH
	cpu.F = FlagRegister(flags)

	cpu.Cycles += 12
}
func (BIT_46) Code() uint8 {
	return 0x46
}
func (BIT_46) String() string {
	return "BIT 0,(HL)"
}

// BIT 7,D    code=0x7a
type BIT_7A struct{}

func (BIT_7A) Exec(cpu *CPU) {
	value := cpu.D
	var flags Flags
	if (value & (1 << 7)) == 0 {
		flags |= FLAGZ
	}
	flags |= FLAGH
	cpu.F = FlagRegister(flags)

	cpu.Cycles += 8
}
func (BIT_7A) Code() uint8 {
	return 0x7A
}
func (BIT_7A) String() string {
	return "BIT 7,D"
}

// BIT 7,(HL)    code=0x7e
type BIT_7E struct{}

func (BIT_7E) Exec(cpu *CPU) {
	value := cpu.loadU8(cpu.HL())
	var flags Flags
	if (value & (1 << 7)) == 0 {
		flags |= FLAGZ
	}
	flags |= FLAGH
	cpu.F = FlagRegister(flags)

	cpu.Cycles += 12
}
func (BIT_7E) Code() uint8 {
	return 0x7E
}
func (BIT_7E) String() string {
	return "BIT 7,(HL)"
}

// RES 6,B    code=0xb0
type RES_B0 struct{}

func (RES_B0) Exec(cpu *CPU) {
	panic("TODO RES_B0")
}
func (RES_B0) Code() uint8 {
	return 0xB0
}
func (RES_B0) String() string {
	return "RES 6,B"
}

// SET 3,(HL)    code=0xde
type SET_DE struct{}

func (SET_DE) Exec(cpu *CPU) {
	panic("TODO SET_DE")
}
func (SET_DE) Code() uint8 {
	return 0xDE
}
func (SET_DE) String() string {
	return "SET 3,(HL)"
}

// SRA C    code=0x29
type SRA_29 struct{}

func (SRA_29) Exec(cpu *CPU) {
	panic("TODO SRA_29")
}
func (SRA_29) Code() uint8 {
	return 0x29
}
func (SRA_29) String() string {
	return "SRA C"
}

// BIT 2,A    code=0x57
type BIT_57 struct{}

func (BIT_57) Exec(cpu *CPU) {
	value := cpu.A
	var flags Flags
	if (value & (1 << 2)) == 0 {
		flags |= FLAGZ
	}
	flags |= FLAGH
	cpu.F = FlagRegister(flags)

	cpu.Cycles += 8
}
func (BIT_57) Code() uint8 {
	return 0x57
}
func (BIT_57) String() string {
	return "BIT 2,A"
}

// BIT 3,A    code=0x5f
type BIT_5F struct{}

func (BIT_5F) Exec(cpu *CPU) {
	value := cpu.A
	var flags Flags
	if (value & (1 << 3)) == 0 {
		flags |= FLAGZ
	}
	flags |= FLAGH
	cpu.F = FlagRegister(flags)

	cpu.Cycles += 8
}
func (BIT_5F) Code() uint8 {
	return 0x5F
}
func (BIT_5F) String() string {
	return "BIT 3,A"
}

// RES 5,C    code=0xa9
type RES_A9 struct{}

func (RES_A9) Exec(cpu *CPU) {
	panic("TODO RES_A9")
}
func (RES_A9) Code() uint8 {
	return 0xA9
}
func (RES_A9) String() string {
	return "RES 5,C"
}

// RES 5,D    code=0xaa
type RES_AA struct{}

func (RES_AA) Exec(cpu *CPU) {
	panic("TODO RES_AA")
}
func (RES_AA) Code() uint8 {
	return 0xAA
}
func (RES_AA) String() string {
	return "RES 5,D"
}

// RL (HL)    code=0x16
type RL_16 struct{}

func (RL_16) Exec(cpu *CPU) {
	res, flags := rotate(cpu.loadU8(cpu.HL()), 0, cpu.F, false)
	cpu.WriteMemory(cpu.HL(), res)
	cpu.F = flags

	cpu.Cycles += 16
}
func (RL_16) Code() uint8 {
	return 0x16
}
func (RL_16) String() string {
	return "RL (HL)"
}

// SRA A    code=0x2f
type SRA_2F struct{}

func (SRA_2F) Exec(cpu *CPU) {
	panic("TODO SRA_2F")
}
func (SRA_2F) Code() uint8 {
	return 0x2F
}
func (SRA_2F) String() string {
	return "SRA A"
}

// SET 2,L    code=0xd5
type SET_D5 struct{}

func (SET_D5) Exec(cpu *CPU) {
	panic("TODO SET_D5")
}
func (SET_D5) Code() uint8 {
	return 0xD5
}
func (SET_D5) String() string {
	return "SET 2,L"
}

// BIT 3,B    code=0x58
type BIT_58 struct{}

func (BIT_58) Exec(cpu *CPU) {
	value := cpu.B
	var flags Flags
	if (value & (1 << 3)) == 0 {
		flags |= FLAGZ
	}
	flags |= FLAGH
	cpu.F = FlagRegister(flags)

	cpu.Cycles += 8
}
func (BIT_58) Code() uint8 {
	return 0x58
}
func (BIT_58) String() string {
	return "BIT 3,B"
}

// RES 0,D    code=0x82
type RES_82 struct{}

func (RES_82) Exec(cpu *CPU) {
	panic("TODO RES_82")
}
func (RES_82) Code() uint8 {
	return 0x82
}
func (RES_82) String() string {
	return "RES 0,D"
}

// RES 6,A    code=0xb7
type RES_B7 struct{}

func (RES_B7) Exec(cpu *CPU) {
	panic("TODO RES_B7")
}
func (RES_B7) Code() uint8 {
	return 0xB7
}
func (RES_B7) String() string {
	return "RES 6,A"
}

// SET 6,A    code=0xf7
type SET_F7 struct{}

func (SET_F7) Exec(cpu *CPU) {
	panic("TODO SET_F7")
}
func (SET_F7) Code() uint8 {
	return 0xF7
}
func (SET_F7) String() string {
	return "SET 6,A"
}

// RES 4,B    code=0xa0
type RES_A0 struct{}

func (RES_A0) Exec(cpu *CPU) {
	panic("TODO RES_A0")
}
func (RES_A0) Code() uint8 {
	return 0xA0
}
func (RES_A0) String() string {
	return "RES 4,B"
}

// RES 4,D    code=0xa2
type RES_A2 struct{}

func (RES_A2) Exec(cpu *CPU) {
	panic("TODO RES_A2")
}
func (RES_A2) Code() uint8 {
	return 0xA2
}
func (RES_A2) String() string {
	return "RES 4,D"
}

var ops = map[uint8]Instruction{
	0xd1: POP_D1{},
	0xda: JP_DA{},
	0x2:  LD_02{},
	0x37: SCF_37{},
	0xf7: RST_F7{},
	0xe6: AND_E6{},
	0xeb: ILLEGAL_EB_EB{},
	0x3:  INC_03{},
	0x41: LD_41{},
	0x48: LD_48{},
	0x7d: LD_7D{},
	0xa9: XOR_A9{},
	0xbb: CP_BB{},
	0x88: ADC_88{},
	0x1a: LD_1A{},
	0x92: SUB_92{},
	0xdc: CALL_DC{},
	0xa7: AND_A7{},
	0xdf: RST_DF{},
	0x70: LD_70{},
	0xae: XOR_AE{},
	0xc8: RET_C8{},
	0x7:  RLCA_07{},
	0x17: RLA_17{},
	0x40: LD_40{},
	0xfe: CP_FE{},
	0x86: ADD_86{},
	0x96: SUB_96{},
	0xa0: AND_A0{},
	0xd4: CALL_D4{},
	0xd9: RETI_D9{},
	0x9:  ADD_09{},
	0xc:  INC_0C{},
	0x3c: INC_3C{},
	0x64: LD_64{},
	0xde: SBC_DE{},
	0xfb: EI_FB{},
	0xfc: ILLEGAL_FC_FC{},
	0x52: LD_52{},
	0xc4: CALL_C4{},
	0xd8: RET_D8{},
	0xdd: ILLEGAL_DD_DD{},
	0x1:  LD_01{},
	0x24: INC_24{},
	0x90: SUB_90{},
	0xc1: POP_C1{},
	0xcf: RST_CF{},
	0xe8: ADD_E8{},
	0x8f: ADC_8F{},
	0xf0: LDH_F0{},
	0x35: DEC_35{},
	0x72: LD_72{},
	0x9a: SBC_9A{},
	0xa6: AND_A6{},
	0xed: ILLEGAL_ED_ED{},
	0xf4: ILLEGAL_F4_F4{},
	0x8:  LD_08{},
	0x13: INC_13{},
	0xa:  LD_0A{},
	0xe:  LD_0E{},
	0x27: DAA_27{},
	0x2b: DEC_2B{},
	0x3a: LD_3A{},
	0x4d: LD_4D{},
	0xfd: ILLEGAL_FD_FD{},
	0x21: LD_21{},
	0x26: LD_26{},
	0x87: ADD_87{},
	0xb9: CP_B9{},
	0xbe: CP_BE{},
	0xe7: RST_E7{},
	0x5a: LD_5A{},
	0xe3: ILLEGAL_E3_E3{},
	0x20: JR_20{},
	0x4c: LD_4C{},
	0x32: LD_32{},
	0x34: INC_34{},
	0xe0: LDH_E0{},
	0xf2: LDH_F2{},
	0x30: JR_30{},
	0x63: LD_63{},
	0x8c: ADC_8C{},
	0xa1: AND_A1{},
	0x6b: LD_6B{},
	0x33: INC_33{},
	0x51: LD_51{},
	0x76: HALT_76{},
	0x81: ADD_81{},
	0x82: ADD_82{},
	0x42: LD_42{},
	0x4e: LD_4E{},
	0x74: LD_74{},
	0x8a: ADC_8A{},
	0x4:  INC_04{},
	0x77: LD_77{},
	0xab: XOR_AB{},
	0xb8: CP_B8{},
	0xd0: RET_D0{},
	0xb1: OR_B1{},
	0x6c: LD_6C{},
	0x2d: DEC_2D{},
	0x50: LD_50{},
	0x7f: LD_7F{},
	0x95: SUB_95{},
	0x9d: SBC_9D{},
	0x9f: SBC_9F{},
	0xaa: XOR_AA{},
	0x1e: LD_1E{},
	0x38: JR_38{},
	0x60: LD_60{},
	0x85: ADD_85{},
	0xc5: PUSH_C5{},
	0xd5: PUSH_D5{},
	0xff: RST_FF{},
	0x6:  LD_06{},
	0x1d: DEC_1D{},
	0x3b: DEC_3B{},
	0x6d: LD_6D{},
	0x9c: SBC_9C{},
	0xd2: JP_D2{},
	0x3f: CCF_3F{},
	0x23: INC_23{},
	0x31: LD_31{},
	0x39: ADD_39{},
	0x4a: LD_4A{},
	0x55: LD_55{},
	0x5b: LD_5B{},
	0x91: SUB_91{},
	0x94: SUB_94{},
	0x3d: DEC_3D{},
	0x57: LD_57{},
	0x6a: LD_6A{},
	0xa4: AND_A4{},
	0xbf: CP_BF{},
	0xcd: CALL_CD{},
	0xf3: DI_F3{},
	0x47: LD_47{},
	0x9b: SBC_9B{},
	0xd7: RST_D7{},
	0xf6: OR_F6{},
	0xfa: LD_FA{},
	0x79: LD_79{},
	0x98: SBC_98{},
	0x7b: LD_7B{},
	0x89: ADC_89{},
	0xf8: LD_F8{},
	0x14: INC_14{},
	0x44: LD_44{},
	0x7e: LD_7E{},
	0xcb: PREFIX_CB{},
	0x0:  NOP_00{},
	0x54: LD_54{},
	0x59: LD_59{},
	0xc2: JP_C2{},
	0x36: LD_36{},
	0x75: LD_75{},
	0xac: XOR_AC{},
	0xd6: SUB_D6{},
	0x2c: INC_2C{},
	0xad: XOR_AD{},
	0xe9: JP_E9{},
	0x28: JR_28{},
	0x58: LD_58{},
	0x8d: ADC_8D{},
	0xd3: ILLEGAL_D3_D3{},
	0x5d: LD_5D{},
	0x67: LD_67{},
	0x80: ADD_80{},
	0xe2: LDH_E2{},
	0xef: RST_EF{},
	0x4f: LD_4F{},
	0xf1: POP_F1{},
	0x15: DEC_15{},
	0x65: LD_65{},
	0x68: LD_68{},
	0x5c: LD_5C{},
	0x84: ADD_84{},
	0x99: SBC_99{},
	0xb7: OR_B7{},
	0xc9: RET_C9{},
	0xe5: PUSH_E5{},
	0xf9: LD_F9{},
	0x1c: INC_1C{},
	0x22: LD_22{},
	0xb6: OR_B6{},
	0xec: ILLEGAL_EC_EC{},
	0xf5: PUSH_F5{},
	0x83: ADD_83{},
	0xe1: POP_E1{},
	0x53: LD_53{},
	0x93: SUB_93{},
	0xb2: OR_B2{},
	0x16: LD_16{},
	0x1b: DEC_1B{},
	0x29: ADD_29{},
	0x2a: LD_2A{},
	0x49: LD_49{},
	0x61: LD_61{},
	0x8b: ADC_8B{},
	0x9e: SBC_9E{},
	0xaf: XOR_AF{},
	0xc7: RST_C7{},
	0x12: LD_12{},
	0x25: DEC_25{},
	0xa8: XOR_A8{},
	0xdb: ILLEGAL_DB_DB{},
	0x45: LD_45{},
	0x71: LD_71{},
	0x7a: LD_7A{},
	0xc6: ADD_C6{},
	0xba: CP_BA{},
	0xca: JP_CA{},
	0xe4: ILLEGAL_E4_E4{},
	0x3e: LD_3E{},
	0x43: LD_43{},
	0x5f: LD_5F{},
	0x69: LD_69{},
	0x6f: LD_6F{},
	0x97: SUB_97{},
	0xbc: CP_BC{},
	0xc3: JP_C3{},
	0x1f: RRA_1F{},
	0x4b: LD_4B{},
	0xb:  DEC_0B{},
	0x10: STOP_10{},
	0x11: LD_11{},
	0x46: LD_46{},
	0xea: LD_EA{},
	0xa2: AND_A2{},
	0xb5: OR_B5{},
	0x18: JR_18{},
	0x19: ADD_19{},
	0x2f: CPL_2F{},
	0x5e: LD_5E{},
	0xcc: CALL_CC{},
	0xee: XOR_EE{},
	0x62: LD_62{},
	0xa3: AND_A3{},
	0xbd: CP_BD{},
	0x5:  DEC_05{},
	0xb3: OR_B3{},
	0xd:  DEC_0D{},
	0x56: LD_56{},
	0xb0: OR_B0{},
	0xce: ADC_CE{},
	0x2e: LD_2E{},
	0x6e: LD_6E{},
	0x73: LD_73{},
	0x78: LD_78{},
	0x7c: LD_7C{},
	0xa5: AND_A5{},
	0xb4: OR_B4{},
	0xc0: RET_C0{},
	0x66: LD_66{},
	0x8e: ADC_8E{},
}
var extOps = map[uint8]Instruction{
	0xb3: RES_B3{},
	0xce: SET_CE{},
	0xfe: SET_FE{},
	0x34: SWAP_34{},
	0x8c: RES_8C{},
	0x5a: BIT_5A{},
	0x9d: RES_9D{},
	0xdc: SET_DC{},
	0x1f: RR_1F{},
	0x56: BIT_56{},
	0xb5: RES_B5{},
	0xe1: SET_E1{},
	0xfa: SET_FA{},
	0x2e: SRA_2E{},
	0x81: RES_81{},
	0xb9: RES_B9{},
	0xdd: SET_DD{},
	0x11: RL_11{},
	0x2b: SRA_2B{},
	0x3f: SRL_3F{},
	0x64: BIT_64{},
	0x8b: RES_8B{},
	0xa7: RES_A7{},
	0x6:  RLC_06{},
	0x41: BIT_41{},
	0x53: BIT_53{},
	0x67: BIT_67{},
	0x77: BIT_77{},
	0xb4: RES_B4{},
	0xe8: SET_E8{},
	0xf8: SET_F8{},
	0x12: RL_12{},
	0x50: BIT_50{},
	0x88: RES_88{},
	0xbf: RES_BF{},
	0xc9: SET_C9{},
	0x94: RES_94{},
	0xcf: SET_CF{},
	0x25: SLA_25{},
	0x3d: SRL_3D{},
	0x55: BIT_55{},
	0x65: BIT_65{},
	0x6c: BIT_6C{},
	0x7c: BIT_7C{},
	0xa4: RES_A4{},
	0xa5: RES_A5{},
	0xc:  RRC_0C{},
	0x1c: RR_1C{},
	0x5c: BIT_5C{},
	0x80: RES_80{},
	0xcc: SET_CC{},
	0xc3: SET_C3{},
	0xd1: SET_D1{},
	0x6d: BIT_6D{},
	0xff: SET_FF{},
	0xc2: SET_C2{},
	0xbd: RES_BD{},
	0x2a: SRA_2A{},
	0x3b: SRL_3B{},
	0x4a: BIT_4A{},
	0xb6: RES_B6{},
	0xdb: SET_DB{},
	0x9:  RRC_09{},
	0x39: SRL_39{},
	0x5e: BIT_5E{},
	0x62: BIT_62{},
	0xc5: SET_C5{},
	0x22: SLA_22{},
	0x60: BIT_60{},
	0x7f: BIT_7F{},
	0x85: RES_85{},
	0xcd: SET_CD{},
	0xeb: SET_EB{},
	0x1b: RR_1B{},
	0x3c: SRL_3C{},
	0xca: SET_CA{},
	0xd2: SET_D2{},
	0x6a: BIT_6A{},
	0x1d: RR_1D{},
	0x59: BIT_59{},
	0x5b: BIT_5B{},
	0x75: BIT_75{},
	0x87: RES_87{},
	0x90: RES_90{},
	0x5:  RLC_05{},
	0x9f: RES_9F{},
	0xad: RES_AD{},
	0x9b: RES_9B{},
	0xe5: SET_E5{},
	0xe9: SET_E9{},
	0xf9: SET_F9{},
	0x91: RES_91{},
	0xb2: RES_B2{},
	0xfc: SET_FC{},
	0xd:  RRC_0D{},
	0x40: BIT_40{},
	0x44: BIT_44{},
	0xac: RES_AC{},
	0x8:  RRC_08{},
	0x66: BIT_66{},
	0x6f: BIT_6F{},
	0xe0: SET_E0{},
	0xb:  RRC_0B{},
	0x63: BIT_63{},
	0x54: BIT_54{},
	0x84: RES_84{},
	0x8a: RES_8A{},
	0xaf: RES_AF{},
	0xf1: SET_F1{},
	0x2:  RLC_02{},
	0x2d: SRA_2D{},
	0x6e: BIT_6E{},
	0xa3: RES_A3{},
	0xdf: SET_DF{},
	0xe3: SET_E3{},
	0x9a: RES_9A{},
	0x0:  RLC_00{},
	0x95: RES_95{},
	0x20: SLA_20{},
	0x2c: SRA_2C{},
	0x86: RES_86{},
	0xbc: RES_BC{},
	0x5d: BIT_5D{},
	0x73: BIT_73{},
	0x7b: BIT_7B{},
	0x14: RL_14{},
	0x47: BIT_47{},
	0x7d: BIT_7D{},
	0x4:  RLC_04{},
	0x27: SLA_27{},
	0x35: SWAP_35{},
	0x78: BIT_78{},
	0x89: RES_89{},
	0xa8: RES_A8{},
	0xf:  RRC_0F{},
	0x10: RL_10{},
	0x1e: RR_1E{},
	0x24: SLA_24{},
	0x28: SRA_28{},
	0xe2: SET_E2{},
	0xe6: SET_E6{},
	0xf3: SET_F3{},
	0x69: BIT_69{},
	0xd7: SET_D7{},
	0x21: SLA_21{},
	0x45: BIT_45{},
	0xc8: SET_C8{},
	0xd0: SET_D0{},
	0x36: SWAP_36{},
	0x79: BIT_79{},
	0x8f: RES_8F{},
	0x98: RES_98{},
	0xf6: SET_F6{},
	0x43: BIT_43{},
	0x49: BIT_49{},
	0x6b: BIT_6B{},
	0xa1: RES_A1{},
	0xb1: RES_B1{},
	0xcb: SET_CB{},
	0xef: SET_EF{},
	0x19: RR_19{},
	0x70: BIT_70{},
	0xb8: RES_B8{},
	0xe7: SET_E7{},
	0xf0: SET_F0{},
	0xf4: SET_F4{},
	0x30: SWAP_30{},
	0x33: SWAP_33{},
	0x3a: SRL_3A{},
	0x76: BIT_76{},
	0xd4: SET_D4{},
	0xfd: SET_FD{},
	0x61: BIT_61{},
	0x68: BIT_68{},
	0x99: RES_99{},
	0xea: SET_EA{},
	0xec: SET_EC{},
	0xf2: SET_F2{},
	0xf5: SET_F5{},
	0x7:  RLC_07{},
	0xae: RES_AE{},
	0xba: RES_BA{},
	0xa:  RRC_0A{},
	0xe:  RRC_0E{},
	0x15: RL_15{},
	0x8d: RES_8D{},
	0x8e: RES_8E{},
	0xab: RES_AB{},
	0xd3: SET_D3{},
	0x18: RR_18{},
	0x23: SLA_23{},
	0x9c: RES_9C{},
	0xa6: RES_A6{},
	0xda: SET_DA{},
	0x4f: BIT_4F{},
	0x1a: RR_1A{},
	0x93: RES_93{},
	0xfb: SET_FB{},
	0x4e: BIT_4E{},
	0x9e: RES_9E{},
	0xc7: SET_C7{},
	0xee: SET_EE{},
	0x13: RL_13{},
	0x31: SWAP_31{},
	0x4c: BIT_4C{},
	0x92: RES_92{},
	0xe4: SET_E4{},
	0x52: BIT_52{},
	0xc6: SET_C6{},
	0xed: SET_ED{},
	0x48: BIT_48{},
	0xbb: RES_BB{},
	0x3e: SRL_3E{},
	0x4d: BIT_4D{},
	0x72: BIT_72{},
	0xbe: RES_BE{},
	0x51: BIT_51{},
	0x71: BIT_71{},
	0x17: RL_17{},
	0x42: BIT_42{},
	0xd6: SET_D6{},
	0x1:  RLC_01{},
	0x4b: BIT_4B{},
	0xd8: SET_D8{},
	0xd9: SET_D9{},
	0x26: SLA_26{},
	0x37: SWAP_37{},
	0xc0: SET_C0{},
	0x3:  RLC_03{},
	0x96: RES_96{},
	0x97: RES_97{},
	0x38: SRL_38{},
	0x74: BIT_74{},
	0x83: RES_83{},
	0xc4: SET_C4{},
	0xc1: SET_C1{},
	0x32: SWAP_32{},
	0x46: BIT_46{},
	0x7a: BIT_7A{},
	0x7e: BIT_7E{},
	0xb0: RES_B0{},
	0xde: SET_DE{},
	0x29: SRA_29{},
	0x57: BIT_57{},
	0x5f: BIT_5F{},
	0xa9: RES_A9{},
	0xaa: RES_AA{},
	0x16: RL_16{},
	0x2f: SRA_2F{},
	0xd5: SET_D5{},
	0x58: BIT_58{},
	0x82: RES_82{},
	0xb7: RES_B7{},
	0xf7: SET_F7{},
	0xa0: RES_A0{},
	0xa2: RES_A2{},
}

// returns code given a string. Useful during testing
func code(s string) uint8 {
	switch s {
	case "POP DE":
		return 0xD1
	case "JP C,a16":
		return 0xDA
	case "LD (BC),A":
		return 0x2
	case "SCF":
		return 0x37
	case "RST $30":
		return 0xF7
	case "AND A,n8":
		return 0xE6
	case "ILLEGAL_EB":
		return 0xEB
	case "INC BC":
		return 0x3
	case "LD B,C":
		return 0x41
	case "LD C,B":
		return 0x48
	case "LD A,L":
		return 0x7D
	case "XOR A,C":
		return 0xA9
	case "CP A,E":
		return 0xBB
	case "ADC A,B":
		return 0x88
	case "LD A,(DE)":
		return 0x1A
	case "SUB A,D":
		return 0x92
	case "CALL C,a16":
		return 0xDC
	case "AND A,A":
		return 0xA7
	case "RST $18":
		return 0xDF
	case "LD (HL),B":
		return 0x70
	case "XOR A,(HL)":
		return 0xAE
	case "RET Z":
		return 0xC8
	case "RLCA":
		return 0x7
	case "RLA":
		return 0x17
	case "LD B,B":
		return 0x40
	case "CP A,n8":
		return 0xFE
	case "ADD A,(HL)":
		return 0x86
	case "SUB A,(HL)":
		return 0x96
	case "AND A,B":
		return 0xA0
	case "CALL NC,a16":
		return 0xD4
	case "RETI":
		return 0xD9
	case "ADD HL,BC":
		return 0x9
	case "INC C":
		return 0xC
	case "INC A":
		return 0x3C
	case "LD H,H":
		return 0x64
	case "SBC A,n8":
		return 0xDE
	case "EI":
		return 0xFB
	case "ILLEGAL_FC":
		return 0xFC
	case "LD D,D":
		return 0x52
	case "CALL NZ,a16":
		return 0xC4
	case "RET C":
		return 0xD8
	case "ILLEGAL_DD":
		return 0xDD
	case "LD BC,n16":
		return 0x1
	case "INC H":
		return 0x24
	case "SUB A,B":
		return 0x90
	case "POP BC":
		return 0xC1
	case "RST $08":
		return 0xCF
	case "ADD SP,e8":
		return 0xE8
	case "ADC A,A":
		return 0x8F
	case "LDH A,(a8)":
		return 0xF0
	case "DEC (HL)":
		return 0x35
	case "LD (HL),D":
		return 0x72
	case "SBC A,D":
		return 0x9A
	case "AND A,(HL)":
		return 0xA6
	case "ILLEGAL_ED":
		return 0xED
	case "ILLEGAL_F4":
		return 0xF4
	case "LD (a16),SP":
		return 0x8
	case "INC DE":
		return 0x13
	case "LD A,(BC)":
		return 0xA
	case "LD C,n8":
		return 0xE
	case "DAA":
		return 0x27
	case "DEC HL":
		return 0x2B
	case "LD A,(HL-)":
		return 0x3A
	case "LD C,L":
		return 0x4D
	case "ILLEGAL_FD":
		return 0xFD
	case "LD HL,n16":
		return 0x21
	case "LD H,n8":
		return 0x26
	case "ADD A,A":
		return 0x87
	case "CP A,C":
		return 0xB9
	case "CP A,(HL)":
		return 0xBE
	case "RST $20":
		return 0xE7
	case "LD E,D":
		return 0x5A
	case "ILLEGAL_E3":
		return 0xE3
	case "JR NZ,e8":
		return 0x20
	case "LD C,H":
		return 0x4C
	case "LD (HL-),A":
		return 0x32
	case "INC (HL)":
		return 0x34
	case "LDH (a8),A":
		return 0xE0
	case "LDH A,(C)":
		return 0xF2
	case "JR NC,e8":
		return 0x30
	case "LD H,E":
		return 0x63
	case "ADC A,H":
		return 0x8C
	case "AND A,C":
		return 0xA1
	case "LD L,E":
		return 0x6B
	case "INC SP":
		return 0x33
	case "LD D,C":
		return 0x51
	case "HALT":
		return 0x76
	case "ADD A,C":
		return 0x81
	case "ADD A,D":
		return 0x82
	case "LD B,D":
		return 0x42
	case "LD C,(HL)":
		return 0x4E
	case "LD (HL),H":
		return 0x74
	case "ADC A,D":
		return 0x8A
	case "INC B":
		return 0x4
	case "LD (HL),A":
		return 0x77
	case "XOR A,E":
		return 0xAB
	case "CP A,B":
		return 0xB8
	case "RET NC":
		return 0xD0
	case "OR A,C":
		return 0xB1
	case "LD L,H":
		return 0x6C
	case "DEC L":
		return 0x2D
	case "LD D,B":
		return 0x50
	case "LD A,A":
		return 0x7F
	case "SUB A,L":
		return 0x95
	case "SBC A,L":
		return 0x9D
	case "SBC A,A":
		return 0x9F
	case "XOR A,D":
		return 0xAA
	case "LD E,n8":
		return 0x1E
	case "JR C,e8":
		return 0x38
	case "LD H,B":
		return 0x60
	case "ADD A,L":
		return 0x85
	case "PUSH BC":
		return 0xC5
	case "PUSH DE":
		return 0xD5
	case "RST $38":
		return 0xFF
	case "LD B,n8":
		return 0x6
	case "DEC E":
		return 0x1D
	case "DEC SP":
		return 0x3B
	case "LD L,L":
		return 0x6D
	case "SBC A,H":
		return 0x9C
	case "JP NC,a16":
		return 0xD2
	case "CCF":
		return 0x3F
	case "INC HL":
		return 0x23
	case "LD SP,n16":
		return 0x31
	case "ADD HL,SP":
		return 0x39
	case "LD C,D":
		return 0x4A
	case "LD D,L":
		return 0x55
	case "LD E,E":
		return 0x5B
	case "SUB A,C":
		return 0x91
	case "SUB A,H":
		return 0x94
	case "DEC A":
		return 0x3D
	case "LD D,A":
		return 0x57
	case "LD L,D":
		return 0x6A
	case "AND A,H":
		return 0xA4
	case "CP A,A":
		return 0xBF
	case "CALL a16":
		return 0xCD
	case "DI":
		return 0xF3
	case "LD B,A":
		return 0x47
	case "SBC A,E":
		return 0x9B
	case "RST $10":
		return 0xD7
	case "OR A,n8":
		return 0xF6
	case "LD A,(a16)":
		return 0xFA
	case "LD A,C":
		return 0x79
	case "SBC A,B":
		return 0x98
	case "LD A,E":
		return 0x7B
	case "ADC A,C":
		return 0x89
	case "LD HL,SP+,e8":
		return 0xF8
	case "INC D":
		return 0x14
	case "LD B,H":
		return 0x44
	case "LD A,(HL)":
		return 0x7E
	case "PREFIX":
		return 0xCB
	case "NOP":
		return 0x0
	case "LD D,H":
		return 0x54
	case "LD E,C":
		return 0x59
	case "JP NZ,a16":
		return 0xC2
	case "LD (HL),n8":
		return 0x36
	case "LD (HL),L":
		return 0x75
	case "XOR A,H":
		return 0xAC
	case "SUB A,n8":
		return 0xD6
	case "INC L":
		return 0x2C
	case "XOR A,L":
		return 0xAD
	case "JP HL":
		return 0xE9
	case "JR Z,e8":
		return 0x28
	case "LD E,B":
		return 0x58
	case "ADC A,L":
		return 0x8D
	case "ILLEGAL_D3":
		return 0xD3
	case "LD E,L":
		return 0x5D
	case "LD H,A":
		return 0x67
	case "ADD A,B":
		return 0x80
	case "LDH (C),A":
		return 0xE2
	case "RST $28":
		return 0xEF
	case "LD C,A":
		return 0x4F
	case "POP AF":
		return 0xF1
	case "DEC D":
		return 0x15
	case "LD H,L":
		return 0x65
	case "LD L,B":
		return 0x68
	case "LD E,H":
		return 0x5C
	case "ADD A,H":
		return 0x84
	case "SBC A,C":
		return 0x99
	case "OR A,A":
		return 0xB7
	case "RET":
		return 0xC9
	case "PUSH HL":
		return 0xE5
	case "LD SP,HL":
		return 0xF9
	case "INC E":
		return 0x1C
	case "LD (HL+),A":
		return 0x22
	case "OR A,(HL)":
		return 0xB6
	case "ILLEGAL_EC":
		return 0xEC
	case "PUSH AF":
		return 0xF5
	case "ADD A,E":
		return 0x83
	case "POP HL":
		return 0xE1
	case "LD D,E":
		return 0x53
	case "SUB A,E":
		return 0x93
	case "OR A,D":
		return 0xB2
	case "LD D,n8":
		return 0x16
	case "DEC DE":
		return 0x1B
	case "ADD HL,HL":
		return 0x29
	case "LD A,(HL+)":
		return 0x2A
	case "LD C,C":
		return 0x49
	case "LD H,C":
		return 0x61
	case "ADC A,E":
		return 0x8B
	case "SBC A,(HL)":
		return 0x9E
	case "XOR A,A":
		return 0xAF
	case "RST $00":
		return 0xC7
	case "LD (DE),A":
		return 0x12
	case "DEC H":
		return 0x25
	case "XOR A,B":
		return 0xA8
	case "ILLEGAL_DB":
		return 0xDB
	case "LD B,L":
		return 0x45
	case "LD (HL),C":
		return 0x71
	case "LD A,D":
		return 0x7A
	case "ADD A,n8":
		return 0xC6
	case "CP A,D":
		return 0xBA
	case "JP Z,a16":
		return 0xCA
	case "ILLEGAL_E4":
		return 0xE4
	case "LD A,n8":
		return 0x3E
	case "LD B,E":
		return 0x43
	case "LD E,A":
		return 0x5F
	case "LD L,C":
		return 0x69
	case "LD L,A":
		return 0x6F
	case "SUB A,A":
		return 0x97
	case "CP A,H":
		return 0xBC
	case "JP a16":
		return 0xC3
	case "RRA":
		return 0x1F
	case "LD C,E":
		return 0x4B
	case "DEC BC":
		return 0xB
	case "STOP":
		return 0x10
	case "LD DE,n16":
		return 0x11
	case "LD B,(HL)":
		return 0x46
	case "LD (a16),A":
		return 0xEA
	case "AND A,D":
		return 0xA2
	case "OR A,L":
		return 0xB5
	case "JR e8":
		return 0x18
	case "ADD HL,DE":
		return 0x19
	case "CPL":
		return 0x2F
	case "LD E,(HL)":
		return 0x5E
	case "CALL Z,a16":
		return 0xCC
	case "XOR A,n8":
		return 0xEE
	case "LD H,D":
		return 0x62
	case "AND A,E":
		return 0xA3
	case "CP A,L":
		return 0xBD
	case "DEC B":
		return 0x5
	case "OR A,E":
		return 0xB3
	case "DEC C":
		return 0xD
	case "LD D,(HL)":
		return 0x56
	case "OR A,B":
		return 0xB0
	case "ADC A,n8":
		return 0xCE
	case "LD L,n8":
		return 0x2E
	case "LD L,(HL)":
		return 0x6E
	case "LD (HL),E":
		return 0x73
	case "LD A,B":
		return 0x78
	case "LD A,H":
		return 0x7C
	case "AND A,L":
		return 0xA5
	case "OR A,H":
		return 0xB4
	case "RET NZ":
		return 0xC0
	case "LD H,(HL)":
		return 0x66
	case "ADC A,(HL)":
		return 0x8E
	case "RES 6,E":
		return 0xB3
	case "SET 1,(HL)":
		return 0xCE
	case "SET 7,(HL)":
		return 0xFE
	case "SWAP H":
		return 0x34
	case "RES 1,H":
		return 0x8C
	case "BIT 3,D":
		return 0x5A
	case "RES 3,L":
		return 0x9D
	case "SET 3,H":
		return 0xDC
	case "RR A":
		return 0x1F
	case "BIT 2,(HL)":
		return 0x56
	case "RES 6,L":
		return 0xB5
	case "SET 4,C":
		return 0xE1
	case "SET 7,D":
		return 0xFA
	case "SRA (HL)":
		return 0x2E
	case "RES 0,C":
		return 0x81
	case "RES 7,C":
		return 0xB9
	case "SET 3,L":
		return 0xDD
	case "RL C":
		return 0x11
	case "SRA E":
		return 0x2B
	case "SRL A":
		return 0x3F
	case "BIT 4,H":
		return 0x64
	case "RES 1,E":
		return 0x8B
	case "RES 4,A":
		return 0xA7
	case "RLC (HL)":
		return 0x6
	case "BIT 0,C":
		return 0x41
	case "BIT 2,E":
		return 0x53
	case "BIT 4,A":
		return 0x67
	case "BIT 6,A":
		return 0x77
	case "RES 6,H":
		return 0xB4
	case "SET 5,B":
		return 0xE8
	case "SET 7,B":
		return 0xF8
	case "RL D":
		return 0x12
	case "BIT 2,B":
		return 0x50
	case "RES 1,B":
		return 0x88
	case "RES 7,A":
		return 0xBF
	case "SET 1,C":
		return 0xC9
	case "RES 2,H":
		return 0x94
	case "SET 1,A":
		return 0xCF
	case "SLA L":
		return 0x25
	case "SRL L":
		return 0x3D
	case "BIT 2,L":
		return 0x55
	case "BIT 4,L":
		return 0x65
	case "BIT 5,H":
		return 0x6C
	case "BIT 7,H":
		return 0x7C
	case "RES 4,H":
		return 0xA4
	case "RES 4,L":
		return 0xA5
	case "RRC H":
		return 0xC
	case "RR H":
		return 0x1C
	case "BIT 3,H":
		return 0x5C
	case "RES 0,B":
		return 0x80
	case "SET 1,H":
		return 0xCC
	case "SET 0,E":
		return 0xC3
	case "SET 2,C":
		return 0xD1
	case "BIT 5,L":
		return 0x6D
	case "SET 7,A":
		return 0xFF
	case "SET 0,D":
		return 0xC2
	case "RES 7,L":
		return 0xBD
	case "SRA D":
		return 0x2A
	case "SRL E":
		return 0x3B
	case "BIT 1,D":
		return 0x4A
	case "RES 6,(HL)":
		return 0xB6
	case "SET 3,E":
		return 0xDB
	case "RRC C":
		return 0x9
	case "SRL C":
		return 0x39
	case "BIT 3,(HL)":
		return 0x5E
	case "BIT 4,D":
		return 0x62
	case "SET 0,L":
		return 0xC5
	case "SLA D":
		return 0x22
	case "BIT 4,B":
		return 0x60
	case "BIT 7,A":
		return 0x7F
	case "RES 0,L":
		return 0x85
	case "SET 1,L":
		return 0xCD
	case "SET 5,E":
		return 0xEB
	case "RR E":
		return 0x1B
	case "SRL H":
		return 0x3C
	case "SET 1,D":
		return 0xCA
	case "SET 2,D":
		return 0xD2
	case "BIT 5,D":
		return 0x6A
	case "RR L":
		return 0x1D
	case "BIT 3,C":
		return 0x59
	case "BIT 3,E":
		return 0x5B
	case "BIT 6,L":
		return 0x75
	case "RES 0,A":
		return 0x87
	case "RES 2,B":
		return 0x90
	case "RLC L":
		return 0x5
	case "RES 3,A":
		return 0x9F
	case "RES 5,L":
		return 0xAD
	case "RES 3,E":
		return 0x9B
	case "SET 4,L":
		return 0xE5
	case "SET 5,C":
		return 0xE9
	case "SET 7,C":
		return 0xF9
	case "RES 2,C":
		return 0x91
	case "RES 6,D":
		return 0xB2
	case "SET 7,H":
		return 0xFC
	case "RRC L":
		return 0xD
	case "BIT 0,B":
		return 0x40
	case "BIT 0,H":
		return 0x44
	case "RES 5,H":
		return 0xAC
	case "RRC B":
		return 0x8
	case "BIT 4,(HL)":
		return 0x66
	case "BIT 5,A":
		return 0x6F
	case "SET 4,B":
		return 0xE0
	case "RRC E":
		return 0xB
	case "BIT 4,E":
		return 0x63
	case "BIT 2,H":
		return 0x54
	case "RES 0,H":
		return 0x84
	case "RES 1,D":
		return 0x8A
	case "RES 5,A":
		return 0xAF
	case "SET 6,C":
		return 0xF1
	case "RLC D":
		return 0x2
	case "SRA L":
		return 0x2D
	case "BIT 5,(HL)":
		return 0x6E
	case "RES 4,E":
		return 0xA3
	case "SET 3,A":
		return 0xDF
	case "SET 4,E":
		return 0xE3
	case "RES 3,D":
		return 0x9A
	case "RLC B":
		return 0x0
	case "RES 2,L":
		return 0x95
	case "SLA B":
		return 0x20
	case "SRA H":
		return 0x2C
	case "RES 0,(HL)":
		return 0x86
	case "RES 7,H":
		return 0xBC
	case "BIT 3,L":
		return 0x5D
	case "BIT 6,E":
		return 0x73
	case "BIT 7,E":
		return 0x7B
	case "RL H":
		return 0x14
	case "BIT 0,A":
		return 0x47
	case "BIT 7,L":
		return 0x7D
	case "RLC H":
		return 0x4
	case "SLA A":
		return 0x27
	case "SWAP L":
		return 0x35
	case "BIT 7,B":
		return 0x78
	case "RES 1,C":
		return 0x89
	case "RES 5,B":
		return 0xA8
	case "RRC A":
		return 0xF
	case "RL B":
		return 0x10
	case "RR (HL)":
		return 0x1E
	case "SLA H":
		return 0x24
	case "SRA B":
		return 0x28
	case "SET 4,D":
		return 0xE2
	case "SET 4,(HL)":
		return 0xE6
	case "SET 6,E":
		return 0xF3
	case "BIT 5,C":
		return 0x69
	case "SET 2,A":
		return 0xD7
	case "SLA C":
		return 0x21
	case "BIT 0,L":
		return 0x45
	case "SET 1,B":
		return 0xC8
	case "SET 2,B":
		return 0xD0
	case "SWAP (HL)":
		return 0x36
	case "BIT 7,C":
		return 0x79
	case "RES 1,A":
		return 0x8F
	case "RES 3,B":
		return 0x98
	case "SET 6,(HL)":
		return 0xF6
	case "BIT 0,E":
		return 0x43
	case "BIT 1,C":
		return 0x49
	case "BIT 5,E":
		return 0x6B
	case "RES 4,C":
		return 0xA1
	case "RES 6,C":
		return 0xB1
	case "SET 1,E":
		return 0xCB
	case "SET 5,A":
		return 0xEF
	case "RR C":
		return 0x19
	case "BIT 6,B":
		return 0x70
	case "RES 7,B":
		return 0xB8
	case "SET 4,A":
		return 0xE7
	case "SET 6,B":
		return 0xF0
	case "SET 6,H":
		return 0xF4
	case "SWAP B":
		return 0x30
	case "SWAP E":
		return 0x33
	case "SRL D":
		return 0x3A
	case "BIT 6,(HL)":
		return 0x76
	case "SET 2,H":
		return 0xD4
	case "SET 7,L":
		return 0xFD
	case "BIT 4,C":
		return 0x61
	case "BIT 5,B":
		return 0x68
	case "RES 3,C":
		return 0x99
	case "SET 5,D":
		return 0xEA
	case "SET 5,H":
		return 0xEC
	case "SET 6,D":
		return 0xF2
	case "SET 6,L":
		return 0xF5
	case "RLC A":
		return 0x7
	case "RES 5,(HL)":
		return 0xAE
	case "RES 7,D":
		return 0xBA
	case "RRC D":
		return 0xA
	case "RRC (HL)":
		return 0xE
	case "RL L":
		return 0x15
	case "RES 1,L":
		return 0x8D
	case "RES 1,(HL)":
		return 0x8E
	case "RES 5,E":
		return 0xAB
	case "SET 2,E":
		return 0xD3
	case "RR B":
		return 0x18
	case "SLA E":
		return 0x23
	case "RES 3,H":
		return 0x9C
	case "RES 4,(HL)":
		return 0xA6
	case "SET 3,D":
		return 0xDA
	case "BIT 1,A":
		return 0x4F
	case "RR D":
		return 0x1A
	case "RES 2,E":
		return 0x93
	case "SET 7,E":
		return 0xFB
	case "BIT 1,(HL)":
		return 0x4E
	case "RES 3,(HL)":
		return 0x9E
	case "SET 0,A":
		return 0xC7
	case "SET 5,(HL)":
		return 0xEE
	case "RL E":
		return 0x13
	case "SWAP C":
		return 0x31
	case "BIT 1,H":
		return 0x4C
	case "RES 2,D":
		return 0x92
	case "SET 4,H":
		return 0xE4
	case "BIT 2,D":
		return 0x52
	case "SET 0,(HL)":
		return 0xC6
	case "SET 5,L":
		return 0xED
	case "BIT 1,B":
		return 0x48
	case "RES 7,E":
		return 0xBB
	case "SRL (HL)":
		return 0x3E
	case "BIT 1,L":
		return 0x4D
	case "BIT 6,D":
		return 0x72
	case "RES 7,(HL)":
		return 0xBE
	case "BIT 2,C":
		return 0x51
	case "BIT 6,C":
		return 0x71
	case "RL A":
		return 0x17
	case "BIT 0,D":
		return 0x42
	case "SET 2,(HL)":
		return 0xD6
	case "RLC C":
		return 0x1
	case "BIT 1,E":
		return 0x4B
	case "SET 3,B":
		return 0xD8
	case "SET 3,C":
		return 0xD9
	case "SLA (HL)":
		return 0x26
	case "SWAP A":
		return 0x37
	case "SET 0,B":
		return 0xC0
	case "RLC E":
		return 0x3
	case "RES 2,(HL)":
		return 0x96
	case "RES 2,A":
		return 0x97
	case "SRL B":
		return 0x38
	case "BIT 6,H":
		return 0x74
	case "RES 0,E":
		return 0x83
	case "SET 0,H":
		return 0xC4
	case "SET 0,C":
		return 0xC1
	case "SWAP D":
		return 0x32
	case "BIT 0,(HL)":
		return 0x46
	case "BIT 7,D":
		return 0x7A
	case "BIT 7,(HL)":
		return 0x7E
	case "RES 6,B":
		return 0xB0
	case "SET 3,(HL)":
		return 0xDE
	case "SRA C":
		return 0x29
	case "BIT 2,A":
		return 0x57
	case "BIT 3,A":
		return 0x5F
	case "RES 5,C":
		return 0xA9
	case "RES 5,D":
		return 0xAA
	case "RL (HL)":
		return 0x16
	case "SRA A":
		return 0x2F
	case "SET 2,L":
		return 0xD5
	case "BIT 3,B":
		return 0x58
	case "RES 0,D":
		return 0x82
	case "RES 6,A":
		return 0xB7
	case "SET 6,A":
		return 0xF7
	case "RES 4,B":
		return 0xA0
	case "RES 4,D":
		return 0xA2

	default:
		panic(fmt.Sprintf("Unknown code for %q", s))
	}
}
func name(code uint8, prefix bool) string {
	if prefix {
		switch code {
		case 179:
			return "RES 6,E"
		case 206:
			return "SET 1,(HL)"
		case 254:
			return "SET 7,(HL)"
		case 52:
			return "SWAP H"
		case 140:
			return "RES 1,H"
		case 90:
			return "BIT 3,D"
		case 157:
			return "RES 3,L"
		case 220:
			return "SET 3,H"
		case 31:
			return "RR A"
		case 86:
			return "BIT 2,(HL)"
		case 181:
			return "RES 6,L"
		case 225:
			return "SET 4,C"
		case 250:
			return "SET 7,D"
		case 46:
			return "SRA (HL)"
		case 129:
			return "RES 0,C"
		case 185:
			return "RES 7,C"
		case 221:
			return "SET 3,L"
		case 17:
			return "RL C"
		case 43:
			return "SRA E"
		case 63:
			return "SRL A"
		case 100:
			return "BIT 4,H"
		case 139:
			return "RES 1,E"
		case 167:
			return "RES 4,A"
		case 6:
			return "RLC (HL)"
		case 65:
			return "BIT 0,C"
		case 83:
			return "BIT 2,E"
		case 103:
			return "BIT 4,A"
		case 119:
			return "BIT 6,A"
		case 180:
			return "RES 6,H"
		case 232:
			return "SET 5,B"
		case 248:
			return "SET 7,B"
		case 18:
			return "RL D"
		case 80:
			return "BIT 2,B"
		case 136:
			return "RES 1,B"
		case 191:
			return "RES 7,A"
		case 201:
			return "SET 1,C"
		case 148:
			return "RES 2,H"
		case 207:
			return "SET 1,A"
		case 37:
			return "SLA L"
		case 61:
			return "SRL L"
		case 85:
			return "BIT 2,L"
		case 101:
			return "BIT 4,L"
		case 108:
			return "BIT 5,H"
		case 124:
			return "BIT 7,H"
		case 164:
			return "RES 4,H"
		case 165:
			return "RES 4,L"
		case 12:
			return "RRC H"
		case 28:
			return "RR H"
		case 92:
			return "BIT 3,H"
		case 128:
			return "RES 0,B"
		case 204:
			return "SET 1,H"
		case 195:
			return "SET 0,E"
		case 209:
			return "SET 2,C"
		case 109:
			return "BIT 5,L"
		case 255:
			return "SET 7,A"
		case 194:
			return "SET 0,D"
		case 189:
			return "RES 7,L"
		case 42:
			return "SRA D"
		case 59:
			return "SRL E"
		case 74:
			return "BIT 1,D"
		case 182:
			return "RES 6,(HL)"
		case 219:
			return "SET 3,E"
		case 9:
			return "RRC C"
		case 57:
			return "SRL C"
		case 94:
			return "BIT 3,(HL)"
		case 98:
			return "BIT 4,D"
		case 197:
			return "SET 0,L"
		case 34:
			return "SLA D"
		case 96:
			return "BIT 4,B"
		case 127:
			return "BIT 7,A"
		case 133:
			return "RES 0,L"
		case 205:
			return "SET 1,L"
		case 235:
			return "SET 5,E"
		case 27:
			return "RR E"
		case 60:
			return "SRL H"
		case 202:
			return "SET 1,D"
		case 210:
			return "SET 2,D"
		case 106:
			return "BIT 5,D"
		case 29:
			return "RR L"
		case 89:
			return "BIT 3,C"
		case 91:
			return "BIT 3,E"
		case 117:
			return "BIT 6,L"
		case 135:
			return "RES 0,A"
		case 144:
			return "RES 2,B"
		case 5:
			return "RLC L"
		case 159:
			return "RES 3,A"
		case 173:
			return "RES 5,L"
		case 155:
			return "RES 3,E"
		case 229:
			return "SET 4,L"
		case 233:
			return "SET 5,C"
		case 249:
			return "SET 7,C"
		case 145:
			return "RES 2,C"
		case 178:
			return "RES 6,D"
		case 252:
			return "SET 7,H"
		case 13:
			return "RRC L"
		case 64:
			return "BIT 0,B"
		case 68:
			return "BIT 0,H"
		case 172:
			return "RES 5,H"
		case 8:
			return "RRC B"
		case 102:
			return "BIT 4,(HL)"
		case 111:
			return "BIT 5,A"
		case 224:
			return "SET 4,B"
		case 11:
			return "RRC E"
		case 99:
			return "BIT 4,E"
		case 84:
			return "BIT 2,H"
		case 132:
			return "RES 0,H"
		case 138:
			return "RES 1,D"
		case 175:
			return "RES 5,A"
		case 241:
			return "SET 6,C"
		case 2:
			return "RLC D"
		case 45:
			return "SRA L"
		case 110:
			return "BIT 5,(HL)"
		case 163:
			return "RES 4,E"
		case 223:
			return "SET 3,A"
		case 227:
			return "SET 4,E"
		case 154:
			return "RES 3,D"
		case 0:
			return "RLC B"
		case 149:
			return "RES 2,L"
		case 32:
			return "SLA B"
		case 44:
			return "SRA H"
		case 134:
			return "RES 0,(HL)"
		case 188:
			return "RES 7,H"
		case 93:
			return "BIT 3,L"
		case 115:
			return "BIT 6,E"
		case 123:
			return "BIT 7,E"
		case 20:
			return "RL H"
		case 71:
			return "BIT 0,A"
		case 125:
			return "BIT 7,L"
		case 4:
			return "RLC H"
		case 39:
			return "SLA A"
		case 53:
			return "SWAP L"
		case 120:
			return "BIT 7,B"
		case 137:
			return "RES 1,C"
		case 168:
			return "RES 5,B"
		case 15:
			return "RRC A"
		case 16:
			return "RL B"
		case 30:
			return "RR (HL)"
		case 36:
			return "SLA H"
		case 40:
			return "SRA B"
		case 226:
			return "SET 4,D"
		case 230:
			return "SET 4,(HL)"
		case 243:
			return "SET 6,E"
		case 105:
			return "BIT 5,C"
		case 215:
			return "SET 2,A"
		case 33:
			return "SLA C"
		case 69:
			return "BIT 0,L"
		case 200:
			return "SET 1,B"
		case 208:
			return "SET 2,B"
		case 54:
			return "SWAP (HL)"
		case 121:
			return "BIT 7,C"
		case 143:
			return "RES 1,A"
		case 152:
			return "RES 3,B"
		case 246:
			return "SET 6,(HL)"
		case 67:
			return "BIT 0,E"
		case 73:
			return "BIT 1,C"
		case 107:
			return "BIT 5,E"
		case 161:
			return "RES 4,C"
		case 177:
			return "RES 6,C"
		case 203:
			return "SET 1,E"
		case 239:
			return "SET 5,A"
		case 25:
			return "RR C"
		case 112:
			return "BIT 6,B"
		case 184:
			return "RES 7,B"
		case 231:
			return "SET 4,A"
		case 240:
			return "SET 6,B"
		case 244:
			return "SET 6,H"
		case 48:
			return "SWAP B"
		case 51:
			return "SWAP E"
		case 58:
			return "SRL D"
		case 118:
			return "BIT 6,(HL)"
		case 212:
			return "SET 2,H"
		case 253:
			return "SET 7,L"
		case 97:
			return "BIT 4,C"
		case 104:
			return "BIT 5,B"
		case 153:
			return "RES 3,C"
		case 234:
			return "SET 5,D"
		case 236:
			return "SET 5,H"
		case 242:
			return "SET 6,D"
		case 245:
			return "SET 6,L"
		case 7:
			return "RLC A"
		case 174:
			return "RES 5,(HL)"
		case 186:
			return "RES 7,D"
		case 10:
			return "RRC D"
		case 14:
			return "RRC (HL)"
		case 21:
			return "RL L"
		case 141:
			return "RES 1,L"
		case 142:
			return "RES 1,(HL)"
		case 171:
			return "RES 5,E"
		case 211:
			return "SET 2,E"
		case 24:
			return "RR B"
		case 35:
			return "SLA E"
		case 156:
			return "RES 3,H"
		case 166:
			return "RES 4,(HL)"
		case 218:
			return "SET 3,D"
		case 79:
			return "BIT 1,A"
		case 26:
			return "RR D"
		case 147:
			return "RES 2,E"
		case 251:
			return "SET 7,E"
		case 78:
			return "BIT 1,(HL)"
		case 158:
			return "RES 3,(HL)"
		case 199:
			return "SET 0,A"
		case 238:
			return "SET 5,(HL)"
		case 19:
			return "RL E"
		case 49:
			return "SWAP C"
		case 76:
			return "BIT 1,H"
		case 146:
			return "RES 2,D"
		case 228:
			return "SET 4,H"
		case 82:
			return "BIT 2,D"
		case 198:
			return "SET 0,(HL)"
		case 237:
			return "SET 5,L"
		case 72:
			return "BIT 1,B"
		case 187:
			return "RES 7,E"
		case 62:
			return "SRL (HL)"
		case 77:
			return "BIT 1,L"
		case 114:
			return "BIT 6,D"
		case 190:
			return "RES 7,(HL)"
		case 81:
			return "BIT 2,C"
		case 113:
			return "BIT 6,C"
		case 23:
			return "RL A"
		case 66:
			return "BIT 0,D"
		case 214:
			return "SET 2,(HL)"
		case 1:
			return "RLC C"
		case 75:
			return "BIT 1,E"
		case 216:
			return "SET 3,B"
		case 217:
			return "SET 3,C"
		case 38:
			return "SLA (HL)"
		case 55:
			return "SWAP A"
		case 192:
			return "SET 0,B"
		case 3:
			return "RLC E"
		case 150:
			return "RES 2,(HL)"
		case 151:
			return "RES 2,A"
		case 56:
			return "SRL B"
		case 116:
			return "BIT 6,H"
		case 131:
			return "RES 0,E"
		case 196:
			return "SET 0,H"
		case 193:
			return "SET 0,C"
		case 50:
			return "SWAP D"
		case 70:
			return "BIT 0,(HL)"
		case 122:
			return "BIT 7,D"
		case 126:
			return "BIT 7,(HL)"
		case 176:
			return "RES 6,B"
		case 222:
			return "SET 3,(HL)"
		case 41:
			return "SRA C"
		case 87:
			return "BIT 2,A"
		case 95:
			return "BIT 3,A"
		case 169:
			return "RES 5,C"
		case 170:
			return "RES 5,D"
		case 22:
			return "RL (HL)"
		case 47:
			return "SRA A"
		case 213:
			return "SET 2,L"
		case 88:
			return "BIT 3,B"
		case 130:
			return "RES 0,D"
		case 183:
			return "RES 6,A"
		case 247:
			return "SET 6,A"
		case 160:
			return "RES 4,B"
		case 162:
			return "RES 4,D"

		}
	}
	switch code {
	case 209:
		return "POP DE"
	case 218:
		return "JP C,a16"
	case 2:
		return "LD (BC),A"
	case 55:
		return "SCF"
	case 247:
		return "RST $30"
	case 230:
		return "AND A,n8"
	case 235:
		return "ILLEGAL_EB"
	case 3:
		return "INC BC"
	case 65:
		return "LD B,C"
	case 72:
		return "LD C,B"
	case 125:
		return "LD A,L"
	case 169:
		return "XOR A,C"
	case 187:
		return "CP A,E"
	case 136:
		return "ADC A,B"
	case 26:
		return "LD A,(DE)"
	case 146:
		return "SUB A,D"
	case 220:
		return "CALL C,a16"
	case 167:
		return "AND A,A"
	case 223:
		return "RST $18"
	case 112:
		return "LD (HL),B"
	case 174:
		return "XOR A,(HL)"
	case 200:
		return "RET Z"
	case 7:
		return "RLCA"
	case 23:
		return "RLA"
	case 64:
		return "LD B,B"
	case 254:
		return "CP A,n8"
	case 134:
		return "ADD A,(HL)"
	case 150:
		return "SUB A,(HL)"
	case 160:
		return "AND A,B"
	case 212:
		return "CALL NC,a16"
	case 217:
		return "RETI"
	case 9:
		return "ADD HL,BC"
	case 12:
		return "INC C"
	case 60:
		return "INC A"
	case 100:
		return "LD H,H"
	case 222:
		return "SBC A,n8"
	case 251:
		return "EI"
	case 252:
		return "ILLEGAL_FC"
	case 82:
		return "LD D,D"
	case 196:
		return "CALL NZ,a16"
	case 216:
		return "RET C"
	case 221:
		return "ILLEGAL_DD"
	case 1:
		return "LD BC,n16"
	case 36:
		return "INC H"
	case 144:
		return "SUB A,B"
	case 193:
		return "POP BC"
	case 207:
		return "RST $08"
	case 232:
		return "ADD SP,e8"
	case 143:
		return "ADC A,A"
	case 240:
		return "LDH A,(a8)"
	case 53:
		return "DEC (HL)"
	case 114:
		return "LD (HL),D"
	case 154:
		return "SBC A,D"
	case 166:
		return "AND A,(HL)"
	case 237:
		return "ILLEGAL_ED"
	case 244:
		return "ILLEGAL_F4"
	case 8:
		return "LD (a16),SP"
	case 19:
		return "INC DE"
	case 10:
		return "LD A,(BC)"
	case 14:
		return "LD C,n8"
	case 39:
		return "DAA"
	case 43:
		return "DEC HL"
	case 58:
		return "LD A,(HL-)"
	case 77:
		return "LD C,L"
	case 253:
		return "ILLEGAL_FD"
	case 33:
		return "LD HL,n16"
	case 38:
		return "LD H,n8"
	case 135:
		return "ADD A,A"
	case 185:
		return "CP A,C"
	case 190:
		return "CP A,(HL)"
	case 231:
		return "RST $20"
	case 90:
		return "LD E,D"
	case 227:
		return "ILLEGAL_E3"
	case 32:
		return "JR NZ,e8"
	case 76:
		return "LD C,H"
	case 50:
		return "LD (HL-),A"
	case 52:
		return "INC (HL)"
	case 224:
		return "LDH (a8),A"
	case 242:
		return "LDH A,(C)"
	case 48:
		return "JR NC,e8"
	case 99:
		return "LD H,E"
	case 140:
		return "ADC A,H"
	case 161:
		return "AND A,C"
	case 107:
		return "LD L,E"
	case 51:
		return "INC SP"
	case 81:
		return "LD D,C"
	case 118:
		return "HALT"
	case 129:
		return "ADD A,C"
	case 130:
		return "ADD A,D"
	case 66:
		return "LD B,D"
	case 78:
		return "LD C,(HL)"
	case 116:
		return "LD (HL),H"
	case 138:
		return "ADC A,D"
	case 4:
		return "INC B"
	case 119:
		return "LD (HL),A"
	case 171:
		return "XOR A,E"
	case 184:
		return "CP A,B"
	case 208:
		return "RET NC"
	case 177:
		return "OR A,C"
	case 108:
		return "LD L,H"
	case 45:
		return "DEC L"
	case 80:
		return "LD D,B"
	case 127:
		return "LD A,A"
	case 149:
		return "SUB A,L"
	case 157:
		return "SBC A,L"
	case 159:
		return "SBC A,A"
	case 170:
		return "XOR A,D"
	case 30:
		return "LD E,n8"
	case 56:
		return "JR C,e8"
	case 96:
		return "LD H,B"
	case 133:
		return "ADD A,L"
	case 197:
		return "PUSH BC"
	case 213:
		return "PUSH DE"
	case 255:
		return "RST $38"
	case 6:
		return "LD B,n8"
	case 29:
		return "DEC E"
	case 59:
		return "DEC SP"
	case 109:
		return "LD L,L"
	case 156:
		return "SBC A,H"
	case 210:
		return "JP NC,a16"
	case 63:
		return "CCF"
	case 35:
		return "INC HL"
	case 49:
		return "LD SP,n16"
	case 57:
		return "ADD HL,SP"
	case 74:
		return "LD C,D"
	case 85:
		return "LD D,L"
	case 91:
		return "LD E,E"
	case 145:
		return "SUB A,C"
	case 148:
		return "SUB A,H"
	case 61:
		return "DEC A"
	case 87:
		return "LD D,A"
	case 106:
		return "LD L,D"
	case 164:
		return "AND A,H"
	case 191:
		return "CP A,A"
	case 205:
		return "CALL a16"
	case 243:
		return "DI"
	case 71:
		return "LD B,A"
	case 155:
		return "SBC A,E"
	case 215:
		return "RST $10"
	case 246:
		return "OR A,n8"
	case 250:
		return "LD A,(a16)"
	case 121:
		return "LD A,C"
	case 152:
		return "SBC A,B"
	case 123:
		return "LD A,E"
	case 137:
		return "ADC A,C"
	case 248:
		return "LD HL,SP+,e8"
	case 20:
		return "INC D"
	case 68:
		return "LD B,H"
	case 126:
		return "LD A,(HL)"
	case 203:
		return "PREFIX"
	case 0:
		return "NOP"
	case 84:
		return "LD D,H"
	case 89:
		return "LD E,C"
	case 194:
		return "JP NZ,a16"
	case 54:
		return "LD (HL),n8"
	case 117:
		return "LD (HL),L"
	case 172:
		return "XOR A,H"
	case 214:
		return "SUB A,n8"
	case 44:
		return "INC L"
	case 173:
		return "XOR A,L"
	case 233:
		return "JP HL"
	case 40:
		return "JR Z,e8"
	case 88:
		return "LD E,B"
	case 141:
		return "ADC A,L"
	case 211:
		return "ILLEGAL_D3"
	case 93:
		return "LD E,L"
	case 103:
		return "LD H,A"
	case 128:
		return "ADD A,B"
	case 226:
		return "LDH (C),A"
	case 239:
		return "RST $28"
	case 79:
		return "LD C,A"
	case 241:
		return "POP AF"
	case 21:
		return "DEC D"
	case 101:
		return "LD H,L"
	case 104:
		return "LD L,B"
	case 92:
		return "LD E,H"
	case 132:
		return "ADD A,H"
	case 153:
		return "SBC A,C"
	case 183:
		return "OR A,A"
	case 201:
		return "RET"
	case 229:
		return "PUSH HL"
	case 249:
		return "LD SP,HL"
	case 28:
		return "INC E"
	case 34:
		return "LD (HL+),A"
	case 182:
		return "OR A,(HL)"
	case 236:
		return "ILLEGAL_EC"
	case 245:
		return "PUSH AF"
	case 131:
		return "ADD A,E"
	case 225:
		return "POP HL"
	case 83:
		return "LD D,E"
	case 147:
		return "SUB A,E"
	case 178:
		return "OR A,D"
	case 22:
		return "LD D,n8"
	case 27:
		return "DEC DE"
	case 41:
		return "ADD HL,HL"
	case 42:
		return "LD A,(HL+)"
	case 73:
		return "LD C,C"
	case 97:
		return "LD H,C"
	case 139:
		return "ADC A,E"
	case 158:
		return "SBC A,(HL)"
	case 175:
		return "XOR A,A"
	case 199:
		return "RST $00"
	case 18:
		return "LD (DE),A"
	case 37:
		return "DEC H"
	case 168:
		return "XOR A,B"
	case 219:
		return "ILLEGAL_DB"
	case 69:
		return "LD B,L"
	case 113:
		return "LD (HL),C"
	case 122:
		return "LD A,D"
	case 198:
		return "ADD A,n8"
	case 186:
		return "CP A,D"
	case 202:
		return "JP Z,a16"
	case 228:
		return "ILLEGAL_E4"
	case 62:
		return "LD A,n8"
	case 67:
		return "LD B,E"
	case 95:
		return "LD E,A"
	case 105:
		return "LD L,C"
	case 111:
		return "LD L,A"
	case 151:
		return "SUB A,A"
	case 188:
		return "CP A,H"
	case 195:
		return "JP a16"
	case 31:
		return "RRA"
	case 75:
		return "LD C,E"
	case 11:
		return "DEC BC"
	case 16:
		return "STOP"
	case 17:
		return "LD DE,n16"
	case 70:
		return "LD B,(HL)"
	case 234:
		return "LD (a16),A"
	case 162:
		return "AND A,D"
	case 181:
		return "OR A,L"
	case 24:
		return "JR e8"
	case 25:
		return "ADD HL,DE"
	case 47:
		return "CPL"
	case 94:
		return "LD E,(HL)"
	case 204:
		return "CALL Z,a16"
	case 238:
		return "XOR A,n8"
	case 98:
		return "LD H,D"
	case 163:
		return "AND A,E"
	case 189:
		return "CP A,L"
	case 5:
		return "DEC B"
	case 179:
		return "OR A,E"
	case 13:
		return "DEC C"
	case 86:
		return "LD D,(HL)"
	case 176:
		return "OR A,B"
	case 206:
		return "ADC A,n8"
	case 46:
		return "LD L,n8"
	case 110:
		return "LD L,(HL)"
	case 115:
		return "LD (HL),E"
	case 120:
		return "LD A,B"
	case 124:
		return "LD A,H"
	case 165:
		return "AND A,L"
	case 180:
		return "OR A,H"
	case 192:
		return "RET NZ"
	case 102:
		return "LD H,(HL)"
	case 142:
		return "ADC A,(HL)"

	default:
		panic(fmt.Sprintf("Unknown code for %d", code))
	}
}
