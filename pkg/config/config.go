package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

// Configure represents the configuration parameters.
type Configure struct {
	APIPORT            string `mapstructure:"APIPORT"`
	ADMINPORT          string `mapstructure:"GRPCADMINPORT"`
	GRPCUSERPORT       string `mapstructure:"GRPCUSERPORT"`
	GRPCCORDINATORPORT string `mapstructure:"GRPCCORDINATORPORT"`
}

// LoadConfigure loads configuration from environment variables.
func LoadConfigure() (*Configure, error) {
	err := godotenv.Load("../../.env")
	if err != nil {
		fmt.Println("Error loading the .env file:", err)
		os.Exit(1)
	}
	cfg := &Configure{
		APIPORT:            os.Getenv("APIPORT"),
		ADMINPORT:          os.Getenv("GRPCADMINPORT"),
		GRPCUSERPORT:       os.Getenv("GRPCUSERPORT"),
		GRPCCORDINATORPORT: os.Getenv("GRPCCORDINATORPORT"),
	}

	return cfg, nil
}
