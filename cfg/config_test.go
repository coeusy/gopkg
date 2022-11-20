package cfg

import (
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
