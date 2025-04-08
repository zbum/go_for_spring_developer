package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, 1*time.Second)
	defer cancel()

	go func() {
		fmt.Println("Starting...")
		select {
		case <-time.After(2 * time.Second):
			fmt.Println("Completed")
		case <-ctx.Done():
			fmt.Println("Cancelled: ", ctx.Err())
		}
	}()

	time.Sleep(3 * time.Second)

}
