package main

import (
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/odanaraujo/golang/hexagonal-news-api/adapter/input/controller/routes"
	"github.com/odanaraujo/golang/hexagonal-news-api/configurations/exception"
	"github.com/odanaraujo/golang/hexagonal-news-api/configurations/logger"
)

func main() {
	logger.Info("Starting the application...")
	if err := godotenv.Load(); err != nil {
		logger.Error("Error trying to load the env file", err)
		exception.NewInternalServerError("Error trying to load the env file")
		return
	}
	r := gin.Default()

	routes.InitRoutes(r)

	if err := r.Run(":8080"); err != nil {
		logger.Error("Error trying to start the application on port 8080", err)
		exception.NewInternalServerError("Error trying to start the application")
		return
	}
}
