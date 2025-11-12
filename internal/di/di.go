package di

import (
	"log"

	"github.com/yourorg/ultimate-rent-car/ultimate-rent-user-service/config"
)

// Init initializes the dependency injection container
func Init() {
	// Load configuration
	cnfg, err := config.LoadConfig()
	// log the config values to verify
	log.Printf("Config Loaded: %+v", cnfg)
	if err != nil {
		panic(err)
	}
}
