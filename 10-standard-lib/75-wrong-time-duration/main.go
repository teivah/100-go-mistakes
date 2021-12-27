package main

import "time"

func listing1() {
	ticker := time.NewTicker(1000)
	for {
		select {
		case <-ticker.C:
			// Do something
		}
	}
}

func listing2() {
	ticker := time.NewTicker(time.Microsecond)
	for {
		select {
		case <-ticker.C:
			// Do something
		}
	}
}
