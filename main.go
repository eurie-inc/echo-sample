package main

import (
	"github.com/eurie-inc/echo-sample/route"
	"github.com/Sirupsen/logrus"
)

func init() {

	logrus.SetLevel(logrus.DebugLevel)
	logrus.SetFormatter(&logrus.JSONFormatter{})
}

func main() {

	router := route.Init()
	router.Run(":8080")
}
