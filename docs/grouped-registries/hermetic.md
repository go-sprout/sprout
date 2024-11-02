---
description: >-
  The Hermetic registry group includes all the registries available in Sprout,
  excluding registries that depend on external services or are influenced by the
  environment where the application is running.
---

# All

{% hint style="info" %}
You can easily import group from the <mark style="color:yellow;">`all`</mark> group by including the following import statement in your code

```go
import "github.com/go-sprout/sprout/group/all"
```
{% endhint %}

### List of registries

* [**checksum**](checksum.md): Tools to generate and verify checksums for data integrity.
* [**conversion**](conversion.md): Functions to convert between different data types within templates.
* [**encoding**](encoding.md): Methods for encoding and decoding data in various formats.
* [**filesystem**](filesystem.md): Functions for interacting with the file system.
* [**maps**](maps.md): Tools to manipulate and interact with map data structures.
* [**numeric**](numeric.md): Utilities for numerical operations and calculations.
* [**reflect**](reflect.md): Tools to inspect and manipulate data types using reflection.
* [**regexp**](regexp.md): Regular expression functions for pattern matching and string manipulation.
* [**semver**](semver.md): Functions to handle semantic versioning and comparison.
* [**slices**](slices.md): Utilities for slice operations, including filtering, sorting, and transforming.
* [**std**](std.md): Standard functions for common operations.
* [**strings**](strings.md): Functions for string manipulation, including formatting, splitting, and joining.
* [**time**](time.md): Tools to handle dates, times, and time-related calculations.
* [**uniqueid**](uniqueid.md): Functions to generate unique identifiers, such as UUIDs.
