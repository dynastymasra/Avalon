package config

import (
	"log"
	"os"
	"strconv"

	"github.com/spf13/viper"
)

type Config struct {
	loggerFormat string
	logLevel     string
	dbConfig     DatabaseConfig
}

var config Config

func Load() {
	viper.SetDefault(envServerAddress, ":8080")
	viper.SetDefault(envLogLevel, "debug")
	viper.SetDefault(envLoggerFormat, "text")

	viper.AutomaticEnv()

	viper.SetConfigName("application")
	viper.AddConfigPath("./")
	viper.AddConfigPath("./../")
	viper.AddConfigPath("../../")
	viper.AddConfigPath("../../../")
	viper.AddConfigPath("../../../../")
	viper.SetConfigType("yaml")

	viper.ReadInConfig()

	config = Config{
		loggerFormat: getString(envLoggerFormat),
		logLevel:     getString(envLogLevel),
		dbConfig: DatabaseConfig{
			host:        getString(envDatabaseHost),
			port:        getInt(envDatabasePort),
			name:        getString(envDatabaseName),
			username:    getString(envDatabaseUsername),
			password:    getString(envDatabasePassword),
			maxOpenConn: getInt(envDatabaseMaxOpenConns),
			maxIdleConn: getInt(envDatabaseMaxIdleConns),
			logEnabled:  getBool(envDatabaseEnableLog),
		},
	}
}

func LoggerFormat() string {
	return config.loggerFormat
}

func LogLevel() string {
	return config.logLevel
}

func ServerAddress() string {
	return getString(envServerAddress)
}

func Database() DatabaseConfig {
	return config.dbConfig
}

func checkEnvKey(key string) {
	if !viper.IsSet(key) && os.Getenv(key) == "" {
		log.Fatalf("%v env key is not set", key)
	}
}

func getString(key string) string {
	checkEnvKey(key)

	return viper.GetString(key)
}

func getInt(key string) int {
	str := getString(key)

	v, err := strconv.Atoi(str)
	if err != nil {
		log.Fatalf("%v key is not valid int", key)
	}

	return v
}

func getBool(key string) bool {
	str := getString(key)

	v, err := strconv.ParseBool(str)
	if err != nil {
		log.Fatalf("%v key with value %s is not valid bool", key, str)
	}

	return v
}
