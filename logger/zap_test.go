package logger

import (
	"flag"
	"testing"

	"go.uber.org/zap"
)

func TestInitZap(t *testing.T) {
	InitZap("test")
	zap.L().Sugar().Debugf("this is a test log: %s", "x")
}

func TestInitZapFromArgs(t *testing.T) {
	_ = flag.Set("log.filepath", "log/ok.log")
	InitZapFromArgs()
	zap.L().Sugar().Infof("test")
}
