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
{{ mulf 1.1 1.1 }} // Output: 1.21
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

### <mark style="color:purple;">sum</mark>

The function returns the total of the provided values as an integer. A single list argument is expanded, so a list of values coming from your data can be totalled directly.

<table data-header-hidden><thead><tr><th width="164">Name</th><th>Value</th></tr></thead><tbody><tr><td>Signature</td><td><pre class="language-go"><code class="lang-go">Sum(values ...any) (int64, error)
</code></pre></td></tr></tbody></table>

{% tabs %}
{% tab title="Template Example" %}
```go
{{ sum 1 2 3 4 }} // Output: 10
{{ sum (list 1 2 3 4) }} // Output: 10
{{ sum }} // Output: 0
{{ sum "invalid" }} // Error
```
{% endtab %}
{% endtabs %}

### <mark style="color:purple;">sumf</mark>

The function returns the total of the provided values as a floating-point number. A single list argument is expanded, so a list of values coming from your data can be totalled directly.

<table data-header-hidden><thead><tr><th width="164">Name</th><th>Value</th></tr></thead><tbody><tr><td>Signature</td><td><pre class="language-go"><code class="lang-go">Sumf(values ...any) (float64, error)
</code></pre></td></tr></tbody></table>

{% tabs %}
{% tab title="Template Example" %}
```go
{{ sumf 1.5 2.25 }} // Output: 3.75
{{ sumf (list 1.1 2.2 3.3) }} // Output: 6.6
{{ sumf 0.1 0.2 }} // Output: 0.3
{{ sumf "invalid" }} // Error
```
{% endtab %}
{% endtabs %}

### <mark style="color:purple;">mean</mark>

The function returns the arithmetic mean of the provided values as an integer. The result is truncated toward zero, so use <mark style="color:yellow;">`meanf`</mark> to keep the fractional part.

<table data-header-hidden><thead><tr><th width="164">Name</th><th>Value</th></tr></thead><tbody><tr><td>Signature</td><td><pre class="language-go"><code class="lang-go">Mean(values ...any) (int64, error)
</code></pre></td></tr></tbody></table>

{% tabs %}
{% tab title="Template Example" %}
```go
{{ mean 10 20 60 }} // Output: 30
{{ mean (list 2 4 6) }} // Output: 4
{{ mean 1 2 }} // Output: 1
{{ mean }} // Error
```
{% endtab %}
{% endtabs %}

### <mark style="color:purple;">meanf</mark>

The function returns the arithmetic mean of the provided values as a floating-point number.

<table data-header-hidden><thead><tr><th width="164">Name</th><th>Value</th></tr></thead><tbody><tr><td>Signature</td><td><pre class="language-go"><code class="lang-go">Meanf(values ...any) (float64, error)
</code></pre></td></tr></tbody></table>

{% tabs %}
{% tab title="Template Example" %}
```go
{{ meanf 1 2 }} // Output: 1.5
{{ meanf (list 1.5 2.5) }} // Output: 2
{{ meanf 10 20 60 }} // Output: 30
{{ meanf }} // Error
```
{% endtab %}
{% endtabs %}

### <mark style="color:purple;">median</mark>

The function returns the middle of the provided values as an integer, or the mean of the two middle values when the count is even. The result is truncated toward zero, so use <mark style="color:yellow;">`medianf`</mark> to keep the fractional part.

<table data-header-hidden><thead><tr><th width="164">Name</th><th>Value</th></tr></thead><tbody><tr><td>Signature</td><td><pre class="language-go"><code class="lang-go">Median(values ...any) (int64, error)
</code></pre></td></tr></tbody></table>

{% tabs %}
{% tab title="Template Example" %}
```go
{{ median 60 10 20 }} // Output: 20
{{ median (list 10 20 30 40) }} // Output: 25
{{ median 1 2 }} // Output: 1
{{ median }} // Error
```
{% endtab %}
{% endtabs %}

