---
description: >-
  The Encoding registry offers methods for encoding and decoding data in
  different formats, allowing for flexible data representation and storage
  within your templates.
---

# Encoding

{% hint style="info" %}
You can easily import all the functions from the <mark style="color:yellow;">`encoding`</mark> registry by including the following import statement in your code

```go
import "github.com/go-sprout/sprout/registry/encoding"
```
{% endhint %}

### <mark style="color:purple;">base64Encode</mark>

The function encodes a given string into its Base64 representation, converting the data into a text format suitable for transmission or storage in systems that support Base64 encoding.

<table data-header-hidden><thead><tr><th width="174">Name</th><th>Value</th></tr></thead><tbody><tr><td>Signature</td><td><pre class="language-go"><code class="lang-go">Base64Encode(s string) string
</code></pre></td></tr></tbody></table>

{% tabs %}
{% tab title="Template Example" %}
```go
{{ "Hello World!" | base64Encode }} // Output: "SGVsbG8gV29ybGQh"
```
{% endtab %}
{% endtabs %}

### <mark style="color:purple;">base64Decode</mark>

The function decodes a Base64 encoded string back to its original form. If the input string is not valid Base64, it returns an error message.

<table data-header-hidden><thead><tr><th width="164">Name</th><th>Value</th></tr></thead><tbody><tr><td>Signature</td><td><pre class="language-go"><code class="lang-go"> Base64Decode(s string) (string, error)
</code></pre></td></tr></tbody></table>

{% tabs %}
{% tab title="Template Example" %}
```go
{{ "SGVsbG8gV29ybGQh" | base64Decode }} // Output: "Hello World!"
```
{% endtab %}
{% endtabs %}

### <mark style="color:purple;">base32Encode</mark>

The function encodes a given string into its Base32 representation, converting the data into a text format that is compatible with systems supporting Base32 encoding.

<table data-header-hidden><thead><tr><th width="164">Name</th><th>Value</th></tr></thead><tbody><tr><td>Signature</td><td><pre class="language-go"><code class="lang-go">Base32Encode(s string) string
</code></pre></td></tr></tbody></table>

{% tabs %}
{% tab title="Template Example" %}
```go
{{ "Hello World!" | base32Encode }} // Output: "JBSWY3DPEBLW64TMMQQQ===="
```
{% endtab %}
{% endtabs %}

### <mark style="color:purple;">base32Decode</mark>

The function decodes a Base32 encoded string back to its original form. If the input is not valid Base32, it returns an error message.

<table data-header-hidden><thead><tr><th width="164">Name</th><th>Value</th></tr></thead><tbody><tr><td>Signature</td><td><pre class="language-go"><code class="lang-go">Base32Decode(s string) (string, error)
</code></pre></td></tr></tbody></table>

{% tabs %}
{% tab title="Template Example" %}
```go
{{ "JBSWY3DPEBLW64TMMQQQ====" | base32Decode }} // Output: "Hello World!"
```
{% endtab %}
{% endtabs %}

### <mark style="color:purple;">fromJSON</mark>

The function converts a JSON string into a corresponding Go data structure, enabling easy manipulation of the JSON data in a Go environment.

<table data-header-hidden><thead><tr><th width="164">Name</th><th>Value</th></tr></thead><tbody><tr><td>Signature</td><td><pre class="language-go"><code class="lang-go">FromJSON(v string) (any, error)
</code></pre></td></tr></tbody></table>

{% tabs %}
{% tab title="Template Example" %}
```go
{{ `{"name":"John", "age":30}` | fromJSON }} // Output: map[age:30 name:John]
{{ "{\\invalid" | fromJSON }} // Error
```
{% endtab %}
{% endtabs %}

### <mark style="color:purple;">toJSON</mark>

The function converts a Go data structure into a JSON string, allowing the data to be easily serialized for storage, transmission, or further processing.

<table data-header-hidden><thead><tr><th width="164">Name</th><th>Value</th></tr></thead><tbody><tr><td>Signature</td><td><pre class="language-go"><code class="lang-go">ToJSON(v any) (string, error)
</code></pre></td></tr></tbody></table>

