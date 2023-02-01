package conv

import (
	"reflect"
	"unsafe"
)

func BytesToString(bytes []byte) string {
	return *(*string)(unsafe.Pointer(&bytes))
}

func Int8sToBytes(int8s []int8) []byte {
	s := (*reflect.SliceHeader)(unsafe.Pointer(&int8s))
	b := reflect.SliceHeader{
		Data: s.Data,
		Len:  s.Len,
		Cap:  s.Len,
	}
	return *(*[]byte)(unsafe.Pointer(&b))
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
