package learn

import (
	"testing"
)

// TestUnsafe test unsafe
func TestUnsafe(t *testing.T) {
	s := UnsafeChangeStructValue()
	if s.age != 12 {
		t.Errorf("expected age is 12, but get %d", s.age)
	}

	if s.name != "chaojie" {
		t.Errorf("expected age is 12, but get %s", s.name)
	}
}

// TestTicker test ticker function
// use go test -v
func TestTicker(t *testing.T) {
	Ticker()
}

func TestWaitGroup(t *testing.T) {
	wg()
}

// /usr/local/go/bin/go test -timeout 30s -run ^TestWaitGroup$ github.com/shiniao/learn -v
func TestSignal(t *testing.T) {
	mySignal()
}

// Sync pool test and benchmark
func BenchmarkWithoutPool(b *testing.B) {
	var s *Small
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for i := 0; i < 1000; i++ {
			s = new(Small)
			s.a++
		}
	}
}

func BenchmarkWithPool(b *testing.B) {
	var s *Small
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for i := 0; i < 1000; i++ {
			s = NewSmallPool()
			s.a++
			s.Free()
		}
	}
}
// bench result:
// go test -bench . -run=none -benchmem
// goos: linux
// goarch: amd64
// pkg: github.com/shiniao/learn
// BenchmarkWithoutPool-8             80280             15546 ns/op            8000 B/op       1000 allocs/op
// BenchmarkWithPool-8                81262             15556 ns/op               0 B/op          0 allocs/op
// PASS
// ok      github.com/shiniao/learn        2.826s
