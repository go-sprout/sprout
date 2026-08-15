---
description: >-
  The Regex registry includes functions for pattern matching and string
  manipulation using regular expressions, where the main parameter is always the
  last one to be usable in a template pipeline.
---

# Regex

{% hint style="info" %}
You can easily import all the functions from the <mark style="color:yellow;">`regex`</mark> registry by including the following import statement in your code

```go
import "github.com/go-sprout/sprout/registry/regex"
```
{% endhint %}

{% hint style="warning" %}
The `regex` registry replaces the deprecated [`regexp`](regexp.md) registry, which keeps the historical sprig signatures. Both registries expose the same function names, so they are **mutually exclusive**: register one or the other, never both.

Only four functions change their signature, the string to work on moves to the last position:

| `regexp` (deprecated)                                   | `regex`                                                 |
| ------------------------------------------------------- | ------------------------------------------------------- |
| `regexFindAll <regex> <value> <n>`                      | `regexFindAll <regex> <n> <value>`                      |
| `regexSplit <regex> <value> <n>`                        | `regexSplit <regex> <n> <value>`                        |
| `regexReplaceAll <regex> <value> <replacedBy>`          | `regexReplaceAll <regex> <replacedBy> <value>`          |
| `regexReplaceAllLiteral <regex> <value> <replacedBy>`   | `regexReplaceAllLiteral <regex> <replacedBy> <value>`   |

All the other functions keep the exact same signature, only the import changes. The `must` prefixed versions are not carried over, use the standard functions instead.
{% endhint %}

### <mark style="color:purple;">regexFind</mark>

The function returns the first match found in the string that corresponds to the specified regular expression pattern.

<table data-header-hidden><thead><tr><th width="174">Name</th><th>Value</th></tr></thead><tbody><tr><td>Signature</td><td><pre class="language-go"><code class="lang-go">RegexFind(regex string, value string) (string, error)
</code></pre></td></tr></tbody></table>

{% tabs %}
{% tab title="Template Example" %}
```go
{{ "hello world" | regexFind "hello" }} // Output: "hello"
{{ "hello world" | regexFind "\\invalid$^///" }} // Error
```
{% endtab %}
{% endtabs %}

### <mark style="color:purple;">regexFindAll</mark>

The function returns all matches of the regex pattern in the string, up to a specified maximum number of matches (`n`).

<table data-header-hidden><thead><tr><th width="164">Name</th><th>Value</th></tr></thead><tbody><tr><td>Signature</td><td><pre class="language-go"><code class="lang-go">RegexFindAll(regex string, n int, value string) ([]string, error)
</code></pre></td></tr></tbody></table>

{% tabs %}
{% tab title="Template Example" %}
```go
{{ "aba acada afa" | regexFindAll "a." 3 }} // Output: [ab a  ac]
{{ "aba acada afa" | regexFindAll "\\invalid$^///" 3 }} // Error
```
{% endtab %}
{% endtabs %}

### <mark style="color:purple;">regexMatch</mark>

The function checks if the entire string matches the given regular expression pattern.

<table data-header-hidden><thead><tr><th width="164">Name</th><th>Value</th></tr></thead><tbody><tr><td>Signature</td><td><pre class="language-go"><code class="lang-go">RegexMatch(regex string, value string) (bool, error)
</code></pre></td></tr></tbody></table>

{% tabs %}
{% tab title="Template Example" %}
```go
{{ "Hello" | regexMatch "^[a-zA-Z]+$" }} // Output: true
{{ "Hello" | regexMatch "\\invalid$^///" }} // Error
```
{% endtab %}
{% endtabs %}

### <mark style="color:purple;">regexSplit</mark>

The function splits the string into substrings based on matches of the regex pattern, performing the split up to `n` times.

<table data-header-hidden><thead><tr><th width="164">Name</th><th>Value</th></tr></thead><tbody><tr><td>Signature</td><td><pre class="language-go"><code class="lang-go">RegexSplit(regex string, n int, value string) ([]string, error)
</code></pre></td></tr></tbody></table>

{% tabs %}
{% tab title="Template Example" %}
```go
{{ "hello world from Go" | regexSplit "\\s+" 2 }} // Output: [hello world from Go]
{{ "hello world from Go" | regexSplit "\\invalid$^///" 2 }} // Error
```
{% endtab %}
{% endtabs %}

{% hint style="info" %}
The output above is a slice of two elements, `hello` and `world from Go`, rendered by the template engine.
{% endhint %}

### <mark style="color:purple;">regexReplaceAll</mark>

The function replaces all occurrences of the regex pattern in the string with the specified replacement string.

<table data-header-hidden><thead><tr><th width="164">Name</th><th>Value</th></tr></thead><tbody><tr><td>Signature</td><td><pre class="language-go"><code class="lang-go">RegexReplaceAll(regex string, replacedBy string, value string) (string, error)
</code></pre></td></tr></tbody></table>

{% tabs %}
{% tab title="Template Example" %}
```go
{{ "R2D2 C3PO" | regexReplaceAll "\\d" "X" }} // Output: "RXDX CXPO"
{{ "R2D2 C3PO" | regexReplaceAll "\\invalid$^///" "X" }} // Error
```
{% endtab %}
{% endtabs %}

### <mark style="color:purple;">regexReplaceAllLiteral</mark>

The function replaces all occurrences of the regex pattern in the string with the specified literal replacement string, without interpreting any special characters in the replacement.

