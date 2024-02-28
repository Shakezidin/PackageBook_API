package config

import (
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
	godotenv.Load("../../.env")

	cfg := &Configure{
		APIPORT:            os.Getenv("APIPORT"),
		ADMINPORT:          os.Getenv("GRPCADMINPORT"),
		GRPCUSERPORT:       os.Getenv("GRPCUSERPORT"),
		GRPCCORDINATORPORT: os.Getenv("GRPCCORDINATORPORT"),
	}

	return cfg, nil
}
