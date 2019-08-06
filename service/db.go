package service

import (
	"OnlineGobangBattle/config"
	"github.com/gomodule/redigo/redis"
	"time"
)

var RedisClient *redis.Pool

func ConnectRedis() {
	maxIdle := config.Conf.RedisMaxIdle
	maxActive := config.Conf.RedisMaxIdle
	RedisClient = &redis.Pool{
		MaxIdle:     maxIdle,
		MaxActive:   maxActive,
		IdleTimeout: 10 * time.Second,
		Wait:        true,
		Dial: func() (conn redis.Conn, e error) {
			con, err := redis.Dial("tcp", config.Conf.RedisHost,
				redis.DialPassword(config.Conf.Redispwd),
				redis.DialDatabase(config.Conf.RedisDb),
				redis.DialConnectTimeout(10*time.Second),
				redis.DialReadTimeout(10*time.Second),
				redis.DialWriteTimeout(10*time.Second))
			if err != nil {
				return nil, err
			}
			return con, nil
		},
	}
}
