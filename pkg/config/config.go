package config

import (
	"fmt"

	"github.com/spf13/viper"
)

type Configure struct {
	APIPORT            string `mapstructure:"APIPORT"`
	ADMINPORT          string `mapstructure:"GRPCADMINPORT"`
	BSERVICEPORT       string `mapstructure:"BSERVICEPORT"`
	SECRETKEY          string `mapstructure:"SECRETKEY"`
	REDISHOST          string `mapstructure:"REDISHOST"`
	GRPCCORDINATORPORT string `mapstructure:"GRPCCORDINATORPORT"`
}

func LoadConfigure() (*Configure, error) {
	var cfg Configure

	viper.SetConfigFile("../../.env")
	err := viper.ReadInConfig()

	err = viper.Unmarshal(&cfg)

	if err != nil {
		return &Configure{}, nil
	}

	fmt.Println(cfg)
	return &cfg, nil
}
