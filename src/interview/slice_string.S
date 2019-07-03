"".slice2String STEXT size=223 args=0x0 locals=0x60
	0x0000 00000 (slice_string.go:5)	TEXT	"".slice2String(SB), $96-0
	0x0000 00000 (slice_string.go:5)	MOVQ	(TLS), CX
	0x0009 00009 (slice_string.go:5)	CMPQ	SP, 16(CX)
	0x000d 00013 (slice_string.go:5)	JLS	213
	0x0013 00019 (slice_string.go:5)	SUBQ	$96, SP
	0x0017 00023 (slice_string.go:5)	MOVQ	BP, 88(SP)
	0x001c 00028 (slice_string.go:5)	LEAQ	88(SP), BP
	0x0021 00033 (slice_string.go:5)	FUNCDATA	$0, gclocals·69c1753bd5f81501d95132d08af04464(SB)
	0x0021 00033 (slice_string.go:5)	FUNCDATA	$1, gclocals·adc9a4e9eec70025782f0016681b6ae2(SB)
	0x0021 00033 (slice_string.go:6)	MOVWLZX	"".statictmp_0(SB), AX
	0x0028 00040 (slice_string.go:6)	MOVBLZX	"".statictmp_0+2(SB), CX
	0x002f 00047 (slice_string.go:6)	MOVW	AX, ""..autotmp_4+53(SP)
	0x0034 00052 (slice_string.go:6)	MOVB	CL, ""..autotmp_4+55(SP)
	0x0038 00056 (slice_string.go:7)	MOVQ	$0, (SP)
	0x0040 00064 (slice_string.go:7)	LEAQ	""..autotmp_4+53(SP), AX
	0x0045 00069 (slice_string.go:7)	MOVQ	AX, 8(SP)
	0x004a 00074 (slice_string.go:7)	MOVQ	$3, 16(SP)
	0x0053 00083 (slice_string.go:7)	MOVQ	$3, 24(SP)
	0x005c 00092 (slice_string.go:7)	PCDATA	$0, $0
	0x005c 00092 (slice_string.go:7)	CALL	runtime.slicebytetostring(SB)
	0x0061 00097 (slice_string.go:7)	MOVQ	32(SP), AX
	0x0066 00102 (slice_string.go:7)	MOVQ	40(SP), CX
	0x006b 00107 (slice_string.go:7)	MOVQ	AX, ""..autotmp_2+56(SP)
	0x0070 00112 (slice_string.go:7)	MOVQ	CX, ""..autotmp_2+64(SP)
	0x0075 00117 (slice_string.go:7)	XORPS	X0, X0
	0x0078 00120 (slice_string.go:7)	MOVUPS	X0, ""..autotmp_1+72(SP)
	0x007d 00125 (slice_string.go:7)	LEAQ	type.string(SB), AX
	0x0084 00132 (slice_string.go:7)	MOVQ	AX, (SP)
	0x0088 00136 (slice_string.go:7)	LEAQ	""..autotmp_2+56(SP), AX
	0x008d 00141 (slice_string.go:7)	MOVQ	AX, 8(SP)
	0x0092 00146 (slice_string.go:7)	PCDATA	$0, $1
	0x0092 00146 (slice_string.go:7)	CALL	runtime.convT2Estring(SB)
	0x0097 00151 (slice_string.go:7)	MOVQ	24(SP), AX
	0x009c 00156 (slice_string.go:7)	MOVQ	16(SP), CX
	0x00a1 00161 (slice_string.go:7)	MOVQ	CX, ""..autotmp_1+72(SP)
	0x00a6 00166 (slice_string.go:7)	MOVQ	AX, ""..autotmp_1+80(SP)
	0x00ab 00171 (slice_string.go:7)	LEAQ	""..autotmp_1+72(SP), AX
	0x00b0 00176 (slice_string.go:7)	MOVQ	AX, (SP)
	0x00b4 00180 (slice_string.go:7)	MOVQ	$1, 8(SP)
	0x00bd 00189 (slice_string.go:7)	MOVQ	$1, 16(SP)
	0x00c6 00198 (slice_string.go:7)	PCDATA	$0, $1
	0x00c6 00198 (slice_string.go:7)	CALL	fmt.Println(SB)
	0x00cb 00203 (slice_string.go:8)	MOVQ	88(SP), BP
	0x00d0 00208 (slice_string.go:8)	ADDQ	$96, SP
	0x00d4 00212 (slice_string.go:8)	RET
	0x00d5 00213 (slice_string.go:8)	NOP
	0x00d5 00213 (slice_string.go:5)	PCDATA	$0, $-1
	0x00d5 00213 (slice_string.go:5)	CALL	runtime.morestack_noctxt(SB)
	0x00da 00218 (slice_string.go:5)	JMP	0
	0x0000 64 48 8b 0c 25 00 00 00 00 48 3b 61 10 0f 86 c2  dH..%....H;a....
	0x0010 00 00 00 48 83 ec 60 48 89 6c 24 58 48 8d 6c 24  ...H..`H.l$XH.l$
	0x0020 58 0f b7 05 00 00 00 00 0f b6 0d 00 00 00 00 66  X..............f
	0x0030 89 44 24 35 88 4c 24 37 48 c7 04 24 00 00 00 00  .D$5.L$7H..$....
	0x0040 48 8d 44 24 35 48 89 44 24 08 48 c7 44 24 10 03  H.D$5H.D$.H.D$..
	0x0050 00 00 00 48 c7 44 24 18 03 00 00 00 e8 00 00 00  ...H.D$.........
	0x0060 00 48 8b 44 24 20 48 8b 4c 24 28 48 89 44 24 38  .H.D$ H.L$(H.D$8
	0x0070 48 89 4c 24 40 0f 57 c0 0f 11 44 24 48 48 8d 05  H.L$@.W...D$HH..
	0x0080 00 00 00 00 48 89 04 24 48 8d 44 24 38 48 89 44  ....H..$H.D$8H.D
	0x0090 24 08 e8 00 00 00 00 48 8b 44 24 18 48 8b 4c 24  $......H.D$.H.L$
	0x00a0 10 48 89 4c 24 48 48 89 44 24 50 48 8d 44 24 48  .H.L$HH.D$PH.D$H
	0x00b0 48 89 04 24 48 c7 44 24 08 01 00 00 00 48 c7 44  H..$H.D$.....H.D
	0x00c0 24 10 01 00 00 00 e8 00 00 00 00 48 8b 6c 24 58  $..........H.l$X
	0x00d0 48 83 c4 60 c3 e8 00 00 00 00 e9 21 ff ff ff     H..`.......!...
	rel 5+4 t=16 TLS+0
	rel 36+4 t=15 "".statictmp_0+0
	rel 43+4 t=15 "".statictmp_0+2
	rel 93+4 t=8 runtime.slicebytetostring+0
	rel 128+4 t=15 type.string+0
	rel 147+4 t=8 runtime.convT2Estring+0
	rel 199+4 t=8 fmt.Println+0
	rel 214+4 t=8 runtime.morestack_noctxt+0
