<div align="center">
  <picture>
    <source media="(prefers-color-scheme: dark)" srcset=".github/profile/images/logo_landing_light.png">
    <source media="(prefers-color-scheme: light)" srcset=".github/profile/images/logo_landing_dark.png">
    <img alt="Sprout Logo" width="700" src="">
  </picture>
  <hr />
</div>

<div align="center">
<a target="_blank" href="https://github.com/42atomys/sprout/actions/workflows/test.yaml"><img src="https://img.shields.io/github/actions/workflow/status/42atomys/sprout/test.yaml?branch=main&label=tests"></a>
<a target="_blank" href="https://goreportcard.com/report/github.com/42atomys/sprout"><img src="https://goreportcard.com/badge/github.com/42atomys/sprout" /></a>
<a target="_blank" href="https://codeclimate.com/github/42atomys/sprout"><img alt="Code Climate maintainability" src="https://img.shields.io/codeclimate/maintainability/42atomys/sprout"></a>
<a target="_blank" href="https://codecov.io/gh/42atomys/sprout"><img alt="Codecov" src="https://img.shields.io/codecov/c/github/42atomys/sprout"></a>
<img src="https://img.shields.io/github/v/release/42atomys/sprout?label=last%20release" alt="GitHub release (latest by date)">
<img src="https://img.shields.io/github/contributors/42atomys/sprout?color=blueviolet" alt="GitHub contributors">
<img src="https://img.shields.io/github/stars/42atomys/sprout?style=flat&color=blueviolet" alt="GitHub Repo stars">
<a target="_blank" href="https://pkg.go.dev/github.com/42atomys/sprout"><img src="https://pkg.go.dev/badge/github.com/42atomys/sprout.svg" alt="Go Reference"></a>
<br />
<h3> <a target="_blank" href="https://docs.atom.codes/sprout">Official Documentation</a></h3>
<hr/>
</div>

