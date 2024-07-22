# Sprout Registry

## Introduction

Sprout allows you to create and manage your own registry of functions. 
This registry is a collection of functions that you can store, share, and use in your projects. 
You can also utilize functions from other users in your own projects.


## Goals

- **Optimize build time**: By using only the registry functions you need, you can reduce build time and the size of your binary.
- **Ease of use**: Import the registry you need with a single line of code.
- **Reusability**: Reuse functions across different projects without duplication.
- **Collaboration**: Utilize registries created by other users in your projects.
- **Sharing**: Share your registry with others or keep it private.
- **Versioning**: Use specific versions of a registry as a Go package to ensure consistency.


## How to create a registry

To create a registry, you can:

- Create a new repository in your GitHub account or organization.
- Use an existing repository.
- Create a pull request (PR) to add your functions to the [official registry](https://github.com/go-sprout/sprout).

If you are contributing to the [official registry](https://github.com/go-sprout/sprout), all registries are located under the registry/ folder. We follow these naming conventions for files:

- `{{registry_name}}.go`: Contains the registry definition, including structs, interfaces, constants, and variables.
- `functions.go`: Contains the implementation of exported functions that can be used by other developers.
- `functions_test.go`: Contains tests for the exported functions to ensure they work correctly.
- `helpers.go`: Contains helper functions that are used internally within the registry but are not exported for public use.
- `helpers_test.go`: Contains tests for the helper functions.


### Create your first registry

> [!NOTE]
> You can found an example of a registry under `registry/_example`.

1. Define the Registry

In your `{{registry_name}}.go` file, create a struct that implements the Registry interface. 
This struct will manage your functions and link them to the handler.


```go
package ownregistry

import (
  "github.com/go-sprout/sprout/registry"
)

// This struct implements the Registry interface and defines your custom registry
// and embeds the Handler to utilize its functionalities
type OwnRegistry struct {
    handler *registry.Handler // Embedding Handler to leverage shared functionality
}

// This function creates and returns a new instance of your registry.
// By convention, the function should be named NewRegistry.
func NewRegistry() *OwnRegistry {
    return &OwnRegistry{}
}

// This function provides a unique identifier for your registry.
func (or *OwnRegistry) Uid() string {
    return "ownRegistry" // Ensure this is unique and in camel case
}

// This function links the Handler to your registry, enabling runtime functionalities.
func (or *OwnRegistry) LinkHandler(fh registry.Handler) {
    or.handler = &fh
}

```

2. Implement Functions


In your `functions.go` file, implement your functions and register them in the registry.


```go
package ownregistry

import (
  "github.com/go-sprout/sprout/registry"
)

// This method registers your functions with the registry, making them available for use.
func (or *OwnRegistry) RegisterFunctions(funcsMap registry.FunctionMap) {
  // Register your functions here
  // "yourFunction" is the name of the function that will be used in the template
  // or.YourFunction is the function that will be executed when the template calls "yourFunction"
  registry.AddFunction(funcsMap, "yourFunction", or.YourFunction)
}

// This is an example function that performs an action and returns a result.
func (or *OwnRegistry) YourFunction() (string, error) {
  return "Hello, World!", nil
}
```

> [!IMPORTANT]
> Ensure to test your functions in functions_test.go to verify their correctness


3. Use Your Registry

You are now ready to use your registry in your project. :tada:


## How to use a registry

You are now ready to use your registry in your project. Here’s how to do it:


```go
package main

import (
  "github.com/go-sprout/sprout"
  "github.com/username/project/ownregistry"
  "html/template"
)

func main() {
  handler := sprout.NewFunctionHandler()
  handler.AddRegistry(ownregistry.NewRegistry())

  tpl := template.Must(
    template.New("base").
      Funcs(handler.Registry()).
      ParseGlob("*.tmpl")
  )
}
```
