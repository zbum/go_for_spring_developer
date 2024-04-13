package main

import (
	"fmt"
	"go_for_spring_developer/05-configuration/e1/configuration"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func init() {
	os.Setenv("ACTIVE_PROFILE", "local")
	configuration.Init()
}

func main() {
	go func() {
		for {
			fmt.Println(configuration.RuntimeConf.Datasource.Address)
			fmt.Println(configuration.RuntimeConf.Server.Port)
			time.Sleep(2 * time.Second)
		}
	}()

	done := make(chan os.Signal, 1)
	signal.Notify(done, syscall.SIGINT, syscall.SIGTERM)
	fmt.Println("Waiting for pressing ctrl+c to exit...")
	<-done
}
