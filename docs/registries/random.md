---
description: >-
  The Random registry provides functions to generate random numbers, strings,
  and other data types, useful for scenarios requiring randomness or unique
  identifiers.
---

# Random

{% hint style="info" %}
You can easily import all the functions from the <mark style="color:yellow;">`random`</mark> registry by including the following import statement in your code

```go
import "github.com/go-sprout/sprout/registry/random"
```
{% endhint %}

### <mark style="color:purple;">randAlphaNum</mark>

The function generates a random alphanumeric string with the specified length, combining both letters and numbers to create a unique sequence.

<table data-header-hidden><thead><tr><th width="174">Name</th><th>Value</th></tr></thead><tbody><tr><td>Signature</td><td><pre class="language-go"><code class="lang-go">RandAlphaNumeric(count int) string
</code></pre></td></tr></tbody></table>

{% tabs %}
{% tab title="Template Example" %}
```go
{{ randAlphaNum 10 }} // Output(will be different): "OnmS3BPwBl"
```
{% endtab %}
{% endtabs %}

### <mark style="color:purple;">randAlpha</mark>

The function generates a random string consisting of only alphabetic characters (letters) with the specified length.

<table data-header-hidden><thead><tr><th width="164">Name</th><th>Value</th></tr></thead><tbody><tr><td>Signature</td><td><pre class="language-go"><code class="lang-go">RandAlpha(count int) string
</code></pre></td></tr></tbody></table>

{% tabs %}
{% tab title="Template Example" %}
```go
{{ 10 | randAlpha }} // Output(will be different): "rBxkROwxav"
```
{% endtab %}
{% endtabs %}

### <mark style="color:purple;">randAscii</mark>

The function generates a random ASCII string of the specified length, using characters within the ASCII range 32 to 126.

<table data-header-hidden><thead><tr><th width="164">Name</th><th>Value</th></tr></thead><tbody><tr><td>Signature</td><td><pre class="language-go"><code class="lang-go">RandAscii(count int) string
</code></pre></td></tr></tbody></table>

{% tabs %}
{% tab title="Template Example" %}
```go
{{ randAscii 10 }} // Output(will be different): "}]~>_<:^%"
```
{% endtab %}
{% endtabs %}

### <mark style="color:purple;">randNumeric</mark>

The function generates a random numeric string consisting only of digits, with the specified length.

<table data-header-hidden><thead><tr><th width="164">Name</th><th>Value</th></tr></thead><tbody><tr><td>Signature</td><td><pre class="language-go"><code class="lang-go">RandNumeric(count int) string
</code></pre></td></tr></tbody></table>

{% tabs %}
{% tab title="Template Example" %}
```go
{{ randNumeric 10 }} // Output(will be different): "3269896295"
```
{% endtab %}
{% endtabs %}

### <mark style="color:purple;">randBytes</mark>

The function generates a random byte array of the specified length and returns it as a Base64 encoded string.

<table data-header-hidden><thead><tr><th width="164">Name</th><th>Value</th></tr></thead><tbody><tr><td>Signature</td><td><pre class="language-go"><code class="lang-go">RandBytes(count int) (string, error)
</code></pre></td></tr></tbody></table>

{% tabs %}
{% tab title="Template Example" %}
```go
{{ randBytes 16 }} // Output(will be different): "c3RhY2thYnVzZSByb2NrcyE="
```
{% endtab %}
{% endtabs %}

### <mark style="color:purple;">randInt</mark>

The function generates a random integer between the specified minimum and maximum values (inclusive). It takes two parameters: `min` for the minimum value and `max` for the maximum value. The function then returns a random integer within this range.

<table data-header-hidden><thead><tr><th width="164">Name</th><th>Value</th></tr></thead><tbody><tr><td>Signature</td><td><pre class="language-go"><code class="lang-go">RandInt(min, max int) int
</code></pre></td></tr></tbody></table>

{% tabs %}
{% tab title="Template Example" %}
```go
{{ randInt 1 10 }} // Output(will be different): 5
```
{% endtab %}
{% endtabs %}
