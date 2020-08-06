package utils

import (
	"github.com/astaxie/beego"
	red "github.com/gomodule/redigo/redis"
	"time"
)

type Redis struct {
	pool *red.Pool
}

var redis *Redis

func InitRedis() *red.Pool {
	host := beego.AppConfig.String("redis_host")
	password := beego.AppConfig.String("redis_password")
	redis = new(Redis)
	redis.pool = &red.Pool{
		MaxIdle:     256,
		MaxActive:   0,
		IdleTimeout: time.Duration(120),
		Dial: func() (red.Conn, error) {
			return red.Dial(
				"tcp",
				host,
				red.DialPassword(password),
				red.DialReadTimeout(time.Duration(1000)*time.Millisecond),
				red.DialWriteTimeout(time.Duration(1000)*time.Millisecond),
				red.DialConnectTimeout(time.Duration(1000)*time.Millisecond),
				red.DialDatabase(0),
				//red.DialPassword(""),
			)
		},
	}
	return redis.pool
}
