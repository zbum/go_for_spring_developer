package main

import (
	"context"
	"fmt"
	"log"
	"sync"
	"time"
)

func main() {
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, 7*time.Second)
	defer cancel()

	wg := sync.WaitGroup{}

	for i := 0; i < 3; i++ {
		wg.Add(1)
		go work(ctx, &wg, i)
	}

	wg.Wait()
	log.Println("done")
}

func work(ctx context.Context, wg *sync.WaitGroup, id int) {
	defer wg.Done()
	switch id {
	case 0:
		ctx, cancel := context.WithTimeout(ctx, time.Second)
		time.AfterFunc(500*time.Millisecond, func() { cancel() })
		slowTask(ctx, id, fmt.Sprintf("worker %d had a timeout of 1 second", id))
	case 1:
		ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
		defer cancel()
		slowTask(ctx, id, fmt.Sprintf("worker %d had a timeout of 10 second", id))
	case 2:
		ctx, cancel := context.WithTimeout(ctx, -1*time.Second)
		defer cancel()
		slowTask(ctx, id, fmt.Sprintf("worker %d had a timeout of -1 second", id))
	}
}

func slowTask(ctx context.Context, id int, prefix string) {
	ctx = context.WithValue(ctx, "id", id)

	fmt.Printf("%d started\n", id)
	select {
	case <-time.Tick(15 * time.Second):
		log.Printf("%s: finished\n", prefix)
	case <-ctx.Done():
		log.Printf("%s: too slow function... returning %s\n", prefix, ctx.Err())
	}
}
