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

### <mark style="color:purple;">regexFind / mustRegexFind</mark>

The function returns the first match found in the string that corresponds to the specified regular expression pattern.

<table data-header-hidden><thead><tr><th width="174">Name</th><th>Value</th></tr></thead><tbody><tr><td>Signature</td><td><pre class="language-go"><code class="lang-go">RegexFind(regex string, s string) string
MustRegexFind(regex string, s string) (string, error)
</code></pre></td></tr><tr><td>Must version</td><td><span data-gb-custom-inline data-tag="emoji" data-code="2705">✅</span></td></tr></tbody></table>

{% tabs %}
{% tab title="Template Example" %}
```go
{{ regexFind "a(b+)" "aaabbb" }} // Output: "abbb"
```
{% endtab %}

{% tab title="Must version" %}
<pre class="language-go"><code class="lang-go">{{ "hello world" | mustRegexFind "hello" }} // Output: "hello", nil
<strong>{{ "hello world" | mustRegexFind "\invalid$^///" }} // Output: "", error
</strong></code></pre>
{% endtab %}
{% endtabs %}

### <mark style="color:purple;">regexFindAll / mustRegexFindAll</mark>

The function returns all matches of the regex pattern in the string, up to a specified maximum number of matches (`n`).

<table data-header-hidden><thead><tr><th width="164">Name</th><th>Value</th></tr></thead><tbody><tr><td>Signature</td><td><pre class="language-go"><code class="lang-go">RegexFindAll(regex string, s string, n int) []string
MustRegexFindAll(regex string, s string, n int) ([]string, error)
</code></pre></td></tr><tr><td>Must version</td><td><span data-gb-custom-inline data-tag="emoji" data-code="2705">✅</span></td></tr></tbody></table>

{% tabs %}
{% tab title="Template Example" %}
```go
{{ regexFindAll "a(b+)" "ababab" 2 }} // Output: ["ab", "ab"]
```
{% endtab %}

{% tab title="Must version" %}
```go
{{ mustRegexFindAll "a.", "aba acada afa", 3 }} // Output: ["ab", "ac", "af"], nil
{{ mustRegexFindAll "\invalid$^///", "aba acada afa", 3 }} // Output: "", error
```
{% endtab %}
{% endtabs %}

### <mark style="color:purple;">regexMatch / mustRegexMatch</mark>

The function checks if the entire string matches the given regular expression pattern.

<table data-header-hidden><thead><tr><th width="164">Name</th><th>Value</th></tr></thead><tbody><tr><td>Signature</td><td><pre class="language-go"><code class="lang-go">RegexMatch(regex string, s string) bool
MustRegexMatch(regex string, s string) (bool, error)
</code></pre></td></tr><tr><td>Must version</td><td><span data-gb-custom-inline data-tag="emoji" data-code="2705">✅</span></td></tr></tbody></table>

{% tabs %}
{% tab title="Template Example" %}
```go
{{ regexMatch "^[a-zA-Z]+$" "Hello" }} // Output: true
```
{% endtab %}

{% tab title="Must version" %}
```go
{{ mustRegexMatch "^[a-zA-Z]+$", "Hello" }} // Output: true, nil
{{ mustRegexMatch "\invalid$^///", "Hello" }} // Output: false, error
```
{% endtab %}
{% endtabs %}

### <mark style="color:purple;">regexSplit / mustRegexSplit</mark>

The function splits the string into substrings based on matches of the regex pattern, performing the split up to `n` times.

<table data-header-hidden><thead><tr><th width="164">Name</th><th>Value</th></tr></thead><tbody><tr><td>Signature</td><td><pre class="language-go"><code class="lang-go">RegexSplit(regex string, s string, n int) []string
MustRegexSplit(regex string, s string, n int) ([]string, error)
</code></pre></td></tr><tr><td>Must version</td><td><span data-gb-custom-inline data-tag="emoji" data-code="2705">✅</span></td></tr></tbody></table>

{% tabs %}
{% tab title="Template Example" %}
```go
{{regexSplit "\\s+" "hello world" -1 }} // Output: ["hello", "world"]
```
{% endtab %}

{% tab title="Must version" %}
```go
{{ mustRegexSplit "\\s+", "hello world from Go", 2 }} // Output: ["hello", "world from Go"], nil
{{ mustRegexSplit "\invalid$^///", "hello world from Go", 2 }} // Output: [], error
```
{% endtab %}
{% endtabs %}

### <mark style="color:purple;">regexReplaceAll / mustRegexReplaceAll</mark>

The function replaces all occurrences of the regex pattern in the string with the specified replacement string.

<table data-header-hidden><thead><tr><th width="164">Name</th><th>Value</th></tr></thead><tbody><tr><td>Signature</td><td><pre class="language-go"><code class="lang-go">RegexReplaceAll(regex string, s string, repl string) string
MustRegexReplaceAll(regex string, s string, repl string) (string, error)
</code></pre></td></tr><tr><td>Must version</td><td><span data-gb-custom-inline data-tag="emoji" data-code="2705">✅</span></td></tr></tbody></table>

{% tabs %}
{% tab title="Template Example" %}
```go
{{ regexReplaceAll "[aeiou]" "hello" "i" }} // Output: "hillo"
```
{% endtab %}

{% tab title="Must version" %}
```go
{{ mustRegexReplaceAll "\\d", "R2D2 C3PO", "X" }} // Output: "RXDX CXPO", nil
{{ mustRegexReplaceAll "\invalid$^///", "R2D2 C3PO", "X" }} // Output: "", error
```
{% endtab %}
{% endtabs %}

### <mark style="color:purple;">regexReplaceAllLiteral / mustRegexReplaceAllLiteral</mark>

The function replaces all occurrences of the regex pattern in the string with the specified literal replacement string, without interpreting any special characters in the replacement.

<table data-header-hidden><thead><tr><th width="164">Name</th><th>Value</th></tr></thead><tbody><tr><td>Signature</td><td><pre class="language-go"><code class="lang-go">RegexReplaceAllLiteral(regex string, s string, repl string) string
MustRegexReplaceAllLiteral(regex string, s string, repl string) (string, error)
</code></pre></td></tr><tr><td>Must version</td><td><span data-gb-custom-inline data-tag="emoji" data-code="2705">✅</span></td></tr></tbody></table>

{% tabs %}
{% tab title="Template Example" %}
```go
{{ regexReplaceAllLiteral "[aeiou]" "hello" "$&" }} // Output: "h$&ll$&"
```
{% endtab %}

{% tab title="Must version" %}
```go
{{ mustRegexReplaceAllLiteral "world", "hello world", "$1" }} // Output: "hello $1", nil
{{ mustRegexReplaceAllLiteral "world", "hello world", "\invalid$^///" }} // Output: "", error
```
{% endtab %}
{% endtabs %}

### <mark style="color:purple;">regexQuoteMeta</mark>

The function returns a version of the provided string that can be used as a literal pattern in a regular expression, escaping any special characters.

<table data-header-hidden><thead><tr><th width="164">Name</th><th>Value</th></tr></thead><tbody><tr><td>Signature</td><td><pre class="language-go"><code class="lang-go">RegexQuoteMeta(str string) string
</code></pre></td></tr><tr><td>Must version</td><td><span data-gb-custom-inline data-tag="emoji" data-code="274c">❌</span></td></tr></tbody></table>

{% tabs %}
{% tab title="Template Example" %}
```go
{{ regexQuoteMeta ".+*?^$()[]{}|" }} // Output: "\.\+\*\?\^\$\(\)\[\]\{\}\|"
```
{% endtab %}
{% endtabs %}

