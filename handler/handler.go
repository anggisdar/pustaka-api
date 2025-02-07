package handler

import (
	"fmt"
	"net/http"

	"github.com/anggisdar/pustaka-api/book"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func RootHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"Name": "Anggis Darmawan",
		"bio":  "A Engineer & Backend Developer enthusiast",
	})
}

func HelloHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"Content":  "Perubahan kecil",
		"subtitle": "golang the best language",
	})
}

func BooksHandler(c *gin.Context) {
	id := c.Param("id")       // tangkap id
	title := c.Param("title") // diambil dari query string

	c.JSON(http.StatusOK, gin.H{"id": id, "title": title})
}

// func queryHandler(c *gin.Context) {
// 	title := c.Param("title") // tangkap title

// 	c.JSON(http.StatusOK, gin.H{"title": title})
// }

func QueryHandler(c *gin.Context) {
	title := c.Query("title") //query?title=value&
	price := c.Query("price")

	c.JSON(http.StatusOK, gin.H{"title": title, "price": price})
}

func PostBooksHandler(c *gin.Context) {
	var bookInput book.BookInput

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
