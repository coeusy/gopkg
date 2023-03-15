package cfg

type InnerConfig struct {
	DataSource DataSource
}

type DataSource struct {
	RDS           RDSConf           `json:"rds"`
	Redis         RedisConf         `json:"redis"`
	Kafka         KafkaConf         `json:"kafka"`
	InfluxDB      InfluxDBConf      `json:"influxdb"`
	ElasticSearch ElasticSearchConf `json:"elasticsearch"`
}

type RDSConf struct {
	Connection RDSConnection `json:"connection"`
	Database   string        `json:"database"`
	Dialect    string        `json:"dialect"`
	Debug      bool          `json:"debug"`
}

type RDSConnection struct {
	Host     string `json:"host"`
	Port     int    `json:"port"`
	User     string `json:"user"`
	Password string `json:"password"`
	Charset  string `json:"charset"`
}

type RDSSetting struct {
	Verbose        bool `json:"verbose"`
	MaxConnections int  `json:"max_connections"`
	PoolRecycle    int  `json:"pool_recycle"`
}

type RedisConf struct {
	Host     string `json:"host"`
	Port     int    `json:"port"`
	Password string `json:"password"`
	DB       int    `json:"db"`
}

type KafkaConf struct {
	Servers []string `json:"servers,omitempty"`
}

type InfluxDBConf struct {
	Host     string `json:"host,omitempty"`
	Port     int    `json:"port,omitempty"`
	User     string `json:"user,omitempty"`
	Password string `json:"password,omitempty"`
}

type ElasticSearchConf struct {
	Hosts    []string `json:"hosts,omitempty"`
	User     string   `json:"user,omitempty"`
	Password string   `json:"password,omitempty"`
}
