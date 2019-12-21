#include "textflag.h"
#include "go_tls.h"

TEXT ·getg(SB),NOSPLIT,$0-8
	get_tls(AX)
	MOVQ	g(AX), AX
	MOVQ	AX, ret+0(FP)
	RET
