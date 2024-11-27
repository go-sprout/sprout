---
description: >-
  The Regexp registry includes functions for pattern matching and string
  manipulation using regular expressions, providing powerful text processing
  capabilities.
---

# Regexp

{% hint style="info" %}
You can easily import all the functions from the <mark style="color:yellow;">`regexp`</mark> registry by including the following import statement in your code

```go
import "github.com/go-sprout/sprout/registry/regexp"
```
{% endhint %}

### <mark style="color:purple;">regexFind</mark>

The function returns the first match found in the string that corresponds to the specified regular expression pattern.

<table data-header-hidden><thead><tr><th width="174">Name</th><th>Value</th></tr></thead><tbody><tr><td>Signature</td><td><pre class="language-go"><code class="lang-go">RegexFind(regex string, s string) (string, error)
</code></pre></td></tr></tbody></table>

{% tabs %}
{% tab title="Template Example" %}
<pre class="language-go"><code class="lang-go">{{ "hello world" | regexFind "hello" }} // Output: "hello", nil
<strong>{{ "hello world" | regexFind "\invalid$^///" }} // Error
</strong></code></pre>
{% endtab %}
{% endtabs %}

### <mark style="color:purple;">regexFindAll</mark>

The function returns all matches of the regex pattern in the string, up to a specified maximum number of matches (`n`).

<table data-header-hidden><thead><tr><th width="164">Name</th><th>Value</th></tr></thead><tbody><tr><td>Signature</td><td><pre class="language-go"><code class="lang-go">RegexFindAll(regex string, s string, n int) ([]string, error)
</code></pre></td></tr></tbody></table>

{% tabs %}
{% tab title="Template Example" %}
```go
{{ regexFindAll "a." "aba acada afa" 3 }} // Output: ["ab", "ac", "af"], nil
{{ regexFindAll "\invalid$^///" "aba acada afa" 3 }} // Error
```
{% endtab %}
{% endtabs %}

### <mark style="color:purple;">regexMatch</mark>

The function checks if the entire string matches the given regular expression pattern.

<table data-header-hidden><thead><tr><th width="164">Name</th><th>Value</th></tr></thead><tbody><tr><td>Signature</td><td><pre class="language-go"><code class="lang-go">RegexMatch(regex string, s string) (bool, error)
</code></pre></td></tr></tbody></table>

{% tabs %}
{% tab title="Template Example" %}
```go
{{ regexMatch "^[a-zA-Z]+$" "Hello" }} // Output: true, nil
{{ regexMatch "\invalid$^///" "Hello" }} // Output: false, error
```
{% endtab %}
{% endtabs %}

### <mark style="color:purple;">regexSplit</mark>

The function splits the string into substrings based on matches of the regex pattern, performing the split up to `n` times.

<table data-header-hidden><thead><tr><th width="164">Name</th><th>Value</th></tr></thead><tbody><tr><td>Signature</td><td><pre class="language-go"><code class="lang-go">RegexSplit(regex string, s string, n int) ([]string, error)
</code></pre></td></tr></tbody></table>

{% tabs %}
{% tab title="Template Example" %}
```go
{{ mustRegexSplit "\\s+" "hello world from Go" 2 }} // Output: ["hello", "world from Go"], nil
{{ mustRegexSplit "\invalid$^///" "hello world from Go" 2 }} // Error
```
{% endtab %}
{% endtabs %}

### <mark style="color:purple;">regexReplaceAll</mark>

The function replaces all occurrences of the regex pattern in the string with the specified replacement string.

<table data-header-hidden><thead><tr><th width="164">Name</th><th>Value</th></tr></thead><tbody><tr><td>Signature</td><td><pre class="language-go"><code class="lang-go">RegexReplaceAll(regex string, s string, repl string) (string, error)
</code></pre></td></tr></tbody></table>

{% tabs %}
{% tab title="Template Example" %}
```go
{{ regexReplaceAll "\\d" "R2D2 C3PO" "X" }} // Output: "RXDX CXPO", nil
{{ regexReplaceAll "\invalid$^///" "R2D2 C3PO" "X" }} // Output: "", error
```
{% endtab %}
{% endtabs %}

### <mark style="color:purple;">regexReplaceAllLiteral</mark>

The function replaces all occurrences of the regex pattern in the string with the specified literal replacement string, without interpreting any special characters in the replacement.

<table data-header-hidden><thead><tr><th width="164">Name</th><th>Value</th></tr></thead><tbody><tr><td>Signature</td><td><pre class="language-go"><code class="lang-go">RegexReplaceAllLiteral(regex string, s string, repl string) (string, error)
</code></pre></td></tr></tbody></table>

{% tabs %}
{% tab title="Template Example" %}
```go
{{ regexReplaceAllLiteral "world" "hello world" "$1" }} // Output: "hello $1", nil
{{ regexReplaceAllLiteral "world" "hello world" "\invalid$^///" }} // Output: "", error
```
{% endtab %}
{% endtabs %}

