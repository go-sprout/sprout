# How to create a handler

## Introduction

Creating a custom handler in Sprout allows you to manage function registries, define custom logging, and handle errors according to your specific needs. Here’s a step-by-step guide on how to create your own custom handler.

### Step 1: Understand the Handler Interface

The `Handler` interface in Sprout defines the basic methods required to manage registries and functions. A typical handler in Sprout must implement the following methods:

* `Logger() *slog.Logger`: Returns the logger instance used for logging.
* `AddRegistry(registry Registry) error`: Adds a single registry to the handler.
* `AddRegistries(registries ...Registry) error`: Adds multiple registries to the handler.
* `Functions() FunctionMap`: Returns the map of registered functions.
* `Aliases() FunctionAliasMap`: Returns the map of function aliases.
* `Build() FunctionMap`: Builds and returns the complete function map, ready to be used in templates.

### Step 2: Create Your Custom Handler Struct

Create a struct that will serve as your custom handler. This struct should store the function map, alias map, and any other configurations you need.

```go
type MyCustomHandler struct {
    logger        *slog.Logger
    registries    []sprout.Registry
    funcsMap      sprout.FunctionMap
    funcsAlias    sprout.FunctionAliasMap
}
```

### Step 3: Implement the Handler Interface

Next, implement the `Handler` interface methods in your custom struct. I take logger as example

```go
func (h *MyCustomHandler) Logger() *slog.Logger {
    return h.logger
}
```

### Step 4: Initialize and Customize Your Handler

Create a function to initialize your custom handler, setting up the logger and any registries you need.

```go
func NewMyCustomHandler() *MyCustomHandler {
    return &MyCustomHandler{
        logger:        slog.New(slog.NewTextHandler(os.Stdout)),
        registries:    make([]sprout.Registry, 0),
        funcsMap:      make(sprout.FunctionMap),
        funcsAlias:    make(sprout.FunctionAliasMap),
    }
}
```

### Step 5: Use Your Custom Handler

With your custom handler in place, you can now use it to register functions and integrate it with your templates.

```go
handler := NewMyCustomHandler()
handler.AddRegistry(std.NewRegistry())

tpl, err := template.New("example").Funcs(handler.Build()).Parse(`{{ myFunc }}`)
if err != nil {
    log.Fatal(err)
}
tpl.Execute(os.Stdout, nil)
```

## Conclusion

Creating a custom handler in Sprout allows you to extend and customize the templating environment according to your needs. By implementing the `Handler` interface, you gain full control over how functions are registered, managed, and used in your templates. This flexibility is key to building robust and maintainable Go applications.
