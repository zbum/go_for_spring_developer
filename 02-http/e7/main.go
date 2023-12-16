package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func main() {
	r := gin.Default()
	r.Handle(http.MethodGet, "/members/:member-id", func(ctx *gin.Context) {
		fmt.Fprintf(ctx.Writer, "your member id is %s", ctx.Param("member-id"))
	})
	log.Fatal(r.Run(":8080"))
}
