package config

import "github.com/spf13/viper"

type Config struct {
	SECERETKEY       string `mapstructure:"JWTKEY"`
	Host             string `mapstructure:"HOST"`
	User             string `mapstructure:"USER"`
	Password         string `mapstructure:"PASSWORD"`
	Database         string `mapstructure:"DBNAME"`
	Port             string `mapstructure:"PORT"`
	Sslmode          string `mapstructure:"SSL"`
	UserPort         string `mapstructure:"USERPORT"`
	AdminPort        string `mapstructure:"ADMINPORT"`
	BookingPort      string `mapstructure:"BOOKINGPORT"`
	NotificationPort string `mapstructure:"NOTIFICATIONPORT"`
	REDISHOST        string `mapstructure:"REDISHOST"`
}

// LoadConfig reads configuration from file or environment variables.
func LoadConfig() (*Config, error) {
	// Load configuration from .env file and environment variables
	var config Config
	// Set the file name of the configurations file
	viper.SetConfigFile(".env")
	// Set the type of the configurations file
	viper.SetConfigType("env")
	// Enable automatic environment variable binding
	viper.AutomaticEnv()

	// Read the configuration file
	if err := viper.ReadInConfig(); err != nil {
		//if no env file found, continue with env variables
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			return nil, err
		}
	}
	// Set default values
	if err := viper.Unmarshal(&config); err != nil {
		return nil, err
	}

	return &config, nil
}
