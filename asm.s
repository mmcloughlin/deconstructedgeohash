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
