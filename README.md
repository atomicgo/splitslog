<!--



┌───────────────────────────────────────────────────────────────────┐
│                                                                   │
│                          IMPORTANT NOTE                           │
│                                                                   │
│               This file is automatically generated                │
│           All manual modifications will be overwritten            │
│                                                                   │
└───────────────────────────────────────────────────────────────────┘



-->

<h1 align="center">AtomicGo | splitslog</h1>

<p align="center">
<img src="https://img.shields.io/endpoint?url=https%3A%2F%2Fatomicgo.dev%2Fapi%2Fshields%2Fsplitslog&style=flat-square" alt="Downloads">

<a href="https://github.com/atomicgo/splitslog/releases">
<img src="https://img.shields.io/github/v/release/atomicgo/splitslog?style=flat-square" alt="Latest Release">
</a>

<a href="https://codecov.io/gh/atomicgo/splitslog" target="_blank">
<img src="https://img.shields.io/github/actions/workflow/status/atomicgo/splitslog/go.yml?style=flat-square" alt="Tests">
</a>

<a href="https://codecov.io/gh/atomicgo/splitslog" target="_blank">
<img src="https://img.shields.io/codecov/c/gh/atomicgo/splitslog?color=magenta&logo=codecov&style=flat-square" alt="Coverage">
</a>

<a href="https://codecov.io/gh/atomicgo/splitslog">
<!-- unittestcount:start --><img src="https://img.shields.io/badge/Unit_Tests-1-magenta?style=flat-square" alt="Unit test count"><!-- unittestcount:end -->
</a>

<a href="https://opensource.org/licenses/MIT" target="_blank">
<img src="https://img.shields.io/badge/License-MIT-yellow.svg?style=flat-square" alt="License: MIT">
</a>
  
<a href="https://goreportcard.com/report/github.com/atomicgo/splitslog" target="_blank">
<img src="https://goreportcard.com/badge/github.com/atomicgo/splitslog?style=flat-square" alt="Go report">
</a>   

</p>

---

<p align="center">
<strong><a href="https://pkg.go.dev/atomicgo.dev/splitslog#section-documentation" target="_blank">Documentation</a></strong>
|
<strong><a href="https://github.com/atomicgo/atomicgo/blob/main/CONTRIBUTING.md" target="_blank">Contributing</a></strong>
|
<strong><a href="https://github.com/atomicgo/atomicgo/blob/main/CODE_OF_CONDUCT.md" target="_blank">Code of Conduct</a></strong>
</p>

---

<p align="center">
  <img src="https://raw.githubusercontent.com/atomicgo/atomicgo/main/assets/header.png" alt="AtomicGo">
</p>

<p align="center">
<table>
<tbody>
</tbody>
</table>
</p>
<h3  align="center"><pre>go get atomicgo.dev/splitslog</pre></h3>
<p align="center">
<table>
<tbody>
</tbody>
</table>
</p>

<!-- gomarkdoc:embed:start -->

<!-- Code generated by gomarkdoc. DO NOT EDIT -->

# splitslog

```go
import "atomicgo.dev/splitslog"
```

Package splitslog provides a handler that splits log records to different handlers based on their level.

The most common use case is to split logs to stdout and stderr based on their level.





```go
package main

import (
	"log/slog"
	"os"

	"atomicgo.dev/splitslog"
)

func main() {
	splitter := splitslog.Splitter{
		// Debug and info messages are printed to stdout.
		slog.LevelDebug: slog.NewJSONHandler(os.Stdout, nil),
		slog.LevelInfo:  slog.NewJSONHandler(os.Stdout, nil),

		// Warn and error messages are printed to stderr.
		slog.LevelWarn:  slog.NewJSONHandler(os.Stderr, nil),
		slog.LevelError: slog.NewJSONHandler(os.Stderr, nil),
	}

	handler := splitslog.NewSplitHandler(splitter)
	logger := slog.New(handler)

	logger.Info("info message prints to stdout")
	logger.Error("error message prints to stderr")

	// stdout: {"time":"2023-09-07T16:56:22.563817+02:00","level":"INFO","msg":"info message prints to stdout"}
	// stderr: {"time":"2023-09-07T16:56:22.564103+02:00","level":"ERROR","msg":"error message prints to stderr"}
}
```



## Index

