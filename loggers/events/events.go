// Package events is a backend for log that uses segmentio/events.
//
// Importing this package automatically assigns to the global Logger, using
// DefaultLogger as the underlying events logger.
package events

import (
	"github.com/segmentio/events/v2"
	"github.com/ucarion/log"
)

// Logger implements log.Logger with a segmentio/events backend.
type Logger struct {
	// The events Logger that is used to handle Log calls.
	Logger *events.Logger
}

// Log implements Logger.
func (l *Logger) Log(msg string, fields map[string]interface{}) {
	l.Logger.Log(msg, events.A(fields))
}

func init() {
	log.Logger = &Logger{Logger: events.DefaultLogger}
}
