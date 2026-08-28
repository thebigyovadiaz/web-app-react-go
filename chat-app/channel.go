package main

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
)

func createChannel(c *gin.Context, db *sql.DB) {
	// Parse JSON request body into Channel struct
	var channel Channel
	if err := c.ShouldBindJSON(&channel); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Insert channel into database
	result, err := db.Exec("INSERT INTO channels (name) VALUES (?)", channel.Name)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Get ID of newly inserted channel
	id, err := result.LastInsertId()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Return ID of newly inserted channel
	c.JSON(http.StatusOK, gin.H{"id": id})
}

func listChannels(c *gin.Context, db *sql.DB) {
	// Query database for channels
	rows, err := db.Query("SELECT id, name FROM channels")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Create slice of channels
	var channels []Channel

	// Iterate over rows
	for rows.Next() {
		// Create new channel
		var channel Channel

		// Scan row into channel
		err := rows.Scan(&channel.ID, &channel.Name)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		// Append channel to slice
		channels = append(channels, channel)
	}

	// Return slice of channels
	c.JSON(http.StatusOK, channels)
}
