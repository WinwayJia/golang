"".Add STEXT nosplit size=15 args=0x10 locals=0x0
	0x0000 00000 (direct_calls.go:4)	TEXT	"".Add(SB), NOSPLIT, $0-16
	0x0000 00000 (direct_calls.go:4)	FUNCDATA	$0, gclocals·f207267fbf96a0178e8758c6e3e0ce28(SB)
	0x0000 00000 (direct_calls.go:4)	FUNCDATA	$1, gclocals·33cdeccccebe80329f1fdbee7f5874cb(SB)
	0x0000 00000 (direct_calls.go:4)	MOVL	"".b+12(SP), AX
	0x0004 00004 (direct_calls.go:4)	MOVL	"".a+8(SP), CX
	0x0008 00008 (direct_calls.go:4)	ADDL	CX, AX
	0x000a 00010 (direct_calls.go:4)	MOVL	AX, "".~r2+16(SP)
	0x000e 00014 (direct_calls.go:4)	RET
	0x0000 8b 44 24 0c 8b 4c 24 08 01 c8 89 44 24 10 c3     .D$..L$....D$..
"".(*Adder).AddPtr STEXT nosplit size=15 args=0x18 locals=0x0
	0x0000 00000 (direct_calls.go:9)	TEXT	"".(*Adder).AddPtr(SB), NOSPLIT, $0-24
	0x0000 00000 (direct_calls.go:9)	FUNCDATA	$0, gclocals·54241e171da8af6ae173d69da0236748(SB)
	0x0000 00000 (direct_calls.go:9)	FUNCDATA	$1, gclocals·33cdeccccebe80329f1fdbee7f5874cb(SB)
	0x0000 00000 (direct_calls.go:9)	MOVL	"".b+20(SP), AX
	0x0004 00004 (direct_calls.go:9)	MOVL	"".a+16(SP), CX
	0x0008 00008 (direct_calls.go:9)	ADDL	CX, AX
	0x000a 00010 (direct_calls.go:9)	MOVL	AX, "".~r2+24(SP)
	0x000e 00014 (direct_calls.go:9)	RET
	0x0000 8b 44 24 14 8b 4c 24 10 01 c8 89 44 24 18 c3     .D$..L$....D$..
"".Adder.AddVal STEXT nosplit size=15 args=0x18 locals=0x0
	0x0000 00000 (direct_calls.go:12)	TEXT	"".Adder.AddVal(SB), NOSPLIT, $0-24
	0x0000 00000 (direct_calls.go:12)	FUNCDATA	$0, gclocals·54241e171da8af6ae173d69da0236748(SB)
	0x0000 00000 (direct_calls.go:12)	FUNCDATA	$1, gclocals·33cdeccccebe80329f1fdbee7f5874cb(SB)
	0x0000 00000 (direct_calls.go:12)	MOVL	"".b+16(SP), AX
	0x0004 00004 (direct_calls.go:12)	MOVL	"".a+12(SP), CX
	0x0008 00008 (direct_calls.go:12)	ADDL	CX, AX
	0x000a 00010 (direct_calls.go:12)	MOVL	AX, "".~r2+24(SP)
	0x000e 00014 (direct_calls.go:12)	RET
	0x0000 8b 44 24 10 8b 4c 24 0c 01 c8 89 44 24 18 c3     .D$..L$....D$..
