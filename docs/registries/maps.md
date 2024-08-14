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
</code></pre></td></tr><tr><td>Must version</td><td><span data-gb-custom-inline data-tag="emoji" data-code="274c">❌</span></td></tr></tbody></table>

{% tabs %}
{% tab title="Template Example" %}
```go
{{ dict "key1", "value1", "key2", "value2" }}
// Output: {"key1": "value1", "key2": "value2"}
```
{% endtab %}
{% endtabs %}

### <mark style="color:purple;">get</mark>

The function retrieves the value associated with a specified key from a dictionary (map). If the key is found, the corresponding value is returned.

<table data-header-hidden><thead><tr><th width="164">Name</th><th>Value</th></tr></thead><tbody><tr><td>Signature</td><td><pre class="language-go"><code class="lang-go">Get(dict map[string]any, key string) any
</code></pre></td></tr><tr><td>Must version</td><td><span data-gb-custom-inline data-tag="emoji" data-code="274c">❌</span></td></tr></tbody></table>

{% tabs %}
{% tab title="Template Example" %}
```go
{{ get {"key": "value"}, "key" }} // Output: "value"
{{ get {"key": "value"}, "invalid" }} // Output: ""
```
{% endtab %}
{% endtabs %}

### <mark style="color:purple;">set</mark>

The function adds a new key-value pair to a dictionary or updates the value associated with an existing key.

<table data-header-hidden><thead><tr><th width="164">Name</th><th>Value</th></tr></thead><tbody><tr><td>Signature</td><td><pre class="language-go"><code class="lang-go">Set(dict map[string]any, key string, value any) map[string]any
</code></pre></td></tr><tr><td>Must version</td><td><span data-gb-custom-inline data-tag="emoji" data-code="274c">❌</span></td></tr></tbody></table>

{% tabs %}
{% tab title="Template Example" %}
```go
{{ set {"key": "oldValue"}, "key", "newValue" }} // Output: {"key": "newValue"}
{{ set {"foo": "bar"}, "far", "boo" }} // Output: {"foo": "bar", "far": "boo"}
```
{% endtab %}
{% endtabs %}

### <mark style="color:purple;">unset</mark>

```go
{{ {"key": "value"}, "key" | unset }} // Output: {}
```

<table data-header-hidden><thead><tr><th width="164">Name</th><th>Value</th></tr></thead><tbody><tr><td>Signature</td><td><pre class="language-go"><code class="lang-go">Unset(dict map[string]any, key string) map[string]any
</code></pre></td></tr><tr><td>Must version</td><td><span data-gb-custom-inline data-tag="emoji" data-code="274c">❌</span></td></tr></tbody></table>

{% tabs %}
{% tab title="Template Example" %}
```go
{{ {"key": "value"}, "key" | unset }} // Output: {}
{{ {"key": "value"}, "invalid" | unset }} // Output: {"key": "value"}
```
{% endtab %}
{% endtabs %}

### <mark style="color:purple;">keys</mark>

The function retrieves all keys from one or more dictionaries, returning them as a list.

<table data-header-hidden><thead><tr><th width="164">Name</th><th>Value</th></tr></thead><tbody><tr><td>Signature</td><td><pre class="language-go"><code class="lang-go">Keys(dicts ...map[string]any) []string
</code></pre></td></tr><tr><td>Must version</td><td><span data-gb-custom-inline data-tag="emoji" data-code="274c">❌</span></td></tr></tbody></table>

{% tabs %}
{% tab title="Template Example" %}
```go
{{ keys {"key1": "value1", "key2": "value2"} }} // Output: ["key1", "key2"]
{{ keys {"key1": "value1"} {"key2": "value2"} }} // Output: ["key1", "key2"]
```
{% endtab %}
{% endtabs %}

### <mark style="color:purple;">values</mark>

The function retrieves all values from a dictionary, returning them as a list.

