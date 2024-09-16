---
description: 'Need just 10 functions or want to create your own? Introducing: the Registry'
---

# Loader System (Registry)

{% hint style="info" %}
This feature was developed based on a Request for Comments (RFC) discussed within the Sprout project community. The RFC outlines the need for a modular registry system to improve the flexibility and maintainability of the codebase. For more detailed information on the discussion and rationale behind this implementation, you can review the full RFC [here](https://github.com/orgs/go-sprout/discussions/31).
{% endhint %}

## Introduction

Sprout empowers you to create and manage your own registry of functions—a versatile collection that you can store, share, and seamlessly integrate into your projects. Additionally, you can leverage functions contributed by other users, enhancing collaboration and expanding the functionality of your applications.

## Goals

* **Optimize Build Time**: Use only the registry functions you need to reduce both build time and binary size.
* **Ease of Use**: Import a registry with just one line of code for quick integration.
* **Reusability**: Reuse functions across multiple projects without redundancy.
* **Collaboration**: Incorporate registries created by others to enhance your projects.
* **Sharing**: Share your registry with others or keep it private.
* **Versioning**: Manage specific versions of a registry as a Go package to ensure consistency across projects.

## How to use a registry

To use a registry in a project, follow these steps:

```go
// Initialize your sprout handler
handler := sprout.New()

// Add one registry
handler.AddRegistry(ownregistry.NewRegistry())
// Add more than one at the same time
handler.AddRegistries(
  ownregistry.NewRegistry(),
  // ...
)

// Build your FunctionMap based on your handler
tpl := template.Must(
    template.New("base").Funcs(handler.Build()).ParseGlob("*.tmpl"),
  )
```

You can also use the option to add registries when initializing the handler:

```go
handler := sprout.New(
  sprout.WithRegistries(reg1.NewRegistry(), reg2.NewRegistry()),
)
```

This code sets up your project to utilize the functions from your custom registry, making it easy to integrate and extend functionality.

## How to create a registry

To show how to create your own registry, go to [how-to-create-a-registry.md](../advanced/how-to-create-a-registry.md "mention")
