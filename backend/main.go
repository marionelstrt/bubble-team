package main

import (
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
	registerRoutes(r, db)
	r.Run()
}
