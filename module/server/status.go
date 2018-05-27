package server

import (
	"github.com/gin-gonic/gin"
	"github.com/olivere/elastic"

	"go-webapp/config"
	"context"
	"reflect"
	"fmt"
	"encoding/json"
	"net/http"
)


// Version endpoint returns the server version and build information.
func Version(c *gin.Context) {
	c.JSON(200, gin.H{
		"source":  "https://github.com/sayonetech/gin-webapp",
		"version": config.GetEnv().VERSION,
	})
}

type Product struct {
	Id uint      `gorm:"primary_key"`
	Title string `gorm:"column:title"`
}

func TestApi(c *gin.Context) {
	// Create a client
	client, err := elastic.NewClient()

	// Search with a term query
	termQuery := elastic.NewMatchAllQuery()
	searchResult, err := client.Search().
		Index("product").
		Query(termQuery).
		Pretty(true).
		Do(context.Background())
	if err != nil {
		panic(err)
	}

	fmt.Printf("\n\nQuery took %d milliseconds\n", searchResult.TookInMillis)

	var ttyp Product
	for _, item := range searchResult.Each(reflect.TypeOf(ttyp)) {
		//Printing results on terminal
		if t, ok := item.(Product); ok {
			fmt.Printf("Product: %s\n", t.Title)
		}

		//Returning results as json
		b, err := json.Marshal(item)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(string(b))
		c.JSON(200, string(b))

	}
	fmt.Printf("Found a total of %d products\n\n", searchResult.TotalHits())

	c.JSON(200, searchResult)

	c.JSON(http.StatusOK, gin.H{
		"message": "ping-pong",
	})
}