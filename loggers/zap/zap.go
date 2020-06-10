// Package zap is a backend to ucarion/log that uses Zap as a backend.
//
// Importing this package automatically assigns to the global Logger, using the
// zap global logger and enabling the DetectErr option afforded by the Logger
// defined in this package.
package zap

import (
	"github.com/ucarion/log"
	"go.uber.org/zap"
)

// Logger implements log.Logger with a Logrus logger backend.
//
// By default, all calls to Log will log at the "info" level, unless one of the
// top-level fields is an error, in which case the "error" level is used
// instead. To disable this error detection, and instead always log at the
// "info" level, set DetectErr to false.
type Logger struct {
	// The zap Logger that is used to handle Log calls.
	Logger *zap.Logger

	// Whether to detect errors in fields and set the log level accordingly.
	DetectErr bool
}

// Log implements Logger.
func (l *Logger) Log(msg string, fields map[string]interface{}) {
	logger := l.Logger
	if logger == nil {
		logger = zap.L()
	}

	if l.DetectErr {
		for _, v := range fields {
			if _, ok := v.(error); ok {
				logger.Error(msg, zap.Reflect("fields", fields))
				return
			}
		}
	}

	logger.Info(msg, zap.Reflect("fields", fields))
}

func init() {
	log.Logger = &Logger{Logger: nil, DetectErr: true}
}
