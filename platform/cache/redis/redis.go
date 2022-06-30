package redis

import (
	"fmt"
	"log"

	goredis "github.com/go-redis/redis/v8"

	"shopping-mono/pkg/configs"
)

type Redis struct {
	RDB *goredis.Client
}

func New(cfg configs.Config) (*Redis, func()) {
	rdb := goredis.NewClient(&goredis.Options{
		Addr:     fmt.Sprintf("%s:%d", cfg.Redis.Host, cfg.Redis.Port),
		Password: cfg.Redis.Password,
		DB:       cfg.Redis.DB,
	})
	log.Println("connected to redis!!!")
	cleanup := func() {
		rdb.Close()
	}
	return &Redis{
		RDB: rdb,
	}, cleanup
}
