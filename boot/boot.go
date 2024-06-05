package boot;

import (
	"github.com/gin-gonic/gin"

	"iform/pkg/managers/db"
	"iform/config"
	"iform/routes"

	"log"
)

func InitServer () {
	currentConfig := config.GetInstance();

	// Step 1: Connect to database 	
	db.InitConnection(currentConfig.SQLDriver, currentConfig.StringConnection);

	if currentConfig.Debug {
		log.Println("Databse Connection Successful")
	}

	if !currentConfig.Debug {
		gin.SetMode(gin.ReleaseMode)
	}

	// Step 2: Init Routes
	server := routes.InitRoutes();

	// Step 3: Start Server
	err := server.Run(currentConfig.AppPort);

	if err != nil && currentConfig.Debug {
		log.Printf("Server Running on port: %s \n", currentConfig.AppPort)
	}
}