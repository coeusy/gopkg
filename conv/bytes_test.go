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
