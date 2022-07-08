package lib

import (
	"context"
	"linkshortener/db"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

// This method performs the redirection of the shortened link.
// Usage:
// Just hit http://localhost:8040/2J2VhO ( use the generated hash )
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

	if len(results) > 0 {
		results[0].Counter = results[0].Counter + 1
		link := results[0]

		filter := bson.M{"url": results[0].URL}
		fields := bson.M{"$set": link}

		linksCollection.UpdateOne(context.TODO(), filter, fields)

		c.Redirect(http.StatusFound, results[0].URL)
	} else {
		c.JSON(http.StatusNotFound, gin.H{"error": "Link not found 404"})
	}
}
