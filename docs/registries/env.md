---
description: >-
  The Env registry provides functions for accessing and managing environment
  variables, enabling dynamic configuration of your applications based on the
  runtime environment.
---

# Env

{% hint style="info" %}
You can easily import all the functions from the <mark style="color:yellow;">`env`</mark> registry by including the following import statement in your code

```go
import "github.com/go-sprout/sprout/registry/env"
```
{% endhint %}

### <mark style="color:purple;">env</mark>

The function retrieves the value of a specified environment variable from the system.

<table data-header-hidden><thead><tr><th width="174">Name</th><th>Value</th></tr></thead><tbody><tr><td>Signature</td><td><pre class="language-go"><code class="lang-go">Env(key string) string
</code></pre></td></tr></tbody></table>

{% tabs %}
{% tab title="Template Example" %}
```go
{{ env "INVALID" }} // Output: ""
{{ "PATH" | env }} // Output(will be different): "/usr/bin:/bin:/usr/sbin:/sbin"
```
{% endtab %}
{% endtabs %}

### <mark style="color:purple;">expandEnv</mark>

The function replaces occurrences of `${var}` or `$var` in a string with the corresponding values from the current environment variables.

<table data-header-hidden><thead><tr><th width="164">Name</th><th>Value</th></tr></thead><tbody><tr><td>Signature</td><td><pre class="language-go"><code class="lang-go">ExpandEnv(str string) string
</code></pre></td></tr></tbody></table>

{% tabs %}
{% tab title="Template Example" %}
```go
{{ "Path is $PATH" | expandEnv }} // Output(will be different): "Path is /usr/bin:/bin:/usr/sbin:/sbin"
```
{% endtab %}
{% endtabs %}
