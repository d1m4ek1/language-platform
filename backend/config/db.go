package config

import (
	"english/backend/shared/errorlog"
	"fmt"
	"github.com/spf13/viper"
)

type DBConfig struct {
	DataBaseURL     string `mapstructure:"DB_URL"`
	Host            string `mapstructure:"DB_HOST"`
	LocalPort       string `mapstructure:"DB_LOCAL_PORT"`
	DockerLocalPort string `mapstructure:"DB_DOCKER_LOCAL_PORT"`
}

func (c *DBConfig) SetDataBaseURL(isDocker bool) {
	if isDocker {
		c.DataBaseURL += fmt.Sprintf(" %s %s", c.Host, c.DockerLocalPort)
		return
	}

	c.DataBaseURL += fmt.Sprintf(" %s", c.DockerLocalPort)
}

func NewDBConfig() (c *DBConfig, err error) {
	viper.AddConfigPath("./config/envs")
	viper.SetConfigName("db")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		return &DBConfig{}, errorlog.NewError(err.Error())
	}

	if err := viper.Unmarshal(&c); err != nil {
		return &DBConfig{}, errorlog.NewError(err.Error())
	}

	return
}
