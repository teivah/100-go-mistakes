package listing1

import (
	"testing"
	"time"
)

func TestCache_TrimOlderThan(t *testing.T) {
	events := []Event{
		{Timestamp: parseTime(t, "2020-01-01T12:00:00.04Z")},
		{Timestamp: parseTime(t, "2020-01-01T12:00:00.05Z")},
		{Timestamp: parseTime(t, "2020-01-01T12:00:00.06Z")},
	}
	cache := &Cache{}
	cache.Add(events)
	cache.TrimOlderThan(parseTime(t, "2020-01-01T12:00:00.06Z"), 15*time.Millisecond)
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
