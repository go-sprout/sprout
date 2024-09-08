---
description: >-
  The Filesystem registry allows for efficient interaction with the file system,
  providing functions to read, write, and manipulate files directly from your
  templates.
---

# Filesystem

{% hint style="info" %}
You can easily import all the functions from the <mark style="color:yellow;">`filesystem`</mark> registry by including the following import statement in your code

```go
import "github.com/go-sprout/sprout/registry/filesystem"
```
{% endhint %}

### <mark style="color:purple;">pathBase</mark>

The function returns the last element of a given path, effectively extracting the file or directory name from the full path.

<table data-header-hidden><thead><tr><th width="174">Name</th><th>Value</th></tr></thead><tbody><tr><td>Signature</td><td><pre class="language-go"><code class="lang-go">PathBase(str string) string
</code></pre></td></tr></tbody></table>

{% tabs %}
{% tab title="Template Example" %}
```go
{{ "/path/to/file.txt" | pathBase }} // Output: "file.txt"
```
{% endtab %}
{% endtabs %}

### <mark style="color:purple;">pathDir</mark>

The function returns all but the last element of a given path, effectively extracting the directory portion of the path.

<table data-header-hidden><thead><tr><th width="164">Name</th><th>Value</th></tr></thead><tbody><tr><td>Signature</td><td><pre class="language-go"><code class="lang-go">PathDir(str string) string
</code></pre></td></tr></tbody></table>

{% tabs %}
{% tab title="Template Example" %}
```go
{{ "/" | pathDir }} // Output: "/"
{{ "/path/to/file.txt" | pathDir }} // Output: "/path/to"
```
{% endtab %}
{% endtabs %}

### <mark style="color:purple;">pathExt</mark>

The function returns the file extension of the given path, identifying the type of file by its suffix.

<table data-header-hidden><thead><tr><th width="164">Name</th><th>Value</th></tr></thead><tbody><tr><td>Signature</td><td><pre class="language-go"><code class="lang-go">PathExt(str string) string
</code></pre></td></tr></tbody></table>

{% tabs %}
{% tab title="Template Example" %}
```go
{{ "/" | pathExt }} // Output: ""
{{ "/path/to/file.txt" | pathExt }} // Output: ".txt"
```
{% endtab %}
{% endtabs %}

### <mark style="color:purple;">pathClean</mark>

The function cleans up a given path by simplifying any redundancies, such as removing unnecessary double slashes or resolving "." and ".." elements, resulting in a standardized and more straightforward path.

<table data-header-hidden><thead><tr><th width="164">Name</th><th>Value</th></tr></thead><tbody><tr><td>Signature</td><td><pre class="language-go"><code class="lang-go">PathClean(str string) string
</code></pre></td></tr></tbody></table>

{% tabs %}
{% tab title="Template Example" %}
```go
{{ "/path//to/file.txt" | pathClean }} // Output: "/path/to/file.txt"
```
{% endtab %}
{% endtabs %}

### <mark style="color:purple;">pathIsAbs</mark>

The function checks whether the given path is an absolute path, returning `true` if it is and `false` if it is not.

<table data-header-hidden><thead><tr><th width="164">Name</th><th>Value</th></tr></thead><tbody><tr><td>Signature</td><td><pre class="language-go"><code class="lang-go">PathIsAbs(str string) bool
</code></pre></td></tr></tbody></table>

{% tabs %}
{% tab title="Template Example" %}
```go
{{ "/path/to/file.txt" | pathIsAbs }} // Output: true
{{ "../file.txt" | pathIsAbs }} // Output: false
```
{% endtab %}
{% endtabs %}

### <mark style="color:purple;">osBase</mark>

The function returns the last element of a given path, using the operating system's specific path separator to determine the path structure.

<table data-header-hidden><thead><tr><th width="164">Name</th><th>Value</th></tr></thead><tbody><tr><td>Signature</td><td><pre class="language-go"><code class="lang-go">OsBase(str string) string
</code></pre></td></tr></tbody></table>

{% tabs %}
{% tab title="Template Example" %}
```go
{{ "C:\\path\\to\\file.txt" | osBase }} // Output: "file.txt"
```
{% endtab %}
{% endtabs %}

### <mark style="color:purple;">osDir</mark>

The function returns all but the last element of a given path, using the operating system's specific path separator to navigate the directory structure.

<table data-header-hidden><thead><tr><th width="164">Name</th><th>Value</th></tr></thead><tbody><tr><td>Signature</td><td><pre class="language-go"><code class="lang-go">OsDir(str string) string
</code></pre></td></tr></tbody></table>

{% tabs %}
{% tab title="Template Example" %}
```go
{{ "C:\\path\\to\\file.txt" | osDir }} // Output: "C:\\path\\to"
```
{% endtab %}
{% endtabs %}

### <mark style="color:purple;">osExt</mark>

The function returns the file extension of a given path, using the operating system's specific path separator to accurately determine the extension.

<table data-header-hidden><thead><tr><th width="164">Name</th><th>Value</th></tr></thead><tbody><tr><td>Signature</td><td><pre class="language-go"><code class="lang-go">OsExt(str string) string
</code></pre></td></tr></tbody></table>

{% tabs %}
{% tab title="Template Example" %}
```go
{{ "C:\\path\\to\\file.txt" | osExt }} // Output: ".txt"
```
{% endtab %}
{% endtabs %}

### <mark style="color:purple;">osClean</mark>

The function cleans up a given path, using the operating system's specific path separator to simplify redundancies and standardize the path structure.

<table data-header-hidden><thead><tr><th width="164">Name</th><th>Value</th></tr></thead><tbody><tr><td>Signature</td><td><pre class="language-go"><code class="lang-go">OsClean(str string) string
</code></pre></td></tr></tbody></table>

{% tabs %}
{% tab title="Template Example" %}
```go
{{ "C:\\path\\\\to\\file.txt" | osClean }} // Output: "C:\\path\\to\\file.txt"
```
{% endtab %}
{% endtabs %}

### <mark style="color:purple;">osIsAbs</mark>

The function checks whether the given path is absolute, using the operating system's specific path separator to determine if the path is absolute or relative.

<table data-header-hidden><thead><tr><th width="164">Name</th><th>Value</th></tr></thead><tbody><tr><td>Signature</td><td><pre class="language-go"><code class="lang-go">OsIsAbs(str string) bool
</code></pre></td></tr></tbody></table>

{% tabs %}
{% tab title="Template Example" %}
```go
{{ "C:\\path\\to\\file.txt" | osIsAbs }} // Output: true
```
{% endtab %}
{% endtabs %}
