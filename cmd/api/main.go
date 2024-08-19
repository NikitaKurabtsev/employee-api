package main

import (
	"log"

	"github.com/NikitaKurabtsev/employee-api.git/internal/handler"
	"github.com/NikitaKurabtsev/employee-api.git/internal/repository"
	"github.com/NikitaKurabtsev/employee-api.git/pkg/logger"
	"github.com/gin-gonic/gin"
)

func main() {
	logger, err := logger.InitLogger()
	if err != nil {
		log.Fatal("failed to init logger: %w", err)
	}

	employeeRepository := repository.NewEmployeeRepository()

	handler := handler.NewHandler(employeeRepository, logger)

	router := gin.Default()

	router.POST("/employee", handler.CreateEmployee)
	router.GET("/employee", handler.GetAllEmployees)
	router.GET("/employee/:id", handler.GetEmployee)
	router.PUT("/employee/:id", handler.UpdateEmployee)
	router.DELETE("/employee/:id", handler.DeleteEmployee)

	// throw the port :80
	// :8080 is default
	router.Run(":80")
}