"".main STEXT size=48 args=0x0 locals=0x8
	0x0000 00000 (slice_string.go:10)	TEXT	"".main(SB), $8-0
	0x0000 00000 (slice_string.go:10)	MOVQ	(TLS), CX
	0x0009 00009 (slice_string.go:10)	CMPQ	SP, 16(CX)
	0x000d 00013 (slice_string.go:10)	JLS	41
	0x000f 00015 (slice_string.go:10)	SUBQ	$8, SP
	0x0013 00019 (slice_string.go:10)	MOVQ	BP, (SP)
	0x0017 00023 (slice_string.go:10)	LEAQ	(SP), BP
	0x001b 00027 (slice_string.go:10)	FUNCDATA	$0, gclocals·33cdeccccebe80329f1fdbee7f5874cb(SB)
	0x001b 00027 (slice_string.go:10)	FUNCDATA	$1, gclocals·33cdeccccebe80329f1fdbee7f5874cb(SB)
	0x001b 00027 (slice_string.go:11)	PCDATA	$0, $0
	0x001b 00027 (slice_string.go:11)	CALL	"".slice2String(SB)
	0x0020 00032 (slice_string.go:12)	MOVQ	(SP), BP
	0x0024 00036 (slice_string.go:12)	ADDQ	$8, SP
	0x0028 00040 (slice_string.go:12)	RET
	0x0029 00041 (slice_string.go:12)	NOP
	0x0029 00041 (slice_string.go:10)	PCDATA	$0, $-1
	0x0029 00041 (slice_string.go:10)	CALL	runtime.morestack_noctxt(SB)
	0x002e 00046 (slice_string.go:10)	JMP	0
	0x0000 64 48 8b 0c 25 00 00 00 00 48 3b 61 10 76 1a 48  dH..%....H;a.v.H
	0x0010 83 ec 08 48 89 2c 24 48 8d 2c 24 e8 00 00 00 00  ...H.,$H.,$.....
	0x0020 48 8b 2c 24 48 83 c4 08 c3 e8 00 00 00 00 eb d0  H.,$H...........
	rel 5+4 t=16 TLS+0
	rel 28+4 t=8 "".slice2String+0
	rel 42+4 t=8 runtime.morestack_noctxt+0
