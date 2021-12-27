package counter_test

import (
	"testing"

	counter "github.com/teivah/100-go-mistakes/11-testing/90-testing-features/different-package"
)

func TestCount(t *testing.T) {
	if counter.Inc() != 1 {
		t.Errorf("expected 1")
	}
}
