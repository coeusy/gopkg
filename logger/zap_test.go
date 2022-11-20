package logger

import (
	"testing"

	"go.uber.org/zap"
)

func TestInitZap(t *testing.T) {
	InitZap("test")
	zap.L().Sugar().Debugf("this is a test log: %s", "x")
}
