package cfg

import (
	"flag"
	"strings"
)

var (
	cmdManager *Manager
)
var (
	cmdOpt                = Option{}
	cmdFiles, cmdFileType string
)

func InitCMD() {
	flag.StringVar(&cmdOpt.Path, "cfg.path", "conf/", "string: path of configuration, default is conf/")
	flag.BoolVar(&cmdOpt.NeedDatasource, "cfg.datasource", true, "boolean: use datasource config or not, default is true")
	flag.StringVar(&cmdFiles, "cfg.files", "", "string: config filename seperated by comma")
	flag.StringVar(&cmdFileType, "cfg.type", "yml", "string: config file type, such as json, yml, yaml, default is yml")
}

func InitConfigFromCMD() {
	if len(cmdFiles) > 0 {
		cmdOpt.ConfigList = strings.Split(cmdFiles, ",")
	}
	cmdOpt.FileType = FileType(cmdFileType)
	cmdManager = NewManager(cmdOpt)
}

func GetConfig() *Manager {
	return cmdManager
}

func GetDatasource() DataSource {
	if cmdManager == nil {
		return DataSource{}
	}
	return cmdManager.datasource
}

func GetKafka() KafkaConf {
	return GetDatasource().Kafka
}

func GetInfluxDB() InfluxDBConf {
	return GetDatasource().InfluxDB
}

func GetES() ElasticSearchConf {
	return GetDatasource().ElasticSearch
}
