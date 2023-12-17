package main

import "fmt"

type Greeter struct {
	message string
}

func NewGreeter() *Greeter {
	return &Greeter{"Hello"}
}

func (g Greeter) Greet(name string) string {
	return fmt.Sprintf("%s, %s!!", g.message, name)
}

func (g *Greeter) ChangeMessage(message string) {
	g.message = message
}
