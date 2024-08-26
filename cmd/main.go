package main

import (
	"github.com/gin-gonic/gin"
	"github.com/onurhanak/songsterrapi/"
)

func main() {
	router := gin.Default()
	router.GET("/search", songsterrapi.SearchRequest)
	router.Run("localhost:8080")
}
