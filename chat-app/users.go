package main

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
)

// createUser: user create endpoint
func createUser(c *gin.Context, db *sql.DB) {
	// Parse JSON request body into User struct
	var user User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Insert user into database
	result, err := db.Exec("INSERT INTO users (username, password) VALUES (?, ?)", user.Username, user.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Get ID of newly inserted user
	id, err := result.LastInsertId()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Return ID of newly inserted user
	c.JSON(http.StatusOK, gin.H{"id": id})
}
