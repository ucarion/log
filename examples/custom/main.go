package main

import (
	"context"
	"fmt"
	"time"

	"github.com/ucarion/log"
)

type fmtLogger struct{}

func (l *fmtLogger) Log(msg string, fields map[string]interface{}) {
	fmt.Println("my custom logger", msg, fields)
}

func main() {
	log.Logger = &fmtLogger{}

	ctx := log.NewContext(context.Background())
	log.Set(ctx, "start_time", time.Now())
	log.Set(ctx, "rpc", "ListUsers")

	foo(log.WithPrefix(ctx, "first_invocation"), "xxx")
	foo(log.WithPrefix(ctx, "second_invocation"), "yyy")

	log.Log(ctx, "example")
}

func foo(ctx context.Context, s string) {
	log.Set(ctx, "foo", s)
}
