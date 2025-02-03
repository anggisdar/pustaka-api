package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/", rootHandler) // root url memanggil fungsi rootHandler yang berisi JSON
	router.GET("/hello", helloHandler)
	router.GET("/books/:id", booksHandler)
	router.GET("/query", queryHandler)

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
	router.Run(":8000")
}

func rootHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"Name": "Anggis Darmawan",
		"bio":  "A Engineer & Backend Developer enthusiast",
	})
}

func helloHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"Content":  "Perubahan kecil",
		"subtitle": "Belajar Golang bareng Agung Setiawan",
	})
}

func booksHandler(c *gin.Context) {
	id := c.Param("id") // tangkap id

	c.JSON(http.StatusOK, gin.H{"id": id})
}

// func queryHandler(c *gin.Context) {
// 	title := c.Param("title") // tangkap title

// 	c.JSON(http.StatusOK, gin.H{"title": title})
// }

func queryHandler(c *gin.Context) {
	title := c.Query("title")
	price := c.Query("price")
	quantity := c.Query("quantity")

	c.JSON(http.StatusOK, gin.H{"title": title, "price": price, "entity": quantity})
}
