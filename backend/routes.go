package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/jmoiron/sqlx"
	"github.com/marionelstrt/bubble-team/models"
)

var secretKey = []byte("secret-key")

func createToken(mail string) (string, error) {
	claims := jwt.MapClaims{
		"mail": mail,
		"exp":  time.Now().Add(time.Hour * 24).Unix(),
		"iat":  time.Now().Unix(),
	}
	t := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return t.SignedString(secretKey)
}

func AuthentificationRequired() gin.HandlerFunc {
	return func(c *gin.Context) {
		cookie, err := c.Cookie("user")
		if err != nil || cookie == "" {
			log.Println("cookie not found")
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "cookie not found"})
			return
		}

		token, err := jwt.Parse(cookie, func(t *jwt.Token) (interface{}, error) {

			if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method")
			}
			return secretKey, nil
		})
		if err != nil || !token.Valid {
			log.Println("invalid token:", err)
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "invalid token"})
			return
		}

		log.Println("cookie valid")
		c.Next()
	}
}

func registerRoutes(r *gin.Engine, db *sqlx.DB) {
	r.POST("/account/create", func(c *gin.Context) {
		var user models.User

		if err := c.BindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if err := user.Create(db); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		if err := SendVerificationEmail(user); err != nil {
			fmt.Println("Failed to send verification email:", err)
		}

		token, err := createToken(user.Email)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "token generation failed"})
			return
		}
		c.SetCookie("user", token, 3600, "/", "", false, true)

		c.JSON(http.StatusCreated, gin.H{
			"message": "created",
			"token":   token,
		})
	})

	r.POST("/account/verify", AuthentificationRequired(), func(c *gin.Context) {
	})
}

/*
   c.SetCookie("user", "", -1, "/", "localhost", false, true)
   c.String(http.StatusOK, "Cookie has been deleted")
*/
