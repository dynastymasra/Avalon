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
		log.WithFields(log.Fields{
			"file":     "create.go",
			"package":  "controller.order",
			"function": "CreateOrderController.BindJSON",
		}).Error(err)
		c.Error(err)
		c.JSON(http.StatusBadRequest, util.FailResponse(err.Error()))
		return
	}

	order.ID = uuid.NewV4().String()
	order.OrderStatus = string(model.OrderStatusNew)
	if err := util.Database.Create(&order).Error; err != nil {
		log.WithFields(log.Fields{
			"file":     "create.go",
			"package":  "controller.order",
			"function": "CreateOrderController.Create",
		}).Error(err)
		c.Error(err)
		c.JSON(http.StatusInternalServerError, util.FailResponse(config.ErrDatabase.Error()))
		return
	}

	// Save order to elatic for search
	go func(id string, order model.Order) {
		docID, err := SaveToSearch(id, order)
		if err != nil {
			log.WithFields(log.Fields{
				"file":     "create.go",
				"package":  "controller.order",
				"function": "CreateOrderController.SaveToSearch",
			}).Warning(err)
		}

		log.WithFields(log.Fields{
			"file":     "create.go",
			"package":  "controller.order",
			"function": "CreateOrderController.SaveToSearch",
		}).Infof("Order id %v success save to search", docID)
	}(order.ShopID, order)

	c.JSON(http.StatusCreated, util.ObjectResponse(order))
}
