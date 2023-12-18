package main

import (
	"go_for_spring_developer/01-class/e2/bean"
	"log"
)

func main() {
	log.Printf("%v \n", bean.UserServiceBean.GetUsers())
}
