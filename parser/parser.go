//TODO error handling of incorrect lines of code.
//TODO resolve indirect or psedo-direct address from a label for I-Type and J-Type instructions
//--Implement via a symmbol table during the parsing phase??
package parser

import (
	"regexp"
	"strings"
)

var counter int = 0

func ResetCounter() {
	counter = 0
}

// \s*(([^\s:]+):*)\s.*
func Parse(input string) interface{} {
	re := regexp.MustCompile("\\s*(([^\\s:]+):*)\\s.*")
	result := re.FindStringSubmatch(input)
	if result == nil {
		return nil
	}
	counter++
	return coreFuncMap[strings.ToLower(result[1])](input)
}

func parseRType(input string) interface{} {
	re := regexp.MustCompile("\\s*(\\S+)\\s+([^,]+),\\s*([^,]+),\\s*([^,]+).*")
	result := re.FindStringSubmatch(input)
	return RType{result[1], result[2], result[3], result[4]}
}

//TODO Modify the regular expressions to support comments adjacent to the last meaningful object in the line.
func parseJType(input string) interface{} {
	re := regexp.MustCompile("\\s*([^\\s]+)\\s([^\\s]+).*")
	result := re.FindStringSubmatch(input)
	return JType{result[1], result[2]}
}

//TODO change regex to accept only numerical values
func parseIDirect(input string) interface{} {
	re := regexp.MustCompile("\\s*(\\S+)\\s+([^,]+),\\s*([^,]+),\\s*([^\\s]+)")
	result := re.FindStringSubmatch(input)
	return IType{result[1], result[2], result[3], result[4]}
}

//\s*(\S+[^:\s*])\s+([^,]+),\s*([^,]+),\s*([^\s]+)????
// !!! Branching routines have Rs and Rt swapped.
func parseIBranch(input string) interface{} {
	re := regexp.MustCompile("\\s*(\\S+)\\s+([^,]+),\\s*([^,]+),\\s*([^\\s]+)")
	result := re.FindStringSubmatch(input)
	return IType{result[1], result[3], result[2], result[4]}
}

// ParseLabel adds a label to the symbol table with its immediate address
//\s*(.+):.*

//\s*(\S+)\s+([^,]+),\s*(\d+)\(([^\)]+)\).*
//parseIIndirect is used for load and store operations
//Load word and store word are different (Order of the operands)
func parseIIndirect(input string) interface{} {
	re := regexp.MustCompile("\\s*(\\S+)\\s+([^,]+),\\s*(\\-\\d+|\\d+)\\(([^\\)]+)\\).*")
	result := re.FindStringSubmatch(input)
	if strings.ToLower(result[1]) == "sw" || strings.ToLower(result[1]) == "sb" {
		return IType{result[1], result[4], result[2], result[3]}
	}
	return IType{result[1], result[2], result[4], result[3]}
}

type RType struct {
	Opcode, Rd, Rs, Rt string
}

type JType struct {
	Opcode, Imm string
}

type IType struct {
	Opcode, Rt, Rs, Imm string
}

/*
type IType struct {
	// Multiple types of I-type instructions
	Opcode, Rs, Rt, Imm string
}*/

//TODO parse MIPS pseudo instructions
//TODO MIPS R-type shifting instructions
var coreFuncMap map[string](func(string) interface{}) = map[string](func(string) interface{}){
	"add":   parseRType,
	"addu":  parseRType,
	"sub":   parseRType,
	"subu":  parseRType,
	"and":   parseRType,
	"or":    parseRType,
	"xor":   parseRType,
	"nor":   parseRType,
	"slt":   parseRType,
	"sltu":  parseRType,
	"jr":    parseRType, // Special case of R-Type
	"beq":   parseIBranch,
	"bne":   parseIBranch,
	"addi":  parseIDirect,
	"addiu": parseIDirect,
	"andi":  parseIDirect,
	"ori":   parseIDirect,
	"lw":    parseIIndirect,
	"sw":    parseIIndirect,
	"lb":    parseIIndirect,
	"sb":    parseIIndirect,
	"slti":  parseIDirect,
	"sltiu": parseIDirect,
	"j":     parseJType,
	"jal":   parseJType,
}