- [type SplitHandler](<#SplitHandler>)
  - [func NewSplitHandler\(splitter Splitter\) \*SplitHandler](<#NewSplitHandler>)
  - [func \(h \*SplitHandler\) Enabled\(ctx context.Context, level slog.Level\) bool](<#SplitHandler.Enabled>)
  - [func \(h \*SplitHandler\) Handle\(ctx context.Context, record slog.Record\) error](<#SplitHandler.Handle>)
  - [func \(h \*SplitHandler\) WithAttrs\(attrs \[\]slog.Attr\) slog.Handler](<#SplitHandler.WithAttrs>)
  - [func \(h \*SplitHandler\) WithGroup\(name string\) slog.Handler](<#SplitHandler.WithGroup>)
- [type Splitter](<#Splitter>)


<a name="SplitHandler"></a>
## type [SplitHandler](<https://github.com/atomicgo/splitslog/blob/main/splitslog.go#L15-L19>)

SplitHandler is a handler that splits log records to different handlers based on their level.

```go
type SplitHandler struct {
    Splitter Splitter
    // contains filtered or unexported fields
}
```

<a name="NewSplitHandler"></a>
### func [NewSplitHandler](<https://github.com/atomicgo/splitslog/blob/main/splitslog.go#L22>)

```go
func NewSplitHandler(splitter Splitter) *SplitHandler
```

NewSplitHandler returns a new SplitHandler.





```go
package main

import (
	"log/slog"
	"os"

	"atomicgo.dev/splitslog"
)

func main() {
	splitter := splitslog.Splitter{
		// Debug and info messages are printed to stdout.
		slog.LevelDebug: slog.NewJSONHandler(os.Stdout, nil),
		slog.LevelInfo:  slog.NewJSONHandler(os.Stdout, nil),

		// Warn and error messages are printed to stderr.
		slog.LevelWarn:  slog.NewJSONHandler(os.Stderr, nil),
		slog.LevelError: slog.NewJSONHandler(os.Stderr, nil),
	}

	handler := splitslog.NewSplitHandler(splitter)
	logger := slog.New(handler)

	logger.Info("info message prints to stdout")
	logger.Error("error message prints to stderr")

	// stdout: {"time":"2023-09-07T16:56:22.563817+02:00","level":"INFO","msg":"info message prints to stdout"}
	// stderr: {"time":"2023-09-07T16:56:22.564103+02:00","level":"ERROR","msg":"error message prints to stderr"}
}
```



<a name="SplitHandler.Enabled"></a>
### func \(\*SplitHandler\) [Enabled](<https://github.com/atomicgo/splitslog/blob/main/splitslog.go#L40>)

```go
func (h *SplitHandler) Enabled(ctx context.Context, level slog.Level) bool
```

Enabled implements Handler.Enabled.

<a name="SplitHandler.Handle"></a>
### func \(\*SplitHandler\) [Handle](<https://github.com/atomicgo/splitslog/blob/main/splitslog.go#L45>)

```go
func (h *SplitHandler) Handle(ctx context.Context, record slog.Record) error
```

Handle implements Handler.Handle.

<a name="SplitHandler.WithAttrs"></a>
### func \(\*SplitHandler\) [WithAttrs](<https://github.com/atomicgo/splitslog/blob/main/splitslog.go#L62>)

```go
func (h *SplitHandler) WithAttrs(attrs []slog.Attr) slog.Handler
```

WithAttrs implements Handler.WithAttrs.

<a name="SplitHandler.WithGroup"></a>
### func \(\*SplitHandler\) [WithGroup](<https://github.com/atomicgo/splitslog/blob/main/splitslog.go#L71>)

```go
func (h *SplitHandler) WithGroup(name string) slog.Handler
```

WithGroup implements Handler.WithGroup.

<a name="Splitter"></a>
## type [Splitter](<https://github.com/atomicgo/splitslog/blob/main/splitslog.go#L12>)

Splitter is a map of log levels to handlers. The default log levels \(slog.LevelDebug, slog.LevelInfo, slog.LevelWarn, slog.LevelError\) must be present, otherwise the SplitHandler panics.

```go
type Splitter map[slog.Level]slog.Handler
```

Generated by [gomarkdoc](<https://github.com/princjef/gomarkdoc>)


<!-- gomarkdoc:embed:end -->

---

> [AtomicGo.dev](https://atomicgo.dev) &nbsp;&middot;&nbsp;
> with ❤️ by [@MarvinJWendt](https://github.com/MarvinJWendt) |
> [MarvinJWendt.com](https://marvinjwendt.com)
