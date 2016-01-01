package parser

import (
	"regexp"
)

// \s*(\S+)\s+([^,]+),\s*([^,]+),\s*([^,]+).*
func parseRType(input string) RType {
	re := regexp.MustCompile("\\s*(\\S+)\\s+([^,]+),\\s*([^,]+),\\s*([^,]+).*")
	thing := re.FindStringSubmatch(input)
	return RType{thing[1], thing[2], thing[3], thing[4]}
}

type RType struct {
	Opcode, Rd, Rs, Rt string
}
