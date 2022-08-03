package lib

import (
	"context"
	"linkshortener/db"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

// This method provides count info for redirections
// Usage:
// http://localhost:8040/api/counter/2J2VhO ( Use your URL hash )
func CounterInfo(c *gin.Context) {
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

	if len(results) > 0 {
		c.JSON(http.StatusOK, gin.H{"url": results[0].URL, "counter": results[0].Counter})
	} else {
		c.JSON(http.StatusNotFound, gin.H{"error": "Link not found 404"})
	}
}
