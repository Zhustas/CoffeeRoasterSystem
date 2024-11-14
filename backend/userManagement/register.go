package userManagement

import (
	"database/sql"
	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

// hashPassword hashes the user's password using bcrypt
func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

// RegisterUser registers a new customer user
func RegisterUser(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var newUser User

		// Parse form data (as registration is typically handled via forms)
		if err := c.ShouldBindJSON(&newUser); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "Invalid input data. Please try again.",
			})
			return
		}

		// Only allow customer registration via the website
		newUser.Role = "customer"

		// Hash the password
		hashedPassword, err := hashPassword(newUser.PasswordHashed)
		if err != nil {
			log.Printf("Error hashing password: %v", err)
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "Internal server error. Please try again later.",
			})
			return
		}

		// Insert the new user into the database
		query := `INSERT INTO users (username, password, name, surname, email, role) VALUES (?, ?, ?, ?, ?, ?)`
		_, err = db.Exec(query, newUser.Username, hashedPassword, newUser.Name, newUser.Surname, newUser.Email, newUser.Role)
		if err != nil {
			// Handle potential errors, such as duplicate usernames or emails
			if strings.Contains(err.Error(), "UNIQUE constraint failed") {
				c.JSON(http.StatusInternalServerError, gin.H{
					"error": "This username or email is already taken. Please choose another.",
				})
			} else {
				log.Printf("Error occurred while registering user: %v", err)
				c.JSON(http.StatusInternalServerError, gin.H{
					"error": "Failed to register user. Please try again later.",
				})
			}
			return
		}

		// Successfully registered user
		c.JSON(http.StatusOK, gin.H{
			"message":        "Registration successful! You can now log in.",
			"dbActionStatus": "SUCCESS",
		})
	}
}
