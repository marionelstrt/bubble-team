package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()

	// temporary, before we have a reverse proxy, just to be able to contact the backend from the frontend
	// TODO: use caddy so that backend and frontend are on the same origin, and remove that
	r.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Origin, Content-Type, Authorization")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
			return
		}
		c.Next()
	})

	r.POST("/account/create", func(c *gin.Context) {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Sorry, didn't implement that yet."})
	})

	r.Run(":8081")
}
