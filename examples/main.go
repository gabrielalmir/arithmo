package main

import (
	"context"
	"fmt"
	"log"

	"github.com/redis/go-redis/v9"
)

var ctx = context.Background()

func main() {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	// Enqueue an item
	err := rdb.LPush(ctx, "queue", "item1").Err()
	if err != nil {
		log.Fatalf("could not enqueue item: %v", err)
	}

	// Dequeue an item
	item, err := rdb.RPop(ctx, "queue").Result()
	if err != nil {
		log.Fatalf("could not dequeue item: %v", err)
	}

	fmt.Printf("Dequeued item: %s\n", item)
}
