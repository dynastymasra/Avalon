package web

import (
	"fmt"
	"net/http"

	"github.com/dynastymasra/avalon/infrastructure/web/controller/order"

	"github.com/dynastymasra/avalon/service"

	"github.com/dynastymasra/avalon/infrastructure/web/middleware"

	"github.com/dynastymasra/avalon/infrastructure/provider"

	"github.com/dynastymasra/avalon/infrastructure/web/controller"

	"github.com/dynastymasra/avalon/config"

	"github.com/dynastymasra/avalon/infrastructure/web/formatter"
	"github.com/gorilla/mux"
	"github.com/urfave/negroni"
)

func Router(provider *provider.Instance, service service.Instance) *mux.Router {
	router := mux.NewRouter().StrictSlash(true).UseEncodedPath()
	commonHandlers := negroni.New(
		middleware.HTTPStatLogger(),
		middleware.RequestID(),
	)

	subRouter := router.PathPrefix("/v1/").Subrouter().UseEncodedPath()

	subRouter.NotFoundHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, formatter.FailResponse(config.ErrDataNotFound.Error()).Stringify())
	})

	// Probes
	subRouter.Handle("/ping", commonHandlers.With(
		negroni.WrapFunc(controller.Ping(provider)),
	)).Methods(http.MethodGet, http.MethodHead)

	// Order group
	subRouter.Handle("/orders", commonHandlers.With(
		negroni.WrapFunc(order.Save(service)),
	)).Methods(http.MethodPost)

	return router
}
