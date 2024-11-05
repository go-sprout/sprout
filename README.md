<div align="center">
  <picture>
    <source media="(prefers-color-scheme: dark)" srcset=".github/profile/images/logo_landing_light.png">
    <source media="(prefers-color-scheme: light)" srcset=".github/profile/images/logo_landing_dark.png">
    <img alt="Sprout Logo" width="700" src="">
  </picture>
  <hr />
</div>

<div align="center">
<a target="_blank" href="https://github.com/go-sprout/sprout/actions/workflows/test.yaml"><img src="https://img.shields.io/github/actions/workflow/status/go-sprout/sprout/test.yaml?branch=main&label=tests"></a>
<a target="_blank" href="https://goreportcard.com/report/github.com/go-sprout/sprout"><img src="https://goreportcard.com/badge/github.com/go-sprout/sprout" /></a>
<a target="_blank" href="https://codeclimate.com/github/go-sprout/sprout"><img alt="Code Climate maintainability" src="https://img.shields.io/codeclimate/maintainability/go-sprout/sprout"></a>
<a target="_blank" href="https://codecov.io/gh/go-sprout/sprout"><img alt="Codecov" src="https://img.shields.io/codecov/c/github/go-sprout/sprout"></a>
<img src="https://img.shields.io/github/v/release/go-sprout/sprout?label=last%20release" alt="GitHub release (latest by date)">
<img src="https://img.shields.io/github/contributors/go-sprout/sprout?color=blueviolet" alt="GitHub contributors">
<img src="https://img.shields.io/github/stars/go-sprout/sprout?style=flat&color=blueviolet" alt="GitHub Repo stars">
<a target="_blank" href="https://pkg.go.dev/github.com/go-sprout/sprout"><img src="https://pkg.go.dev/badge/github.com/go-sprout/sprout.svg" alt="Go Reference"></a>
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
  - [Creating a Handler](#creating-a-handler)
  - [Customizing the Handler](#customizing-the-handler)
  - [Working with Registries](#working-with-registries)
  - [Working with Registries Groups](#working-with-registries-groups)
  - [Building Function Maps](#building-function-maps)
  - [Working with Templates](#working-with-templates)
- [Usage: Quick Example (code only)](#usage-quick-example)
- [Performance Benchmarks](#performance-benchmarks)
  - [Sprig v3.2.3 vs Sprout v0.2](#sprig-v323-vs-sprout-v02)
- [Development Philosophy (Currently in reflexion to create our)](#development-philosophy-currently-in-reflexion-to-create-our)


## Transitioning from Sprig

Sprout provide a package `sprigin` to provide a drop-in replacement for Sprig in the v1.0, with the same function names and behavior. To use Sprout in your project, simply replace the Sprig import with sprigin:

> [!IMPORTANT]
> The `sprigin` package is a temporary solution to provide backward compatibility with Sprig. We recommend updating your code to use the Sprout package directly to take advantage of the new features and improvements.
>
> A complete guide is available in the [documentation](https://docs.atom.codes/sprout/migration-from-sprig).

```diff
import (
-  "github.com/Masterminds/sprig/v3"
+  "github.com/go-sprout/sprout/sprigin"
)

tpl := template.Must(
  template.New("base").
-   Funcs(sprig.FuncMap()).
+   Funcs(sprigin.FuncMap()).
    ParseGlob("*.tmpl")
)
```

## Usage

### Creating a Handler
A handler in Sprout is responsible for managing the function registries and functions. The DefaultHandler is the primary implementation provided by Sprout.

```go
import "github.com/go-sprout/sprout"

handler := sprout.New()
``` 

### Customizing the Handler

Sprout supports various customization options using handler options:

```go
handler := sprout.New(
  // Add your logger to the handler to log errors and debug information using the
  // standard slog package or any other logger that implements the slog.Logger interface.
  // By default, Sprout uses a slog.TextHandler.
  sprout.WithLogger(slogLogger),
  // Set the alias for a function. By default, Sprout use alias for some functions for backward compatibility with Sprig.
  sprout.WithAlias("hello", "hi"),
)
```

### Working with Registries
Registries in Sprout are groups of functions that can be added to a handler. They help organize functions and optimize template performance.

You can retrieve all built-in registries and functions under [Registries](https://docs.atom.codes/sprout/registries/list-of-all-registries).

```go
import (
  "github.com/go-sprout/sprout/registry/conversion" // toString, toInt, toBool, ...
  "github.com/go-sprout/sprout/registry/std" // default, empty, any, all, ...
)

//...

handler.AddRegistries(
  conversion.NewRegistry(),
  std.NewRegistry(),
)
```

### Working with Registries Groups
In some cases, you can use a group of registries to add multiple registries at once.

You can retrieve all built-in registries groups under [Registry Groups](https://docs.atom.codes/sprout/groups/list-of-all-registry-groups).

```go
import (
  "github.com/go-sprout/sprout/group/all"
)

//...

handler.AddGroup(
  all.RegistryGroup(),
)
```

### Building Function Maps

To use Sprout with templating engines like `html/template` or `text/template`, you need to build the function map:
```go
funcs := handler.Build()
tpl := template.New("example").Funcs(funcs).Parse(`{{ hello }}`)
```

### Working with Templates
Once your function map is ready, you can use it to render templates:
```go
tpl, err := template.New("example").Funcs(funcs).Parse(`{{ myFunc }}`)
if err != nil {
    log.Fatal(err)
}
tpl.Execute(os.Stdout, nil)
``` 
This will render the template with all functions and aliases available.


## Usage: Quick Example 

Here is a quick example of how to use Sprout with the `text/template` package:
```go
package main

import (
	"os"
	"text/template"

	"github.com/go-sprout/sprout"
	"github.com/go-sprout/sprout/registry/std"
)

func main() {
	handler := sprout.New()
	handler.AddRegistry(std.NewRegistry())

	tpl := template.Must(
    template.New("example").Funcs(handler.Build()).Parse(`{{ hello }}`),
  )
	tpl.Execute(os.Stdout, nil)
}
```

## Performance Benchmarks

To see all the benchmarks, please refer to the [benchmarks](benchmarks/README.md) directory.

## Sprig v3.2.3 vs Sprout v0.5
```
goos: linux
goarch: amd64
pkg: sprout_benchmarks
cpu: Intel(R) Core(TM) i9-9900K CPU @ 3.60GHz
BenchmarkSprig-16              1        2991811373 ns/op        50522680 B/op      32649 allocs/op
BenchmarkSprout-16             1        1638797544 ns/op        42171152 B/op      18061 allocs/op
PASS
ok      sprout_benchmarks       4.921s
```

**Time improvement**: ((2991811373 - 1638797544) / 2991811373) * 100 = 45.3%
**Memory improvement**: ((50522680 - 42171152) / 50522680) * 100 = 16.5%

So, Sprout v0.5 is approximately 45.3% faster and uses 16.5% less memory than Sprig v3.2.3. 🚀

You can see the full benchmark results [here](benchmarks/README.md).

## Development Philosophy (Currently in consideration to create ours)

Our approach to extending and refining Sprout was guided by several key principles:

- Build on the principles of simplicity, flexibility, and consistency. 
- Empower developers to create robust templates without sacrificing performance or usability. 
- Adheres strictly to Go's templating conventions, ensuring a seamless experience for those familiar with Go's native tools.
- Naming conventions across functions are standardized for predictability and ease of use.
- Emphasizes error handling, preferring to safe defaults over panics.
- Provide a clear and comprehensive documentation to help users understand the library and its features.
- Maintain a high level of code quality, ensuring that the library is well-tested, performant, and reliable.
- Continuously improve and optimize the library to meet the needs of the community.
- Avoids any external dependencies within template functions, ensuring all operations are self-contained and reliable.
- Performance is a key consideration, with a focus on optimizing the library for speed and efficiency.
