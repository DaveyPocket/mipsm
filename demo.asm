# Short demo for remote presentations

begin:	#Initialize
	ori $t0, $zero, 48
	lw $s0, 0($t0)	#Read input switches; address 48
	sw $s0, 4($t0)	#Display number; address 52

loop:
	addi $s0, $s0, -1	#Decrement
	sw $s0, 4($t0)		#Display
	bne $s0, $zero, loop	#Conditional
	
	#Magic constant
	ori $t5, $zero, 49153
	sw $t5, 4($t0)

end:
	#Loop forever....
	beq $zero, $zero, end
