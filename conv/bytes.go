package conv

import "unsafe"

func BytesToString(bytes []byte) string {
	return *(*string)(unsafe.Pointer(&bytes))
}
