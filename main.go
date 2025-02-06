package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/", rootHandler) // root url yang berisi JSON
	router.GET("/hello", helloHandler)
	router.GET("/books/:id/:title", booksHandler)
	router.GET("/query", queryHandler)
	router.POST("/books", PostBooksHandler)

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
		"subtitle": "golang the best language",
	})
}

func booksHandler(c *gin.Context) {
	id := c.Param("id")       // tangkap id
	title := c.Param("title") // diambil dari query string

	c.JSON(http.StatusOK, gin.H{"id": id, "title": title})
}

// func queryHandler(c *gin.Context) {
// 	title := c.Param("title") // tangkap title

// 	c.JSON(http.StatusOK, gin.H{"title": title})
// }

func queryHandler(c *gin.Context) {
	title := c.Query("title") //query?title=value&
	price := c.Query("price")

	c.JSON(http.StatusOK, gin.H{"title": title, "price": price})
}

// post & binding json
type BookInput struct { // struct
	Title    string `json:"title" binding:"required"`        // apabila required tidak terpenuhi maka sever akan berhenti
	Price    int    `json:"price" binding:"required,number"` // int "required,number"
	Subtitle string `json:"sub_title"`                       // alias
}

func PostBooksHandler(c *gin.Context) {
	var bookInput BookInput

	err := c.ShouldBindJSON(&bookInput)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		fmt.Println(err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"title":     bookInput.Title,
		"price":     bookInput.Price,
		"sub_title": bookInput.Subtitle,
	})

}

// http post & json upload
