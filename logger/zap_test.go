package logger

import (
	"go.uber.org/zap"
	"os"
	"testing"
)

func TestInitZap(t *testing.T) {
	InitZap("test")
	zap.L().Sugar().Debugf("this is a test log: %s", "x")
}

func TestInitZapFromArgs(t *testing.T) {
	os.Args = []string{"log.filepath=log/ok.log"}
	InitZapFromArgs()
	zap.L().Sugar().Infof("test")
}
