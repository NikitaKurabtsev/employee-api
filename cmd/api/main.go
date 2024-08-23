package main

import (
	"github.com/NikitaKurabtsev/employee-api/internal/handler"
	"github.com/NikitaKurabtsev/employee-api/internal/repository"
	"github.com/NikitaKurabtsev/employee-api/internal/responses"
	"github.com/NikitaKurabtsev/employee-api/internal/validation"
	"github.com/NikitaKurabtsev/employee-api/pkg/applog"
	"github.com/gin-gonic/gin"
	"log"
	"os"
)

func main() {
	fileLogger, err := applog.NewFileLogger("api_log.log")
	if err != nil {
		log.Fatalf("failed to init logger: %v", err)
	}

	gin.DisableConsoleColor()
	gin.DefaultWriter = os.Stdout

	employeeRepository := repository.NewEmployeeRepository()
	employeeResponse := responses.NewEmployeeResponse()
	employeeValidator := validation.NewEmployeeValidator()
	employeeHandler := handler.NewHandler(employeeRepository, employeeResponse, fileLogger, employeeValidator)

	router := gin.Default()

	router.Use(applog.ConsoleLogger())
	router.Use(gin.Logger())

	router.Use(gin.Recovery())

	router.POST("/employee", employeeHandler.CreateEmployee)
	router.GET("/employee", employeeHandler.GetAllEmployees)
	router.GET("/employee/:id", employeeHandler.GetEmployee)
	router.PUT("/employee/:id", employeeHandler.UpdateEmployee)
	router.DELETE("/employee/:id", employeeHandler.DeleteEmployee)

	// throw the port :80
	// :8080 is default
	router.Run(":80")
}
