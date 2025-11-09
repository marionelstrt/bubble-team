package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"github.com/marionelstrt/bubble-team/models"
)

func registerRoutes(r *gin.Engine, db *sqlx.DB) {
	r.POST("/account/create", func(c *gin.Context) {
		var user models.User

		err := c.BindJSON(&user)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		err = user.Create(db)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		err = SendVerificationEmail(user)
		if err != nil {
			fmt.Println("Failed to send verification email:", err)
		}

		c.JSON(http.StatusCreated, gin.H{"message": "created"})
	})

	r.POST("/account/verify", func(c *gin.Context) {
	})
}
