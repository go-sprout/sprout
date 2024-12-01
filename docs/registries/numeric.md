---
description: >-
  The Numeric registry includes a range of utilities for performing numerical
  operations and calculations, making it easier to handle numbers and perform
  math functions in your templates.
---

# Numeric

{% hint style="info" %}
You can easily import all the functions from the <mark style="color:yellow;">`numeric`</mark> registry by including the following import statement in your code

```go
import "github.com/go-sprout/sprout/registry/numeric"
```
{% endhint %}

### <mark style="color:purple;">floor</mark>

The function returns the largest integer that is less than or equal to the provided number.

<table data-header-hidden><thead><tr><th width="174">Name</th><th>Value</th></tr></thead><tbody><tr><td>Signature</td><td><pre class="language-go"><code class="lang-go">Floor(num any) (float64, error)
</code></pre></td></tr></tbody></table>

{% tabs %}
{% tab title="Template Example" %}
```go
{{ 3.7 | floor }} // Output: 3
{{ floor 1.5 }} // Output: 1
{{ floor 123.9999 }} // Output: 123
{{ floor 123.0001 }} // Output: 123
{{ floor "invalid" }} // Error
```
{% endtab %}
{% endtabs %}

### <mark style="color:purple;">ceil</mark>

The function returns the smallest integer that is greater than or equal to the provided number.

<table data-header-hidden><thead><tr><th width="164">Name</th><th>Value</th></tr></thead><tbody><tr><td>Signature</td><td><pre class="language-go"><code class="lang-go">Ceil(num any) (float64, error)
</code></pre></td></tr></tbody></table>

{% tabs %}
{% tab title="Template Example" %}
```go
{{ 3.1 | ceil }} // Output: 4
{{ ceil 1.5 }} // Output: 2
{{ ceil 123.9999 }} // Output: 124
{{ ceil 123.0001 }} // Output: 124
{{ ceil "invalid" }} // Error
```
{% endtab %}
{% endtabs %}

### <mark style="color:purple;">round</mark>

The function rounds a number to a specified precision, allowing control over the number of decimal places. It also considers an optional rounding threshold to determine whether to round up or down (default to 0.5).

<table data-header-hidden><thead><tr><th width="136">Name</th><th>Value</th></tr></thead><tbody><tr><td>Signature</td><td><pre class="language-go"><code class="lang-go">Round(num any, poww int, roundOpts ...float64) (float64, error)
</code></pre></td></tr></tbody></table>

{% tabs %}
{% tab title="Template Example" %}
```go
{{ round 3.746 2 }} // Output: 3.75
{{ round 3.746 2 0.5 }} // Output: 3.75
{{ round "123.5555" 3 }} // Output: 123.556
{{ round 123.49999999 0 }} // Output: 123
{{ round 123.2329999 2 .3 }} // Output: 123.23
```
{% endtab %}
{% endtabs %}

### <mark style="color:purple;">add / addf</mark>

The function performs addition on a slice of values, summing all elements in the slice and returning the total.

<table data-header-hidden><thead><tr><th width="164">Name</th><th>Value</th></tr></thead><tbody><tr><td>Signature</td><td><pre class="language-go"><code class="lang-go">Add(values ...any) (any, error)
</code></pre></td></tr></tbody></table>

{% tabs %}
{% tab title="Template Example" %}
```go
{{ add }} // Output: 0
{{ add 1 }} // Output: 1
{{ add 1 2 3 4 5 6 7 8 9 10 }} // Output: 55
{{ add 1.1 2.2 3.3 4.4 5.5 6.6 7.7 8.8 9.9 10.1 }} // Output: 59.6
```
{% endtab %}
{% endtabs %}

### <mark style="color:purple;">add1 / add1f</mark>

The function performs a unary addition, incrementing the provided value by one.

<table data-header-hidden><thead><tr><th width="164">Name</th><th>Value</th></tr></thead><tbody><tr><td>Signature</td><td><pre class="language-go"><code class="lang-go">Add1(x any) (any, error)
</code></pre></td></tr></tbody></table>

{% tabs %}
{% tab title="Template Example" %}
```go
{{ add1 -1 }} // Output: 0
{{ add1 1 }} // Output: 2
{{ add1 1.1 }} // Output: 2.1
```
{% endtab %}
{% endtabs %}

### <mark style="color:purple;">sub / subf</mark>

The function performs subtraction on a slice of values, starting with the first value and subtracting each subsequent value from it.

<table data-header-hidden><thead><tr><th width="164">Name</th><th>Value</th></tr></thead><tbody><tr><td>Signature</td><td><pre class="language-go"><code class="lang-go">Sub(values ...any) (any, error)
</code></pre></td></tr></tbody></table>

{% tabs %}
{% tab title="Template Example" %}
```go
{{ sub 10 3 2 }} // Output: 5
```
{% endtab %}
{% endtabs %}

### <mark style="color:purple;">mul</mark>

The function multiplies a sequence of values together and returns the result as an `int64`.

