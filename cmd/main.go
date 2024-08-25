package main

import (
	"github.com/gin-gonic/gin"
	"github.com/onurhanak/SongsterrAPI/internal"
)

func main() {
	router := gin.Default()
	router.GET("/search", internal.SearchRequest)
	router.Run("localhost:8080")
}
