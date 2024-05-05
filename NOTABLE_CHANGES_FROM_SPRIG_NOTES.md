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
