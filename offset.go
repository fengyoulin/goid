package goid

import (
	"github.com/fengyoulin/inspect"
	"reflect"
)

var offset uintptr

func init() {
	typ := inspect.TypeOf("runtime.g")
	if typ == nil {
		panic("runtime.g not found")
	}
	for i := 0; i < typ.NumField(); i++ {
		f := typ.Field(i)
		if f.Name == "goid" && f.Type == reflect.TypeOf(int64(0)) {
			offset = f.Offset
			break
		}
	}
	if offset == 0 {
		panic("runtime.g.goid not found")
	}
}
