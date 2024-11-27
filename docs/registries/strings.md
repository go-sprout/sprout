---
description: >-
  The Strings registry offers a comprehensive set of functions for manipulating
  strings, including formatting, splitting, joining, and other common string
  operations.
---

# Strings

{% hint style="info" %}
You can easily import all the functions from the <mark style="color:yellow;">`strings`</mark> registry by including the following import statement in your code

```go
import "github.com/go-sprout/sprout/registry/strings"
```
{% endhint %}

### <mark style="color:purple;">nospace</mark>

The function removes all whitespace characters from the provided string, eliminating any spaces, tabs, or line breaks.

<table data-header-hidden><thead><tr><th width="174">Name</th><th>Value</th></tr></thead><tbody><tr><td>Signature</td><td><pre class="language-go"><code class="lang-go">Nospace(str string) string
</code></pre></td></tr></tbody></table>

{% tabs %}
{% tab title="Template Example" %}
```go
{{ "Hello World" | nospace }} // Output: "HelloWorld"
```
{% endtab %}
{% endtabs %}

### <mark style="color:purple;">trim</mark>

The function removes any leading and trailing whitespace from the provided string.

<table data-header-hidden><thead><tr><th width="164">Name</th><th>Value</th></tr></thead><tbody><tr><td>Signature</td><td><pre class="language-go"><code class="lang-go">Trim(str string) string
</code></pre></td></tr></tbody></table>

{% tabs %}
{% tab title="Template Example" %}
```go
{{ " Hello World " | trim }} // Output: "Hello World"
```
{% endtab %}
{% endtabs %}

### <mark style="color:purple;">trimAll</mark>

The function removes all instances of any characters in the 'cutset' from both the beginning and the end of the provided string.

<table data-header-hidden><thead><tr><th width="164">Name</th><th>Value</th></tr></thead><tbody><tr><td>Signature</td><td><pre class="language-go"><code class="lang-go">TrimAll(cutset string, str string)
</code></pre></td></tr></tbody></table>

{% tabs %}
{% tab title="Template Example" %}
```go
{{ "xyzHelloxyz" | trimAll "xyz" }} // Output: "Hello"
```
{% endtab %}
{% endtabs %}

### <mark style="color:purple;">trimPrefix</mark>

The function removes the specified 'prefix' from the start of the provided string if it is present.

<table data-header-hidden><thead><tr><th width="164">Name</th><th>Value</th></tr></thead><tbody><tr><td>Signature</td><td><pre class="language-go"><code class="lang-go">TrimPrefix(prefix string, str string) string
</code></pre></td></tr></tbody></table>

{% tabs %}
{% tab title="Template Example" %}
<pre class="language-go"><code class="lang-go"><strong>{{ "HelloWorld" | trimPrefix "Hello" }} // Output: "World"
</strong>{{ "HelloWorld" | trimPrefix "World" }} // Output: "HelloWorld"
</code></pre>
{% endtab %}
{% endtabs %}

### <mark style="color:purple;">trimSuffix</mark>

The function removes the specified 'suffix' from the end of the provided string if it is present.

<table data-header-hidden><thead><tr><th width="164">Name</th><th>Value</th></tr></thead><tbody><tr><td>Signature</td><td><pre class="language-go"><code class="lang-go">TrimSuffix(suffix string, str string) string
</code></pre></td></tr></tbody></table>

{% tabs %}
{% tab title="Template Example" %}
```go
{{ "HelloWorld" | trimSuffix "Hello" }} // Output: "HelloWorld"
{{ "HelloWorld" | trimSuffix "World" }} // Output: "Hello"
```
{% endtab %}
{% endtabs %}

### <mark style="color:purple;">contains</mark>

The function checks whether the provided string contains the specified substring.

<table data-header-hidden><thead><tr><th width="164">Name</th><th>Value</th></tr></thead><tbody><tr><td>Signature</td><td><pre class="language-go"><code class="lang-go">Contains(substring string, str string) bool
</code></pre></td></tr></tbody></table>

{% tabs %}
{% tab title="Template Example" %}
```go
{{ "Hello" | contains "ell" }} // Output: true
```
{% endtab %}
{% endtabs %}

### <mark style="color:purple;">hasPrefix</mark>

