package main

import (
	"fmt"
	"sync"
	"time"
)

func listing1() {
	type Donation struct {
		mu      sync.RWMutex
		balance int
	}
	donation := &Donation{}

	// Listener goroutines
	f := func(goal int) {
		donation.mu.RLock()
		for donation.balance < goal {
			donation.mu.RUnlock()
			donation.mu.RLock()
		}
		fmt.Printf("$%d goal reached\n", donation.balance)
		donation.mu.RUnlock()
	}
	go f(10)
	go f(15)

	// Updater goroutine
	go func() {
		for {
			time.Sleep(time.Second)
			donation.mu.Lock()
			donation.balance++
			donation.mu.Unlock()
		}
	}()
}

func listing2() {
	type Donation struct {
		balance int
		ch      chan int
	}

	donation := &Donation{ch: make(chan int)}

	// Listener goroutines
	f := func(goal int) {
		for balance := range donation.ch {
			if balance >= goal {
				fmt.Printf("$%d goal reached\n", balance)
				return
			}
		}
	}
	go f(10)
	go f(15)

	// Updater goroutine
	for {
		time.Sleep(time.Second)
		donation.balance++
		donation.ch <- donation.balance
	}
}

func listing3() {
	type Donation struct {
		cond    *sync.Cond
		balance int
	}

	donation := &Donation{
		cond: sync.NewCond(&sync.Mutex{}),
	}

	// Listener goroutines
	f := func(goal int) {
		donation.cond.L.Lock()
		for donation.balance < goal {
			donation.cond.Wait()
		}
		fmt.Printf("%d$ goal reached\n", donation.balance)
		donation.cond.L.Unlock()
	}
	go f(10)
	go f(15)

	// Updater goroutine
	for {
		time.Sleep(time.Second)
		donation.cond.L.Lock()
		donation.balance++
		donation.cond.L.Unlock()
		donation.cond.Broadcast()
	}
}

func listing4() {
	type Donation struct {
		cond     *sync.Cond
		jobDone  chan struct{}
		waitDone chan struct{}
		balance  int
	}

	donation := &Donation{
		cond:     sync.NewCond(&sync.Mutex{}),
		jobDone:  make(chan struct{}),
		waitDone: make(chan struct{}),
	}

	// Listener goroutines
	f := func(goal int) {
		for donation.balance < goal {
			donation.cond.Wait()
			if donation.balance < goal {
				donation.waitDone <- struct{}{}
			}
		}
		fmt.Printf("%d$ goal reached\n", donation.balance)
		donation.jobDone <- struct{}{}
		donation.cond.L.Unlock()
	}

	donation.cond.L.Lock()
	go f(10)
	donation.cond.L.Lock()
	go f(15)

	jobNumber := 2
	// Updater goroutine
	for jobNumber != 0 {
		donation.cond.L.Lock()
		// after get the lock, means all living jobs have been released the lock and in waiting status.
		waitNumber := jobNumber
		donation.balance++
		donation.cond.L.Unlock()
		donation.cond.Broadcast()
		for waitNumber != 0 {
			// no matter job done or wait done, means a job waked and get the lock.
			select {
			case <-donation.jobDone:
				jobNumber -= 1
			case <-donation.waitDone:
			}
			waitNumber -= 1
		}
	}
}
