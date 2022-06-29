package test

import (
	"context"
	"github.com/go-redis/redis/v8"
	"testing"
	"time"
)

var ctx = context.Background()
var rdb = redis.NewClient(&redis.Options{
	Addr:     "106.55.3.215:6377",
	Password: "123456@tengxunyun",
	DB:       0,
})

func TestSetValue(t *testing.T) {
	err := rdb.Set(ctx, "key", "value", time.Second*10).Err()
	if err != nil {
		t.Fatal(err)
	}
}
