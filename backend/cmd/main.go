package main

import (
	"log"
	"net/http"

	"database/sql"
	"main/endpoints"

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

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	r.POST("/login", endpoints.LoginUser(db))
	r.POST("/register", endpoints.Register(db))
	r.GET("/coffeeinventory", endpoints.DisplayCoffeeList(db))
	r.POST("/coffeeinventoryrefresh", endpoints.DisplayCoffeeList(db))
	r.POST("/addcoffee", endpoints.ManageNewCoffee(db))
	r.POST("/deletecoffee/:id", endpoints.ManageNewCoffee(db))
	// r.GET("/debuggingusers", endpoints.DebugUsers(db))
	// r.POST("/verifyauth", endpoints.VerifyAuthentication(db))
	// r.GET("/genereatepassword", endpoints.CreatePassword)

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
