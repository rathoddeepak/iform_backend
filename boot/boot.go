package boot;

import (
	"github.com/gin-gonic/gin"

	"iform/internal/migrations"
	"iform/pkg/managers/db"	
	"iform/config"
	"iform/routes"

	"log"
	"os"
)

func initDB () *config.CacheConfig {
	currentConfig := config.GetInstance();	
	db.InitConnection(currentConfig.SQLDriver, os.Getenv("DATABASE_URL"));

	if currentConfig.Debug {
		log.Println("Databse Connection Successful")
	}

	if !currentConfig.Debug {
		gin.SetMode(gin.ReleaseMode)
	}

	return currentConfig;
}

func InitServer () {
	// Step 1: Connect to database
	currentConfig := initDB();

	// Step 2: Init Routes
	server := routes.InitRoutes();

	// Step 3: Start Server
	err := server.Run(currentConfig.AppPort);

	if err != nil && currentConfig.Debug {
		log.Printf("Server Running on port: %s \n", currentConfig.AppPort)
	}
}

func MigrateDatabase () {
	
	initDB();

	err := migrations.StartAutoMigrate();
	if err != nil {
		log.Fatal("Unable to migrate database: ", err);
	} else {
		log.Println("Database migrated successfully!");
	}

}