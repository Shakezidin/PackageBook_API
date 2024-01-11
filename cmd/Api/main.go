package main

import (
	"log"

	"github.com/Shakezidin/pkg/admin"
	cnfg "github.com/Shakezidin/pkg/config"
	"github.com/Shakezidin/pkg/server"
)

func main() {
	config, err := cnfg.LoadConfigure()
	if err != nil {
		log.Printf("Error Loading Config Files, error: %v", err)
	}

	server := server.Server()
	admin.NewAdminRoutes(server.R, *config)
	// coordinator.NewCoordinatorRoute(server.R, *config)
	server.StartServer(config.APIPORT)
}
