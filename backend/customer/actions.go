package customer

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateOrder(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Extract user_id from context
		userID, exists := c.Get("user_id")
		if !exists {
			fmt.Println(userID)
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			return
		}

		type OrderRequest struct {
			Items []struct {
				CoffeeID int64 `json:"coffee_id"`
				Quantity int   `json:"quantity"`
			} `json:"items"`
		}

		var orderRequest OrderRequest
		if err := c.ShouldBindJSON(&orderRequest); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request format"})
			return
		}

		tx, err := db.Begin()
		if err != nil {
			log.Printf("Failed to start transaction: %v", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create order"})
			return
		}

		defer func() {
			if p := recover(); p != nil {
				tx.Rollback()
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create order"})
			}
		}()

		var totalAmount float64
		for _, item := range orderRequest.Items {
			var coffeeStock int
			var unitPrice float64

			err = tx.QueryRow("SELECT stock, price FROM coffee WHERE id = ?", item.CoffeeID).Scan(&coffeeStock, &unitPrice)
			if err != nil {
				tx.Rollback()
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid coffee item"})
				return
			}

			if coffeeStock < item.Quantity {
				tx.Rollback()
				c.JSON(http.StatusConflict, gin.H{"error": "Insufficient stock for item"})
				return
			}

			totalAmount += float64(item.Quantity) * unitPrice

			_, err = tx.Exec("UPDATE coffee SET stock = stock - ?, updated_at = CURRENT_TIMESTAMP WHERE id = ?", item.Quantity, item.CoffeeID)
			if err != nil {
				tx.Rollback()
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update stock"})
				return
			}
		}

		result, err := tx.Exec("INSERT INTO orders (user_id, total_amount, status) VALUES (?, ?, 'pending')", userID, totalAmount)
		if err != nil {
			tx.Rollback()
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create order"})
			return
		}

		orderID, _ := result.LastInsertId()

		for _, item := range orderRequest.Items {
			_, err := tx.Exec("INSERT INTO order_items (order_id, coffee_id, quantity, unit_price) VALUES (?, ?, ?, ?)",
				orderID, item.CoffeeID, item.Quantity, item.Quantity)
			if err != nil {
				tx.Rollback()
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create order items"})
				return
			}
		}

		if err := tx.Commit(); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to finalize order"})
			return
		}

		c.JSON(http.StatusCreated, gin.H{"message": "Order created successfully", "order_id": orderID})
	}
}
