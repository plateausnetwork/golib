//go:generate mockgen -source=${GOFILE} -package=${GOPACKAGE} -destination=${GOPACKAGE}_mock.go
package cache

import (
	"context"
	"crypto/tls"
	"time"

	"github.com/go-redis/redis/v8"
)

const (
	ProviderRedis = "redis"
)

const (
	V1 = iota + 1
	V2
)

type (
	Options struct {
		ProviderType string
		Addr         string
		Password     string
		Username     string
		DB           int
		Timeout      time.Duration
		Expiration   time.Duration
		SkipVerify   bool
	}
	Cache interface {
		Ping(ctx context.Context) error
		Get(ctx context.Context, key string, v interface{}) error
		Set(ctx context.Context, key string, v interface{}, expire time.Duration) error
		Del(ctx context.Context, key string) error
	}
)

func New(opts Options) (Cache, error) {
	switch opts.ProviderType {
	case ProviderRedis:
		redisOpts := &redis.Options{
			Addr:        opts.Addr,
			Password:    opts.Password,
			Username:    opts.Username,
			DB:          opts.DB,
			IdleTimeout: opts.Timeout,
		}

		if opts.Timeout > 0 {
			redisOpts.DialTimeout = opts.Timeout
		}

		if opts.SkipVerify {
			redisOpts.TLSConfig = &tls.Config{
				InsecureSkipVerify: true, // Set to false if you want to reject unauthorized connections
			}
		}

		implRedis := &implRedis{
			defaultExpiration: opts.Expiration,
			client:            redis.NewClient(redisOpts),
		}
		if err := implRedis.Ping(context.Background()); err != nil {
			return nil, ErrPing(err)
		}
		return implRedis, nil
	default:
		return nil, ErrInvlaidProvider(opts.ProviderType)
	}
}
