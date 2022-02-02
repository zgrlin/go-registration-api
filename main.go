package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Welcome to the Registration API",
		})
	})

	r.POST("/register", func(c *gin.Context) {
		fullname := c.PostForm("fullname")
		phone := c.PostForm("phone")
		email := c.PostForm("email")

		c.JSON(200, gin.H{
			"status":   "success",
			"message":  "User details successfully posted",
			"fullname": fullname,
			"phone":    phone,
			"email":    email,
		})
	})
	r.Run()
}
