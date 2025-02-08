package main

import (
	"log"

	"github.com/anggisdar/pustaka-api/book"
	"github.com/anggisdar/pustaka-api/handler"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	router := gin.Default()
	dsn := "root:05072001@(127.0.0.1:3306)/pustaka_api?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Db connection error")
	}

	db.AutoMigrate(&book.Book{}) // table auto
	// fmt.Println("Database connection succed")

	v1 := router.Group("v1") // versioning

	v1.GET("/", handler.RootHandler) // root url yang berisi JSON
	v1.GET("/hello", handler.HelloHandler)
	v1.GET("/books/:id/:title", handler.BooksHandler)
	v1.GET("/query", handler.QueryHandler)
	v1.POST("/books", handler.PostBooksHandler)

	// v2 := router.Group("v2")

	// v2.GET("/")

	// // root url "/"
	// router.GET("/", func(c *gin.Context) {
	// 	c.JSON(http.StatusOK, gin.H{
	// 		"Name": "Anggis Darmawan",
	// 		"bio":  "A Engineer & Backend Developer enthusiast",
	// 	})
	// })

	// router.GET("/hello", func(c *gin.Context) {
	// 	c.JSON(http.StatusOK, gin.H{
	// 		"Content":  "Perubahan kecil",
	// 		"subtitle": "Belajar Golang bareng Agung Setiawan",
	// 	})
	// })

	//port 8080
	router.Run()
}
