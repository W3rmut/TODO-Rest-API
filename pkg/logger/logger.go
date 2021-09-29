package logger

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"restapi/pkg/config"
)

var Logger *logrus.Logger

func InitLogger() {
	Logger = logrus.New()
	level, err := logrus.ParseLevel(config.ResultConfig.LogsConfig.LogLevel)
	if err != nil {
		fmt.Errorf("Error parsing config file gor logrus")
	}
	Logger.SetLevel(level)

}
