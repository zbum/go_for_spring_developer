package main

import (
	"fmt"
	"go.uber.org/fx"
)

func main() {
	fx.New(
		fx.Invoke(PrintHello),
	).
		Run()
}

func PrintHello() {
	fmt.Println("Hello")
}
