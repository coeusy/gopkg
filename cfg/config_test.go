package cfg

import (
	"testing"

	jsoniter "github.com/json-iterator/go"
	"go.uber.org/zap"

	"github.com/coeusy/gopkg/logger"
)

func TestNewManager(t *testing.T) {
	logger.InitZap("test")
	opt := NewOption(WithConfigList([]string{"runtime"}))
	zap.L().Sugar().Infof("%+v", opt)
	m := NewManager(opt)
	conf, _ := jsoniter.MarshalToString(m.AllSettings())
	zap.L().Sugar().Infof(conf)
}

func TestGetConfig(t *testing.T) {
	logger.InitZap("test")
	InitConfigFromArgs()
	m := GetConfig()
	conf, _ := jsoniter.MarshalToString(m.AllSettings())
	zap.L().Sugar().Infof(conf)
}
