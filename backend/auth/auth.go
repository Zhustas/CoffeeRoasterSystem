package auth

import (
	"crypto/rand"
	"database/sql"
	"encoding/base64"
	"errors"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// SessionConfig holds configuration for session cookies
type SessionConfig struct {
	CookieName     string
	CookieLifetime time.Duration
	IsSecure       bool
	Domain         string
}

type Session struct {
	ID        int64
	UserID    int64
	Token     string
	Role      string // Add user role
	CreatedAt time.Time
	ExpiresAt time.Time
	IsActive  bool
}

// DefaultSessionConfig returns default session configuration
func DefaultSessionConfig() SessionConfig {
	return SessionConfig{
		CookieName:     "session_token",
		CookieLifetime: 24 * time.Hour,
		IsSecure:       true,
		Domain:         "",
	}
}

// GenerateSessionToken creates a secure random token
func GenerateSessionToken() (string, error) {
	b := make([]byte, 32)
	_, err := rand.Read(b)
	if err != nil {
		return "", err
	}
	return base64.URLEncoding.EncodeToString(b), nil
}

// CreateSession creates a new session in the database
func CreateSession(db *sql.DB, userID int64, token string, lifetime time.Duration) (*Session, error) {
	session := &Session{
		UserID:    userID,
		Token:     token,
		CreatedAt: time.Now(),
		ExpiresAt: time.Now().Add(lifetime),
		IsActive:  true,
	}

	query := `
        INSERT INTO sessions (user_id, token, expires_at, is_active)
        VALUES (?, ?, ?, ?)
    `

	result, err := db.Exec(query, session.UserID, session.Token, session.ExpiresAt, session.IsActive)
	if err != nil {
		return nil, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}
	session.ID = id

	return session, nil
}

// ValidateSession checks if a session is valid and active
func ValidateSession(db *sql.DB, token string) (*Session, error) {
	session := &Session{}
	query := `
        SELECT s.id, s.user_id, s.token, u.role, s.created_at, s.expires_at, s.is_active
        FROM sessions s
        INNER JOIN users u ON s.user_id = u.id
        WHERE s.token = ? AND s.is_active = TRUE AND s.expires_at > datetime('now')
    `

	err := db.QueryRow(query, token).Scan(
		&session.ID,
		&session.UserID,
		&session.Token,
		&session.Role, // Fetch user role
		&session.CreatedAt,
		&session.ExpiresAt,
		&session.IsActive,
	)

	if err == sql.ErrNoRows {
		return nil, errors.New("session not found or expired")
	}
	if err != nil {
		return nil, err
	}

	return session, nil
}

// DeactivateSession marks a session as inactive
func DeactivateSession(db *sql.DB, token string) error {
	query := `
        UPDATE sessions 
        SET is_active = FALSE 
        WHERE token = ?
    `

	result, err := db.Exec(query, token)
	if err != nil {
		return err
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rows == 0 {
		return errors.New("session not found")
	}

	return nil
}

// CleanupExpiredSessions removes expired sessions from the database
func CleanupExpiredSessions(db *sql.DB) error {
	query := `
        DELETE FROM sessions 
        WHERE expires_at <= datetime('now')
        OR is_active = FALSE
    `

	_, err := db.Exec(query)
	return err
}

// SetSessionCookie creates a session and sets the cookie
func SetSessionCookie(c *gin.Context, db *sql.DB, config SessionConfig, userID int64, username string) error {
	// Generate session token
	token, err := GenerateSessionToken()
	if err != nil {
		return err
	}

	// Create session in database
	session, err := CreateSession(db, userID, token, config.CookieLifetime)
	if err != nil {
		return err
	}

	// Set cookie
	c.SetCookie(
		config.CookieName,
		token,
		int(config.CookieLifetime.Seconds()),
		"/",
		config.Domain,
		config.IsSecure,
		true,
	)

	// Send response
	c.JSON(http.StatusOK, gin.H{
		"username":      username,
		"session_token": token,
		"expires_at":    session.ExpiresAt,
	})

	return nil
}

// AuthMiddleware verifies the session cookie against the database
func AuthMiddleware(db *sql.DB, config SessionConfig) gin.HandlerFunc {
	return func(c *gin.Context) {
		token, err := c.Cookie(config.CookieName)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "Unauthorized: No valid session",
			})
			c.Abort()
			return
		}

		// Validate session in database
		session, err := ValidateSession(db, token)
		if err != nil {
			c.SetCookie(config.CookieName, "", -1, "/", config.Domain, config.IsSecure, true)
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "Unauthorized: Invalid or expired session",
			})
			c.Abort()
			return
		}

		// Add session info to context
		c.Set("user_id", session.UserID)
		c.Set("user_role", session.Role) // Add user role to context
		c.Set("session_token", token)
		c.Next()
	}
}

// LogoutHandler handles user logout
func LogoutHandler(db *sql.DB, config SessionConfig) gin.HandlerFunc {
	return func(c *gin.Context) {
		token, err := c.Cookie(config.CookieName)
		if err == nil {
			// Deactivate session in database
			_ = DeactivateSession(db, token)
		}

		// Clear cookie
		c.SetCookie(config.CookieName, "", -1, "/", config.Domain, config.IsSecure, true)

		c.JSON(http.StatusOK, gin.H{
			"message": "Logged out successfully",
		})
	}
}

func RoleMiddleware(allowedRoles ...string) gin.HandlerFunc {
	return func(c *gin.Context) {
		role, exists := c.Get("user_role")
		if !exists {
			c.JSON(http.StatusForbidden, gin.H{"error": "Forbidden: No role found"})
			c.Abort()
			return
		}

		userRole := role.(string)
		for _, allowedRole := range allowedRoles {
			if userRole == allowedRole {
				c.Next()
				return
			}
		}

		c.JSON(http.StatusForbidden, gin.H{"error": "Forbidden: Insufficient permissions"})
		c.Abort()
	}
}
