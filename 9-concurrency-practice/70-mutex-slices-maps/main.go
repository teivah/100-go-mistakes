package main

import "sync"

type Cache struct {
	mu       sync.Mutex
	balances map[string]float64
}

func (c *Cache) AddBalance(id string, balance float64) {
	c.mu.Lock()
	c.balances[id] = balance
	c.mu.Unlock()
}

func (c *Cache) AverageBalance1() float64 {
	c.mu.Lock()
	balances := c.balances
	c.mu.Unlock()

	sum := 0.
	for _, balance := range balances {
		sum += balance
	}
	return sum / float64(len(balances))
}

func (c *Cache) AverageBalance2() float64 {
	c.mu.Lock()
	defer c.mu.Unlock()

	sum := 0.
	for _, balance := range c.balances {
		sum += balance
	}
	return sum / float64(len(c.balances))
}

func (c *Cache) AverageBalance3() float64 {
	c.mu.Lock()
	m := make(map[string]float64, len(c.balances))
	for k, v := range c.balances {
		m[k] = v
	}
	c.mu.Unlock()

	sum := 0.
	for _, balance := range m {
		sum += balance
	}
	return sum / float64(len(c.balances))
}
