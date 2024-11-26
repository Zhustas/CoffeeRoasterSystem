package orders

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Order struct {
	ID          int64   `json:"id"`
	UserID      int64   `json:"user_id"`
	TotalAmount float64 `json:"total_amount"`
	Status      string  `json:"status"`
	CreatedAt   string  `json:"created_at"`
	UpdatedAt   string  `json:"updated_at"`
}

// InspectOrders fetches all pending orders and returns them
func InspectOrders(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Fetch pending orders
		rows, err := db.Query(`
            SELECT id, user_id, total_amount, status, created_at, updated_at
            FROM orders
            WHERE status = 'pending'
        `)
		if err != nil {
			log.Printf("Error fetching orders: %v", err)
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "Failed to load orders. Please try again later.",
			})
			return
		}
		defer rows.Close()

		var orders []Order
		for rows.Next() {
			var order Order
			if err := rows.Scan(&order.ID, &order.UserID, &order.TotalAmount, &order.Status, &order.CreatedAt, &order.UpdatedAt); err != nil {
				log.Printf("Error scanning order: %v", err)
				c.JSON(http.StatusInternalServerError, gin.H{
					"error": "Error loading order data.",
				})
				return
			}
			orders = append(orders, order)
		}

		c.JSON(http.StatusOK, gin.H{
			"title":  "Pending Orders",
			"orders": orders,
		})
	}
}

// UpdateOrderStatus updates the status of a specific order
func UpdateOrderStatus(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		orderID := c.Param("id")
		var request struct {
			Status string `json:"status" binding:"required"`
		}

		// Parse the new status from the request body
		if err := c.ShouldBindJSON(&request); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "Invalid input: Status is required",
			})
			return
		}

		// Update order status in the database
		_, err := db.Exec(`
            UPDATE orders
            SET status = ?, updated_at = datetime('now')
            WHERE id = ?
        `, request.Status, orderID)

		if err != nil {
			log.Printf("Error updating order status: %v", err)
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "Failed to update order status. Please try again later.",
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"message": "Order status updated successfully",
		})
	}
}
