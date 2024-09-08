---
description: >-
  The Slices registry provides utilities for working with slice data structures,
  including functions for filtering, sorting, and transforming slices in a
  flexible manner.
---

# Slices

{% hint style="info" %}
You can easily import all the functions from the <mark style="color:yellow;">`slices`</mark> registry by including the following import statement in your code

```go
import "github.com/go-sprout/sprout/registry/slices"
```
{% endhint %}

### <mark style="color:purple;">list</mark>

The function creates a list from the provided elements, collecting them into a single array-like structure.

<table data-header-hidden><thead><tr><th width="174">Name</th><th>Value</th></tr></thead><tbody><tr><td>Signature</td><td><pre class="language-go"><code class="lang-go">List(values ...any) []any
</code></pre></td></tr></tbody></table>

{% tabs %}
{% tab title="Template Example" %}
```go
{{ 1, 2, 3 | list }} // Output: [1, 2, 3]
```
{% endtab %}
{% endtabs %}

### <mark style="color:purple;">append</mark>

The function adds an element to the end of an existing list, extending the list by one item.

<table data-header-hidden><thead><tr><th width="164">Name</th><th>Value</th></tr></thead><tbody><tr><td>Signature</td><td><pre class="language-go"><code class="lang-go">Append(v any, list any) ([]any, error)
</code></pre></td></tr></tbody></table>

{% tabs %}
{% tab title="Template Example" %}
```go
{{ ["a", "b"] | append "c"  }} // Output: ["a", "b", "c"], nil
{{ nil | append "c"  }} // Error
```
{% endtab %}
{% endtabs %}

### <mark style="color:purple;">prepend</mark>

The function adds an element to the beginning of an existing list, placing the new item before all others.

<table data-header-hidden><thead><tr><th width="164">Name</th><th>Value</th></tr></thead><tbody><tr><td>Signature</td><td><pre class="language-go"><code class="lang-go">Prepend(v any, list any) ([]any, error)
</code></pre></td></tr></tbody></table>

{% tabs %}
{% tab title="Template Example" %}
```go
{{ ["b", "c"] | prepend "a" }} // Output: ["a", "b", "c"]
{{ nil | prepend "c" }} // Error
```
{% endtab %}
{% endtabs %}

### <mark style="color:purple;">concat</mark>

The function merges multiple lists into a single, unified list, combining all elements from the provided lists.

<table data-header-hidden><thead><tr><th width="164">Name</th><th>Value</th></tr></thead><tbody><tr><td>Signature</td><td><pre class="language-go"><code class="lang-go">Concat(lists ...any) any
</code></pre></td></tr></tbody></table>

{% tabs %}
{% tab title="Template Example" %}
```go
{{ ["c", "d"] | concat ["a", "b"] }} // Output: ["a", "b", "c", "d"]
```
{% endtab %}
{% endtabs %}

### <mark style="color:purple;">chunk</mark>

The function divides a list into smaller, equally sized chunks based on the specified size, breaking the original list into manageable sublists.

<table data-header-hidden><thead><tr><th width="164">Name</th><th>Value</th></tr></thead><tbody><tr><td>Signature</td><td><pre class="language-go"><code class="lang-go">Chunk(size int, list any) ([][]any, error)
</code></pre></td></tr></tbody></table>

{% tabs %}
{% tab title="Template Example" %}
```go
{{ ["a", "b", "c", "d"] | chunk 2 }}
// Output: [["a", "b"], ["c", "d"]], nil
{{ chunk 2 nil }}
// Output: nil, error
```
{% endtab %}
{% endtabs %}

### <mark style="color:purple;">uniq</mark>

The function removes duplicate elements from a list, ensuring that each element appears only once in the resulting list.

<table data-header-hidden><thead><tr><th width="164">Name</th><th>Value</th></tr></thead><tbody><tr><td>Signature</td><td><pre class="language-go"><code class="lang-go">Uniq(list any) ([]any, error)
</code></pre></td></tr></tbody></table>

{% tabs %}
{% tab title="Template Example" %}
```go
{{ ["a", "b", "a", "c"] | mustUniq }} // Output: ["a", "b", "c"], nil
{{ nil | mustUniq }} // Error
```
{% endtab %}
{% endtabs %}

### <mark style="color:purple;">compact</mark>

The function removes `nil` and zero-value elements from a list, leaving only non-empty and meaningful values.

<table data-header-hidden><thead><tr><th width="164">Name</th><th>Value</th></tr></thead><tbody><tr><td>Signature</td><td><pre class="language-go"><code class="lang-go">Compact(list any) ([]any, error)
</code></pre></td></tr></tbody></table>

