// Package log provides structured, canonical, Context-based logging.
package log

import (
	"context"
	"sync"
)

// Logger is the interface that gets called when you call Log.
var Logger interface {
	Log(msg string, fields map[string]interface{})
}

type ctxKey struct{}

type ctxValue struct {
	Path   []string
	Fields *map[string]interface{}
	*sync.Mutex
}

// NewContext constructs a new context with an empty set of log fields. Calls to
// Set with this context will add fields at the root level.
func NewContext(ctx context.Context) context.Context {
	return context.WithValue(ctx, ctxKey{}, ctxValue{
		Fields: &map[string]interface{}{},
		Mutex:  &sync.Mutex{},
	})
}

// WithPrefix constructs a new context where calls to Set will write to a
// sub-object at the given prefix.
//
// If you never call Set with the returned context, then no sub-object will be
// created in the logger's fields.
//
// WithPrefix will panic if passed a context that wasn't created through
// NewContext.
func WithPrefix(ctx context.Context, prefix string) context.Context {
	val := ctx.Value(ctxKey{}).(ctxValue)
	return context.WithValue(ctx, ctxKey{}, ctxValue{
		Path:   append(val.Path, prefix),
		Fields: val.Fields,
		Mutex:  val.Mutex,
	})
}

// Set assigns the field at key with value.
//
// If an entry with the given key already exists, then it is overwritten. Use
// WithPrefix to write to non-root-level field properties.
//
// Set will panic if passed a context that wasn't created through NewContext.
func Set(ctx context.Context, key string, value interface{}) {
	val := ctx.Value(ctxKey{}).(ctxValue)
	val.Lock()
	defer val.Unlock()

	fields := *val.Fields
	for _, s := range val.Path {
		if _, ok := fields[s]; !ok {
			fields[s] = map[string]interface{}{}
		}

		fields = fields[s].(map[string]interface{})
	}

	fields[key] = value
}

// Log passes the given message, and well as the fields accrued through Set
// calls, to the global Logger.
//
// Log does not remove existing fields on the context. It is valid to call Log
// on the same context multiple times.
//
// Log will panic if passed a context that wasn't created through NewContext.
func Log(ctx context.Context, msg string) {
	val := ctx.Value(ctxKey{}).(ctxValue)
	val.Lock()
	defer val.Unlock()

	Logger.Log(msg, *val.Fields)
}
