package redis

import (
	"OceanShipBeidouMS/biz/config"
	"context"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/go-redis/redis/v8"
)

func init() {
	initRedisClient()
}

var redisClient *redis.Client

func initRedisClient() {
	// 配置Redis连接参数
	opts, err := redis.ParseURL("redis://:@" + config.GetConfig().Redis.TCP)
	if err != nil {
		hlog.Fatalf("connect redis error:[%#v]", err)
		return
	}

	// 创建Redis客户端实例
	redisClient = redis.NewClient(opts)
	println(redisClient.Ping(context.Background()))
}

func GetClient() *redis.Client {
	return redisClient
}
