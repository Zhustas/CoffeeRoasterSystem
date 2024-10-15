package endpoints

import (
	"database/sql"
	"log"
	"net/http"

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
		if err := c.ShouldBind(&creds); err != nil {
			log.Printf("Error occurred while parsing request: %v", err)
			c.HTML(http.StatusBadRequest, "login.html", gin.H{
				"error": "Invalid request body",
			})
			return
		}

		// Fetch user from database
		user, err := userManagement.GetUserData(db, creds.Username)
		if err != nil {
			log.Printf("Error fetching user data: %v", err)
			c.HTML(http.StatusUnauthorized, "login.html", gin.H{
				"error": "Invalid username or password",
			})
			return
		}

		// Check if the provided password matches the hashed password in the database
		if CheckPasswordMatchLogin(creds.Password, user.PasswordHashed) {
			// Successful login: Render a welcome page with the user's info
			c.HTML(http.StatusOK, "welcome.html", gin.H{
				"username": user.Username,
				"role":     user.Role,
				"user_id":  user.Id,
			})
		} else {
			// Password doesn't match: Render login page with error
			c.HTML(http.StatusUnauthorized, "login.html", gin.H{
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

func VerifyAuthentication(db *sql.DB) gin.HandlerFunc {
	return userManagement.VerifyCode(db)
}