### <mark style="color:purple;">medianf</mark>

The function returns the middle of the provided values as a floating-point number, or the mean of the two middle values when the count is even. Unlike the mean, it is not pulled by a single outlying value.

<table data-header-hidden><thead><tr><th width="164">Name</th><th>Value</th></tr></thead><tbody><tr><td>Signature</td><td><pre class="language-go"><code class="lang-go">Medianf(values ...any) (float64, error)
</code></pre></td></tr></tbody></table>

{% tabs %}
{% tab title="Template Example" %}
```go
{{ medianf 60 10 20 }} // Output: 20
{{ medianf (list 10 20 30 40) }} // Output: 25
{{ medianf 1 2 }} // Output: 1.5
{{ medianf }} // Error
```
{% endtab %}
{% endtabs %}

### <mark style="color:purple;">mode</mark>

The function returns the most frequent of the provided values as an integer. When several values are equally frequent, the one appearing first is returned.

<table data-header-hidden><thead><tr><th width="164">Name</th><th>Value</th></tr></thead><tbody><tr><td>Signature</td><td><pre class="language-go"><code class="lang-go">Mode(values ...any) (int64, error)
</code></pre></td></tr></tbody></table>

{% tabs %}
{% tab title="Template Example" %}
```go
{{ mode 25 50 25 }} // Output: 25
{{ mode (list 9 9 1) }} // Output: 9
{{ mode 5 3 5 3 }} // Output: 5
{{ mode }} // Error
```
{% endtab %}
{% endtabs %}

### <mark style="color:purple;">modef</mark>

The function returns the most frequent of the provided values as a floating-point number. When several values are equally frequent, the one appearing first is returned.

<table data-header-hidden><thead><tr><th width="164">Name</th><th>Value</th></tr></thead><tbody><tr><td>Signature</td><td><pre class="language-go"><code class="lang-go">Modef(values ...any) (float64, error)
</code></pre></td></tr></tbody></table>

{% tabs %}
{% tab title="Template Example" %}
```go
{{ modef 2.5 1.75 2.5 }} // Output: 2.5
{{ modef (list 1.1 2.2 1.1) }} // Output: 1.1
{{ modef }} // Error
```
{% endtab %}
{% endtabs %}

### <mark style="color:purple;">spread</mark>

The function returns the distance between the largest and smallest of the provided values as an integer. It is named <mark style="color:yellow;">`spread`</mark> rather than <mark style="color:yellow;">`range`</mark> because <mark style="color:yellow;">`range`</mark> is a keyword of the template language and cannot be used as a function name.

<table data-header-hidden><thead><tr><th width="164">Name</th><th>Value</th></tr></thead><tbody><tr><td>Signature</td><td><pre class="language-go"><code class="lang-go">Spread(values ...any) (int64, error)
</code></pre></td></tr></tbody></table>

{% tabs %}
{% tab title="Template Example" %}
```go
{{ spread 10 60 20 }} // Output: 50
{{ spread (list 3 9 1) }} // Output: 8
{{ spread -10 10 }} // Output: 20
{{ spread }} // Error
```
{% endtab %}
{% endtabs %}

### <mark style="color:purple;">spreadf</mark>

The function returns the distance between the largest and smallest of the provided values as a floating-point number.

<table data-header-hidden><thead><tr><th width="164">Name</th><th>Value</th></tr></thead><tbody><tr><td>Signature</td><td><pre class="language-go"><code class="lang-go">Spreadf(values ...any) (float64, error)
</code></pre></td></tr></tbody></table>

{% tabs %}
{% tab title="Template Example" %}
```go
{{ spreadf 1.5 4.25 }} // Output: 2.75
{{ spreadf (list 2.5 0.5) }} // Output: 2
{{ spreadf 7 }} // Output: 0
{{ spreadf }} // Error
```
{% endtab %}
{% endtabs %}
