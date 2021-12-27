package main

import (
	"encoding/json"
	"fmt"
	"time"
)

func main() {
	if err := listing1(); err != nil {
		panic(err)
	}
	if err := listing2(); err != nil {
		panic(err)
	}
}

type Event1 struct {
	ID int
	time.Time
}

func listing1() error {
	event := Event1{
		ID:   1234,
		Time: time.Now(),
	}

	b, err := json.Marshal(event)
	if err != nil {
		return err
	}

	fmt.Println(string(b))
	return nil
}

type Event2 struct {
	ID   int
	Time time.Time
}

func listing2() error {
	event := Event2{
		ID:   1234,
		Time: time.Now(),
	}

	b, err := json.Marshal(event)
	if err != nil {
		return err
	}

	fmt.Println(string(b))
	return nil
}
