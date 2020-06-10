// Package log is a backend to ucarion/log that uses the standard library's log
// as a backend.
//
// Importing this package automatically assigns to the global Logger, using the
// standard logger as the underlying stdlib log logger.
package log

import (
	upstream "log"

	"github.com/ucarion/log"
)

// Logger implements log.Logger with a stdlib logger backend.
type Logger struct {
	// The stdlib Logger that is used to handle Log calls. If nil, the standard
	// logger is used instead.
	Logger *upstream.Logger
}

// Log implements Logger.
func (l *Logger) Log(msg string, fields map[string]interface{}) {
	if l.Logger == nil {
		upstream.Println(msg, fields)
	} else {
		l.Logger.Println(msg, fields)
	}
}

func init() {
	log.Logger = &Logger{}
}
