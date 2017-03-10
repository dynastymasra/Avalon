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

// SearchOrderController seacrh order by order id, shop id and customer id
func SearchOrderController(c *gin.Context) {
	var orders []model.Order

	search := c.Param("search")

	if err := util.Database.Where("order_status = ?", strings.ToUpper(search)).Find(&orders).Error; err != nil {
		log.WithFields(log.Fields{"file": "search.go", "package": "controller.order"}).Errorf("Where %v", err)
		c.Error(err)
		c.JSON(http.StatusInternalServerError, util.FailResponse(config.ErrDatabase.Error()))
		return
	}

	if len(orders) < 1 {
		log.WithFields(log.Fields{"file": "search.go", "package": "controller.order"}).Warning("Data order not found")
		c.Error(config.ErrorRecordNotFound)
		c.JSON(http.StatusNotFound, util.FailResponse(config.ErrorRecordNotFound.Error()))
		return
	}

	c.JSON(http.StatusOK, util.ObjectResponse(orders))
}
