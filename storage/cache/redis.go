package cache

import (
	"fmt"

	"github.com/go-redis/redis"

	"github.com/coeusy/gopkg/cfg"
)

func NewRedisClientFromConf(conf cfg.RedisConf) *redis.Client {
	opt := &redis.Options{
		Addr:       fmt.Sprintf("%s:%d", conf.Host, conf.Port),
		Password:   conf.Password,
		DB:         conf.DB,
		MaxRetries: 2,
	}
	return redis.NewClient(opt)
}
