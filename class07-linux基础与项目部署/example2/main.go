package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()
	r.GET("/hello", func(context *gin.Context) {
		context.JSON(200, gin.H{
			"status": http.StatusOK,
			"info":   "hello",
		})
	})
	r.Run(":8077")
}
