---
description: >-
  The Maps registry offers tools for creating, manipulating, and interacting
  with map data structures, facilitating efficient data organization and
  retrieval.
---

# Maps

{% hint style="info" %}
You can easily import all the functions from the <mark style="color:yellow;">`maps`</mark> registry by including the following import statement in your code

```go
import "github.com/go-sprout/sprout/registry/maps"
```
{% endhint %}

### <mark style="color:purple;">dict</mark>

The function creates a dictionary (map) from a list of alternating keys and values, pairing each key with its corresponding value.

<table data-header-hidden><thead><tr><th width="174">Name</th><th>Value</th></tr></thead><tbody><tr><td>Signature</td><td><pre class="language-go"><code class="lang-go">Dict(values ...any) map[string]any
</code></pre></td></tr></tbody></table>

{% tabs %}
{% tab title="Template Example" %}
```go
{{ dict "key1" "value1" "key2" "value2" }}
// Output: {"key1": "value1", "key2": "value2"}
```
{% endtab %}
{% endtabs %}

### <mark style="color:purple;">get</mark>

The function retrieves the value associated with a specified key from a dictionary (map). If the key is found, the corresponding value is returned.

<table data-header-hidden><thead><tr><th width="164">Name</th><th>Value</th></tr></thead><tbody><tr><td>Signature</td><td><pre class="language-go"><code class="lang-go">Get(key string, dict map[string]any) (any, error)
</code></pre></td></tr></tbody></table>

{% tabs %}
{% tab title="Template Example" %}
```go
{{ {"key": "value"} | get "key" }} // Output: "value"
{{ {"key": "value"} | get "invalid" }} // Output: ""
```
{% endtab %}
{% endtabs %}

### <mark style="color:purple;">set</mark>

The function adds a new key-value pair to a dictionary or updates the value associated with an existing key.

<table data-header-hidden><thead><tr><th width="164">Name</th><th>Value</th></tr></thead><tbody><tr><td>Signature</td><td><pre class="language-go"><code class="lang-go">Set(key string, value any, dict map[string]any) (map[string]any, error)
</code></pre></td></tr></tbody></table>

{% tabs %}
{% tab title="Template Example" %}
```go
{{ {"key": "oldValue"} | set "key" "newValue" }} // Output: {"key": "newValue"}
{{ {"foo": "bar"} | set "far" "boo" }} // Output: {"foo": "bar", "far": "boo"}
```
{% endtab %}
{% endtabs %}

### <mark style="color:purple;">unset</mark>

The function remove the key-value pair to a dictionary associated with an existing key.

<table data-header-hidden><thead><tr><th width="164">Name</th><th>Value</th></tr></thead><tbody><tr><td>Signature</td><td><pre class="language-go"><code class="lang-go">Unset(key string, dict map[string]any) map[string]any
</code></pre></td></tr></tbody></table>

{% tabs %}
{% tab title="Template Example" %}
```go
{{ {"key": "value"} | unset "key" }} // Output: {}
{{ {"key": "value"} | unset "invalid" }} // Output: {"key": "value"}
```
{% endtab %}
{% endtabs %}

### <mark style="color:purple;">keys</mark>

The function retrieves all keys from one or more dictionaries, returning them as a list.

<table data-header-hidden><thead><tr><th width="164">Name</th><th>Value</th></tr></thead><tbody><tr><td>Signature</td><td><pre class="language-go"><code class="lang-go">Keys(dicts ...map[string]any) []string
</code></pre></td></tr></tbody></table>

{% tabs %}
{% tab title="Template Example" %}
```go
{{ $d := dict "key1" "value1" "key2" "value2"}}
{{ keys $d }} // Output: ["key1", "key2"]

{{ $d1 := dict "key1" "value1" }}
{{ $d2 := dict "key2" "value2" }}
{{ keys $d1 $d2 }} // Output: ["key1", "key2"]
```
{% endtab %}
{% endtabs %}

### <mark style="color:purple;">values</mark>

The function retrieves all values from one or more dictionaries, returning them as a list.

<table data-header-hidden><thead><tr><th width="164">Name</th><th>Value</th></tr></thead><tbody><tr><td>Signature</td><td><pre class="language-go"><code class="lang-go">Values(dicts ...map[string]any) []any
</code></pre></td></tr></tbody></table>

{% tabs %}
{% tab title="Template Example" %}
<pre class="language-go"><code class="lang-go">{{ $d := dict "key1" "value1" "key2" "value2"}}
<strong>{{ values $d }} // Output: ["value1", "value2"]
</strong></code></pre>
{% endtab %}
{% endtabs %}

### <mark style="color:purple;">pluck</mark>

The function extracts values associated with a specified key from a list of dictionaries, returning a list of those values.

<table data-header-hidden><thead><tr><th width="164">Name</th><th>Value</th></tr></thead><tbody><tr><td>Signature</td><td><pre class="language-go"><code class="lang-go">Pluck(key string, dicts ...map[string]any) []any
</code></pre></td></tr></tbody></table>

