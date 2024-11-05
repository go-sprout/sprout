---
description: 'Managing multiple registries can be simplified with RegistryGroup feature.'
---

# Loader System (Registry Groups)

## Introduction

The `RegistryGroup` struct enables the grouping of multiple registries, simplifying their management and enhancing modularity within the `Handler`. This approach promotes cleaner code organization and facilitates the reuse of registry configurations across different projects.

{% hint style="info" %}
**Considerations**

Using a `RegistryGroup` can help you manage multiple registries. But can lead to registering undesired functions if you are not careful. Be sure to review the functions in each registry to avoid conflicts or unexpected behavior.
{% endhint %}

## How to use a registry group

If you have created a `RegistryGroup` or want to use a built-in group, you can add it to your handler using the `AddGroups` method:

```go
err := handler.AddGroups(group1, group2)
if err != nil {
  // Handle error
}
```
You can also use the option to add registries when initializing the handler:

```go	
handler := sprout.New(
  sprout.WithGroups(group1, group2),
)
```

In this example, `group1` and `group2` are added to the `handler`, allowing all registries within these groups to be registered collectively.



## How to create a registry group

To show how to create your own registry group, go to [how-to-create-a-registry-group.md](../advanced/how-to-create-a-registry-group.md "mention")
