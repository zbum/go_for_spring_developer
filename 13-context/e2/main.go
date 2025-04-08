package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)

	go func() {
		fmt.Println("Starting...")
		select {
		case <-time.After(2 * time.Second):
			fmt.Println("Completed")
		case <-ctx.Done():
			fmt.Println("Cancelled: ", ctx.Err())
		}
	}()

	time.Sleep(1 * time.Second)
	cancel()

	time.Sleep(1 * time.Second)
}