<table data-header-hidden><thead><tr><th width="164">Name</th><th>Value</th></tr></thead><tbody><tr><td>Signature</td><td><pre class="language-go"><code class="lang-go">MulInt(values ...any) (int64, error)
</code></pre></td></tr></tbody></table>

{% tabs %}
{% tab title="Template Example" %}
```go
{{ mul 1 1 }} // Output: 1
{{ mul 1.1 1.1 }} // Output: 1
{{ 3 | mul 14 }} // Output: 42
```
{% endtab %}
{% endtabs %}

### <mark style="color:purple;">mulf</mark>

The function multiplies a sequence of values and returns the result as a `float64`.

<table data-header-hidden><thead><tr><th width="164">Name</th><th>Value</th></tr></thead><tbody><tr><td>Signature</td><td><pre class="language-go"><code class="lang-go">Mulf(values ...any) (any, error)
</code></pre></td></tr></tbody></table>

{% tabs %}
{% tab title="Template Example" %}
```go
{{ mulf 1.1 1.1 }} // Output: 1.2100000000000002
```
{% endtab %}
{% endtabs %}

### <mark style="color:purple;">div</mark>

The function divides a sequence of values and returns the result as an `int64`.

<table data-header-hidden><thead><tr><th width="164">Name</th><th>Value</th></tr></thead><tbody><tr><td>Signature</td><td><pre class="language-go"><code class="lang-go">DivInt(values ...any) (int64, error)
</code></pre></td></tr></tbody></table>

{% tabs %}
{% tab title="Template Example" %}
```go
{{ div 1 1 }} // Output: 1
{{ div 1.1 1.1 }} // Output: 1
```
{% endtab %}
{% endtabs %}

### <mark style="color:purple;">divf</mark>

The function divides a sequence of values, starting with the first value, and returns the result as a `float64`.

<table data-header-hidden><thead><tr><th width="164">Name</th><th>Value</th></tr></thead><tbody><tr><td>Signature</td><td><pre class="language-go"><code class="lang-go">Divf(values ...any) (any, error)
</code></pre></td></tr></tbody></table>

{% tabs %}
{% tab title="Template Example" %}
```go
{{ divf 30.0 3.0 2.0 }} // Output: 5
{{ 2 | divf 5 4 }} // Output: 0.625
```
{% endtab %}
{% endtabs %}

### <mark style="color:purple;">mod</mark>

The function returns the remainder of the division of `x` by `y`.

<table data-header-hidden><thead><tr><th width="164">Name</th><th>Value</th></tr></thead><tbody><tr><td>Signature</td><td><pre class="language-go"><code class="lang-go">Mod(x any, y any) (any, error)
</code></pre></td></tr></tbody></table>

{% tabs %}
{% tab title="Template Example" %}
```go
{{ 4 | mod 10 }} // Output: 2
```
{% endtab %}
{% endtabs %}

### <mark style="color:purple;">min</mark>

The function returns the minimum value among the provided arguments.

<table data-header-hidden><thead><tr><th width="164">Name</th><th>Value</th></tr></thead><tbody><tr><td>Signature</td><td><pre class="language-go"><code class="lang-go">Min(a any, i ...any) (int64, error)
</code></pre></td></tr></tbody></table>

{% tabs %}
{% tab title="Template Example" %}
```go
{{ min 5 3 8 2 }} // Output: 2
```
{% endtab %}
{% endtabs %}

### <mark style="color:purple;">minf</mark>

The function returns the minimum value among the provided floating-point arguments.

<table data-header-hidden><thead><tr><th width="164">Name</th><th>Value</th></tr></thead><tbody><tr><td>Signature</td><td><pre class="language-go"><code class="lang-go">Minf(a any, i ...any) (float64, error)
</code></pre></td></tr></tbody></table>

{% tabs %}
{% tab title="Template Example" %}
```go
{{ minf 5.2 3.8 8.1 2.6 }} // Output: 2.6
```
{% endtab %}
{% endtabs %}

### <mark style="color:purple;">max</mark>

The function returns the maximum value among the provided arguments.

<table data-header-hidden><thead><tr><th width="164">Name</th><th>Value</th></tr></thead><tbody><tr><td>Signature</td><td><pre class="language-go"><code class="lang-go">Max(a any, i ...any) (int64, error)
</code></pre></td></tr></tbody></table>

{% tabs %}
{% tab title="Template Example" %}
```go
{{ max 5 3 8 2 }} // Output: 8
```
{% endtab %}
{% endtabs %}

### <mark style="color:purple;">maxf</mark>

The function returns the maximum value among the provided floating-point arguments.

<table data-header-hidden><thead><tr><th width="164">Name</th><th>Value</th></tr></thead><tbody><tr><td>Signature</td><td><pre class="language-go"><code class="lang-go">Maxf(a any, i ...any) (float64, error)
</code></pre></td></tr></tbody></table>

{% tabs %}
{% tab title="Template Example" %}
```go
{{ maxf 5.2 3.8 8.1 2.6 }} // Output: 8.1
```
{% endtab %}
{% endtabs %}
