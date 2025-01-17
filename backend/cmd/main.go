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
	db, err := sql.Open("sqlite3", "../db/coffee.db") // Open DB connection
	if err != nil {
		log.Fatal("Error opening database: ", err)
	}
	defer db.Close() // Close DB connection

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
		// Coffee inventory accessible to all authenticated users
		protected.GET("/coffeeinventory", endpoints.DisplayCoffeeList(db))
		protected.POST("/fetchcoffee/:id", endpoints.GetCoffeeWithId(db))
		protected.POST("/fetchuser/:session_token", endpoints.GetSingleUser(db))
		protected.POST("/getorderitems/:id", endpoints.GetAllOrderItems(db))

		// Admin and roaster routes
		adminOrRoaster := protected.Group("/")
		adminOrRoaster.Use(auth.RoleMiddleware("admin", "roaster"))
		{
			adminOrRoaster.POST("/registernewuser", endpoints.AdminRegister(db))
			adminOrRoaster.POST("/coffeeinventoryrefresh", endpoints.DisplayCoffeeList(db))
			adminOrRoaster.POST("/addcoffee", endpoints.ManageNewCoffee(db))
			adminOrRoaster.POST("/deletecoffee/:id", endpoints.DeleteOldCoffee(db))
			adminOrRoaster.GET("/orderlist", endpoints.ViewOrders(db))
			adminOrRoaster.GET("/allorderlist", endpoints.ViewAllOrders(db))
			adminOrRoaster.POST("/updateorders/:id", endpoints.UpdateOrders(db))
			adminOrRoaster.POST("/deleteorder/:id", endpoints.DeleteOldOrder(db))
			adminOrRoaster.GET("/fetchallusers", endpoints.GetAllUsers(db))
			adminOrRoaster.POST("/deleteuser/:id", endpoints.DeleteAUser(db))
		}

		// Customer routes
		customer := protected.Group("/")
		customer.Use(auth.RoleMiddleware("customer"))
		{
			customer.POST("/order", endpoints.CreateNewOrder(db))
			customer.POST("/fetchuserorder/:id", endpoints.GetAllUserOrders(db))
		}
	}

	r.Run()
}
