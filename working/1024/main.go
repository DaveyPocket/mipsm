package main

import ("fmt"
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

func main() {
	/*f, err := os.Open(os.Args[1])
	if err != nil {
		panic(err)
	}
	*/
	f_assemble("SUB $t0, $t1, $t0")
	f_assemble("ADD $t1, $t0, $t0")
	f_assemble("ADD $t2, $t0, $t1")
	f_assemble("ADD $t3, $t2, $t2")
	f_assemble("ADD $t4, $t3, $t2")
	f_assemble("ADD $t5, $t4, $t4")
	f_assemble("ADD $t6, $t5, $t4")
	f_assemble("ADD $t7, $t6, $t6")
	f_assemble("ADD $t8, $t7, $t6")
	fmt.Println(prog)
	fmt.Println(instrType["BEQ"])
}

func f_assemble(in string) {
	var t_opcode, t_function, t_rs, t_rd, t_rt, t_shamt uint8
	switch in[0:3] {
	case "ADD":
		t_opcode = 0
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
	}
	// The pretty-print below is of key importance. Correctly segments the fields of an R-Type instruction into a proper 32-bit hex string.
	prog += fmt.Sprintf("x\"%x%x\",x\"%x%x\",x\"%x%x\",x\"%x%x\", \n", (t_opcode&0x3C)>>2, (t_opcode&3) << 2 + (t_rs&0x18)>>3, (t_rs&0x07)<<1 + (t_rt&0x10)>>4, (t_rt&0x0F), (t_rd&0x1E)>>1, (t_rd&0x01)<<3 + (t_shamt&0x1C)>>2, (t_shamt&0x03)<<2 + (t_function&0x38)>>4, (t_function&0x0F))
}

func f_getRType(in string) (uint8, uint8, uint8){
	return f_getReg(in, 0), f_getReg(in, 1), f_getReg(in, 2)
}

func f_getReg(in string, c int) uint8 {
	var count int = -1
	for i, val := range in {
		if val == '$' {
			count++
		}
		if count == c {
			fmt.Print(in[i:i+3], " ")
			return regAddr[in[i:i+3]]
		}
	}
	return 255
}
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
	"BEQ": {t_I, 0x04, 0},
}

var regAddr map[string]uint8 = map[string]uint8 {
	"$zero": 0,
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
