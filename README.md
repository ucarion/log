# log

[![go.dev reference](https://img.shields.io/badge/go.dev-reference-007d9c?logo=go&logoColor=white&style=flat-square)](https://pkg.go.dev/mod/github.com/ucarion/log?tab=overview)
[![GitHub Workflow Status](https://img.shields.io/github/workflow/status/ucarion/log/tests?label=tests&logo=github&style=flat-square)](https://github.com/ucarion/log/actions)

`github.com/ucarion/log` provides structured, canonical, `context`-based logs in
Golang on top of any existing logging backend. The core idea of canonical
logging is to output fewer, better logs -- say, one per HTTP request served.
With this package, you will add a bit more context to the log in many places:

```go
// Set a field to the current canonical log line -- this does not output anything.
log.Set("user_cache_latency_seconds", 5)
```

And then at the end of every request, you fire off a single log:

```go
// This will write out all of the fields you've added using Set, along with a message.
log.Log(ctx, "http_request")
```

You'll find this approach will give you far more useful logs. For instance, you
can easily load your logs into an analytics database, and then perform queries
to better understand the performance of your application.

## Installation

To install this package, run:

```bash
go get github.com/ucarion/log
```

## Usage

First, you need to choose what backend you want to use. See [the Backends
section in this README for more](#backends), but supposing you wanted to use
[Logrus](https://github.com/sirupsen/logrus):

```go
// All you have to do to enable Logrus as your logger is import:
import _ "github.com/ucarion/log/loggers/logrus"
```

In most of your code, using `github.com/ucarion/log` looks like this:

```go
// Prepare a context to keep track of the log-line properties.
ctx := log.NewContext(context.Background())

// Add some properties to the message.
log.Set(ctx, "start_time", time.Now())
log.Set(ctx, "rpc", "ListUsers")
log.Set(ctx, "err", err)

// Fire off a canonical log line.
log.Log(ctx, "request")
```

For a more advanced usage, you can also have nested properties by using
`WithPrefix`:

```go
ctx1 := log.NewContext(context.Background())
log.Set(ctx1, "root", "level")

ctx2 := log.WithPrefix(ctx1, "nested")
log.Set(ctx2, "foo", "bar")

// Outputs something like, depending your backend:
//
// { "msg": "example", "data": { "root": "level", "nested": { "foo": "bar" }}}}
log.Log(ctx, "example")
```

## Backends

`github.com/ucarion/log` is agnostic to what backend you want to use. But for
your convenience, it ships with support for a few popular backends.

### The standard library's `log`

To use `github.com/ucarion/log` with [the standard library `log`
package](https://golang.org/pkg/log/), just import:

```go
import _ "github.com/ucarion/log/loggers/log"
```

See [`examples/log`](./examples/log/main.go) for a working example you can play
with. See the GoDocs for details on how you can use other loggers than the
standard stdlib logger.

### sirupsen/logrus

To use `github.com/ucarion/log` with
[Logrus](https://github.com/sirupsen/logrus), just import:

```go
import _ "github.com/ucarion/log/loggers/logrus"
```

See [`examples/logrus`](./examples/logrus/main.go) for a working example you can
play with. See the GoDocs for details on how you can use other loggers than the
standard Logrus logger.

### uber/zap

To use `github.com/ucarion/log` with [Uber's
Zap](https://github.com/uber-go/zap), just import:

```go
import _ "github.com/ucarion/log/loggers/zap"
```

You'll also need to update the global logger, if you haven't done so already.
You can do this by running:

```go
// You could also use NewDevelopment or NewProduction, or a custom Zap logger.
zap.ReplaceGlobals(zap.NewExample())
```

See [`examples/zap`](./examples/zap/main.go) for a working example you can play
with. See the GoDocs for details on how you can use other loggers than the
global Zap logger.

### segmentio/events

To use `github.com/ucarion/log` with [Segment's
`events`](https://github.com/segmentio/events), just import:

```go
import _ "github.com/ucarion/log/loggers/events"
```

See [`examples/events`](./examples/events/main.go) for a working example you can
play with. See the GoDocs for details on how you can use other loggers than the
standard `events` logger.

## Implementing your own backend

Implementing your own custom backend for `github.com/ucarion/log` is just a
question of updating `log.Logger`, which just requires confirming to an
interface:

```go
var Logger interface {
	Log(msg string, fields map[string]interface{})
}
```

For example, here's some working code that uses `fmt.Println` as a logging
backend:

```go
type fmtLogger struct{}

func (l *fmtLogger) Log(msg string, fields map[string]interface{}) {
	fmt.Println("my custom logger", msg, fields)
}

func main() {
	log.Logger = &fmtLogger{}

	// ...
}
```

See [`examples/custom`](./examples/custom/main.go) for a working example you can
play with.
