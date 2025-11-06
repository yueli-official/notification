package config

import (
	"fmt"

	"github.com/spf13/viper"
)

type Config struct {
	Server Server          `mapstructure:"server"`
	Email  []EmailProvider `mapstructure:"email"`
	SMS    []SMSProvider   `mapstructure:"sms"`
}

type Server struct {
	Port string `mapstructure:"port"`
}

type EmailProvider struct {
	Name     string `mapstructure:"name"`
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	Username string `mapstructure:"username"`
	Password string `mapstructure:"password"`
	From     string `mapstructure:"from"`
	Enabled  bool   `mapstructure:"enabled"`
}

type SMSProvider struct {
	Name      string `mapstructure:"name"`
	Provider  string `mapstructure:"provider"` // aliyun, tencent, etc.
	AccessKey string `mapstructure:"access_key"`
	SecretKey string `mapstructure:"secret_key"`
	SignName  string `mapstructure:"sign_name"`
	Enabled   bool   `mapstructure:"enabled"`
}

func LoadConfig() (*Config, error) {
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	viper.AddConfigPath("./config")

	// 设置环境变量
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		return nil, fmt.Errorf("读取配置文件失败: %w", err)
	}

	var config Config
	if err := viper.Unmarshal(&config); err != nil {
		return nil, fmt.Errorf("解析配置失败: %w", err)
	}

	return &config, nil
}
