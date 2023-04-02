# Logrus Secret Remover Hook
`logrus-secret-remover-hook` is a Logrus hook for removing sensitive information from log entries. This module can help you protect secrets in your logs by redacting them before the logs are written.

## Installation

To install the `logrus-secret-remover-hook`, run the following command:

```sh
$ go get github.com/eddort/go-secret-remover-hook
```

## Usage
To use the go-secret-remover-hook, import it in your Go code and add it to your Logrus logger. Here's an example:

```go
package main

import (
	"github.com/sirupsen/logrus"
	"github.com/eddort/logrus-secret-remover-hook"
)

func main() {
	logger := logrus.New()

	secrets := []string{"supersecret", "anothersecret"}
	secretHook := secretremoverhook.New(secrets, "***")

	logger.AddHook(secretHook)

	logger.Info("This log entry contains a supersecret token.")
}

```

In this example, we create a new Logrus logger, create a new secret remover hook with a list of secrets to redact, and add the hook to the logger. When we log an entry containing one of the secrets, the hook will replace the secret with `***`.

## Configuration

The logrus-secret-remover-hook takes an array of strings as first argument and `replacer` as the second argument. These strings represent the secrets you want to remove from your log entries. You can add as many secrets as needed:

```go
secrets := []string{"secret1", "secret2", "secret3"}
secretHook := secretremoverhook.New(secrets, "***")
```

## Contributing
If you'd like to contribute to the project, feel free to submit a pull request or open an issue.

## License
This module is released under the [MIT License](https://opensource.org/license/mit/).