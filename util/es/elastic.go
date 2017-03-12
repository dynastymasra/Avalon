package es

import (
	"github.com/spf13/viper"

	log "github.com/Sirupsen/logrus"
	elastic "gopkg.in/olivere/elastic.v5"
)

var (
	// ElasticConnector is elasticsearch client
	ElasticConnector *elastic.Client
)

// NewElasticClient function for create new connection elasticsearch
func NewElasticClient() (*elastic.Client, error) {
	connect, err := elastic.NewClient(elastic.SetURL(viper.GetString("ELASTIC_ADDRESS")), elastic.SetSniff(false))
	if err != nil {
		log.WithFields(log.Fields{
			"file":            "elastic.go",
			"package":         "util.es",
			"elastic_address": viper.GetString("ELASTIC_ADDRESS"),
		}).Errorf("%v", err)

		return nil, err
	}

	return connect, nil
}
