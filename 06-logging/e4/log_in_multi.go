package main

import (
	"io"
	"log"
	"os"
)

func main() {
	// 로그파일 오픈
	fpLog, err := os.OpenFile("logfile.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		panic(err)
	}
	defer fpLog.Close()

	log.SetOutput(io.MultiWriter(os.Stdout, fpLog))

	run()

	log.Println("End of Program")
}

func run() {
	log.Print("Test")
}