{% tabs %}
{% tab title="Template Example" %}
```go
{{- $d := dict "key1" "value1" "key2" "value2" "key3" "value3" -}}
{{ toJSON $d }} // Output: {\"key1\":\"value1\",\"key2\":\"value2\",\"key3\":\"value3\"}
```
{% endtab %}
{% endtabs %}

### <mark style="color:purple;">toPrettyJSON</mark>

The function converts a Go data structure into a pretty-printed JSON string, formatting the output with indentation and line breaks for better readability.

<table data-header-hidden><thead><tr><th width="164">Name</th><th>Value</th></tr></thead><tbody><tr><td>Signature</td><td><pre class="language-go"><code class="lang-go">ToPrettyJSON(v any) (string, error)
</code></pre></td></tr></tbody></table>

{% tabs %}
{% tab title="Template Example" %}
```go
{{- $d := dict "key1" "value1" "key2" "value2" "key3" "value3" -}}
{{ toPrettyJSON $d }} // Output: "{\n  \"key1\": \"value1\",\n  \"key2\": \"value2\",\n  \"key3\": \"value3\"\n}"
```
{% endtab %}
{% endtabs %}

### <mark style="color:purple;">toRawJSON</mark>

The function converts a Go data structure into a JSON string without escaping HTML characters, preserving the raw content as it is.

<table data-header-hidden><thead><tr><th width="164">Name</th><th>Value</th></tr></thead><tbody><tr><td>Signature</td><td><pre class="language-go"><code class="lang-go">ToRawJSON(v any) (string, error)
</code></pre></td></tr></tbody></table>

{% tabs %}
{% tab title="Template Example" %}
```go
{{- $d := dict "content" "<p>Hello World</p>" -}}
{{ toRawJSON $d }} // Output: {\"content\":\"<p>Hello World</p>\"}
```
{% endtab %}
{% endtabs %}

### <mark style="color:purple;">fromYAML</mark>

The function deserializes a YAML string into a Go map, allowing the structured data from YAML to be used and manipulated within a Go program.

<table data-header-hidden><thead><tr><th width="164">Name</th><th>Value</th></tr></thead><tbody><tr><td>Signature</td><td><pre class="language-go"><code class="lang-go">FromYAML(v string) (any, error)
</code></pre></td></tr></tbody></table>

{% tabs %}
{% tab title="Template Example" %}
```go

{{ "name: John Doe\nage: 30" | fromYAML }} // Output: map[age:30 name:John Doe]
```
{% endtab %}
{% endtabs %}

### <mark style="color:purple;">toYAML</mark>

The function serializes a Go data structure into a YAML string, converting the data into a format suitable for YAML representation.

<table data-header-hidden><thead><tr><th width="164">Name</th><th>Value</th></tr></thead><tbody><tr><td>Signature</td><td><pre class="language-go"><code class="lang-go">ToYAML(v any) (string, error)
</code></pre></td></tr></tbody></table>

{% tabs %}
{% tab title="Template Example" %}
```go
{{- $d := dict "name" "John Doe" "age" 30 -}}
{{ $d | toYAML }} // Output: age: 30\nname: John Doe
```
{% endtab %}
{% endtabs %}

### <mark style="color:purple;">toIndentYAML</mark>

The function serializes a Go data structure into a YAML string, converting the data into a format suitable for YAML representation. In addition to toYAML, toIndentYAML takes a parameter to define the indentation width in spaces.

<table data-header-hidden><thead><tr><th width="164">Name</th><th>Value</th></tr></thead><tbody><tr><td>Signature</td><td><pre class="language-go"><code class="lang-go">ToIndentYAML(indent int, v any) (string, error)
</code></pre></td></tr></tbody></table>

{% tabs %}
{% tab title="Template Example" %}
```go
{{ $person := dict "name" "John Doe" "age" 30 "location" (dict "country" "US" "planet" "Earth") }}
{{ $person | toIndentYAML 2 }} // Output: \nage: 30\nlocation:\n  country: US\n  planet: Earth\nname: John Doe
```
{% endtab %}
{% endtabs %}


