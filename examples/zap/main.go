package main

import (
	"context"
	"time"

	"github.com/ucarion/log"
	_ "github.com/ucarion/log/loggers/zap"
	"go.uber.org/zap"
)

func main() {
	zap.ReplaceGlobals(zap.NewExample())

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
