package listing1

import (
	"testing"
	"time"
)

func TestCache_TrimOlderThan(t *testing.T) {
	events := []Event{
		{Timestamp: time.Now().Add(-20 * time.Millisecond)},
		{Timestamp: time.Now().Add(-10 * time.Millisecond)},
		{Timestamp: time.Now().Add(10 * time.Millisecond)},
	}
	cache := &Cache{}
	cache.Add(events)
	cache.TrimOlderThan(15 * time.Millisecond)
	got := cache.GetAll()
	expected := 2
	if len(got) != expected {
		t.Fatalf("expected %d, got %d", expected, len(got))
	}
}
