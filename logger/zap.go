package logger

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func InitZap(filename string) {
	cores := make([]zapcore.Core, 0)
	cores = append(cores, newFileOutput(filename))
	zap.ReplaceGlobals(zap.New(zapcore.NewTee(cores...), zap.AddCaller()))
}

func InitZapFromArgs() {
	fileCore := newFileOutput(fmt.Sprintf("log/info.log"))
	for _, val := range os.Args {
		if !strings.Contains(val, "=") {
			continue
		}
		arg := strings.Split(val, "=")
		switch arg[0] {
		case "log.filepath":
			fileCore = newFileOutput(arg[1])
		}
	}
	zap.ReplaceGlobals(zap.New(zapcore.NewTee(fileCore), zap.AddCaller()))
}

func newLumberJackWriter(filepath string) zapcore.WriteSyncer {
	lumberJackLogger := &lumberjack.Logger{
		Filename:   fmt.Sprintf(filepath), // 日志文件位置
		MaxSize:    100,
		MaxBackups: 7,
		MaxAge:     7,
		Compress:   false,
	}
	return zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout), zapcore.AddSync(lumberJackLogger))
}

func newFileOutput(filename string) zapcore.Core {
	encoderConfig := zap.NewDevelopmentEncoderConfig()
	encoderConfig.EncodeTime = func(t time.Time, encoder zapcore.PrimitiveArrayEncoder) {
		encoder.AppendString(t.Format("2006-01-02 15:04:05.000"))
	}
	encoder := zapcore.NewConsoleEncoder(encoderConfig)
	return zapcore.NewCore(encoder, newLumberJackWriter(filename), zapcore.DebugLevel)
}
