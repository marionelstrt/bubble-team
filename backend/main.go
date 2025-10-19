package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"fmt"
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
		var json struct {
			Boba string `json:"boba" binding:"required"`
			Name string `json:"name" binding:"required"`	
			LastName string `json:"lastName" binding:"required"`
			Password string `json:"password" binding:"required"`
			Email string `json:"email" binding:"required"`
		}

		if c.Bind(&json) == nil {
			c.JSON(http.StatusOK, gin.H{"status": "ok"})
			fmt.Println("boba", json.Boba) 
			fmt.Println("name", json.Name)
			fmt.Println("lastName", json.LastName) 
			fmt.Println("password", json.Password) 
			fmt.Println("email", json.Email)
		}
	})


	r.Run(":8081")
}
