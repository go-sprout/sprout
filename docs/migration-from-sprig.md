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

Choose your migration path based on your situation:

### Option A: New Projects or Breaking Changes Acceptable

If you're starting a new project or your users accept breaking changes, use Sprout directly:

```go
import "github.com/go-sprout/sprout"

handler := sprout.New()
funcs := handler.Build()
```

This gives you all the improvements (bug fixes, proper piping support, better error handling) without any compatibility overhead.

### Option B: Non-Breaking Migration for End-Users

If you have existing end-users and need a gradual migration path:

**Phase 1: Replace Sprig with Sprigin**

```go
// Before (Sprig)
import "github.com/Masterminds/sprig/v3"
funcs := sprig.FuncMap()

// After (Sprigin)
import "github.com/go-sprout/sprout/sprigin"
funcs := sprigin.FuncMap()
```

Sprigin provides full backward compatibility while logging deprecation warnings. Your end-users will see warnings in logs about deprecated functions and signature changes, giving them time to update their templates.

**Customizing the Logger**

By default, Sprigin logs warnings to the standard `slog` default handler. You can provide your own logger to integrate with your application's logging system:

```go
import (
    "log/slog"
    "os"

    "github.com/go-sprout/sprout/sprigin"
)

// Use a custom logger for deprecation warnings
logger := slog.New(slog.NewJSONHandler(os.Stderr, &slog.HandlerOptions{
    Level: slog.LevelWarn,
}))
funcs := sprigin.FuncMap(sprigin.WithLogger(logger))
```

This is useful when you want to:
- Route deprecation warnings to a specific log destination
- Use a structured logging format (JSON, etc.)
- Filter or aggregate warnings in your logging infrastructure

**Phase 2: Keep Sprigin for X Versions/Months**

Maintain sprigin for your defined deprecation period (e.g., 3-6 months or 2-3 versions) to:
- Respect your breaking change policy
- Allow end-users to see and act on deprecation warnings
- Ensure a smooth transition

**Phase 3: Switch to Sprout**

Once your deprecation period ends, replace sprigin with sprout:

```go
// Final migration
import "github.com/go-sprout/sprout"

handler := sprout.New()
funcs := handler.Build()
```

## <mark style="color:purple;">How to Transition for your end-users</mark>

You use Sprig or Sprout for end-users and want to migrate? Here's a detailed plan to ensure confidence during the migration:

