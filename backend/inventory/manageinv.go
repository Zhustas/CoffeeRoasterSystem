package inventory

import (
	"database/sql"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
)

type Coffee struct {
	ID          int     `json:"ID"`
	Name        string  `json:"name"`
	RoastType   string  `json:"roast_type"`
	Description string  `json:"description"`
	Stock       int     `json:"stock"`
	Price       float64 `json:"price"`
}

func GetInventory(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Query to get all coffee entries
		rows, err := db.Query("SELECT id, name, roast_type, description, stock, price FROM coffee")
		if err != nil {
			log.Printf("Error fetching inventory: %v", err)
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "Failed to load inventory. Please try again later.",
			})
			return
		}
		defer rows.Close()

		var coffeeList []Coffee
		for rows.Next() {
			var coffee Coffee
			if err := rows.Scan(&coffee.ID, &coffee.Name, &coffee.Description, &coffee.RoastType, &coffee.Stock, &coffee.Price); err != nil {
				log.Printf("Error scanning row: %v", err)
				c.JSON(http.StatusInternalServerError, gin.H{
					"error": "Error loading inventory data.",
				})
				return
			}
			coffeeList = append(coffeeList, coffee)
		}

		c.JSON(http.StatusOK, gin.H{
			"title":      "RoastingRooster Coffee Inventory",
			"coffeeList": coffeeList,
		})
	}
}

func AddCoffee(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var newCoffee Coffee

		log.Printf("Form data: %+v", c.Request.Form)

		// Bind form data to the newCoffee struct
		if err := c.ShouldBindJSON(&newCoffee); err != nil {
			log.Printf("Error binding form data: %v", err)
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "Invalid input. Please fill out all required fields.",
			})
			return
		}

		// Check that required fields are not empty or zero
		if newCoffee.Name == "" || newCoffee.RoastType == "" || newCoffee.Stock <= 0 || newCoffee.Price <= 0 {
			log.Printf("Invalid form data: %+v", newCoffee)
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "All fields must be filled out correctly.",
			})
			return
		}

		// SQL query to insert a new coffee item
		query := `
			INSERT INTO coffee (name, roast_type, description, stock, price, created_at, updated_at)
			VALUES (?, ?, ?, ?, ?, ?, ?)
		`

		_, err := db.Exec(query, newCoffee.Name, newCoffee.RoastType, newCoffee.Description, newCoffee.Stock, newCoffee.Price, time.Now(), time.Now())
		if err != nil {
			log.Printf("Error adding new coffee: %v", err)
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "Failed to add new coffee. Please try again.",
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"title":          "RoastingRooster Coffee Inventory",
			"dbActionStatus": "SUCCESS",
		})
	}
}

func DeleteCoffee(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		if id == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid coffee ID"})
			return
		}

		query := "DELETE FROM coffee WHERE id = ?"
		res, err := db.Exec(query, id)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete coffee"})
			return
		}

		rowsAffected, _ := res.RowsAffected()
		if rowsAffected == 0 {
			c.JSON(http.StatusNotFound, gin.H{"error": "Coffee not found"})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"title":          "RoastingRooster Coffee Inventory",
			"dbActionStatus": "SUCCESS",
		})
	}
}

func GetCoffee(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		if id == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid coffee ID"})
			return
		}
		var coffee Coffee
		err := db.QueryRow("SELECT name, roast_type, description, stock, price FROM coffee WHERE id = ?", id).
			Scan(&coffee.Name, &coffee.RoastType, &coffee.Description, &coffee.Stock, &coffee.Price)
		if err != nil {
			if err == sql.ErrNoRows {
				c.JSON(http.StatusNotFound, gin.H{"error": "Coffee not found"})
			} else {
				log.Printf("Error creating a coffee object: %v", err)
				c.JSON(http.StatusInternalServerError, gin.H{
					"error": "Error creating a coffee object.",
				})
			}
			return
		}

		coffee.ID, err = strconv.Atoi(id)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Couldn't convert to string. Check the value of id"})
		}

		c.JSON(http.StatusOK, gin.H{
			"title":        "RoastingRooster Coffee Inventory",
			"coffeeObject": coffee,
		})
	}
}
