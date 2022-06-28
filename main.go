package main

import (
	"linkshortener/lib"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"data": "hello world"})
	})

	router.POST("/api/shorten", lib.ShortenLink)

	router.Run(":8040")
}
