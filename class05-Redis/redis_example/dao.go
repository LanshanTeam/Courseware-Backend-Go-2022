package main

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"log"
)

var Rdb *redis.Client

func InitRedis() {
	Rdb = redis.NewClient(&redis.Options{
		//连接信息
		//Network:  "tcp",              //网络类型，tcp or unix，默认tcp
		Addr:     "127.0.0.1:6379", //主机名+冒号+端口，默认localhost:6379
		Password: "123456",         //密码
		DB:       0,                // redis数据库index

	})

	_, err := Rdb.Ping(context.Background()).Result()
	if err != nil {
		log.Print(err)
	}
	fmt.Println("redis 链接成功")
}
