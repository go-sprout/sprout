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
</code></pre></td></tr><tr><td>Must version</td><td><span data-gb-custom-inline data-tag="emoji" data-code="274c">❌</span></td></tr></tbody></table>

{% tabs %}
{% tab title="Template Example" %}
```go
{{ 1, 2, 3 | list }} // Output: [1, 2, 3]
```
{% endtab %}
{% endtabs %}

### <mark style="color:purple;">append / mustAppend</mark>

The function adds an element to the end of an existing list, extending the list by one item.

<table data-header-hidden><thead><tr><th width="164">Name</th><th>Value</th></tr></thead><tbody><tr><td>Signature</td><td><pre class="language-go"><code class="lang-go">Append(list any, v any) []any
MustAppend(list any, v any) ([]any, error)
</code></pre></td></tr><tr><td>Must version</td><td><span data-gb-custom-inline data-tag="emoji" data-code="2705">✅</span></td></tr></tbody></table>

{% tabs %}
{% tab title="Template Example" %}
```go
{{ append ["a", "b"], "c" }} // Output: ["a", "b", "c"]
```
{% endtab %}

{% tab title="Must version" %}
```go
{{ mustAppend ["a", "b"], "c"  }} // Output: ["a", "b", "c"], nil
{{ mustAppend nil, "c"  }} // Output: nil, error
```
{% endtab %}
{% endtabs %}

### <mark style="color:purple;">prepend / mustPrepend</mark>

The function adds an element to the beginning of an existing list, placing the new item before all others.

<table data-header-hidden><thead><tr><th width="164">Name</th><th>Value</th></tr></thead><tbody><tr><td>Signature</td><td><pre class="language-go"><code class="lang-go">Prepend(list any, v any) []any
MustPrepend(list any, v any) ([]any, error)
</code></pre></td></tr><tr><td>Must version</td><td><span data-gb-custom-inline data-tag="emoji" data-code="2705">✅</span></td></tr></tbody></table>

{% tabs %}
{% tab title="Template Example" %}
```go
{{ prepend  ["b", "c"], "a" }} // Output: ["a", "b", "c"]
```
{% endtab %}

{% tab title="Must version" %}
```go
{{ mustPrepend ["b", "c"], "a" }} // Output: ["a", "b", "c"], nil
```
{% endtab %}
{% endtabs %}

### <mark style="color:purple;">concat</mark>

The function merges multiple lists into a single, unified list, combining all elements from the provided lists.

<table data-header-hidden><thead><tr><th width="164">Name</th><th>Value</th></tr></thead><tbody><tr><td>Signature</td><td><pre class="language-go"><code class="lang-go">Concat(lists ...any) any
</code></pre></td></tr><tr><td>Must version</td><td><span data-gb-custom-inline data-tag="emoji" data-code="274c">❌</span></td></tr></tbody></table>

{% tabs %}
{% tab title="Template Example" %}
```go
{{ ["c", "d"] | concat ["a", "b"] }} // Output: ["a", "b", "c", "d"]
```
{% endtab %}
{% endtabs %}

### <mark style="color:purple;">chunk / mustChunk</mark>

The function divides a list into smaller, equally sized chunks based on the specified size, breaking the original list into manageable sublists.

<table data-header-hidden><thead><tr><th width="164">Name</th><th>Value</th></tr></thead><tbody><tr><td>Signature</td><td><pre class="language-go"><code class="lang-go">Chunk(size int, list any) [][]any
MustChunk(size int, list any) ([][]any, error)
</code></pre></td></tr><tr><td>Must version</td><td><span data-gb-custom-inline data-tag="emoji" data-code="2705">✅</span></td></tr></tbody></table>

{% tabs %}
{% tab title="Template Example" %}
```go
{{ chunk 2, ["a", "b", "c", "d"] }} // Output: [["a", "b"], ["c", "d"]]
```
{% endtab %}