<table data-header-hidden><thead><tr><th width="164">Name</th><th>Value</th></tr></thead><tbody><tr><td>Signature</td><td><pre class="language-go"><code class="lang-go">RegexReplaceAllLiteral(regex string, replacedBy string, value string) (string, error)
</code></pre></td></tr></tbody></table>

{% tabs %}
{% tab title="Template Example" %}
```go
{{ "hello world" | regexReplaceAllLiteral "world" "$1" }} // Output: "hello $1"
{{ "hello world" | regexReplaceAllLiteral "none" "all" }} // Output: "hello world"
{{ "hello world" | regexReplaceAllLiteral "\\invalid$^///" "all" }} // Error
```
{% endtab %}
{% endtabs %}

### <mark style="color:purple;">regexQuoteMeta</mark>

The function returns a version of the provided string that can be used as a literal pattern in a regular expression, escaping any special characters.

<table data-header-hidden><thead><tr><th width="164">Name</th><th>Value</th></tr></thead><tbody><tr><td>Signature</td><td><pre class="language-go"><code class="lang-go">RegexQuoteMeta(value string) string
</code></pre></td></tr></tbody></table>

{% tabs %}
{% tab title="Template Example" %}
```go
{{ regexQuoteMeta ".+*?^$()[]{}|" }}
// Output: \\.\\+\\*\\?\\^\\$\\(\\)\\[\\]\\{\\}\\|
```
{% endtab %}
{% endtabs %}

### <mark style="color:purple;">regexFindGroups</mark>

The function finds the first match of a regex pattern in a string and returns the matched groups, with error handling.

<table data-header-hidden><thead><tr><th width="164">Name</th><th>Value</th></tr></thead><tbody><tr><td>Signature</td><td><pre class="language-go"><code class="lang-go">RegexFindGroups(regex string, value string) ([]string, error)
</code></pre></td></tr></tbody></table>

{% tabs %}
{% tab title="Template Example" %}
```go
{{ "aaabbb" | regexFindGroups "(a+)(b+)" }} // Output: [aaabbb aaa bbb]
{{ "aaabbb" | regexFindGroups "\\invalid$^///" }} // Error
```
{% endtab %}
{% endtabs %}

### <mark style="color:purple;">regexFindAllGroups</mark>

The function finds all matches of a regex pattern in a string up to a specified limit and returns the matched groups, with error handling.

<table data-header-hidden><thead><tr><th width="164">Name</th><th>Value</th></tr></thead><tbody><tr><td>Signature</td><td><pre class="language-go"><code class="lang-go">RegexFindAllGroups(regex string, n int, value string) ([][]string, error)
</code></pre></td></tr></tbody></table>

{% tabs %}
{% tab title="Template Example" %}
```go
{{ "aaabbb aab aaabbb" | regexFindAllGroups "(a+)(b+)" -1 }} // Output: [[aaabbb aaa bbb] [aab aa b] [aaabbb aaa bbb]]
{{ "aaabbb aab aaabbb" | regexFindAllGroups "(a+)(b+)" 1 }} // Output: [[aaabbb aaa bbb]]
{{ "aaabbb" | regexFindAllGroups "\\invalid$^///" -1 }} // Error
```
{% endtab %}
{% endtabs %}

### <mark style="color:purple;">regexFindNamed</mark>

The function finds the first match of a regex pattern with named capturing groups in a string and returns a map of group names to matched strings, with error handling.

<table data-header-hidden><thead><tr><th width="164">Name</th><th>Value</th></tr></thead><tbody><tr><td>Signature</td><td><pre class="language-go"><code class="lang-go">RegexFindNamed(regex string, value string) (map[string]string, error)
</code></pre></td></tr></tbody></table>

{% tabs %}
{% tab title="Template Example" %}
```go
{{ "aaabbb" | regexFindNamed "(?P<first>a+)(?P<second>b+)" }} // Output: map[first:aaa second:bbb]
{{ "aaabbb" | regexFindNamed "(?P<first>a+)(b+)" }} // Output: map[first:aaa]
{{ "bbb" | regexFindNamed "(?P<first>a+)" }} // Output: map[]
{{ "aaabbb" | regexFindNamed "\\invalid$^///" }} // Error
```
{% endtab %}
{% endtabs %}

### <mark style="color:purple;">regexFindAllNamed</mark>

The function finds all matches of a regex pattern with named capturing groups in a string up to a specified limit and returns a slice of maps of group names to matched strings, with error handling.

<table data-header-hidden><thead><tr><th width="164">Name</th><th>Value</th></tr></thead><tbody><tr><td>Signature</td><td><pre class="language-go"><code class="lang-go">RegexFindAllNamed(regex string, n int, value string) ([]map[string]string, error)
</code></pre></td></tr></tbody></table>

{% tabs %}
{% tab title="Template Example" %}
```go
{{ "var1=value1&var2=value2" | regexFindAllNamed "(?P<param>\\w+)=(?P<value>\\w+)" -1 }} // Output: [map[param:var1 value:value1] map[param:var2 value:value2]]
{{ "var1=value1&var2=value2" | regexFindAllNamed "(?P<param>\\w+)=(?P<value>\\w+)" 1 }} // Output: [map[param:var1 value:value1]]
{{ "var1+value1" | regexFindAllNamed "(?P<param>\\w+)=(?P<value>\\w+)" -1 }} // Output: []
{{ "var1=value1" | regexFindAllNamed "\\invalid$^///" -1 }} // Error
```
{% endtab %}
{% endtabs %}
