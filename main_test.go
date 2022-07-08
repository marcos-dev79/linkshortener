package main

import (
	"bou.ke/monkey"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/assert/v2"
	"go.mongodb.org/mongo-driver/mongo/integration/mtest"
	"linkshortener/db"
	"linkshortener/lib"
	"testing"
	"time"
)

func SetUpRouter() *gin.Engine {
	router := gin.Default()
	return router
}

func TestShortenLink(t *testing.T) {
	var request db.Request
	request.URL = "https://github.com/victorneuret/mongo-go-driver-mock/blob/master/insert_test.go"

	mockTime := time.Date(2022, 7, 8, 20, 0, 0, 0, time.UTC)
	patch := monkey.Patch(time.Now, func() time.Time { return mockTime })
	defer patch.Unpatch()

	mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))
	defer mt.Close()

	mt.Run("Mocking Shortener", func(mt *mtest.T) {
		mt.AddMockResponses(mtest.CreateSuccessResponse())

		link := lib.ShortenLink(request)

		assert.Equal(t, db.Link{
			URL:       request.URL,
			Hash:      "a93bd42d81cfbfc8961d9eba9ecc8da2fdc582b5",
			Shorthash: "sQ5ay",
			Created:   link.Created,
		}, link)
	})
}
