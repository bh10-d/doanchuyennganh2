package main

import (
	"fmt"
	"golang/entity"
	"log"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// "github.com/HelloWorld/goProductAPI/handlers"

func main() {
	//gorm will work with database while gin work with route
	//dsn is Data Source Name Configuration
	// 	mysql.Open(dsn): Creates a MySQL database connection using the provided DSN
	// &gorm.Config{}: An optional configuration object for GORM, the ORM (Object-Relational Mapping) library
	// dsn := "root:root@tcp(docker.for.mac.localhost:3306)/dacn2?charset=utf8mb4&parseTime=True&loc=Local"
	dsn := "root:root@tcp(172.26.0.2:3306)/dacn2?charset=utf8mb4&parseTime=True&loc=Local"
	// dsn := "root@tcp(127.0.0.1:3306)/dacn2?charset=utf8mb4&parseTime=True&loc=Local" // using when outside docker
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	// gorm.Open wull return 2 value
	// the first vallue is pointer to a *gorm.DB established database connection
	// the second value will return the err if cannot connect to db
	if err != nil {
		log.Fatalln("Cannot connect to MySQL:", err)
	}

	log.Println("Connected to MySQL:", db)

	router := gin.Default()

	//Config

	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	config.AllowMethods = []string{"POST", "GET", "PUT", "OPTIONS", "DELETE"}
	config.AllowHeaders = []string{"Origin", "Content-Type", "Authorization", "Accept", "User-Agent", "Cache-Control", "Pragma"}
	config.ExposeHeaders = []string{"Content-Length"}
	config.AllowCredentials = true
	config.MaxAge = 12 * time.Hour

	router.Use(cors.New(config))

	// example i have a path: v1/items or v1/product
	// and i using golang to get the path and redirect request to server own this function

	v1 := router.Group("/v1")
	{
		v1.POST("/items", entity.CreateItem(db)) // create item
		// v1.GET("/items", entity.GetListOfItems(db))        // list items
		v1.GET("/items/:id", entity.ReadItemById(db)) // get an item by ID
		v1.PUT("/items/:id", entity.EditItemById(db)) // edit an item by ID
		// v1.DELETE("/items/:id", entity.DeleteItemById(db)) // delete an item by ID

		//api for nodejs server
		v1.GET("/items", entity.GetData())           // test call api from another server
		v1.DELETE("/items/:id", entity.DeleteData()) // delete an item by ID
	}

	router.Run() // default port is 8080 and if i custom the port i will change follow format router.Run(":<Port>")

	fmt.Print("Main run")
}
