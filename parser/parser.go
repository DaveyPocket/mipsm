package parser

import (
	"regexp"
	"strings"
)

func Parse(input string) interface{} {
	re := regexp.MustCompile("\\s*([^\\s]+)\\s.*")
	result := re.FindStringSubmatch(input)
	return coreFuncMap[strings.ToLower(result[1])](input)

}

func parseRType(input string) interface{} {
	re := regexp.MustCompile("\\s*(\\S+)\\s+([^,]+),\\s*([^,]+),\\s*([^,]+).*")
	result := re.FindStringSubmatch(input)
	return RType{result[1], result[2], result[3], result[4]}
}

type RType struct {
	Opcode, Rd, Rs, Rt string
}

/*
type IType struct {
	// Multiple types of I-type instructions
	Opcode, Rs, Rt, Imm string
}

type JType struct {
	Opcode, Imm string
}
*/

//TODO parse MIPS pseudo instructions

var coreFuncMap map[string](func(string) interface{}) = map[string](func(string) interface{}){
	"add": parseRType,
	/*	"addu":
		"sub":
		"subu":
		"and":
		"or":
		"xor":
		"nor":
		"slt":
		"sltu":
		"jr":
		"beq":
		"bne":
		"addi":
		"addiu":
		"andi":
		"ori":
		"lw":
		"sw":
		"slti":
		"sltiu":
		"j":
		"jal":   */
}
