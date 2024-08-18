package main

import "github.com/gin-gonic/gin"

func main() {
	router := gin.Default()

	router.POST("/employee")
	router.GET("/employee")
	router.GET("/employee/:id")
	router.PUT("/employee/:id")
	router.DELETE("/employee/:id")

	// throw the port 80
	// default is 8080
	router.Run(":80")
}
