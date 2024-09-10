package e1

import "fmt"

func Foo(i int, s string) (string, error) {
	fmt.Printf("%d, %s \n", i, s)
	return s, nil

}
