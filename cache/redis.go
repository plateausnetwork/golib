package cache

import (
	"context"
	"encoding/json"
	"reflect"
	"time"

	"github.com/go-redis/redis/v8"
)

type (
	implRedis struct {
		client            *redis.Client
		defaultExpiration time.Duration
	}
)

func (i implRedis) Ping(ctx context.Context) error {
	return i.client.Ping(ctx).Err()
}

func (i implRedis) Get(ctx context.Context, key string, v interface{}) error {
	if reflect.TypeOf(v).Kind() != reflect.Ptr {
		return ErrValueIsNotPointer()
	}
	cmd := i.client.Get(ctx, key)
	if err := cmd.Err(); err != nil {
		return err
	}
	b := []byte(cmd.Val())
	return json.Unmarshal(b, v)
}

func (i implRedis) Set(ctx context.Context, key string, v interface{}, expire time.Duration) error {
	bytes, err := json.Marshal(v)
	if err != nil {
		return err
	}
	if expire <= 0 {
		expire = i.defaultExpiration
	}
	return i.client.Set(ctx, key, bytes, expire).Err()
}

func (i implRedis) Del(ctx context.Context, key string) error {
	return i.client.Del(ctx, key).Err()
}