### <mark style="color:purple;">regexQuoteMeta</mark>

The function returns a version of the provided string that can be used as a literal pattern in a regular expression, escaping any special characters.

<table data-header-hidden><thead><tr><th width="164">Name</th><th>Value</th></tr></thead><tbody><tr><td>Signature</td><td><pre class="language-go"><code class="lang-go">RegexQuoteMeta(str string) string
</code></pre></td></tr></tbody></table>

{% tabs %}
{% tab title="Template Example" %}
```go
{{ regexQuoteMeta ".+*?^$()[]{}|" }} // Output: "\.\+\*\?\^\$\(\)\[\]\{\}\|"
```
{% endtab %}
{% endtabs %}

### <mark style="color:purple;">regexFindGroups</mark>

The function finds the first match of a regex pattern in a string and returns the matched groups, with error handling.

<table data-header-hidden><thead><tr><th width="164">Name</th><th>Value</th></tr></thead><tbody><tr><td>Signature</td><td><pre class="language-go"><code class="lang-go">RegexFindGroups(regex string, str string) ([]string, error)
</code></pre></td></tr></tbody></table>

{% tabs %}
{% tab title="Template Example" %}
```go
{{ "aaabbb" | regexFindGroups "(a+)(b+)" }} // Output: ["aaabbb", "aaa", "bbb"], nil
{{ "aaabbb" | regexFindGroups "\invalid$^///" }} // Error
```
{% endtab %}
{% endtabs %}

### <mark style="color:purple;">regexFindAllGroups</mark>

The function finds all matches of a regex pattern in a string up to a specified limit and returns the matched groups, with error handling.

<table data-header-hidden><thead><tr><th width="164">Name</th><th>Value</th></tr></thead><tbody><tr><td>Signature</td><td><pre class="language-go"><code class="lang-go">RegexFindAllGroups(regex string, n int, str string) ([]string, error)
</code></pre></td></tr></tbody></table>

{% tabs %}
{% tab title="Template Example" %}
```go
{{ "aaabbb aab aaabbb" | regexFindAllGroups "(a+)(b+)" -1 }} // Output: [["aaabbb", "aaa", "bbb"], ["aab", "aa", "b"], ["aaabbb", "aaa", "bbb"]], nil
{{ "aaabbb aab aaabbb" | regexFindAllGroups "(a+)(b+)" 1 }} // Output: [["aaabbb", "aaa", "bbb"]], nil
{{ "aaabbb" | regexFindAllGroups "\invalid$^///" }} // Error
```
{% endtab %}
{% endtabs %}

### <mark style="color:purple;">regexFindNamed</mark>

The function finds the first match of a regex pattern with named capturing groups in a string and returns a map of group names to matched strings, with error handling.

<table data-header-hidden><thead><tr><th width="164">Name</th><th>Value</th></tr></thead><tbody><tr><td>Signature</td><td><pre class="language-go"><code class="lang-go">RegexFindNamed(regex string, str string) (map[string]string, error)
</code></pre></td></tr></tbody></table>

{% tabs %}
{% tab title="Template Example" %}
```go
{{ "aaabbb" | regexFindNamed "(?P<first>a+)(?P<second>b+)" }} // Output: map["first":"aaa", "second":"bbb"], nil
{{ "aaabbb" | regexFindNamed "(?P<first>a+)(b+)" }} // Output: map["first":"aaa"], nil
{{ "bbb" | regexFindNamed "(?P<first>a+)" }} // Output: map[], nil
{{ "aaabbb" | regexFindNamed "\invalid$^///" }} // Error
```
{% endtab %}
{% endtabs %}

### <mark style="color:purple;">regexFindAllNamed</mark>

The function finds all matches of a regex pattern with named capturing groups in a string up to a specified limit and returns a slice of maps of group names to matched strings, with error handling.

<table data-header-hidden><thead><tr><th width="164">Name</th><th>Value</th></tr></thead><tbody><tr><td>Signature</td><td><pre class="language-go"><code class="lang-go">RegexFindAllNamed(regex string, n int, str string) ([]map[string]string, error)
</code></pre></td></tr></tbody></table>

{% tabs %}
{% tab title="Template Example" %}
```go
{{ "var1=value1&var2=value2" | regexFindAllNamed "(?P<param>\\w+)=(?P<value>\\w+)" -1 }} // Output: [map[param:var1 value:value1] map[param:var2 value:value2]], nil
{{ "var1=value1&var2=value2" | regexFindAllNamed "(?P<param>\\w+)=(?P<value>\\w+)" 1 }} // Output: [map[param:var1 value:value1]], nil
{{ "var1+value1" | regexFindAllNamed "(?P<param>\\w+)=(?P<value>\\w+)" -1 }} // Output: map[], nil
{{ "var1=value1" | regexFindAllNamed "\invalid$^///" -1 }} // Error
```
{% endtab %}
{% endtabs %}
