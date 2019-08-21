package postgres

import (
	"errors"

	"github.com/dynastymasra/avalon/config"
	"github.com/jinzhu/gorm"
	"github.com/sirupsen/logrus"

	"github.com/matryer/resync"

	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var (
	db   *gorm.DB
	err  error
	once resync.Once
)

type Connector struct {
	Postgres *gorm.DB
}

func Connect() (*Connector, error) {
	dbURL := config.Database().ConnectionString()

	once.Do(func() {
		db, err = gorm.Open("postgres", dbURL)
		if err != nil {
			logrus.WithError(err).WithField("db_url", dbURL).Errorln("Cannot connect to DB")
			return
		}

		db.DB().SetMaxIdleConns(config.Database().MaxIdleConns())
		db.DB().SetMaxOpenConns(config.Database().MaxOpenConns())

		if err := db.DB().Ping(); err != nil {
			logrus.WithError(err).Errorln("Cannot ping database")
			return
		}

		db.LogMode(config.Database().LogEnabled())
	})

	return &Connector{Postgres: db}, err
}

func (c *Connector) Ping() error {
	if c.Postgres == nil {
		return errors.New("does't have database data")
	}
	return c.Postgres.DB().Ping()
}

func (c *Connector) Close() error {
	if c.Postgres == nil {
		return errors.New("does't have database data")
	}
	return db.Close()
}
