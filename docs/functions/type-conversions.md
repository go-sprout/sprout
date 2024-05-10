---
description: Utility functions are used to convert one type to another in your templates.
---

# Type Conversions

### toBool

toBool converts a value from any types reasonably be converted to a boolean. _Using the_ [_cast_ ](https://github.com/spf13/cast)_package._

<table data-header-hidden><thead><tr><th width="193">Name</th><th>Value</th></tr></thead><tbody><tr><td>Group</td><td><code>conversion</code></td></tr><tr><td>Signature</td><td><pre class="language-go"><code class="lang-go">toBool(v any) bool
</code></pre></td></tr><tr><td>Must version</td><td><span data-gb-custom-inline data-tag="emoji" data-code="274c">❌</span></td></tr></tbody></table>

{% tabs %}
{% tab title="Template Example" %}
```go
{{ "true" | toBool }} // Output: true
{{ "t" | toBool }} // Output: true
{{ 1 | toBool }} // Output: true
{{ 0.0 | toBool }} // Output: false
{{ "invalid" | toBool }} // Output: false
```
{% endtab %}
{% endtabs %}

### toInt

toInt converts a value into a int. _Using the_ [_cast_ ](https://github.com/spf13/cast)_package._

<table data-header-hidden><thead><tr><th width="193">Name</th><th>Value</th></tr></thead><tbody><tr><td>Group</td><td><code>conversion</code></td></tr><tr><td>Signature</td><td><pre class="language-go"><code class="lang-go">toInt(v any) int
</code></pre></td></tr><tr><td>Must version</td><td><span data-gb-custom-inline data-tag="emoji" data-code="274c">❌</span></td></tr></tbody></table>

{% tabs %}
{% tab title="Template Example" %}
```go
{{ "1" | toInt }} // Output: 1
{{ 1.1 | toInt }} // Output: 1
{{ true | toInt }} // Output: 1
{{ "invalid" | toInt }} // Output: 0
```
{% endtab %}
{% endtabs %}

### toInt64

toInt64 converts a value into a int64. _Using the_ [_cast_ ](https://github.com/spf13/cast)_package._

<table data-header-hidden><thead><tr><th width="193">Name</th><th>Value</th></tr></thead><tbody><tr><td>Group</td><td><code>conversion</code></td></tr><tr><td>Signature</td><td><pre class="language-go"><code class="lang-go">toInt64(v any) int64
</code></pre></td></tr><tr><td>Must version</td><td><span data-gb-custom-inline data-tag="emoji" data-code="274c">❌</span></td></tr></tbody></table>

{% tabs %}
{% tab title="Template Example" %}
```go
{{ "1" | toInt }} // Output: 1
{{ 1.1 | toInt }} // Output: 1
{{ true | toInt }} // Output: 1
{{ "invalid" | toInt }} // Output: 0
```
{% endtab %}
{% endtabs %}

### toUint

toUint converts a value into a uint. Utilizes the [cast](https://github.com/spf13/cast) package for conversion.

<table data-header-hidden><thead><tr><th width="193">Name</th><th>Value</th></tr></thead><tbody><tr><td>Group</td><td><code>conversion</code></td></tr><tr><td>Signature</td><td><pre class="language-go"><code class="lang-go">toUint(v any) uint
</code></pre></td></tr><tr><td>Must version</td><td><span data-gb-custom-inline data-tag="emoji" data-code="274c">❌</span></td></tr></tbody></table>

{% tabs %}
{% tab title="Template Example" %}
```go
{{ "1" | toUint }} // Output: 1
{{ 1.1 | toUint }} // Output: 1
{{ true | toUint }} // Output: 1
{{ "invalid" | toUint }} // Output: 0
```
{% endtab %}
{% endtabs %}

### toUint64

toUint64 converts a value into a uint64. Utilizes the [cast](https://github.com/spf13/cast) package for conversion.

<table data-header-hidden><thead><tr><th width="193">Name</th><th>Value</th></tr></thead><tbody><tr><td>Group</td><td><code>conversion</code></td></tr><tr><td>Signature</td><td><pre class="language-go"><code class="lang-go">toUint64(v any) uint64
</code></pre></td></tr><tr><td>Must version</td><td><span data-gb-custom-inline data-tag="emoji" data-code="274c">❌</span></td></tr></tbody></table>

{% tabs %}
{% tab title="Template Example" %}
<pre class="language-go"><code class="lang-go">{{ "1" | toUint64 }} // Output: 1
<strong>{{ 1.1 | toUint64 }} // Output: 1
</strong>{{ true | toUint64 }} // Output: 1
{{ "invalid" | toUint64 }} // Output: 0
</code></pre>
{% endtab %}
{% endtabs %}

### toFloat64

toFloat64 converts a value into a float64. Utilizes the [cast](https://github.com/spf13/cast) package for conversion.

<table data-header-hidden><thead><tr><th width="193">Name</th><th>Value</th></tr></thead><tbody><tr><td>Group</td><td><code>conversion</code></td></tr><tr><td>Signature</td><td><pre class="language-go"><code class="lang-go">toFloat64(v any) float64
</code></pre></td></tr><tr><td>Must version</td><td><span data-gb-custom-inline data-tag="emoji" data-code="274c">❌</span></td></tr></tbody></table>

{% tabs %}
{% tab title="Template Example" %}
```go
{{ "1" | toFloat64 }} // Output: 1
{{ 1.42 | toFloat64 }} // Output: 1.42
{{ true | toFloat64 }} // Output: 1
{{ "invalid" | toFloat64 }} // Output: 0
```
{% endtab %}
{% endtabs %}

### toOctal

toOctal parses a value as an octal (base 8) integer.

<table data-header-hidden><thead><tr><th width="193">Name</th><th>Value</th></tr></thead><tbody><tr><td>Group</td><td><code>conversion</code></td></tr><tr><td>Signature</td><td><pre class="language-go"><code class="lang-go">toOctal(v any) int64
</code></pre></td></tr><tr><td>Must version</td><td><span data-gb-custom-inline data-tag="emoji" data-code="274c">❌</span></td></tr></tbody></table>

{% tabs %}
{% tab title="Template Example" %}
```go
{{ 777 | toOctal }} // Output: "511"
{{ "770" | toOctal }} // Output: "504"
{{ true | toOctal }} // Output: "1"
{{ "invalid" | toOctal }} // Output: "0"
```
{% endtab %}
{% endtabs %}

### toString

toString converts a value to a string, handling various types effectively.

<table data-header-hidden><thead><tr><th width="193">Name</th><th>Value</th></tr></thead><tbody><tr><td>Group</td><td><code>conversion</code></td></tr><tr><td>Signature</td><td><pre class="language-go"><code class="lang-go">toString(v any) string
</code></pre></td></tr><tr><td>Must version</td><td><span data-gb-custom-inline data-tag="emoji" data-code="274c">❌</span></td></tr></tbody></table>

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

* `error` and output `err.Error()`&#x20;
* `fmt.Stringer` and output `o.String()`
{% endhint %}

### toDate / toMustDate

toDate converts a string to a `time.Time` object based on a format specification.

<table data-header-hidden><thead><tr><th width="193">Name</th><th>Value</th></tr></thead><tbody><tr><td>Group</td><td><code>conversion</code></td></tr><tr><td>Signature</td><td><pre class="language-go"><code class="lang-go">toDate(layout, value string) time.Time
toMustDate(layout string, value string) (time.Time, error)
</code></pre></td></tr><tr><td>Must version</td><td><span data-gb-custom-inline data-tag="emoji" data-code="2705">✅</span></td></tr></tbody></table>

{% tabs %}
{% tab title="Template Example" %}
```go
{{ toDate "2006-01-02", "2024-05-10 11:12:42" }}
// Output: 2024-05-10 00:00:00 +0000 UTC
```

_This example will takes the_ `"2024-05-10 11:12:42"` _string and convert it with the layout_ `"2006-01-02"`.
{% endtab %}

{% tab title="Must version" %}
```go
{{ toMustDate "2006-01-02", "2024-05-10 11:12:42" }}
// Output: 2024-05-10 00:00:00 +0000 UTC, nil
```

_This example will takes the_ `"2024-05-10 11:12:42"` _string and convert it with the layout_ `"2006-01-02"`.&#x20;

:heavy\_check\_mark: Native Go Template error handling. _In case of error, the template rendering stop._
{% endtab %}
{% endtabs %}

{% hint style="info" %}
See more about Golang Layout on the [official documentation](https://go.dev/src/time/format.go).
{% endhint %}

### toDuration

toDuration converts a value to a `time.Duration`. Taking a possibly signed sequence of decimal numbers, each optional fraction and a unit suffix, such `300ms`, `-1.5h` or `2h45m`.

Valid time units are `ns`, `us` (or `µs`), `ms`, `s`, `m` and `h`.

<table data-header-hidden><thead><tr><th width="193">Name</th><th>Value</th></tr></thead><tbody><tr><td>Group</td><td><code>conversion</code></td></tr><tr><td>Signature</td><td><pre class="language-go"><code class="lang-go">toDuration(v any) time.Duration
</code></pre></td></tr><tr><td>Must version</td><td><span data-gb-custom-inline data-tag="emoji" data-code="274c">❌</span></td></tr></tbody></table>

{% tabs %}
{% tab title="Template Example" %}
```go
{{ 1 | toDuration }} // Output: 1ns
{{ (1000.0 * 1000.0) | toDuration }} // Output: 1ms
{{ "1m" | toDuration }} // Output: 1m
{{ "invalid" | toDuration }} // Output: 0s
```
{% endtab %}
{% endtabs %}

## Deprecated functions

### atoi :warning:

{% hint style="warning" %}
**\[DEPRECATED]** Use [`toInt`](type-conversions.md#toint)instead.
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
**\[DEPRECATED]** Use [`toInt`](type-conversions.md#toint)instead.
{% endhint %}

{% tabs %}
{% tab title="Template Example" %}
<pre class="language-go"><code class="lang-go"><strong>`{{ "42" | int }} // Output: 42
</strong></code></pre>
{% endtab %}
{% endtabs %}

### int64 :warning:

{% hint style="warning" %}
**\[DEPRECATED]** Use [`toInt64`](type-conversions.md#toint64)instead.
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
**\[DEPRECATED]** Use [`toFloat64`](type-conversions.md#tofloat64)instead.
{% endhint %}

{% tabs %}
{% tab title="Template Example" %}
<pre class="language-go"><code class="lang-go"><strong>`{{ "42.42" | float64 }} // Output: 42.42
</strong></code></pre>
{% endtab %}
{% endtabs %}
