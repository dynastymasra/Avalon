package middleware

import (
	"net/http"
	"strings"
	"time"

	"github.com/sirupsen/logrus"
	"github.com/urfave/negroni"
)

func HTTPStatLogger() negroni.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
		startTime := time.Now()

		next(w, r)

		responseTime := time.Now()
		deltaTime := responseTime.Sub(startTime)

		if !IsHealthCheckURL(r.URL.String()) {
			logrus.WithFields(logrus.Fields{
				"request_time":   startTime.Format(time.RFC3339),
				"delta_time":     deltaTime,
				"response_time":  responseTime.Format(time.RFC3339),
				"request_proxy":  r.RemoteAddr,
				"url":            r.URL.Path,
				"method":         r.Method,
				"request_source": r.Header.Get("X-FORWARDED-FOR"),
				"headers":        r.Header,
			}).Infoln("HTTP Request")
		}
	}
}

func IsHealthCheckURL(url string) bool {
	switch {
	case strings.Contains(url, "/ping"):
		return true
	default:
		return false
	}
}
