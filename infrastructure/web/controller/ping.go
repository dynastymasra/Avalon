package controller

import (
	"fmt"
	"net/http"

	"github.com/dynastymasra/avalon/infrastructure/provider"

	"github.com/dynastymasra/avalon/config"

	"github.com/sirupsen/logrus"

	"github.com/dynastymasra/avalon/infrastructure/web/formatter"
)

func Ping(provider *provider.Instance) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json; charset=utf-8")

		log := logrus.WithField(config.RequestID, r.Context().Value(config.HeaderRequestID))

		if err := provider.Postgres.Ping(); err != nil {
			log.WithError(err).Infoln("Failed ping postgres")

			w.WriteHeader(http.StatusServiceUnavailable)
			fmt.Fprintf(w, formatter.FailResponse(err.Error()).Stringify())
			return
		}

		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, formatter.SuccessResponse().Stringify())
	}
}
