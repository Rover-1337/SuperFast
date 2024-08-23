package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	_ "github.com/glebarez/go-sqlite"
)

var secret = []byte("secret")
const userkey = "user"

func initDB() *sql.DB {
	db, err := sql.Open("sqlite", "./init.db")

	if err != nil {
		log.Fatal(err)
	}
	return db
}

func AuthRequired(c *gin.Context) {
	session := sessions.Default(c)
	user := session.Get(userkey)
	if user == nil {
		// Abort the request with the appropriate error code
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}
	c.Next()
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
	  port = "8000"
	}

	var db = initDB()
	defer db.Close()

	r := gin.Default()

    corsConfig := cors.Config{
        AllowOrigins: []string{"http://localhost:5173"},
    }

    // Use the custom CORS configuration
    r.Use(cors.New(corsConfig))


	r.Use(sessions.Sessions("session", cookie.NewStore(secret)))


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

	r.GET("/api/blogPosts/:id", func(c *gin.Context) {
		id := c.Param("id")
		row := db.QueryRow("SELECT id, title, content, description, created_at FROM blogPosts WHERE id = ?", id)

		post := make(map[string]interface{})
		var dbID int
		var title, content, description, createdAt string

		err := row.Scan(&dbID, &title, &content, &description, &createdAt)

		if err != nil {
			log.Fatal(err)
		}

		post["id"] = dbID
		post["title"] = title
		post["content"] = content
		post["description"] = description
		post["created_at"] = createdAt

		c.JSON(http.StatusOK, gin.H{"data": post})
	})

	r.GET("/api/deleteBlogPost/:id", func(c *gin.Context) {
		id := c.Param("id")
		_, err := db.Exec("DELETE FROM blogPosts WHERE id = ?", id)

		if err != nil {
			log.Fatal(err)
		}

		c.JSON(http.StatusOK, gin.H{"message": "Post deleted"})
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

	r.POST("/api/login", func(c *gin.Context) {
		session := sessions.Default(c)

		var user struct {
			Username string `json:"username" binding:"required"`
			Password string `json:"password" binding:"required"`
		}

		if err := c.BindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		result, err := db.Exec("SELECT * FROM users WHERE username = ? AND password = ?", user.Username, user.Password)

		if err != nil {
			log.Fatal(err)
		}

		id, err := result.LastInsertId()
		if err != nil {
			log.Fatal(err)
		}

		session.Set(userkey, id)
		session.Save()

		c.JSON(http.StatusFound, gin.H{"loggedin": true})
	})

	private := r.Group("/api/auth")
	private.Use(AuthRequired)
	{
		private.GET("/dashboard", func(ctx *gin.Context) {
			var users []map[string]int16

			rows, err := db.Query("SELECT username FROM users")

			if err != nil {
				log.Fatal(err)
			}

			for rows.Next() {
				var username int16
				if err := rows.Scan(&username); err != nil {
					log.Fatal(err)
				}
				username *= 2
				users = append(users, map[string]int16{"username": username})
			}

			ctx.JSON(http.StatusOK, gin.H{"data": users})
		})
		private.GET("/status", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{"loggedin": true})
		})
	}

	// Listen and serve on 0.0.0.0:8080
	r.Run(fmt.Sprintf(":%s", port))
}