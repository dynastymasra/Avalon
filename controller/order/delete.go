package order

import (
	"avalon/config"
	"avalon/model"
	"avalon/util"
	"net/http"

	log "github.com/Sirupsen/logrus"
	"gopkg.in/gin-gonic/gin.v1"
)

// DeleteOrderController control
func DeleteOrderController(c *gin.Context) {
	id := c.Param("id")

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

	c.JSON(http.StatusOK, util.SuccessResponse())
}
