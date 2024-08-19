package main

import (
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	logger, err := InitLogger()
	if err != nil {
		log.Fatal("failed to init logger: %w", err)
	}

	mapMemoryStorage := NewMapMemoryStorage()
	handler := NewHandler(mapMemoryStorage, logger)

	router := gin.Default()

	router.POST("/employee", handler.CreateEmployee)
	router.GET("/employee", handler.GetAllEmployees)
	router.GET("/employee/:id", handler.GetEmployee)
	router.PUT("/employee/:id", handler.UpdateEmployee)
	router.DELETE("/employee/:id", handler.DeleteEmployee)

	// throw the port 80
	// default is 8080
	router.Run(":80")
}
