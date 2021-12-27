package main

import (
	"errors"
	"fmt"
	"sync"
)

type Customer struct {
	mutex sync.RWMutex
	id    string
	age   int
}

func (c *Customer) UpdateAge1(age int) error {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	if age < 0 {
		return errors.New("age should be positive")
	}

	c.age = age
	return nil
}

func (c *Customer) UpdateAge2(age int) error {
	if age < 0 {
		return errors.New("age should be positive")
	}

	c.mutex.Lock()
	defer c.mutex.Unlock()

	c.age = age
	return nil
}

func (c *Customer) String() string {
	c.mutex.RLock()
	defer c.mutex.RUnlock()
	return fmt.Sprintf("id %s, age %d", c.id, c.age)
}
