package es

import (
	"fmt"

	"avalon/config"

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
			"function":      "IsIndexExist.IndexExists",
			"elastic_index": index,
		}).Errorf("%v", err)
		return false, err
	}

	return exist, nil
}

// CreateIndex is function for create index in elasticsearch
func CreateIndex(index string, mapping ...string) error {
	isExist, err := IsIndexExist(index)
	if err != nil {
		log.WithFields(log.Fields{
			"file":          "index.go",
			"package":       "util.es",
			"function":      "CreateIndex.IsIndexExist",
			"elastic_index": index,
		}).Errorf("%v", err)

		return err
	}

	if isExist {
		log.WithFields(log.Fields{
			"file":          "index.go",
			"package":       "util.es",
			"function":      "CreateIndex.isExist",
			"elastic_index": index,
			"isExist":       isExist,
		}).Warning("%v", err)

		return fmt.Errorf("Index is exist %v, cannot create again", isExist)
	}

	if mapping != nil {
		indexCreated, err := ElasticConnector.CreateIndex(index).BodyString(mapping[0]).Do(context.Background())
		if err != nil {
			log.WithFields(log.Fields{
				"file":          "index.go",
				"package":       "util.es",
				"function":      "CreateIndex.CreateIndex",
				"elastic_index": index,
			}).Error(err)

			return err
		}

		if !indexCreated.Acknowledged {
			log.WithFields(log.Fields{
				"file":          "index.go",
				"package":       "util.es",
				"function":      "CreateIndex.Acknowledged",
				"elastic_index": index,
				"acknowledged":  indexCreated.Acknowledged,
			}).Error(err)

			return config.ErrorNotAcknowledgedIndex
		}

		return nil
	}

	indexCreated, err := ElasticConnector.CreateIndex(index).Do(context.Background())
	if err != nil {
		log.WithFields(log.Fields{
			"file":          "index.go",
			"package":       "util.es",
			"function":      "CreateIndex.CreateIndex",
			"elastic_index": index,
		}).Error(err)

		return err
	}

	if !indexCreated.Acknowledged {
		log.WithFields(log.Fields{
			"file":          "index.go",
			"package":       "util.es",
			"function":      "CreateIndex.Acknowledged",
			"elastic_index": index,
			"acknowledged":  indexCreated.Acknowledged,
		}).Error(err)

		return config.ErrorNotAcknowledgedIndex
	}

	return nil
}
