package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
    app := gin.Default()
    app.GET("/", func(ctx *gin.Context)  {
        ctx.String(http.StatusOK, "hello gin!")
        fmt.Println("hello is requested")
    })
    app.GET("/id/:id/hello", func(ctx *gin.Context)  {
        ctx.JSON(http.StatusOK, gin.H{"requestId": ctx.Param("id"), "query": ctx.Query("name")})
    })
    app.Run(":3000")
}