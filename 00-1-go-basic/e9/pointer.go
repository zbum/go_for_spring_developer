package main

import "fmt"

func main() {
	var foo int = 23
	fmt.Println(foo, &foo)

	var fooPtr *int = &foo
	fmt.Println(fooPtr)

	fmt.Println(*fooPtr)

	*fooPtr = *fooPtr / 2
	fmt.Println(foo)

	var value = 1
	plusOneValue(value)
	fmt.Println(value)

	plusOnePointer(&value)
	fmt.Println(value)
}

func plusOneValue(v int) {
	v += 1
	fmt.Println(v)
}

func plusOnePointer(v *int) {
	*v += 1
	fmt.Println(*v)
}
