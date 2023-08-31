package config

import (
	"english/backend/shared/errorlog"
	"github.com/spf13/viper"
)

type SVCConfig struct {
	Port         string `mapstructure:"PORT"`
	ProxyPort    string `mapstructure:"PROXY_PORT"`
	JWTSecretKey string `mapstructure:"JWT_SECRET_KEY"`
	IsDocker     bool   `mapstructure:"IS_DOCKER"`
	DBConfig     *DBConfig
}

func NewSVCConfig(envfilename string) (c SVCConfig, err error) {
	viper.AddConfigPath("./config/envs")
	viper.SetConfigName(envfilename)
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		return SVCConfig{}, errorlog.NewError(err.Error())
	}

	if err := viper.Unmarshal(&c); err != nil {
		return SVCConfig{}, errorlog.NewError(err.Error())
	}

	return
}
