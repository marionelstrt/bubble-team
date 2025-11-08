package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"github.com/marionelstrt/bubble-team/models"
)

func registerRoutes(r *gin.Engine, db *sqlx.DB) {
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
}
