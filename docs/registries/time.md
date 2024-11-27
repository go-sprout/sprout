---
description: >-
  The Time registry provides tools to manage and manipulate dates, times, and
  time-related calculations, making it easy to handle time-based data in your
  projects.
---

# Time

{% hint style="info" %}
You can easily import all the functions from the <mark style="color:yellow;">`time`</mark> registry by including the following import statement in your code

```go
import "github.com/go-sprout/sprout/registry/time"
```
{% endhint %}

### <mark style="color:purple;">date</mark>

The function formats a given date or the current time into a specified format string.

<table data-header-hidden><thead><tr><th width="174">Name</th><th>Value</th></tr></thead><tbody><tr><td>Signature</td><td><pre class="language-go"><code class="lang-go"> Date(fmt string, date any) (string, error)
</code></pre></td></tr></tbody></table>

{% tabs %}
{% tab title="Template Example" %}
```go
{{ "2023-05-04T15:04:05Z" | toDate "2006-01-02T15:04:05Z" | date "Jan 2, 2006" }} // Output: "May 4, 2023"
```
{% endtab %}
{% endtabs %}

### <mark style="color:purple;">dateInZone</mark>

The function formats a given date or the current time into a specified format string for a specified timezone.

<table data-header-hidden><thead><tr><th width="124">Name</th><th>Value</th></tr></thead><tbody><tr><td>Signature</td><td><pre class="language-go"><code class="lang-go">DateInZone(fmt string, date any, zone string) (string, error)
</code></pre></td></tr></tbody></table>

{% tabs %}
{% tab title="Template Example" %}
```go
{{- $date := "2023-05-04T15:04:05Z" | toDate "2006-01-02T15:04:05Z" -}}
{{ dateInZone "Jan 2, 2006" $date "UTC" }} // Output: "May 4, 2023"

{{- $date := "2023-05-04T15:04:05Z" | toDate "2006-01-02T15:04:05Z" -}}
{{ dateInZone "Jan 2, 2006" $date "invalid" }} // Error
```
{% endtab %}
{% endtabs %}

### <mark style="color:purple;">duration</mark>

The function converts a given number of seconds into a human-readable duration string.

<table data-header-hidden><thead><tr><th width="164">Name</th><th>Value</th></tr></thead><tbody><tr><td>Signature</td><td><pre class="language-go"><code class="lang-go">Duration(sec any) string
</code></pre></td></tr></tbody></table>

{% tabs %}
{% tab title="Template Example" %}
```go
{{ 3661 | duration }} // Output: "1h1m1s"
```
{% endtab %}
{% endtabs %}

### <mark style="color:purple;">dateAgo</mark>

The function calculates the time elapsed since a given date.

<table data-header-hidden><thead><tr><th width="164">Name</th><th>Value</th></tr></thead><tbody><tr><td>Signature</td><td><pre class="language-go"><code class="lang-go">DateAgo(date any) string
</code></pre></td></tr></tbody></table>

{% tabs %}
{% tab title="Template Example" %}
```go
{{- $date := now | dateModify "-10h" -}}
{{ $date | dateAgo }} // Output: 10h0m0s
```
{% endtab %}
{% endtabs %}

{% hint style="info" %}
`dateAgo` can receive multiples input types like:

* `time.Time` object (object or pointer)
* `int`, `int32`, `int64` converted as Unix
{% endhint %}

### <mark style="color:purple;">now</mark>

The function returns the current time.

<table data-header-hidden><thead><tr><th width="164">Name</th><th>Value</th></tr></thead><tbody><tr><td>Signature</td><td><pre class="language-go"><code class="lang-go">Now() time.Time
</code></pre></td></tr></tbody></table>

{% tabs %}
{% tab title="Template Example" %}
```go
{{ now }} // Output(will be different): "2023-05-07T15:04:05Z"
```
{% endtab %}
{% endtabs %}

### <mark style="color:purple;">unixEpoch</mark>

The function returns the Unix epoch timestamp for a given date.

