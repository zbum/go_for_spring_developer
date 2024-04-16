package main

import (
	"log"
	"os"
)

var myFileLogger *log.Logger

func main() {
	// 로그파일 오픈
	fpLog, err := os.OpenFile("logfile.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		panic(err)
	}
	defer fpLog.Close()

	myFileLogger = log.New(fpLog, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)

	run()

	myFileLogger.Println("End of Program")
}

func run() {
	myFileLogger.Print("Test")
}
