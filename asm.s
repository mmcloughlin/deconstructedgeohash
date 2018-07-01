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

// // func EncodeInt(lat, lng float64) uint64
// TEXT ·EncodeInt(SB), NOSPLIT, $0
// 	CMPB ·useAsm(SB), $1
// 	JNE  fallback
//
// #define LATF	X0
// #define LATI	R8
// #define LNGF	X1
// #define LNGI	R9
// #define TEMP	R10
// #define GHSH	R11
// #define MASK	BX
//
// 	MOVSD lat+0(FP), LATF
// 	MOVSD lng+8(FP), LNGF
//
// 	MOVQ $0x5555555555555555, MASK
//
// 	MULSD $(0.005555555555555556), LATF
// 	ADDSD $(1.5), LATF
//
// 	MULSD $(0.002777777777777778), LNGF
// 	ADDSD $(1.5), LNGF
//
// 	MOVQ LNGF, LNGI
// 	SHRQ $20, LNGI
//
// 	MOVQ  LATF, LATI
// 	SHRQ  $20, LATI
// 	PDEPQ MASK, LATI, GHSH
//
// 	PDEPQ MASK, LNGI, TEMP
//
// 	SHLQ $1, TEMP
// 	XORQ TEMP, GHSH
//
// 	MOVQ GHSH, ret+16(FP)
// 	RET
