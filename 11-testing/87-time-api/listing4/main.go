package listing1

import (
	"sync"
	"time"
)

type Cache struct {
	mu     sync.RWMutex
	events []Event
}

type Event struct {
	Timestamp time.Time
	Data      string
}

func (c *Cache) TrimOlderThan(t time.Time) {
	c.mu.RLock()
	defer c.mu.RUnlock()

	for i := 0; i < len(c.events); i++ {
		if c.events[i].Timestamp.After(t) {
			c.events = c.events[i:]
			return
		}
	}
}

func (c *Cache) Add(events []Event) {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.events = append(c.events, events...)
}

func (c *Cache) GetAll() []Event {
	c.mu.RLock()
	defer c.mu.RUnlock()

	return c.events
}
