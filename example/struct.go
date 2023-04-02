package main

import (
	logredact "github.com/eddort/logrus-secret-remover-hook"
	"github.com/sirupsen/logrus"
)

type MyStruct struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func LogStruct() {
	logger := logrus.New()
	logger.SetFormatter(&logrus.JSONFormatter{})

	secrets := []string{"mysecret1", "mysecret2"}

	hook := logredact.New(secrets, "***")

	logger.AddHook(hook)

	logger.Info("This is a log entry with mysecret1 and mysecret2")

	myData := MyStruct{
		Username: "myuser",
		Password: "mysecret1",
	}

	// Logging an object with sensitive data using WithFields
	logger.WithFields(logrus.Fields{
		"my_struct": myData,
	}).Info("Logging a Go struct with sensitive data")
}