<table data-header-hidden><thead><tr><th width="164">Name</th><th>Value</th></tr></thead><tbody><tr><td>Signature</td><td><pre class="language-go"><code class="lang-go">UnixEpoch(date time.Time) string
</code></pre></td></tr></tbody></table>

{% tabs %}
{% tab title="Template Example" %}
```go
{{ now | unixEpoch }} // Output(will be different): 1683306245
```
{% endtab %}
{% endtabs %}

### <mark style="color:purple;">dateModify</mark>

The function adjusts a given date by a specified duration, returning the modified date. If the duration format is incorrect, it returns the original date without any changes, in case of must version, an error is returned.

<table data-header-hidden><thead><tr><th width="164">Name</th><th>Value</th></tr></thead><tbody><tr><td>Signature</td><td><pre class="language-go"><code class="lang-go">DateModify(fmt string, date time.Time) (time.Time, error)
</code></pre></td></tr></tbody></table>

{% tabs %}
{% tab title="Template Example" %}
```go
{{- $date := "2023-05-04T15:04:05Z" | toDate "2006-01-02T15:04:05Z" -}}
{{ $date | dateModify "48h" }}
// Output: 2023-05-06 15:04:05 +0000 UTC
{{- $date := "2023-05-04T15:04:05Z" | toDate "2006-01-02T15:04:05Z" -}}
{{ $date | dateModify "0z+" }}
// Error
```
{% endtab %}
{% endtabs %}

### <mark style="color:purple;">durationRound</mark>

The function rounds a duration to the nearest significant time unit, such as years or seconds.

<table data-header-hidden><thead><tr><th width="164">Name</th><th>Value</th></tr></thead><tbody><tr><td>Signature</td><td><pre class="language-go"><code class="lang-go">DurationRound(duration any) string
</code></pre></td></tr></tbody></table>

{% tabs %}
{% tab title="Template Example" %}
```go
{{ "9600h" | durationRound }} // Output: "1y"
{{ "960h" | durationRound }} // Output: "1mo"
{{ "192h" | durationRound }} // Output: "8d"
{{ "3600s" | durationRound }} // Output: "60m"
{{ "300s" | durationRound }} // Output: "5m"
{{ "61s" | durationRound }} // Output: "1m"
{{ "59s" | durationRound }} // Output: "59s"
```
{% endtab %}
{% endtabs %}

### <mark style="color:purple;">htmlDate</mark>

The function formats a date into the standard HTML date format (YYYY-MM-DD).

<table data-header-hidden><thead><tr><th width="164">Name</th><th>Value</th></tr></thead><tbody><tr><td>Signature</td><td><pre class="language-go"><code class="lang-go">HtmlDate(date any) (string, error)
</code></pre></td></tr></tbody></table>

{% tabs %}
{% tab title="Template Example" %}
```go
{{- $date := "2023-05-04T15:04:05Z" | toDate "2006-01-02T15:04:05Z" -}}
{{ $date | htmlDate }} // Output: "2023-05-04"
```
{% endtab %}
{% endtabs %}

{% hint style="info" %}
_This basically call_ `dateInZone("2006-01-02", date, "Local")`
{% endhint %}

### <mark style="color:purple;">htmlDateInZone</mark>

The function formats a date into the standard HTML date format (YYYY-MM-DD) based on a specified timezone.

<table data-header-hidden><thead><tr><th width="164">Name</th><th>Value</th></tr></thead><tbody><tr><td>Signature</td><td><pre class="language-go"><code class="lang-go">HtmlDateInZone(date any, zone string) (string, error)
</code></pre></td></tr></tbody></table>

{% tabs %}
{% tab title="Template Example" %}
```go
{{- $date := "2023-05-04T15:04:05Z" | toDate "2006-01-02T15:04:05Z" -}}
{{ htmlDateInZone $date "UTC" }} // Output: "2023-05-04"
```
{% endtab %}
{% endtabs %}

{% hint style="info" %}
_This basically call_ `dateInZone("2006-01-02", date, zone)`
{% endhint %}
