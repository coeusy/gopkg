package cfg

import (
	"os"
	"strings"

	jsoniter "github.com/json-iterator/go"
	"github.com/spf13/cast"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

var (
	manager *Manager
)

func InitConfigFromArgs() {
	manager = NewManagerFromArgs()
}

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

	manager := Manager{
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
			_ = jsoniter.Unmarshal(b, &manager.datasource)
		}
		manager.Viper.Set(key, conf[key])
	}
	return &manager
}

func NewManagerFromArgs() *Manager {
	optFn := make([]OptionFn, 0)
	for _, val := range os.Args {
		if !strings.Contains(val, "=") {
			continue
		}
		arg := strings.Split(val, "=")
		switch arg[0] {
		case "cfg.path":
			optFn = append(optFn, WithPath(arg[1]))
		case "cfg.datasource":
			if cast.ToBool(arg[1]) {
				optFn = append(optFn, WithoutDatasource())
			}
		case "cfg.files":
			optFn = append(optFn, WithConfigList(strings.Split(arg[1], ",")))
		case "cfg.type":
			optFn = append(optFn, WithFileType(FileType(arg[1])))
		}
	}
	return NewManager(NewOption(optFn...))
}

func GetConfig() *Manager {
	return manager
}

func GetDatasource() DataSource {
	if manager == nil {
		return DataSource{}
	}
	return manager.datasource
}

func GetKafka() KafkaConf {
	return GetDatasource().Kafka
}
