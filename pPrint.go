package mipsm

import (
	"fmt"
)

func PrintRType(t_opcode, t_function, t_rs, t_rd, t_rt, t_shamt uint8) string {
	return fmt.Sprintf("x\"%x%x\",x\"%x%x\",x\"%x%x\",x\"%x%x\", \n", (t_opcode&0x3C)>>2, (t_opcode&3)<<2+(t_rs&0x18)>>3, (t_rs&0x07)<<1+(t_rt&0x10)>>4, (t_rt & 0x0F), (t_rd&0x1E)>>1, (t_rd&0x01)<<3+(t_shamt&0x1C)>>2, (t_shamt&0x03)<<2+(t_function&0x38)>>4, (t_function & 0x0F))
}

func PrintIType(t_opcode, t_rs, t_rt uint8, t_imm uint32) string {
	return fmt.Sprintf("x\"%x%x\",x\"%x%x\",x\"%x%x\",x\"%x%x\", \n", (t_opcode&0x3C)>>2, (t_opcode&3)<<2+(t_rs&0x18)>>3, (t_rs&0x07)<<1+(t_rt&0x10)>>4, (t_rt & 0x0F), (t_imm&0xF000)>>12, (t_imm&0x0F00)>>8, (t_imm&0x00F0)>>4, (t_imm & 0x000F))
}

func PrintJType(t_opcode uint8, t_imm uint32) string {
	return fmt.Sprintf("x\"%x%x\",x\"%x%x\",x\"%x%x\",x\"%x%x\", \n", (t_opcode&0x3C)>>2, (t_opcode&3)<<2+uint8(((t_imm2&0x0F000000)>>24)&2), (t_imm2&0x00F00000)>>20, (t_imm2&0x000F0000)>>16, (t_imm2&0x0000F000)>>12, (t_imm2&0x00000F00)>>8, (t_imm2&0x000000F0)>>4, (t_imm2 & 0x0000000F))
}
