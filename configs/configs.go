package configs

import (
	"fmt"
	"github.com/spf13/viper"
)

type Config struct {
	App      `yaml:"app"`
	GRPC     `yaml:"grpc"`
	Firebase `yaml:"firebase"`
}

type App struct {
	Name string `yaml:"name"`
}

type GRPC struct {
	Addr string `yaml:"addr"`
	Port string `yaml:"port"`
}

type Firebase struct {
	AgentFile string `yaml:"agentFile"`
}

func InitConfigs() (Config, error) {
	viper.SetConfigName("configs")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("configs")
	if err := viper.ReadInConfig(); err != nil {
		return Config{}, fmt.Errorf("init config error:%v", err)
	}
	var configs Config
	if err := viper.Unmarshal(&configs); err != nil {
		return Config{}, err
	}
	return configs, nil
}
