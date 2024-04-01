# Sprout 🌱

> [!NOTE]
> Sprout is an evolved variant of the [Masterminds/sprig](https://github.com/Masterminds/sprig) library, reimagined for modern Go versions. It introduces fresh functionalities and commits to maintaining the library, picking up where Sprig left off. Notably, Sprig had not seen updates for two years and was not compatible beyond Golang 1.13, necessitating the creation of Sprout.

## Table of Contents

- [Table of Contents](#table-of-contents)
- [Transitioning from Sprig](#transitioning-from-sprig)
- [Performence Benchmarks](#performence-benchmarks)
  - [Sprig v3.2.3 vs Sprout v0.1](#sprig-v323-vs-sprout-v01)
- [Usage](#usage)
  - [Integrating the Sprout Library](#integrating-the-sprout-library)
  - [Template Function Invocation](#template-function-invocation)
- [Development Philosophy (Currently in reflexion to create our)](#development-philosophy-currently-in-reflexion-to-create-our)

## Transitioning from Sprig

For those looking to switch from Sprig to Sprout, the process is straightforward and involves just a couple of steps:
1. Ensure your project uses Sprig's last version (v3.2.3).
2. Update your import statements and package references as shown below:
```diff
import (
-  "github.com/Masterminds/sprig/v3"
+  "github.com/42atomys/sprout"

  "html/template"
)

tpl := template.Must(
  template.New("base").
-   Funcs(sprig.FuncMap()).
+   Funcs(sprout.FuncMap()).
    ParseGlob("*.html")
)
```

## Performence Benchmarks

To see all the benchmarks, please refer to the [benchmarks](benchmarks/README.md) directory.

### Sprig v3.2.3 vs Sprout v0.1
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

So, Sprout v0.1 is approximately 53.1% faster and uses 15.7% less memory than Sprig v3.2.3. 🚀


## Usage

**For Template Creators**: Refer to the comprehensive function guide in Sprig's documentation for detailed instructions and examples across over 100 template functions.

**For Go Developers**: Integrate Sprout into your applications by consulting our API documentation available on GoDoc.org.

For general library usage, proceed as follows.

### Integrating the Sprout Library
To utilize Sprout's functions within your templates:


```golang
import (
  "github.com/42atomys/sprout"
  "html/template"
)

// Ensure the FuncMap is set before loading the templates.
tpl := template.Must(
  template.New("base").Funcs(sprout.FuncMap()).ParseGlob("*.html")
)
```

### Template Function Invocation
Adhering to Go's conventions, all Sprout functions are lowercase, differing from method naming which employs TitleCase. For instance, this template snippet:


```golang
{{ "hello!" | upper | repeat 5 }}
```
Will output:
```
HELLO!HELLO!HELLO!HELLO!HELLO!
```

## Development Philosophy (Currently in reflexion to create our)

Our approach to extending and refining Sprout was guided by several key principles:

- Empowering layout construction through template functions.
- Designing template functions that avoid returning errors when possible, instead displaying default values for smoother user experiences.
- Ensuring template functions operate solely on provided data, without external data fetching.
- Maintaining the integrity of core Go template functionalities without overrides.






