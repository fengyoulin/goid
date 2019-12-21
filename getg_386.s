#include "textflag.h"
#include "go_tls.h"

TEXT ·getg(SB),NOSPLIT,$0-4
	get_tls(AX)
	MOVL	g(AX), AX
	MOVL	AX, ret+0(FP)
	RET
