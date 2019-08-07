package zap

import (
	"github.com/ucarion/log"
	"go.uber.org/zap"
)

type Logger struct {
	Logger    *zap.Logger
	DetectErr bool
}

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
	log.DefaultLogger = &Logger{Logger: nil, DetectErr: true}
}
