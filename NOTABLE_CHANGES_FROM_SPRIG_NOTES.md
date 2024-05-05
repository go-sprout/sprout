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
