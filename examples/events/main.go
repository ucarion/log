package main

import (
	"context"
	"time"

	_ "github.com/segmentio/events/v2/ecslogs"
	_ "github.com/segmentio/events/v2/text"
	"github.com/ucarion/log"
	_ "github.com/ucarion/log/loggers/events"
)

func main() {
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