The function checks whether the provided string starts with the specified prefix.

<table data-header-hidden><thead><tr><th width="164">Name</th><th>Value</th></tr></thead><tbody><tr><td>Signature</td><td><pre class="language-go"><code class="lang-go">HasPrefix(prefix string, str string) bool
</code></pre></td></tr></tbody></table>

{% tabs %}
{% tab title="Template Example" %}
```go
{{ "HelloWorld" | hasPrefix "Hello" }} // Output: true
```
{% endtab %}
{% endtabs %}

### <mark style="color:purple;">hasSuffix</mark>

The function checks whether the provided string ends with the specified suffix.

<table data-header-hidden><thead><tr><th width="164">Name</th><th>Value</th></tr></thead><tbody><tr><td>Signature</td><td><pre class="language-go"><code class="lang-go">HasSuffix(suffix string, str string) bool
</code></pre></td></tr></tbody></table>

{% tabs %}
{% tab title="Template Example" %}
```go
{{ "HelloWorld" | hasSuffix "World" }} // Output: true
```
{% endtab %}
{% endtabs %}

### <mark style="color:purple;">toLower</mark>

The function converts all characters in the provided string to lowercase.

<table data-header-hidden><thead><tr><th width="164">Name</th><th>Value</th></tr></thead><tbody><tr><td>Signature</td><td><pre class="language-go"><code class="lang-go">ToLower(str string) string
</code></pre></td></tr></tbody></table>

{% tabs %}
{% tab title="Template Example" %}
```go
{{ "HELLO WORLD" | toLower }} // Output: "hello world"
```
{% endtab %}
{% endtabs %}

### <mark style="color:purple;">toUpper</mark>

The function converts all characters in the provided string to uppercase.

<table data-header-hidden><thead><tr><th width="164">Name</th><th>Value</th></tr></thead><tbody><tr><td>Signature</td><td><pre class="language-go"><code class="lang-go">ToUpper(str string) string
</code></pre></td></tr></tbody></table>

{% tabs %}
{% tab title="Template Example" %}
```go
{{ "hello world" | toUpper }} // Output: "HELLO WORLD"
```
{% endtab %}
{% endtabs %}

### <mark style="color:purple;">replace</mark>

The function replaces all occurrences of a specified substring ('old') in the source string with a new substring ('new').

<table data-header-hidden><thead><tr><th width="164">Name</th><th>Value</th></tr></thead><tbody><tr><td>Signature</td><td><pre class="language-go"><code class="lang-go">Replace(old string, new string, src string) string
</code></pre></td></tr></tbody></table>

{% tabs %}
{% tab title="Template Example" %}
```go
{{ "banana" | replace "a" "o" }} // Output: "bonono"
```
{% endtab %}
{% endtabs %}

### <mark style="color:purple;">repeat</mark>

The function repeats the provided string a specified number of times.

<table data-header-hidden><thead><tr><th width="164">Name</th><th>Value</th></tr></thead><tbody><tr><td>Signature</td><td><pre class="language-go"><code class="lang-go">Repeat(count int, str string) string
</code></pre></td></tr></tbody></table>

{% tabs %}
{% tab title="Template Example" %}
```go
{{ "ha" | repeat 3 }} // Output: "hahaha"
```
{% endtab %}
{% endtabs %}

### <mark style="color:purple;">join</mark>

The function concatenates elements of a slice into a single string, with each element separated by a specified delimiter. It can convert various slice types to a slice of strings if needed before joining.

<table data-header-hidden><thead><tr><th width="164">Name</th><th>Value</th></tr></thead><tbody><tr><td>Signature</td><td><pre class="language-go"><code class="lang-go">Join(sep string, v any) string
</code></pre></td></tr></tbody></table>

{% tabs %}
{% tab title="Template Example" %}
```go
{{ $list := slice "apple" "banana" "cherry" }}
{{ $list | join ", " }} // Output: "apple, banana, cherry"
```
{% endtab %}
{% endtabs %}

### <mark style="color:purple;">trunc</mark>

The function truncates the provided string to a maximum specified length. If the length is negative, it removes the specified number of characters from the beginning of the string.

<table data-header-hidden><thead><tr><th width="164">Name</th><th>Value</th></tr></thead><tbody><tr><td>Signature</td><td><pre class="language-go"><code class="lang-go">Trunc(count int, str string) string
</code></pre></td></tr></tbody></table>

