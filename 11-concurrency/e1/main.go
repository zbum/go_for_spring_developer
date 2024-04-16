package main

import "fmt"

func main() {
	c := gen(2, 3)
	out := sq(c)

	fmt.Println(<-out)
	fmt.Println(<-out)
}

func gen(nums ...int) <-chan int {
	out := make(chan int)
	go func() {
		for _, num := range nums {
			out <- num
		}
		close(out)
	}()
	return out
}

func sq(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for n := range in {
			out <- n * n
		}
	}()
	return out
}
