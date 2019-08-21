package order

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/dynastymasra/avalon/config"
	"github.com/dynastymasra/avalon/domain"
	"github.com/dynastymasra/avalon/infrastructure/web/formatter"
	"github.com/dynastymasra/avalon/service"
	"github.com/sirupsen/logrus"
	"gopkg.in/go-playground/validator.v9"
)

func Save(service service.Instance) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		var requestBody domain.Order

		log := logrus.WithFields(logrus.Fields{
			config.RequestID: r.Context().Value(config.HeaderRequestID),
		})

		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			log.WithError(err).Errorln("Unable to read request body")

			response := formatter.FailResponse(err.Error()).Stringify()
			formatter.WriteResponse(w, r, http.StatusBadRequest, response)
			return
		}

		if err := json.Unmarshal(body, &requestBody); err != nil {
			log.WithError(err).WithField("body", string(body)).Errorln("Unable to parse request body")

			response := formatter.FailResponse(err.Error()).Stringify()
			formatter.WriteResponse(w, r, http.StatusBadRequest, response)
			return
		}

		validate := validator.New()
		if err := validate.Struct(&requestBody); err != nil {
			log.WithError(err).WithField("body", requestBody).Errorln("Failed validate order request")

			response := formatter.FailResponse(err.Error()).Stringify()
			formatter.WriteResponse(w, r, http.StatusBadRequest, response)
			return
		}

		order, err := service.OrderServicer.CreateOrder(r.Context(), requestBody)
		if err != nil {
			log.WithError(err).WithField("body", requestBody).Errorln("Failed create new order")

			response := formatter.FailResponse(err.Error()).Stringify()
			formatter.WriteResponse(w, r, http.StatusInternalServerError, response)
			return
		}

		response := formatter.ObjectResponse(order).Stringify()
		formatter.WriteResponse(w, r, http.StatusCreated, response)
	}
}
