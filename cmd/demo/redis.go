package main

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"time"
)

func main() {
	rdb := connRdb()
	ctx := context.Background()
	err := rdb.Set(ctx, "session_id:admin", "session_id", 5*time.Second).Err()
	if err != nil {
		panic(err)
	}

	result, err := rdb.Get(ctx, "session_id:admin").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println(result)
}
