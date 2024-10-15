package main

import (
	"log"
	"net/http"

	"database/sql"
	"main/endpoints"
	"main/pages"

	"github.com/gin-gonic/gin"

	// "database/sql"

	_ "github.com/mattn/go-sqlite3"
)

func main() {

	db, err := sql.Open("sqlite3", "db/coffee.db") //open db connection
	if err != nil {
		log.Fatal("Error opening database: ", err)
	}
	defer db.Close() // close db

	r := gin.Default()

	r.Static("/static", "../frontend/static")
	r.LoadHTMLGlob("../frontend/templates/*")

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	r.GET("/", pages.ViewMainPage())
	r.POST("/login", endpoints.LoginUser(db))
	r.GET("/loginview", pages.LoginPage())
	r.POST("/register", endpoints.Register(db))
	r.GET("/registerview", pages.RegisterPage())
	// r.GET("/debuggingusers", endpoints.DebugUsers(db))
	// r.POST("/verifyauth", endpoints.VerifyAuthentication(db))
	// r.GET("/genereatepassword", endpoints.CreatePassword)

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