"".main STEXT size=163 args=0x0 locals=0x28
	0x0000 00000 (direct_calls.go:14)	TEXT	"".main(SB), $40-0
	0x0000 00000 (direct_calls.go:14)	MOVQ	(TLS), CX
	0x0009 00009 (direct_calls.go:14)	CMPQ	SP, 16(CX)
	0x000d 00013 (direct_calls.go:14)	JLS	153
	0x0013 00019 (direct_calls.go:14)	SUBQ	$40, SP
	0x0017 00023 (direct_calls.go:14)	MOVQ	BP, 32(SP)
	0x001c 00028 (direct_calls.go:14)	LEAQ	32(SP), BP
	0x0021 00033 (direct_calls.go:14)	FUNCDATA	$0, gclocals·33cdeccccebe80329f1fdbee7f5874cb(SB)
	0x0021 00033 (direct_calls.go:14)	FUNCDATA	$1, gclocals·33cdeccccebe80329f1fdbee7f5874cb(SB)
	0x0021 00033 (direct_calls.go:15)	MOVQ	$137438953482, AX
	0x002b 00043 (direct_calls.go:15)	MOVQ	AX, (SP)
	0x002f 00047 (direct_calls.go:15)	PCDATA	$0, $0
	0x002f 00047 (direct_calls.go:15)	CALL	"".Add(SB)
	0x0034 00052 (direct_calls.go:17)	MOVL	$6754, "".adder+28(SP)
	0x003c 00060 (direct_calls.go:18)	LEAQ	"".adder+28(SP), AX
	0x0041 00065 (direct_calls.go:18)	MOVQ	AX, (SP)
	0x0045 00069 (direct_calls.go:18)	MOVQ	$137438953482, AX
	0x004f 00079 (direct_calls.go:18)	MOVQ	AX, 8(SP)
	0x0054 00084 (direct_calls.go:18)	PCDATA	$0, $0
	0x0054 00084 (direct_calls.go:18)	CALL	"".(*Adder).AddPtr(SB)
	0x0059 00089 (direct_calls.go:19)	MOVL	"".adder+28(SP), AX
	0x005d 00093 (direct_calls.go:19)	MOVL	AX, (SP)
	0x0060 00096 (direct_calls.go:19)	MOVQ	$137438953482, AX
	0x006a 00106 (direct_calls.go:19)	MOVQ	AX, 4(SP)
	0x006f 00111 (direct_calls.go:19)	PCDATA	$0, $0
	0x006f 00111 (direct_calls.go:19)	CALL	"".Adder.AddVal(SB)
	0x0074 00116 (direct_calls.go:21)	MOVL	"".adder+28(SP), AX
	0x0078 00120 (direct_calls.go:21)	MOVL	AX, (SP)
	0x007b 00123 (direct_calls.go:21)	MOVQ	$137438953482, AX
	0x0085 00133 (direct_calls.go:21)	MOVQ	AX, 4(SP)
	0x008a 00138 (direct_calls.go:21)	PCDATA	$0, $0
	0x008a 00138 (direct_calls.go:21)	CALL	"".Adder.AddVal(SB)
	0x008f 00143 (direct_calls.go:22)	MOVQ	32(SP), BP
	0x0094 00148 (direct_calls.go:22)	ADDQ	$40, SP
	0x0098 00152 (direct_calls.go:22)	RET
	0x0099 00153 (direct_calls.go:22)	NOP
	0x0099 00153 (direct_calls.go:14)	PCDATA	$0, $-1
	0x0099 00153 (direct_calls.go:14)	CALL	runtime.morestack_noctxt(SB)
	0x009e 00158 (direct_calls.go:14)	JMP	0
	0x0000 64 48 8b 0c 25 00 00 00 00 48 3b 61 10 0f 86 86  dH..%....H;a....
	0x0010 00 00 00 48 83 ec 28 48 89 6c 24 20 48 8d 6c 24  ...H..(H.l$ H.l$
	0x0020 20 48 b8 0a 00 00 00 20 00 00 00 48 89 04 24 e8   H..... ...H..$.
	0x0030 00 00 00 00 c7 44 24 1c 62 1a 00 00 48 8d 44 24  .....D$.b...H.D$
	0x0040 1c 48 89 04 24 48 b8 0a 00 00 00 20 00 00 00 48  .H..$H..... ...H
	0x0050 89 44 24 08 e8 00 00 00 00 8b 44 24 1c 89 04 24  .D$.......D$...$
	0x0060 48 b8 0a 00 00 00 20 00 00 00 48 89 44 24 04 e8  H..... ...H.D$..
	0x0070 00 00 00 00 8b 44 24 1c 89 04 24 48 b8 0a 00 00  .....D$...$H....
	0x0080 00 20 00 00 00 48 89 44 24 04 e8 00 00 00 00 48  . ...H.D$......H
	0x0090 8b 6c 24 20 48 83 c4 28 c3 e8 00 00 00 00 e9 5d  .l$ H..(.......]
	0x00a0 ff ff ff                                         ...
	rel 5+4 t=16 TLS+0
	rel 48+4 t=8 "".Add+0
	rel 85+4 t=8 "".(*Adder).AddPtr+0
	rel 112+4 t=8 "".Adder.AddVal+0
	rel 139+4 t=8 "".Adder.AddVal+0
	rel 154+4 t=8 runtime.morestack_noctxt+0
"".init STEXT size=79 args=0x0 locals=0x8
	0x0000 00000 (<autogenerated>:1)	TEXT	"".init(SB), $8-0
	0x0000 00000 (<autogenerated>:1)	MOVQ	(TLS), CX
	0x0009 00009 (<autogenerated>:1)	CMPQ	SP, 16(CX)
	0x000d 00013 (<autogenerated>:1)	JLS	72
	0x000f 00015 (<autogenerated>:1)	SUBQ	$8, SP
	0x0013 00019 (<autogenerated>:1)	MOVQ	BP, (SP)
	0x0017 00023 (<autogenerated>:1)	LEAQ	(SP), BP
	0x001b 00027 (<autogenerated>:1)	FUNCDATA	$0, gclocals·33cdeccccebe80329f1fdbee7f5874cb(SB)
	0x001b 00027 (<autogenerated>:1)	FUNCDATA	$1, gclocals·33cdeccccebe80329f1fdbee7f5874cb(SB)
	0x001b 00027 (<autogenerated>:1)	MOVBLZX	"".initdone·(SB), AX
	0x0022 00034 (<autogenerated>:1)	CMPB	AL, $1
	0x0024 00036 (<autogenerated>:1)	JLS	47
	0x0026 00038 (<autogenerated>:1)	MOVQ	(SP), BP
	0x002a 00042 (<autogenerated>:1)	ADDQ	$8, SP
	0x002e 00046 (<autogenerated>:1)	RET
	0x002f 00047 (<autogenerated>:1)	JNE	56
	0x0031 00049 (<autogenerated>:1)	PCDATA	$0, $0
	0x0031 00049 (<autogenerated>:1)	CALL	runtime.throwinit(SB)
	0x0036 00054 (<autogenerated>:1)	UNDEF
	0x0038 00056 (<autogenerated>:1)	MOVB	$2, "".initdone·(SB)
	0x003f 00063 (<autogenerated>:1)	MOVQ	(SP), BP
	0x0043 00067 (<autogenerated>:1)	ADDQ	$8, SP
	0x0047 00071 (<autogenerated>:1)	RET
	0x0048 00072 (<autogenerated>:1)	NOP
	0x0048 00072 (<autogenerated>:1)	PCDATA	$0, $-1
	0x0048 00072 (<autogenerated>:1)	CALL	runtime.morestack_noctxt(SB)
	0x004d 00077 (<autogenerated>:1)	JMP	0
	0x0000 64 48 8b 0c 25 00 00 00 00 48 3b 61 10 76 39 48  dH..%....H;a.v9H
	0x0010 83 ec 08 48 89 2c 24 48 8d 2c 24 0f b6 05 00 00  ...H.,$H.,$.....
	0x0020 00 00 3c 01 76 09 48 8b 2c 24 48 83 c4 08 c3 75  ..<.v.H.,$H....u
	0x0030 07 e8 00 00 00 00 0f 0b c6 05 00 00 00 00 02 48  ...............H
	0x0040 8b 2c 24 48 83 c4 08 c3 e8 00 00 00 00 eb b1     .,$H...........
	rel 5+4 t=16 TLS+0
	rel 30+4 t=15 "".initdone·+0
	rel 50+4 t=8 runtime.throwinit+0
	rel 58+4 t=15 "".initdone·+-1
	rel 73+4 t=8 runtime.morestack_noctxt+0
"".(*Adder).AddVal STEXT dupok size=121 args=0x18 locals=0x20
	0x0000 00000 (<autogenerated>:1)	TEXT	"".(*Adder).AddVal(SB), DUPOK|WRAPPER, $32-24
	0x0000 00000 (<autogenerated>:1)	MOVQ	(TLS), CX
	0x0009 00009 (<autogenerated>:1)	CMPQ	SP, 16(CX)
	0x000d 00013 (<autogenerated>:1)	JLS	99
	0x000f 00015 (<autogenerated>:1)	SUBQ	$32, SP
	0x0013 00019 (<autogenerated>:1)	MOVQ	BP, 24(SP)
	0x0018 00024 (<autogenerated>:1)	LEAQ	24(SP), BP
	0x001d 00029 (<autogenerated>:1)	MOVQ	32(CX), BX
	0x0021 00033 (<autogenerated>:1)	TESTQ	BX, BX
	0x0024 00036 (<autogenerated>:1)	JNE	106
	0x0026 00038 (<autogenerated>:1)	NOP
	0x0026 00038 (<autogenerated>:1)	FUNCDATA	$0, gclocals·e6397a44f8e1b6e77d0f200b4fba5269(SB)
	0x0026 00038 (<autogenerated>:1)	FUNCDATA	$1, gclocals·69c1753bd5f81501d95132d08af04464(SB)
	0x0026 00038 (<autogenerated>:1)	MOVQ	""..this+40(SP), AX
	0x002b 00043 (<autogenerated>:1)	TESTQ	AX, AX
	0x002e 00046 (<autogenerated>:1)	JEQ	92
	0x0030 00048 (<autogenerated>:1)	MOVL	(AX), AX
	0x0032 00050 (<autogenerated>:1)	MOVL	AX, (SP)
	0x0035 00053 (<autogenerated>:1)	MOVL	"".a+48(SP), AX
	0x0039 00057 (<autogenerated>:1)	MOVL	AX, 4(SP)
	0x003d 00061 (<autogenerated>:1)	MOVL	"".b+52(SP), AX
	0x0041 00065 (<autogenerated>:1)	MOVL	AX, 8(SP)
	0x0045 00069 (<autogenerated>:1)	PCDATA	$0, $1
	0x0045 00069 (<autogenerated>:1)	CALL	"".Adder.AddVal(SB)
	0x004a 00074 (<autogenerated>:1)	MOVL	16(SP), AX
	0x004e 00078 (<autogenerated>:1)	MOVL	AX, "".~r2+56(SP)
	0x0052 00082 (<autogenerated>:1)	MOVQ	24(SP), BP
	0x0057 00087 (<autogenerated>:1)	ADDQ	$32, SP
	0x005b 00091 (<autogenerated>:1)	RET
	0x005c 00092 (<autogenerated>:1)	PCDATA	$0, $1
	0x005c 00092 (<autogenerated>:1)	CALL	runtime.panicwrap(SB)
	0x0061 00097 (<autogenerated>:1)	UNDEF
	0x0063 00099 (<autogenerated>:1)	NOP
	0x0063 00099 (<autogenerated>:1)	PCDATA	$0, $-1
	0x0063 00099 (<autogenerated>:1)	CALL	runtime.morestack_noctxt(SB)
	0x0068 00104 (<autogenerated>:1)	JMP	0
	0x006a 00106 (<autogenerated>:1)	LEAQ	40(SP), DI
	0x006f 00111 (<autogenerated>:1)	CMPQ	(BX), DI
	0x0072 00114 (<autogenerated>:1)	JNE	38
	0x0074 00116 (<autogenerated>:1)	MOVQ	SP, (BX)
	0x0077 00119 (<autogenerated>:1)	JMP	38
	0x0000 64 48 8b 0c 25 00 00 00 00 48 3b 61 10 76 54 48  dH..%....H;a.vTH
	0x0010 83 ec 20 48 89 6c 24 18 48 8d 6c 24 18 48 8b 59  .. H.l$.H.l$.H.Y
	0x0020 20 48 85 db 75 44 48 8b 44 24 28 48 85 c0 74 2c   H..uDH.D$(H..t,
	0x0030 8b 00 89 04 24 8b 44 24 30 89 44 24 04 8b 44 24  ....$.D$0.D$..D$
	0x0040 34 89 44 24 08 e8 00 00 00 00 8b 44 24 10 89 44  4.D$.......D$..D
	0x0050 24 38 48 8b 6c 24 18 48 83 c4 20 c3 e8 00 00 00  $8H.l$.H.. .....
	0x0060 00 0f 0b e8 00 00 00 00 eb 96 48 8d 7c 24 28 48  ..........H.|$(H
	0x0070 39 3b 75 b2 48 89 23 eb ad                       9;u.H.#..
	rel 5+4 t=16 TLS+0
	rel 70+4 t=8 "".Adder.AddVal+0
	rel 93+4 t=8 runtime.panicwrap+0
	rel 100+4 t=8 runtime.morestack_noctxt+0
go.info."".Add SDWARFINFO size=69
	0x0000 02 22 22 2e 41 64 64 00 00 00 00 00 00 00 00 00  ."".Add.........
	0x0010 00 00 00 00 00 00 00 00 01 9c 00 00 00 00 01 0e  ................
	0x0020 61 00 00 04 00 00 00 00 01 9c 0e 62 00 00 04 00  a..........b....
	0x0030 00 00 00 02 91 04 0e 7e 72 32 00 01 04 00 00 00  .......~r2......
	0x0040 00 02 91 08 00                                   .....
	rel 8+8 t=1 "".Add+0
	rel 16+8 t=1 "".Add+15
	rel 26+4 t=29 gofile../home/wenhuijia/go/src/golang/src/interview/asm/direct_calls.go+0
	rel 36+4 t=28 go.info.int32+0
	rel 47+4 t=28 go.info.int32+0
	rel 61+4 t=28 go.info.int32+0
go.range."".Add SDWARFRANGE size=0
go.info."".(*Adder).AddPtr SDWARFINFO size=97
	0x0000 02 22 22 2e 28 2a 41 64 64 65 72 29 2e 41 64 64  ."".(*Adder).Add
	0x0010 50 74 72 00 00 00 00 00 00 00 00 00 00 00 00 00  Ptr.............
	0x0020 00 00 00 00 01 9c 00 00 00 00 01 0e 61 00 00 09  ............a...
	0x0030 00 00 00 00 02 91 08 0e 61 64 64 65 72 00 00 09  ........adder...
	0x0040 00 00 00 00 01 9c 0e 62 00 00 09 00 00 00 00 02  .......b........
	0x0050 91 0c 0e 7e 72 32 00 01 09 00 00 00 00 02 91 10  ...~r2..........
	0x0060 00                                               .
	rel 20+8 t=1 "".(*Adder).AddPtr+0
	rel 28+8 t=1 "".(*Adder).AddPtr+15
	rel 38+4 t=29 gofile../home/wenhuijia/go/src/golang/src/interview/asm/direct_calls.go+0
	rel 48+4 t=28 go.info.int32+0
	rel 64+4 t=28 go.info.*"".Adder+0
	rel 75+4 t=28 go.info.int32+0
	rel 89+4 t=28 go.info.int32+0
go.range."".(*Adder).AddPtr SDWARFRANGE size=0
go.info."".Adder.AddVal SDWARFINFO size=94
	0x0000 02 22 22 2e 41 64 64 65 72 2e 41 64 64 56 61 6c  ."".Adder.AddVal
	0x0010 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0020 00 01 9c 00 00 00 00 01 0e 61 00 00 0c 00 00 00  .........a......
	0x0030 00 02 91 04 0e 61 64 64 65 72 00 00 0c 00 00 00  .....adder......
	0x0040 00 01 9c 0e 62 00 00 0c 00 00 00 00 02 91 08 0e  ....b...........
	0x0050 7e 72 32 00 01 0c 00 00 00 00 02 91 10 00        ~r2...........
	rel 17+8 t=1 "".Adder.AddVal+0
	rel 25+8 t=1 "".Adder.AddVal+15
	rel 35+4 t=29 gofile../home/wenhuijia/go/src/golang/src/interview/asm/direct_calls.go+0
	rel 45+4 t=28 go.info.int32+0
	rel 61+4 t=28 go.info."".Adder+0
	rel 72+4 t=28 go.info.int32+0
	rel 86+4 t=28 go.info.int32+0
go.range."".Adder.AddVal SDWARFRANGE size=0
go.info."".main SDWARFINFO size=48
	0x0000 02 22 22 2e 6d 61 69 6e 00 00 00 00 00 00 00 00  ."".main........
	0x0010 00 00 00 00 00 00 00 00 00 01 9c 00 00 00 00 01  ................
	0x0020 09 61 64 64 65 72 00 11 00 00 00 00 02 91 6c 00  .adder........l.
	rel 9+8 t=1 "".main+0
	rel 17+8 t=1 "".main+163
	rel 27+4 t=29 gofile../home/wenhuijia/go/src/golang/src/interview/asm/direct_calls.go+0
	rel 40+4 t=28 go.info."".Adder+0
go.range."".main SDWARFRANGE size=0
go.info."".init SDWARFINFO size=33
	0x0000 02 22 22 2e 69 6e 69 74 00 00 00 00 00 00 00 00  ."".init........
	0x0010 00 00 00 00 00 00 00 00 00 01 9c 00 00 00 00 01  ................
	0x0020 00                                               .
	rel 9+8 t=1 "".init+0
	rel 17+8 t=1 "".init+79
	rel 27+4 t=29 gofile..<autogenerated>+0
go.range."".init SDWARFRANGE size=0
"".initdone· SNOPTRBSS size=1
runtime.gcbits.01 SRODATA dupok size=1
	0x0000 01                                               .
go.info."".(*Adder).AddVal SDWARFINFO dupok size=97
	0x0000 02 22 22 2e 28 2a 41 64 64 65 72 29 2e 41 64 64  ."".(*Adder).Add
	0x0010 56 61 6c 00 00 00 00 00 00 00 00 00 00 00 00 00  Val.............
	0x0020 00 00 00 00 01 9c 00 00 00 00 01 0e 2e 74 68 69  .............thi
	0x0030 73 00 00 01 00 00 00 00 01 9c 0e 61 00 00 01 00  s..........a....
	0x0040 00 00 00 02 91 08 0e 62 00 00 01 00 00 00 00 02  .......b........
	0x0050 91 0c 0e 7e 72 32 00 01 01 00 00 00 00 02 91 10  ...~r2..........
	0x0060 00                                               .
	rel 20+8 t=1 "".(*Adder).AddVal+0
	rel 28+8 t=1 "".(*Adder).AddVal+121
	rel 38+4 t=29 gofile..<autogenerated>+0
	rel 52+4 t=28 go.info.*"".Adder+0
	rel 63+4 t=28 go.info.int32+0
	rel 75+4 t=28 go.info.int32+0
	rel 89+4 t=28 go.info.int32+0
go.range."".(*Adder).AddVal SDWARFRANGE dupok size=0
type..namedata.*main.Adder. SRODATA dupok size=14
	0x0000 01 00 0b 2a 6d 61 69 6e 2e 41 64 64 65 72        ...*main.Adder
type..namedata.*func(*main.Adder, int32, int32) int32- SRODATA dupok size=41
	0x0000 00 00 26 2a 66 75 6e 63 28 2a 6d 61 69 6e 2e 41  ..&*func(*main.A
	0x0010 64 64 65 72 2c 20 69 6e 74 33 32 2c 20 69 6e 74  dder, int32, int
	0x0020 33 32 29 20 69 6e 74 33 32                       32) int32
type.*func(*"".Adder, int32, int32) int32 SRODATA dupok size=56
	0x0000 08 00 00 00 00 00 00 00 08 00 00 00 00 00 00 00  ................
	0x0010 35 26 2b 33 00 08 08 36 00 00 00 00 00 00 00 00  5&+3...6........
	0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0030 00 00 00 00 00 00 00 00                          ........
	rel 24+8 t=1 runtime.algarray+80
	rel 32+8 t=1 runtime.gcbits.01+0
	rel 40+4 t=5 type..namedata.*func(*main.Adder, int32, int32) int32-+0
	rel 48+8 t=1 type.func(*"".Adder, int32, int32) int32+0
type.func(*"".Adder, int32, int32) int32 SRODATA dupok size=88
	0x0000 08 00 00 00 00 00 00 00 08 00 00 00 00 00 00 00  ................
	0x0010 d7 38 ec a1 02 08 08 33 00 00 00 00 00 00 00 00  .8.....3........
	0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0030 03 00 01 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0040 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0050 00 00 00 00 00 00 00 00                          ........
	rel 24+8 t=1 runtime.algarray+0
	rel 32+8 t=1 runtime.gcbits.01+0
	rel 40+4 t=5 type..namedata.*func(*main.Adder, int32, int32) int32-+0
	rel 44+4 t=6 type.*func(*"".Adder, int32, int32) int32+0
	rel 56+8 t=1 type.*"".Adder+0
	rel 64+8 t=1 type.int32+0
	rel 72+8 t=1 type.int32+0
	rel 80+8 t=1 type.int32+0
type..namedata.AddPtr. SRODATA dupok size=9
	0x0000 01 00 06 41 64 64 50 74 72                       ...AddPtr
type..namedata.*func(int32, int32) int32- SRODATA dupok size=28
	0x0000 00 00 19 2a 66 75 6e 63 28 69 6e 74 33 32 2c 20  ...*func(int32, 
	0x0010 69 6e 74 33 32 29 20 69 6e 74 33 32              int32) int32
type.*func(int32, int32) int32 SRODATA dupok size=56
	0x0000 08 00 00 00 00 00 00 00 08 00 00 00 00 00 00 00  ................
	0x0010 c0 41 58 87 00 08 08 36 00 00 00 00 00 00 00 00  .AX....6........
	0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0030 00 00 00 00 00 00 00 00                          ........
	rel 24+8 t=1 runtime.algarray+80
	rel 32+8 t=1 runtime.gcbits.01+0
	rel 40+4 t=5 type..namedata.*func(int32, int32) int32-+0
	rel 48+8 t=1 type.func(int32, int32) int32+0
type.func(int32, int32) int32 SRODATA dupok size=80
	0x0000 08 00 00 00 00 00 00 00 08 00 00 00 00 00 00 00  ................
	0x0010 13 9f 60 7f 02 08 08 33 00 00 00 00 00 00 00 00  ..`....3........
	0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0030 02 00 01 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0040 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	rel 24+8 t=1 runtime.algarray+0
	rel 32+8 t=1 runtime.gcbits.01+0
	rel 40+4 t=5 type..namedata.*func(int32, int32) int32-+0
	rel 44+4 t=6 type.*func(int32, int32) int32+0
	rel 56+8 t=1 type.int32+0
	rel 64+8 t=1 type.int32+0
	rel 72+8 t=1 type.int32+0
type..namedata.AddVal. SRODATA dupok size=9
	0x0000 01 00 06 41 64 64 56 61 6c                       ...AddVal
type.*"".Adder SRODATA size=104
	0x0000 08 00 00 00 00 00 00 00 08 00 00 00 00 00 00 00  ................
	0x0010 46 c2 6b c4 01 08 08 36 00 00 00 00 00 00 00 00  F.k....6........
	0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0030 00 00 00 00 00 00 00 00 00 00 00 00 02 00 00 00  ................
	0x0040 10 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0050 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0060 00 00 00 00 00 00 00 00                          ........
	rel 24+8 t=1 runtime.algarray+80
	rel 32+8 t=1 runtime.gcbits.01+0
	rel 40+4 t=5 type..namedata.*main.Adder.+0
	rel 48+8 t=1 type."".Adder+0
	rel 56+4 t=5 type..importpath."".+0
	rel 72+4 t=5 type..namedata.AddPtr.+0
	rel 76+4 t=24 type.func(int32, int32) int32+0
	rel 80+4 t=24 "".(*Adder).AddPtr+0
	rel 84+4 t=24 "".(*Adder).AddPtr+0
	rel 88+4 t=5 type..namedata.AddVal.+0
	rel 92+4 t=24 type.func(int32, int32) int32+0
	rel 96+4 t=24 "".(*Adder).AddVal+0
	rel 100+4 t=24 "".(*Adder).AddVal+0
runtime.gcbits. SRODATA dupok size=0
type..namedata.*func(main.Adder, int32, int32) int32- SRODATA dupok size=40
	0x0000 00 00 25 2a 66 75 6e 63 28 6d 61 69 6e 2e 41 64  ..%*func(main.Ad
	0x0010 64 65 72 2c 20 69 6e 74 33 32 2c 20 69 6e 74 33  der, int32, int3
	0x0020 32 29 20 69 6e 74 33 32                          2) int32
type.*func("".Adder, int32, int32) int32 SRODATA dupok size=56
	0x0000 08 00 00 00 00 00 00 00 08 00 00 00 00 00 00 00  ................
	0x0010 ec 3b d4 ed 00 08 08 36 00 00 00 00 00 00 00 00  .;.....6........
	0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0030 00 00 00 00 00 00 00 00                          ........
	rel 24+8 t=1 runtime.algarray+80
	rel 32+8 t=1 runtime.gcbits.01+0
	rel 40+4 t=5 type..namedata.*func(main.Adder, int32, int32) int32-+0
	rel 48+8 t=1 type.func("".Adder, int32, int32) int32+0
type.func("".Adder, int32, int32) int32 SRODATA dupok size=88
	0x0000 08 00 00 00 00 00 00 00 08 00 00 00 00 00 00 00  ................
	0x0010 80 6d 70 15 02 08 08 33 00 00 00 00 00 00 00 00  .mp....3........
	0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0030 03 00 01 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0040 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0050 00 00 00 00 00 00 00 00                          ........
	rel 24+8 t=1 runtime.algarray+0
	rel 32+8 t=1 runtime.gcbits.01+0
	rel 40+4 t=5 type..namedata.*func(main.Adder, int32, int32) int32-+0
	rel 44+4 t=6 type.*func("".Adder, int32, int32) int32+0
	rel 56+8 t=1 type."".Adder+0
	rel 64+8 t=1 type.int32+0
	rel 72+8 t=1 type.int32+0
	rel 80+8 t=1 type.int32+0
type..namedata.id- SRODATA dupok size=5
	0x0000 00 00 02 69 64                                   ...id
type."".Adder SRODATA size=136
	0x0000 04 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0010 8a 3d 5f 61 07 04 04 99 00 00 00 00 00 00 00 00  .=_a............
	0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0030 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0040 01 00 00 00 00 00 00 00 01 00 00 00 00 00 00 00  ................
	0x0050 00 00 00 00 01 00 00 00 28 00 00 00 00 00 00 00  ........(.......
	0x0060 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0070 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0080 00 00 00 00 00 00 00 00                          ........
	rel 24+8 t=1 runtime.algarray+64
	rel 32+8 t=1 runtime.gcbits.+0
	rel 40+4 t=5 type..namedata.*main.Adder.+0
	rel 44+4 t=5 type.*"".Adder+0
	rel 48+8 t=1 type..importpath."".+0
	rel 56+8 t=1 type."".Adder+96
	rel 80+4 t=5 type..importpath."".+0
	rel 96+8 t=1 type..namedata.id-+0
	rel 104+8 t=1 type.int32+0
	rel 120+4 t=5 type..namedata.AddVal.+0
	rel 124+4 t=24 type.func(int32, int32) int32+0
	rel 128+4 t=24 "".(*Adder).AddVal+0
	rel 132+4 t=24 "".Adder.AddVal+0
gclocals·f207267fbf96a0178e8758c6e3e0ce28 SRODATA dupok size=9
	0x0000 01 00 00 00 02 00 00 00 00                       .........
gclocals·33cdeccccebe80329f1fdbee7f5874cb SRODATA dupok size=8
	0x0000 01 00 00 00 00 00 00 00                          ........
gclocals·54241e171da8af6ae173d69da0236748 SRODATA dupok size=9
	0x0000 01 00 00 00 03 00 00 00 00                       .........
gclocals·e6397a44f8e1b6e77d0f200b4fba5269 SRODATA dupok size=10
	0x0000 02 00 00 00 03 00 00 00 01 00                    ..........
gclocals·69c1753bd5f81501d95132d08af04464 SRODATA dupok size=8
	0x0000 02 00 00 00 00 00 00 00                          ........
