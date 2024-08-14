---
description: Need two names for one function in your code? Aliases are the solution.
---

# Function Aliases

The function aliasing feature introduces a seamless way for developers to maintain backward compatibility while transitioning to new function names in their codebases. This mechanism is designed to ensure that older code remains functional without immediate modifications, even as the library evolves.

{% hint style="info" %}
This feature is crucial for migrating from Sprig v3.2.3 or when upgrading between Sprout versions.
{% endhint %}

To configure aliases, you must use the Sprout function handler.&#x20;

## How It Works

* **Definition:** An alias acts as a secondary name for a function, referring to the original implementation in memory, the function are not duplicated at runtime.
* **Usage:** When the deprecated (aliased) function names are used, the code behaves as if the new function names were called, ensuring compatibility.
* **Example:** Suppose `oldFunc` is deprecated in favor of `newFunc` in your template. The alias allows code calling `oldFunc` to execute `newFunc` transparently.

## Usage

### Add one alias

To use the function aliases feature, you just need to use the configuration function `WithAlias()` as shown below:

<pre class="language-go"><code class="lang-go"><strong>handler:= sprout.New(sprout.WithAlias("newFunc", "oldFunc"))
</strong>
template.New("base").Funcs(handler.Build()).Parse("{{ newFunc }}")
</code></pre>

This creates a mapping between an old function name (`oldFunc`) and a new one (`newFunc`). Calls to `oldFunc` within the template are redirected to execute `newFunc`. This enables the template to parse and execute using the new function name seamlessly.

### Add more than one aliases for te same function

To add more aliases for the same original function, simply add more parameters to the `WithAlias` function:

<pre class="language-go"><code class="lang-go"><strong>handler := sprout.New(sprout.WithAlias("newFunc", "oldFunc", "secondAlias"))
</strong></code></pre>

This creates two aliases for the function `newFunc`. Calling `oldFunc` or `secondAlias` will execute `newFunc`.

### Add aliases for multiples function at once

To map multiple functions, use the same strategy we use for backward compatibility with Sprig by creating a `FunctionAliasMap` and injecting it into your Function Handler with the `WithAliases` function:

```go
var myAliases = sprout.FunctionAliasMap{
	"newFunc": {"oldFunc", "secondAlias"},
	"hello":   {"hi", "greet"},
}

handler := sprout.New(sprout.WithAliases(myAliases))
```

This creates two aliases for two methodes (4 in total). Calling `oldFunc` or `secondAlias` with execute `newFunc` and calling `hi` or `greet` will execute `hello`.

## Best Practices

* **Documentation:** Clearly document all aliases on your codebase to avoid confusion.
* **Deprecation Notices:** Use comments or documentation to inform users of deprecated functions and encourage the adoption of new names.
* **Gradual Transition:** Allow a transition period where both old and new function names are supported. For my projects, I use a version frame of 5 minors or 1 major, if this can guide you.

## Add aliases on your registry

To add aliases on your registry, see [how-to-create-a-registry.md](../advanced/how-to-create-a-registry.md "mention")page.
