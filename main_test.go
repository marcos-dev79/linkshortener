package main

import (
	"linkshortener/db"
	"linkshortener/lib"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/assert/v2"
	"go.mongodb.org/mongo-driver/mongo/integration/mtest"
)

func SetUpRouter() *gin.Engine {
	router := gin.Default()
	return router
}

func TestShortenLink(t *testing.T) {
	var request db.Request
	request.URL = "https://github.com/victorneuret/mongo-go-driver-mock/blob/master/insert_test.go"

	mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))
	defer mt.Close()

	mt.Run("Mocking Shortener", func(mt *mtest.T) {
		mt.AddMockResponses(mtest.CreateSuccessResponse())

		link := lib.ShortenLink(request)

		assert.Equal(t, db.Link{
			URL:       request.URL,
			Hash:      "fb84a1b0e9baba5686fe9a27a47e8594c0f77525",
			Shorthash: "fb84a1b0e9baba56",
			Created:   link.Created,
		}, link)
	})
}
