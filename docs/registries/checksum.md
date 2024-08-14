---
description: >-
  The Checksum registry offers functions to generate and verify checksums,
  ensuring data integrity. It supports various algorithms for reliable error
  detection and data validation.
---

# Checksum

{% hint style="info" %}
You can easily import all the functions from the <mark style="color:yellow;">`checksum`</mark> registry by including the following import statement in your code

```go
import "github.com/go-sprout/sprout/registry/checksum"
```
{% endhint %}

### <mark style="color:purple;">sha1sum</mark>

Sha1sum calculates the SHA-1 hash of the input string and returns it as a hexadecimal encoded string.

<table data-header-hidden><thead><tr><th width="174">Name</th><th>Value</th></tr></thead><tbody><tr><td>Signature</td><td><pre class="language-go"><code class="lang-go">SHA1Sum(input string) string
</code></pre></td></tr><tr><td>Must version</td><td><span data-gb-custom-inline data-tag="emoji" data-code="274c">❌</span></td></tr></tbody></table>

{% tabs %}
{% tab title="Template Example" %}
<pre class="language-go"><code class="lang-go"><strong>{{ sha1sum "" }} // Output: da39a3ee5e6b4b0d3255bfef95601890afd80709
</strong>{{ sha1sum "Hello, World!" }} // Output: 0a0a9f2a6772942557ab5355d76af442f8f65e01
</code></pre>
{% endtab %}
{% endtabs %}

### <mark style="color:purple;">sha256sum</mark>

Sha256sum calculates the SHA-256 hash of the input string and returns it as a hexadecimal encoded string.

<table data-header-hidden><thead><tr><th width="174">Name</th><th>Value</th></tr></thead><tbody><tr><td>Signature</td><td><pre class="language-go"><code class="lang-go">SHA256Sum(input string) string
</code></pre></td></tr><tr><td>Must version</td><td><span data-gb-custom-inline data-tag="emoji" data-code="274c">❌</span></td></tr></tbody></table>

{% tabs %}
{% tab title="Template Example" %}
```go
{{ sha256sum "" }} // Output: e3b0c44298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b7852b855
{{ sha256sum "Hello, World!" }} // Output: dffd6021bb2bd5b0af676290809ec3a53191dd81c7f70a4b28688a362182986f
```
{% endtab %}
{% endtabs %}

### <mark style="color:purple;">adler32sum</mark>

Adler32sum calculates the Adler-32 checksum of the input string and returns it as a hexadecimal encoded string.

<table data-header-hidden><thead><tr><th width="174">Name</th><th>Value</th></tr></thead><tbody><tr><td>Signature</td><td><pre class="language-go"><code class="lang-go">Adler32Sum(input string) string
</code></pre></td></tr><tr><td>Must version</td><td><span data-gb-custom-inline data-tag="emoji" data-code="274c">❌</span></td></tr></tbody></table>

{% tabs %}
{% tab title="Tempalte Example" %}
```go
{{ adler32sum "" }} // Output: 00000001
{{ adler32sum "Hello, World!" }} // Output: 1f9e046a
```
{% endtab %}
{% endtabs %}

### <mark style="color:purple;">md5sum</mark>

Md5sum calculates the MD5 hash of the input string and returns it as a hexadecimal encoded string.

<table data-header-hidden><thead><tr><th width="174">Name</th><th>Value</th></tr></thead><tbody><tr><td>Signature</td><td><pre class="language-go"><code class="lang-go">MD5Sum(input string) string
</code></pre></td></tr><tr><td>Must version</td><td><span data-gb-custom-inline data-tag="emoji" data-code="274c">❌</span></td></tr></tbody></table>

{% tabs %}
{% tab title="Template Example" %}
```go
{{ md5sum "" }} // Output: d41d8cd98f00b204e9800998ecf8427e
{{ md5sum "Hello, World!" }} // Output: 65a8e27d8879283831b664bd8b7f0ad4
```
{% endtab %}
{% endtabs %}
