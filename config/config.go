package config

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/spf13/viper"
)

type Config struct {
	Server Server          `mapstructure:"server"`
	Email  []EmailProvider `mapstructure:"email"`
	SMS    []SMSProvider   `mapstructure:"sms"`
}

type Server struct {
	Port   string `mapstructure:"port"`
	APIKey string `mapstructure:"api_key"`
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

	wd, _ := os.Getwd()

	parent := filepath.Dir(wd) // 获取上一级目录

	viper.SetConfigName("config")
	viper.AddConfigPath(".")                             // 当前目录
	viper.AddConfigPath("./config")                      // 当前目录下的 config/
	viper.AddConfigPath(parent)                          // 上一级目录
	viper.AddConfigPath(filepath.Join(parent, "config")) // 上级目录的 config/

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		return nil, fmt.Errorf("读取配置文件失败: %w", err)
	}

	var cfg Config
	if err := viper.Unmarshal(&cfg); err != nil {
		return nil, fmt.Errorf("解析配置失败: %w", err)
	}

	log.Printf("使用的配置文件: %s", viper.ConfigFileUsed())

	return &cfg, nil
}
