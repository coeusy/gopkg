package cache

import (
	"fmt"
	"github.com/coeusy/gopkg/cfg"
	"testing"
)

func TestNewRedisClientFromConf(t *testing.T) {
	cli := NewRedisClientFromConf(cfg.RedisConf{
		Host:     "192.168.1.150",
		Port:     56379,
		Password: "*****",
		DB:       3,
	})
	fmt.Println(cli.Ping().Result())
}
