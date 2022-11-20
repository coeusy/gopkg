package cfg

import (
	"os"
	"testing"

	jsoniter "github.com/json-iterator/go"
	"go.uber.org/zap"

	"github.com/coeusy/gopkg/logger"
)

func TestNewManager(t *testing.T) {
	logger.InitZap("../log/test")
	opt := NewOption(WithConfigList([]string{"runtime"}))
	zap.L().Sugar().Infof("%+v", opt)
	m := NewManager(opt)
	conf, _ := jsoniter.MarshalToString(m.AllSettings())
	zap.L().Sugar().Infof(conf)
}

func TestGetConfig(t *testing.T) {
	os.Args = []string{"cfg.path=../conf", "cfg.files=runtime"}
	logger.InitZap("../log/test")
	InitConfigFromArgs()
	m := GetConfig()
	conf, _ := jsoniter.MarshalToString(m.AllSettings())
	zap.L().Sugar().Infof(conf)
	zap.L().Sugar().Infof("%+v", m.GetDatasource())
}
