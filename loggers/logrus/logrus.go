package logrus

import (
	upstream "github.com/sirupsen/logrus"
	"github.com/ucarion/log"
)

type Logger struct {
	Logger    *upstream.Logger
	DetectErr bool
}

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
	log.DefaultLogger = &Logger{Logger: upstream.StandardLogger(), DetectErr: true}
}