<table data-header-hidden><thead><tr><th width="164">Name</th><th>Value</th></tr></thead><tbody><tr><td>Signature</td><td><pre class="language-go"><code class="lang-go">Values(dict map[string]any) []any
</code></pre></td></tr><tr><td>Must version</td><td><span data-gb-custom-inline data-tag="emoji" data-code="274c">❌</span></td></tr></tbody></table>

{% tabs %}
{% tab title="Template Example" %}
```go
{{ values {"key1": "value1", "key2": "value2"} }} // Output: ["value1", "value2"]
```
{% endtab %}
{% endtabs %}

### <mark style="color:purple;">pluck</mark>

The function extracts values associated with a specified key from a list of dictionaries, returning a list of those values.

<table data-header-hidden><thead><tr><th width="164">Name</th><th>Value</th></tr></thead><tbody><tr><td>Signature</td><td><pre class="language-go"><code class="lang-go">Pluck(key string, dicts ...map[string]any) []any
</code></pre></td></tr><tr><td>Must version</td><td><span data-gb-custom-inline data-tag="emoji" data-code="274c">❌</span></td></tr></tbody></table>

{% tabs %}
{% tab title="Template Example" %}
<pre class="language-go"><code class="lang-go"><strong>{{ $d1 := dict "key" "value1"}}
</strong>{{ $d2 := dict "key" "value2" }}
{{ pluck "key"	$d1 $d2 }} // Output: ["value1", "value2"]
</code></pre>
{% endtab %}
{% endtabs %}

### <mark style="color:purple;">pick</mark>

The function creates a new dictionary that includes only the specified keys from the original dictionary, effectively filtering out all other keys and their associated values.

<table data-header-hidden><thead><tr><th width="164">Name</th><th>Value</th></tr></thead><tbody><tr><td>Signature</td><td><pre class="language-go"><code class="lang-go">Pick(dict map[string]any, keys ...string) map[string]any
</code></pre></td></tr><tr><td>Must version</td><td><span data-gb-custom-inline data-tag="emoji" data-code="274c">❌</span></td></tr></tbody></table>

{% tabs %}
{% tab title="Template Example" %}
```go
{{ $d := dict "key1" "value1" "key2" "value2" "key3" "value3" }}
{{ pick $d "key1" "key3" }} // Output: {"key1": "value1", "key3": "value3"}
```
{% endtab %}
{% endtabs %}

### <mark style="color:purple;">omit</mark>

The function creates a new dictionary by excluding the specified keys from the original dictionary, effectively removing those key-value pairs from the resulting dictionary.

<table data-header-hidden><thead><tr><th width="164">Name</th><th>Value</th></tr></thead><tbody><tr><td>Signature</td><td><pre class="language-go"><code class="lang-go">Omit(dict map[string]any, keys ...string) map[string]any
</code></pre></td></tr><tr><td>Must version</td><td><span data-gb-custom-inline data-tag="emoji" data-code="274c">❌</span></td></tr></tbody></table>

{% tabs %}
{% tab title="Template Example" %}
```go
{{ $d := dict "key1" "value1" "key2" "value2" "key3" "value3" }}
{{ omit $d "key1" "key3" }} // Output: {"key2": "value2"}
```
{% endtab %}
{% endtabs %}

### <mark style="color:purple;">dig</mark>

The function navigates through a nested dictionary structure using a sequence of keys and returns the value found at the specified path, allowing access to deeply nested data. The last argument must be the map.

<table data-header-hidden><thead><tr><th width="164">Name</th><th>Value</th></tr></thead><tbody><tr><td>Signature</td><td><pre class="language-go"><code class="lang-go">Dig(args ...any) (any, error)
</code></pre></td></tr><tr><td>Must version</td><td><span data-gb-custom-inline data-tag="emoji" data-code="274c">❌</span></td></tr></tbody></table>

{% tabs %}
{% tab title="Template Example" %}
```go
{{ $nest := dict "foo" "bar" }}
{{ $d := dict "key1" "value1" "nested" $nest }}
{{ $d | dig "nested" "foo" }} // Output: "bar"
```
{% endtab %}
{% endtabs %}

