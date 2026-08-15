---
description: >-
  The Uniqueid registry offers functions to generate unique identifiers, such as
  UUIDs, which are essential for creating distinct and traceable entities in
  your applications.
---

# Uniqueid

{% hint style="info" %}
You can easily import all the functions from the <mark style="color:yellow;">`uniqueid`</mark> registry by including the following import statement in your code

```go
import "github.com/go-sprout/sprout/registry/uniqueid"
```
{% endhint %}

### <mark style="color:purple;">uuidv4</mark>

Uuidv4 generates a new random UUID (Universally Unique Identifier) version 4.

<table data-header-hidden><thead><tr><th width="174">Name</th><th>Value</th></tr></thead><tbody><tr><td>Signature</td><td><pre class="language-go"><code class="lang-go">Uuidv4() string
</code></pre></td></tr></tbody></table>

{% tabs %}
{% tab title="Template Example" %}
```go
{{ uuidv4 }} // Output(will be different): 3f0c463e-53f5-4f05-a2ec-3c083aa8f937
```
{% endtab %}
{% endtabs %}

### <mark style="color:purple;">uuidv7</mark>

Uuidv7 generates a new UUID version 7, based on the current Unix time in milliseconds. Unlike a version 4, successive UUIDs are sortable by generation time.

<table data-header-hidden><thead><tr><th width="174">Name</th><th>Value</th></tr></thead><tbody><tr><td>Signature</td><td><pre class="language-go"><code class="lang-go">Uuidv7() (string, error)
</code></pre></td></tr></tbody></table>

{% tabs %}
{% tab title="Template Example" %}
```go
{{ uuidv7 }} // Output(will be different): 018f5395-4e88-7c2a-9f3b-1d7e4a6c8b90
```
{% endtab %}
{% endtabs %}

{% hint style="info" %}
Prefer `uuidv7` over `uuidv4` when the identifiers are used as database keys, the time ordering keeps the index locality that a fully random UUID destroys.
{% endhint %}

### <mark style="color:purple;">uuidv5</mark>

Uuidv5 generates a UUID version 5, derived from a namespace and a name using SHA-1. The same namespace and name always produce the same UUID.

<table data-header-hidden><thead><tr><th width="164">Name</th><th>Value</th></tr></thead><tbody><tr><td>Signature</td><td><pre class="language-go"><code class="lang-go">Uuidv5(namespace string, name string) (string, error)
</code></pre></td></tr></tbody></table>

{% tabs %}
{% tab title="Template Example" %}
```go
{{ "python.org" | uuidv5 "dns" }} // Output: 886313e1-3b8a-5372-9b90-0c9aee199e5d
{{ "python.org" | uuidv5 "6ba7b810-9dad-11d1-80b4-00c04fd430c8" }} // Output: 886313e1-3b8a-5372-9b90-0c9aee199e5d
{{ "python.org" | uuidv5 "invalid" }} // Error
```
{% endtab %}
{% endtabs %}

{% hint style="info" %}
The namespace is either one of the predefined names `dns`, `url`, `oid`, `x500` (case insensitive), or any valid UUID to use a namespace of your own.
{% endhint %}

### <mark style="color:purple;">uuidv3</mark>

Uuidv3 generates a UUID version 3, derived from a namespace and a name using MD5. The same namespace and name always produce the same UUID.

<table data-header-hidden><thead><tr><th width="164">Name</th><th>Value</th></tr></thead><tbody><tr><td>Signature</td><td><pre class="language-go"><code class="lang-go">Uuidv3(namespace string, name string) (string, error)
</code></pre></td></tr></tbody></table>

{% tabs %}
{% tab title="Template Example" %}
```go
{{ "python.org" | uuidv3 "dns" }} // Output: 6fa459ea-ee8a-3ca4-894e-db77e160355e
{{ "python.org" | uuidv3 "invalid" }} // Error
```
{% endtab %}
{% endtabs %}

{% hint style="warning" %}
Prefer `uuidv5` for new usages, the version 3 only exists for compatibility with systems already relying on MD5 derived UUIDs.
{% endhint %}

### <mark style="color:purple;">uuidNil</mark>

UuidNil returns the nil UUID, the UUID with all its bits set to zero.

<table data-header-hidden><thead><tr><th width="164">Name</th><th>Value</th></tr></thead><tbody><tr><td>Signature</td><td><pre class="language-go"><code class="lang-go">UuidNil() string
</code></pre></td></tr></tbody></table>

{% tabs %}
{% tab title="Template Example" %}
```go
{{ uuidNil }} // Output: 00000000-0000-0000-0000-000000000000
```
{% endtab %}
{% endtabs %}

### <mark style="color:purple;">isUUID</mark>

IsUUID checks if the given value is a valid UUID.

<table data-header-hidden><thead><tr><th width="164">Name</th><th>Value</th></tr></thead><tbody><tr><td>Signature</td><td><pre class="language-go"><code class="lang-go">IsUUID(value string) bool
</code></pre></td></tr></tbody></table>

{% tabs %}
{% tab title="Template Example" %}
```go
{{ "886313e1-3b8a-5372-9b90-0c9aee199e5d" | isUUID }} // Output: true
{{ "urn:uuid:886313e1-3b8a-5372-9b90-0c9aee199e5d" | isUUID }} // Output: true
{{ "886313e13b8a53729b900c9aee199e5d" | isUUID }} // Output: true
{{ "not-a-uuid" | isUUID }} // Output: false
```
{% endtab %}
{% endtabs %}

{% hint style="info" %}
On top of the canonical form, three alternative forms are accepted: the urn prefixed one (`urn:uuid:886313e1-…`), the braced one (`{886313e1-…}`) and the one without hyphen (`886313e13b8a…`). Use `uuidVersion` if you also need to validate which version you received.
{% endhint %}

### <mark style="color:purple;">uuidVersion</mark>

UuidVersion returns the version of the given UUID.

<table data-header-hidden><thead><tr><th width="164">Name</th><th>Value</th></tr></thead><tbody><tr><td>Signature</td><td><pre class="language-go"><code class="lang-go">UuidVersion(value string) (int, error)
</code></pre></td></tr></tbody></table>

{% tabs %}
{% tab title="Template Example" %}
```go
{{ "886313e1-3b8a-5372-9b90-0c9aee199e5d" | uuidVersion }} // Output: 5
{{ "not-a-uuid" | uuidVersion }} // Error
```
{% endtab %}
{% endtabs %}

### <mark style="color:purple;">uuidTime</mark>

UuidTime returns the time embedded in the given UUID. Only the versions carrying a timestamp are supported, namely the versions 1, 2, 6 and 7.

<table data-header-hidden><thead><tr><th width="164">Name</th><th>Value</th></tr></thead><tbody><tr><td>Signature</td><td><pre class="language-go"><code class="lang-go">UuidTime(value string) (time.Time, error)
</code></pre></td></tr></tbody></table>

{% tabs %}
{% tab title="Template Example" %}
```go
{{ dateInZone "2006-01-02 15:04:05" ("018f5395-4e88-7000-8000-000000000000" | uuidTime) "UTC" }} // Output: 2024-05-07 15:04:05
{{ "886313e1-3b8a-5372-9b90-0c9aee199e5d" | uuidTime }} // Error
{{ "not-a-uuid" | uuidTime }} // Error
```
{% endtab %}
{% endtabs %}

