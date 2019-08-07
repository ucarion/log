package log

import (
	upstream "log"

	"github.com/ucarion/log"
)

type Logger struct {
	Logger *upstream.Logger
}

func (l *Logger) Log(msg string, fields map[string]interface{}) {
	if l.Logger == nil {
		upstream.Println(msg, fields)
	} else {
		l.Logger.Println(msg, fields)
	}
}

func init() {
	log.DefaultLogger = &Logger{Logger: nil}
}
