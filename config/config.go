package config

import (
	"errors"

	"github.com/spf13/viper"
)

type ConfigInterface interface {
	Read() (*Config, error)
}

// Config is
type Config struct {
	FirebaseConfig firebaseConfig
}

type firebaseConfig struct {
	ProjectId string
}

// NewConfig is
func NewConfig() ConfigInterface {
	return &Config{}
}

// Read is
func (c Config) Read() (config *Config, err error) {
	viper.SetConfigName("./configs/secret/service-account.json")
	viper.SetConfigType("json")
	viper.AddConfigPath(".")

	err = viper.ReadInConfig()
	if err != nil {
		return nil, err
	}

	// Firebase
	firebaseProjectId := viper.GetString("project_id")
	if firebaseProjectId == "" {
		return nil, errors.New("invalid project id")
	}
	c.FirebaseConfig.ProjectId = firebaseProjectId

	return &c, nil
}
