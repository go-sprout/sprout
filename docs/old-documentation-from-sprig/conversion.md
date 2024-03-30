# Type Conversion Functions

The following type conversion functions are provided by Sprig:

* `toDecimal`: Convert a unix octal to a `int64`.
* `toStrings`: Convert a list, slice, or array to a list of strings.

## toStrings

Given a list-like collection, produce a slice of strings.

```
list 1 2 3 | toStrings
```

The above converts `1` to `"1"`, `2` to `"2"`, and so on, and then returns them as a list.

## toDecimal

Given a unix octal permission, produce a decimal.

```
"0777" | toDecimal
```

The above converts `0777` to `511` and returns the value as an int64.
