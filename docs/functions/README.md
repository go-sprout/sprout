---
description: A documented page for all functions usable in the template with examples
---

# 💻 Functions

{% hint style="info" %}
This page is currently under construction, functions migrated to this page are deleted from individuals pages under the [Old documentation from sprig](../old-documentation-from-sprig/) section. Thanks for your comprehension :pray:
{% endhint %}

Every function is categorized into a group and may include a **'must'** version. This 'must' version utilizes the native error handling of Go templates to manage errors that occur within your method.

### List of groups

* [**Conversions**](../old-documentation-from-sprig/conversion.md): Utility functions are used to convert one type to another in your templates.
* <mark style="color:red;">**Encoding**</mark>: Functions designed to handle the encoding and decoding of data formats.
* <mark style="color:red;">**Filesystem**</mark>: Tools to interact with and manipulate the file system.
* <mark style="color:red;">**Maps**</mark>: Functions to facilitate operations and manipulations on map data structures.
* <mark style="color:red;">**Misc**</mark>: A collection of miscellaneous functions that do not fit into the other categories.
* <mark style="color:red;">**Numeric**</mark>: Functions focused on numeric calculations and conversions
* <mark style="color:red;">**Random**</mark>: Tools to generate random things.
* <mark style="color:red;">**Regexp**</mark>: Functions that provide support for regular expression processing.
* <mark style="color:red;">**Slices**</mark>: Utilities to manage and manipulate slices.
* <mark style="color:red;">**Strings**</mark>: Functions dedicated to string manipulation and analysis.
* <mark style="color:red;">**Time**</mark>: Tools to handle dates, times, and time-related calculations.

### Must version

The **Must** version of each function is essentially a safer variant that ensures error handling is integrated into the function's execution. When you use a function prefixed by `Must` (e.g., `MustEncode`, `MustConvert`), the template engine automatically checks for and handles any errors that might occur during the function's execution.&#x20;

This is particularly useful in scenarios where failing silently is not an option, and you need immediate feedback if something goes wrong. Utilizing the **Must** versions helps to maintain the robustness and reliability of your template code.

