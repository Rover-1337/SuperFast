package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	_ "github.com/glebarez/go-sqlite"
)

func initDB() *sql.DB {
	db, err := sql.Open("sqlite", "./init.db")

	if err != nil {
		log.Fatal(err)
	}
	return db
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
	  port = "8000"
	}

	var db = initDB()
	defer db.Close()

	r := gin.Default()
	r.Use(cors.Default())

	r.GET("/api/blogPosts", func(c *gin.Context) {
		rows, err := db.Query("SELECT id, title, content FROM blogPosts")
		if err != nil {
			log.Fatal(err)
		}
		defer rows.Close()

		var posts []map[string]string

		for rows.Next() {
			var id, title, content string
			if err := rows.Scan(&id, &title, &content); err != nil {
				log.Fatal(err)
			}
			posts = append(posts, map[string]string{"id": id, "title": title, "content": content})
		}
		if err := rows.Err(); err != nil {
			log.Fatal(err)
		}

		c.JSON(http.StatusOK, gin.H{"data": posts})
	})


	r.POST("/api/blogPosts", func(c *gin.Context) {
		var post struct {
			Title   string `json:"title" binding:"required"`
			Description string `json:"description" binding:"required"`
			Content string `json:"content" binding:"required"`
		}
		if err := c.BindJSON(&post); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		result, err := db.Exec("INSERT INTO blogPosts (title, description, content, created_at) VALUES (?, ?, ?, ?)", post.Title, post.Description, post.Content, time.Now())

		if err != nil {
			log.Fatal(err)
		}

		id, err := result.LastInsertId()
		if err != nil {
			log.Fatal(err)
		}

		c.JSON(http.StatusCreated, gin.H{"id": id})
	})

	// Listen and serve on 0.0.0.0:8080
	r.Run(fmt.Sprintf(":%s", port))
}