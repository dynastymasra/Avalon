package es

import (
	log "github.com/Sirupsen/logrus"
	"golang.org/x/net/context"
)

// IsIndexExist function for check index elasticsearch
func IsIndexExist(index string) (bool, error) {
	exist, err := ElasticConnector.IndexExists(index).Do(context.Background())
	if err != nil {
		log.WithFields(log.Fields{
			"file":          "index.go",
			"package":       "util.es",
			"elastic_index": index,
		}).Errorf("%v", err)
		return false, err
	}

	return exist, nil
}
