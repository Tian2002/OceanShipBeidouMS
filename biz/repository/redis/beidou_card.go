package redis

import (
	"OceanShipBeidouMS/biz/common"
	"context"
	"errors"
	"fmt"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/go-redis/redis/v8"
)

type BeidouCard interface {
	getSetSendMessageCardListKey(station int) string
	getSetSendMessageCardListUrgentKey(station int) string
	SetSendMessageCard(ctx context.Context, station int, cards []string) error
	GetSendMessageCard(ctx context.Context, station int) (string, error)
	GetSendMessageCardUrgent(ctx context.Context, station int) (string, error)
	ReturnSendMessageCard(ctx context.Context, station int, card string) error
}

func (repo *redisRepo) getSetSendMessageCardListKey(station int) string {
	return fmt.Sprintf("sendMessageCardListStation:%d", station)
}

func (repo *redisRepo) getSetSendMessageCardListUrgentKey(station int) string {
	return fmt.Sprintf("sendMessageCardListNumberStation:%d", station)
}

func (repo *redisRepo) SetSendMessageCard(ctx context.Context, station int, cards []string) error {
	var cs []interface{}
	for _, card := range cards {
		cs = append(cs, card)
	}
	pipe := repo.cli.Pipeline()
	pipe.Del(ctx, repo.getSetSendMessageCardListKey(station), repo.getSetSendMessageCardListUrgentKey(station))
	pipe.LPush(ctx, repo.getSetSendMessageCardListKey(station), cs[:common.UrgentCartNumber])
	pipe.LPush(ctx, repo.getSetSendMessageCardListUrgentKey(station), cs[common.UrgentCartNumber:])
	_, err := pipe.Exec(ctx)
	if err != nil {
		hlog.Fatalf("set send message card station:[%v] error:[%v]", station, err)
		return err
	}
	return nil
}

func (repo *redisRepo) GetSendMessageCard(ctx context.Context, station int) (string, error) {
	result, err := repo.cli.RPop(ctx, repo.getSetSendMessageCardListKey(station)).Result()
	if err != nil {
		if errors.Is(err, redis.Nil) {
			hlog.Infof("get send message card info:nil")
			return "", nil
		}
		hlog.Errorf("get send message card error:[%v]", err)
		return "", err
	}
	return result, nil
}
func (repo *redisRepo) GetSendMessageCardUrgent(ctx context.Context, station int) (string, error) {
	card, err := repo.GetSendMessageCard(ctx, station)
	if err != nil {
		return "", err
	}
	if card != "" {
		return card, nil
	}
	result, err := repo.cli.RPop(ctx, repo.getSetSendMessageCardListUrgentKey(station)).Result()
	if err != nil {
		if errors.Is(err, redis.Nil) {
			hlog.Infof("get send message urgent card info:nil")
			return "", nil
		}
		hlog.Errorf("get send message urgent card error:[%v]", err)
		return "", err
	}
	return result, nil
}

func (repo *redisRepo) ReturnSendMessageCard(ctx context.Context, station int, card string) error {
	len, err := repo.cli.LLen(ctx, repo.getSetSendMessageCardListUrgentKey(station)).Result()
	if err != nil {
		hlog.Errorf("return send message card error:[%v]", err)
		return err
	}
	if len < common.UrgentCartNumber {
		err = repo.cli.LPush(ctx, repo.getSetSendMessageCardListUrgentKey(station), card).Err()
		if err != nil {
			hlog.Errorf("return send message card error:[%v]", err)
			return err
		}
		return nil
	}
	err = repo.cli.LPush(ctx, repo.getSetSendMessageCardListKey(station), card).Err()
	if err != nil {
		hlog.Errorf("return send message card error:[%v]", err)
		return err
	}
	return nil
}
