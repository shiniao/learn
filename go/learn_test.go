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

// Sync
var urls = []string{
	"http://www.golang.org/",
	"http://www.google.com/",
	"http://www.somestupidname.com/",
}

func TestSyncPool(t *testing.T) {

	for _, url := range urls {
		syncPool(url)
	}

}

func TestSyncNoPool(t *testing.T) {

	for _, url := range urls {
		syncNoPool(url)
	}
}

func BenchmarkSyncPool(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		// for _, url := range urls {
		// 	syncPool(url)
		// }
		syncPool(urls[0])
	}
}

func BenchmarkSyncNoPool(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		// for _, url := range urls {
		// 	syncNoPool(url)
		// }
		syncNoPool(urls[0])
	}
}
