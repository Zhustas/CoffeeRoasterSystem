package endpoints

import (
	"database/sql"
	"log"
	"net/http"

	"main/inventory"
	"main/userManagement"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type LoginCredentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func LoginUser(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var creds LoginCredentials

		// Parse the form data from the request body
		if err := c.ShouldBindJSON(&creds); err != nil {
			log.Printf("Error occurred while parsing request: %v", err)
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "Invalid request body",
			})
			return
		}

		// Fetch user from database
		user, err := userManagement.GetUserData(db, creds.Username)
		if err != nil {
			log.Printf("Error fetching user data: %v", err)
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "Invalid username or password",
			})
			return
		}

		// Check if the provided password matches the hashed password in the database
		if CheckPasswordMatchLogin(creds.Password, user.PasswordHashed) {
			// Successful login: Render a welcome page with the user's info
			c.JSON(http.StatusOK, gin.H{
				"username": user.Username,
				"role":     user.Role,
				"user_id":  user.Id,
			})
		} else {
			// Password doesn't match: Render login page with error
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "Invalid username or password",
			})
		}
	}
}

// Helper function to compare password with hashed password
func CheckPasswordMatchLogin(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func Register(db *sql.DB) gin.HandlerFunc {
	return userManagement.RegisterUser(db)
}

func DebugUsers(db *sql.DB) gin.HandlerFunc {
	return userManagement.FetchUsers(db)
}

func ManageNewCoffee(db *sql.DB) gin.HandlerFunc {
	return inventory.AddCoffee(db)
}

func DeleteOldCoffee(db *sql.DB) gin.HandlerFunc {
	return inventory.DeleteCoffee(db)
}

func DisplayCoffeeList(db *sql.DB) gin.HandlerFunc {
	return inventory.GetInventory(db)
}
