package main

import (
	"log"
	"os"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"github.com/bersennaidoo/muzeekstore/backend/config"
	"github.com/bersennaidoo/muzeekstore/backend/server"
)

func main() {
	log.Println("Starting Muzeekstore App")
	
	log.Println("Initializing configuration")
	config := config.InitConfig(getConfigFileName())

	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	log.Println("Initialize HTTP server")
	httpServer := server.InitHttpServer(config, db)

	httpServer.Start()
}

func getConfigFileName() string {
	env := os.Getenv("ENV")

	if env != "" {
		return "muzeekstore-" + env
	}

	return "muzeekstore"
}