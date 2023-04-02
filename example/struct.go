package main

import (
	"github.com/eddort/logredact"
	"github.com/sirupsen/logrus"
)

type MyStruct struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// rename it to main
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

	logger.WithFields(logrus.Fields{
		"my_struct": myData,
	}).Info("Logging a Go struct with sensitive data")
}
