package logger

import (
	"flag"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var cmdFilepath string

func InitCMD() {
	flag.StringVar(&cmdFilepath, "log.filepath", "log/info.log", "string: path for log")
}

func InitZapFromCMD() {
	zap.ReplaceGlobals(zap.New(zapcore.NewTee(newFileOutput(cmdFilepath)), zap.AddCaller()))
}
