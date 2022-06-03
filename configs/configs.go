package configs

import (
	"github.com/spf13/viper"
	"log"
)

type Config struct {
	GRPC `yaml:"grpc"`
}

type GRPC struct {
	Addr string `yaml:"addr"`
	Port string `yaml:"port"`
}

func InitConfigs() Config {
	viper.SetConfigName("configs")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("configs")
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("init config error:%v", err)
	}
	var configs Config
	if err := viper.Unmarshal(&configs); err != nil {
		log.Fatalf("init config error:%v", err)
	}
	return configs
}
