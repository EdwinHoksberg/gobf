mov x9, #0
mov x10, #0
mov x11, #0
mov x15, x0
add x9, x9, #1
ldrb w11, [x15, x9]
add x11, x11, #1
strb w11, [x15, x9]
ldrb w11, [x15, x9]
add x11, x11, #1
strb w11, [x15, x9]
ldrb w11, [x15, x9]
cbz w11, #0xec
sub x9, x9, #1
ldrb w11, [x15, x9]
add x11, x11, #1
strb w11, [x15, x9]
ldrb w11, [x15, x9]
add x11, x11, #1
strb w11, [x15, x9]
ldrb w11, [x15, x9]
add x11, x11, #1
strb w11, [x15, x9]
ldrb w11, [x15, x9]
add x11, x11, #1
strb w11, [x15, x9]
ldrb w11, [x15, x9]
add x11, x11, #1
strb w11, [x15, x9]
ldrb w11, [x15, x9]
add x11, x11, #1
strb w11, [x15, x9]
ldrb w11, [x15, x9]
add x11, x11, #1
strb w11, [x15, x9]
ldrb w11, [x15, x9]
add x11, x11, #1
strb w11, [x15, x9]
ldrb w11, [x15, x9]
add x11, x11, #1
strb w11, [x15, x9]
ldrb w11, [x15, x9]
add x11, x11, #1
strb w11, [x15, x9]
ldrb w11, [x15, x9]
add x11, x11, #1
strb w11, [x15, x9]
ldrb w11, [x15, x9]
add x11, x11, #1
strb w11, [x15, x9]
ldrb w11, [x15, x9]
add x11, x11, #1
strb w11, [x15, x9]
add x9, x9, #1
ldrb w11, [x15, x9]
sub x11, x11, #1
strb w11, [x15, x9]
ldrb w11, [x15, x9]
cbnz w11, #0x34
sub x9, x9, #1
ldrb w11, [x15, x9]
cbz w11, #0x6ec
ldrb w11, [x15, x9]
cbz w11, #0x13c
add x9, x9, #1
ldrb w11, [x15, x9]
add x11, x11, #1
strb w11, [x15, x9]
add x9, x9, #1
ldrb w11, [x15, x9]
add x11, x11, #1
strb w11, [x15, x9]
sub x9, x9, #1
sub x9, x9, #1
ldrb w11, [x15, x9]
sub x11, x11, #1
strb w11, [x15, x9]
ldrb w11, [x15, x9]
cbnz w11, #0x100
add x9, x9, #1
ldrb w11, [x15, x9]
cbz w11, #0x170
sub x9, x9, #1
ldrb w11, [x15, x9]
add x11, x11, #1
strb w11, [x15, x9]
add x9, x9, #1
ldrb w11, [x15, x9]
sub x11, x11, #1
strb w11, [x15, x9]
ldrb w11, [x15, x9]
cbnz w11, #0x148
ldrb w11, [x15, x9]
add x11, x11, #1
strb w11, [x15, x9]
ldrb w11, [x15, x9]
add x11, x11, #1
strb w11, [x15, x9]
ldrb w11, [x15, x9]
add x11, x11, #1
strb w11, [x15, x9]
ldrb w11, [x15, x9]
add x11, x11, #1
strb w11, [x15, x9]
ldrb w11, [x15, x9]
add x11, x11, #1
strb w11, [x15, x9]
ldrb w11, [x15, x9]
add x11, x11, #1
strb w11, [x15, x9]
ldrb w11, [x15, x9]
add x11, x11, #1
strb w11, [x15, x9]
ldrb w11, [x15, x9]
add x11, x11, #1
strb w11, [x15, x9]
ldrb w11, [x15, x9]
cbz w11, #0x254
add x9, x9, #1
ldrb w11, [x15, x9]
add x11, x11, #1
strb w11, [x15, x9]
ldrb w11, [x15, x9]
add x11, x11, #1
strb w11, [x15, x9]
ldrb w11, [x15, x9]
add x11, x11, #1
strb w11, [x15, x9]
ldrb w11, [x15, x9]
add x11, x11, #1
strb w11, [x15, x9]
ldrb w11, [x15, x9]
add x11, x11, #1
strb w11, [x15, x9]
ldrb w11, [x15, x9]
add x11, x11, #1
strb w11, [x15, x9]
ldrb w11, [x15, x9]
add x11, x11, #1
strb w11, [x15, x9]
ldrb w11, [x15, x9]
add x11, x11, #1
strb w11, [x15, x9]
sub x9, x9, #1
ldrb w11, [x15, x9]
sub x11, x11, #1
strb w11, [x15, x9]
ldrb w11, [x15, x9]
cbnz w11, #0x1d8
add x9, x9, #1
mov x0, #1
mov x1, x15
add x1, x1, x9
mov x2, #1
mov x16, #4
svc #0x80
ldrb w11, [x15, x9]
cbz w11, #0x28c
ldrb w11, [x15, x9]
sub x11, x11, #1
strb w11, [x15, x9]
ldrb w11, [x15, x9]
cbnz w11, #0x278
sub x9, x9, #1
sub x9, x9, #1
add x9, x9, #1
ldrb w11, [x15, x9]
add x11, x11, #1
strb w11, [x15, x9]
ldrb w11, [x15, x9]
add x11, x11, #1
strb w11, [x15, x9]
ldrb w11, [x15, x9]
add x11, x11, #1
strb w11, [x15, x9]
ldrb w11, [x15, x9]
add x11, x11, #1
strb w11, [x15, x9]
ldrb w11, [x15, x9]
add x11, x11, #1
strb w11, [x15, x9]
ldrb w11, [x15, x9]
add x11, x11, #1
strb w11, [x15, x9]
ldrb w11, [x15, x9]
add x11, x11, #1
strb w11, [x15, x9]
ldrb w11, [x15, x9]
add x11, x11, #1
strb w11, [x15, x9]
ldrb w11, [x15, x9]
add x11, x11, #1
strb w11, [x15, x9]
ldrb w11, [x15, x9]
add x11, x11, #1
strb w11, [x15, x9]
ldrb w11, [x15, x9]
cbz w11, #0x6d4
add x9, x9, #1
ldrb w11, [x15, x9]
add x11, x11, #1
strb w11, [x15, x9]
ldrb w11, [x15, x9]
add x11, x11, #1
strb w11, [x15, x9]
ldrb w11, [x15, x9]
add x11, x11, #1
strb w11, [x15, x9]
ldrb w11, [x15, x9]
add x11, x11, #1
strb w11, [x15, x9]
ldrb w11, [x15, x9]
add x11, x11, #1
strb w11, [x15, x9]
ldrb w11, [x15, x9]
add x11, x11, #1
strb w11, [x15, x9]
ldrb w11, [x15, x9]
add x11, x11, #1
strb w11, [x15, x9]
ldrb w11, [x15, x9]
add x11, x11, #1
strb w11, [x15, x9]
ldrb w11, [x15, x9]
add x11, x11, #1
strb w11, [x15, x9]
ldrb w11, [x15, x9]
add x11, x11, #1
strb w11, [x15, x9]
ldrb w11, [x15, x9]
cbz w11, #0x6bc
add x9, x9, #1
ldrb w11, [x15, x9]
add x11, x11, #1
strb w11, [x15, x9]
ldrb w11, [x15, x9]
add x11, x11, #1
strb w11, [x15, x9]
ldrb w11, [x15, x9]
add x11, x11, #1
strb w11, [x15, x9]
ldrb w11, [x15, x9]
add x11, x11, #1
strb w11, [x15, x9]
ldrb w11, [x15, x9]
add x11, x11, #1
strb w11, [x15, x9]
ldrb w11, [x15, x9]
add x11, x11, #1
strb w11, [x15, x9]
ldrb w11, [x15, x9]
add x11, x11, #1
strb w11, [x15, x9]
ldrb w11, [x15, x9]
add x11, x11, #1
strb w11, [x15, x9]
ldrb w11, [x15, x9]
add x11, x11, #1
strb w11, [x15, x9]
ldrb w11, [x15, x9]
add x11, x11, #1
strb w11, [x15, x9]
ldrb w11, [x15, x9]
cbz w11, #0x6a4
add x9, x9, #1
ldrb w11, [x15, x9]
add x11, x11, #1
strb w11, [x15, x9]
ldrb w11, [x15, x9]
add x11, x11, #1
strb w11, [x15, x9]
ldrb w11, [x15, x9]
add x11, x11, #1
strb w11, [x15, x9]
ldrb w11, [x15, x9]
add x11, x11, #1
strb w11, [x15, x9]
ldrb w11, [x15, x9]
add x11, x11, #1
strb w11, [x15, x9]
ldrb w11, [x15, x9]
add x11, x11, #1
strb w11, [x15, x9]
ldrb w11, [x15, x9]
add x11, x11, #1
strb w11, [x15, x9]
ldrb w11, [x15, x9]
add x11, x11, #1
strb w11, [x15, x9]
ldrb w11, [x15, x9]
add x11, x11, #1
strb w11, [x15, x9]
ldrb w11, [x15, x9]
add x11, x11, #1
strb w11, [x15, x9]
ldrb w11, [x15, x9]
cbz w11, #0x68c
add x9, x9, #1
ldrb w11, [x15, x9]
add x11, x11, #1
strb w11, [x15, x9]
ldrb w11, [x15, x9]
add x11, x11, #1
strb w11, [x15, x9]
ldrb w11, [x15, x9]
add x11, x11, #1
strb w11, [x15, x9]
ldrb w11, [x15, x9]
add x11, x11, #1
strb w11, [x15, x9]
ldrb w11, [x15, x9]
add x11, x11, #1
strb w11, [x15, x9]
ldrb w11, [x15, x9]
add x11, x11, #1
strb w11, [x15, x9]
ldrb w11, [x15, x9]
add x11, x11, #1
strb w11, [x15, x9]
ldrb w11, [x15, x9]
add x11, x11, #1
strb w11, [x15, x9]
ldrb w11, [x15, x9]
add x11, x11, #1
strb w11, [x15, x9]
ldrb w11, [x15, x9]
add x11, x11, #1
strb w11, [x15, x9]
ldrb w11, [x15, x9]
cbz w11, #0x674
add x9, x9, #1
ldrb w11, [x15, x9]
add x11, x11, #1
strb w11, [x15, x9]
ldrb w11, [x15, x9]
add x11, x11, #1
strb w11, [x15, x9]
ldrb w11, [x15, x9]
add x11, x11, #1
strb w11, [x15, x9]
ldrb w11, [x15, x9]
add x11, x11, #1
strb w11, [x15, x9]
ldrb w11, [x15, x9]
add x11, x11, #1
strb w11, [x15, x9]
ldrb w11, [x15, x9]
add x11, x11, #1
strb w11, [x15, x9]
ldrb w11, [x15, x9]
add x11, x11, #1
strb w11, [x15, x9]
ldrb w11, [x15, x9]
add x11, x11, #1
strb w11, [x15, x9]
ldrb w11, [x15, x9]
add x11, x11, #1
strb w11, [x15, x9]
ldrb w11, [x15, x9]
add x11, x11, #1
strb w11, [x15, x9]
ldrb w11, [x15, x9]
cbz w11, #0x65c
add x9, x9, #1
ldrb w11, [x15, x9]
add x11, x11, #1
strb w11, [x15, x9]
ldrb w11, [x15, x9]
add x11, x11, #1
strb w11, [x15, x9]
ldrb w11, [x15, x9]
add x11, x11, #1
strb w11, [x15, x9]
ldrb w11, [x15, x9]
add x11, x11, #1
strb w11, [x15, x9]
ldrb w11, [x15, x9]
add x11, x11, #1
strb w11, [x15, x9]
ldrb w11, [x15, x9]
add x11, x11, #1
strb w11, [x15, x9]
ldrb w11, [x15, x9]
add x11, x11, #1
strb w11, [x15, x9]
ldrb w11, [x15, x9]
add x11, x11, #1
strb w11, [x15, x9]
ldrb w11, [x15, x9]
add x11, x11, #1
strb w11, [x15, x9]
ldrb w11, [x15, x9]
add x11, x11, #1
strb w11, [x15, x9]
ldrb w11, [x15, x9]
cbz w11, #0x644
ldrb w11, [x15, x9]
sub x11, x11, #1
strb w11, [x15, x9]
ldrb w11, [x15, x9]
cbnz w11, #0x630
sub x9, x9, #1
ldrb w11, [x15, x9]
sub x11, x11, #1
strb w11, [x15, x9]
ldrb w11, [x15, x9]
cbnz w11, #0x5ac
sub x9, x9, #1
ldrb w11, [x15, x9]
sub x11, x11, #1
strb w11, [x15, x9]
ldrb w11, [x15, x9]
cbnz w11, #0x528
sub x9, x9, #1
ldrb w11, [x15, x9]
sub x11, x11, #1
strb w11, [x15, x9]
ldrb w11, [x15, x9]
cbnz w11, #0x4a4
sub x9, x9, #1
ldrb w11, [x15, x9]
sub x11, x11, #1
strb w11, [x15, x9]
ldrb w11, [x15, x9]
cbnz w11, #0x420
sub x9, x9, #1
ldrb w11, [x15, x9]
sub x11, x11, #1
strb w11, [x15, x9]
ldrb w11, [x15, x9]
cbnz w11, #0x39c
sub x9, x9, #1
ldrb w11, [x15, x9]
sub x11, x11, #1
strb w11, [x15, x9]
ldrb w11, [x15, x9]
cbnz w11, #0x318
sub x9, x9, #1
ldrb w11, [x15, x9]
sub x11, x11, #1
strb w11, [x15, x9]
ldrb w11, [x15, x9]
cbnz w11, #0xf8
ldrb w11, [x15, x9]
add x11, x11, #1
strb w11, [x15, x9]
ldrb w11, [x15, x9]
add x11, x11, #1
strb w11, [x15, x9]
ldrb w11, [x15, x9]
add x11, x11, #1
strb w11, [x15, x9]
ldrb w11, [x15, x9]
add x11, x11, #1
strb w11, [x15, x9]
ldrb w11, [x15, x9]
add x11, x11, #1
strb w11, [x15, x9]
ldrb w11, [x15, x9]
add x11, x11, #1
strb w11, [x15, x9]
ldrb w11, [x15, x9]
add x11, x11, #1
strb w11, [x15, x9]
ldrb w11, [x15, x9]
add x11, x11, #1
strb w11, [x15, x9]
ldrb w11, [x15, x9]
add x11, x11, #1
strb w11, [x15, x9]
ldrb w11, [x15, x9]
add x11, x11, #1
strb w11, [x15, x9]
mov x0, #1
mov x1, x15
add x1, x1, x9
mov x2, #1
mov x16, #4
svc #0x80
mov x0, x15
ret