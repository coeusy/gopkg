package logger

import (
	"go.uber.org/zap"
	"testing"
)

func TestInitZap(t *testing.T) {
	InitZap("test")
	zap.L().Sugar().Debugf("this is a test log: %s", "x")
}
