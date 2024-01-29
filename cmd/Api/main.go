package main

import (
	"log"

	"github.com/Shakezidin/pkg/admin"
	cnfg "github.com/Shakezidin/pkg/config"
	"github.com/Shakezidin/pkg/coordinator"
	"github.com/Shakezidin/pkg/server"
	"github.com/Shakezidin/pkg/user"
)

func main() {
	config, err := cnfg.LoadConfigure()
	if err != nil {
		log.Printf("Error Loading Config Files, error: %v", err)
	}

	server := server.Server()
	server.R.LoadHTMLGlob("../../templates/*")
	admin.NewAdminRoutes(server.R, *config)
	coordinator.NewCoordinatorRoute(server.R, *config)
	user.NewUserRoute(server.R, *config)
	server.StartServer(config.APIPORT)
}
