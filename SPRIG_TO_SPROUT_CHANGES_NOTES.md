# Migration Notes for Sprout Library
This document outlines the key differences and migration changes between the
Sprig and Sprout libraries. The changes are designed to enhance stability and 
usability in the Sprout library.

This document will help contributors and maintainers understand the changes made
between the fork date and version 1.0.0 of the Sprout library. 

It will be updated to reflect changes in future versions of the library. 

This document will also assist in creating the migration guide when version 1.0 is ready.


## Error Handling Enhancements
### General Error Handling
In Sprig, errors within certain functions cause a panic. 
In contrast, Sprout opts for returning nil or an empty value, improving safety
and predictability.

**Old Behavior (Sprig)**: Triggers a panic on error
```go
if err != nil {
  panic("deepCopy error: " + err.Error())
}
```

**New Behavior (Sprout)**: Returns nil or an empty value on error
```go
if err != nil {
  return nil, err
}
```

Methods that previously caused a panic in Sprig :
- DeepCopy
- MustDeepCopy
- ToRawJson
- Append
- Prepend
- Concat
- Chunk
- Uniq
- Compact
- Slice
- Without
- Rest
- Initial
- Reverse
- First
- Last
- Has
- Dig
- RandAlphaNumeric
- RandAlpha
- RandAscii
- RandNumeric
- RandBytes

## Function-Specific Changes

### MustDeepCopy

- **Sprig**: Accepts `nil` input, causing an internal panic.
- **Sprout**: Returns `nil` if input is `nil`, avoiding panic.

## Rand Functions

- **Sprig**: Causes an internal panic if the length parameter is zero.
- **Sprout**: Returns an empty string if the length is zero, ensuring stability.

## DateAgo

- **Sprig**: Does not support int32 and *time.Time; returns "0s".
- **Sprout**: Supports int32 and *time.Time and returns the correct duration.

## DateRound
- **Sprig**: Returns a corrected duration in positive form, even for negative inputs.
- **Sprout**: Accurately returns the duration, preserving the sign of the input.

## Base32Decode / Base64Decode
- **Sprig**: Decoding functions return the error string when the input is not a valid base64 encoded string.
- **Sprout**: Decoding functions return an empty string if the input is not a valid base64 encoded string, simplifying error handling.

## Dig 
> Consider the example dictionary defined as follows:
> ```go
> dict := map[string]any{
>   "a": map[string]any{
>     "b": 2,
>   },
> }
> ```

- **Sprig**: Previously, the `dig` function would return the last map in the access chain.
```go
{{ $dict | dig "a" "b" }} // Output: map[b:2]
```
- **Sprout**: Now, the `dig` function returns the final object in the chain, regardless of its type (map, array, string, etc.).
```go
{{ $dict | dig "a" "b" }} // Output: 2
```

## ToCamelCase / ToPascalCase
- **Sprig**: The `toCamelCase` return value are in PascalCase. No `toPascalCase` function is available.
- **Sprout**: The `toCamelCase` function returns camelCase strings, while the `toPascalCase` function returns PascalCase strings.

## Merge / MergeOverwrite
- **Sprig**: The `merge` and `mergeOverwrite` functions does dereferencing when second value are the default golang value (example: `0` for int).
- **Sprout**: The `merge` and `mergeOverwrite` functions does not dereference and keep the second value as is (example: `0` for int).
