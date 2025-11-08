package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/marionelstrt/bubble-team/models"

	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)

func initDB() *sqlx.DB {
	db := sqlx.MustConnect("sqlite3", "db.sqlite3")
	models.CreateUserTable(db)
	return db
}

func main() {
	db := initDB()

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
		var user models.User

		if err := c.BindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		_, err := user.Save(db)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		} else {
			c.JSON(http.StatusCreated, gin.H{"message": "created"})
		}
	})

	_ = r.Run(":8081")
}
