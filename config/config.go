package config

import (
	"fmt"
	"github.com/spf13/viper"
)

var (
	Data Config
)

type Config struct {
	UidUrl string `mapstructure:"UID_URL"`
}

func init() {
	var err error
	Data, err = LoadConfig()
	if err != nil {
		fmt.Println(fmt.Sprintf("Reading config failed: %s", err.Error()))
	}
}

func LoadConfig() (config Config, err error) {

	viper.AddConfigPath("./config")
	viper.SetConfigName("conf")
	viper.SetConfigType("yaml")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	if err != nil {
		return
	}

	return
}
