package helper

import (
	"Service-API/config"
	"Service-API/model"
	"context"
	"fmt"
	"time"

	"github.com/goccy/go-json"
)

func InsertRedis(set model.SetDataRedis) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	client := config.NewRedisDB()
	jsonData, err := json.Marshal(set.Data)

	if err != nil {
		fmt.Println(err)
		return
	}

	err = client.Set(ctx, set.Key, jsonData, set.Exp).Err()
	if err != nil {
		fmt.Println(err)
		return
	}
}

func GetRedis[T any](key string) (cek bool, raw T) {
	var redis = config.NewRedisDB()

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	value, _ := redis.Get(ctx, key).Result()
	if value == "" {
		return false, raw
	}

	_ = json.Unmarshal([]byte(value), &raw)

	return true, raw
}
func DelRedis(key string) {
	client := config.NewRedisDB()
	ctx := context.Background()

	client.Eval(ctx, "for i, name in ipairs(redis.call('KEYS', '"+key+"')) do redis.call('expire', name, 0); end", []string{"*"})
}
