package lib

import (
	"context"
	"linkshortener/db"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

// This method deletes the redirection
// Usage:
// Send a http delete call to
// http://localhost:8040/api/delete/102f01e97fa6239c ( Use your URL hash )
func DeleteLink(c *gin.Context) {
	shorthash := c.Param("hash")

	mongoclient := db.Mongo_connect()
	linksCollection := mongoclient.Database("linkshortener").Collection("links")

	filter := bson.M{"shorthash": shorthash}
	_, err := linksCollection.DeleteOne(context.TODO(), filter)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		panic(err)
	}

	c.JSON(http.StatusInternalServerError, gin.H{"response": "Deleted successfully"})

}