{% tab title="Must version" %}
```go
{{ ["a", "b", "c", "d"] | mustChunk 2 }}
// Output: [["a", "b"], ["c", "d"]], nil
{{ mustChunk 2 nil }}
// Output: nil, error
```
{% endtab %}
{% endtabs %}

### <mark style="color:purple;">uniq / mustUniq</mark>

The function removes duplicate elements from a list, ensuring that each element appears only once in the resulting list.

<table data-header-hidden><thead><tr><th width="164">Name</th><th>Value</th></tr></thead><tbody><tr><td>Signature</td><td><pre class="language-go"><code class="lang-go">Uniq(list any) []any
MustUniq(list any) ([]any, error)
</code></pre></td></tr><tr><td>Must version</td><td><span data-gb-custom-inline data-tag="emoji" data-code="2705">✅</span></td></tr></tbody></table>

{% tabs %}
{% tab title="Template Example" %}
```go
{{ ["a", "b", "a", "c"] | uniq }} // Output: ["a", "b", "c"]
{{ nil | uniq }} // Output: []
```
{% endtab %}

{% tab title="Must version" %}
```go
{{ ["a", "b", "a", "c"] | mustUniq }} // Output: ["a", "b", "c"], nil
{{ nil | mustUniq }} // Output: nil, error
```
{% endtab %}
{% endtabs %}

### <mark style="color:purple;">compact / mustCompact</mark>

The function removes `nil` and zero-value elements from a list, leaving only non-empty and meaningful values.

<table data-header-hidden><thead><tr><th width="164">Name</th><th>Value</th></tr></thead><tbody><tr><td>Signature</td><td><pre class="language-go"><code class="lang-go">Compact(list any) []any
MustCompact(list any) ([]any, error)
</code></pre></td></tr><tr><td>Must version</td><td><span data-gb-custom-inline data-tag="emoji" data-code="2705">✅</span></td></tr></tbody></table>

{% tabs %}
{% tab title="Template Example" %}
```go
{{ [0, 1, nil, 2, "", 3] | compact }} // Output: [1, 2, 3]
{{ nil | compact }} // Output: []
```
{% endtab %}

{% tab title="Must version" %}
```go
{{ [0, 1, nil, 2, "", 3] | mustCompact }} // Output: [1, 2, 3], nil
{{ nil | mustCompact }} // Output: nil, error
```
{% endtab %}
{% endtabs %}

### <mark style="color:purple;">slice / mustSlice</mark>

The function extracts a portion of a list, creating a new slice based on the specified start and end indices.

<table data-header-hidden><thead><tr><th width="164">Name</th><th>Value</th></tr></thead><tbody><tr><td>Signature</td><td><pre class="language-go"><code class="lang-go">Slice(list any, indices ...any) any
MustSlice(list any, indices ...any) (any, error)
</code></pre></td></tr><tr><td>Must version</td><td><span data-gb-custom-inline data-tag="emoji" data-code="2705">✅</span></td></tr></tbody></table>

{% tabs %}
{% tab title="Template Example" %}
```go
{{ slice [1, 2, 3, 4, 5], 1, 3 }} // Output: [2, 3]
{{ slice 1 }} // Output: []
```
{% endtab %}

{% tab title="Must version" %}
```go
{ mustSlice [1, 2, 3, 4, 5], 1, 3 }} // Output: [2, 3], nil
{{ mustSlice 1 }} // Output: nil, error
```
{% endtab %}
{% endtabs %}

### <mark style="color:purple;">has / mustHas</mark>

The function checks if a specified element is present within a collection, returning true if the element is found.

<table data-header-hidden><thead><tr><th width="164">Name</th><th>Value</th></tr></thead><tbody><tr><td>Signature</td><td><pre class="language-go"><code class="lang-go">Has(element any, list any) bool
MustHas(element any, list any) (bool, error)
</code></pre></td></tr><tr><td>Must version</td><td><span data-gb-custom-inline data-tag="emoji" data-code="2705">✅</span></td></tr></tbody></table>

