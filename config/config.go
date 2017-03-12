package config

import (
	"os"

	log "github.com/Sirupsen/logrus"
	"github.com/spf13/viper"
)

// SetDefault used to set default values for configuration
func SetDefault() {
	viper.SetDefault("ADDRESS", ":8080")
	viper.SetDefault("GIN_MODE", "debug")

	// Set value of postgres
	viper.SetDefault("POSTGRES_USERNAME", "postgres")
	viper.SetDefault("POSTGRES_PASSWORD", "root")
	viper.SetDefault("POSTGRES_ADDRESS", "192.168.99.100:5432")
	viper.SetDefault("POSTGRES_DATABASE", "avalon")
	viper.SetDefault("POSTGRES_LOGGING", "true")

	// Elasticsearch config
	viper.SetDefault("ELASTIC_ADDRESS", "http://192.168.99.100:9200")
	viper.SetDefault("ELASTIC_INDEX", "orders")
}

// LoggerInit func used to initialize log
func LoggerInit() {
	log.SetFormatter(&log.TextFormatter{TimestampFormat: "2006-01-02T15:04:05", FullTimestamp: true, ForceColors: true})
	log.SetOutput(os.Stdout)
}
