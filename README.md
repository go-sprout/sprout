# Sprout 🌱

> [!NOTE]
> Sprout is an evolved variant of the [Masterminds/sprig](https://github.com/Masterminds/sprig) library, reimagined for modern Go versions. It introduces fresh functionalities and commits to maintaining the library, picking up where Sprig left off. Notably, Sprig had not seen updates for two years and was not compatible beyond Golang 1.13, necessitating the creation of Sprout.

# Table of Content 



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

### Development Philosophy (Currently in reflexion to create our)

Our approach to extending and refining Sprout was guided by several key principles:

- Empowering layout construction through template functions.
- Designing template functions that avoid returning errors when possible, instead displaying default values for smoother user experiences.
- Ensuring template functions operate solely on provided data, without external data fetching.
- Maintaining the integrity of core Go template functionalities without overrides.






