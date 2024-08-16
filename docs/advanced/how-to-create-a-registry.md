# How to create a registry

## File Naming Conventions

* `{{registry_name}}.go`: This file defines the registry, including key components like structs, interfaces, constants, and variables.
* `functions.go`: Contains the implementation of exported functions, making them accessible to other developers.
* `functions_test.go`: Includes tests for the exported functions to ensure they function as expected.
* `helpers.go`: Contains internal helper functions that support the registry but are not exposed for public use.
* `helpers_test.go`: Holds tests for the helper functions to validate their reliability.

{% hint style="info" %}
This structure ensures consistency and maintainability across different registries, making it easier for developers to contribute and collaborate effectively.
{% endhint %}

## Creating a Registry

1. **New Repository**: You can start by creating a new repository on your GitHub account or organization. This will house your registry functions.
2. **Contributing to Official Registry**: To add your functions to the official registry, submit a pull request (PR). The official registries are organized under the `registry/` folder.

{% hint style="info" %}
You can found an example of a registry under `registry/_example`.
{% endhint %}

To start, in your `{{registry_name}}.go` file, start by creating a struct that implements the `Registry` interface. This struct will manage your custom functions and connect them to the handler.

```go
package ownregistry

import (
  "github.com/go-sprout/sprout"
)

// OwnRegistry struct implements the Registry interface, embedding the Handler to access shared functionalities.
type OwnRegistry struct {
  handler sprout.Handler // Embedding Handler for shared functionality
}

// NewRegistry initializes and returns a new instance of your registry.
func NewRegistry() *OwnRegistry {
  return &OwnRegistry{}
}

// Uid provides a unique identifier for your registry.
func (or *OwnRegistry) Uid() string {
  return "ownRegistry" // Ensure this identifier is unique and uses camelCase
}

// LinkHandler connects the Handler to your registry, enabling runtime functionalities.
func (or *OwnRegistry) LinkHandler(fh sprout.Handler) error {
  or.handler = fh
  return nil
}

// RegisterFunctions adds the provided functions into the given function map.
// This method is called by an Handler to register all functions of a registry.
func (or *OwnRegistry) RegisterFunctions(funcsMap sprout.FunctionMap) {
  // Example of registering a function
  sprout.AddFunction(funcsMap, "yourFunction", or.YourFunction)
}

// OPTIONAL: Your registry don't needs to register aliases to work.
// RegisterAliases adds the provided aliases into the given alias map.
// method is called by an Handler to register all aliases of a registry.
func (or *OwnRegistry) RegisterAliases(aliasMap sprout.FunctionAliasMap) error {
  // Example of registering an alias
  sprout.AddAlias(aliasMap, "yourFunction", "yourAlias")
}
```

After create your registry structure and implement the `Registry` interface, you can start to define your functions in `functions.go`, you can access all features of the handler through&#x20;

```go
// YourFunction is an example function that returns a string and an error.
func (or *OwnRegistry) YourFunction() (string, error) {
  return "Hello, World!", nil
}
```

{% hint style="danger" %}
**Important:** Make sure to write tests for your functions in `functions_test.go` to validate their functionality.
{% endhint %}

Once your registry is defined and functions are implemented, you can start using it in your projects. 🎉
