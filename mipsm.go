//TODO
//Clean
//Order for rt and rs fields for BEQ and BNE instructions
//,, for instructions lacking register fields
//Immediate sizes

package mipsm

import ("fmt"
			"strconv"
		//	"os"
		)
var prog string
/* Manual

f_assemble - Assemble a single assemble line into machine code

Key

f_* - Temporary name for functions
t_* - Temporary name for types

*/
type t_instrType struct {
	fam t_family
	opcode uint8
	funct uint8
}

type t_family uint8
const(t_R t_family = iota
		t_I
		t_J
		)
func F_getProgString() string {
	return prog
}

func F_assemble(in string) {
	var t_opcode, t_function, t_rs, t_rd, t_rt, t_shamt uint8
	var t_imm uint16
	var t_imm2 uint32
	var inst string
	for _, val := range in {
		if val != ' ' {
			inst += string(val)
		} else {
			break
		}
	}
	t_thing := instrType[inst]
	inst = ""
	switch t_thing.fam {
	case t_R:
		t_opcode, t_function = t_thing.opcode, t_thing.funct
		t_rd, t_rs, t_rt = f_getRType(in)
		prog += fmt.Sprintf("x\"%x%x\",x\"%x%x\",x\"%x%x\",x\"%x%x\", \n", (t_opcode&0x3C)>>2, (t_opcode&3) << 2 + (t_rs&0x18)>>3, (t_rs&0x07)<<1 + (t_rt&0x10)>>4, (t_rt&0x0F), (t_rd&0x1E)>>1, (t_rd&0x01)<<3 + (t_shamt&0x1C)>>2, (t_shamt&0x03)<<2 + (t_function&0x38)>>4, (t_function&0x0F))
	case t_I:
		t_opcode = t_thing.opcode
		t_rt, t_rs, t_imm = f_getIType(in)
		prog += fmt.Sprintf("x\"%x%x\",x\"%x%x\",x\"%x%x\",x\"%x%x\", \n", (t_opcode&0x3C)>>2, (t_opcode&3)<<2 + (t_rs&0x18)>>3, (t_rs&0x07)<<1 + (t_rt&0x10)>>4, (t_rt&0x0F), (t_imm&0xF000) >> 12, (t_imm&0x0F00) >> 8, (t_imm&0x00F0) >> 4, (t_imm&0x000F))
	case t_J:
		t_opcode = t_thing.opcode
		t_imm2 = (f_getJType(in))
		prog += fmt.Sprintf("x\"%x%x\",x\"%x%x\",x\"%x%x\",x\"%x%x\", \n", (t_opcode&0x3C)>>2, (t_opcode&3) << 2 + uint8(((t_imm2&0x0F000000) >> 24)&2), (t_imm2&0x00F00000) >> 20,(t_imm2&0x000F0000) >> 16, (t_imm2&0x0000F000) >> 12, (t_imm2&0x00000F00) >> 8, (t_imm2&0x000000F0) >> 4, (t_imm2&0x0000000F))
	}
/*	switch in[0:3] {
	case "ADD":
		t_thing := instrType[in[0:3]]
		t_function = 0x20
		fmt.Print(in[0:3], " ")
		t_rd, t_rs, t_rt = f_getRType(in)
		fmt.Print(" - ")
		fmt.Printf("%x %x %x %x\n", t_function, t_rd, t_rs, t_rt)
	case "ADDU":
		t_opcode = 0
		t_function = 0x21
	case "SUB":
		t_opcode = 0
		t_function = 0x22
		fmt.Print(in[0:3], " ")
		t_rd, t_rs, t_rt = f_getRType(in)
		fmt.Print(" - ")
		fmt.Printf("%x %x %x %x\n", t_function, t_rd, t_rs, t_rt)
	case "SUBU":
		t_opcode = 0
		t_function = 0x23
	case "AND":
		t_opcode = 0
		t_function = 0x24
	case "NOR":
		t_opcode = 0
		t_function = 0x27
	case "OR":
		t_opcode = 0
		t_function = 0x25
	case "XOR":
		t_opcode = 0
		t_function = 0x26
	case "SLT":
		t_opcode = 0
		t_function = 0x2A
	case "SLTU":
		t_opcode = 0
		t_function = 0x2B
	}*/
	// The pretty-print below is of key importance. Correctly segments the fields of an R-Type instruction into a proper 32-bit hex string.
}

