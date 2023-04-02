package logredact_test

import (
	"bytes"
	"fmt"
	"testing"

	"github.com/eddort/logredact"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
)

func TestSecretRemoverHook(t *testing.T) {
	tests := []struct {
		name          string
		input         interface{}
		expected      string
		shouldContain bool
	}{
		{
			name:          "string with secret",
			input:         "my-password-is-hunter2",
			expected:      "hunter2",
			shouldContain: false,
		},
		{
			name:          "string without secret",
			input:         "just-an-ordinary-string",
			expected:      "just-an-ordinary-string",
			shouldContain: true,
		},
		{
			name:          "integer",
			input:         42,
			expected:      "42",
			shouldContain: true,
		},
		{
			name:          "float",
			input:         3.14,
			expected:      "3.14",
			shouldContain: true,
		},
		{
			name:          "array",
			input:         []string{"first-string", "password123", "third-string"},
			expected:      "password123",
			shouldContain: false,
		},
		{
			name:          "map",
			input:         map[string]string{"key1": "value1", "key2": "mysecret123"},
			expected:      "mysecret123",
			shouldContain: false,
		},
		{
			name:          "struct",
			input:         struct{ Field1, Field2 string }{"data1", "hiddensecret"},
			expected:      "hiddensecret",
			shouldContain: false,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			buf := new(bytes.Buffer)
			logger := logrus.New()
			logger.SetOutput(buf)
			logger.AddHook(logredact.New([]string{"hunter2", "password123", "mysecret123", "hiddensecret"}, "***"))

			logger.WithField("input", tc.input).Info("Test log entry")

			if tc.shouldContain {
				assert.Contains(t, buf.String(), tc.expected, fmt.Sprintf("Log entry should contain '%s'", tc.expected))
			} else {
				assert.NotContains(t, buf.String(), tc.expected, fmt.Sprintf("Log entry should not contain '%s'", tc.expected))
			}
		})
	}
}
