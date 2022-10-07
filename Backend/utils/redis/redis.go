package redis

import (
	"betxin/utils"
	"context"
	"time"

	"github.com/go-redis/redis/v8"
)

type RedisClient struct {
	ctx         context.Context
	redisClient *redis.Client
}

func NewRedisClient(ctx context.Context) *RedisClient {
	return &RedisClient{
		ctx: ctx,
		redisClient: redis.NewClient(&redis.Options{
			Addr:     utils.RedisHost + ":" + utils.RedisPort,
			Password: utils.RedisPassword, // no password set
			DB:       0,                   // use default DB
		}),
	}
}

func (r *RedisClient) Get(key string) *redis.StringCmd {
	return r.redisClient.Get(r.ctx, key)
}
func (r *RedisClient) Exists(key string) bool {
	return r.redisClient.Exists(r.ctx, key).Val() != 0
}

func (r *RedisClient) Set(key string, value interface{}, expiration time.Duration) {
	if r.Exists(key) {
		r.redisClient.Expire(r.ctx, key, expiration)
		return
	}
	r.redisClient.Set(r.ctx, key, value, expiration)
}

func (r *RedisClient) Del(key string)  {
	r.redisClient.Del(r.ctx, key)
}

func (r *RedisClient) Increment(key string) {
	r.redisClient.Incr(r.ctx, key)
}

func (r *RedisClient) SAdd(key string, members interface{}) {
	r.redisClient.SAdd(r.ctx, key, members)
}

func (r *RedisClient) SRem(key string, members interface{}) {
	r.redisClient.SRem(r.ctx, key, members)
}

func (r *RedisClient) Smembers(key string) *redis.StringSliceCmd {
	return r.redisClient.SMembers(r.ctx, key)
}
