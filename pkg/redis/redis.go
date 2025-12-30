package redis

import (
	"context"
	"kumiko/pkg/logger"
	"time"

	"github.com/redis/go-redis/v9"
	"github.com/spf13/viper"
)

var Client *redis.Client
var Ctx = context.Background()

func InitRedis() {
	addr := viper.GetString("redis.addr")
	password := viper.GetString("redis.password")
	db := viper.GetInt("redis.db")

	Client = redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: password,
		DB:       db,
	})

	if err := Client.Ping(Ctx).Err(); err != nil {
		logger.StdError("Redis连接失败: %v", err)
	}
}

// Set 封装，带统一错误处理
func Set(key string, value interface{}, expiration time.Duration) bool {
	err := Client.Set(Ctx, key, value, expiration).Err()
	if err != nil {
		logger.StdError("Redis Set 错误: %v")
		return false
	}
	return true
}

// Get 封装，带统一错误处理
func Get(key string) (string, bool) {
	val, err := Client.Get(Ctx, key).Result()
	if err == redis.Nil {
		return "", false // key不存在
	} else if err != nil {
		logger.StdError("Redis Get 错误: %v", err)
		return "", false
	}
	return val, true
}

// Del 封装
func Del(key string) bool {
	err := Client.Del(Ctx, key).Err()
	if err != nil {
		logger.StdError("Redis Del 错误: %v", err)
		return false
	}
	return true
}
