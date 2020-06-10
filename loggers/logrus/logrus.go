// Package logrus is a backend to ucarion/log that uses Logrus as a backend.
//
// Importing this package automatically assigns to the global Logger, using the
// logrus StandardLogger and enabling the DetectErr option afforded by the
// Logger defined in this package.
package logrus

import (
	"github.com/sirupsen/logrus"
	"github.com/ucarion/log"
)

// Logger implements log.Logger with a Logrus logger backend.
//
// By default, all calls to Log will log at the "info" level, unless one of the
// top-level fields is an error, in which case the "error" level is used
// instead. To disable this error detection, and instead always log at the
// "info" level, set DetectErr to false.
type Logger struct {
	// The logrus Logger that is used to handle Log calls.
	Logger *logrus.Logger

	// Whether to detect errors in fields and set the log level accordingly.
	DetectErr bool
}

// Log implements Logger.
func (l *Logger) Log(msg string, fields map[string]interface{}) {
	if l.DetectErr {
		for _, v := range fields {
			if _, ok := v.(error); ok {
				l.Logger.WithFields(fields).Error(msg)
				return
			}
		}
	}

	l.Logger.WithFields(fields).Info(msg)
}

func init() {
	log.Logger = &Logger{Logger: logrus.StandardLogger(), DetectErr: true}
}
