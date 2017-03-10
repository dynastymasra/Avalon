package util

import (
	"fmt"

	log "github.com/Sirupsen/logrus"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/spf13/viper"
)

// Global variable database connection
var (
	Database *gorm.DB
)

// Connect database postgres
func Connect() (*gorm.DB, error) {
	config := fmt.Sprintf("postgres://%v:%v@%v/%v?sslmode=disable", viper.GetString("POSTGRES_USERNAME"),
		viper.GetString("POSTGRES_PASSWORD"), viper.GetString("POSTGRES_ADDRESS"), viper.GetString("POSTGRES_DATABASE"))
	db, err := gorm.Open("postgres", config)
	if err != nil {
		log.WithFields(log.Fields{
			"file":             "connector.go",
			"package":          "util",
			"postgre_username": viper.GetString("POSTGRES_USERNAME"),
			"postgre_password": viper.GetString("POSTGRES_PASSWORD"),
			"postgre_address":  viper.GetString("POSTGRES_ADDRESS"),
			"postgre_database": viper.GetString("POSTGRES_DATABASE"),
		}).Errorf("%v", err)
		return nil, err
	}

	db.LogMode(viper.GetString("POSTGRES_LOGGING") == "true")

	return db, nil
}
