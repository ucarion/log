package log

import (
	"context"
	"sync"
)

type Logger interface {
	Log(msg string, fields map[string]interface{})
}

var DefaultLogger Logger

type ctxKey struct{}

type ctxValue struct {
	fields map[string]interface{}
	*sync.Mutex
}

func NewContext(ctx context.Context) context.Context {
	return context.WithValue(ctx, ctxKey{}, ctxValue{
		map[string]interface{}{},
		&sync.Mutex{},
	})
}

func Set(ctx context.Context, key string, value interface{}) {
	ctxVal := ctx.Value(ctxKey{}).(ctxValue)
	ctxVal.Lock()
	ctxVal.fields[key] = value
	ctxVal.Unlock()
}

func Log(ctx context.Context, msg string) {
	ctxVal := ctx.Value(ctxKey{}).(ctxValue)
	ctxVal.Lock()
	DefaultLogger.Log(msg, ctxVal.fields)
	ctxVal.fields = map[string]interface{}{}
	ctxVal.Unlock()
}
