package userManagement

import (
	"database/sql"
	"errors"
	"log"
	"net/http"
	"time"

	"golang.org/x/crypto/bcrypt"

	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
)

type User struct {
	Id             int64
	Username       string `json:"username"`
	PasswordHashed string `json:"password"`
	Name           string `json:"name"`
	Surname        string `json:"surname"`
	Email          string `json:"email"`
	Role           string `json:"role"` // Role can be customer, roaster, or admin
}

type SessionFetch struct {
	ID        int64
	UserID    int64
	Token     string
	CreatedAt time.Time
	ExpiresAt time.Time
	IsActive  bool
}

// GetUserData fetches user data based on the username
func GetUserData(db *sql.DB, username string) (*User, error) {
	var user User

	// Adjust query to match new database schema
	err := db.QueryRow("SELECT id, username, password, name, surname, email, role FROM users WHERE username = ?", username).Scan(&user.Id, &user.Username, &user.PasswordHashed, &user.Name, &user.Surname, &user.Email, &user.Role)
	if err != nil {
		if err == sql.ErrNoRows {
			// Return an error instead of logging and terminating the program
			return nil, errors.New("user does not exist")
		}
		return nil, err
	}
	return &user, nil
}

// CheckPasswordMatchLogin compares the hashed password with the provided one
func CheckPasswordMatchLogin(password, hash string) bool {
	// Check if the provided password matches the stored hashed password
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func GetSingleUserData(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Param("session_token")
		if token == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid session token"})
			return
		}

		var session SessionFetch
		err := db.QueryRow("SELECT id, user_id, token, created_at, expires_at, is_active FROM sessions WHERE token = ?", token).
			Scan(&session.ID, &session.UserID, &session.Token, &session.CreatedAt, &session.ExpiresAt, &session.IsActive)
		if err != nil {
			if err == sql.ErrNoRows {
				c.JSON(http.StatusNotFound, gin.H{"error": "Session not found"})
			} else {
				log.Printf("Error creating session object: %v", err)
				c.JSON(http.StatusInternalServerError, gin.H{
					"error": "Error creating session object.",
				})
			}
			return
		}

		var user User
		err1 := db.QueryRow("SELECT id, username, password, name, surname, email, role FROM users WHERE id = ?", session.UserID).
			Scan(&user.Id, &user.Username, &user.PasswordHashed, &user.Name, &user.Surname, &user.Email, &user.Role)
		if err1 != nil {
			if err1 == sql.ErrNoRows {
				c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
			} else {
				log.Printf("Error creating a user object: %v", err)
				c.JSON(http.StatusInternalServerError, gin.H{
					"error": "Error creating a user object.",
				})
			}
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"title":      "RoastingRooster Coffee Inventory",
			"userObject": user,
		})
	}
}
