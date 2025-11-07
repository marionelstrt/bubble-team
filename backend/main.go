package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	r := gin.Default()

	db, err := sqlx.Connect("sqlite3", "db.sqlite3")
	if err != nil {
		log.Fatalln(err)
	}
	CreateTable(db)
	db.MustExec(schema)

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

	_ = db.MustBegin()

	r.POST("/account/create", func(c *gin.Context) {
		var user User

		if err := c.BindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		_, err = user.Save(db)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"status": "error"})
		} else {
			c.JSON(http.StatusCreated, gin.H{"status": "created"})
		}
	})

	_ = r.Run(":8081")
}
