package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func main() {
	router := gin.Default()

	v1 := router.Group("v1") // versioning

	v1.GET("/", rootHandler) // root url yang berisi JSON
	v1.GET("/hello", helloHandler)
	v1.GET("/books/:id/:title", booksHandler)
	v1.GET("/query", queryHandler)
	v1.POST("/books", PostBooksHandler)

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
	Title    string      `json:"title" binding:"required"`        // apabila required tidak terpenuhi maka sever akan berhenti
	Price    json.Number `json:"price" binding:"required,number"` // int "required,number" // json.Number support int "10" "bukanini"
	Subtitle string      `json:"sub_title"`                       // alias
}

func PostBooksHandler(c *gin.Context) {
	var bookInput BookInput

	err := c.ShouldBindJSON(&bookInput)
	if err != nil {

		errorMessages := []string{} //slice
		for _, e := range err.(validator.ValidationErrors) {
			errorMessage := fmt.Sprintf("Error on filed %s, condition: %s", e.Field(), e.ActualTag()) // menampilan error message
			errorMessages = append(errorMessages, errorMessage)                                       // append
			// fmt.Println(err)
		}
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": errorMessages,
		})
		return

	}

	c.JSON(http.StatusOK, gin.H{
		"title":     bookInput.Title,
		"price":     bookInput.Price,
		"sub_title": bookInput.Subtitle,
	})

}