{% tabs %}
{% tab title="Template Example" %}
```go
{{ [0, 1, nil, 2, "", 3] | compact }} // Output: [1, 2, 3], nil
{{ nil | compact }} // Error
```
{% endtab %}
{% endtabs %}

### <mark style="color:purple;">slice</mark>

The function extracts a portion of a list, creating a new slice based on the specified start and end indices.

<table data-header-hidden><thead><tr><th width="164">Name</th><th>Value</th></tr></thead><tbody><tr><td>Signature</td><td><pre class="language-go"><code class="lang-go">Slice(indices ...any, list any) (any, error)
</code></pre></td></tr></tbody></table>

{% tabs %}
{% tab title="Template Example" %}
```go
{{ [1, 2, 3, 4, 5] | slice 1, 3 }} // Output: [2, 3]
{{ slice 1 }} // Error
```
{% endtab %}
{% endtabs %}

### <mark style="color:purple;">has</mark>

The function checks if a specified element is present within a collection, returning true if the element is found.

<table data-header-hidden><thead><tr><th width="164">Name</th><th>Value</th></tr></thead><tbody><tr><td>Signature</td><td><pre class="language-go"><code class="lang-go">Has(element any, list any) (bool, error)
</code></pre></td></tr></tbody></table>

{% tabs %}
{% tab title="Template Example" %}
```go
{{ [1, 2, 3, 4] | has 3 }} // Output: true, nil
{{ 3 | has 3 }} // Error
```
{% endtab %}
{% endtabs %}

### <mark style="color:purple;">without</mark>

The function returns a new list that excludes the specified elements, effectively filtering out unwanted items from the original list.

<table data-header-hidden><thead><tr><th width="164">Name</th><th>Value</th></tr></thead><tbody><tr><td>Signature</td><td><pre class="language-go"><code class="lang-go">Without(omit ...any, list any) ([]any, error)
</code></pre></td></tr></tbody></table>

{% tabs %}
{% tab title="Template Example" %}
```go
{{ [1, 2, 3, 4] | without 2, 4 }} // Output: [1, 3], nil
{{ without nil, nil }} // Error
```
{% endtab %}
{% endtabs %}

### <mark style="color:purple;">rest</mark>

The function returns all elements of a list except for the first one, effectively giving you the "rest" of the list.

<table data-header-hidden><thead><tr><th width="164">Name</th><th>Value</th></tr></thead><tbody><tr><td>Signature</td><td><pre class="language-go"><code class="lang-go">Rest(list any) ([]any, error)
</code></pre></td></tr></tbody></table>

{% tabs %}
{% tab title="Template Example" %}
```go
{{ [1, 2, 3, 4] | rest }} // Output: [2, 3, 4], nil
{{ rest 1 }} // Error
```
{% endtab %}
{% endtabs %}

### <mark style="color:purple;">initial</mark>

The function returns all elements of a list except the last one, effectively providing the "initial" portion of the list.

<table data-header-hidden><thead><tr><th width="164">Name</th><th>Value</th></tr></thead><tbody><tr><td>Signature</td><td><pre class="language-go"><code class="lang-go">Initial(list any) ([]any, error)
</code></pre></td></tr></tbody></table>

{% tabs %}
{% tab title="Template Example" %}
```go
{{ [1, 2, 3, 4] | initial }} // Output: [1, 2, 3]
{{ initial 1 }} // Error
```
{% endtab %}
{% endtabs %}

### <mark style="color:purple;">first</mark>

The function returns the first element of a list.

<table data-header-hidden><thead><tr><th width="164">Name</th><th>Value</th></tr></thead><tbody><tr><td>Signature</td><td><pre class="language-go"><code class="lang-go">First(list any) (any, error)
</code></pre></td></tr></tbody></table>

{% tabs %}
{% tab title="Template Example" %}
```go
{{ [1, 2, 3, 4] | first }} // Output: 1
{{ first nil }} // Error
```
{% endtab %}
{% endtabs %}

### <mark style="color:purple;">last</mark>

The function returns the last element of a list.

<table data-header-hidden><thead><tr><th width="164">Name</th><th>Value</th></tr></thead><tbody><tr><td>Signature</td><td><pre class="language-go"><code class="lang-go">Last(list any) (any, error)
</code></pre></td></tr></tbody></table>

{% tabs %}
{% tab title="Template Example" %}
```go
{{ [1, 2, 3, 4] | last }} // Output: 4
{{ last nil }} // Error
```
{% endtab %}
{% endtabs %}

