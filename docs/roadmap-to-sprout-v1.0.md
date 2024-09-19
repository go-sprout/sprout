---
icon: rocket-launch
description: The roadmap to grow the sprout
---

# Roadmap to Sprout v1.0

## Key Objectives

{% hint style="info" %}
All objectives are get from feedback, suggestions and personal knowledge. You can discuss about the v1.0 directly [in the issue on GitHub](https://github.com/go-sprout/sprout/issues/1).
{% endhint %}

### :white\_check\_mark: **Minimize Dependencies - **<mark style="color:green;">**DONE**</mark>

Reduce the number of external dependencies to mitigate frequent update cycles, making Sprout more stable and lightweight.

{% hint style="success" %}
Dependencies have been minimized and optimized across all registries.
{% endhint %}

### :white\_check\_mark: **Enhanced Documentation - **<mark style="color:green;">**DONE**</mark>

Provide comprehensive, easy-to-understand documentation that covers all functionalities, use cases, and examples to improve the developer experience.

{% hint style="success" %}
This feature are implemented on v0.5.0, documentations can be found here:&#x20;

<mark style="color:green;">You are on the official documentation site</mark> :tada:
{% endhint %}

### :white\_check\_mark: **Conventional Function Naming - **<mark style="color:green;">**DONE**</mark>

Establish clear, consistent naming conventions for functions to enhance code readability and maintainability. Unlike Sprig, where function naming varies between camelCase, and snake\_case, and similar functions lack consistent prefixing, Sprout will introduce a standardized approach to function naming. This will make the library more intuitive and reduce the learning curve for new users.

{% hint style="success" %}
This convention are defined and available here:

[templating-conventions.md](introduction/templating-conventions.md "mention")
{% endhint %}

### :white\_check\_mark: **Reduce memory fingerprint - **<mark style="color:green;">**DONE**</mark>

Aim to minimize memory allocations as much as possible to alleviate the burden on the garbage collector in large-scale applications. By optimizing the way memory is used within the framework, we ensure that Sprout is not only efficient in its functionality but also in its resource consumption. This approach contributes to overall better performance and scalability of applications using Sprout.

### :white\_check\_mark: **Native Error Handling - **<mark style="color:green;">**DONE**</mark>

Follow default go template error handling mechanisms for all functions to ensure that errors are managed gracefully and efficiently.

{% hint style="success" %}
This feature are implemented on v0.6.0, documentation can be found here:

[safe-functions.md](features/safe-functions.md "mention")
{% endhint %}

### :hourglass:**Expanded Function Set - **<mark style="color:orange;">**IN PROGRESS**</mark>

Add a broader array of functions without imposing limitations, enabling users to accomplish more tasks directly within the framework.

### :white\_check\_mark: **Customizable Function Loading - **<mark style="color:green;">**DONE**</mark>

Allow users to customize which functions to load into their runtime environment, preventing unnecessary resource consumption and enhancing performance.

{% hint style="success" %}
This feature are implemented on v0.5.0, documentations can be found here:&#x20;

[loader-system-registry.md](features/loader-system-registry.md "mention")&#x20;

[how-to-create-a-registry.md](advanced/how-to-create-a-registry.md "mention")
{% endhint %}

### :white\_check\_mark: **Function Aliasing - **<mark style="color:green;">**DONE**</mark>

Enable the creation of aliases for functions outside of the library, providing flexibility and convenience in how functions are accessed and utilized.

{% hint style="success" %}
This feature are implemented on v0.3.0, documentation can be found here :&#x20;

[function-aliases.md](features/function-aliases.md "mention")
{% endhint %}

### :white\_check\_mark: **Function Notices - **<mark style="color:green;">**DONE**</mark>

When you are a middle-app (between sprout and the user how write the template), you need to be careful when you upgrade a template library due to potential breaking changes or deprecated functions. \
The solution are to embed a notice system in the template library to warn the end-user of a deprecation and let x versions between the deprecation notice and the replacement / removal of the function.

{% hint style="success" %}
This feature are implemented on v0.6.0, documentation can be found here : \
[Broken link](broken-reference "mention")
{% endhint %}

### :hourglass:**Advanced Error Handling Strategy - **<mark style="color:red;">**DISCONTINUED**</mark>

Implement a custom error handling framework utilising channels for improved error reporting and handling on the Go side, reducing the risk of template crashes.

{% hint style="danger" %}
This feature has been discontinued due to its complexity, which benefits less than 1% of users. If truly necessary, a [custom handler](advanced/how-to-create-a-handler.md) can be implemented to achieve the same functionality.
{% endhint %}

## Compatibility between spring and sprout

{% hint style="success" %}
:white\_check\_mark: All functions present on sprig v3.2.3 are available without breaking changes on sprout v1. Only deprecation warning due to following convention are present.

***

This page will be updated each time function are re-implemented correctly in Sprout v1
{% endhint %}

## Functions added to Sprout v1

A list of functions wanted for v1 based on issues, pull requests from sprig, feedback on sprout. All functions listed here will be implemented for the v1.

<table><thead><tr><th width="94" data-type="checkbox">DONE</th><th>Functions</th><th>Description</th></tr></thead><tbody><tr><td>true</td><td><code>toYaml</code></td><td>Convert a struct to a YAML String</td></tr><tr><td>true</td><td><code>fromYaml</code></td><td>Convert YAML String to a struct</td></tr><tr><td>true</td><td><code>toBool</code></td><td>Convert any to a boolean</td></tr><tr><td>true</td><td><code>toDuration</code></td><td>Convert any to a <code>time.Duration</code></td></tr><tr><td>true</td><td><code>default</code>,<code>empty</code>,<code>coalesce</code></td><td>Dont trigger default go value as false</td></tr><tr><td>true</td><td><code>dig</code></td><td>Dig into a map without crashes in format <code>book.author.name</code></td></tr><tr><td>true</td><td><code>sha512sum</code></td><td>Support of SHA512</td></tr><tr><td>true</td><td><code>md5sum</code></td><td>Support of md5 hash</td></tr><tr><td>true</td><td><code>hasField</code></td><td>Detect if a field are present in an object using reflect. <a href="https://github.com/Masterminds/sprig/issues/401">Source</a></td></tr><tr><td>true</td><td><code>toDuration</code></td><td>convert a value to a <code>time.Duration</code></td></tr><tr><td>true</td><td><code>toCamelCase</code>, <code>toPascalCase</code>, <code>toKebakCase</code>, <code>toDotCase</code>, <code>topathCase</code>, <code>toConstantCase</code>,<code>toSnakeCase</code>,<code>toTitleCase</code></td><td>A batch of functions to change casing of a string to aby casing you want.</td></tr><tr><td>true</td><td><code>capitalize</code>, <code>uncapitalize</code></td><td>Capitalize / Uncapitalize a string (Upper/lower only the first character)</td></tr><tr><td>true</td><td><code>flatten</code></td><td>Flatten nested list be one level</td></tr><tr><td>true</td><td><code>regexpFindSubmatch</code>, <code>regexpAllSubmatches</code>,<code>regexpFindNamedSubmatch</code>, <code>regexpAllNamedSubmatches</code></td><td>Collection of function to found and retrieve submatches and named submatches</td></tr><tr><td>false</td><td><code>cidrhost</code>,<code>cidrnetmask</code>,<code>cidrsubnet</code>,<code>cidrsubnets</code></td><td>A collection of functions for network ip manipulation</td></tr><tr><td>true</td><td><code>toLocalDate</code></td><td>Convert to a <code>time.Time</code> with a timezone support</td></tr></tbody></table>

