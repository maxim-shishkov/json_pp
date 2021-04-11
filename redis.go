package main

import (
	"context"
	"fmt"
	"github.com/ReneKroon/ttlcache/v2"
	"github.com/go-redis/redis/v8"
	"sync"
	"time"
)

type hacker struct {
	Name string `json:"name"`
	Score float64 `json:"score"`
}
type hackers []hacker

var (
	mu sync.Mutex
	ctx = context.Background()
	client *redis.Client
	cache = ttlcache.NewCache()
)
const KEY = "key"

func init() {
	client = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	cache.SetTTL(time.Duration(1 * time.Second))
}


func (h *hackers)getHacker() {
	mu.Lock()
	defer mu.Unlock()

	var data *redis.ZSliceCmd

	// проверяем наличее в кэше
	if val, err := cache.Get(KEY); err != nil {
		data = client.ZRangeWithScores(ctx, "hackers", 0, -1)
		cache.Set(KEY, data)
	} else {
		data = val.(*redis.ZSliceCmd)
	}

	// разбираем и упаковываем данные
	for _, z := range data.Val() {
		*h = append(*h, hacker{
			Name:  fmt.Sprintf("%v", z.Member ),
			Score: z.Score,
		})
	}
}
