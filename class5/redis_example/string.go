package main

import (
	"context"
	"time"
)

func GetRedisValue(ctx context.Context, key string) (string, error) {
	GetKey := Rdb.Get(ctx, key)

	if GetKey.Err() != nil {
		return "", GetKey.Err()
	}
	return GetKey.Val(), nil
}

func SetRedisValue(ctx context.Context, key string, value string, expiration time.Duration) error {
	SetKV := Rdb.Set(ctx, key, value, expiration)
	return SetKV.Err()
}