### <mark style="color:purple;">hasKey</mark>

The function checks whether a specified key exists in the dictionary, returning `true` if the key is found and `false` otherwise.

<table data-header-hidden><thead><tr><th width="164">Name</th><th>Value</th></tr></thead><tbody><tr><td>Signature</td><td><pre class="language-go"><code class="lang-go">HasKey(dict map[string]any, key string) bool
</code></pre></td></tr><tr><td>Must version</td><td><span data-gb-custom-inline data-tag="emoji" data-code="274c">❌</span></td></tr></tbody></table>

{% tabs %}
{% tab title="Template Example" %}
```go
{{ $d := dict "key1" "value1" }}
{{ hasKey $d "key1" }} // Output: true
{{ hasKey $d "key2" }} // Output: false
```
{% endtab %}
{% endtabs %}

### <mark style="color:purple;">merge / mustMerge</mark>

The function combines multiple source maps into a single destination map, adding new key-value pairs without overwriting any existing keys in the destination map.

<table data-header-hidden><thead><tr><th width="164">Name</th><th>Value</th></tr></thead><tbody><tr><td>Signature</td><td><pre class="language-go"><code class="lang-go">Merge(dest map[string]any, srcs ...map[string]any) any
MustMerge(dest map[string]any, srcs ...map[string]any) (any, error)
</code></pre></td></tr><tr><td>Must version</td><td><span data-gb-custom-inline data-tag="emoji" data-code="2705">✅</span></td></tr></tbody></table>

{% tabs %}
{% tab title="Template Example" %}
<pre class="language-go"><code class="lang-go"><strong>{{ $d1 := dict "a" 1 "b" 2 }}
</strong><strong>{{ $d2 := dict "b" 3 "c" 4 }}
</strong><strong>{{ merge $d1, $d2 }} // Output: {"a": 1, "b": 2, "c": 4}, nil
</strong></code></pre>
{% endtab %}

{% tab title="Must version" %}
<pre class="language-go"><code class="lang-go"><strong>{{ $d1 := dict "a" 1 "b" 2 }}
</strong><strong>{{ $d2 := dict "b" 3 "c" 4 }}
</strong><strong>{{ mustMerge $d1, $d2 }} // Output: {"a": 1, "b": 2, "c": 4}, nil
</strong></code></pre>
{% endtab %}
{% endtabs %}

### <mark style="color:purple;">mergeOverwrite / mustMergeOverwrite</mark>

The function combines multiple source maps into a destination map, overwriting existing keys with values from the source maps.

<table data-header-hidden><thead><tr><th width="164">Name</th><th>Value</th></tr></thead><tbody><tr><td>Signature</td><td><pre class="language-go"><code class="lang-go">MergeOverwrite(dest map[string]any, srcs ...map[string]any) any
MustMergeOverwrite(dest map[string]any, srcs ...map[string]any) (any, error)
</code></pre></td></tr><tr><td>Must version</td><td><span data-gb-custom-inline data-tag="emoji" data-code="2705">✅</span></td></tr></tbody></table>

{% tabs %}
{% tab title="Template Example" %}
```go
{{ $d1 := dict "a" 1 "b" 2 }}
{{ $d2 := dict "b" 3 "c" 4 }}
{{ mergeOverwrite $d1, $d2 }} // Output: {"a": 1, "b": 3, "c": 4}, nil
```
{% endtab %}

{% tab title="Must version" %}
<pre class="language-go"><code class="lang-go"><strong>{{ $d1 := dict "a" 1 "b" 2 }}
</strong><strong>{{ $d2 := dict "b" 3 "c" 4 }}
</strong><strong>{{ mustMergeOverwrite $d1, $d2 }} // Output: {"a": 1, "b": 3, "c": 4}, nil
</strong></code></pre>
{% endtab %}
{% endtabs %}
