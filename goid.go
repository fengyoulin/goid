package goid

import "unsafe"

// ID returns current goroutine's runtime ID
func ID() int64 {
	gp := getg()
	return *(*int64)(unsafe.Pointer(gp + offset))
}
