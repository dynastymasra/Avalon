package config

import "github.com/spf13/viper"

// SetDefault used to set default values for configuration
func SetDefault() {
	viper.SetDefault("ADDRESS", ":8080")
	viper.SetDefault("GIN_MODE", "release")

	// Set value of postgres
	viper.SetDefault("POSTGRES_USERNAME", "postgres")
	viper.SetDefault("POSTGRES_PASSWORD", "root")
	viper.SetDefault("POSTGRES_ADDRESS", "192.168.99.100:5432")
	viper.SetDefault("POSTGRES_DATABASE", "avalon")
	viper.SetDefault("POSTGRES_LOGGING", "false")
}
