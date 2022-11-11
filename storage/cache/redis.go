package cache

import (
	"fmt"

	"github.com/go-redis/redis"

	"github.com/coeusy/gopkg/cfg"
)

func ConvConfToOpt(conf cfg.RedisConf) *redis.Options {
	return &redis.Options{
		Addr:       fmt.Sprintf("%s:%d", conf.Host, conf.Port),
		Password:   conf.Password,
		DB:         conf.DB,
		MaxRetries: 2,
	}
}
