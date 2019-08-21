package config

import (
	"io/ioutil"
	"os"
	"strings"

	"github.com/sirupsen/logrus"
)

func SetupLogger() {
	level, err := logrus.ParseLevel(LogLevel())
	if err != nil {
		logrus.Fatalln("Unable to parse log level", err.Error())
	}

	logrus.SetOutput(os.Stderr)
	logrus.SetLevel(level)

	switch strings.ToLower(LoggerFormat()) {
	case "json":
		logrus.SetFormatter(&logrus.JSONFormatter{})
	default:
		logrus.SetFormatter(&logrus.TextFormatter{})
	}
}

func SetupTestLogger() {
	logrus.SetOutput(ioutil.Discard)
}
