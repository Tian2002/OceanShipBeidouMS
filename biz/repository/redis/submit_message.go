package redis

import (
	"context"
	"errors"
	"fmt"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/go-redis/redis/v8"
	"time"
)

type SubmitMessageKey interface {
	getSubmitMessageKey(messageID int64) string
	CheckSubmitMessage(ctx context.Context, messageID int64) (bool, error)
	TrySubmitMessage(ctx context.Context, messageID int64) (bool, error)
}

func (repo *redisRepo) getSubmitMessageKey(messageID int64) string {
	return fmt.Sprintf("submitMessage:%d", messageID)
}

func (repo *redisRepo) CheckSubmitMessage(ctx context.Context, messageID int64) (bool, error) {
	err := repo.cli.Get(ctx, repo.getSubmitMessageKey(messageID)).Err()
	if err != nil {
		//如果没有就证明没提交过
		if errors.Is(err, redis.Nil) {
			return false, nil
		}
		hlog.Errorf("check submit message error:[%v]", err)
		return false, err
	}
	return true, nil
}

func (repo *redisRepo) TrySubmitMessage(ctx context.Context, messageID int64) (bool, error) {
	result, err := repo.cli.SetNX(ctx, repo.getSubmitMessageKey(messageID), time.Now().String(), time.Hour*24*3).Result()
	if err != nil {
		hlog.Errorf("try submit message error:[%v]", err)
		return false, err
	}
	return result, nil
}
