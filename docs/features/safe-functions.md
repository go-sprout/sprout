---
description: >-
  When you don't want to stop the execution of your template when an error
  occurs, you can use the function-safe feature.
---

# Safe Functions

The **Safe Functions** feature in Sprout allows for more resilient template rendering by ensuring that errors in function execution do not halt the rendering process. Instead, these functions return a default value when an error occurs. This feature is disabled by default but can be enabled as needed.

## How It Works

When Safe Functions are enabled, for every function in your template, a corresponding "safe" version is automatically created.&#x20;

For example, if you have a function called `toInt`, enabling Safe Functions will also create `safeToInt`.&#x20;

The "safe" version of the function will catch any errors that occur during execution and return a default value instead of stopping the rendering process.

## Usage

To enable Safe Functions, you need to configure your handler using the `WithSafeFuncs` option when creating the handler:

```go
handler := sprout.New(sprout.WithSafeFuncs(true))
```

You activate Safe Functions by setting the `enabled` parameter to `true` when you create your handler.

## Usage in Templates

Once Safe Functions are enabled, you can use them in your templates by simply prefixing the original function name with `safe`. For instance:

* Instead of using `toInt`, you would use `safeToInt`.
* If the function `safeToInt` encounters an error, it will return a default value rather than causing the template rendering to fail.
* You can continue to use `toInt` with the default error mechanism.

## Important Considerations

* **Disclaimer:** Enabling Safe Functions will effectively double the number of functions available in your template, as each original function will now have a corresponding safe version.
* **Performance:** While Safe Functions improve the robustness of template rendering, they may introduce some overhead due to the additional error handling. Use them judiciously based on your needs.
