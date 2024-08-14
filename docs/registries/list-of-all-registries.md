---
icon: list-radio
---

# List of all registries

Every function is categorized into a registry and may include a [**'must'** version](list-of-all-registries.md#must-version). This 'must' version utilizes the native error handling of Go templates to manage errors that occur within your method.

### List of registries

* [**backward**](backward.md): Functions to maintain backward compatibility with sprig.
* [**checksum**](checksum.md): Tools to generate and verify checksums for data integrity.
* [**conversion**](conversion.md): Functions to convert between different data types within templates.
* [**crypto**](crypto.md): Cryptographic utilities for encryption, hashing, and security.
* [**encoding**](encoding.md): Methods for encoding and decoding data in various formats.
* [**env**](env.md): Access and manipulate environment variables within templates.
* [**filesystem**](filesystem.md): Functions for interacting with the file system.
* [**maps**](maps.md): Tools to manipulate and interact with map data structures.
* [**numeric**](numeric.md): Utilities for numerical operations and calculations.
* [**random**](random.md): Functions to generate random numbers, strings, and other data.
* [**reflect**](reflect.md): Tools to inspect and manipulate data types using reflection.
* [**regexp**](regexp.md): Regular expression functions for pattern matching and string manipulation.
* [**semver**](semver.md): Functions to handle semantic versioning and comparison.
* [**slices**](slices.md): Utilities for slice operations, including filtering, sorting, and transforming.
* [**std**](std.md): Standard functions for common operations.
* [**strings**](strings.md): Functions for string manipulation, including formatting, splitting, and joining.
* [**time**](time.md): Tools to handle dates, times, and time-related calculations.
* [**uniqueid**](uniqueid.md): Functions to generate unique identifiers, such as UUIDs.

### Must version

{% hint style="warning" %}
<mark style="color:yellow;">The</mark> <mark style="color:yellow;"></mark><mark style="color:yellow;">`Must`</mark> <mark style="color:yellow;"></mark><mark style="color:yellow;">strategy is currently under discussion and may be subject to change in the future. An RFC is currently open for feedback and discussion. You can view and participate in the RFC</mark> [here](https://github.com/orgs/go-sprout/discussions/32)<mark style="color:yellow;">.</mark>
{% endhint %}

The **Must** version of each function is essentially a safer variant that ensures error handling is integrated into the function's execution. When you use a function prefixed by `Must` (e.g., `MustEncode`, `MustConvert`), the template engine automatically checks for and handles any errors that might occur during the function's execution.&#x20;

This is particularly useful in scenarios where failing silently is not an option, and you need immediate feedback if something goes wrong. Utilizing the **Must** versions helps to maintain the robustness and reliability of your template code.
