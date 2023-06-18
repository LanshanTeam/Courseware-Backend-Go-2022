package main

import (
	"context"
	"github.com/go-redis/redis/v8"
)

type RedisSet struct {
	Id      int64
	Object  string
	Conn    *redis.Client
	Context context.Context
}

func NewRedisSet(context context.Context, Objet string, Id int64, Conn *redis.Client) *RedisSet {
	return &RedisSet{
		Id:      Id,
		Object:  Objet,
		Conn:    Conn,
		Context: context,
	}
}

var x map[string]string

func (rs *RedisSet) SetNumberByUid() (int64, error) {
	val, err := rs.Conn.SCard(rs.Context, rs.Object).Result()
	rs.Conn.HSet(rs.Context, "hset", x)
	if err != nil {
		return val, err
	}
	return val, err
}
