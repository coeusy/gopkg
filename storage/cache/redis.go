package cache

import (
	"fmt"

	"github.com/go-redis/redis"

	"github.com/coeusy/gopkg/cfg"
)

func NewRedisClientFromConf(conf cfg.RedisConf) (*redis.Client, error) {
	opt := &redis.Options{
		Addr:       fmt.Sprintf("%s:%d", conf.Host, conf.Port),
		Password:   conf.Password,
		DB:         conf.DB,
		MaxRetries: 2,
	}
	cli := redis.NewClient(opt)
	_, err := cli.Ping().Result()
	if err != nil {
		return nil, err
	}
	return cli, nil
}
