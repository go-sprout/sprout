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
{{ uuidv4 }} // Output: "3f0c463e-53f5-4f05-a2ec-3c083aa8f937"
```
{% endtab %}
{% endtabs %}

