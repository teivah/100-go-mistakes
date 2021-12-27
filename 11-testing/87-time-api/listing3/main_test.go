package listing1

import (
	"testing"
	"time"
)

func TestCache_TrimBefore(t *testing.T) {
	events := []Event{
		{Timestamp: time.Now().Add(-20 * time.Millisecond)},
		{Timestamp: time.Now().Add(-10 * time.Millisecond)},
		{Timestamp: time.Now().Add(10 * time.Millisecond)},
	}
	cache := &Cache{}
	cache.Add(events)
	cache.TrimBefore(parseTime(t, "2020-01-01T12:00:00.06Z"), 15*time.Millisecond)
	got := cache.GetAll()
	expected := 2
	if len(got) != expected {
		t.Fatalf("expected %d, got %d", expected, len(got))
	}
}

func parseTime(t *testing.T, timestamp string) time.Time {
	ts, err := time.Parse(time.RFC3339, timestamp)
	if err != nil {
		t.FailNow()
	}
	return ts
}