{% tabs %}
{% tab title="Template Example" %}
```go
{{ "Hello World" | trunc 5 }} // Output: "Hello"
{{ "Hello World" | trunc -5 }} // Output: "World"
```
{% endtab %}
{% endtabs %}

### <mark style="color:purple;">shuffle</mark>

The function randomly rearranges the characters in the provided string, producing a shuffled version of the original string.

<table data-header-hidden><thead><tr><th width="164">Name</th><th>Value</th></tr></thead><tbody><tr><td>Signature</td><td><pre class="language-go"><code class="lang-go">Shuffle(str string) string
</code></pre></td></tr></tbody></table>

{% tabs %}
{% tab title="Template Example" %}
```go
{{ "hello" | shuffle }} // Output: "loleh" (output may vary due to randomness)
```
{% endtab %}
{% endtabs %}

### <mark style="color:purple;">ellipsis</mark>

The function truncates a string to a specified maximum width and appends an ellipsis ("...") if the string exceeds that width.

<table data-header-hidden><thead><tr><th width="164">Name</th><th>Value</th></tr></thead><tbody><tr><td>Signature</td><td><pre class="language-go"><code class="lang-go">Ellipsis(maxWidth int, str string) string
</code></pre></td></tr></tbody></table>

{% tabs %}
{% tab title="Template Example" %}
```go
{{ "Hello World" | ellipsis 10 }} // Output: "Hello W..."
```
{% endtab %}
{% endtabs %}

### <mark style="color:purple;">ellipsisBoth</mark>

The function truncates a string from both ends, preserving the middle portion and adding ellipses ("...") to both ends if the string exceeds the specified length.

<table data-header-hidden><thead><tr><th width="164">Name</th><th>Value</th></tr></thead><tbody><tr><td>Signature</td><td><pre class="language-go"><code class="lang-go">EllipsisBoth(offset int, maxWidth int, str string) string
</code></pre></td></tr></tbody></table>

{% tabs %}
{% tab title="Template Example" %}
```go
{{ "Hello World" | ellipsisBoth 1 10 }} // Output: "...lo Wor..."
```
{% endtab %}
{% endtabs %}

### <mark style="color:purple;">initials</mark>

The function extracts initials from a string, optionally using specified delimiters to identify word boundaries.

<table data-header-hidden><thead><tr><th width="164">Name</th><th>Value</th></tr></thead><tbody><tr><td>Signature</td><td><pre class="language-go"><code class="lang-go">Initials(str string) string
</code></pre></td></tr></tbody></table>

{% tabs %}
{% tab title="Template Example" %}
```go
{{ "John Doe" | initials }} // Output: "JD"
```
{% endtab %}
{% endtabs %}

### <mark style="color:purple;">plural</mark>

The function returns a specified string ('one') if the count is 1; otherwise, it returns an alternative string ('many').

<table data-header-hidden><thead><tr><th width="164">Name</th><th>Value</th></tr></thead><tbody><tr><td>Signature</td><td><pre class="language-go"><code class="lang-go">Plural(one, many string, count int) string
</code></pre></td></tr></tbody></table>

{% tabs %}
{% tab title="Template Example" %}
```go
{{ 1 | plural "apple" "apples" }} // Output: "apple"
{{ 2 | plural "apple" "apples" }} // Output: "apples"
```
{% endtab %}
{% endtabs %}

### <mark style="color:purple;">wrap</mark>

The function breaks a string into lines, ensuring that each line does not exceed a specified maximum length. It avoids splitting words across lines unless absolutely necessary.

<table data-header-hidden><thead><tr><th width="164">Name</th><th>Value</th></tr></thead><tbody><tr><td>Signature</td><td><pre class="language-go"><code class="lang-go">Wrap(length int, str string) string
</code></pre></td></tr></tbody></table>

{% tabs %}
{% tab title="Template Example" %}
```go
{{ "This is a long string that needs to be wrapped." | wrap 10 }}
// Output: "This is a\nlong\nstring\nthat needs\nto be\nwrapped."
```
{% endtab %}
{% endtabs %}

### <mark style="color:purple;">wrapWith</mark>

