---
description: >-
  The Semver registry is designed to handle semantic versioning, offering
  functions to compare and manage version numbers consistently across your
  projects.
---

# Semver

{% hint style="info" %}
You can easily import all the functions from the <mark style="color:yellow;">`semver`</mark> registry by including the following import statement in your code

```go
import "github.com/go-sprout/sprout/registry/semver"
```
{% endhint %}

{% hint style="info" %}
This registry utilizing the original [Semver package](https://github.com/Masterminds/semver) created by Masterminds, which adheres to the Semantic Versioning specification.
{% endhint %}

### <mark style="color:purple;">semver</mark>

The function creates a new semantic version object from a given version string, allowing for the structured handling and comparison of software versioning according to semantic versioning principles.

<table data-header-hidden><thead><tr><th width="174">Name</th><th>Value</th></tr></thead><tbody><tr><td>Signature</td><td><pre class="language-go"><code class="lang-go">Semver(version string) (*semver.Version, error)
</code></pre></td></tr></tbody></table>

{% tabs %}
{% tab title="Template Example" %}
```go
{{ semver "1.0.0" }} // Output: 1.0.0
{{ semver "1.0.0-alpha" }} // Output: 1.0.0-alpha
{{ (semver "2.1.0").Major }} // Output: 2
```
{% endtab %}
{% endtabs %}

### <mark style="color:purple;">semverCompare</mark>

The function checks whether a given version string satisfies a specified semantic version constraint, ensuring that the version meets the defined requirements according to the Semantic Versioning rules.

<table data-header-hidden><thead><tr><th width="164">Name</th><th>Value</th></tr></thead><tbody><tr><td>Signature</td><td><pre class="language-go"><code class="lang-go">SemverCompare(constraint, version string) (bool, error)
</code></pre></td></tr></tbody></table>

{% tabs %}
{% tab title="Template Example" %}
```go
{{ semverCompare ">=1.0.0" "1.0.0" }} // Output: true
{{ semverCompare "1.0.0" "1.0.0" }} // Output: true
{{ semverCompare "1.0.0" "1.0.1" }} // Output: false
{{ semverCompare "~1.0.0" "1.0.0" }} // Output: true
{{ semverCompare ">1.0.0-alpha" "1.0.0-alpha.1" }} // Output: true
{{ semverCompare "1.0.0-alpha.1" "1.0.0-alpha" }} // Output: false
{{ semverCompare "1.0.0-alpha.1" "1.0.0-alpha.1" }} // Output: true
```
{% endtab %}
{% endtabs %}
