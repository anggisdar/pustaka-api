package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/", rootHandler)

	router.GET("/hello", helloHandler)

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
