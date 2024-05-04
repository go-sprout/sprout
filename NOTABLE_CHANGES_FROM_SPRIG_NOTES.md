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
