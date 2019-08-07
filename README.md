# log

`log` helps you implement **canonical log lines** in Golang.

See [this blog post from Stripe][stripe] for more about what canonical log
lines, and how you can use them. But the basic idea is to log few,
highly-detailed messages that you can then ETL into your observability tools.

[stripe]: https://stripe.com/blog/canonical-log-lines

`log` helps you implement canonical log lines by exposing a trivial interface:

```go
// Prepare a context.
ctx := log.NewContext(r.Context())

// Add some properties to the canonical log line.
log.Set(ctx, "start_time", time.Now())
log.Set(ctx, "rpc", "ListUsers")
log.Set(ctx, "err", err)

// Fire off a canonical log line.
log.Log("request")
```

This package is not a one-stop logging solution. All you're getting here is:

1. A simple solution for keeping log fields in `ctx`, and logging them out.
2. A set of integrations with existing logging backends.

The guts of this package is basically twenty lines of Golang. So if you like the
idea, but have incompatible requirements, consider just copy-pasting this
package!

Because of how simple the interface is, it's easy to add your own middleware as
required. This package is meant to be a cog in your wider observability
implementation.

## Backends

Today, this package supports the following logging backends:

* `log`
* `github.com/sirupsen/logrus`
* `go.uber.org/zap`
* `github.com/segmentio/events`

Adding more is easy. This is the interface to satisfy:

```go
type Logger interface {
  Log(msg string, fields map[string]interface{})
}
```
