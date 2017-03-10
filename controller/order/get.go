package order

import (
	"avalon/config"
	"avalon/model"
	"avalon/util"
	"net/http"
	"strings"

	log "github.com/Sirupsen/logrus"
	"gopkg.in/gin-gonic/gin.v1"
)

// GetAllOrderController controller
func GetAllOrderController(c *gin.Context) {
	orderStatus := c.Query("status")
	orders := []model.Order{}

	if orderStatus != "" {
		if err := model.CheckOrderStatus(orderStatus); err != nil {
			log.WithFields(log.Fields{"file": "get.go", "package": "controller.order"}).Errorf("CheckOrderStatus %v", err)
			c.Error(err)
			c.JSON(http.StatusBadRequest, util.FailResponse(err.Error()))
			return
		}

		if err := util.Database.Where("order_status = ?", strings.ToUpper(orderStatus)).Find(&orders).Error; err != nil {
			log.WithFields(log.Fields{"file": "get.go", "package": "controller.order"}).Errorf("Find %v", err)
			c.Error(err)
			c.JSON(http.StatusInternalServerError, util.FailResponse(config.ErrDatabase.Error()))
			return
		}
	} else {
		if err := util.Database.Find(&orders).Error; err != nil {
			log.WithFields(log.Fields{"file": "get.go", "package": "controller.order"}).Errorf("Find %v", err)
			c.Error(err)
			c.JSON(http.StatusInternalServerError, util.FailResponse(config.ErrDatabase.Error()))
			return
		}
	}

	if len(orders) < 1 {
		log.WithFields(log.Fields{"file": "get.go", "package": "controller.order"}).Warning("Data order not found")
		c.Error(config.ErrorRecordNotFound)
		c.JSON(http.StatusNotFound, util.FailResponse(config.ErrorRecordNotFound.Error()))
		return
	}

	c.JSON(http.StatusOK, util.ObjectResponse(orders))
}
