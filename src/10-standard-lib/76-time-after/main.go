package main

import (
	"context"
	"log"
	"time"
)

func consumer1(ch <-chan Event) {
	for {
		select {
		case event := <-ch:
			handle(event)
		case <-time.After(time.Hour):
			log.Println("warning: no messages received")
		}
	}
}

func consumer2(ch <-chan Event) {
	for {
		ctx, cancel := context.WithTimeout(context.Background(), time.Hour)
		select {
		case event := <-ch:
			cancel()
			handle(event)
		case <-ctx.Done():
			log.Println("warning: no messages received")
		}
	}
}

func consumer3(ch <-chan Event) {
	timerDuration := 1 * time.Hour
	timer := time.NewTimer(timerDuration)

	for {
		timer.Reset(timerDuration)
		select {
		case event := <-ch:
			handle(event)
		case <-timer.C:
			log.Println("warning: no messages received")
		}
	}
}

type Event struct{}

func handle(Event) {
}
