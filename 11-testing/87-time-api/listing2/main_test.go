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
	cache := &Cache{now: func() time.Time {
		return parseTime(t, "2020-01-01T12:00:00.06Z")
	}}
	cache.Add(events)
	cache.TrimOlderThan(15 * time.Millisecond)
	// ...
}

func parseTime(t *testing.T, timestamp string) time.Time {
	ts, err := time.Parse(time.RFC3339, timestamp)
	if err != nil {
		t.FailNow()
	}
	return ts
}
