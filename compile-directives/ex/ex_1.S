"".appendStr STEXT size=94 args=0x10 locals=0x1c
	0x0000 00000 (ex_1.go:6)	TEXT	"".appendStr(SB), ABIInternal, $28-16
	0x0000 00000 (ex_1.go:6)	MOVL	TLS, CX
	0x0007 00007 (ex_1.go:6)	MOVL	(CX)(TLS*2), CX
	0x000d 00013 (ex_1.go:6)	CMPL	SP, 8(CX)
	0x0010 00016 (ex_1.go:6)	JLS	87
	0x0012 00018 (ex_1.go:6)	SUBL	$28, SP
	0x0015 00021 (ex_1.go:6)	FUNCDATA	$0, gclocals·9fad110d66c97cf0b58d28cccea80b12(SB)
	0x0015 00021 (ex_1.go:6)	FUNCDATA	$1, gclocals·7d2d5fca80364273fb07d5820a76fef4(SB)
	0x0015 00021 (ex_1.go:6)	FUNCDATA	$2, gclocals·9fb7f0986f647f17cb53dda1484e0f7a(SB)
	0x0015 00021 (ex_1.go:7)	PCDATA	$0, $0
	0x0015 00021 (ex_1.go:7)	PCDATA	$1, $0
	0x0015 00021 (ex_1.go:7)	MOVL	$0, (SP)
	0x001c 00028 (ex_1.go:7)	PCDATA	$0, $1
	0x001c 00028 (ex_1.go:7)	LEAL	go.string."new "(SB), AX
	0x0022 00034 (ex_1.go:7)	PCDATA	$0, $0
	0x0022 00034 (ex_1.go:7)	MOVL	AX, 4(SP)
	0x0026 00038 (ex_1.go:7)	MOVL	$4, 8(SP)
	0x002e 00046 (ex_1.go:7)	PCDATA	$0, $1
	0x002e 00046 (ex_1.go:7)	MOVL	"".word+32(SP), AX
	0x0032 00050 (ex_1.go:7)	PCDATA	$0, $0
	0x0032 00050 (ex_1.go:7)	MOVL	AX, 12(SP)
	0x0036 00054 (ex_1.go:7)	PCDATA	$1, $1
	0x0036 00054 (ex_1.go:7)	MOVL	"".word+36(SP), AX
	0x003a 00058 (ex_1.go:7)	MOVL	AX, 16(SP)
	0x003e 00062 (ex_1.go:7)	CALL	runtime.concatstring2(SB)
	0x0043 00067 (ex_1.go:7)	PCDATA	$0, $1
	0x0043 00067 (ex_1.go:7)	MOVL	20(SP), AX
	0x0047 00071 (ex_1.go:7)	MOVL	24(SP), CX
	0x004b 00075 (ex_1.go:7)	PCDATA	$0, $0
	0x004b 00075 (ex_1.go:7)	PCDATA	$1, $2
	0x004b 00075 (ex_1.go:7)	MOVL	AX, "".~r1+40(SP)
	0x004f 00079 (ex_1.go:7)	MOVL	CX, "".~r1+44(SP)
	0x0053 00083 (ex_1.go:7)	ADDL	$28, SP
	0x0056 00086 (ex_1.go:7)	RET
	0x0057 00087 (ex_1.go:7)	NOP
	0x0057 00087 (ex_1.go:6)	PCDATA	$1, $-1
	0x0057 00087 (ex_1.go:6)	PCDATA	$0, $-1
	0x0057 00087 (ex_1.go:6)	CALL	runtime.morestack_noctxt(SB)
	0x005c 00092 (ex_1.go:6)	JMP	0
	0x0000 65 8b 0d 00 00 00 00 8b 89 00 00 00 00 3b 61 08  e............;a.
	0x0010 76 45 83 ec 1c c7 04 24 00 00 00 00 8d 05 00 00  vE.....$........
	0x0020 00 00 89 44 24 04 c7 44 24 08 04 00 00 00 8b 44  ...D$..D$......D
	0x0030 24 20 89 44 24 0c 8b 44 24 24 89 44 24 10 e8 00  $ .D$..D$$.D$...
	0x0040 00 00 00 8b 44 24 14 8b 4c 24 18 89 44 24 28 89  ....D$..L$..D$(.
	0x0050 4c 24 2c 83 c4 1c c3 e8 00 00 00 00 eb a2        L$,...........
	rel 9+4 t=16 TLS+0
	rel 30+4 t=1 go.string."new "+0
	rel 63+4 t=8 runtime.concatstring2+0
	rel 88+4 t=8 runtime.morestack_noctxt+0
"".appendStr2 STEXT size=94 args=0x10 locals=0x1c
	0x0000 00000 (ex_1.go:11)	TEXT	"".appendStr2(SB), ABIInternal, $28-16
	0x0000 00000 (ex_1.go:11)	MOVL	TLS, CX
	0x0007 00007 (ex_1.go:11)	MOVL	(CX)(TLS*2), CX
	0x000d 00013 (ex_1.go:11)	CMPL	SP, 8(CX)
	0x0010 00016 (ex_1.go:11)	JLS	87
	0x0012 00018 (ex_1.go:11)	SUBL	$28, SP
	0x0015 00021 (ex_1.go:11)	FUNCDATA	$0, gclocals·9fad110d66c97cf0b58d28cccea80b12(SB)
	0x0015 00021 (ex_1.go:11)	FUNCDATA	$1, gclocals·7d2d5fca80364273fb07d5820a76fef4(SB)
	0x0015 00021 (ex_1.go:11)	FUNCDATA	$2, gclocals·9fb7f0986f647f17cb53dda1484e0f7a(SB)
	0x0015 00021 (ex_1.go:12)	PCDATA	$0, $0
	0x0015 00021 (ex_1.go:12)	PCDATA	$1, $0
	0x0015 00021 (ex_1.go:12)	MOVL	$0, (SP)
	0x001c 00028 (ex_1.go:12)	PCDATA	$0, $1
	0x001c 00028 (ex_1.go:12)	LEAL	go.string."new "(SB), AX
	0x0022 00034 (ex_1.go:12)	PCDATA	$0, $0
	0x0022 00034 (ex_1.go:12)	MOVL	AX, 4(SP)
	0x0026 00038 (ex_1.go:12)	MOVL	$4, 8(SP)
	0x002e 00046 (ex_1.go:12)	PCDATA	$0, $1
	0x002e 00046 (ex_1.go:12)	MOVL	"".word+32(SP), AX
	0x0032 00050 (ex_1.go:12)	PCDATA	$0, $0
	0x0032 00050 (ex_1.go:12)	MOVL	AX, 12(SP)
	0x0036 00054 (ex_1.go:12)	PCDATA	$1, $1
	0x0036 00054 (ex_1.go:12)	MOVL	"".word+36(SP), AX
	0x003a 00058 (ex_1.go:12)	MOVL	AX, 16(SP)
	0x003e 00062 (ex_1.go:12)	CALL	runtime.concatstring2(SB)
	0x0043 00067 (ex_1.go:12)	PCDATA	$0, $1
	0x0043 00067 (ex_1.go:12)	MOVL	20(SP), AX
	0x0047 00071 (ex_1.go:12)	MOVL	24(SP), CX
	0x004b 00075 (ex_1.go:12)	PCDATA	$0, $0
	0x004b 00075 (ex_1.go:12)	PCDATA	$1, $2
	0x004b 00075 (ex_1.go:12)	MOVL	AX, "".~r1+40(SP)
	0x004f 00079 (ex_1.go:12)	MOVL	CX, "".~r1+44(SP)
	0x0053 00083 (ex_1.go:12)	ADDL	$28, SP
	0x0056 00086 (ex_1.go:12)	RET
	0x0057 00087 (ex_1.go:12)	NOP
	0x0057 00087 (ex_1.go:11)	PCDATA	$1, $-1
	0x0057 00087 (ex_1.go:11)	PCDATA	$0, $-1
	0x0057 00087 (ex_1.go:11)	CALL	runtime.morestack_noctxt(SB)
	0x005c 00092 (ex_1.go:11)	JMP	0
	0x0000 65 8b 0d 00 00 00 00 8b 89 00 00 00 00 3b 61 08  e............;a.
	0x0010 76 45 83 ec 1c c7 04 24 00 00 00 00 8d 05 00 00  vE.....$........
	0x0020 00 00 89 44 24 04 c7 44 24 08 04 00 00 00 8b 44  ...D$..D$......D
	0x0030 24 20 89 44 24 0c 8b 44 24 24 89 44 24 10 e8 00  $ .D$..D$$.D$...
	0x0040 00 00 00 8b 44 24 14 8b 4c 24 18 89 44 24 28 89  ....D$..L$..D$(.
	0x0050 4c 24 2c 83 c4 1c c3 e8 00 00 00 00 eb a2        L$,...........
	rel 9+4 t=16 TLS+0
	rel 30+4 t=1 go.string."new "+0
	rel 63+4 t=8 runtime.concatstring2+0
	rel 88+4 t=8 runtime.morestack_noctxt+0
"".test STEXT size=102 args=0x0 locals=0x3c
	0x0000 00000 (ex_1.go:15)	TEXT	"".test(SB), ABIInternal, $60-0
	0x0000 00000 (ex_1.go:15)	MOVL	TLS, CX
	0x0007 00007 (ex_1.go:15)	MOVL	(CX)(TLS*2), CX
	0x000d 00013 (ex_1.go:15)	CMPL	SP, 8(CX)
	0x0010 00016 (ex_1.go:15)	JLS	95
	0x0012 00018 (ex_1.go:15)	SUBL	$60, SP
	0x0015 00021 (ex_1.go:15)	FUNCDATA	$0, gclocals·33cdeccccebe80329f1fdbee7f5874cb(SB)
	0x0015 00021 (ex_1.go:15)	FUNCDATA	$1, gclocals·33cdeccccebe80329f1fdbee7f5874cb(SB)
	0x0015 00021 (ex_1.go:15)	FUNCDATA	$2, gclocals·9fb7f0986f647f17cb53dda1484e0f7a(SB)
	0x0015 00021 (ex_1.go:16)	PCDATA	$0, $0
	0x0015 00021 (ex_1.go:16)	PCDATA	$1, $0
	0x0015 00021 (<unknown line number>)	NOP
	0x0015 00021 (ex_1.go:16)	PCDATA	$0, $1
	0x0015 00021 (ex_1.go:16)	LEAL	""..autotmp_3+28(SP), AX
	0x0019 00025 (ex_1.go:7)	PCDATA	$0, $0
	0x0019 00025 (ex_1.go:7)	MOVL	AX, (SP)
	0x001c 00028 (ex_1.go:7)	PCDATA	$0, $1
	0x001c 00028 (ex_1.go:7)	LEAL	go.string."new "(SB), AX
	0x0022 00034 (ex_1.go:7)	PCDATA	$0, $0
	0x0022 00034 (ex_1.go:7)	MOVL	AX, 4(SP)
	0x0026 00038 (ex_1.go:7)	MOVL	$4, 8(SP)
	0x002e 00046 (ex_1.go:7)	PCDATA	$0, $1
	0x002e 00046 (ex_1.go:7)	LEAL	go.string."11"(SB), AX
	0x0034 00052 (ex_1.go:7)	PCDATA	$0, $0
	0x0034 00052 (ex_1.go:7)	MOVL	AX, 12(SP)
	0x0038 00056 (ex_1.go:7)	MOVL	$2, 16(SP)
	0x0040 00064 (ex_1.go:7)	CALL	runtime.concatstring2(SB)
	0x0045 00069 (ex_1.go:17)	PCDATA	$0, $1
	0x0045 00069 (ex_1.go:17)	LEAL	go.string."22"(SB), AX
	0x004b 00075 (ex_1.go:17)	PCDATA	$0, $0
	0x004b 00075 (ex_1.go:17)	MOVL	AX, (SP)
	0x004e 00078 (ex_1.go:17)	MOVL	$2, 4(SP)
	0x0056 00086 (ex_1.go:17)	CALL	"".appendStr2(SB)
	0x005b 00091 (ex_1.go:18)	ADDL	$60, SP
	0x005e 00094 (ex_1.go:18)	RET
	0x005f 00095 (ex_1.go:18)	NOP
	0x005f 00095 (ex_1.go:15)	PCDATA	$1, $-1
	0x005f 00095 (ex_1.go:15)	PCDATA	$0, $-1
	0x005f 00095 (ex_1.go:15)	CALL	runtime.morestack_noctxt(SB)
	0x0064 00100 (ex_1.go:15)	JMP	0
	0x0000 65 8b 0d 00 00 00 00 8b 89 00 00 00 00 3b 61 08  e............;a.
	0x0010 76 4d 83 ec 3c 8d 44 24 1c 89 04 24 8d 05 00 00  vM..<.D$...$....
	0x0020 00 00 89 44 24 04 c7 44 24 08 04 00 00 00 8d 05  ...D$..D$.......
	0x0030 00 00 00 00 89 44 24 0c c7 44 24 10 02 00 00 00  .....D$..D$.....
	0x0040 e8 00 00 00 00 8d 05 00 00 00 00 89 04 24 c7 44  .............$.D
	0x0050 24 04 02 00 00 00 e8 00 00 00 00 83 c4 3c c3 e8  $............<..
	0x0060 00 00 00 00 eb 9a                                ......
	rel 9+4 t=16 TLS+0
	rel 30+4 t=1 go.string."new "+0
	rel 48+4 t=1 go.string."11"+0
	rel 65+4 t=8 runtime.concatstring2+0
	rel 71+4 t=1 go.string."22"+0
	rel 87+4 t=8 "".appendStr2+0
	rel 96+4 t=8 runtime.morestack_noctxt+0
go.cuinfo.packagename. SDWARFINFO dupok size=0
	0x0000 65 78                                            ex
go.info."".appendStr$abstract SDWARFINFO dupok size=26
	0x0000 04 2e 61 70 70 65 6e 64 53 74 72 00 01 01 11 77  ..appendStr....w
	0x0010 6f 72 64 00 00 00 00 00 00 00                    ord.......
	rel 21+4 t=28 go.info.string+0
go.string."new " SRODATA dupok size=4
	0x0000 6e 65 77 20                                      new 
go.loc."".appendStr SDWARFLOC size=33
	0x0000 ff ff ff ff 00 00 00 00 00 00 00 00 5e 00 00 00  ............^...
	0x0010 07 00 9c 93 04 91 04 93 04 00 00 00 00 00 00 00  ................
	0x0020 00                                               .
	rel 4+4 t=1 "".appendStr+0
go.info."".appendStr SDWARFINFO size=37
	0x0000 05 00 00 00 00 00 00 00 00 00 00 00 00 01 9c 13  ................
	0x0010 00 00 00 00 00 00 00 00 0f 7e 72 31 00 01 06 00  .........~r1....
	0x0020 00 00 00 00 00                                   .....
	rel 1+4 t=28 go.info."".appendStr$abstract+0
	rel 5+4 t=1 "".appendStr+0
	rel 9+4 t=1 "".appendStr+94
	rel 16+4 t=28 go.info."".appendStr$abstract+14
	rel 20+4 t=28 go.loc."".appendStr+0
	rel 31+4 t=28 go.info.string+0
go.range."".appendStr SDWARFRANGE size=0
go.isstmt."".appendStr SDWARFMISC size=0
	0x0000 04 12 04 03 03 07 01 22 02 05 01 14 02 07 00     .......".......
go.loc."".appendStr2 SDWARFLOC size=33
	0x0000 ff ff ff ff 00 00 00 00 00 00 00 00 5e 00 00 00  ............^...
	0x0010 07 00 9c 93 04 91 04 93 04 00 00 00 00 00 00 00  ................
	0x0020 00                                               .
	rel 4+4 t=1 "".appendStr2+0
go.info."".appendStr2 SDWARFINFO size=59
	0x0000 03 22 22 2e 61 70 70 65 6e 64 53 74 72 32 00 00  ."".appendStr2..
	0x0010 00 00 00 00 00 00 00 01 9c 00 00 00 00 01 10 77  ...............w
	0x0020 6f 72 64 00 00 0b 00 00 00 00 00 00 00 00 0f 7e  ord............~
	0x0030 72 31 00 01 0b 00 00 00 00 00 00                 r1.........
	rel 15+4 t=1 "".appendStr2+0
	rel 19+4 t=1 "".appendStr2+94
	rel 25+4 t=29 gofile../Users/shichaopeng/Work/golang-exercise/compile-directives/ex/ex_1.go+0
	rel 38+4 t=28 go.info.string+0
	rel 42+4 t=28 go.loc."".appendStr2+0
	rel 53+4 t=28 go.info.string+0
go.range."".appendStr2 SDWARFRANGE size=0
go.isstmt."".appendStr2 SDWARFMISC size=0
	0x0000 04 12 04 03 03 07 01 22 02 05 01 14 02 07 00     .......".......
go.string."11" SRODATA dupok size=2
	0x0000 31 31                                            11
go.string."22" SRODATA dupok size=2
	0x0000 32 32                                            22
go.loc."".test SDWARFLOC size=0
go.info."".test SDWARFINFO size=50
	0x0000 03 22 22 2e 74 65 73 74 00 00 00 00 00 00 00 00  ."".test........
	0x0010 00 01 9c 00 00 00 00 01 06 00 00 00 00 00 00 00  ................
	0x0020 00 00 00 00 00 00 00 00 00 10 12 00 00 00 00 00  ................
	0x0030 00 00                                            ..
	rel 9+4 t=1 "".test+0
	rel 13+4 t=1 "".test+102
	rel 19+4 t=29 gofile../Users/shichaopeng/Work/golang-exercise/compile-directives/ex/ex_1.go+0
	rel 25+4 t=28 go.info."".appendStr$abstract+0
	rel 29+4 t=1 "".test+25
	rel 33+4 t=1 "".test+69
	rel 37+4 t=29 gofile../Users/shichaopeng/Work/golang-exercise/compile-directives/ex/ex_1.go+0
	rel 43+4 t=28 go.info."".appendStr$abstract+14
go.range."".test SDWARFRANGE size=0
go.isstmt."".test SDWARFMISC size=0
	0x0000 04 12 04 03 03 07 01 24 02 0b 01 0b 02 10 00     .......$.......
runtime.gcbits.01 SRODATA dupok size=1
	0x0000 01                                               .
type..namedata.*[]uint8- SRODATA dupok size=11
	0x0000 00 00 08 2a 5b 5d 75 69 6e 74 38                 ...*[]uint8
type.*[]uint8 SRODATA dupok size=36
	0x0000 04 00 00 00 04 00 00 00 a5 8e d0 69 00 04 04 36  ...........i...6
	0x0010 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0020 00 00 00 00                                      ....
	rel 16+4 t=1 runtime.algarray+32
	rel 20+4 t=1 runtime.gcbits.01+0
	rel 24+4 t=5 type..namedata.*[]uint8-+0
	rel 32+4 t=1 type.[]uint8+0
type.[]uint8 SRODATA dupok size=36
	0x0000 0c 00 00 00 04 00 00 00 df 7e 2e 38 02 04 04 17  .........~.8....
	0x0010 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0020 00 00 00 00                                      ....
	rel 16+4 t=1 runtime.algarray+0
	rel 20+4 t=1 runtime.gcbits.01+0
	rel 24+4 t=5 type..namedata.*[]uint8-+0
	rel 28+4 t=6 type.*[]uint8+0
	rel 32+4 t=1 type.uint8+0
type..hashfunc32 SRODATA dupok size=8
	0x0000 00 00 00 00 20 00 00 00                          .... ...
	rel 0+4 t=1 runtime.memhash_varlen+0
type..eqfunc32 SRODATA dupok size=8
	0x0000 00 00 00 00 20 00 00 00                          .... ...
	rel 0+4 t=1 runtime.memequal_varlen+0
type..alg32 SRODATA dupok size=8
	0x0000 00 00 00 00 00 00 00 00                          ........
	rel 0+4 t=1 type..hashfunc32+0
	rel 4+4 t=1 type..eqfunc32+0
type..namedata.*[32]uint8- SRODATA dupok size=13
	0x0000 00 00 0a 2a 5b 33 32 5d 75 69 6e 74 38           ...*[32]uint8
type.*[32]uint8 SRODATA dupok size=36
	0x0000 04 00 00 00 04 00 00 00 f4 c7 79 15 00 04 04 36  ..........y....6
	0x0010 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0020 00 00 00 00                                      ....
	rel 16+4 t=1 runtime.algarray+32
	rel 20+4 t=1 runtime.gcbits.01+0
	rel 24+4 t=5 type..namedata.*[32]uint8-+0
	rel 32+4 t=1 type.[32]uint8+0
runtime.gcbits. SRODATA dupok size=0
type.[32]uint8 SRODATA dupok size=44
	0x0000 20 00 00 00 00 00 00 00 9c 59 ff a8 02 01 01 11   ........Y......
	0x0010 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0020 00 00 00 00 00 00 00 00 20 00 00 00              ........ ...
	rel 16+4 t=1 type..alg32+0
	rel 20+4 t=1 runtime.gcbits.+0
	rel 24+4 t=5 type..namedata.*[32]uint8-+0
	rel 28+4 t=6 type.*[32]uint8+0
	rel 32+4 t=1 type.uint8+0
	rel 36+4 t=1 type.[]uint8+0
gclocals·9fad110d66c97cf0b58d28cccea80b12 SRODATA dupok size=11
	0x0000 03 00 00 00 03 00 00 00 01 00 04                 ...........
gclocals·7d2d5fca80364273fb07d5820a76fef4 SRODATA dupok size=8
	0x0000 03 00 00 00 00 00 00 00                          ........
gclocals·9fb7f0986f647f17cb53dda1484e0f7a SRODATA dupok size=10
	0x0000 02 00 00 00 01 00 00 00 00 01                    ..........
gclocals·33cdeccccebe80329f1fdbee7f5874cb SRODATA dupok size=8
	0x0000 01 00 00 00 00 00 00 00                          ........
