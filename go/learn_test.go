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
