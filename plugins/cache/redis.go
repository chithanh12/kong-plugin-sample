package cache

import (
	"context"
	"github.com/redis/go-redis/v9"
	"sync"
	"time"
)

type (
	Cache interface {
		Get(key string) string
	}

	RedisConfig struct {
		Host     string `json:"host"`
		Password string `json:"password"`
		DB       int    `json:"DB"`
	}

	redisRepo struct {
		client *redis.Client
	}
)

var (
	repoInstance Cache
	doOnce       sync.Once
)

func initRedis(cfg RedisConfig) error {
	redisCache := &redisRepo{}

	redisCache.client = redis.NewClient(&redis.Options{
		Addr:     cfg.Host,
		Password: cfg.Password, // no password set
		DB:       cfg.DB,       // use default DB
	})

	ctx, cancelCtx := context.WithTimeout(context.TODO(), 1500*time.Millisecond)
	defer cancelCtx()
	_, err := redisCache.client.Ping(ctx).Result()

	repoInstance = redisCache
	return err
}

func Instance(cfg RedisConfig) Cache {
	doOnce.Do(func() {
		// init redis for the first time
		if err := initRedis(cfg); err != nil {
			panic(err)
		}
	})

	return repoInstance
}

func (r *redisRepo) Get(key string) string {
	rs, err := r.client.Get(context.Background(), key).Result()
	if err != nil {
		return ""
	}
	return rs
}
