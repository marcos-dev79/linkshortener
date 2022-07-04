package lib

import (
	"context"
	"linkshortener/db"
	"net/http"

	"github.com/gin-gonic/gin"
)

// This method saves the redirection into the mongo database.
// Usage:
// To shorten an URL, just http POST to http://localhost:8040/api/shorten with the following json payload (example):
//
// {
//    "url":"https://www.vultr.com/docs/create-a-crud-application-with-golang-and-mongodb-on-ubuntu-20-04/?utm_source=performance-max-latam&utm_medium=paidmedia&obility_id=17096555207&utm_adgroup=&utm_campaign=&utm_term=&utm_content=&gclid=CjwKCAjwk_WVBhBZEiwAUHQCme_6kOgaeQWOKjalscslO99kCatxV5FJFdtFbqGv1127YkYBURCQ0BoCHnMQAvD_BwE"
// }
func InsertLink(c *gin.Context) {
	var url db.Request

	if err := c.BindJSON(&url); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		panic(err)
	}

	link := ShortenLink(url)
	link.Counter = 0

	mongoclient := db.Mongo_connect()

	linksCollection := mongoclient.Database("linkshortener").Collection("links")
	result, err := linksCollection.InsertOne(context.TODO(), link)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		panic(err)
	}

	c.JSON(http.StatusCreated, gin.H{"shortened_url": "http://localhost:8040/" + link.Shorthash, "id": result.InsertedID})
}
