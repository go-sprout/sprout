# Type Conversion Functions

The following type conversion functions are provided by Sprig:

* `toStrings`: Convert a list, slice, or array to a list of strings.

## toStrings

Given a list-like collection, produce a slice of strings.

```
list 1 2 3 | toStrings
```

The above converts `1` to `"1"`, `2` to `"2"`, and so on, and then returns them as a list.

