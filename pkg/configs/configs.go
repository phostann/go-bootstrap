package configs

import (
	"github.com/spf13/viper"
)

const configPath = "./configs"

type Config struct {
	Server   ServerConfig   `json:"server" mapstructure:"server"`
	Database DatabaseConfig `json:"database" mapstructure:"database"`
	JWT      JWTConfig      `json:"jwt" mapstructure:"jwt"`
}

type ServerConfig struct {
	Host string `json:"host" mapstructure:"host"`
	Port string `json:"port" mapstructure:"port"`
}

type DatabaseConfig struct {
	Driver    string `json:"driver" mapstructure:"driver"`
	Host      string `json:"host" mapstructure:"host"`
	Port      int    `json:"port" mapstructure:"port"`
	User      string `json:"user" mapstructure:"user"`
	Password  string `json:"password" mapstructure:"password"`
	DB        string `json:"db" mapstructure:"db"`
	Migration string `json:"migration" mapstructure:"migration"`
}

type JWTConfig struct {
	Secret string `json:"secret" mapstructure:"secret"`
	Issuer string `json:"issuer" mapstructure:"issuer"`
}

func Parse() (*Config, error) {
	var config Config
	err := ReadConfig(&config)
	return &config, err
}

// ReadConfig use viper to read config file
func ReadConfig(cfg *Config) error {
	viper.SetConfigName("configs")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(configPath)
	err := viper.ReadInConfig()
	if err != nil {
		return err
	}
	return viper.Unmarshal(cfg)
}