The function breaks a string into lines with a specified maximum length, using a custom newline character to separate the lines. It only wraps words when they exceed the maximum line length.

<table data-header-hidden><thead><tr><th width="164">Name</th><th>Value</th></tr></thead><tbody><tr><td>Signature</td><td><pre class="language-go"><code class="lang-go">WrapWith(length int, newLineCharacter string, str string) string
</code></pre></td></tr></tbody></table>

{% tabs %}
{% tab title="Template Example" %}
```go
{{ "This is a long string that needs to be wrapped." | wrapWith 10 "<br>" }}
// Output: "This is a<br>long<br>string<br>that needs<br>to be<br>wrapped."
```
{% endtab %}
{% endtabs %}

### <mark style="color:purple;">quote</mark>

The function wraps each element in a provided list with double quotes and separates them with spaces.

<table data-header-hidden><thead><tr><th width="164">Name</th><th>Value</th></tr></thead><tbody><tr><td>Signature</td><td><pre class="language-go"><code class="lang-go">Quote(elements ...any) string
</code></pre></td></tr></tbody></table>

{% tabs %}
{% tab title="Template Example" %}
```go
{{ $list := slice "hello" "world" 123 }}
{{ $list | quote }}
// Output: "hello" "world" "123"
```
{% endtab %}
{% endtabs %}

### <mark style="color:purple;">squote</mark>

The function wraps each element in the provided list with single quotes and separates them with spaces.

<table data-header-hidden><thead><tr><th width="164">Name</th><th>Value</th></tr></thead><tbody><tr><td>Signature</td><td><pre class="language-go"><code class="lang-go">Squote(elements ...any) string
</code></pre></td></tr></tbody></table>

{% tabs %}
{% tab title="Template Example" %}
```go
{{ $list := slice "hello" "world" 123 }}
{{ $list | squote }}
// Output: 'hello' 'world' '123'
```
{% endtab %}
{% endtabs %}

### <mark style="color:purple;">toCamelCase</mark>

Converts a string to `camelCase` format.

<table data-header-hidden><thead><tr><th width="164">Name</th><th>Value</th></tr></thead><tbody><tr><td>Signature</td><td><pre class="language-go"><code class="lang-go">ToCamelCase(str string) string
</code></pre></td></tr></tbody></table>

{% tabs %}
{% tab title="Template Example" %}
```go
{{ "hello world" | toCamelCase }} // Output: "helloWorld"
```
{% endtab %}
{% endtabs %}

### <mark style="color:purple;">toKebabCase</mark>

Converts a string to `kebab-case` format.

<table data-header-hidden><thead><tr><th width="164">Name</th><th>Value</th></tr></thead><tbody><tr><td>Signature</td><td><pre class="language-go"><code class="lang-go">ToKebabCase(str string) string
</code></pre></td></tr></tbody></table>

{% tabs %}
{% tab title="Template Example" %}
```go
{{ "hello world" | toKebabCase }} // Output: "hello-world"
```
{% endtab %}
{% endtabs %}

### <mark style="color:purple;">toPascalCase</mark>

Converts a string to `PascalCase` format.

<table data-header-hidden><thead><tr><th width="164">Name</th><th>Value</th></tr></thead><tbody><tr><td>Signature</td><td><pre class="language-go"><code class="lang-go">ToPascalCase(str string) string
</code></pre></td></tr></tbody></table>

{% tabs %}
{% tab title="Template Example" %}
```go
{{ "hello world" | toPascalCase }} // Output: "HelloWorld"
```
{% endtab %}
{% endtabs %}

### <mark style="color:purple;">toDotCase</mark>

Converts a string to `dot.case` format.

<table data-header-hidden><thead><tr><th width="164">Name</th><th>Value</th></tr></thead><tbody><tr><td>Signature</td><td><pre class="language-go"><code class="lang-go">ToDotCase(str string) string
</code></pre></td></tr></tbody></table>

{% tabs %}
{% tab title="Template Example" %}
```go
{{ "hello world" | toDotCase }} // Output: "hello.world"
```
{% endtab %}
{% endtabs %}

### <mark style="color:purple;">toPathCase</mark>

Converts a string to `path/case` format.

