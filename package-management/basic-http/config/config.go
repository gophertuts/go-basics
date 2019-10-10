package config

import "github.com/spf13/viper"

const (
	listen = "listen"
)

type Manager struct{}

func New() (Manager, error) {
	viper.SetConfigFile("config/config.yaml")
	viper.SetConfigType("yaml")
	if err := viper.ReadInConfig(); err != nil {
		return Manager{}, err
	}
	return Manager{}, nil
}

func (m Manager) Listen() string {
	return viper.GetString(listen)
}