{% hint style="info" %}
Need more information? Contact maintainers or open [a discussion on the repository](https://github.com/orgs/go-sprout/discussions/categories/q-a).
{% endhint %}

***

### The Sprigin Compatibility Layer

The `sprigin` package is specifically designed for this use case. It:

- **Supports both signatures**: Automatically detects whether you're using the old Sprig signature (`get $dict "key"`) or the new Sprout piping signature (`$dict | get "key"`)
- **Logs deprecation warnings**: When old signatures or deprecated functions are used, warnings are logged to inform your end-users (see [Customizing the Logger](#customizing-the-logger) to integrate with your logging system)
- **Preserves Sprig behavior**: Bug-for-bug compatible to avoid breaking existing templates

### Migration Steps

1. **Communicate the Purpose of the Migration**\
   Explain the reasons for switching to Sprout, emphasizing improvements such as better performance, modular function registries, enhanced error handling, and new features like function notices and safe functions.

2. **Replace Sprig with Sprigin**\
   Simply replace your import and function call:
   ```go
   // Before
   import "github.com/Masterminds/sprig/v3"
   funcs := sprig.FuncMap()

   // After
   import "github.com/go-sprout/sprout/sprigin"
   funcs := sprigin.FuncMap()
   ```

3. **Monitor Deprecation Warnings**\
   Sprigin logs warnings for deprecated functions and signature changes. Share these logs with your end-users so they can update their templates. Use `sprigin.WithLogger()` to route warnings to your preferred logging destination.

4. **Keep Sprigin for X Versions/Months**\
   Maintain the compatibility layer according to your breaking change policy. We recommend 3-6 months or 2-3 versions.

5. **Final Switch to Sprout**\
   Once your deprecation period ends and end-users have migrated their templates, switch to Sprout directly.

***

{% hint style="info" %}
As a library developer, you can extend Sprout by creating your [own function registry](advanced/how-to-create-a-registry.md). Additionally, you can use [**notices**](features/function-notices.md) to inform your end-users about important updates during template execution.
{% endhint %}

{% hint style="success" %}
Our maintainers and collaborators can assist you if you have questions. Don't hesitate to [open a discussion on GitHub](https://github.com/orgs/go-sprout/discussions/categories/q-a)!
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

In Sprig, errors within certain functions cause a panic. Both **Sprout and Sprigin** fix all panics, returning errors properly instead.

**Old Behavior (Sprig)**: Triggers a panic on error

```go
if err != nil {
  panic("deepCopy error: " + err.Error())
}
```

**New Behavior (Sprout & Sprigin)**: Returns nil or an empty value on error

```go
if err != nil {
  return nil, err
}
```

{% hint style="success" %}
**Migration Tip**

Whether you use `sprout` or `sprigin`, all panics are fixed. You can safely migrate without worrying about panics crashing your application.
{% endhint %}

#### Methods that previously caused a panic in Sprig:

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

### Signature Changes (Argument Order)

Sprout reorders function arguments to support Go template piping conventions. The target (map/list) is now the **last** argument instead of the first.

{% hint style="info" %}
If you use `sprigin.FuncMap()`, both signatures are supported automatically. Sprigin detects which signature you're using and logs a warning when the old Sprig signature is detected.
{% endhint %}

#### Map Functions: get, set, unset, hasKey, pick, omit

| Function | Sprig Signature | Sprout Signature |
|----------|-----------------|------------------|
| `get` | `{{ get $dict "key" }}` | `{{ $dict \| get "key" }}` |
| `set` | `{{ set $dict "key" "value" }}` | `{{ $dict \| set "key" "value" }}` |
| `unset` | `{{ unset $dict "key" }}` | `{{ $dict \| unset "key" }}` |
| `hasKey` | `{{ hasKey $dict "key" }}` | `{{ $dict \| hasKey "key" }}` |
| `pick` | `{{ pick $dict "k1" "k2" }}` | `{{ $dict \| pick "k1" "k2" }}` |
| `omit` | `{{ omit $dict "k1" "k2" }}` | `{{ $dict \| omit "k1" "k2" }}` |

#### List Functions: append, prepend, slice, without

| Function | Sprig Signature | Sprout Signature |
|----------|-----------------|------------------|
| `append` | `{{ append $list "value" }}` | `{{ $list \| append "value" }}` |
| `prepend` | `{{ prepend $list "value" }}` | `{{ $list \| prepend "value" }}` |
| `slice` | `{{ slice $list 1 3 }}` | `{{ $list \| slice 1 3 }}` |
| `without` | `{{ without $list "a" "b" }}` | `{{ $list \| without "a" "b" }}` |

#### Dig Function

* **Sprig**: `{{ dig "key1" "key2" "default" $dict }}` - default value is the second-to-last argument
* **Sprout**: `{{ $dict | dig "key1" "key2" | default "default" }}` - use `default` filter separately

Additional differences:
* **Sprig**: Dots in keys are treated literally (`dig "a.b"` looks for key `"a.b"`)
* **Sprout**: Keys are split on dots (`dig "a.b"` is equivalent to `dig "a" "b"`)

---

### Behavior Changes

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

#### Date

* **Sprig**: Uses local timezone for formatting.
* **Sprout**: Uses the timezone from the time value itself.

{% hint style="info" %}
If you use `sprigin.FuncMap()`, the `date` function uses local timezone like Sprig for backward compatibility.
{% endhint %}

#### Base32Decode / Base64Decode

* **Sprig**: Returns the error message string when decoding fails.
* **Sprout**: Returns an empty string when decoding fails.

{% hint style="info" %}
If you use `sprigin.FuncMap()`, the decoding functions return the error message like Sprig for backward compatibility.
{% endhint %}

#### ToCamelCase / ToPascalCase

* **Sprig**: The `camelcase` function returns PascalCase (e.g., `"hello_world"` → `"HelloWorld"`). No true camelCase function exists.
* **Sprout**:
  * `toCamelCase` returns true camelCase (e.g., `"hello_world"` → `"helloWorld"`)
  * `toPascalCase` returns PascalCase (e.g., `"hello_world"` → `"HelloWorld"`)
  * The alias `camelcase` points to `toPascalCase` for backward compatibility.

#### ToTitleCase / Title

* **Sprig**: Uses the deprecated `strings.Title` function which:
  * Doesn't lowercase letters before capitalizing (`"HELLO"` stays `"HELLO"`)
  * Treats apostrophes as word separators (`"it's"` becomes `"It'S"`)
* **Sprout**: Uses proper Unicode title casing (`"HELLO"` → `"Hello"`, `"it's"` → `"It's"`)

{% hint style="info" %}
If you use `sprigin.FuncMap()`, the `toTitleCase`/`title` function uses the old `strings.Title` behavior with a warning about the upcoming change.
{% endhint %}

#### Substr

* **Sprig**: `substr 0 0 "hello"` returns `""` (empty string)
* **Sprout**: `substr 0 0 "hello"` returns `"hello"` (full string, as end=0 means "to the end")

{% hint style="info" %}
If you use `sprigin.FuncMap()`, the old behavior is preserved with warnings about the upcoming change.
{% endhint %}

#### KindOf

* **Sprig**: `kindOf nil` returns `"invalid"`
* **Sprout**: `kindOf nil` returns an error

{% hint style="info" %}
If you use `sprigin.FuncMap()`, the old behavior is preserved with a warning about the upcoming change.
{% endhint %}

#### EllipsisBoth / Abbrevboth

* **Sprig**: Has a bug where offset <= 4 is ignored (truncates from right only)
* **Sprout**: Respects the offset parameter correctly

{% hint style="info" %}
If you use `sprigin.FuncMap()`, the buggy behavior is preserved with a warning when offset <= 4.
{% endhint %}

#### Nospace / Ellipsis (Abbrev)

* **Sprig**: Has bugs with Unicode characters (e.g., `nospace "α β γ"`)
* **Sprout**: Correctly handles Unicode characters

#### Merge / MergeOverwrite

* **Sprig**: Dereferences when the second value is the default Go value (e.g., `0` for int is treated as "not set").
* **Sprout**: Does not dereference; keeps the second value as-is (e.g., `0` for int is preserved).

## <mark style="color:purple;">Deprecated Features</mark>

### Functions to be Removed

The following functions are marked with `// ! DEPRECATED` and will be removed in the next major version:

| Function | Reason | Alternative |
|----------|--------|-------------|
| `fail` | Limited use case | Use Go error handling |
| `urlParse` | Moving to dedicated URL package | Will be replaced |
| `urlJoin` | Moving to dedicated URL package | Will be replaced |
| `getHostByName` | Non-deterministic, security concern | Handle DNS outside templates |

{% hint style="warning" %}
Perform cryptographic operations (listed in `crypto` package) outside of templates. The [`crypto` registry](registries/crypto.md) will be dropped in future versions.
{% endhint %}

### Renamed Functions (Deprecated Aliases)

The following function names are deprecated. They still work but will log deprecation warnings. Use the new names instead:

#### String Functions

| Deprecated | New Name |
|------------|----------|
| `upper`, `toupper`, `uppercase` | `toUpper` |
| `lower`, `tolower`, `lowercase` | `toLower` |
| `title`, `titlecase` | `toTitleCase` |
| `camelcase` | `toPascalCase` |
| `snake`, `snakecase` | `toSnakeCase` |
| `kebab`, `kebabcase` | `toKebabCase` |
| `swapcase` | `swapCase` |
| `abbrev` | `ellipsis` |
| `abbrevboth` | `ellipsisBoth` |
| `trimall` | `trimAll` |

#### List Functions

| Deprecated | New Name |
|------------|----------|
| `push` | `append` |
| `mustPush` | `mustAppend` |
| `tuple` | `list` |
| `biggest` | `max` |

#### Encoding Functions

| Deprecated | New Name |
|------------|----------|
| `b64enc` | `base64Encode` |
| `b64dec` | `base64Decode` |
| `b32enc` | `base32Encode` |
| `b32dec` | `base32Decode` |

#### Path Functions

| Deprecated | New Name |
|------------|----------|
| `base` | `pathBase` |
| `dir` | `pathDir` |
| `ext` | `pathExt` |
| `clean` | `pathClean` |
| `isAbs` | `pathIsAbs` |

#### Date Functions

| Deprecated | New Name |
|------------|----------|
| `date_modify` | `dateModify` |
| `date_in_zone` | `dateInZone` |
| `must_date_modify` | `mustDateModify` |
| `ago` | `dateAgo` |

#### Type Conversion Functions

| Deprecated | New Name |
|------------|----------|
| `int`, `atoi` | `toInt` |
| `int64` | `toInt64` |
| `float64` | `toFloat64` |
| `toDecimal` | `toOctal` |
| `toStrings` | `strSlice` |

#### Math Functions

| Deprecated | New Name |
|------------|----------|
| `addf` | `add` |
| `add1f` | `add1` |
| `subf` | `sub` |

#### Other Functions

| Deprecated | New Name |
|------------|----------|
| `expandenv` | `expandEnv` |

{% hint style="info" %}
**Migration Tip**

All deprecated aliases are flagged with `// ! deprecated` (lowercase) in sprigin for renamed functions.\
Functions marked for total removal use `// ! DEPRECATED` (uppercase) in the codebase.
{% endhint %}

## <mark style="color:purple;">Conclusion</mark>

Migrating from Sprig to Sprout offers significant benefits, including improved error handling, modular function management, and enhanced compatibility with modern Go practices. While the `sprigin` package provides a bridge for backward compatibility, fully embracing Sprout’s native capabilities will lead to a more stable and maintainable codebase. For further details on Sprout’s features and API, consult the [official Sprout documentation](https://docs.atom.codes/sprout).
