package main

import (
	"linkshortener/lib"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/:hash", lib.Redirect)
	router.GET("/api/counter/:hash", lib.CounterInfo)
	router.POST("/api/shorten", lib.InsertLink)
	router.DELETE("/api/delete/:hash", lib.DeleteLink)

	router.Run(":8040")
}
