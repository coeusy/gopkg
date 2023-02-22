package cfg

type InnerConfig struct {
	DataSource DataSource
}

type DataSource struct {
	RDS   RDSConf   `json:"rds"`
	Redis RedisConf `json:"cache"`
	Kafka KafkaConf `json:"kafka"`
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
	Producer KProducerConf `json:"producer"`
}

type KProducerConf struct {
	Servers []string `json:"servers,omitempty"`
}
