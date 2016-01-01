//TODO
//Clean
//Order for rt and rs fields for BEQ and BNE instructions
//,, for instructions lacking register fields
//Immediate sizes

package mipsm

import (
	"fmt"
	"pPrint"
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
//	t_instrType is used to identify the instruction addressing mode being used, the opcode, and the (R-Type) function code.
type t_instrType struct {
	fam    t_family
	opcode uint8
	funct  uint8
}

//	The t_family type
type t_family uint8

const (
	t_R t_family = iota
	t_I
	t_J
)

func F_getProgString() string {
	return prog
}

// F_assemble is an exported function of the mipsm package. The function takes in a single line of MIPS assembly code and pretty-prints a string of the assembled machine code line.
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
		pPrint.PrintRType(t_opcode, t_function, t_rs, t_rd, t_rt, t_shamt)
	case t_I:
		t_opcode = t_thing.opcode
		t_rt, t_rs, t_imm = f_getIType(in)
		pPrint.PrintIType(t_opcode, t_rs, t_rt, t_imm)
	case t_J:
		t_opcode = t_thing.opcode
		t_imm2 = (f_getJType(in))
		pPrint.PrintJType(t_opcode, t_imm)
	}
}

//	f_getRType returns three 8-bit unsigned integer register indices corresponding to the fields of an R-Type instruction.
func f_getRType(in string) (uint8, uint8, uint8) {
	return f_getReg(in, 0), f_getReg(in, 1), f_getReg(in, 2)
}

//	f_GetIType returns two register indices and a 16-bit unsigned integer corresponding to the fields of an I-Type instruction.
func f_getIType(in string) (uint8, uint8, uint16) {
	return f_getReg(in, 0), f_getReg(in, 1), uint16(f_getImm(in))
}

// f_getJType returns a 32-bit immediate value associated with the pseudo-direct address of the J-Type instruction.
// TODO - Strive for correctness.
func f_getJType(in string) uint32 {
	return f_getImm(in)
}

// f_getReg returns the register index associated with c'th register denoted in the assembly instruction line.
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

// f_getImm returns an immediate value obtained from a simplified instruction string.
func f_getImm(in string) uint32 {
	var count int = 0
	for i, val := range in {
		if val == ',' {
			count++
		}
		if count == 2 {
			num, _ := strconv.ParseInt(in[i+2:len(in)], 10, 32)
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

// The map below maps instructions to an t_instrType datatype that contains the opcode/function, as well as the addressing mode of the instruction.
var instrType map[string]t_instrType = map[string]t_instrType{
	"ADD":   {t_R, 0, 0x20},
	"ADDU":  {t_R, 0, 0x21},
	"SUB":   {t_R, 0, 0x22},
	"SUBU":  {t_R, 0, 0x23},
	"AND":   {t_R, 0, 0x24},
	"OR":    {t_R, 0, 0x25},
	"XOR":   {t_R, 0, 0x26},
	"NOR":   {t_R, 0, 0x27},
	"SLT":   {t_R, 0, 0x2A},
	"SLTU":  {t_R, 0, 0x2B},
	"JR":    {t_R, 0, 0x08},
	"BEQ":   {t_I, 0x04, 0},
	"BNE":   {t_I, 0x05, 0},
	"ADDI":  {t_I, 0x08, 0},
	"ADDIU": {t_I, 0x09, 0},
	"ANDI":  {t_I, 0x0C, 0},
	"ORI":   {t_I, 0x0F, 0},
	"LW":    {t_I, 0x23, 0},
	"SW":    {t_I, 0x2B, 0},
	"SLTI":  {t_I, 0x0A, 0},
	"SLTIU": {t_I, 0x0B, 0},
	"J":     {t_J, 0x02, 0},
	"JAL":   {t_J, 0x03, 0},
}

// The map below maps the register strings to their corresponding 8-bit unsigned integer index.
var regAddr map[string]uint8 = map[string]uint8{
	"$zero": 0,
	"$ze":   0,
	"$at":   1,
	"$v0":   2,
	"$v1":   3,
	"$a0":   4,
	"$a1":   5,
	"$a2":   6,
	"$a3":   7,
	"$t0":   8,
	"$t1":   9,
	"$t2":   10,
	"$t3":   11,
	"$t4":   12,
	"$t5":   13,
	"$t6":   14,
	"$t7":   15,
	"$s0":   16,
	"$s1":   17,
	"$s2":   18,
	"$s3":   19,
	"$s4":   20,
	"$s5":   21,
	"$s6":   22,
	"$s7":   23,
	"$t8":   24,
	"$t9":   25,
	"$k0":   26,
	"$k1":   27,
	"$gp":   28,
	"$sp":   29,
	"$fp":   30,
	"$ra":   31,
}
