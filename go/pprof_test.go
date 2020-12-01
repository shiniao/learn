package learn

import "testing"

func Fib(n int) int {
	if n < 2 {
		return n
	}
	return Fib(n-1) + Fib(n-2)
}

func BenchmarkFib(b *testing.B)  {
	for i := 0; i < b.N; i++ {
		Fib(10)
	}
}

// pprof 性能测试工具

// go test --run=XXX -bench . -benchmem -cpuprofile profile.out -memprofile memprofile.out

// 在 web 中查看
// go tool pprof -http=:5000 profile.out
