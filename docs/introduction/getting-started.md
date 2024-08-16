---
description: A quick start guide to understand and use Sprout in your project
---

# Getting Started

{% hint style="info" %}
This page evolves with each version based on modifications and community feedback. Having trouble following this guide?&#x20;

[Open an issue](https://github.com/go-sprout/sprout/issues/new/choose) to get help, and contribute to improving this guide for future users. :seedling: :purple\_heart:&#x20;
{% endhint %}

## Introduction

Sprout is a powerful and flexible templating engine designed to help you manage and organize template functions efficiently. It allows developers to register and manage functions across different registries, offering features such as aliasing, error handling, and logging.

## Installation

To get started with Sprout, first install the package:

```bash
get -u github.com/go-sprout/sprout
```

## Usage

### Creating a Handler

A handler in Sprout is responsible for managing the function registries and functions. The `DefaultHandler` is the primary implementation provided by Sprout.

```go
import "github.com/go-sprout/sprout"

handler := sprout.New()
```

### Customizing the Handler

Sprout supports various customization options using handler options:

*   **Logger Configuration:**

    You can customize the logging behavior by providing a custom logger:

    ```go
    logger := slog.New(slog.NewTextHandler(os.Stdout))
    handler := sprout.New(sprout.WithLogger(logger))
    ```
*   **Aliases Management:**\
    You can specify your custom aliases directly on your handler:

    ```go
    handler := sprout.New(sprout.WithAlias("originalFunc", "alias"))
    ```

    See more below or in dedicated page [function-aliases.md](../features/function-aliases.md "mention").

### Working with Registries

Registries in Sprout are groups of functions that can be added to a handler. They help organize functions and optimize template performance.

#### Using a built-in registry

You can retrieve all built-ins registries and functions under [Broken link](broken-reference "mention").

#### Create your own registry

To create your own, see the dedicated page [how-to-create-a-registry.md](../advanced/how-to-create-a-registry.md "mention").

### Adding a Registry to a Handler

Once your registry is implemented, you can add it to a handler:

```go
import (
  "github.com/go-sprout/sprout"
  "github.com/go-sprout/sprout/registry/conversion" // toString, toInt, toBool, ...
  "github.com/go-sprout/sprout/registry/std" // default, empty, any, all, ...
)
handler := sprout.New()
handler.AddRegistry(conversion.NewRegistry())
handler.AddRegistry(std.NewRegistry())
```

You can also add multiple registries at once:

```go
handler.AddRegistries(conversion.NewRegistry(), std.NewRegistry())
```

### Function Aliases

Sprout supports function aliases, allowing you to call the same function by different names.

#### Adding Aliases

You can add aliases for functions in your handler configuration:

```go
handler := sprout.New(
    sprout.WithAlias("originalFunc", "alias1", "alias2"),
)
```

#### Using Aliases

Aliases are automatically resolved when you build your handler’s function map.

### Building Function Maps

To use Sprout with templating engines like `html/template` or `text/template`, you need to build the function map:

```go
funcs := handler.Build()
tpl := template.New("example").Funcs(funcs).Parse(`{{ hello }}`)
```

This prepares all registered functions and aliases for use in templates.

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

## Conclusion

Sprout provides a structured and powerful way to manage template functions in Go, making it easier to build, maintain, and extend templating functionality. With features like custom registries, aliases, and configurable error handling, Sprout can significantly enhance your templating experience.

For more informations or questions, refer to the [Sprout GitHub repository](https://github.com/go-sprout/sprout).
