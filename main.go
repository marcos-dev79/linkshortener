package main

import (
	"linkshortener/lib"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/:hash", lib.Redirect)
	router.POST("/api/shorten", lib.InsertLink)

	router.Run(":8040")
}
