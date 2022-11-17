package conv

import (
	"bytes"
	"encoding/binary"
)

func Int32ToBytesL(i int32) []byte {
	buf := bytes.NewBuffer([]byte{})
	_ = binary.Write(buf, binary.LittleEndian, i)
	return buf.Bytes()
}