{% tabs %}
{% tab title="Template Example" %}
```go
{{ ["value", "other"] | has "value" }} // Output: true
```
{% endtab %}

{% tab title="Must version" %}
```go
{{ [1, 2, 3, 4] | mustHas 3 }} // Output: true, nil
{{ 3 | mustHas 3 }} // Output: false, error
```
{% endtab %}
{% endtabs %}

### <mark style="color:purple;">without / mustWithout</mark>

The function returns a new list that excludes the specified elements, effectively filtering out unwanted items from the original list.

<table data-header-hidden><thead><tr><th width="164">Name</th><th>Value</th></tr></thead><tbody><tr><td>Signature</td><td><pre class="language-go"><code class="lang-go">Without(list any, omit ...any) []any
MustWithout(list any, omit ...any) ([]any, error)
</code></pre></td></tr><tr><td>Must version</td><td><span data-gb-custom-inline data-tag="emoji" data-code="2705">✅</span></td></tr></tbody></table>

{% tabs %}
{% tab title="Template Example" %}
```go
{{ without [1, 2, 3, 4], 2, 4 }} // Output: [1, 3]
{{ without nil, nil }} // Output: []
```
{% endtab %}

{% tab title="Must version" %}
```go
{{ mustWithout [1, 2, 3, 4], 2, 4 }} // Output: [1, 3], nil
{{ mustWithout nil, nil }} // Output: nil, error
```
{% endtab %}
{% endtabs %}

### <mark style="color:purple;">rest / mustRest</mark>

The function returns all elements of a list except for the first one, effectively giving you the "rest" of the list.

<table data-header-hidden><thead><tr><th width="164">Name</th><th>Value</th></tr></thead><tbody><tr><td>Signature</td><td><pre class="language-go"><code class="lang-go">Rest(list any) []any
MustRest(list any) ([]any, error)
</code></pre></td></tr><tr><td>Must version</td><td><span data-gb-custom-inline data-tag="emoji" data-code="2705">✅</span></td></tr></tbody></table>

{% tabs %}
{% tab title="Template Example" %}
```go
{{ [1, 2, 3, 4] | rest }} // Output: [2, 3, 4]
{{ rest nil }} // Output: []
```
{% endtab %}

{% tab title="Must version" %}
```go
{{ [1, 2, 3, 4] | mustRest }} // Output: [2, 3, 4], nil
{{ mustRest 1 }} // Output: nil, error
```
{% endtab %}
{% endtabs %}

### <mark style="color:purple;">initial / mustInitial</mark>

The function returns all elements of a list except the last one, effectively providing the "initial" portion of the list.

<table data-header-hidden><thead><tr><th width="164">Name</th><th>Value</th></tr></thead><tbody><tr><td>Signature</td><td><pre class="language-go"><code class="lang-go">Initial(list any) []any
MustInitial(list any) ([]any, error)
</code></pre></td></tr><tr><td>Must version</td><td><span data-gb-custom-inline data-tag="emoji" data-code="2705">✅</span></td></tr></tbody></table>

{% tabs %}
{% tab title="Template Example" %}
```go
{{ [1, 2, 3, 4] | initial }} // Output: [1, 2, 3]
{{ initial 1 }} // Output: []
```
{% endtab %}

{% tab title="Must version" %}
```go
{{ [1, 2, 3, 4] | mustInitial }} // Output: [1, 2, 3], nil
{{ mustInitial 1 }} // Output: nil, error
```
{% endtab %}
{% endtabs %}

### <mark style="color:purple;">first / mustFirst</mark>

The function returns the first element of a list.

<table data-header-hidden><thead><tr><th width="164">Name</th><th>Value</th></tr></thead><tbody><tr><td>Signature</td><td><pre class="language-go"><code class="lang-go">First(list any) any
<strong>MustFirst(list any) (any, error)
</strong></code></pre></td></tr><tr><td>Must version</td><td><span data-gb-custom-inline data-tag="emoji" data-code="2705">✅</span></td></tr></tbody></table>

