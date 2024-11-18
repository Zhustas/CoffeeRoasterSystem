package userManagement

import (
	"database/sql"
	"errors"

	"golang.org/x/crypto/bcrypt"

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
