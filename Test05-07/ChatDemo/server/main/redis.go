package main

import (
	"github.com/gomodule/redigo/redis"
	"time"
)

var pool *redis.Pool

func initPool(addr string, maxIdle, maxActive int, idleTimeout time.Duration) {
	pool = &redis.Pool{
		//表示和数据库的最大连接数,0表示没有限制
		MaxActive: maxActive,
		//最大空闲数
		MaxIdle: maxIdle,
		//最大空闲时间
		IdleTimeout: idleTimeout,
		//初始化连接
		Dial: func() (redis.Conn, error) {
			return redis.Dial("tcp", addr)
		},
	}
}
