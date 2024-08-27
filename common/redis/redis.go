package redis

import (
	"context"
	"time"

	"github.com/ceuloong/fil-admin-core/sdk/config"
	"github.com/redis/go-redis/v9"
)

func connect() *redis.Client {
	// 连接redis
	//log.Println("redis connect success " + config.CacheConfig.Redis.Addr)
	rdb := redis.NewClient(&redis.Options{
		Addr:     config.CacheConfig.Redis.Addr,
		Password: "", // 密码
		DB:       0,  // 数据库
		PoolSize: 20, // 连接池大小
	})
	return rdb
}

func SetRedis(key string, value interface{}, expiration time.Duration) error {
	rdb := connect()
	defer rdb.Close()

	ctx, cancel := context.WithTimeout(context.Background(), 500*time.Millisecond)
	defer cancel()

	// 设置值
	err := rdb.Set(ctx, key, value, expiration).Err()
	if err != nil {
		return err
	}
	return nil
}

func GetRedis(key string) (string, error) {
	rdb := connect()
	defer rdb.Close()

	ctx, cancel := context.WithTimeout(context.Background(), 500*time.Millisecond)
	defer cancel()

	// 获取值
	val, err := rdb.Get(ctx, key).Result()
	if err != nil {
		return "", err
	}
	return val, nil
}

func RPushRedis(key string, value interface{}) error {
	rdb := connect()
	defer rdb.Close()

	ctx, cancel := context.WithTimeout(context.Background(), 500*time.Millisecond)
	defer cancel()

	// 添加值
	err := rdb.RPush(ctx, key, value).Err()
	if err != nil {
		return err
	}
	return nil
}

func LPopRedis(key string) (string, error) {
	rdb := connect()
	defer rdb.Close()

	ctx, cancel := context.WithTimeout(context.Background(), 500*time.Millisecond)
	defer cancel()

	// 弹出值
	val, err := rdb.LPop(ctx, key).Result()
	if err != nil {
		return "", err
	}
	return val, nil
}

func RPopRedis(key string) (string, error) {
	rdb := connect()
	defer rdb.Close()

	ctx, cancel := context.WithTimeout(context.Background(), 500*time.Millisecond)
	defer cancel()

	// 弹出值
	val, err := rdb.RPop(ctx, key).Result()
	if err != nil {
		return "", err
	}
	return val, nil
}

func LRangeRedis(key string, start, stop int64) ([]string, error) {
	rdb := connect()
	defer rdb.Close()

	ctx, cancel := context.WithTimeout(context.Background(), 500*time.Millisecond)
	defer cancel()

	// 获取范围
	val, err := rdb.LRange(ctx, key, start, stop).Result()
	if err != nil {
		return nil, err
	}
	return val, nil
}

func LLenRedis(key string) (int64, error) {
	rdb := connect()
	defer rdb.Close()

	ctx, cancel := context.WithTimeout(context.Background(), 500*time.Millisecond)
	defer cancel()

	// 获取长度
	val, err := rdb.LLen(ctx, key).Result()
	if err != nil {
		return 0, err
	}
	return val, nil
}
