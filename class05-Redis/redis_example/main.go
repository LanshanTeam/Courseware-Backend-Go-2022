package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	InitRedis()
	String()

}

func String() {
	err := SetRedisValue(context.Background(), "gocybee", "666", 24*time.Hour)
	if err != nil {
		fmt.Println(err)
	}

	val := Rdb.Get(context.Background(), "lmj")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(val)
}

func Set() {
	rs := NewRedisSet(context.Background(), "article:1", 1, Rdb)
	_, err := rs.Conn.SAdd(rs.Context, rs.Object, rs.Id).Result()
	if err != nil {
		fmt.Println(err)
	}
}
