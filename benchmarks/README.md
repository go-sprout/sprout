# Benchmarks outputs

## Sprig v3.2.3 vs Sprout v0.2
```
go test -count=1 -bench ^Benchmark -benchmem -cpuprofile cpu.out -memprofile mem.out
goos: linux
goarch: amd64
pkg: sprout_benchmarks
cpu: Intel(R) Core(TM) i7-9750H CPU @ 2.60GHz
BenchmarkSprig-12              1        3869134593 ns/op        45438616 B/op      24098 allocs/op
BenchmarkSprout-12             1        1814126036 ns/op        38284040 B/op      11627 allocs/op
PASS
ok      sprout_benchmarks       5.910s
```

**Time improvement**: ((3869134593 - 1814126036) / 3869134593) * 100 = 53.1%
**Memory improvement**: ((45438616 - 38284040) / 45438616) * 100 = 15.7%

So, Sprout v0.2 is approximately 53.1% faster and uses 15.7% less memory than Sprig v3.2.3.

## Sprig v3.2.3 vs Sprout v0.3

```
go test -count=1 -bench ^Benchmark -benchmem -cpuprofile cpu.out -memprofile mem.out
goos: linux
goarch: amd64
pkg: sprout_benchmarks
cpu: Intel(R) Core(TM) i9-9900K CPU @ 3.60GHz
BenchmarkSprig-16              1        2152506421 ns/op        44671136 B/op      21938 allocs/op
BenchmarkSprout-16             1        1020721871 ns/op        37916984 B/op      11173 allocs/op
PASS
ok      sprout_benchmarks       3.720s
```

**Time improvement**: ((2152506421 - 1020721871) / 2152506421) * 100 = 52.6%
**Memory improvement**: ((44671136 - 37916984) / 44671136) * 100 = 15.1%

So, Sprout v0.3 is approximately 52.6% faster and uses 15.1% less memory than Sprig v3.2.3.
