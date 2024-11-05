---
description: >-
  The Conversion registry includes a collection of functions designed to convert
  one data type to another directly within your templates. This allows for
  seamless type transformations.
---

# Conversion

{% hint style="info" %}
You can easily import all the functions from the <mark style="color:yellow;">`conversion`</mark> registry by including the following import statement in your code

```go
import "github.com/go-sprout/sprout/registry/conversion"
```
{% endhint %}

### <mark style="color:purple;">toBool</mark>

toBool converts a value from any types reasonably be converted to a boolean value. _Using the_ [_cast_ ](https://github.com/spf13/cast)_package._

<table data-header-hidden><thead><tr><th width="193">Name</th><th>Value</th></tr></thead><tbody><tr><td>Signature</td><td><pre class="language-go"><code class="lang-go">ToBool(v any) (bool, error)
</code></pre></td></tr></tbody></table>

{% tabs %}
{% tab title="Template Example" %}
```go
{{ "true" | toBool }} // Output: true
{{ "t" | toBool }} // Output: true
{{ 1 | toBool }} // Output: true
{{ 0.0 | toBool }} // Output: false
{{ "invalid" | toBool }} // Error
```
{% endtab %}
{% endtabs %}

### <mark style="color:purple;">toInt</mark>

toInt converts a value into an `int`. _Using the_ [_cast_ ](https://github.com/spf13/cast)_package._

<table data-header-hidden><thead><tr><th width="193">Name</th><th>Value</th></tr></thead><tbody><tr><td>Signature</td><td><pre class="language-go"><code class="lang-go">ToInt(v any) (int, error)
</code></pre></td></tr></tbody></table>

{% tabs %}
{% tab title="Template Example" %}
```go
{{ "1" | toInt }} // Output: 1
{{ 1.1 | toInt }} // Output: 1
{{ true | toInt }} // Output: 1
{{ "invalid" | toInt }} // Error
```
{% endtab %}
{% endtabs %}

### <mark style="color:purple;">toInt64</mark>

toInt64 converts a value into an `int64`. _Using the_ [_cast_ ](https://github.com/spf13/cast)_package._

<table data-header-hidden><thead><tr><th width="193">Name</th><th>Value</th></tr></thead><tbody><tr><td>Signature</td><td><pre class="language-go"><code class="lang-go">ToInt64(v any) (int64, error)
</code></pre></td></tr></tbody></table>

{% tabs %}
{% tab title="Template Example" %}
```go
{{ "1" | toInt }} // Output: 1
{{ 1.1 | toInt }} // Output: 1
{{ true | toInt }} // Output: 1
{{ "invalid" | toInt }} // Error
```
{% endtab %}
{% endtabs %}

### <mark style="color:purple;">toUint</mark>

toUint converts a value into a `uint`. Utilizes the [cast](https://github.com/spf13/cast) package for conversion.

<table data-header-hidden><thead><tr><th width="193">Name</th><th>Value</th></tr></thead><tbody><tr><td>Signature</td><td><pre class="language-go"><code class="lang-go">ToUint(v any) (uint, error)
</code></pre></td></tr></tbody></table>

{% tabs %}
{% tab title="Template Example" %}
```go
{{ "1" | toUint }} // Output: 1
{{ 1.1 | toUint }} // Output: 1
{{ true | toUint }} // Output: 1
{{ "invalid" | toUint }} // Error
```
{% endtab %}
{% endtabs %}

### <mark style="color:purple;">toUint64</mark>

toUint64 converts a value into a `uint64`. Utilizes the [cast](https://github.com/spf13/cast) package for conversion.

<table data-header-hidden><thead><tr><th width="193">Name</th><th>Value</th></tr></thead><tbody><tr><td>Signature</td><td><pre class="language-go"><code class="lang-go">ToUint64(v any) (uint64, error)
</code></pre></td></tr></tbody></table>

{% tabs %}
{% tab title="Template Example" %}
<pre class="language-go"><code class="lang-go">{{ "1" | toUint64 }} // Output: 1
<strong>{{ 1.1 | toUint64 }} // Output: 1
</strong>{{ true | toUint64 }} // Output: 1
{{ "invalid" | toUint64 }} // Error
</code></pre>
{% endtab %}
{% endtabs %}

### <mark style="color:purple;">toFloat64</mark>

toFloat64 converts a value into a float64. Utilizes the [cast](https://github.com/spf13/cast) package for conversion.

<table data-header-hidden><thead><tr><th width="193">Name</th><th>Value</th></tr></thead><tbody><tr><td>Signature</td><td><pre class="language-go"><code class="lang-go">ToFloat64(v any) (float64, error)
</code></pre></td></tr></tbody></table>

{% tabs %}
{% tab title="Template Example" %}
```go
{{ "1" | toFloat64 }} // Output: 1
{{ 1.42 | toFloat64 }} // Output: 1.42
{{ true | toFloat64 }} // Output: 1
{{ "invalid" | toFloat64 }} // Error
```
{% endtab %}
{% endtabs %}

### <mark style="color:purple;">toOctal</mark>

toOctal parses a value as an octal (base 8) integer.

<table data-header-hidden><thead><tr><th width="193">Name</th><th>Value</th></tr></thead><tbody><tr><td>Signature</td><td><pre class="language-go"><code class="lang-go">ToOctal(v any) (int64, error)
</code></pre></td></tr></tbody></table>

{% tabs %}
{% tab title="Template Example" %}
```go
{{ 777 | toOctal }} // Output: "511"
{{ "770" | toOctal }} // Output: "504"
{{ true | toOctal }} // Output: "1"
{{ "invalid" | toOctal }} // Error
```
{% endtab %}
{% endtabs %}

### <mark style="color:purple;">toString</mark>

toString converts a value to a string, handling various types effectively.

<table data-header-hidden><thead><tr><th width="193">Name</th><th>Value</th></tr></thead><tbody><tr><td>Signature</td><td><pre class="language-go"><code class="lang-go">ToString(v any) string
</code></pre></td></tr></tbody></table>

{% tabs %}
{% tab title="Template Example" %}
```go
{{ 1 | toString }} // Output: "1"
{{ 1.42 | toString }} // Output: "1.42"
{{ true | toString }} // Output: "true"
{{ nil | toString }} // Output: "<nil>"
```
{% endtab %}
{% endtabs %}

{% hint style="info" %}
**Note**: toString can handle various types as:

* `error` and output `err.Error()`
* `fmt.Stringer` and output `o.String()`
{% endhint %}

### <mark style="color:purple;">toDate</mark>

toDate converts a string to a `time.Time` object based on a format specification.

<table data-header-hidden><thead><tr><th width="162">Name</th><th>Value</th></tr></thead><tbody><tr><td>Signature</td><td><pre class="language-go"><code class="lang-go">ToDate(layout string, value string) (time.Time, error)
</code></pre></td></tr></tbody></table>

{% tabs %}
{% tab title="Template Example" %}
```go
{{ toDate "2006-01-02", "2024-05-10 11:12:42" }}
// Output: 2024-05-10 00:00:00 +0000 UTC, nil
```

_This example will takes the_ `"2024-05-10 11:12:42"` _string and convert it with the layout_ `"2006-01-02"`.
{% endtab %}
{% endtabs %}

{% hint style="info" %}
See more about Golang Layout on the [official documentation](https://go.dev/src/time/format.go).
{% endhint %}

### <mark style="color:purple;">toLocalDate</mark>

toLocalDate converts a string to a time.Time object based on a format specification and the local timezone.

<table data-header-hidden><thead><tr><th width="162">Name</th><th>Value</th></tr></thead><tbody><tr><td>Signature</td><td><pre class="language-go"><code class="lang-go">ToLocalDate(fmt, timezone, str string) (time.Time, error)
</code></pre></td></tr></tbody></table>

{% tabs %}
{% tab title="Template Example" %}
```go
{{ "2024-09-17 11:12:42" | toLocalDate "2006-01-02" "Europe/Paris" }}
// Output: 2024-09-17 00:00:00 +0200 CEST, nil
{{ "2024-09-17 11:12:42" | toLocalDate "2006-01-02" "MST" }}
// Output: 2024-09-17 00:00:00 -0700 MST, nil
{{ "2024-09-17 11:12:42" | toLocalDate "2006-01-02" "invalid" }}
// Error
```
{% endtab %}
{% endtabs %}

{% hint style="info" %}
See more about Golang Layout on the [official documentation](https://go.dev/src/time/format.go).
{% endhint %}

### <mark style="color:purple;">toDuration</mark>

toDuration converts a value to a `time.Duration`. Taking a possibly signed sequence of decimal numbers, each optional fraction and a unit suffix, such `300ms`, `-1.5h` or `2h45m`.

Valid time units are `ns`, `us` (or `µs`), `ms`, `s`, `m` and `h`.

<table data-header-hidden><thead><tr><th width="193">Name</th><th>Value</th></tr></thead><tbody><tr><td>Signature</td><td><pre class="language-go"><code class="lang-go">ToDuration(v any) (time.Duration, error)
</code></pre></td></tr></tbody></table>

{% tabs %}
{% tab title="Template Example" %}
```go
{{ 1 | toDuration }} // Output: 1ns
{{ (1000.0 * 1000.0) | toDuration }} // Output: 1ms
{{ "1m" | toDuration }} // Output: 1m
{{ "invalid" | toDuration }} // Error
{{ (toDuration "1h30m").Seconds }} // Output: 5400
```
{% endtab %}
{% endtabs %}

## <mark style="color:red;">Deprecated functions</mark>

### atoi :warning:

{% hint style="warning" %}
**\[DEPRECATED]** Use [`toInt`](conversion.md#toint)instead.
{% endhint %}

{% tabs %}
{% tab title="Template Example" %}
<pre class="language-go"><code class="lang-go"><strong>{{ "42" | atoi }} // Output: 42
</strong></code></pre>

:x: No error handling
{% endtab %}
{% endtabs %}

### int :warning:

{% hint style="warning" %}
**\[DEPRECATED]** Use [`toInt`](conversion.md#toint)instead.
{% endhint %}

{% tabs %}
{% tab title="Template Example" %}
<pre class="language-go"><code class="lang-go"><strong>`{{ "42" | int }} // Output: 42
</strong></code></pre>
{% endtab %}
{% endtabs %}

### int64 :warning:

{% hint style="warning" %}
**\[DEPRECATED]** Use [`toInt64`](conversion.md#toint64)instead.
{% endhint %}

{% tabs %}
{% tab title="Template Example" %}
```go
{{ "42" | int64 }} // Output: 42
```
{% endtab %}
{% endtabs %}

### float64 :warning:

{% hint style="warning" %}
**\[DEPRECATED]** Use [`toFloat64`](conversion.md#tofloat64)instead.
{% endhint %}

{% tabs %}
{% tab title="Template Example" %}
<pre class="language-go"><code class="lang-go"><strong>`{{ "42.42" | float64 }} // Output: 42.42
</strong></code></pre>
{% endtab %}
{% endtabs %}
