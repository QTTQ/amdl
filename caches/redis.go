package caches

import (
	"github.com/go-redis/redis"
)

var (
	Cache    = New()
	Addr     = "127.0.0.1:6379"
	Password = ""
	DB       = 0
)

func New() *redis.Client {
	// configTree := config.Conf.Get("redis").(*toml.Tree)
	return redis.NewClient(&redis.Options{
		Addr:     Addr,
		Password: Password, // no password set
		DB:       int(DB),  // 因为系统是64位的，所以默认的 int 型是 int64
	})
}
