package endpoints

import (
	"database/sql"
	"log"
	"net/http"

	"main/auth"
	"main/customer"
	"main/inventory"
	"main/orders"
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
		sessionConfig := auth.DefaultSessionConfig()

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

		// Check if the provided password matches
		if CheckPasswordMatchLogin(creds.Password, user.PasswordHashed) {
			// Set session cookie
			err := auth.SetSessionCookie(c, db, sessionConfig, user.Id, user.Username)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{
					"error": "Failed to create session",
				})
				return
			}

			// The response is now handled by SetSessionCookie
		} else {
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

func AdminRegister(db *sql.DB) gin.HandlerFunc {
	return userManagement.AdminRegisterUser(db)
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

func CreateNewOrder(db *sql.DB) gin.HandlerFunc {
	return customer.CreateOrder(db)
}

func ViewOrders(db *sql.DB) gin.HandlerFunc {
	return orders.InspectOrders(db)
}

func UpdateOrders(db *sql.DB) gin.HandlerFunc {
	return orders.UpdateOrderStatus(db)
}

func DeleteOldOrder(db *sql.DB) gin.HandlerFunc {
	return orders.DeleteOrder(db)
}

func GetCoffeeWithId(db *sql.DB) gin.HandlerFunc {
	return inventory.GetCoffee(db)
}

func GetSingleUser(db *sql.DB) gin.HandlerFunc {
	return userManagement.GetSingleUserData(db)
}
