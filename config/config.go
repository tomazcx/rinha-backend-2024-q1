package config

import "github.com/spf13/viper"

type AppConfig struct {
	DBName     string `mapstructure:"DB_NAME"`
	DBPort     string `mapstructure:"DB_PORT"`
	DBUser     string `mapstructure:"DB_USER"`
	DBPassword string `mapstructure:"DB_PASSWORD"`
	DBHost     string `mapstructure:"DB_HOST"`
	WebPort    string `mapstructure:"WEB_PORT"`
}

var appCfg AppConfig

func LoadConfig() (*AppConfig, error) {
	viper.SetConfigName("app_config")
	viper.SetConfigType("env")
	viper.AddConfigPath(".")
	viper.SetConfigFile(".env")
	viper.AutomaticEnv()

	err := viper.ReadInConfig()

	if err != nil {
		return nil, err
	}

	err = viper.Unmarshal(&appCfg)

	if err != nil {
		return nil, err
	}

	return &appCfg, nil
}