{% tabs %}
{% tab title="Template Example" %}
```go
{{ [1, 2, 3, 4] | first }} // Output: 1
{{ first 1 }} // Output: nil
```
{% endtab %}

{% tab title="Must version" %}
```go
{{ [1, 2, 3, 4] | mustFirst }} // Output: 1, nil
{{ mustFirst 1 }} // Output: nil, error
```
{% endtab %}
{% endtabs %}

### <mark style="color:purple;">last / mustLast</mark>

The function returns the last element of a list.

<table data-header-hidden><thead><tr><th width="164">Name</th><th>Value</th></tr></thead><tbody><tr><td>Signature</td><td><pre class="language-go"><code class="lang-go">Last(list any) any
MustLast(list any) (any, error)
</code></pre></td></tr><tr><td>Must version</td><td><span data-gb-custom-inline data-tag="emoji" data-code="2705">✅</span></td></tr></tbody></table>

{% tabs %}
{% tab title="Template Example" %}
```go
{{ [1, 2, 3, 4] | last }} // Output: 4
{{ last 1 }} // Output: nil
```
{% endtab %}

{% tab title="Must version" %}
```go
{{ [1, 2, 3, 4] | mustLast }} // Output: 4, nil
{{ mustLast 1 }} // Output: nil, error
```
{% endtab %}
{% endtabs %}

### <mark style="color:purple;">reverse / mustReverse</mark>

The function returns a new list with the elements in reverse order, flipping the sequence of the original list.

<table data-header-hidden><thead><tr><th width="164">Name</th><th>Value</th></tr></thead><tbody><tr><td>Signature</td><td><pre class="language-go"><code class="lang-go">Reverse(list any) []any
MustReverse(list any) ([]any, error)
</code></pre></td></tr><tr><td>Must version</td><td><span data-gb-custom-inline data-tag="emoji" data-code="2705">✅</span></td></tr></tbody></table>

{% tabs %}
{% tab title="Template Example" %}
```go
{{ [1, 2, 3, 4] | reverse }} // Output: [4, 3, 2, 1]
{{ reverse nil }} // Output: []
```
{% endtab %}

{% tab title="Must version" %}
```go
{{ [1, 2, 3, 4] | mustReverse }} // Output: [4, 3, 2, 1], nil
{{ mustReverse nil }} // Output: nil, error
```
{% endtab %}
{% endtabs %}

### <mark style="color:purple;">sortAlpha</mark>

The function sorts a list of strings in alphabetical order.

<table data-header-hidden><thead><tr><th width="164">Name</th><th>Value</th></tr></thead><tbody><tr><td>Signature</td><td><pre class="language-go"><code class="lang-go">SortAlpha(list any) []string
</code></pre></td></tr><tr><td>Must version</td><td><span data-gb-custom-inline data-tag="emoji" data-code="274c">❌</span></td></tr></tbody></table>

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
</code></pre></td></tr><tr><td>Must version</td><td><span data-gb-custom-inline data-tag="emoji" data-code="274c">❌</span></td></tr></tbody></table>

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
</code></pre></td></tr><tr><td>Must version</td><td><span data-gb-custom-inline data-tag="emoji" data-code="274c">❌</span></td></tr></tbody></table>

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
</code></pre></td></tr><tr><td>Must version</td><td><span data-gb-custom-inline data-tag="emoji" data-code="274c">❌</span></td></tr></tbody></table>

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
</code></pre></td></tr><tr><td>Must version</td><td><span data-gb-custom-inline data-tag="emoji" data-code="274c">❌</span></td></tr></tbody></table>

{% tabs %}
{% tab title="Template Example" %}
```go
{{ 0, 10, 2 | untilStep }} // Output: [0 2 4 6 8]
{{ 10, 0, -2 | untilStep }} // Output: [10 8 6 4 2]
```
{% endtab %}
{% endtabs %}