func f_getRType(in string) (uint8, uint8, uint8){
	return f_getReg(in, 0), f_getReg(in, 1), f_getReg(in, 2)
}


func f_getIType(in string) (uint8, uint8, uint16){
	return f_getReg(in, 0), f_getReg(in, 1), uint16(f_getImm(in))
}

func f_getJType(in string) (uint32){
	return f_getImm(in)
}

func f_getReg(in string, c int) uint8 {
	var count int = -1
	for i, val := range in {
		if val == '$' {
			count++
		}
		if count == c {
			//fmt.Print(in[i:i+3], " ")
			return regAddr[in[i:i+3]]
		}
	}
	return 255
}

func f_getImm(in string) uint32 {
	var count int = 0
	for i, val := range in {
		if val == ',' {
			count++
		}
		if count == 2 {
			num, _ := strconv.ParseInt(in[i + 2:len(in)], 10, 32)
			return uint32(num)
			//return regAddr[in[i:len(in) - i]]
		}
	}
	return 255
}
//	MIPS R-Type
// OPCODE - RS - RT - RD - SHAMT - FUNCT
// MIPS I-TYPE
// OPCODE - RS - RT - IMMEDIATE
// MIPS J-TYPE
// OPCODE - ADDRESS

// Lexer
// Parser
// Code generator

// Assembly programs are one-to-one with machine code.

// Instruction - opcode - function(hex)
// ADD	0	0x20
// SUB	0	0x22
// AND	0	0x24
// ADDU	0	0x21
// NOR	0	0x27
// OR		0	0x25
// XOR	0	0x26
// SUBU	0	0x23
// SLT	0	0x2A
// SLTU	0	0x2B
var instrType map[string]t_instrType = map[string]t_instrType {
	"ADD": {t_R, 0, 0x20},
	"ADDU": {t_R, 0, 0x21},
	"SUB": {t_R, 0, 0x22},
	"SUBU": {t_R, 0, 0x23},
	"AND": {t_R, 0, 0x24},
	"OR": {t_R, 0, 0x25},
	"XOR": {t_R, 0, 0x26},
	"NOR": {t_R, 0, 0x27},
	"SLT": {t_R, 0, 0x2A},
	"SLTU": {t_R, 0, 0x2B},
	"JR": {t_R, 0, 0x08},
	"BEQ": {t_I, 0x04, 0},
	"BNE": {t_I, 0x05, 0},
	"ADDI": {t_I, 0x08, 0},
	"ADDIU": {t_I, 0x09, 0},
	"ANDI": {t_I, 0x0C, 0},
	"ORI": {t_I, 0x0F, 0},
	"LW": {t_I, 0x23, 0},
	"SW": {t_I, 0x2B, 0},
	"SLTI": {t_I, 0x0A, 0},
	"SLTIU": {t_I, 0x0B, 0},
	"J": {t_J, 0x02, 0},
	"JAL": {t_J, 0x03, 0},
}

var regAddr map[string]uint8 = map[string]uint8 {
	"$zero": 0,
	"$ze": 0,
	"$at": 1,
	"$v0": 2,
	"$v1": 3,
	"$a0": 4,
	"$a1": 5,
	"$a2": 6,
	"$a3": 7,
	"$t0": 8,
	"$t1": 9,
	"$t2": 10,
	"$t3": 11,
	"$t4": 12,
	"$t5": 13,
	"$t6": 14,
	"$t7": 15,
	"$s0": 16,
	"$s1": 17,
	"$s2": 18,
	"$s3": 19,
	"$s4": 20,
	"$s5": 21,
	"$s6": 22,
	"$s7": 23,
	"$t8": 24,
	"$t9": 25,
	"$k0": 26,
	"$k1": 27,
	"$gp": 28,
	"$sp": 29,
	"$fp": 30,
	"$ra": 31,
}
