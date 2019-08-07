package events

import (
	"github.com/segmentio/events"
	"github.com/ucarion/log"
)

type Logger struct {
	Logger *events.Logger
}

func (l *Logger) Log(msg string, fields map[string]interface{}) {
	l.Logger.Log(msg, events.A(fields))
}

func init() {
	log.DefaultLogger = &Logger{Logger: events.DefaultLogger}
}
