package conv

import (
	"reflect"
	"unsafe"
)

func BytesToString(bytes []byte) string {
	return *(*string)(unsafe.Pointer(&bytes))
}

func StringToBytes(str string) []byte {
	s := (*reflect.StringHeader)(unsafe.Pointer(&str))
	b := reflect.SliceHeader{
		Data: s.Data,
		Len:  s.Len,
		Cap:  s.Len,
	}
	return *(*[]byte)(unsafe.Pointer(&b))
}
