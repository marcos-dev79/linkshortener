package lib

import (
	"context"
	"linkshortener/db"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

func Redirect(c *gin.Context) {
	shorthash := c.Param("hash")

	mongoclient := db.Mongo_connect()
	linksCollection := mongoclient.Database("linkshortener").Collection("links")
	cursor, err := linksCollection.Find(context.TODO(), bson.M{"shorthash": shorthash})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		panic(err)
	}

	var results []db.Link
	if err = cursor.All(context.TODO(), &results); err != nil {
		panic(err)
	}

	c.Redirect(http.StatusFound, results[0].URL)
}
