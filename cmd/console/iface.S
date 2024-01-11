main.(*Cat).Quack STEXT size=143 args=0x8 locals=0x70 funcid=0x0 align=0x0
	0x0000 00000 (/Users/caozhengcheng/go/src/github.com/CarlCao17/go-gears/cmd/console/main.go:12)	TEXT	main.(*Cat).Quack(SB), ABIInternal, $112-8
	0x0000 00000 (/Users/caozhengcheng/go/src/github.com/CarlCao17/go-gears/cmd/console/main.go:12)	CMPQ	SP, 16(R14)
	0x0004 00004 (/Users/caozhengcheng/go/src/github.com/CarlCao17/go-gears/cmd/console/main.go:12)	PCDATA	$0, $-2
	0x0004 00004 (/Users/caozhengcheng/go/src/github.com/CarlCao17/go-gears/cmd/console/main.go:12)	JLS	121
	0x0006 00006 (/Users/caozhengcheng/go/src/github.com/CarlCao17/go-gears/cmd/console/main.go:12)	PCDATA	$0, $-1
	0x0006 00006 (/Users/caozhengcheng/go/src/github.com/CarlCao17/go-gears/cmd/console/main.go:12)	SUBQ	$112, SP
	0x000a 00010 (/Users/caozhengcheng/go/src/github.com/CarlCao17/go-gears/cmd/console/main.go:12)	MOVQ	BP, 104(SP)
	0x000f 00015 (/Users/caozhengcheng/go/src/github.com/CarlCao17/go-gears/cmd/console/main.go:12)	LEAQ	104(SP), BP
	0x0014 00020 (/Users/caozhengcheng/go/src/github.com/CarlCao17/go-gears/cmd/console/main.go:12)	FUNCDATA	$0, gclocals·ZzMiPAiVBg7DJ6dh/CjSag==(SB)
	0x0014 00020 (/Users/caozhengcheng/go/src/github.com/CarlCao17/go-gears/cmd/console/main.go:12)	FUNCDATA	$1, gclocals·Kw04D3lmrk4S2VQALtp7iQ==(SB)
	0x0014 00020 (/Users/caozhengcheng/go/src/github.com/CarlCao17/go-gears/cmd/console/main.go:12)	FUNCDATA	$5, main.(*Cat).Quack.arginfo1(SB)
	0x0014 00020 (/Users/caozhengcheng/go/src/github.com/CarlCao17/go-gears/cmd/console/main.go:12)	MOVQ	AX, main.c+120(SP)
	0x0019 00025 (/Users/caozhengcheng/go/src/github.com/CarlCao17/go-gears/cmd/console/main.go:13)	TESTB	AL, (AX)
	0x001b 00027 (/Users/caozhengcheng/go/src/github.com/CarlCao17/go-gears/cmd/console/main.go:13)	MOVQ	(AX), BX
	0x001e 00030 (/Users/caozhengcheng/go/src/github.com/CarlCao17/go-gears/cmd/console/main.go:13)	MOVQ	8(AX), CX
	0x0022 00034 (/Users/caozhengcheng/go/src/github.com/CarlCao17/go-gears/cmd/console/main.go:13)	MOVQ	BX, main..autotmp_3+72(SP)
	0x0027 00039 (/Users/caozhengcheng/go/src/github.com/CarlCao17/go-gears/cmd/console/main.go:13)	MOVQ	CX, main..autotmp_3+80(SP)
	0x002c 00044 (/Users/caozhengcheng/go/src/github.com/CarlCao17/go-gears/cmd/console/main.go:13)	LEAQ	main..autotmp_2+40(SP), AX
	0x0031 00049 (/Users/caozhengcheng/go/src/github.com/CarlCao17/go-gears/cmd/console/main.go:13)	LEAQ	go:string." meow"(SB), DI
	0x0038 00056 (/Users/caozhengcheng/go/src/github.com/CarlCao17/go-gears/cmd/console/main.go:13)	MOVL	$5, SI
	0x003d 00061 (/Users/caozhengcheng/go/src/github.com/CarlCao17/go-gears/cmd/console/main.go:13)	PCDATA	$1, $1
	0x003d 00061 (/Users/caozhengcheng/go/src/github.com/CarlCao17/go-gears/cmd/console/main.go:13)	NOP
	0x0040 00064 (/Users/caozhengcheng/go/src/github.com/CarlCao17/go-gears/cmd/console/main.go:13)	CALL	runtime.concatstring2(SB)
	0x0045 00069 (/Users/caozhengcheng/go/src/github.com/CarlCao17/go-gears/cmd/console/main.go:13)	MOVQ	AX, main..autotmp_1+88(SP)
	0x004a 00074 (/Users/caozhengcheng/go/src/github.com/CarlCao17/go-gears/cmd/console/main.go:13)	MOVQ	BX, main..autotmp_1+96(SP)
	0x004f 00079 (/Users/caozhengcheng/go/src/github.com/CarlCao17/go-gears/cmd/console/main.go:13)	PCDATA	$1, $2
	0x004f 00079 (/Users/caozhengcheng/go/src/github.com/CarlCao17/go-gears/cmd/console/main.go:13)	CALL	runtime.printlock(SB)
	0x0054 00084 (/Users/caozhengcheng/go/src/github.com/CarlCao17/go-gears/cmd/console/main.go:13)	MOVQ	main..autotmp_1+88(SP), AX
	0x0059 00089 (/Users/caozhengcheng/go/src/github.com/CarlCao17/go-gears/cmd/console/main.go:13)	MOVQ	main..autotmp_1+96(SP), BX
	0x005e 00094 (/Users/caozhengcheng/go/src/github.com/CarlCao17/go-gears/cmd/console/main.go:13)	PCDATA	$1, $1
	0x005e 00094 (/Users/caozhengcheng/go/src/github.com/CarlCao17/go-gears/cmd/console/main.go:13)	NOP
	0x0060 00096 (/Users/caozhengcheng/go/src/github.com/CarlCao17/go-gears/cmd/console/main.go:13)	CALL	runtime.printstring(SB)
	0x0065 00101 (/Users/caozhengcheng/go/src/github.com/CarlCao17/go-gears/cmd/console/main.go:13)	CALL	runtime.printnl(SB)
	0x006a 00106 (/Users/caozhengcheng/go/src/github.com/CarlCao17/go-gears/cmd/console/main.go:13)	CALL	runtime.printunlock(SB)
	0x006f 00111 (/Users/caozhengcheng/go/src/github.com/CarlCao17/go-gears/cmd/console/main.go:14)	MOVQ	104(SP), BP
	0x0074 00116 (/Users/caozhengcheng/go/src/github.com/CarlCao17/go-gears/cmd/console/main.go:14)	ADDQ	$112, SP
	0x0078 00120 (/Users/caozhengcheng/go/src/github.com/CarlCao17/go-gears/cmd/console/main.go:14)	RET
	0x0079 00121 (/Users/caozhengcheng/go/src/github.com/CarlCao17/go-gears/cmd/console/main.go:14)	NOP
	0x0079 00121 (/Users/caozhengcheng/go/src/github.com/CarlCao17/go-gears/cmd/console/main.go:12)	PCDATA	$1, $-1
	0x0079 00121 (/Users/caozhengcheng/go/src/github.com/CarlCao17/go-gears/cmd/console/main.go:12)	PCDATA	$0, $-2
	0x0079 00121 (/Users/caozhengcheng/go/src/github.com/CarlCao17/go-gears/cmd/console/main.go:12)	MOVQ	AX, 8(SP)
	0x007e 00126 (/Users/caozhengcheng/go/src/github.com/CarlCao17/go-gears/cmd/console/main.go:12)	NOP
	0x0080 00128 (/Users/caozhengcheng/go/src/github.com/CarlCao17/go-gears/cmd/console/main.go:12)	CALL	runtime.morestack_noctxt(SB)
	0x0085 00133 (/Users/caozhengcheng/go/src/github.com/CarlCao17/go-gears/cmd/console/main.go:12)	MOVQ	8(SP), AX
	0x008a 00138 (/Users/caozhengcheng/go/src/github.com/CarlCao17/go-gears/cmd/console/main.go:12)	PCDATA	$0, $-1
	0x008a 00138 (/Users/caozhengcheng/go/src/github.com/CarlCao17/go-gears/cmd/console/main.go:12)	JMP	0
	0x0000 49 3b 66 10 76 73 48 83 ec 70 48 89 6c 24 68 48  I;f.vsH..pH.l$hH
	0x0010 8d 6c 24 68 48 89 44 24 78 84 00 48 8b 18 48 8b  .l$hH.D$x..H..H.
	0x0020 48 08 48 89 5c 24 48 48 89 4c 24 50 48 8d 44 24  H.H.\$HH.L$PH.D$
	0x0030 28 48 8d 3d 00 00 00 00 be 05 00 00 00 0f 1f 00  (H.=............
	0x0040 e8 00 00 00 00 48 89 44 24 58 48 89 5c 24 60 e8  .....H.D$XH.\$`.
	0x0050 00 00 00 00 48 8b 44 24 58 48 8b 5c 24 60 66 90  ....H.D$XH.\$`f.
	0x0060 e8 00 00 00 00 e8 00 00 00 00 e8 00 00 00 00 48  ...............H
	0x0070 8b 6c 24 68 48 83 c4 70 c3 48 89 44 24 08 66 90  .l$hH..p.H.D$.f.
	0x0080 e8 00 00 00 00 48 8b 44 24 08 e9 71 ff ff ff     .....H.D$..q...
	rel 52+4 t=14 go:string." meow"+0
	rel 65+4 t=7 runtime.concatstring2+0
	rel 80+4 t=7 runtime.printlock+0
	rel 97+4 t=7 runtime.printstring+0
	rel 102+4 t=7 runtime.printnl+0
	rel 107+4 t=7 runtime.printunlock+0
	rel 129+4 t=7 runtime.morestack_noctxt+0
main.main STEXT size=110 args=0x0 locals=0x50 funcid=0x0 align=0x0
	0x0000 00000 (/Users/caozhengcheng/go/src/github.com/CarlCao17/go-gears/cmd/console/main.go:16)	TEXT	main.main(SB), ABIInternal, $80-0
	0x0000 00000 (/Users/caozhengcheng/go/src/github.com/CarlCao17/go-gears/cmd/console/main.go:16)	CMPQ	SP, 16(R14)
	0x0004 00004 (/Users/caozhengcheng/go/src/github.com/CarlCao17/go-gears/cmd/console/main.go:16)	PCDATA	$0, $-2
	0x0004 00004 (/Users/caozhengcheng/go/src/github.com/CarlCao17/go-gears/cmd/console/main.go:16)	JLS	103
	0x0006 00006 (/Users/caozhengcheng/go/src/github.com/CarlCao17/go-gears/cmd/console/main.go:16)	PCDATA	$0, $-1
	0x0006 00006 (/Users/caozhengcheng/go/src/github.com/CarlCao17/go-gears/cmd/console/main.go:16)	SUBQ	$80, SP
	0x000a 00010 (/Users/caozhengcheng/go/src/github.com/CarlCao17/go-gears/cmd/console/main.go:16)	MOVQ	BP, 72(SP)
	0x000f 00015 (/Users/caozhengcheng/go/src/github.com/CarlCao17/go-gears/cmd/console/main.go:16)	LEAQ	72(SP), BP
	0x0014 00020 (/Users/caozhengcheng/go/src/github.com/CarlCao17/go-gears/cmd/console/main.go:16)	FUNCDATA	$0, gclocals·g2BeySu+wFnoycgXfElmcg==(SB)
	0x0014 00020 (/Users/caozhengcheng/go/src/github.com/CarlCao17/go-gears/cmd/console/main.go:16)	FUNCDATA	$1, gclocals·/ydTHfVJHvKeH/UP4dRKSQ==(SB)
	0x0014 00020 (/Users/caozhengcheng/go/src/github.com/CarlCao17/go-gears/cmd/console/main.go:16)	FUNCDATA	$2, main.main.stkobj(SB)
	0x0014 00020 (/Users/caozhengcheng/go/src/github.com/CarlCao17/go-gears/cmd/console/main.go:17)	MOVUPS	X15, main..autotmp_3+56(SP)
	0x001a 00026 (/Users/caozhengcheng/go/src/github.com/CarlCao17/go-gears/cmd/console/main.go:17)	LEAQ	main..autotmp_3+56(SP), DX
	0x001f 00031 (/Users/caozhengcheng/go/src/github.com/CarlCao17/go-gears/cmd/console/main.go:17)	MOVQ	DX, main..autotmp_2+24(SP)
	0x0024 00036 (/Users/caozhengcheng/go/src/github.com/CarlCao17/go-gears/cmd/console/main.go:17)	TESTB	AL, (DX)
	0x0026 00038 (/Users/caozhengcheng/go/src/github.com/CarlCao17/go-gears/cmd/console/main.go:17)	MOVQ	$6, main..autotmp_3+64(SP)
	0x002f 00047 (/Users/caozhengcheng/go/src/github.com/CarlCao17/go-gears/cmd/console/main.go:17)	LEAQ	go:string."draven"(SB), SI
	0x0036 00054 (/Users/caozhengcheng/go/src/github.com/CarlCao17/go-gears/cmd/console/main.go:17)	MOVQ	SI, main..autotmp_3+56(SP)
	0x003b 00059 (/Users/caozhengcheng/go/src/github.com/CarlCao17/go-gears/cmd/console/main.go:17)	MOVQ	DX, main..autotmp_1+32(SP)
	0x0040 00064 (/Users/caozhengcheng/go/src/github.com/CarlCao17/go-gears/cmd/console/main.go:17)	LEAQ	go:itab.*<unlinkable>.Cat,<unlinkable>.Duck(SB), AX
	0x0047 00071 (/Users/caozhengcheng/go/src/github.com/CarlCao17/go-gears/cmd/console/main.go:17)	MOVQ	AX, main.c+40(SP)
	0x004c 00076 (/Users/caozhengcheng/go/src/github.com/CarlCao17/go-gears/cmd/console/main.go:17)	MOVQ	DX, main.c+48(SP)
	0x0051 00081 (/Users/caozhengcheng/go/src/github.com/CarlCao17/go-gears/cmd/console/main.go:18)	JMP	83
	0x0053 00083 (/Users/caozhengcheng/go/src/github.com/CarlCao17/go-gears/cmd/console/main.go:18)	LEAQ	type:*main.Cat(SB), BX
	0x005a 00090 (/Users/caozhengcheng/go/src/github.com/CarlCao17/go-gears/cmd/console/main.go:18)	LEAQ	type:main.Duck(SB), CX
	0x0061 00097 (/Users/caozhengcheng/go/src/github.com/CarlCao17/go-gears/cmd/console/main.go:18)	PCDATA	$1, $0
	0x0061 00097 (/Users/caozhengcheng/go/src/github.com/CarlCao17/go-gears/cmd/console/main.go:18)	CALL	runtime.panicdottypeI(SB)
	0x0066 00102 (/Users/caozhengcheng/go/src/github.com/CarlCao17/go-gears/cmd/console/main.go:18)	XCHGL	AX, AX
	0x0067 00103 (/Users/caozhengcheng/go/src/github.com/CarlCao17/go-gears/cmd/console/main.go:18)	NOP
	0x0067 00103 (/Users/caozhengcheng/go/src/github.com/CarlCao17/go-gears/cmd/console/main.go:16)	PCDATA	$1, $-1
	0x0067 00103 (/Users/caozhengcheng/go/src/github.com/CarlCao17/go-gears/cmd/console/main.go:16)	PCDATA	$0, $-2
	0x0067 00103 (/Users/caozhengcheng/go/src/github.com/CarlCao17/go-gears/cmd/console/main.go:16)	CALL	runtime.morestack_noctxt(SB)
	0x006c 00108 (/Users/caozhengcheng/go/src/github.com/CarlCao17/go-gears/cmd/console/main.go:16)	PCDATA	$0, $-1
	0x006c 00108 (/Users/caozhengcheng/go/src/github.com/CarlCao17/go-gears/cmd/console/main.go:16)	JMP	0
	0x0000 49 3b 66 10 76 61 48 83 ec 50 48 89 6c 24 48 48  I;f.vaH..PH.l$HH
	0x0010 8d 6c 24 48 44 0f 11 7c 24 38 48 8d 54 24 38 48  .l$HD..|$8H.T$8H
	0x0020 89 54 24 18 84 02 48 c7 44 24 40 06 00 00 00 48  .T$...H.D$@....H
	0x0030 8d 35 00 00 00 00 48 89 74 24 38 48 89 54 24 20  .5....H.t$8H.T$ 
	0x0040 48 8d 05 00 00 00 00 48 89 44 24 28 48 89 54 24  H......H.D$(H.T$
	0x0050 30 eb 00 48 8d 1d 00 00 00 00 48 8d 0d 00 00 00  0..H......H.....
	0x0060 00 e8 00 00 00 00 90 e8 00 00 00 00 eb 92        ..............
	rel 2+0 t=23 type:*main.Cat+0
	rel 50+4 t=14 go:string."draven"+0
	rel 67+4 t=14 go:itab.*<unlinkable>.Cat,<unlinkable>.Duck+0
	rel 86+4 t=14 type:*main.Cat+0
	rel 93+4 t=14 type:main.Duck+0
	rel 98+4 t=7 runtime.panicdottypeI+0
	rel 104+4 t=7 runtime.morestack_noctxt+0
main.Duck.Quack STEXT dupok size=108 args=0x10 locals=0x10 funcid=0x15 align=0x0
	0x0000 00000 (<autogenerated>:1)	TEXT	main.Duck.Quack(SB), DUPOK|WRAPPER|ABIInternal, $16-16
	0x0000 00000 (<autogenerated>:1)	CMPQ	SP, 16(R14)
	0x0004 00004 (<autogenerated>:1)	PCDATA	$0, $-2
	0x0004 00004 (<autogenerated>:1)	JLS	60
	0x0006 00006 (<autogenerated>:1)	PCDATA	$0, $-1
	0x0006 00006 (<autogenerated>:1)	SUBQ	$16, SP
	0x000a 00010 (<autogenerated>:1)	MOVQ	BP, 8(SP)
	0x000f 00015 (<autogenerated>:1)	LEAQ	8(SP), BP
	0x0014 00020 (<autogenerated>:1)	MOVQ	32(R14), R12
	0x0018 00024 (<autogenerated>:1)	TESTQ	R12, R12
	0x001b 00027 (<autogenerated>:1)	JNE	87
	0x001d 00029 (<autogenerated>:1)	NOP
	0x001d 00029 (<autogenerated>:1)	FUNCDATA	$0, gclocals·IuErl7MOXaHVn7EZYWzfFA==(SB)
	0x001d 00029 (<autogenerated>:1)	FUNCDATA	$1, gclocals·J5F+7Qw7O7ve2QcWC7DpeQ==(SB)
	0x001d 00029 (<autogenerated>:1)	FUNCDATA	$5, main.Duck.Quack.arginfo1(SB)
	0x001d 00029 (<autogenerated>:1)	MOVQ	AX, main..this+24(SP)
	0x0022 00034 (<autogenerated>:1)	MOVQ	BX, main..this+32(SP)
	0x0027 00039 (<autogenerated>:1)	TESTB	AL, (AX)
	0x0029 00041 (<autogenerated>:1)	MOVQ	24(AX), CX
	0x002d 00045 (<autogenerated>:1)	MOVQ	BX, AX
	0x0030 00048 (<autogenerated>:1)	PCDATA	$1, $1
	0x0030 00048 (<autogenerated>:1)	CALL	CX
	0x0032 00050 (<autogenerated>:1)	MOVQ	8(SP), BP
	0x0037 00055 (<autogenerated>:1)	ADDQ	$16, SP
	0x003b 00059 (<autogenerated>:1)	RET
	0x003c 00060 (<autogenerated>:1)	NOP
	0x003c 00060 (<autogenerated>:1)	PCDATA	$1, $-1
	0x003c 00060 (<autogenerated>:1)	PCDATA	$0, $-2
	0x003c 00060 (<autogenerated>:1)	MOVQ	AX, 8(SP)
	0x0041 00065 (<autogenerated>:1)	MOVQ	BX, 16(SP)
	0x0046 00070 (<autogenerated>:1)	CALL	runtime.morestack_noctxt(SB)
	0x004b 00075 (<autogenerated>:1)	MOVQ	8(SP), AX
	0x0050 00080 (<autogenerated>:1)	MOVQ	16(SP), BX
	0x0055 00085 (<autogenerated>:1)	PCDATA	$0, $-1
	0x0055 00085 (<autogenerated>:1)	JMP	0
	0x0057 00087 (<autogenerated>:1)	LEAQ	24(SP), R13
	0x005c 00092 (<autogenerated>:1)	NOP
	0x0060 00096 (<autogenerated>:1)	CMPQ	(R12), R13
	0x0064 00100 (<autogenerated>:1)	JNE	29
	0x0066 00102 (<autogenerated>:1)	MOVQ	SP, (R12)
	0x006a 00106 (<autogenerated>:1)	JMP	29
	0x0000 49 3b 66 10 76 36 48 83 ec 10 48 89 6c 24 08 48  I;f.v6H...H.l$.H
	0x0010 8d 6c 24 08 4d 8b 66 20 4d 85 e4 75 3a 48 89 44  .l$.M.f M..u:H.D
	0x0020 24 18 48 89 5c 24 20 84 00 48 8b 48 18 48 89 d8  $.H.\$ ..H.H.H..
	0x0030 ff d1 48 8b 6c 24 08 48 83 c4 10 c3 48 89 44 24  ..H.l$.H....H.D$
	0x0040 08 48 89 5c 24 10 e8 00 00 00 00 48 8b 44 24 08  .H.\$......H.D$.
	0x0050 48 8b 5c 24 10 eb a9 4c 8d 6c 24 18 0f 1f 40 00  H.\$...L.l$...@.
	0x0060 4d 39 2c 24 75 b7 49 89 24 24 eb b1              M9,$u.I.$$..
	rel 2+0 t=24 type:main.Duck+96
	rel 48+0 t=10 +0
	rel 71+4 t=7 runtime.morestack_noctxt+0
go:cuinfo.producer.<unlinkable> SDWARFCUINFO dupok size=0
	0x0000 2d 4e 20 2d 6c 20 72 65 67 61 62 69              -N -l regabi
runtime.memequal64·f SRODATA dupok size=8
	0x0000 00 00 00 00 00 00 00 00                          ........
	rel 0+8 t=1 runtime.memequal64+0
runtime.gcbits.0100000000000000 SRODATA dupok size=8
	0x0000 01 00 00 00 00 00 00 00                          ........
type:.namedata.*func()- SRODATA dupok size=9
	0x0000 00 07 2a 66 75 6e 63 28 29                       ..*func()
type:*func() SRODATA dupok size=56
	0x0000 08 00 00 00 00 00 00 00 08 00 00 00 00 00 00 00  ................
	0x0010 75 ac 29 27 08 08 08 36 00 00 00 00 00 00 00 00  u.)'...6........
	0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0030 00 00 00 00 00 00 00 00                          ........
	rel 24+8 t=1 runtime.memequal64·f+0
	rel 32+8 t=1 runtime.gcbits.0100000000000000+0
	rel 40+4 t=5 type:.namedata.*func()-+0
	rel 48+8 t=1 type:func()+0
type:func() SRODATA dupok size=56
	0x0000 08 00 00 00 00 00 00 00 08 00 00 00 00 00 00 00  ................
	0x0010 fe fa b9 80 02 08 08 33 00 00 00 00 00 00 00 00  .......3........
	0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0030 00 00 00 00                                      ....
	rel 32+8 t=1 runtime.gcbits.0100000000000000+0
	rel 40+4 t=5 type:.namedata.*func()-+0
	rel 44+4 t=-32763 type:*func()+0
runtime.interequal·f SRODATA dupok size=8
	0x0000 00 00 00 00 00 00 00 00                          ........
	rel 0+8 t=1 runtime.interequal+0
type:.namedata.*main.Duck. SRODATA dupok size=12
	0x0000 01 0a 2a 6d 61 69 6e 2e 44 75 63 6b              ..*main.Duck
type:*<unlinkable>.Duck SRODATA size=56
	0x0000 08 00 00 00 00 00 00 00 08 00 00 00 00 00 00 00  ................
	0x0010 00 e7 d4 d2 08 08 08 36 00 00 00 00 00 00 00 00  .......6........
	0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0030 00 00 00 00 00 00 00 00                          ........
	rel 24+8 t=1 runtime.memequal64·f+0
	rel 32+8 t=1 runtime.gcbits.0100000000000000+0
	rel 40+4 t=5 type:.namedata.*main.Duck.+0
	rel 48+8 t=1 type:<unlinkable>.Duck+0
runtime.gcbits.0200000000000000 SRODATA dupok size=8
	0x0000 02 00 00 00 00 00 00 00                          ........
type:.importpath.<unlinkable>. SRODATA dupok size=14
	0x0000 00 0c 3c 75 6e 6c 69 6e 6b 61 62 6c 65 3e        ..<unlinkable>
type:.namedata.Quack. SRODATA dupok size=7
	0x0000 01 05 51 75 61 63 6b                             ..Quack
type:<unlinkable>.Duck SRODATA size=104
	0x0000 10 00 00 00 00 00 00 00 10 00 00 00 00 00 00 00  ................
	0x0010 31 33 89 84 07 08 08 14 00 00 00 00 00 00 00 00  13..............
	0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0030 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0040 01 00 00 00 00 00 00 00 01 00 00 00 00 00 00 00  ................
	0x0050 00 00 00 00 00 00 00 00 18 00 00 00 00 00 00 00  ................
	0x0060 00 00 00 00 00 00 00 00                          ........
	rel 24+8 t=1 runtime.interequal·f+0
	rel 32+8 t=1 runtime.gcbits.0200000000000000+0
	rel 40+4 t=5 type:.namedata.*main.Duck.+0
	rel 44+4 t=5 type:*<unlinkable>.Duck+0
	rel 48+8 t=1 type:.importpath.<unlinkable>.+0
	rel 56+8 t=1 type:<unlinkable>.Duck+96
	rel 80+4 t=5 type:.importpath.<unlinkable>.+0
	rel 96+4 t=5 type:.namedata.Quack.+0
	rel 100+4 t=5 type:func()+0
runtime.strequal·f SRODATA dupok size=8
	0x0000 00 00 00 00 00 00 00 00                          ........
	rel 0+8 t=1 runtime.strequal+0
type:.namedata.*main.Cat. SRODATA dupok size=11
	0x0000 01 09 2a 6d 61 69 6e 2e 43 61 74                 ..*main.Cat
type:.namedata.Name. SRODATA dupok size=6
	0x0000 01 04 4e 61 6d 65                                ..Name
type:<unlinkable>.Cat SRODATA size=120
	0x0000 10 00 00 00 00 00 00 00 08 00 00 00 00 00 00 00  ................
	0x0010 5e 02 d0 de 07 08 08 19 00 00 00 00 00 00 00 00  ^...............
	0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0030 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0040 01 00 00 00 00 00 00 00 01 00 00 00 00 00 00 00  ................
	0x0050 00 00 00 00 00 00 00 00 28 00 00 00 00 00 00 00  ........(.......
	0x0060 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0070 00 00 00 00 00 00 00 00                          ........
	rel 24+8 t=1 runtime.strequal·f+0
	rel 32+8 t=1 runtime.gcbits.0100000000000000+0
	rel 40+4 t=5 type:.namedata.*main.Cat.+0
	rel 44+4 t=5 type:*<unlinkable>.Cat+0
	rel 56+8 t=1 type:<unlinkable>.Cat+96
	rel 80+4 t=5 type:.importpath.<unlinkable>.+0
	rel 96+8 t=1 type:.namedata.Name.+0
	rel 104+8 t=1 type:string+0
type:.namedata.*func(*main.Cat)- SRODATA dupok size=18
	0x0000 00 10 2a 66 75 6e 63 28 2a 6d 61 69 6e 2e 43 61  ..*func(*main.Ca
	0x0010 74 29                                            t)
type:*func(*<unlinkable>.Cat) SRODATA dupok size=56
	0x0000 08 00 00 00 00 00 00 00 08 00 00 00 00 00 00 00  ................
	0x0010 28 fa eb 06 08 08 08 36 00 00 00 00 00 00 00 00  (......6........
	0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0030 00 00 00 00 00 00 00 00                          ........
	rel 24+8 t=1 runtime.memequal64·f+0
	rel 32+8 t=1 runtime.gcbits.0100000000000000+0
	rel 40+4 t=5 type:.namedata.*func(*main.Cat)-+0
	rel 48+8 t=1 type:func(*<unlinkable>.Cat)+0
type:func(*<unlinkable>.Cat) SRODATA dupok size=64
	0x0000 08 00 00 00 00 00 00 00 08 00 00 00 00 00 00 00  ................
	0x0010 fd 6c c5 67 02 08 08 33 00 00 00 00 00 00 00 00  .l.g...3........
	0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0030 01 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	rel 32+8 t=1 runtime.gcbits.0100000000000000+0
	rel 40+4 t=5 type:.namedata.*func(*main.Cat)-+0
	rel 44+4 t=-32763 type:*func(*<unlinkable>.Cat)+0
	rel 56+8 t=1 type:*<unlinkable>.Cat+0
type:*<unlinkable>.Cat SRODATA size=88
	0x0000 08 00 00 00 00 00 00 00 08 00 00 00 00 00 00 00  ................
	0x0010 63 c9 44 17 09 08 08 36 00 00 00 00 00 00 00 00  c.D....6........
	0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0030 00 00 00 00 00 00 00 00 00 00 00 00 01 00 01 00  ................
	0x0040 10 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0050 00 00 00 00 00 00 00 00                          ........
	rel 24+8 t=1 runtime.memequal64·f+0
	rel 32+8 t=1 runtime.gcbits.0100000000000000+0
	rel 40+4 t=5 type:.namedata.*main.Cat.+0
	rel 48+8 t=1 type:<unlinkable>.Cat+0
	rel 56+4 t=5 type:.importpath.<unlinkable>.+0
	rel 72+4 t=5 type:.namedata.Quack.+0
	rel 76+4 t=26 type:func()+0
	rel 80+4 t=26 <unlinkable>.(*Cat).Quack+0
	rel 84+4 t=26 <unlinkable>.(*Cat).Quack+0
go:itab.*<unlinkable>.Cat,<unlinkable>.Duck SRODATA dupok size=32
	0x0000 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0010 63 c9 44 17 00 00 00 00 00 00 00 00 00 00 00 00  c.D.............
	rel 0+8 t=1 type:<unlinkable>.Duck+0
	rel 8+8 t=1 type:*<unlinkable>.Cat+0
	rel 24+8 t=-32767 <unlinkable>.(*Cat).Quack+0
go:cuinfo.packagename.main SDWARFCUINFO dupok size=0
	0x0000 6d 61 69 6e                                      main
main..inittask SNOPTRDATA size=24
	0x0000 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0010 00 00 00 00 00 00 00 00                          ........
go:string." meow" SRODATA dupok size=5
	0x0000 20 6d 65 6f 77                                    meow
go:string."draven" SRODATA dupok size=6
	0x0000 64 72 61 76 65 6e                                draven
type:*main.Duck SRODATA size=56
	0x0000 08 00 00 00 00 00 00 00 08 00 00 00 00 00 00 00  ................
	0x0010 8c 88 4d 63 08 08 08 36 00 00 00 00 00 00 00 00  ..Mc...6........
	0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0030 00 00 00 00 00 00 00 00                          ........
	rel 24+8 t=1 runtime.memequal64·f+0
	rel 32+8 t=1 runtime.gcbits.0100000000000000+0
	rel 40+4 t=5 type:.namedata.*main.Duck.+0
	rel 48+8 t=1 type:main.Duck+0
type:main.Duck SRODATA size=104
	0x0000 10 00 00 00 00 00 00 00 10 00 00 00 00 00 00 00  ................
	0x0010 5d e2 e8 d7 07 08 08 14 00 00 00 00 00 00 00 00  ]...............
	0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0030 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0040 01 00 00 00 00 00 00 00 01 00 00 00 00 00 00 00  ................
	0x0050 00 00 00 00 00 00 00 00 18 00 00 00 00 00 00 00  ................
	0x0060 00 00 00 00 00 00 00 00                          ........
	rel 24+8 t=1 runtime.interequal·f+0
	rel 32+8 t=1 runtime.gcbits.0200000000000000+0
	rel 40+4 t=5 type:.namedata.*main.Duck.+0
	rel 44+4 t=5 type:*main.Duck+0
	rel 48+8 t=1 type:.importpath.<unlinkable>.+0
	rel 56+8 t=1 type:main.Duck+96
	rel 80+4 t=5 type:.importpath.<unlinkable>.+0
	rel 96+4 t=5 type:.namedata.Quack.+0
	rel 100+4 t=5 type:func()+0
type:main.Cat SRODATA size=120
	0x0000 10 00 00 00 00 00 00 00 08 00 00 00 00 00 00 00  ................
	0x0010 85 bd ce 45 07 08 08 19 00 00 00 00 00 00 00 00  ...E............
	0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0030 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0040 01 00 00 00 00 00 00 00 01 00 00 00 00 00 00 00  ................
	0x0050 00 00 00 00 00 00 00 00 28 00 00 00 00 00 00 00  ........(.......
	0x0060 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0070 00 00 00 00 00 00 00 00                          ........
	rel 24+8 t=1 runtime.strequal·f+0
	rel 32+8 t=1 runtime.gcbits.0100000000000000+0
	rel 40+4 t=5 type:.namedata.*main.Cat.+0
	rel 44+4 t=5 type:*main.Cat+0
	rel 56+8 t=1 type:main.Cat+96
	rel 80+4 t=5 type:.importpath.<unlinkable>.+0
	rel 96+8 t=1 type:.namedata.Name.+0
	rel 104+8 t=1 type:string+0
type:*func(*main.Cat) SRODATA dupok size=56
	0x0000 08 00 00 00 00 00 00 00 08 00 00 00 00 00 00 00  ................
	0x0010 f6 f1 7f da 08 08 08 36 00 00 00 00 00 00 00 00  .......6........
	0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0030 00 00 00 00 00 00 00 00                          ........
	rel 24+8 t=1 runtime.memequal64·f+0
	rel 32+8 t=1 runtime.gcbits.0100000000000000+0
	rel 40+4 t=5 type:.namedata.*func(*main.Cat)-+0
	rel 48+8 t=1 type:func(*main.Cat)+0
type:func(*main.Cat) SRODATA dupok size=64
	0x0000 08 00 00 00 00 00 00 00 08 00 00 00 00 00 00 00  ................
	0x0010 db 2d 6a 47 02 08 08 33 00 00 00 00 00 00 00 00  .-jG...3........
	0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0030 01 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	rel 32+8 t=1 runtime.gcbits.0100000000000000+0
	rel 40+4 t=5 type:.namedata.*func(*main.Cat)-+0
	rel 44+4 t=-32763 type:*func(*main.Cat)+0
	rel 56+8 t=1 type:*main.Cat+0
type:*main.Cat SRODATA size=88
	0x0000 08 00 00 00 00 00 00 00 08 00 00 00 00 00 00 00  ................
	0x0010 4a ae 86 a6 09 08 08 36 00 00 00 00 00 00 00 00  J......6........
	0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0030 00 00 00 00 00 00 00 00 00 00 00 00 01 00 01 00  ................
	0x0040 10 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0050 00 00 00 00 00 00 00 00                          ........
	rel 24+8 t=1 runtime.memequal64·f+0
	rel 32+8 t=1 runtime.gcbits.0100000000000000+0
	rel 40+4 t=5 type:.namedata.*main.Cat.+0
	rel 48+8 t=1 type:main.Cat+0
	rel 56+4 t=5 type:.importpath.<unlinkable>.+0
	rel 72+4 t=5 type:.namedata.Quack.+0
	rel 76+4 t=26 type:func()+0
	rel 80+4 t=26 main.(*Cat).Quack+0
	rel 84+4 t=26 main.(*Cat).Quack+0
go:itab.*main.Cat,main.Duck SRODATA dupok size=32
	0x0000 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0010 4a ae 86 a6 00 00 00 00 00 00 00 00 00 00 00 00  J...............
	rel 0+8 t=1 type:main.Duck+0
	rel 8+8 t=1 type:*main.Cat+0
	rel 24+8 t=-32767 main.(*Cat).Quack+0
type:.eqfunc32 SRODATA dupok size=16
	0x0000 00 00 00 00 00 00 00 00 20 00 00 00 00 00 00 00  ........ .......
	rel 0+8 t=1 runtime.memequal_varlen+0
type:.namedata.*[32]uint8- SRODATA dupok size=12
	0x0000 00 0a 2a 5b 33 32 5d 75 69 6e 74 38              ..*[32]uint8
type:*[32]uint8 SRODATA dupok size=56
	0x0000 08 00 00 00 00 00 00 00 08 00 00 00 00 00 00 00  ................
	0x0010 26 0d b3 9e 08 08 08 36 00 00 00 00 00 00 00 00  &......6........
	0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0030 00 00 00 00 00 00 00 00                          ........
	rel 24+8 t=1 runtime.memequal64·f+0
	rel 32+8 t=1 runtime.gcbits.0100000000000000+0
	rel 40+4 t=5 type:.namedata.*[32]uint8-+0
	rel 48+8 t=1 type:[32]uint8+0
runtime.gcbits. SRODATA dupok size=0
type:[32]uint8 SRODATA dupok size=72
	0x0000 20 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00   ...............
	0x0010 aa ab 93 92 0a 01 01 11 00 00 00 00 00 00 00 00  ................
	0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0030 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0040 20 00 00 00 00 00 00 00                           .......
	rel 24+8 t=1 type:.eqfunc32+0
	rel 32+8 t=1 runtime.gcbits.+0
	rel 40+4 t=5 type:.namedata.*[32]uint8-+0
	rel 44+4 t=-32763 type:*[32]uint8+0
	rel 48+8 t=1 type:uint8+0
	rel 56+8 t=1 type:[]uint8+0
gclocals·ZzMiPAiVBg7DJ6dh/CjSag== SRODATA dupok size=11
	0x0000 03 00 00 00 01 00 00 00 01 00 00                 ...........
gclocals·Kw04D3lmrk4S2VQALtp7iQ== SRODATA dupok size=11
	0x0000 03 00 00 00 04 00 00 00 00 00 04                 ...........
main.(*Cat).Quack.arginfo1 SRODATA static dupok size=3
	0x0000 00 08 ff                                         ...
gclocals·g2BeySu+wFnoycgXfElmcg== SRODATA dupok size=8
	0x0000 01 00 00 00 00 00 00 00                          ........
gclocals·/ydTHfVJHvKeH/UP4dRKSQ== SRODATA dupok size=9
	0x0000 01 00 00 00 06 00 00 00 00                       .........
main.main.stkobj SRODATA static size=24
	0x0000 01 00 00 00 00 00 00 00 f0 ff ff ff 10 00 00 00  ................
	0x0010 08 00 00 00 00 00 00 00                          ........
	rel 20+4 t=5 runtime.gcbits.0100000000000000+0
gclocals·IuErl7MOXaHVn7EZYWzfFA== SRODATA dupok size=10
	0x0000 02 00 00 00 02 00 00 00 02 00                    ..........
gclocals·J5F+7Qw7O7ve2QcWC7DpeQ== SRODATA dupok size=8
	0x0000 02 00 00 00 00 00 00 00                          ........
main.Duck.Quack.arginfo1 SRODATA static dupok size=7
	0x0000 fe 00 08 08 08 fd ff                             .......