{% tabs %}
{% tab title="Template Example" %}
<pre class="language-go"><code class="lang-go"><strong>{{ $d1 := dict "key" "value1"}}
</strong>{{ $d2 := dict "key" "value2" }}
{{ pluck "key" $d1 $d2 }} // Output: ["value1", "value2"]
</code></pre>
{% endtab %}
{% endtabs %}

### <mark style="color:purple;">pick</mark>

The function creates a new dictionary that includes only the specified keys from the original dictionary, effectively filtering out all other keys and their associated values.

<table data-header-hidden><thead><tr><th width="164">Name</th><th>Value</th></tr></thead><tbody><tr><td>Signature</td><td><pre class="language-go"><code class="lang-go">Pick(keys ...string, dict map[string]any) (map[string]any, error)
</code></pre></td></tr></tbody></table>

{% tabs %}
{% tab title="Template Example" %}
```go
{{ $d := dict "key1" "value1" "key2" "value2" "key3" "value3" }}
{{ $d | pick "key1" "key3" }} // Output: {"key1": "value1", "key3": "value3"}
```
{% endtab %}
{% endtabs %}

### <mark style="color:purple;">omit</mark>

The function creates a new dictionary by excluding the specified keys from the original dictionary, effectively removing those key-value pairs from the resulting dictionary.

<table data-header-hidden><thead><tr><th width="164">Name</th><th>Value</th></tr></thead><tbody><tr><td>Signature</td><td><pre class="language-go"><code class="lang-go">Omit(keys ...string, dict map[string]any) (map[string]any, error)
</code></pre></td></tr></tbody></table>

{% tabs %}
{% tab title="Template Example" %}
```go
{{ $d := dict "key1" "value1" "key2" "value2" "key3" "value3" }}
{{ $d | omit "key1" "key3" }} // Output: {"key2": "value2"}
```
{% endtab %}
{% endtabs %}

### <mark style="color:purple;">dig</mark>

The function navigates through a nested dictionary structure using a sequence of keys and returns the value found at the specified path, allowing access to deeply nested data. The last argument must be the map.

<table data-header-hidden><thead><tr><th width="164">Name</th><th>Value</th></tr></thead><tbody><tr><td>Signature</td><td><pre class="language-go"><code class="lang-go">Dig(args ...any) (any, error)
</code></pre></td></tr></tbody></table>

{% tabs %}
{% tab title="Template Example" %}
```go
{{ $nest := dict "foo" "bar" }}
{{ $d := dict "key1" "value1" "nested" $nest }}
{{ $d | dig "nested" "foo" }} // Output: "bar"
{{ $d | dig "nested.foo" }} // Output: "bar"
```
{% endtab %}
{% endtabs %}

### <mark style="color:purple;">hasKey</mark>

The function checks whether a specified key exists in the dictionary, returning `true` if the key is found and `false` otherwise.

<table data-header-hidden><thead><tr><th width="164">Name</th><th>Value</th></tr></thead><tbody><tr><td>Signature</td><td><pre class="language-go"><code class="lang-go">HasKey(key string, dict map[string]any) (bool, error)
</code></pre></td></tr></tbody></table>

{% tabs %}
{% tab title="Template Example" %}
```go
{{ $d := dict "key1" "value1" }}
{{ $d | hasKey "key1" }} // Output: true
{{ $d | hasKey "key2" }} // Output: false
```
{% endtab %}
{% endtabs %}

### <mark style="color:purple;">merge</mark>

The function combines multiple source maps into a single destination map, adding new key-value pairs without overwriting any existing keys in the destination map.

<table data-header-hidden><thead><tr><th width="164">Name</th><th>Value</th></tr></thead><tbody><tr><td>Signature</td><td><pre class="language-go"><code class="lang-go">Merge(dest map[string]any, srcs ...map[string]any) (any, error)
</code></pre></td></tr></tbody></table>

{% tabs %}
{% tab title="Template Example" %}
<pre class="language-go"><code class="lang-go"><strong>{{ $d1 := dict "a" 1 "b" 2 }}
</strong><strong>{{ $d2 := dict "b" 3 "c" 4 }}
</strong><strong>{{ merge $d1 $d2 }} // Output: {"a": 1, "b": 2, "c": 4}, nil
</strong></code></pre>
{% endtab %}
{% endtabs %}

### <mark style="color:purple;">mergeOverwrite</mark>

The function combines multiple source maps into a destination map, overwriting existing keys with values from the source maps.

<table data-header-hidden><thead><tr><th width="164">Name</th><th>Value</th></tr></thead><tbody><tr><td>Signature</td><td><pre class="language-go"><code class="lang-go">MergeOverwrite(dest map[string]any, srcs ...map[string]any) (any, error)
</code></pre></td></tr></tbody></table>

{% tabs %}
{% tab title="Template Example" %}
```go
{{ $d1 := dict "a" 1 "b" 2 }}
{{ $d2 := dict "b" 3 "c" 4 }}
{{ mergeOverwrite $d1 $d2 }} // Output: {"a": 1, "b": 3, "c": 4}, nil
```
{% endtab %}
{% endtabs %}