### <mark style="color:purple;">reverse</mark>

The function returns a new list with the elements in reverse order, flipping the sequence of the original list.

<table data-header-hidden><thead><tr><th width="164">Name</th><th>Value</th></tr></thead><tbody><tr><td>Signature</td><td><pre class="language-go"><code class="lang-go">Reverse(list any) ([]any, error)
</code></pre></td></tr></tbody></table>

{% tabs %}
{% tab title="Template Example" %}
```go
{{ [1, 2, 3, 4] | reverse }} // Output: [4, 3, 2, 1]
{{ reverse nil }} // Error
```
{% endtab %}
{% endtabs %}

### <mark style="color:purple;">sortAlpha</mark>

The function sorts a list of strings in alphabetical order.

<table data-header-hidden><thead><tr><th width="164">Name</th><th>Value</th></tr></thead><tbody><tr><td>Signature</td><td><pre class="language-go"><code class="lang-go">SortAlpha(list any) []string
</code></pre></td></tr></tbody></table>

{% tabs %}
{% tab title="Template Example" %}
```go
{{ ["d", "b", "a", "c"] | sortAlpha }} // Output: ["a", "b", "c", "d"]
{{ [4, 3, 2, 1, "a"] | sortAlpha }} // Output: ["1", "2", "3", "4", "a"]
```
{% endtab %}
{% endtabs %}

### <mark style="color:purple;">splitList</mark>

The function splits a string into a slice of substrings based on the specified separator.

{% hint style="warning" %}
This function may be renamed in the future to better reflect its purpose.
{% endhint %}

<table data-header-hidden><thead><tr><th width="164">Name</th><th>Value</th></tr></thead><tbody><tr><td>Signature</td><td><pre class="language-go"><code class="lang-go">SplitList(sep string, str string) []string
</code></pre></td></tr></tbody></table>

{% tabs %}
{% tab title="Template Example" %}
```go
{{ "one, two, three" | splitList ", " }} // Output: ["one", "two", "three"]
```
{% endtab %}
{% endtabs %}

### <mark style="color:purple;">strSlice</mark>

The function converts a given value into a slice of strings, handling various input types including `[]string`, `[]any`, and other slice types, ensuring flexible type conversion to a string slice.

<table data-header-hidden><thead><tr><th width="164">Name</th><th>Value</th></tr></thead><tbody><tr><td>Signature</td><td><pre class="language-go"><code class="lang-go">StrSlice(value any) []string
</code></pre></td></tr></tbody></table>

{% tabs %}
{% tab title="Template Example" %}
```go
{{ strSlice ["a", "b", "c"] }} // Output: ["a", "b", "c"]
{{ strSlice [5, "a", true, nil, 1] }} // Output: ["5", "a", "true", "1"]
```
{% endtab %}
{% endtabs %}

### <mark style="color:purple;">until</mark>

The function generates a slice of integers starting from 0 up to, but not including, the specified `count`. If `count` is negative, it produces a descending slice from 0 down to `count`, inclusive, stepping by -1. The function utilizes [`UntilStep`](slices.md#untilstep) to dynamically determine the range and step size.

<table data-header-hidden><thead><tr><th width="164">Name</th><th>Value</th></tr></thead><tbody><tr><td>Signature</td><td><pre class="language-go"><code class="lang-go">Until(count int) []int
</code></pre></td></tr></tbody></table>

{% tabs %}
{% tab title="Template Example" %}
```go
{{ 5 | until }} // Output: [0 1 2 3 4]
{{ -3 | until }} // Output: [0 -1 -2]
```
{% endtab %}
{% endtabs %}

### <mark style="color:purple;">untilStep</mark>

The function generates a slice of integers from `start` to `stop` (exclusive), incrementing by the specified `step`. If `step` is positive, the sequence ascends; if negative, it descends. The function returns an empty slice if the sequence is logically invalid, such as when a positive step is used but `start` is greater than `stop`, or vice versa.

<table data-header-hidden><thead><tr><th width="164">Name</th><th>Value</th></tr></thead><tbody><tr><td>Signature</td><td><pre class="language-go"><code class="lang-go">UntilStep(start, stop, step int) []int
</code></pre></td></tr></tbody></table>

{% tabs %}
{% tab title="Template Example" %}
```go
{{ 0, 10, 2 | untilStep }} // Output: [0 2 4 6 8]
{{ 10, 0, -2 | untilStep }} // Output: [10 8 6 4 2]
```
{% endtab %}
{% endtabs %}