"".init STEXT size=91 args=0x0 locals=0x8
	0x0000 00000 (<autogenerated>:1)	TEXT	"".init(SB), $8-0
	0x0000 00000 (<autogenerated>:1)	MOVQ	(TLS), CX
	0x0009 00009 (<autogenerated>:1)	CMPQ	SP, 16(CX)
	0x000d 00013 (<autogenerated>:1)	JLS	84
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
	0x0038 00056 (<autogenerated>:1)	MOVB	$1, "".initdone·(SB)
	0x003f 00063 (<autogenerated>:1)	PCDATA	$0, $0
	0x003f 00063 (<autogenerated>:1)	CALL	fmt.init(SB)
	0x0044 00068 (<autogenerated>:1)	MOVB	$2, "".initdone·(SB)
	0x004b 00075 (<autogenerated>:1)	MOVQ	(SP), BP
	0x004f 00079 (<autogenerated>:1)	ADDQ	$8, SP
	0x0053 00083 (<autogenerated>:1)	RET
	0x0054 00084 (<autogenerated>:1)	NOP
	0x0054 00084 (<autogenerated>:1)	PCDATA	$0, $-1
	0x0054 00084 (<autogenerated>:1)	CALL	runtime.morestack_noctxt(SB)
	0x0059 00089 (<autogenerated>:1)	JMP	0
	0x0000 64 48 8b 0c 25 00 00 00 00 48 3b 61 10 76 45 48  dH..%....H;a.vEH
	0x0010 83 ec 08 48 89 2c 24 48 8d 2c 24 0f b6 05 00 00  ...H.,$H.,$.....
	0x0020 00 00 3c 01 76 09 48 8b 2c 24 48 83 c4 08 c3 75  ..<.v.H.,$H....u
	0x0030 07 e8 00 00 00 00 0f 0b c6 05 00 00 00 00 01 e8  ................
	0x0040 00 00 00 00 c6 05 00 00 00 00 02 48 8b 2c 24 48  ...........H.,$H
	0x0050 83 c4 08 c3 e8 00 00 00 00 eb a5                 ...........
	rel 5+4 t=16 TLS+0
	rel 30+4 t=15 "".initdone·+0
	rel 50+4 t=8 runtime.throwinit+0
	rel 58+4 t=15 "".initdone·+-1
	rel 64+4 t=8 fmt.init+0
	rel 70+4 t=15 "".initdone·+-1
	rel 85+4 t=8 runtime.morestack_noctxt+0
go.info."".slice2String SDWARFINFO size=41
	0x0000 02 22 22 2e 73 6c 69 63 65 32 53 74 72 69 6e 67  ."".slice2String
	0x0010 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0020 00 01 9c 00 00 00 00 01 00                       .........
	rel 17+8 t=1 "".slice2String+0
	rel 25+8 t=1 "".slice2String+223
	rel 35+4 t=29 gofile../home/wenhuijia/go/src/golang/src/interview/slice_string.go+0
go.range."".slice2String SDWARFRANGE size=0
go.info."".main SDWARFINFO size=33
	0x0000 02 22 22 2e 6d 61 69 6e 00 00 00 00 00 00 00 00  ."".main........
	0x0010 00 00 00 00 00 00 00 00 00 01 9c 00 00 00 00 01  ................
	0x0020 00                                               .
	rel 9+8 t=1 "".main+0
	rel 17+8 t=1 "".main+48
	rel 27+4 t=29 gofile../home/wenhuijia/go/src/golang/src/interview/slice_string.go+0
go.range."".main SDWARFRANGE size=0
go.info."".init SDWARFINFO size=33
	0x0000 02 22 22 2e 69 6e 69 74 00 00 00 00 00 00 00 00  ."".init........
	0x0010 00 00 00 00 00 00 00 00 00 01 9c 00 00 00 00 01  ................
	0x0020 00                                               .
	rel 9+8 t=1 "".init+0
	rel 17+8 t=1 "".init+91
	rel 27+4 t=29 gofile..<autogenerated>+0