> [!NOTE]
> Sprout is an evolved variant of the [Masterminds/sprig](https://github.com/Masterminds/sprig) library, reimagined for modern Go versions. It introduces fresh functionalities and commits to maintaining the library, picking up where Sprig left off. Notably, Sprig had not seen updates for two years and was not compatible beyond Golang 1.13, necessitating the creation of Sprout.

## Motivation

Sprout was born out of the need for a modernized, maintained, and performant template function library. Sprig, the predecessor to Sprout, had not seen updates for two years and was not optimized for later versions of Golang. Sprout aims to fill this gap by providing a library that is actively maintained, compatible with the latest Go versions, and optimized for performance.

## Roadmap to Sprout v1.0

You can track our progress towards Sprout v1.0 by following the documentation page
[here](https://docs.atom.codes/sprout/roadmap-to-sprout-v1.0).

## Table of Contents

- [Motivation](#motivation)
- [Roadmap to Sprout v1.0](#roadmap-to-sprout-v10)
- [Transitioning from Sprig](#transitioning-from-sprig)
- [Usage](#usage)
  - [Usage: Logger](#usage-logger)
  - [Usage: Alias](#usage-alias)
  - [Usage: Error Handling](#usage-error-handling)
    - [Default Value](#default-value)
    - [Panic](#panic)
    - [Error Channel](#error-channel)
- [Performence Benchmarks](#performence-benchmarks)
  - [Sprig v3.2.3 vs Sprout v0.2](#sprig-v323-vs-sprout-v02)
- [Development Philosophy (Currently in reflexion to create our)](#development-philosophy-currently-in-reflexion-to-create-our)


## Transitioning from Sprig

Sprout is designed to be a drop-in replacement for Sprig in the v1.0, with the same function names and behavior. To use Sprout in your project, simply replace the Sprig import with Sprout:

```diff
import (
-  "github.com/Masterminds/sprig/v3"
+  "github.com/42atomys/sprout"
)

tpl := template.Must(
  template.New("base").
-   Funcs(sprig.FuncMap()).
+   Funcs(sprout.FuncMap()).
    ParseGlob("*.tmpl")
)
```

## Usage

To use Sprout in your project, import the library and use the `FuncMap` function to add the template functions to your template:

```go
import (
  "github.com/42atomys/sprout"
  "text/template"
)

tpl := template.Must(
  template.New("base").
    Funcs(sprout.FuncMap()).
    ParseGlob("*.tmpl")
)
```

You can customize the behavior of Sprout by creating a `FunctionHandler` and passing it to the `FuncMap` function or using the configuration functions provided by Sprout:

```go
handler := sprout.NewFunctionHandler(
  // Add your logger to the handler to log errors and debug information using the
  // standard slog package or any other logger that implements the slog.Logger interface.
  // By default, Sprout uses a slog.TextHandler.
  sprout.WithLogger(slogLogger),
  // Set the error handling behavior for the handler. By default, Sprout returns the default value of the return type without crashes or panics.
  sprout.WithErrHandling(sprout.ErrHandlingReturnDefaultValue),
  // Set the error channel for the handler. By default, Sprout does not use an error channel. If you set an error channel, Sprout will send errors to it.
  // This options is only used when the error handling behavior is set to
  // `ErrHandlingErrorChannel`
  sprout.WithErrorChannel(errChan),
  // Set the alias for a function. By default, Sprout use alias for some functions for backward compatibility with Sprig.
  sprout.WithAlias("hello", "hi"),
)

// Use the handler with the FuncMap function. The handler will be used to handle all template functions.
tpl := template.Must(
  template.New("base").
    Funcs(sprout.FuncMap(sprout.WithFunctionHandler(handler))).
    ParseGlob("*.tmpl")
)
```
### Usage: Logger

Sprout uses the `slog` package for logging. You can pass your logger
to the `WithLogger` configuration function to log errors and debug information:

```go
// Create a new logger using the slog package.
logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))

// Use the handler with the FuncMap function.
tpl := template.Must(
  template.New("base").
    Funcs(sprout.FuncMap(sprout.WithLogger(logger))).
    ParseGlob("*.tmpl")
)
```

### Usage: Alias

Sprout provides the ability to set an alias for a function. This feature is useful for backward compatibility with Sprig. You can set an alias for a function using the `WithAlias` or `WithAliases` configuration functions.

See more about the alias in the [documentation](https://docs.atom.codes/sprout/function-aliases).

```go
sprout.NewFunctionHandler(
  sprout.WithAlias("hello", "hi"),
)
```

### Usage: Error Handling

Sprout provides three error handling behaviors:
- `ErrHandlingReturnDefaultValue`: Sprout returns the default value of the return type without crashes or panics.
- `ErrHandlingPanic`: Sprout panics when an error occurs.
- `ErrHandlingErrorChannel`: Sprout sends errors to the error channel.

You can set the error handling behavior using the `WithErrHandling` configuration function:

```go
sprout.NewFunctionHandler(
  sprout.WithErrHandling(sprout.ErrHandlingReturnDefaultValue),
)
```

#### Default Value

If you set the error handling behavior to `ErrHandlingReturnDefaultValue`, Sprout will return the default value of the return type without crashes or panics to ensure a smooth user experience when an error occurs.

#### Panic

If you set the error handling behavior to `ErrHandlingPanic`, Sprout will panic when an error occurs to ensure that the error is not ignored and sended back to template execution.

#### Error Channel

If you set the error handling behavior to `ErrHandlingErrorChannel`, you can pass an error channel to the `WithErrorChannel` configuration function. Sprout will send errors to the error channel:

```go
errChan := make(chan error)

sprout.NewFunctionHandler(
  sprout.WithErrHandling(sprout.ErrHandlingErrorChannel),
  sprout.WithErrorChannel(errChan),
)
```

## Performence Benchmarks

To see all the benchmarks, please refer to the [benchmarks](benchmarks/README.md) directory.

### Sprig v3.2.3 vs Sprout v0.2
```
goos: linux
goarch: amd64
pkg: sprout_benchmarks
cpu: Intel(R) Core(TM) i7-9750H CPU @ 2.60GHz
BenchmarkSprig-12              1        3869134593 ns/op        45438616 B/op      24098 allocs/op
BenchmarkSprout-12             1        1814126036 ns/op        38284040 B/op      11627 allocs/op
PASS
ok      sprout_benchmarks       5.910s
```

**Time improvement**: 53.1%
**Memory improvement**: 15.7%

So, Sprout v0.2 is approximately 53.1% faster and uses 15.7% less memory than Sprig v3.2.3. 🚀

## Development Philosophy (Currently in reflexion to create our)

Our approach to extending and refining Sprout was guided by several key principles:

- Empowering layout construction through template functions.
- Designing template functions that avoid returning errors when possible, instead displaying default values for smoother user experiences.
- Ensuring template functions operate solely on provided data, without external data fetching.
- Maintaining the integrity of core Go template functionalities without overrides.






