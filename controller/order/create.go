package order

import (
	"avalon/config"
	"avalon/model"
	"avalon/util"
	"net/http"

	"github.com/satori/go.uuid"

	log "github.com/Sirupsen/logrus"
	"gopkg.in/gin-gonic/gin.v1"
)

// CreateOrderController function
func CreateOrderController(c *gin.Context) {
	var order model.Order

	c.Header("Content-Type", "application/json")
	if err := c.BindJSON(&order); err != nil {
		log.WithFields(log.Fields{"file": "create.go", "package": "controller.order"}).Errorf("Bind JSON %v", err)
		c.Error(err)
		c.JSON(http.StatusBadRequest, util.FailResponse(err.Error()))
		return
	}

	order.ID = uuid.NewV4().String()
	order.OrderStatus = string(model.OrderStatusNew)
	if err := util.Database.Create(&order).Error; err != nil {
		log.WithFields(log.Fields{"file": "create.go", "package": "controller.order"}).Errorf("Create %v", err)
		c.Error(err)
		c.JSON(http.StatusInternalServerError, util.FailResponse(config.ErrDatabase.Error()))
		return
	}

	c.JSON(http.StatusCreated, util.ObjectResponse(order))
}