go.range."".init SDWARFRANGE size=0
"".statictmp_0 SRODATA size=3
	0x0000 31 32 33                                         123
"".initdone· SNOPTRBSS size=1
runtime.gcbits.01 SRODATA dupok size=1
	0x0000 01                                               .
type..namedata.*interface {}- SRODATA dupok size=16
	0x0000 00 00 0d 2a 69 6e 74 65 72 66 61 63 65 20 7b 7d  ...*interface {}
type.*interface {} SRODATA dupok size=56
	0x0000 08 00 00 00 00 00 00 00 08 00 00 00 00 00 00 00  ................
	0x0010 4f 0f 96 9d 00 08 08 36 00 00 00 00 00 00 00 00  O......6........
	0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0030 00 00 00 00 00 00 00 00                          ........
	rel 24+8 t=1 runtime.algarray+80
	rel 32+8 t=1 runtime.gcbits.01+0
	rel 40+4 t=5 type..namedata.*interface {}-+0
	rel 48+8 t=1 type.interface {}+0
runtime.gcbits.03 SRODATA dupok size=1
	0x0000 03                                               .
type.interface {} SRODATA dupok size=80
	0x0000 10 00 00 00 00 00 00 00 10 00 00 00 00 00 00 00  ................
	0x0010 e7 57 a0 18 02 08 08 14 00 00 00 00 00 00 00 00  .W..............
	0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0030 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0040 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	rel 24+8 t=1 runtime.algarray+144
	rel 32+8 t=1 runtime.gcbits.03+0
	rel 40+4 t=5 type..namedata.*interface {}-+0
	rel 44+4 t=6 type.*interface {}+0
	rel 56+8 t=1 type.interface {}+80
type..namedata.*[]interface {}- SRODATA dupok size=18
	0x0000 00 00 0f 2a 5b 5d 69 6e 74 65 72 66 61 63 65 20  ...*[]interface 
	0x0010 7b 7d                                            {}
type.*[]interface {} SRODATA dupok size=56
	0x0000 08 00 00 00 00 00 00 00 08 00 00 00 00 00 00 00  ................
	0x0010 f3 04 9a e7 00 08 08 36 00 00 00 00 00 00 00 00  .......6........
	0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0030 00 00 00 00 00 00 00 00                          ........
	rel 24+8 t=1 runtime.algarray+80
	rel 32+8 t=1 runtime.gcbits.01+0
	rel 40+4 t=5 type..namedata.*[]interface {}-+0
	rel 48+8 t=1 type.[]interface {}+0
type.[]interface {} SRODATA dupok size=56
	0x0000 18 00 00 00 00 00 00 00 08 00 00 00 00 00 00 00  ................
	0x0010 70 93 ea 2f 02 08 08 17 00 00 00 00 00 00 00 00  p../............
	0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0030 00 00 00 00 00 00 00 00                          ........
	rel 24+8 t=1 runtime.algarray+0
	rel 32+8 t=1 runtime.gcbits.01+0
	rel 40+4 t=5 type..namedata.*[]interface {}-+0
	rel 44+4 t=6 type.*[]interface {}+0
	rel 48+8 t=1 type.interface {}+0
type..namedata.*[1]interface {}- SRODATA dupok size=19
	0x0000 00 00 10 2a 5b 31 5d 69 6e 74 65 72 66 61 63 65  ...*[1]interface
	0x0010 20 7b 7d                                          {}
type.*[1]interface {} SRODATA dupok size=56
	0x0000 08 00 00 00 00 00 00 00 08 00 00 00 00 00 00 00  ................
	0x0010 bf 03 a8 35 00 08 08 36 00 00 00 00 00 00 00 00  ...5...6........
	0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0030 00 00 00 00 00 00 00 00                          ........
	rel 24+8 t=1 runtime.algarray+80
	rel 32+8 t=1 runtime.gcbits.01+0
	rel 40+4 t=5 type..namedata.*[1]interface {}-+0
	rel 48+8 t=1 type.[1]interface {}+0
