package conv

import (
	"fmt"
	"testing"
)

func TestStringToBytes(t *testing.T) {
	fmt.Println(StringToBytes("123124"))
}

func TestBytesToString(t *testing.T) {
	fmt.Println(BytesToString([]byte{1, 2, 3, 4, 5}))
}

func TestInt8sToBytes(t *testing.T) {
	fmt.Println(Int8sToBytes([]int8{1, 2, 3, 4, 5}))
}