<table data-header-hidden><thead><tr><th width="164">Name</th><th>Value</th></tr></thead><tbody><tr><td>Signature</td><td><pre class="language-go"><code class="lang-go">ToPathCase(str string) string
</code></pre></td></tr></tbody></table>

{% tabs %}
{% tab title="Template Example" %}
```go
{{ "hello world" | toPathCase }} // Output: "hello/world"
```
{% endtab %}
{% endtabs %}

### <mark style="color:purple;">toConstantCase</mark>

Converts a string to `CONSTANT_CASE` format.

<table data-header-hidden><thead><tr><th width="164">Name</th><th>Value</th></tr></thead><tbody><tr><td>Signature</td><td><pre class="language-go"><code class="lang-go">ToConstantCase(str string) string
</code></pre></td></tr></tbody></table>

{% tabs %}
{% tab title="Template Example" %}
```go
{{ "hello world" | toConstantCase }} // Output: "HELLO_WORLD"
```
{% endtab %}
{% endtabs %}

### <mark style="color:purple;">toSnakeCase</mark>

Converts a string to `snake_case` format.

<table data-header-hidden><thead><tr><th width="164">Name</th><th>Value</th></tr></thead><tbody><tr><td>Signature</td><td><pre class="language-go"><code class="lang-go">ToSnakeCase(str string) string
</code></pre></td></tr></tbody></table>

{% tabs %}
{% tab title="Template Example" %}
```go
{{ "hello world" | toSnakeCase }} // Output: "hello_world"
```
{% endtab %}
{% endtabs %}

### <mark style="color:purple;">toTitleCase</mark>

Converts a string to `Title Case` format.

<table data-header-hidden><thead><tr><th width="164">Name</th><th>Value</th></tr></thead><tbody><tr><td>Signature</td><td><pre class="language-go"><code class="lang-go">ToTitleCase(str string) string
</code></pre></td></tr></tbody></table>

{% tabs %}
{% tab title="Template Example" %}
```go
{{ "hello world" | toTitleCase }} // Output: "Hello World"
```
{% endtab %}
{% endtabs %}

### <mark style="color:purple;">untitle</mark>

Converts the first letter of each word in a string to lowercase.

<table data-header-hidden><thead><tr><th width="164">Name</th><th>Value</th></tr></thead><tbody><tr><td>Signature</td><td><pre class="language-go"><code class="lang-go">Untitle(str string) string
</code></pre></td></tr></tbody></table>

{% tabs %}
{% tab title="Template Example" %}
```go
{{ "Hello World" | untitle }} // Output: "hello world"
```
{% endtab %}
{% endtabs %}

### <mark style="color:purple;">swapCase</mark>

Switches the case of each letter in a string, converting lowercase to uppercase and vice versa.

<table data-header-hidden><thead><tr><th width="164">Name</th><th>Value</th></tr></thead><tbody><tr><td>Signature</td><td><pre class="language-go"><code class="lang-go">SwapCase(str string) string
</code></pre></td></tr></tbody></table>

{% tabs %}
{% tab title="Template Example" %}
```go
{{ "Hello World" | swapCase }} // Output: "hELLO wORLD"
```
{% endtab %}
{% endtabs %}

### <mark style="color:purple;">capitalize</mark>

Uppercases the first letter of a string while leaving the rest of the string unchanged.

<table data-header-hidden><thead><tr><th width="164">Name</th><th>Value</th></tr></thead><tbody><tr><td>Signature</td><td><pre class="language-go"><code class="lang-go">capitalize(str string) string
</code></pre></td></tr></tbody></table>

{% tabs %}
{% tab title="Template Example" %}
```go
{{ "hello world" | capitalize }} // Output: "Hello world"
{{ "123boo_bar" | capitalize }} // Output: 123Boo_bar
{{ " Fe bar" | capitalize }} // Output: " Fe bar"
```
{% endtab %}
{% endtabs %}

### <mark style="color:purple;">uncapitalize</mark>

Lowercases the first letter of a string while leaving the rest of the string unchanged.

<table data-header-hidden><thead><tr><th width="164">Name</th><th>Value</th></tr></thead><tbody><tr><td>Signature</td><td><pre class="language-go"><code class="lang-go">uncapitalize(str string) string
</code></pre></td></tr></tbody></table>

