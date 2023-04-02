package main

import (
	logredact "github.com/eddort/logrus-secret-remover-hook"
	"github.com/sirupsen/logrus"
)

func main() {
	logger := logrus.New()

	secrets := []string{"supersecret", "anothersecret"}
	secretHook := logredact.New(secrets, "***")

	logger.AddHook(secretHook)

	logger.Info("This log entry contains a supersecret token.")
}
