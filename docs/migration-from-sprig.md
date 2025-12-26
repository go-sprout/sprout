---
icon: route
description: >-
  Coming from Sprig and looking to use Sprout? You're in the right place for a
  complete guide on making the transition. This guideline will help you navigate
  the differences.
---

# Migration from Sprig

{% hint style="info" %}
This page evolves with each version based on modifications and community feedback. Having trouble following this guide?

[Open an issue](https://github.com/go-sprout/sprout/issues/new/choose) to get help, and contribute to improving this guide for future users. :seedling: :purple\_heart:
{% endhint %}

## <mark style="color:purple;">Introduction</mark>

Sprout is a modern templating engine inspired by [Sprig ](https://github.com/Masterminds/sprig)but with enhanced features and a more modular approach. Migrating from Sprig to Sprout involves understanding these differences and adjusting your templates and code accordingly.

## <mark style="color:purple;">Key Differences</mark>

### 1. **Registry System**

* **Sprig:** Functions are globally available.
* **Sprout:** Functions are grouped into registries for modularity. Each registry can be added to the handler as needed.

{% hint style="info" %}
**Migration Tip**

List the registries needed for your project and register them with the handler. If you're unsure, you can safely register all built-in registries.
{% endhint %}

### 2. **Handler and Function Management**

* **Sprig:** Functions are accessed directly.
* **Sprout:** Functions are managed by a handler, allowing for better control over function availability, error handling, and logging.

{% hint style="success" %}
**Migration Top**

Nothing to do, using the new handler is enough.
{% endhint %}

### 3. **Error Handling**

* **Sprig:** Limited and inconsistent error handling, with some functions causing panics (see [#panicking-functions](migration-from-sprig.md#panicking-functions "mention")), and not fully adhering to Go template standards.
* **Sprout:** Offers configurable error handling strategies, including returning default values, or return error to stop template generation (default), providing a more consistent and flexible approach.

{% hint style="success" %}
**Migration Tip**

You can learn mote about the safe strategy here: [safe-functions.md](features/safe-functions.md "mention")
{% endhint %}

### 4. **Function Aliases**

* **Sprig:** Functions are accessed by a single name, deprecated functions are duplicated in code.
* **Sprout:** Supports function aliases, allowing multiple names for the same function.

{% hint style="info" %}
**Migration Tip**

Use `WithAlias` or `WithAliases` to set up function aliases as needed when creating your handler.
{% endhint %}

### 5. Following the Go Template Convention

* **Sprig**: No error for naming convention are followed.
* **Sprout**: Design to adhere closely to Go's native template conventions, ensuring compatibility and a consistent experience for developers familiar with Go templates. This adherence helps reduce surprises and ensures that templates behave as expected according to Go's standard practices.

{% hint style="info" %}
**Migration Tip**\
Takes a look over renamed functions to change it in your template or use aliases.
{% endhint %}

## <mark style="color:purple;">How to Transition</mark>

### Step 1: Identify and Organize Functions

Review your existing templates and identify the functions you rely on from Sprig. Determine whether these functions have direct equivalents in Sprout or if they require using the `sprigin` package for backward compatibility.

### Step 2: Refactor Templates

Update your templates to replace Sprig function calls with their Sprout equivalents. If a direct replacement is unavailable, consider using aliases or the `sprigin` package to maintain functionality.

### Step 3: Register Functions and Handlers

Set up your Sprout environment by creating a handler and registering the necessary function registries. This step ensures that your templates have access to the required functions.

### Step 4: Test and Validate

Thoroughly test your migrated templates to ensure that all functions behave as expected. Pay particular attention to error handling and any deprecated functions that may require adjustments.

## <mark style="color:purple;">How to Transition for your end-users</mark>

You use sprig or sprout for end-users and want to migrate ?\
Transitioning from Sprig to Sprout can be smooth with proper guidance. Here's a detailed plan to ensure confidence during the migration:

{% hint style="info" %}
You need more information now, contact maintainers or open [a discussion on the repository](https://github.com/orgs/go-sprout/discussions/categories/q-a).
{% endhint %}

***

1. **Communicate the Purpose of the Migration**\
   Explain the reasons for switching to Sprout, emphasizing improvements such as better performance, modular function registries, enhanced error handling, and new features like function notices and safe functions. Providing context will minimize resistance and help users understand the long-term benefits.
2. **Attach Your Logger to Sprout**\
   Starting from Sprout v1.0.0, deprecated functions are flagged through the logger. This ensures end-users are informed of upcoming deprecations without breaking the existing code.
3. **Update Frequently**\
   We follow a no-breaking-change policy for five minor versions. For example, a function deprecated in v1.2 will be removed in v1.7. End-of-life notices for functions will be automatically logged, so you don’t need to intervene.
4. **Provide Feedback**\
   If you or your end-users encounter any feedback during the migration—whether it’s about missing functions, bugs, or positive experiences—reach out to the maintainers or contribute to the project on GitHub. Your input helps improve the library for everyone.

***

{% hint style="info" %}
As a library developer, you can extend Sprout by creating your [own function registry](advanced/how-to-create-a-registry.md). This allows you to add custom functions to meet specific use cases. Additionally, you can use [**notices** ](features/function-notices.md)to inform your end-users about important updates or issues during template execution, such as deprecations or debugging hints. These notices can enhance the user experience by providing real-time feedback, ensuring a smoother transition to the new system, and allowing users to handle function changes more efficiently.
{% endhint %}

{% hint style="success" %}
Our maintainers and collaborators can assist you if you have question, don't hesitate to [open a discussion on GitHub](https://github.com/orgs/go-sprout/discussions/categories/q-a) !
{% endhint %}

## <mark style="color:purple;">Migrating Common Functions</mark>

Many functions in Sprig have direct equivalents in Sprout, but they might be organized differently or require registration in a handler.

### Example: Simple Function Migration

**Sprig:**

```go
{{ upper "hello" }}
```

**Sprout**:

```go
{{ toUpper "hello" }}
```

### Example: Using Aliases

**Sprig:**

```go
{{ upper "hello" }}
```

**Sprout**:

```go
handler := sprout.New(
    sprout.WithAlias("toUpper", "upper"),
)
funcs := handler.Build()
```

You can continue to use the same function name inside your template

```go
{{ upper "hello" }}
```

## <mark style="color:purple;">Panicking Functions</mark>

In Sprig, errors within certain functions cause a panic. In contrast, Sprout opts for returning nil or an empty value, improving safety and predictability.

**Old Behavior (Sprig)**: Triggers a panic on error

```go
if err != nil {
  panic("deepCopy error: " + err.Error())
}
```

**New Behavior (Sprout)**: Returns nil or an empty value on error

```go
if err != nil {
  return nil, err
}
```

#### Methods that previously caused a panic in Sprig :

* DeepCopy
* MustDeepCopy
* ToRawJson
* Append
* Prepend
* Concat
* Chunk
* Uniq
* Compact
* Slice
* Without
* Rest
* Initial
* Reverse
* First
* Last
* Has
* Dig
* RandAlphaNumeric
* RandAlpha
* RandAscii
* RandNumeric
* RandBytes

## <mark style="color:purple;">Function-Specific Changes</mark>

#### MustDeepCopy

* **Sprig**: Accepts `nil` input, causing an internal panic.
* **Sprout**: Returns `nil` if input is `nil`, avoiding panic.

#### Rand Functions

* **Sprig**: Causes an internal panic if the length parameter is zero.
* **Sprout**: Returns an empty string if the length is zero, ensuring stability.

#### DateAgo

* **Sprig**: Does not support int32 and \*time.Time; returns "0s".
* **Sprout**: Supports int32 and \*time.Time and returns the correct duration.

#### DateRound

* **Sprig**: Returns a corrected duration in positive form, even for negative inputs.
* **Sprout**: Accurately returns the duration, preserving the sign of the input.

#### Base32Decode / Base64Decode

* **Sprig**: Decoding functions return the error string when the input is not a valid base64 encoded string.
* **Sprout**: Decoding functions return an empty string if the input is not a valid base64 encoded string, simplifying error handling.

#### Dig

* **Sprig**: The `dig` function requires a default value as the second-to-last argument: `{{ dig "a" "b" "default" $dict }}`. It returns the last map in the access chain instead of the final value.
* **Sprout**: The `dig` function does not include a default value: `{{ $dict | dig "a" "b" }}`. Use the `default` filter instead: `{{ $dict | dig "a" "b" | default "default" }}`. It returns the final value in the chain, regardless of its type.

{% hint style="info" %}
If you use `sprigin.FuncMap()`, the `dig` function retains Sprig's signature with the default value parameter for backward compatibility.
{% endhint %}

#### ToCamelCase / ToPascalCase

* **Sprig**: The `toCamelCase` return value are in PascalCase. No `toPascalCase` function is available.
* **Sprout**: The `toCamelCase` function returns camelCase strings, while the `toPascalCase` function returns PascalCase strings.

#### Merge / MergeOverwrite

* **Sprig**: The `merge` and `mergeOverwrite` functions does dereference when second value are the default golang value (example: `0` for int).
* **Sprout**: The `merge` and `mergeOverwrite` functions does not dereference and keep the second value as is (example: `0` for int).

## <mark style="color:purple;">Deprecated Features</mark>

Sprout has deprecated certain features for better security and performance. For example, **direct cryptographic operations in templates are discouraged.**

{% hint style="info" %}
**Migration Tip**

Review your template functions and avoid using deprecated features.\
Move critical operations outside of templates to maintain security.
{% endhint %}

{% hint style="warning" %}
Perform cryptographic operations (listed in `crypto` package) outside of templates. the [`crypto`registry ](registries/crypto.md)will be drop in few versions.
{% endhint %}

All deprecated features are flagged with <mark style="color:red;">`// ! DEPRECATED`</mark> in codebase.\
A complete list will be available here when the v1 of Sprout are released.

## <mark style="color:purple;">Conclusion</mark>

Migrating from Sprig to Sprout offers significant benefits, including improved error handling, modular function management, and enhanced compatibility with modern Go practices. While the `sprigin` package provides a bridge for backward compatibility, fully embracing Sprout’s native capabilities will lead to a more stable and maintainable codebase. For further details on Sprout’s features and API, consult the [official Sprout documentation](https://docs.atom.codes/sprout).
