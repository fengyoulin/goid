package goid

import (
	"github.com/fengyoulin/inspect"
	"reflect"
)

var (
	offset uintptr
	kind   reflect.Kind
)

func init() {
	typ := inspect.TypeOf("runtime.g")
	if typ == nil {
		panic("runtime.g not found")
	}
	for i := 0; i < typ.NumField(); i++ {
		f := typ.Field(i)
		if f.Name == "goid" {
			kind = f.Type.Kind()
			offset = f.Offset
			break
		}
	}
	if offset == 0 {
		panic("runtime.g.goid not found")
	}
	if kind != reflect.Int64 && kind != reflect.Uint64 {
		panic("runtime.g.goid is not int64 or uint64, got: " + kind.String())
	}
}
