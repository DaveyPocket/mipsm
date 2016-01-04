/*
   MIPSm - A "MIPS" assembler.
   Copyright (C) 2015-2016 Bradley Boccuzzi

   This program is free software: you can redistribute it and/or modify
   it under the terms of the GNU General Public License as published by
   the Free Software Foundation, either version 3 of the License, or
   (at your option) any later version.

   This program is distributed in the hope that it will be useful,
   but WITHOUT ANY WARRANTY; without even the implied warranty of
   MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
   GNU General Public License for more details.

   You should have received a copy of the GNU General Public License
   along with this program.  If not, see <http://www.gnu.org/licenses/>.
*/
//	mipsm.go - The main program.
//TODO
//Clean
//Order for rt and rs fields for BEQ and BNE instructions
//,, for instructions lacking register fields
//Immediate sizes
//Support for pseudo instructions
package main

import (
	"bufio" // Reading one byte per time from the file
	"fmt"
	"mipsm/parser"
	"mipsm/pretty"
	"os" //Include os for reading stuff
	"strconv"
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

func main() {
	f, err := os.Open("program.asm")
	defer f.Close() //	Close the file at the end of the program
	if err != nil {
		panic(err)
	}
	buf := bufio.NewScanner(f)
	for buf.Scan() {
		//	Fill the symbol table.
		parser.Parse(buf.Text())
	}
	if err := buf.Err(); err != nil {
		//	Scanner returns an non-nil error after stopping a scan not due to EOF.
		panic(err)
	} else {
		//	EOF
		fmt.Println("Initial parse successful.")
	}
	f.Seek(0, 0) // Reset seeking pointer for reading in file into bufio buffer.
	//	Add error handling
	buf = bufio.NewScanner(f)
	//	Using default split function (Change this?)
	parser.ResetCounter()
	for buf.Scan() {
		f_assemble(buf.Text())
	}
	if err := buf.Err(); err != nil {
		//	Scanner returns an non-nil error after stopping a scan not due to EOF.
		panic(err)
	} else {
		//	EOF
		fmt.Println("Assembly successful.")
	}

	fmt.Printf("%v", parser.GetEntireTable())
}

//	F_assemble is an exported function of the mipsm package. The function takes in a single line of MIPS assembly code and pretty-prints a string of the assembled machine code line.
func f_assemble(in string) {
	var t_opcode, t_function, t_rs, t_rd, t_rt, t_shamt uint8
	var t_imm uint16
	var t_imm2 uint32
	instruction := parser.Parse(in)
	switch instType := instruction.(type) {
	case parser.RType:
		t_meta := coreInstrType[instType.Opcode]
		t_opcode, t_function = t_meta.opcode, t_meta.funct
		t_rd, t_rs, t_rt = f_getRType(instType)
		fmt.Println(pretty.PrintRType(t_opcode, t_function, t_rs, t_rd, t_rt, t_shamt))
	case parser.IType:
		t_meta := coreInstrType[instType.Opcode]
		t_opcode = t_meta.opcode
		t_rd, t_rs, t_imm = f_getIType(instType)
		fmt.Println(pretty.PrintIType(t_opcode, t_rs, t_rt, t_imm))
	case parser.JType:
		t_meta := coreInstrType[instType.Opcode]
		t_opcode = t_meta.opcode
		t_imm2 = f_getJType(instType)
		fmt.Println(pretty.PrintJType(t_opcode, t_imm2))
	}
}

//	f_getRType returns three 8-bit unsigned integer register indices corresponding to the fields of an R-Type instruction.
func f_getRType(in parser.RType) (uint8, uint8, uint8) {
	return f_getReg(in.Rd), f_getReg(in.Rs), f_getReg(in.Rt)
}

//	f_GetIType returns two register indices and a 16-bit unsigned integer corresponding to the fields of an I-Type instruction.
func f_getIType(in parser.IType) (uint8, uint8, uint16) {
	return f_getReg(in.Rt), f_getReg(in.Rs), uint16(f_getImm(in.Imm))
}

// f_getJType returns a 32-bit immediate value associated with the pseudo-direct address of the J-Type instruction.
// TODO - Strive for correctness.
func f_getJType(in parser.JType) uint32 {
	return f_getImm(in.Imm)
}

// f_getJLiteral gets resolves a pseudo-direct address for the given label. Symbol table contains labels and associated address values. DOES NOT HANDLE CLEANING UP STRING!!!
//func f_getJLiteral(in string) uint32 {
//}

// f_getReg returns the register index associated with c'th register denoted in the assembly instruction line.
func f_getReg(in string) uint8 {
	return regAddr[in]
}

// f_getImm returns an immediate value obtained from a simplified instruction string.
func f_getImm(in string) uint32 {
	num, err := strconv.ParseInt(in, 10, 32)
	if err != nil {
		panic(err)
	}
	return uint32(num)
}

// MIPS R-Type
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
// OR	0	0x25
// XOR	0	0x26
// SUBU	0	0x23
// SLT	0	0x2A
// SLTU	0	0x2B

// The map below maps instructions to an t_instrType datatype that contains the opcode/function, as well as the addressing mode of the instruction.
var coreInstrType map[string]t_instrType = map[string]t_instrType{
	"add":   {t_R, 0, 0x20},
	"addu":  {t_R, 0, 0x21},
	"sub":   {t_R, 0, 0x22},
	"subu":  {t_R, 0, 0x23},
	"and":   {t_R, 0, 0x24},
	"or":    {t_R, 0, 0x25},
	"xor":   {t_R, 0, 0x26},
	"nor":   {t_R, 0, 0x27},
	"slt":   {t_R, 0, 0x2A},
	"sltu":  {t_R, 0, 0x2B},
	"jr":    {t_R, 0, 0x08},
	"beq":   {t_I, 0x04, 0},
	"bne":   {t_I, 0x05, 0},
	"addi":  {t_I, 0x08, 0},
	"addiu": {t_I, 0x09, 0},
	"andi":  {t_I, 0x0C, 0},
	"ori":   {t_I, 0x0F, 0},
	"lw":    {t_I, 0x23, 0},
	"sw":    {t_I, 0x2B, 0},
	"slti":  {t_I, 0x0A, 0},
	"sltiu": {t_I, 0x0B, 0},
	"j":     {t_J, 0x02, 0},
	"jal":   {t_J, 0x03, 0},
}

// The map below maps the register strings to their corresponding 8-bit unsigned integer index.
var regAddr map[string]uint8 = map[string]uint8{
	"$zero": 0,
	//	"$ze":   0,
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
