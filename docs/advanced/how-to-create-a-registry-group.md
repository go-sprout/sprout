# How to create a registry group

## File Naming Conventions

* `{{registry_group_name}}.go`: This file defines the registry group, including key components like structs, interfaces, constants, and variables.
* `{{registry_group_name}}_test.go`: Includes tests for the registry group to ensure it functions as expected.

{% hint style="info" %}
This structure ensures consistency and maintainability across different registries, making it easier for developers to contribute and collaborate effectively.\
\
For the rest of conventions please read [templating-conventions.md](../introduction/templating-conventions.md "mention").
{% endhint %}


## Creating a Registry Group

1. **New Repository**: You can start by creating a new repository on your GitHub account or organization. This will house your registry group functions.
2. **Contributing to Official Registry**: To add a new registry group to the official registry, submit a pull request (PR) after a discussion inside an issue. The official registry groups are organized under the `group/` folder.

To start, in your `{{registry_group_name}}.go` file, create a function called `RegistryGroup` respecting the following signature:

```go
func RegistryGroup() *sprout.RegistryGroup {
  return sprout.NewRegistryGroup(
    // Add your registries here
    // registry.NewRegistry(),
  )
}
```

You can also create a group directly in your codebase without the need to create a new package. This is useful when you want to group registries that are specific to a project or a use case.
```go
group := sprout.NewRegistryGroup(registry.NewRegistry())

handler.AddGroups(group)
```

Once your group is defined, you can start using it in your projects. 🎉
