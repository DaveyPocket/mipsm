//TODO error handling of incorrect lines of code.
//TODO resolve indirect or psedo-direct address from a label for I-Type and J-Type instructions
//--Implement via a symmbol table during the parsing phase??
package parser

import (
	"errors"
	"regexp"
	"strconv"
	"strings"
)

// Counter keeps the current value of the current line being parsed (of relevant assembly code).
var counter int = 0

const pseudoDirectMax int = 67108864 //	2^26, MIPS "Green card"
const wordSize int = 4

func ResetCounter() {
	counter = 0
}

//	Create an empty symbol table map.
//	A map to empty interfaces may not be the best idea...
//	Currently being used to determine an invalid entry in the symbol table
var symTable map[string]int = map[string]int{}

func tableAdd(in string, address int) {
	symTable[in] = address
}

func stripComments(in string) string {
	re := regexp.MustCompile("(.*)\\s*#.*")
	out := re.FindStringSubmatch(in)
	if out == nil {
		return in
	}
	return out[1]
}

func tableApply(in string, address int) {
	if _, ok := symTable[in]; !ok {
		tableAdd(in, address)
	}
}

//	E.g. jal or j instructions.
func resolvePseudoDirect(in string) int {
	pda := TableGet(in)
	if pda >= pseudoDirectMax {
		panic(errors.New("Target pseudo-direct address exceeds maximum allowable size."))
	}
	return pda
}

//	E.g. beq or bne instructions.
func resolvePCRelative(in string) int {
	return (TableGet(in) - counter)
}

//	E.g. loading a constant from memory.
func resolveBase(in string) int {
	return TableGet(in)
}

func TableGet(in string) int {
	return symTable[in]
}

//	Should only be used for debugging/crash report
func GetEntireTable() map[string]int {
	return symTable
}

//	Increments the counter by one word (four bytes)
func incrementCounter() {
	counter += wordSize
}

// \s*(([^\s:]+):*)\s.*
func Parse(input string) interface{} {
	clean := stripComments(input)
	re := regexp.MustCompile("\\s*(([^\\s:]+):*)\\s*([^#]*).*")
	result := re.FindStringSubmatch(clean)
	if result == nil {
		return nil
	}
	//	result[1] - label with colon
	//	result[2] - label without colon
	if result[1] != result[2] {
		// Add to the symbol table (If the symbol does not already exist)
		tableApply(result[2], counter)
		return Parse(result[3])
	}
	incrementCounter()
	return coreFuncMap[strings.ToLower(result[1])](clean)
}

func parseRType(input string) interface{} {
	re := regexp.MustCompile("\\s*(\\S+)\\s+([^,]+),\\s*([^,]+),\\s*([^\\s#]+).*")
	result := re.FindStringSubmatch(input)
	return RType{result[1], result[2], result[3], result[4]}
}

//TODO Modify the regular expressions to support comments adjacent to the last meaningful object in the line.
func parseJType(input string) interface{} {
	re := regexp.MustCompile("\\s*([^\\s]+)\\s([^\\s#]+).*")
	result := re.FindStringSubmatch(input)
	pda := resolvePseudoDirect(result[2])
	return JType{result[1], strconv.Itoa(pda)}
}

func parseIImmediate(input string) interface{} {
	re := regexp.MustCompile("\\s*(\\S+)\\s+([^,]+),\\s*([^,]+),\\s*([^\\s#]+)")
	result := re.FindStringSubmatch(input)
	return IType{result[1], result[2], result[3], result[4]}
}

//\s*(\S+[^:\s*])\s+([^,]+),\s*([^,]+),\s*([^\s]+)????
// !!! Branching routines have Rs and Rt swapped.
func parseIBranch(input string) interface{} {
	re := regexp.MustCompile("\\s*(\\S+)\\s+([^,]+),\\s*([^,]+),\\s*([^\\s#]+).*")
	result := re.FindStringSubmatch(input)
	pcr := resolvePCRelative(result[4])

	return IType{result[1], result[3], result[2], strconv.Itoa(pcr)}
}

// ParseLabel adds a label to the symbol table with its immediate address
//\s*(.+):.*

//\s*(\S+)\s+([^,]+),\s*(\d+)\(([^\)]+)\).*
//parseIBaseOffset is used for load and store operations
//Load word and store word are different (Order of the operands)
func parseIBaseOffset(input string) interface{} {
	re := regexp.MustCompile("\\s*(\\S+)\\s+([^,]+),\\s*(\\-\\d+|\\d+)\\(([^\\)]+)\\).*")
	result := re.FindStringSubmatch(input)
	//if strings.ToLower(result[1]) == "sw" || strings.ToLower(result[1]) == "sb" {
	//		return IType{result[1], result[4], result[2], result[3]}
	//}
	// Is reversing rt and rs necessary?
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
	"addi":  parseIImmediate,
	"addiu": parseIImmediate,
	"andi":  parseIImmediate,
	"ori":   parseIImmediate,
	"lw":    parseIBaseOffset,
	"sw":    parseIBaseOffset,
	"lb":    parseIBaseOffset,
	"sb":    parseIBaseOffset,
	"slti":  parseIImmediate,
	"sltiu": parseIImmediate,
	"j":     parseJType,
	"jal":   parseJType,
}
