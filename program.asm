label:add $t0, $t1, $s0

   

loop: addi $t0, $t4, 5
addi $t0, $t4, 5
addi $t0, $t4, 5
addi $t0, $t4, 5
addi $t0, $t4, 5
addi $t0, $t4, 5
addi $t0, $t4, 5
addi $t0, $t4, 5
addi $t0, $t4, 5
addi $t0, $t4, 5
addi $t0, $t4, 5
addi $t0, $t4, 5
addi $t0, $t4, 5
addi $t0, $t4, 5
addi $t0, $t4, -1

beq $s0, $s5, loop
more_label:sltiu $s5, $ra, 5
