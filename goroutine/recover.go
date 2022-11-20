package goroutine

import (
	"runtime"

	"go.uber.org/zap"

	"github.com/coeusy/gopkg/conv"
)

func GoRecover() {
	if err := recover(); err != nil {
		var buf [4096]byte
		n := runtime.Stack(buf[:], false)
		zap.S().Errorf("recover err=%+v, stack=%s", err, conv.BytesToString(buf[:n]))
	}
}
