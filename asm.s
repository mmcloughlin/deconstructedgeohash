#include "textflag.h"

// func QuantizeLatAsm(lat float64) uint32
TEXT ·QuantizeLatAsm(SB), NOSPLIT, $0
	MOVSD lat+0(FP), X0

	MULSD $(0.005555555555555556), X0
	ADDSD $(1.5), X0
	MOVQ  X0, AX
	SHRQ  $20, AX

	MOVL AX, ret+8(FP)
	RET

// func InterleaveAsm(x, y uint32) uint64
TEXT ·InterleaveAsm(SB), NOSPLIT, $0
	MOVL x+0(FP), AX
	MOVL y+4(FP), BX

	MOVQ  $0x5555555555555555, CX
	PDEPQ CX, AX, AX
	PDEPQ CX, BX, BX

	SHLQ $1, BX
	XORQ BX, AX

	MOVQ AX, ret+8(FP)
	RET

// func EncodeIntAsm(lat, lng float64) uint64
TEXT ·EncodeIntAsm(SB), NOSPLIT, $0
#define LATF	X0
#define LATI	R8
#define LNGF	X1
#define LNGI	R9
#define TEMP	R10
#define GHSH	R11
#define MASK	BX

	MOVSD lat+0(FP), LATF
	MOVSD lng+8(FP), LNGF

	MOVQ $0x5555555555555555, MASK

	MULSD $(0.005555555555555556), LATF
	ADDSD $(1.5), LATF

	MULSD $(0.002777777777777778), LNGF
	ADDSD $(1.5), LNGF

	MOVQ LNGF, LNGI
	SHRQ $20, LNGI

	MOVQ  LATF, LATI
	SHRQ  $20, LATI
	PDEPQ MASK, LATI, GHSH

	PDEPQ MASK, LNGI, TEMP

	SHLQ $1, TEMP
	XORQ TEMP, GHSH

	MOVQ GHSH, ret+16(FP)
	RET

#include "constants.h"

// func encodeIntAVX2(lat, lng *float64, hash *uint64)
TEXT ·encodeIntAVX2(SB), NOSPLIT, $0
	MOVQ lat+0(FP), AX
	MOVQ lng+8(FP), BX
	MOVQ hash+16(FP), CX

	//  4:	c4 e2 7d 19 05 ab 00 00 00 	vbroadcastsd	171(%rip), %ymm0
	VBROADCASTSD reciprocal180+0x00(SB), Y0

	//  d:	c5 fd 59 07 	vmulpd	(%rdi), %ymm0, %ymm0
	VMULPD (AX), Y0, Y0

	// 11:	c4 e2 7d 19 0d a6 00 00 00 	vbroadcastsd	166(%rip), %ymm1
	VBROADCASTSD onepointfive+0x00(SB), Y1

	// 1a:	c5 fd 58 c1 	vaddpd	%ymm1, %ymm0, %ymm0
	VADDPD Y1, Y0, Y0

	// // 1e:	c5 fd 73 d0 14 	vpsrlq	$20, %ymm0, %ymm0
	VPSRLQ $20, Y0, Y0

	// 23:	c4 e2 7d 19 15 9c 00 00 00 	vbroadcastsd	156(%rip), %ymm2
	VBROADCASTSD reciprocal360+0x00(SB), Y2

	// 2c:	c5 ed 59 16 	vmulpd	(%rsi), %ymm2, %ymm2
	VMULPD (BX), Y2, Y2

	// 30:	c5 ed 58 c9 	vaddpd	%ymm1, %ymm2, %ymm1
	VADDPD Y1, Y2, Y1

	// // 34:	c5 f5 73 d1 14 	vpsrlq	$20, %ymm1, %ymm1
	VPSRLQ $20, Y1, Y1

	// 39:	c5 fd 6f 15 9f 00 00 00 	vmovdqa	159(%rip), %ymm2
	VMOVDQU spreadbytes+0x00(SB), Y2

	// 41:	c4 e2 7d 00 c2 	vpshufb	%ymm2, %ymm0, %ymm0
	VPSHUFB Y2, Y0, Y0

	// 46:	c5 fd 6f 1d b2 00 00 00 	vmovdqa	178(%rip), %ymm3
	VBROADCASTSD lonibblemask+0x00(SB), Y3

	// 4e:	c5 fd db e3 	vpand	%ymm3, %ymm0, %ymm4
	VPAND Y3, Y0, Y4

	// 52:	c5 fd 6f 2d c6 00 00 00 	vmovdqa	198(%rip), %ymm5
	VMOVDQU spreadnibbleslut+0x00(SB), Y5

	// 5a:	c4 e2 55 00 e4 	vpshufb	%ymm4, %ymm5, %ymm4
	VPSHUFB Y4, Y5, Y4

	// 5f:	c4 e2 7d 59 35 68 00 00 00 	vpbroadcastq	104(%rip), %ymm6
	VBROADCASTSD hinibblemask+0x00(SB), Y6

	// 68:	c5 fd db c6 	vpand	%ymm6, %ymm0, %ymm0
	VPAND Y6, Y0, Y0

	// 6c:	c5 fd 73 d0 04 	vpsrlq	$4, %ymm0, %ymm0
	VPSRLQ $4, Y0, Y0

	// 71:	c4 e2 55 00 c0 	vpshufb	%ymm0, %ymm5, %ymm0
	VPSHUFB Y0, Y5, Y0

	// 76:	c5 fd 73 f0 08 	vpsllq	$8, %ymm0, %ymm0
	VPSLLQ $8, Y0, Y0

	// 7b:	c5 fd eb c4 	vpor	%ymm4, %ymm0, %ymm0
	VPOR Y4, Y0, Y0

	// 7f:	c4 e2 75 00 ca 	vpshufb	%ymm2, %ymm1, %ymm1
	VPSHUFB Y2, Y1, Y1

	// 84:	c5 f5 db d3 	vpand	%ymm3, %ymm1, %ymm2
	VPAND Y3, Y1, Y2

	// 88:	c4 e2 55 00 d2 	vpshufb	%ymm2, %ymm5, %ymm2
	VPSHUFB Y2, Y5, Y2

	// 8d:	c5 f5 db ce 	vpand	%ymm6, %ymm1, %ymm1
	VPAND Y6, Y1, Y1

	// 91:	c5 f5 73 d1 04 	vpsrlq	$4, %ymm1, %ymm1
	VPSRLQ $4, Y1, Y1

	// 96:	c4 e2 55 00 c9 	vpshufb	%ymm1, %ymm5, %ymm1
	VPSHUFB Y1, Y5, Y1

	// 9b:	c5 f5 73 f1 08 	vpsllq	$8, %ymm1, %ymm1
	VPSLLQ $8, Y1, Y1

	// a0:	c5 f5 eb ca 	vpor	%ymm2, %ymm1, %ymm1
	VPOR Y2, Y1, Y1

	// a4:	c5 f5 d4 c9 	vpaddq	%ymm1, %ymm1, %ymm1
	VPADDQ Y1, Y1, Y1

	// a8:	c5 fd eb c1 	vpor	%ymm1, %ymm0, %ymm0
	VPOR Y1, Y0, Y0

	// ac:	c5 fe 7f 02 	vmovdqu	%ymm0, (%rdx)
	VMOVDQU Y0, (CX)

	RET
