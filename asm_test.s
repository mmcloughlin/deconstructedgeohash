#include "textflag.h"

// func NoopAsm(lat, lng float64) uint64
TEXT ·NoopAsm(SB), NOSPLIT, $0
	RET
