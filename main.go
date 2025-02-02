package main

import (
	"answer/pkg/answer"
	"github.com/gin-gonic/gin"
)

func handler(c *gin.Context) {
	answer.SendResponseSuccess(c, "ddd", 404)
}
func main() {
	r := gin.Default()
	r.GET("/", handler)
	r.Run(":8080")
}
