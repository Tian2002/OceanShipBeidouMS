package redis

import (
	"context"
	"errors"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/go-redis/redis/v8"
	"time"
)

func (repo *redisRepo) KeepListAlive(ctx context.Context, key string, aliveTime time.Duration, checkTime time.Duration, sleepTime time.Duration) {
	for {
		ttl, err := repo.cli.TTL(ctx, key).Result()
		if err != nil {
			if errors.Is(err, redis.Nil) {
				hlog.Fatalf("keep list alive error:[key:%v not exist]", key)
				return
			}
			hlog.Errorf("keep list alive,key:[%v] error:[%v]", key, err)
			continue
		}
		// 根据ttl的值输出不同的信息
		switch {
		case ttl == -2:
			hlog.Infof("keep list alive,key:[%v],key exists but no expiration time is set", key)
			return
		case ttl == -1:
			hlog.Warnf("keep list alive,key:[%v],Key has expired", key)
			time.Sleep(sleepTime)
		default:
			if ttl < checkTime {
				err := repo.cli.Expire(ctx, key, aliveTime).Err()
				if err != nil {
					hlog.Errorf("keep list alive,key:[%v],Key has expired", key)
					continue
				}
				time.Sleep(sleepTime)
			}
		}
	}
	return
}
