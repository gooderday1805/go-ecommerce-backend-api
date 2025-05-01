package initialize

import (
	"context"
	"fmt"

	"github.com/gooderday1805/go-ecommerce-backend-api/global"
	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
)

func InitRedis() {
	r := global.Config.Redis
	rdb := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%v", r.Host, r.Port),
		Password: r.Password, // no password set
		DB:       0,          // use default DB
		PoolSize: 10,         // use default PoolSize
	})

	var ctx = context.Background()

	_, err := rdb.Ping(ctx).Result()
	if err != nil {
		global.Logger.Error("Initialize redis failed", zap.Error(err))
		panic("Redis connection failed")
	}

	global.Rdb = rdb
	global.Logger.Info("Initialize redis success")
}
