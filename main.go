package main

import (
	"log"
	"log/slog"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	logFile, err := os.OpenFile("employee-api.log",
		os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal(err)
	}
	defer logFile.Close()

	router := gin.Default()

	router.POST("/employee")
	router.GET("/employee")
	router.GET("/employee/:id")
	router.PUT("/employee/:id")
	router.DELETE("/employee/:id")

	// throw the port 80
	// default is 8080
	router.Run(":80")

	logger := slog.New(slog.NewJSONHandler(logFile, nil))
}
