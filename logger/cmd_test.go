package logger

import (
	"flag"
	"fmt"
	"testing"

	"go.uber.org/zap"
)

func TestInitZapFromCMD(t *testing.T) {
	InitCMD()
	fmt.Println(flag.Set("log.filepath", "log/ok.log"))
	flag.Parse()
	InitZapFromCMD()
	zap.L().Sugar().Infof("test")
}
