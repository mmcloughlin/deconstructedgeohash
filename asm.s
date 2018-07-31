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

// func EncodeIntSimd(lat, lng []float64, hash []uint64)
TEXT ·EncodeIntSimd(SB), NOSPLIT, $0
	MOVQ lat+0(FP), AX
	MOVQ lng+24(FP), BX
	MOVQ hash+48(FP), CX

	VBROADCASTSD reciprocal180<>+0x00(SB), Y0
	VMULPD       (AX), Y0, Y0
	VBROADCASTSD onepointfive<>+0x00(SB), Y1
	VADDPD       Y1, Y0, Y0
	VPSRLQ       $20, Y0, Y0
	VBROADCASTSD reciprocal360<>+0x00(SB), Y2
	VMULPD       (BX), Y2, Y2
	VADDPD       Y1, Y2, Y1
	VPSRLQ       $20, Y1, Y1
	VMOVDQU      spreadbyte<>+0x00(SB), Y2
	VPSHUFB      Y2, Y0, Y0
	VBROADCASTSD lonibblemask<>+0x00(SB), Y3
	VPAND        Y3, Y0, Y4
	VMOVDQU      spreadnibblelut<>+0x00(SB), Y5
	VPSHUFB      Y4, Y5, Y4
	VBROADCASTSD hinibblemask<>+0x00(SB), Y6
	VPAND        Y6, Y0, Y0
	VPSRLQ       $4, Y0, Y0
	VPSHUFB      Y0, Y5, Y0
	VPSLLQ       $8, Y0, Y0
	VPOR         Y4, Y0, Y0
	VPSHUFB      Y2, Y1, Y1
	VPAND        Y3, Y1, Y2
	VPSHUFB      Y2, Y5, Y2
	VPAND        Y6, Y1, Y1
	VPSRLQ       $4, Y1, Y1
	VPSHUFB      Y1, Y5, Y1
	VPSLLQ       $8, Y1, Y1
	VPOR         Y2, Y1, Y1
	VPADDQ       Y1, Y1, Y1
	VPOR         Y1, Y0, Y0
	VMOVDQU      Y0, (CX)

	RET
