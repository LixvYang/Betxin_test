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

var r *RedisClient

func NewRedisClient(ctx context.Context) *RedisClient {
	r = &RedisClient{
		ctx: ctx,
		redisClient: redis.NewClient(&redis.Options{
			Addr:     utils.RedisHost + ":" + utils.RedisPort,
			Password: utils.RedisPassword, // no password set
			DB:       0,                   // use default DB
		},
		),
	}
	return r
}

func Get(key string) *redis.StringCmd {
	return r.redisClient.Get(r.ctx, key)
}
func Exists(key string) bool {
	return r.redisClient.Exists(r.ctx, key).Val() != 0
}

func Set(key string, value interface{}, expiration time.Duration) {
	if Exists(key) {
		r.redisClient.Expire(r.ctx, key, expiration)
		return
	}
	r.redisClient.Set(r.ctx, key, value, expiration)
}

func DelKeys(keys ...string) {
	for i := 0; i < len(keys); i++ {
		Del(keys[i])
	}
}

func Del(key string) {
	if !Exists(key) {
		return
	}
	r.redisClient.Del(r.ctx, key)
}

// func (r *RedisClient) Increment(key string) {
// 	r.redisClient.Incr(r.ctx, key)
// }

// func (r *RedisClient) SAdd(key string, members interface{}) {
// 	r.redisClient.SAdd(r.ctx, key, members)
// }

// func (r *RedisClient) SRem(key string, members interface{}) {
// 	r.redisClient.SRem(r.ctx, key, members)
// }

// func (r *RedisClient) Smembers(key string) *redis.StringSliceCmd {
// 	return r.redisClient.SMembers(r.ctx, key)
// }

// func (r *RedisClient) HSet(key string, expiration time.Duration, values ...interface{}) {
// 	if r.Exists(key) {
// 		r.redisClient.Expire(r.ctx, key, expiration)
// 		return
// 	}
// 	r.redisClient.HSet(r.ctx, key, values...)
// }

// func (r *RedisClient) HGet(key string, field ...string) *redis.SliceCmd {
// 	return r.redisClient.MGet(r.ctx, field...)
// }

// func (r *RedisClient) HMGet(key string, field ...string) *redis.SliceCmd {
// 	return r.redisClient.HMGet(r.ctx, key, field...)
// }
