package order

import (
	"avalon/config"
	"avalon/model"
	"avalon/util"
	"net/http"

	log "github.com/Sirupsen/logrus"
	"gopkg.in/gin-gonic/gin.v1"
)

// UpdateOrderController controller
func UpdateOrderController(c *gin.Context) {
	var order model.Order

	id := c.Param("id")

	if validUUID := util.IsValidUUID(id); !validUUID {
		log.WithFields(log.Fields{"file": "update.go", "package": "controller.order"}).Errorf("IsValidUUID %v", config.ErrorNotValidUUID)
		c.Error(config.ErrorNotValidUUID)
		c.JSON(http.StatusBadRequest, util.FailResponse(config.ErrorNotValidUUID.Error()))
		return
	}

	c.Header("Content-Type", "application/json")
	if err := c.BindJSON(&order); err != nil {
		log.WithFields(log.Fields{"file": "update.go", "package": "controller.order"}).Errorf("BindJSON %v", err)
		c.Error(err)
		c.JSON(http.StatusBadRequest, util.FailResponse(err.Error()))
		return
	}

	var or model.Order
	recordNotFound := util.Database.Where("id = ?", id).First(&or).RecordNotFound()
	if recordNotFound {
		log.WithFields(log.Fields{"file": "update.go", "package": "controller.order"}).Warningf("Order id %v not found", id)
		c.Error(config.ErrorRecordNotFound)
		c.JSON(http.StatusNotFound, util.FailResponse(config.ErrorRecordNotFound.Error()))
		return
	}

	order.ID = id
	if err := util.Database.Save(&order).Error; err != nil {
		log.WithFields(log.Fields{"file": "update.go", "package": "controller.order"}).Errorf("Save %v", err)
		c.Error(err)
		c.JSON(http.StatusInternalServerError, util.FailResponse(config.ErrDatabase.Error()))
		return
	}

	c.JSON(http.StatusOK, util.ObjectResponse(order))
}
