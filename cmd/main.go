package main

import (
	"github.com/gin-gonic/gin"
	"github.com/yourusername/yourmodulename/internal"
)

func main() {
	router := gin.Default()
	router.GET("/search", internal.SearchRequest)
	router.Run("localhost:8080")
}
