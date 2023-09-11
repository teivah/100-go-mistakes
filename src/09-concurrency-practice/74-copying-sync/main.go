package main

import (
	"sync"
	"time"
)

func main() {
	counter := NewCounter()

	go func() {
		counter.Increment1("foo")
	}()
	go func() {
		counter.Increment1("bar")
	}()

	time.Sleep(10 * time.Millisecond)
}

type Counter struct {
	mu       sync.Mutex
	counters map[string]int
}

func NewCounter() Counter {
	return Counter{counters: map[string]int{}}
}

func (c Counter) Increment1(name string) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.counters[name]++
}

func (c *Counter) Increment2(name string) {
	// Same code
}

type Counter2 struct {
	mu       *sync.Mutex
	counters map[string]int
}

func NewCounter2() Counter2 {
	return Counter2{
		mu:       &sync.Mutex{},
		counters: map[string]int{},
	}
}
