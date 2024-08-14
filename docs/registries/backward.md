---
description: >-
  The Backward registry offers functions to maintain compatibility with older
  Sprig versions, ensuring legacy templates continue to work seamlessly after
  updates.
---

# Backward

{% hint style="info" %}
You can easily import all the functions from the <mark style="color:yellow;">`backward`</mark> registry by including the following import statement in your code

```go
import "github.com/go-sprout/sprout/registry/backward"
```
{% endhint %}

## <mark style="color:red;">Deprecated functions</mark>

### fail ⚠️

The `Fail` function creates an error with a specified message and returns a `nil` pointer along with the error. It is generally used to indicate failure in functions that return both a pointer and an error.

{% hint style="warning" %}
**\[DEPRECATED]** No replacement are scheduled yet. If you need this function, open an issue to clearly explain the usage.
{% endhint %}

<table data-header-hidden><thead><tr><th width="164">Name</th><th>Value</th></tr></thead><tbody><tr><td>Signature</td><td><pre class="language-go"><code class="lang-go">Fail(message string) (*uint, error)
</code></pre></td></tr><tr><td>Must version</td><td><span data-gb-custom-inline data-tag="emoji" data-code="274c">❌</span></td></tr></tbody></table>

{% tabs %}
{% tab title="Template Example" %}
```go
{{ fail "Operation failed" }} // Output: nil, error with "Operation failed"
```
{% endtab %}
{% endtabs %}

### urlParse ⚠️

The function parses a given URL string and returns a map containing its components, such as the scheme, host, path, query parameters, and more, allowing easy access and manipulation of the different parts of a URL.

{% hint style="warning" %}
**\[DEPRECATED]** No replacement are scheduled yet. If you need this function, open an issue to clearly explain the usage.
{% endhint %}

<table data-header-hidden><thead><tr><th width="164">Name</th><th>Value</th></tr></thead><tbody><tr><td>Signature</td><td><pre class="language-go"><code class="lang-go">UrlParse(v string) map[string]any
</code></pre></td></tr><tr><td>Must version</td><td><span data-gb-custom-inline data-tag="emoji" data-code="274c">❌</span></td></tr></tbody></table>

{% tabs %}
{% tab title="Template Example" %}
```go
{{ "https://example.com/path?query=1#fragment" | urlParse }}
// Output: map[fragment:fragment host:example.com hostname:example.com path:path query:query scheme:https]
```
{% endtab %}
{% endtabs %}

### urlJoin ⚠️

The function constructs a URL string from a given map of URL components, assembling the various parts such as scheme, host, path, and query parameters into a complete URL.

{% hint style="warning" %}
**\[DEPRECATED]** No replacement are scheduled yet. If you need this function, open an issue to clearly explain the usage.
{% endhint %}

<table data-header-hidden><thead><tr><th width="164">Name</th><th>Value</th></tr></thead><tbody><tr><td>Signature</td><td><pre class="language-go"><code class="lang-go">UrlJoin(d map[string]any) string
</code></pre></td></tr><tr><td>Must version</td><td><span data-gb-custom-inline data-tag="emoji" data-code="274c">❌</span></td></tr></tbody></table>

{% tabs %}
{% tab title="Template Example" %}
```go
{{ dict scheme="https" host="example.com" path="/path" query="query=1" opaque="opaque" fragment="fragment" | urlJoin }}
// Output: "https://example.com/path?query=1#fragment"

```
{% endtab %}
{% endtabs %}

### getHostByName⚠️

The function returns a random IP address associated with a given hostname, providing a way to resolve a hostname to one of its corresponding IP addresses.

{% hint style="warning" %}
**\[DEPRECATED]** No replacement are scheduled yet. If you need this function, open an issue to clearly explain the usage.
{% endhint %}

<table data-header-hidden><thead><tr><th width="164">Name</th><th>Value</th></tr></thead><tbody><tr><td>Signature</td><td><pre class="language-go"><code class="lang-go">GetHostByName(name string) string
</code></pre></td></tr><tr><td>Must version</td><td><span data-gb-custom-inline data-tag="emoji" data-code="274c">❌</span></td></tr></tbody></table>

{% tabs %}
{% tab title="Template Example" %}
```go
{{ getHostByName "localhost" }} // Output: "127.0.0.1"
```
{% endtab %}
{% endtabs %}
