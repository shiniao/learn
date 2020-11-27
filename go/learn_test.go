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
}
