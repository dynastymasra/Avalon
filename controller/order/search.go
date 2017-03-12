package order

import (
	"avalon/model"
	modelEs "avalon/model/es"
	"avalon/util"
	"avalon/util/es"
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"gopkg.in/olivere/elastic.v5"

	log "github.com/Sirupsen/logrus"

	"github.com/spf13/viper"
	"gopkg.in/gin-gonic/gin.v1"
)

func createSearchIndex(consumerID string) string {
	mapper, err := modelEs.BuildMappingToString(modelEs.SearchSetting(), modelEs.SearchMappings(consumerID))
	if err != nil {
		log.WithFields(log.Fields{
			"file":     "search.go",
			"package":  "controller.order",
			"function": "createSearchIndex.BuildMappingToString",
		}).Error(err)
		return ""
	}

	if err := es.CreateIndex(viper.GetString("ELASTIC_INDEX"), mapper); err != nil {
		log.WithFields(log.Fields{
			"file":     "search.go",
			"package":  "controller.order",
			"function": "createSearchIndex.BuildMappingToString",
			"mapper":   mapper,
		}).Error(err)
	}

	return mapper
}

// SearchOrderController seacrh order by order id, shop id and customer id
func SearchOrderController(c *gin.Context) {
	var orders []model.Order

	shopID := c.Param("shopId")
	val := c.Query("search")

	if val == "" {
		log.WithFields(log.Fields{
			"file":     "search.go",
			"package":  "controller.order",
			"function": "SearchOrderController.Query",
		}).Error("Query is empty")
		c.Error(fmt.Errorf("Query is empty"))
		c.JSON(http.StatusBadRequest, util.FailResponse("Query is empty"))
		return
	}

	search := es.ElasticConnector.Search().Index(viper.GetString("ELASTIC_INDEX")).Type(shopID)
	query := elastic.NewQueryStringQuery(val)

	result, err := search.Query(query).Do(context.Background())
	if err != nil {
		log.WithFields(log.Fields{
			"file":     "search.go",
			"package":  "controller.order",
			"function": "SearchOrderController.Query",
			"query":    query,
		}).Error(err)
		c.Error(err)
		c.JSON(http.StatusBadRequest, util.FailResponse(err.Error()))
		return
	}

	for _, hit := range result.Hits.Hits {
		doc := model.Order{}
		if err := json.Unmarshal(*hit.Source, &doc); err != nil {
			log.WithFields(log.Fields{
				"file":     "search.go",
				"package":  "controller.order",
				"function": "SearchOrderController.Unmarshal",
			}).Error(err)
			continue
		}
		orders = append(orders, doc)
	}

	c.JSON(http.StatusOK, util.ObjectResponse(orders))
}

// SaveToSearch is function save data to elastic
func SaveToSearch(consumerID string, order model.Order) (string, error) {
	createSearchIndex(consumerID)

	payload, err := json.Marshal(order)
	if err != nil {
		log.WithFields(log.Fields{
			"file":     "search.go",
			"package":  "controller.order",
			"function": "SaveToSearch.Marshal",
		}).Error(err)
		return "", err
	}

	doc, err := es.ElasticConnector.Index().Index(viper.GetString("ELASTIC_INDEX")).Type(consumerID).Id(order.ID).
		BodyString(string(payload)).Do(context.Background())
	if err != nil {
		log.WithFields(log.Fields{
			"file":     "search.go",
			"package":  "controller.order",
			"function": "SaveToSearch.Index",
		}).Error(err)
		return "", err
	}

	return doc.Id, nil
}
