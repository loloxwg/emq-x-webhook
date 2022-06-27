package main

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"time"
)

var ctx = context.Background()

func ExampleClient() {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "124.222.47.219:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	err := rdb.Set(ctx, "key", "value", 0).Err()
	if err != nil {
		panic(err)
	}

	val, err := rdb.Get(ctx, "key").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("key", val)

	val2, err := rdb.Get(ctx, "key2").Result()
	if err == redis.Nil {
		fmt.Println("key2 does not exist")
	} else if err != nil {
		panic(err)
	} else {
		fmt.Println("key2", val2)
	}
	// Output: key value
	// key2 does not exist
	pipe := rdb.TxPipeline()
	for {
		go func() {
			incr := pipe.Incr(ctx, "counter")
			_, exer := pipe.Exec(ctx)
			if exer != nil {
				fmt.Println(exer)
			}
			fmt.Println(incr.Val())
		}()

	}

	pipe.Expire(ctx, "key", time.Hour)

}
func main() {
	ExampleClient()
}
