# Migration notes

## DeepCopy

```go
if err != nil {
	panic("deepCopy error: " + err.Error())
}
```
changed to 
```go
if err != nil {
	return nil
}
```

## MustDeepCopy
In sprig MustDeepCopy accept nil as input and cause internal panic. In sprout MustDeepCopy return nil if input is nil.


## all rand functions
In sprig all rand functions cause internal panic if the length is equal to 0. In sprout all rand functions return empty string if the length is equal to 0.

<!-- ## Encoding Decode
In sprig Decoding functions return the error string instead of empty string if the input is not a valid base64 encoded string.
In sprout Decoding functions return empty string if the input is not a valid base64 encoded string. -->

## ToRawJson
in sprig code panic
```go
func toRawJson(v interface{}) string {
	output, err := mustToRawJson(v)
	if err != nil {
		panic(err)
	}
	return string(output)
}
```

in sprout code follow the same pattern as other functions

```go
func (fh *FunctionHandler) ToRawJson(v any) string {
	output, _ := fh.MustToRawJson(v)
	return output
}
```

### DateAgo

In sprig this function dont support int32 and *time.Time and cause result to "0s"
In sprout this function support int32 and *time.Time and return the correct result

### DateRound 
In sprig When we pass a negative value, it will return the correct duration but in positive value.
In sprout When we pass a negative value, it will return the correct duration with in negative value.

### Append, Prepend, Concat, Chunk, Uniq, Compact, Slice, Without, Rest, Initial, Reverse
In sprig all these functions cause internal panic.
In sprout all these functions return empty slice when an error occurs.

### First, Last
In sprig all these functions cause internal panic.
In sprout all these functions return nil when an error occurs.

### MustAppend, MustPrepend, MustConcat, MustChunk, MustUniq, MustCompact, MustSlice, MustWithout, MustRest, MustInitial, MustReverse
In sprig all these functions cause segfault when lsit are nil.
In sprout all these functions return nil and an error when an error occurs.

### Has, Dig
In sprig this function cause internal panic.
In sprout this function return false when an error occurs.
