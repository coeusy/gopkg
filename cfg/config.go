package cfg

import (
	jsoniter "github.com/json-iterator/go"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

type Manager struct {
	*viper.Viper
	datasource DataSource
}

func (m *Manager) GetDatasource() DataSource {
	if m == nil {
		return DataSource{}
	}
	return m.datasource
}

type configType string

const (
	configTypeDatasource configType = "datasource"
)

func NewManager(options ...Option) *Manager {
	opt := NewDefaultOption()
	if len(options) > 0 {
		opt = options[0]
	}
	path := opt.Path

	m := Manager{
		Viper:      viper.New(),
		datasource: DataSource{},
	}
	zap.L().Sugar().Infof("cfg path=%s", opt.Path)
	reader := viper.New()
	reader.AddConfigPath(path)
	reader.SetConfigType(string(opt.FileType))
	if opt.NeedDatasource {
		opt.ConfigList = append(opt.ConfigList, string(configTypeDatasource))
	}
	for _, key := range opt.ConfigList {
		zap.L().Sugar().Infof("load cfg from %s", key)
		reader.SetConfigName(key)
		err := reader.ReadInConfig()
		if err != nil {
			zap.S().Errorf("load cfg from %s failed: %s", key, err)
		}
		conf := map[string]interface{}{}
		if reader.IsSet(key) {
			conf[key] = reader.Get(key)
		} else {
			conf[key] = reader.AllSettings()
		}
		switch key {
		case string(configTypeDatasource):
			b, _ := jsoniter.Marshal(conf[key])
			_ = jsoniter.Unmarshal(b, &m.datasource)
		}
		m.Viper.Set(key, conf[key])
	}
	return &m
}
