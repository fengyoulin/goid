package goid

// G returns the address of current g
func G() uintptr {
	return getg()
}

func getg() uintptr
