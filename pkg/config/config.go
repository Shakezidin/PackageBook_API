package config

import (
	"fmt"
	"os"
)

type Configure struct {
	APIPORT            string `mapstructure:"APIPORT"`
	ADMINPORT          string `mapstructure:"GRPCADMINPORT"`
	BSERVICEPORT       string `mapstructure:"BSERVICEPORT"`
	SECRETKEY          string `mapstructure:"SECRETKEY"`
	REDISHOST          string `mapstructure:"REDISHOST"`
	GRPCCORDINATORPORT string `mapstructure:"GRPCCORDINATORPORT"`
	GRPCUSERPORT       string `mapstructure:"GRPCUSERPORT"`
}

func LoadConfigure() (*Configure, error) {
	cfg := &Configure{
		APIPORT:            os.Getenv("APIPORT"),
		ADMINPORT:          os.Getenv("GRPCADMINPORT"),
		BSERVICEPORT:       os.Getenv("BSERVICEPORT"),
		SECRETKEY:          os.Getenv("SECRETKEY"),
		REDISHOST:          os.Getenv("REDISHOST"),
		GRPCCORDINATORPORT: os.Getenv("GRPCCORDINATORPORT"),
		GRPCUSERPORT:       os.Getenv("GRPCUSERPORT"),
	}
	fmt.Println(cfg)
	return cfg, nil
}
