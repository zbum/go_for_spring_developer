package main

import (
	"context"
	"fmt"
)

func main() {
	ctx := context.Background()
	ctx = context.WithValue(ctx, "one", 1)
	ctx = context.WithValue(ctx, "two", 2)

	fmt.Println(ctx.Value("one"))
	fmt.Println(ctx.Value("two"))
	fmt.Println(ctx.Value("three"))
}
