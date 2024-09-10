package goid

import (
	"reflect"
	"unsafe"
)

// ID returns current goroutine's runtime ID
func ID() uint64 {
	gp := getg()
	switch kind {
	case reflect.Int64:
		return uint64(*(*int64)(unsafe.Pointer(gp + offset)))
	case reflect.Uint64:
		return *(*uint64)(unsafe.Pointer(gp + offset))
	default:
		panic("unexpected kind: " + kind.String())
	}
}
