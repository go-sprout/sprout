---
description: A documented page for all functions usable in the template with examples
---

# 💻 Functions

{% hint style="info" %}
This page is currently under construction, functions migrated to this page are deleted from individuals pages under the [Old documentation from sprig](old-documentation-from-sprig/) section. Thanks for your comprehension :pray:
{% endhint %}

## Conversions functions

### toString

Signature: `toSring(value any) string`

{% tabs %}
{% tab title="Template Example" %}
```go
`{{ 42 | toString }}`
```

_This example will takes the 42 integer and convert it to a string, the output will be `"42"`._

:x: No error handling
{% endtab %}
{% endtabs %}

### toInt

Signature: `toInt(value any) int`

{% tabs %}
{% tab title="Template Example" %}
<pre class="language-go"><code class="lang-go"><strong>`{{ "42"| toInt }}`
</strong></code></pre>

_This example will takes the 42 string and convert it to an int, the output will be `42`._

:x: No error handling
{% endtab %}
{% endtabs %}

### toInt64

Signature: `toInt64(value any) int64`

{% tabs %}
{% tab title="Template Example" %}
<pre class="language-go"><code class="lang-go"><strong>`{{ "42"| toInt64 }}`
</strong></code></pre>

_This example will takes the 42 string and convert it to an int64, the output will be `42`._

:x: No error handling
{% endtab %}
{% endtabs %}

### toFloat64

Signature: `toFloat64(value any) float64`

{% tabs %}
{% tab title="Template Example" %}
<pre class="language-go"><code class="lang-go"><strong>`{{ "42.42"| toFloat64 }}`
</strong></code></pre>

_This example will takes the 42.43 string and convert it to a float64, the output will be `42.42`._

:x: No error handling
{% endtab %}
{% endtabs %}

### toDate

Signature: `toDate(layout string, value string) time.Time`

{% tabs %}
{% tab title="Template Example" %}
```go
`{{ "2024-03-29 11:12:42" | toDate "2006-01-02 15:04:05" }}`
```

_This example will takes the_ `"2024-03-29 11:12:42"` _string and convert it with the layout_ `"2006-01-02 15:04:05"` , the output will be a time.Time object with value `2024-03-29 11:12:42`.

{% hint style="info" %}
See more about Golang Layout on the [official documentation](https://go.dev/src/time/format.go).
{% endhint %}

:x: No error handling
{% endtab %}
{% endtabs %}

### toMustDate

Signature: `toMustDate(layout string, value string) (time.Time, error)`

{% tabs %}
{% tab title="Template Example" %}
```go
`{{ "2024-03-29 11:12:42" | toMustDate "2006-01-02 15:04:05" }}`
```

_This example will takes the_ `"2024-03-29 11:12:42"` _string and convert it with the layout_ `"2006-01-02 15:04:05"` , the output will be a time.Time object with value `2024-03-29 11:12:42`.&#x20;

{% hint style="info" %}
See more about Golang Layout on the [official documentation](https://go.dev/src/time/format.go).
{% endhint %}

:heavy\_check\_mark: Native Go Template error handling. _In case of error, the template rendering stop._toMustDate
{% endtab %}
{% endtabs %}

### atoi :warning:

Signature: `atoi(value any) int`

{% hint style="warning" %}
**\[DEPRECATED]** Use [`toInt`](functions.md#toint)instead.
{% endhint %}

{% tabs %}
{% tab title="Template Example" %}
<pre class="language-go"><code class="lang-go"><strong>`{{ "42"| atoi }}`
</strong></code></pre>

_This example will takes the 42 string and convert it to an int, the output will be `42`._

:x: No error handling
{% endtab %}
{% endtabs %}

### int :warning:

Signature: `int(value any) int`

{% hint style="warning" %}
**\[DEPRECATED]** Use [`toInt`](functions.md#toint)instead.
{% endhint %}

{% tabs %}
{% tab title="Template Example" %}
<pre class="language-go"><code class="lang-go"><strong>`{{ "42"| int }}`
</strong></code></pre>

_This example will takes the 42 string and convert it to an int, the output will be `42`._

:x: No error handling
{% endtab %}
{% endtabs %}

### int64 :warning:

Signature: `int64(value any) int64`

{% hint style="warning" %}
**\[DEPRECATED]** Use [`toInt64`](functions.md#toint64)instead.
{% endhint %}

{% tabs %}
{% tab title="Template Example" %}
<pre class="language-go"><code class="lang-go"><strong>`{{ "42"| int64 }}`
</strong></code></pre>

_This example will takes the 42 string and convert it to an int64, the output will be `42`._

:x: No error handling
{% endtab %}
{% endtabs %}

### float64 :warning:

Signature: `float64(value any) float64`

{% hint style="warning" %}
**\[DEPRECATED]** Use [`toFloat64`](functions.md#tofloat64)instead.
{% endhint %}

{% tabs %}
{% tab title="Template Example" %}
<pre class="language-go"><code class="lang-go"><strong>`{{ "42.42"| float64 }}`
</strong></code></pre>

_This example will takes the 42.42 string and convert it to a float64, the output will be `42.42`._

:x: No error handling
{% endtab %}
{% endtabs %}

## Date Manipulation

## JSON Manipulation

