package main

import "fmt"

type Greeter struct {
}

func NewGreeter() *Greeter {
	return &Greeter{}
}

func (g *Greeter) Greet(name string) string {
	return fmt.Sprintf("Hello, %s!!", name)
}