type.[1]interface {} SRODATA dupok size=72
	0x0000 10 00 00 00 00 00 00 00 10 00 00 00 00 00 00 00  ................
	0x0010 50 91 5b fa 02 08 08 11 00 00 00 00 00 00 00 00  P.[.............
	0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0030 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0040 01 00 00 00 00 00 00 00                          ........
	rel 24+8 t=1 runtime.algarray+144
	rel 32+8 t=1 runtime.gcbits.03+0
	rel 40+4 t=5 type..namedata.*[1]interface {}-+0
	rel 44+4 t=6 type.*[1]interface {}+0
	rel 48+8 t=1 type.interface {}+0
	rel 56+8 t=1 type.[]interface {}+0
type..namedata.*[]uint8- SRODATA dupok size=11
	0x0000 00 00 08 2a 5b 5d 75 69 6e 74 38                 ...*[]uint8
type.*[]uint8 SRODATA dupok size=56
	0x0000 08 00 00 00 00 00 00 00 08 00 00 00 00 00 00 00  ................
	0x0010 a5 8e d0 69 00 08 08 36 00 00 00 00 00 00 00 00  ...i...6........
	0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0030 00 00 00 00 00 00 00 00                          ........
	rel 24+8 t=1 runtime.algarray+80
	rel 32+8 t=1 runtime.gcbits.01+0
	rel 40+4 t=5 type..namedata.*[]uint8-+0
	rel 48+8 t=1 type.[]uint8+0
type.[]uint8 SRODATA dupok size=56
	0x0000 18 00 00 00 00 00 00 00 08 00 00 00 00 00 00 00  ................
	0x0010 df 7e 2e 38 02 08 08 17 00 00 00 00 00 00 00 00  .~.8............
	0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0030 00 00 00 00 00 00 00 00                          ........
	rel 24+8 t=1 runtime.algarray+0
	rel 32+8 t=1 runtime.gcbits.01+0
	rel 40+4 t=5 type..namedata.*[]uint8-+0
	rel 44+4 t=6 type.*[]uint8+0
	rel 48+8 t=1 type.uint8+0
type..hashfunc3 SRODATA dupok size=16
	0x0000 00 00 00 00 00 00 00 00 03 00 00 00 00 00 00 00  ................
	rel 0+8 t=1 runtime.memhash_varlen+0
type..eqfunc3 SRODATA dupok size=16
	0x0000 00 00 00 00 00 00 00 00 03 00 00 00 00 00 00 00  ................
	rel 0+8 t=1 runtime.memequal_varlen+0
type..alg3 SRODATA dupok size=16
	0x0000 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	rel 0+8 t=1 type..hashfunc3+0
	rel 8+8 t=1 type..eqfunc3+0
type..namedata.*[3]uint8- SRODATA dupok size=12
	0x0000 00 00 09 2a 5b 33 5d 75 69 6e 74 38              ...*[3]uint8
type.*[3]uint8 SRODATA dupok size=56
	0x0000 08 00 00 00 00 00 00 00 08 00 00 00 00 00 00 00  ................
	0x0010 69 04 66 6c 00 08 08 36 00 00 00 00 00 00 00 00  i.fl...6........
	0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0030 00 00 00 00 00 00 00 00                          ........
	rel 24+8 t=1 runtime.algarray+80
	rel 32+8 t=1 runtime.gcbits.01+0
	rel 40+4 t=5 type..namedata.*[3]uint8-+0
	rel 48+8 t=1 type.[3]uint8+0
runtime.gcbits. SRODATA dupok size=0
type.[3]uint8 SRODATA dupok size=72
	0x0000 03 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0010 c2 b9 52 dd 02 01 01 91 00 00 00 00 00 00 00 00  ..R.............
	0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0030 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0040 03 00 00 00 00 00 00 00                          ........
	rel 24+8 t=1 type..alg3+0
	rel 32+8 t=1 runtime.gcbits.+0
	rel 40+4 t=5 type..namedata.*[3]uint8-+0
	rel 44+4 t=6 type.*[3]uint8+0
	rel 48+8 t=1 type.uint8+0
	rel 56+8 t=1 type.[]uint8+0
type..importpath.fmt. SRODATA dupok size=6
	0x0000 00 00 03 66 6d 74                                ...fmt
gclocals·69c1753bd5f81501d95132d08af04464 SRODATA dupok size=8
	0x0000 02 00 00 00 00 00 00 00                          ........
gclocals·adc9a4e9eec70025782f0016681b6ae2 SRODATA dupok size=10
	0x0000 02 00 00 00 04 00 00 00 00 0d                    ..........
gclocals·33cdeccccebe80329f1fdbee7f5874cb SRODATA dupok size=8
	0x0000 01 00 00 00 00 00 00 00                          ........
