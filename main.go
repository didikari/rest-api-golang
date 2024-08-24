package main

import (
	"myapp/controllers"
	"myapp/models"

	"github.com/gin-gonic/gin"
)

func main() {
	//inisialiasai Gin
	router := gin.Default()

	//panggil koneksi database
	models.ConnectDatabase()

	//membuat route dengan method GET
	router.GET("/", func(c *gin.Context) {

		//return response JSON
		c.JSON(200, gin.H{
			"message": "Hello World!",
		})
	})

	router.GET("/api/posts", controllers.FindPosts)

	router.POST("/api/posts", controllers.StorePost)

	router.GET("/api/posts/:id", controllers.FindPostById)

	router.PUT("/api/posts/:id", controllers.UpdatePost)

	router.DELETE("api/posts/:id", controllers.DeletePost)

	//mulai server dengan port 3000
	router.Run(":3000")
}
