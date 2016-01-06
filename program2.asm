.text
.global main
#$s0 - 'Previous' Value
#$s1 - 'Current' value
main:
	addi $s1, $zero, 1 #Initialize current value with 1.
	loop:					#Begin loop.
	add $s0, $s1, $s1
	or $a0, $s0, $zero
	jal EVEN
	add $s1, $s0, $s1
	or $a0, $s1, $zero
	jal EVEN
	j loop		#Loop.

EVEN:
	nor $a0, $a0, $zero
	andi $v0, $a0, 1	#Bitmask zeroth bit of current value.
	jr $ra
