package initialize

import (
	"common-lib/redis"
	"jaden/we-life/common"
	"time"
)

func SetUpRedis() {
	common.Redis = redis.New(redis.Config{
		Address:     common.Cfg.GetString("redis.addr"),
		Password:    common.Cfg.GetString("redis.password"),
		Database:    common.Cfg.GetInt("redis.database"),
		MaxActive:   common.Cfg.GetInt("redis.maxActive"),
		IdleTimeout: time.Duration(common.Cfg.GetInt("redis.idleTimeout")) * time.Second,
	})
}
