package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	_ "github.com/glebarez/go-sqlite"
)

type User struct {
	ID       int    `json:"id,omitempty"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type Channel struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type Message struct {
	ID        int    `json:"id"`
	ChannelID int    `json:"channel_id"`
	UserID    int    `json:"user_id"`
	Username  string `json:"username"`
	Text      string `json:"text"`
}

func main() {
	// Get working directory
	wd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Working directory:", wd)

	// Open the SQLite database file
	db, err := sql.Open("sqlite", wd+"/chat-app/database.db")
	if err != nil {
		log.Fatal(err)
	}

	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(db)

	// Create the Gin ruter
	router := gin.Default()

	// Creation Endpoints
	router.POST("/users", func(c *gin.Context) { createUser(c, db) })
	router.POST("/channels", func(c *gin.Context) { createChannel(c, db) })
	router.POST("/messages", func(c *gin.Context) { createMessage(c, db) })

	// Listing Endpoints
	router.GET("/channels", func(c *gin.Context) { listChannels(c, db) })
	router.GET("/messages", func(c *gin.Context) { listMessages(c, db) })

	// Login Endpoint
	router.POST("/login", func(c *gin.Context) { login(c, db) })

	err = router.Run(":8080")
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Server Listening on port:8080")
}
