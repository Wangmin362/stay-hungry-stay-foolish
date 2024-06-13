package main

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v9"
)

func main() {
	ctx := context.Background()
	rdb := redis.NewClient(&redis.Options{
		Username: "",
		Password: "123456",
		Addr:     "172.30.3.224:6379",
		DB:       1,
		PoolSize: 5,
	})
	// ping一下检查是否连通
	pingResult, err := rdb.Ping(ctx).Result()
	if err != nil {
		panic(err)
	}
	// PONG
	fmt.Println(pingResult)

	rdb.Set(ctx, "aaa", "bbb", 0)
	get := rdb.Get(ctx, "aaa")
	fmt.Println(get.Val())

}
