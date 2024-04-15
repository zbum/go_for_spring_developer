package main

import (
	"fmt"
	"github.com/redis/go-redis/v9"
)

func main() {
	opt, err := redis.ParseURL("redis://localhost:6379/0")
	if err != nil {
		panic(err)
	}

	client := redis.NewClient(opt)

	fmt.Printf("%v\n", client)
}
