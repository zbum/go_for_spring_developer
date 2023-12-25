package main

import (
	"context"
	"log"
	"sync"
	"time"
)

func main() {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	wg := sync.WaitGroup{}

	for i := 0; i < 2; i++ {
		wg.Add(1)
		go func(i int) {
			work(ctx, i)
			wg.Done()
		}(i)
	}

	time.AfterFunc(4*time.Second, cancel)
	wg.Wait()
	log.Println("completed")
}

func work(ctx context.Context, i int) {
	ctx, cancel := context.WithCancel(ctx)
	//	time.AfterFunc(time.Duration(i+1)*time.Second, cancel)
	defer cancel()
	slowFn(ctx, i)
}

func slowFn(ctx context.Context, i int) {
	ctx = context.WithValue(ctx, "one", 1)
	ctx = context.WithValue(ctx, "two", 2)

	log.Printf("slow function %d started. \n", i)
	select {
	case <-time.Tick(3 * time.Second):
		log.Printf("slow function %d finished\n", i)
	case <-ctx.Done():
		log.Printf("slow function %d too slow: %s \n", i, ctx.Err())
	}

}
