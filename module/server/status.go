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
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
	"database/sql"
)

const (
	DB_USER     = "akhil"
	DB_PASSWORD = "sayone"
	DB_NAME     = "gin"
)


// Version endpoint returns the server version and build information.
func Version(c *gin.Context) {
	c.JSON(200, gin.H{
		"source":  "https://github.com/sayonetech/gin-webapp",
		"version": config.GetEnv().VERSION,
	})
}

type Product struct {
	gorm.Model
	ID     int 		 	 `bson:"_id"`
	Title  string        `json:"title"`
	Price   int       `json:"price"`
	Rating  int        `json:"rating"`
}


type mProduct struct {
	ID     int
	Name  string
	Age   int
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


//func GetProducts(c *gin.Context) {
//	var qn []Product
//	dbinfo := "user=root password=root dbname=gin sslmode=disable"
//
//	//db, err := gorm.Open("postgres", dbinfo)
//	db, err := sql.Open("postgres", dbinfo)
//	if err != nil {
//		fmt.Println(dbinfo)
//		panic("failed to connect database")
//	}
//	defer db.Close()
//
//	rows, err := db.Query("SELECT * FROM product")
//	for rows.Next() {
//		rows.Scan(&qn)
//		fmt.Println(qn)
//	}
//	fmt.Println(err)
//	c.JSON(200, gin.H{"qn": qn})
//
//}


func GetProducts(c *gin.Context) {
	dbinfo := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable",
		DB_USER, DB_PASSWORD, DB_NAME)
	db, err := sql.Open("postgres", dbinfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()
	//var data []Product
	//rows, err := db.Find(&data)("SELECT * FROM company")
	rows, err := db.Query("SELECT * FROM company")

	if err != nil {
		panic(err)
	}

	for rows.Next() {
		var id int
		var name string
		var age int
		err = rows.Scan(&id, &name, &age)
		if err != nil {
			panic(err)
		}

		fmt.Println("id | name | age")
		fmt.Printf("%3v | %8v | %6v\n", id, name, age)
		for rows.Next() {
			s := mProduct{}
			if err := rows.Scan(&s); err != nil {
				return err
			}

			// Do something with 's'
		}

		b, err := json.Marshal(rows.Scan())
		fmt.Println(b)

		if err != nil {
			panic(err)
		}
	}


	c.JSON(200, rows)

}