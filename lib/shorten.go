package lib

import (
	"context"
	"crypto/sha1"
	"encoding/hex"
	"linkshortener/db"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func ShortenLink(c *gin.Context) {

	var url db.Request

	if err := c.BindJSON(&url); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		panic(err)
	}

	now := time.Now()
	sec := now.Unix()
	nsec := now.UnixNano()
	hash := sha1.Sum([]byte(url.URL + string(rune(nsec))))
	shorthash := hash[0:8]

	var link db.Link
	link.Created = sec
	link.Hash = hex.EncodeToString(hash[:])
	link.Shorthash = hex.EncodeToString(shorthash[:])
	link.URL = url.URL

	mongoclient := db.Mongo_connect()

	linksCollection := mongoclient.Database("linkshortener").Collection("links")
	result, err := linksCollection.InsertOne(context.TODO(), link)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		panic(err)
	}

	c.JSON(http.StatusOK, gin.H{"shortened_url": "http://localhost:8040/" + link.Shorthash, "id": result.InsertedID})

}
