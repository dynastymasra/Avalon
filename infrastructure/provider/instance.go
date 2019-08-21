package provider

import (
	"github.com/dynastymasra/avalon/infrastructure/provider/postgres"
)

type Instance struct {
	Postgres *postgres.Connector
}

func NewInstance(postgres *postgres.Connector) *Instance {
	return &Instance{Postgres: postgres}
}