{% tabs %}
{% tab title="Template Example" %}
```go
{{ "Hello World" | uncapitalize }} // Output: "hello World"
{{ "123Boo_bar" | uncapitalize }} // Output: 123boo_bar
{{ " Fe bar" | uncapitalize }} // Output: " fe bar"
```
{% endtab %}
{% endtabs %}

### <mark style="color:purple;">split</mark>

Divides a string into a map of parts based on a specified separator, returning a collection of the split components.

<table data-header-hidden><thead><tr><th width="164">Name</th><th>Value</th></tr></thead><tbody><tr><td>Signature</td><td><pre class="language-go"><code class="lang-go">Split(sep, orig string) map[string]string
</code></pre></td></tr></tbody></table>

{% tabs %}
{% tab title="Template Example" %}
```go
{{ "apple,banana,cherry" | split "," }}
// Output: { "_0":"apple", "_1":"banana", "_2":"cherry" }
```
{% endtab %}
{% endtabs %}

### <mark style="color:purple;">splitn</mark>

Splits a string into a specified number of parts using a separator, returning a map with up to `n` elements.

<table data-header-hidden><thead><tr><th width="164">Name</th><th>Value</th></tr></thead><tbody><tr><td>Signature</td><td><pre class="language-go"><code class="lang-go">Splitn(sep string, n int, orig string) map[string]string
</code></pre></td></tr></tbody></table>

{% tabs %}
{% tab title="Template Example" %}
```go
{{ "apple,banana,cherry" | split "," 2 }}
// Output: { "_0":"apple", "_1":"banana,cherry" }
```
{% endtab %}
{% endtabs %}

### <mark style="color:purple;">substr</mark>

Extracts a portion of a string based on given start and end positions, with support for negative indices to count from the end.

<table data-header-hidden><thead><tr><th width="164">Name</th><th>Value</th></tr></thead><tbody><tr><td>Signature</td><td><pre class="language-go"><code class="lang-go">Substring(start int, end int, str string) string
</code></pre></td></tr></tbody></table>

{% tabs %}
{% tab title="Template Example" %}
```go
{{ "Hello World" | substr 0 5 }} // Output: "Hello"
```
{% endtab %}
{% endtabs %}

### <mark style="color:purple;">indent</mark>

Adds spaces to the beginning of each line in a string, effectively indenting the text.

<table data-header-hidden><thead><tr><th width="164">Name</th><th>Value</th></tr></thead><tbody><tr><td>Signature</td><td><pre class="language-go"><code class="lang-go">Indent(spaces int, str string) string
</code></pre></td></tr></tbody></table>

{% tabs %}
{% tab title="Template Example" %}
```go
{{ "Hello\nWorld" | indent 4 }} // Output: "    Hello\n    World"
```
{% endtab %}
{% endtabs %}

### <mark style="color:purple;">nindent</mark>

Similar to `Indent`, but also adds a newline before the indented lines.

<table data-header-hidden><thead><tr><th width="164">Name</th><th>Value</th></tr></thead><tbody><tr><td>Signature</td><td><pre class="language-go"><code class="lang-go">Nindent(spaces int, str string) string
</code></pre></td></tr></tbody></table>

{% tabs %}
{% tab title="Template Example" %}
```go
{{ "Hello\nWorld" | nindent 4 }} // Output: "\n    Hello\n    World"
```
{% endtab %}
{% endtabs %}

### <mark style="color:purple;">seq</mark>

Generates a sequence of numbers as a string, allowing for customizable start, end, and step values, similar to the Unix `seq` command.

<table data-header-hidden><thead><tr><th width="164">Name</th><th>Value</th></tr></thead><tbody><tr><td>Signature</td><td><pre class="language-go"><code class="lang-go">Seq(params ...int) string
</code></pre></td></tr></tbody></table>

{% tabs %}
{% tab title="Template Example" %}
```go
{{ seq 1 2 10 }} // Output: "1 3 5 7 9"
{{ seq 3 -3 2 }} // Output: "3"
{{ seq }} // Output: ""
{{ seq 0 4 }} // Output: "0 1 2 3 4"
{{ seq -5 }} // Output: "1 0 -1 -2 -3 -4 -5"
{{ seq 0 -4 }} // Output: "0 -1 -2 -3 -4"
```
{% endtab %}
{% endtabs %}
