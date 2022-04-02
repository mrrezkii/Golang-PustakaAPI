package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"name": "Muhamamd Rezki Ananda",
			"bio":  "Professional Software Engineer",
		})
	})

	router.GET("/hello", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"title":    "Helo World",
			"subtitle": "Belajar Golang bareng Agung Setiawan",
		})
	})

	router.Run(":8888")
}
