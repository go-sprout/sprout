bench:
	go test -run=^$$ -bench ^Benchmark -benchmem -cpuprofile cpu.out -memprofile mem.out

compare:
	go test -cpuprofile cpu.out -memprofile mem.out
