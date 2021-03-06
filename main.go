package main

import (
	"avalon/config"
	"net/http"
	"runtime"

	"avalon/model"
	"avalon/routes"
	"avalon/util"
	"avalon/util/es"

	log "github.com/Sirupsen/logrus"
	"github.com/spf13/viper"
	gin "gopkg.in/gin-gonic/gin.v1"
)

func init() {
	config.LoggerInit()
	config.SetDefault()
	viper.AutomaticEnv()

	cpu := runtime.NumCPU()
	runtime.GOMAXPROCS(cpu)

	gin.SetMode(viper.GetString("GIN_MODE"))

	log.WithFields(log.Fields{
		"file":        "main.go",
		"package":     "main",
		"cpu":         cpu,
		"mode":        viper.GetString("GIN_MODE"),
		"postgre_log": viper.GetString("POSTGRES_LOGGING"),
		"port":        viper.GetString("ADDRESS"),
	}).Info("Success initialize all init value")
}

func main() {
	var err error
	util.Database, err = util.Connect()
	if err != nil {
		log.WithFields(log.Fields{"file": "main.go", "package": "main"}).Fatalf("Database connection %v", err)
	}
	defer util.Database.Close()

	es.ElasticConnector, err = es.NewElasticClient()
	if err != nil {
		log.WithFields(log.Fields{"file": "main.go", "package": "main"}).Fatalf("NewElasticClient %v", err)
	}

	if err := util.Database.AutoMigrate(&model.Order{}).Error; err != nil {
		log.WithFields(log.Fields{"file": "main.go", "package": "main"}).Fatalf("Auto migration %v", err)
	}

	router := gin.Default()

	router.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, util.FailResponse(config.ErrorNotFound.Error()))
	})

	v1 := router.Group("/v1")
	{
		routes.AvalonRouters(v1)
	}

	router.Run(viper.GetString("ADDRESS"))
}
