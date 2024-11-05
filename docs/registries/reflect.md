---
description: >-
  The Reflect registry offers tools for inspecting and manipulating data types
  using reflection, enabling advanced dynamic type handling within your
  projects.
---

# Reflect

{% hint style="info" %}
You can easily import all the functions from the <mark style="color:yellow;">`reflect`</mark> registry by including the following import statement in your code

```go
import "github.com/go-sprout/sprout/registry/reflect"
```
{% endhint %}

### <mark style="color:purple;">typeIs</mark>

The function compares the type of the provided value (`src`) to a specified target type string (`target`). It returns `true` if the type of `src` matches the target type.

<table data-header-hidden><thead><tr><th width="174">Name</th><th>Value</th></tr></thead><tbody><tr><td>Signature</td><td><pre class="language-go"><code class="lang-go">TypeIs(target string, src any) bool
</code></pre></td></tr></tbody></table>

{% tabs %}
{% tab title="Template Example" %}
```go
{{ 42 | typeIs "int" }} // Output: true
```
{% endtab %}
{% endtabs %}

### <mark style="color:purple;">typeIsLike</mark>

The function compares the type of the provided value (`src`) to a target type string (`target`), with an option for a wildcard `*` prefix (pointer). It returns `true` if `src` matches `target` or `*target`, which is useful for checking if a variable is of a specific type or a pointer to that type.

<table data-header-hidden><thead><tr><th width="164">Name</th><th>Value</th></tr></thead><tbody><tr><td>Signature</td><td><pre class="language-go"><code class="lang-go">TypeIsLike(target string, src any) bool
</code></pre></td></tr></tbody></table>

{% tabs %}
{% tab title="Template Example" %}
```go
{{ 42 | typeIsLike "*int" }} // Output: true
```
{% endtab %}
{% endtabs %}

### <mark style="color:purple;">typeOf</mark>

The function returns the type of the provided value (`src`) as a string, giving you a textual representation of its data type.

<table data-header-hidden><thead><tr><th width="164">Name</th><th>Value</th></tr></thead><tbody><tr><td>Signature</td><td><pre class="language-go"><code class="lang-go">TypeOf(src any) string
</code></pre></td></tr></tbody></table>

{% tabs %}
{% tab title="Template Example" %}
```go
{{ 42 | typeOf }} // Output: "int"
```
{% endtab %}
{% endtabs %}

### <mark style="color:purple;">kindIs</mark>

The function compares the kind (category) of a given value (`src`) to a target kind string (`target`). It returns `true` if the kind of `src` matches the specified target kind.

<table data-header-hidden><thead><tr><th width="164">Name</th><th>Value</th></tr></thead><tbody><tr><td>Signature</td><td><pre class="language-go"><code class="lang-go">KindIs(target string, src any) (bool, error)
</code></pre></td></tr></tbody></table>

{% tabs %}
{% tab title="Template Example" %}
```go
{{ 42 | kindIs "int" }} // Output: true
{{ nil | kindIs "int" }} // Error
```
{% endtab %}
{% endtabs %}

### <mark style="color:purple;">kindOf</mark>

The function returns the kind (category) of the provided value (`src`) as a string, giving a general classification like "int," "struct," or "slice."

<table data-header-hidden><thead><tr><th width="164">Name</th><th>Value</th></tr></thead><tbody><tr><td>Signature</td><td><pre class="language-go"><code class="lang-go">KindOf(src any) (string, error)
</code></pre></td></tr></tbody></table>

{% tabs %}
{% tab title="Template Example" %}
```go
{{ 42 | kindOf }} // Output: "int"
{{ nil | kindOf }} // Error
```
{% endtab %}
{% endtabs %}

### <mark style="color:purple;">hasField</mark>

The function checks the presence of a field with the specified name (`name`) in the provided struct (`src`). It returns `true` if the field exists.

<table data-header-hidden><thead><tr><th width="164">Name</th><th>Value</th></tr></thead><tbody><tr><td>Signature</td><td><pre class="language-go"><code class="lang-go">HasField(name string, src any) (bool, error)
</code></pre></td></tr></tbody></table>

{% tabs %}
{% tab title="Template Example" %}
```go
{{ hasField "someExistingField" .someStruct }} // Output: true
{{ hasField "someNonExistingField" .someStruct }} // Output: false
{{ .someStruct | hasField "someExistingField" }} // Output: true
{{ .someStruct | hasField "someNonExistingField" }} // Output: false
```
{% endtab %}
{% endtabs %}

### <mark style="color:purple;">deepEqual</mark>

The function checks if two variables, `x` and `y`, are deeply equal by comparing their values and structures using `reflect.DeepEqual`.

<table data-header-hidden><thead><tr><th width="164">Name</th><th>Value</th></tr></thead><tbody><tr><td>Signature</td><td><pre class="language-go"><code class="lang-go">DeepEqual(x, y any) bool
</code></pre></td></tr></tbody></table>

{% tabs %}
{% tab title="Template Example" %}
```go
{{ {"a":1}, {"a":1} | deepEqual }} // Output: true
```
{% endtab %}
{% endtabs %}

### <mark style="color:purple;">deepCopy</mark>

The function performs a deep copy of the provided `element`, creating an exact duplicate of its structure and data. It uses `MustDeepCopy` internally to manage the copy process and handle any potential errors. This use the [copystructure package](https://github.com/mitchellh/copystructure) internally.

<table data-header-hidden><thead><tr><th width="164">Name</th><th>Value</th></tr></thead><tbody><tr><td>Signature</td><td><pre class="language-go"><code class="lang-go">DeepCopy(element any) (any, error)
</code></pre></td></tr></tbody></table>

{% tabs %}
{% tab title="Template Example" %}
```go
{{ {"name":"John"} | deepCopy }} // Output: {"name":"John"}
{{ nil | deepCopy }} // Output: nil, error
```
{% endtab %}
{% endtabs %}
