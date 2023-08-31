package config

import "github.com/spf13/viper"

type MainConfig struct {
	Port      string
	ProxyPort string `mapstructure:"PROXY_PORT"`

	AuthSvcURL   string `mapstructure:"AUTH_SVC_URL"`
	CourseSvcURL string `mapstructure:"COURSE_SVC_URL"`
	ClientSvcURL string `mapstructure:"CLIENT_SVC_URL"`
}

func NewMainConfig() (c MainConfig, err error) {
	viper.AddConfigPath("./config/envs")
	viper.SetConfigName("main")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		return MainConfig{}, err
	}

	if err := viper.Unmarshal(&c); err != nil {
		return MainConfig{}, err
	}

	return
}
