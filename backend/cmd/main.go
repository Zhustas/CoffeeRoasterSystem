package main

import (
	"log"
	"time"

	"main/auth"

	"database/sql"
	"main/endpoints"

	"github.com/gin-gonic/gin"

	// "database/sql"

	_ "github.com/mattn/go-sqlite3"
)

func setupSessionCleanup(db *sql.DB) {
	ticker := time.NewTicker(1 * time.Hour)
	go func() {
		for range ticker.C {
			if err := auth.CleanupExpiredSessions(db); err != nil {
				log.Printf("Error cleaning up sessions: %v", err)
			}
		}
	}()
}

func main() {

	db, err := sql.Open("sqlite3", "db/coffee.db") //open db connection
	if err != nil {
		log.Fatal("Error opening database: ", err)
	}
	defer db.Close() // close db

	// Start session cleanup routine
	setupSessionCleanup(db)

	r := gin.Default()

	// Session configuration
	sessionConfig := auth.DefaultSessionConfig()

	// Public routes
	r.POST("/login", endpoints.LoginUser(db))
	r.POST("/register", endpoints.Register(db))

	// Protected routes
	protected := r.Group("/")
	protected.Use(auth.AuthMiddleware(db, sessionConfig))
	{
		protected.GET("/coffeeinventory", endpoints.DisplayCoffeeList(db))
		protected.POST("/coffeeinventoryrefresh", endpoints.DisplayCoffeeList(db))
		protected.POST("/addcoffee", endpoints.ManageNewCoffee(db))
		protected.POST("/deletecoffee/:id", endpoints.ManageNewCoffee(db))
	}

	r.Run()
}
