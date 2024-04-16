package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func main() {
	// gin.SetMode(gin.ReleaseMode)
	r := gin.Default()

	//  curl http://localhost:8080/members/1
	r.Handle(http.MethodGet, "/members/:member-id", func(ctx *gin.Context) {
		fmt.Fprintf(ctx.Writer, "your member id is %s", ctx.Param("member-id"))
	})
	log.Fatal(r.Run(":8080"))
}
