package db;

import (
	"errors"	
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	
	"iform/config"
)

var currentConnection *gorm.DB;

func InitConnection (driver string, dsn string) *gorm.DB {
	
	if driver != "postgres" {
		log.Fatal(errors.New("We currently support postgres only"));
	}

	connection, err := gorm.Open(postgres.Open(dsn), &gorm.Config{});

	if err != nil {
		log.Fatal(errors.New("Unable to connect to database!"));
	}
	
	currentConnection = connection;

	return connection;
}

func GetConnection () *gorm.DB {
	currentConfig := config.GetInstance();
	if currentConnection == nil && currentConfig.Debug {
		log.Println("Trying to access database before connecting");
	}
	return currentConnection;
}