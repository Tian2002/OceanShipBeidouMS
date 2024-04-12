package redis

import (
	"context"
	"testing"
)

func TestGetClient(t *testing.T) {
	initRedisClient()
	t.Log(GetClient().Ping(context.Background()).Result())
}
