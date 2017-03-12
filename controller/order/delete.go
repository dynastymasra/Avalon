package order

import (
	"avalon/config"
	"avalon/model"
	"avalon/util"
	"avalon/util/es"
	"context"
	"net/http"

	log "github.com/Sirupsen/logrus"
	"github.com/spf13/viper"
	"gopkg.in/gin-gonic/gin.v1"
)

// DeleteOrderController control
func DeleteOrderController(c *gin.Context) {
	id := c.Param("id")
	header := c.Request.Header.Get("X-Consumer-ID")

	if validUUID := util.IsValidUUID(id); !validUUID {
		log.WithFields(log.Fields{"file": "delete.go", "package": "controller.order"}).Errorf("IsValidUUID %v", config.ErrorNotValidUUID)
		c.Error(config.ErrorNotValidUUID)
		c.JSON(http.StatusBadRequest, util.FailResponse(config.ErrorNotValidUUID.Error()))
		return
	}

	var order model.Order
	recordNotFound := util.Database.Where("id = ?", id).First(&order).RecordNotFound()
	if recordNotFound {
		log.WithFields(log.Fields{"file": "delete.go", "package": "controller.order"}).Warningf("Order id %v not found", id)
		c.Error(config.ErrorRecordNotFound)
		c.JSON(http.StatusNotFound, util.FailResponse(config.ErrorRecordNotFound.Error()))
		return
	}

	if err := util.Database.Delete(&order).Error; err != nil {
		log.WithFields(log.Fields{"file": "delete.go", "package": "controller.order"}).Errorf("Delete %v", err)
		c.Error(err)
		c.JSON(http.StatusInternalServerError, util.FailResponse(config.ErrDatabase.Error()))
		return
	}

	// Update order to elastic for search
	go func(consumer, id string) {
		deleted, err := es.ElasticConnector.Delete().Index(viper.GetString("ELASTIC_INDEX")).Type(consumer).Id(id).Do(context.Background())
		if err != nil {
			log.WithFields(log.Fields{
				"file":     "delete.go",
				"package":  "controller.order",
				"function": "DeleteOrderController.Delete",
			}).Warning(err)
		}

		if !deleted.Found {
			log.WithFields(log.Fields{
				"file":     "delete.go",
				"package":  "controller.order",
				"function": "DeleteOrderController.deleted",
				"deleted":  deleted.Found,
			}).Warning("Deleted order not found in search")
		}
	}(header, id)

	c.JSON(http.StatusOK, util.SuccessResponse())
}
