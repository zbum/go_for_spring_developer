package main

import (
	"fmt"
)

func main() {
	//r := bean.NewUserRepository()
	//s := bean.NewUserService(r)
	s := initializeUserService()
	fmt.Println(s.GetUsers())
}
