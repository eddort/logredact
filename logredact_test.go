package logredact_test

import (
	"io/ioutil"
	"testing"

	logredact "github.com/eddort/logrus-secret-remover-hook"
	"github.com/eddort/logrus-secret-remover-hook/memoryhook"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
)

type TestData struct {
	Secret string
	Public string
}

func TestSecretRemoverHook(t *testing.T) {
	t.Run("removes secret from log message", func(t *testing.T) {
		logger, memoryHook := setupLoggerWithMemoryHook([]string{"mysecret"})

		logger.Info("This is a message with a secret: mysecret")

		entry := getLastLogEntry(memoryHook)
		assert.Equal(t, "This is a message with a secret: ********", entry.Message)
	})

	t.Run("removes secret from log data", func(t *testing.T) {
		logger, memoryHook := setupLoggerWithMemoryHook([]string{"mysecret"})

		data := TestData{
			Secret: "mysecret",
			Public: "public",
		}

		logger.WithFields(logrus.Fields{"data": data}).Info("Logging a struct with secrets")

		entry := getLastLogEntry(memoryHook)
		assert.Equal(t, logrus.Fields{
			"data": TestData{
				Secret: "********",
				Public: "public",
			},
		}, entry.Data)
	})
}

func setupLoggerWithMemoryHook(secrets []string) (*logrus.Logger, *memoryhook.MemoryHook) {
	logger := logrus.New()
	logger.SetOutput(ioutil.Discard)

	memoryHook := memoryhook.New()
	secretHook := logredact.New(secrets, "********")

	logger.AddHook(memoryHook)
	logger.AddHook(secretHook)

	return logger, memoryHook
}

func getLastLogEntry(hook *memoryhook.MemoryHook) *logrus.Entry {
	entries := hook.Entries()
	return entries[len(entries)-1]
}
