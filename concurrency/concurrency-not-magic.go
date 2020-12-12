package concurrency

import (
	"bufio"
	"io"
	"runtime"
	"strconv"
	"strings"
	"sync"
	"time"
)

type customer struct {
	id        string
	firstName string
	lastName  string
	ts        time.Time
}

func parseFileSequential(reader *bufio.Reader) ([]customer, error) {
	customers := make([]customer, 0)
	for {
		// Read a line
		line, _, err := reader.ReadLine()
		if err != nil {
			switch err {
			default:
				return nil, err
			case io.EOF:
				// If io.EOF, we reached the end of the input
				return customers, nil
			}
		}
		// Call parseLineSequential and add another customer
		customers = append(customers, parseLineSequential(string(line)))
	}
	return customers, nil
}

func parseLineSequential(line string) customer {
	tokens := strings.Split(line, ",")
	ts, _ := strconv.ParseInt(tokens[3], 10, 64)
	return customer{
		id:        tokens[0],
		firstName: tokens[1],
		lastName:  tokens[2],
		ts:        time.Unix(ts, 0),
	}
}

func parseFileConcurrentV1(reader *bufio.Reader) ([]customer, error) {
	results := make(chan customer, 1024)
	customers := make([]customer, 0)
	wg := sync.WaitGroup{}

loop:
	for {
		line, _, err := reader.ReadLine()
		if err != nil {
			switch err {
			default:
				return nil, err
			case io.EOF:
				break loop
			}
		}

		// Add 1 to the wait group
		wg.Add(1)
		// Spin up a new goroutine
		go parseLineConcurrentV1(&wg, string(line), results)
	}

	go func() {
		// Wait for all the goroutines to complete before closing the channel
		wg.Wait()
		close(results)
	}()

	// Gather the results
	for customer := range results {
		customers = append(customers, customer)
	}

	return customers, nil
}

func parseLineConcurrentV1(wg *sync.WaitGroup, line string, output chan<- customer) {
	tokens := strings.Split(line, ",")
	ts, _ := strconv.ParseInt(tokens[3], 10, 64)
	output <- customer{
		id:        tokens[0],
		firstName: tokens[1],
		lastName:  tokens[2],
		ts:        time.Unix(ts, 0),
	}
	wg.Done()
}

func parseFileConcurrentV2(reader *bufio.Reader) ([]customer, error) {
	inputs := make(chan string, 1024)
	results := make(chan customer, 1024)
	customers := make([]customer, 0)

	// Spin up multiple worker goroutines
	workerWg := sync.WaitGroup{}
	for i := 0; i < runtime.NumCPU(); i++ {
		workerWg.Add(1)
		go parseLineConcurrentV2(&workerWg, inputs, results)
	}

	// Gather
	gatherWg := sync.WaitGroup{}
	gatherWg.Add(1)
	go func() {
		for customer := range results {
			customers = append(customers, customer)
		}
		gatherWg.Done()
	}()

	// When workers are complete, we close the channel.
	go func() {
		workerWg.Wait()
		close(results)
	}()

loop:
	for {
		line, _, err := reader.ReadLine()
		if err != nil {
			switch err {
			default:
				return nil, err
			case io.EOF:
				break loop
			}
		}
		// Scatter
		inputs <- string(line)
	}
	close(inputs)

	// Wait for the gather goroutine to complete.
	gatherWg.Wait()
	return customers, nil
}

func parseLineConcurrentV2(wg *sync.WaitGroup, input <-chan string, output chan<- customer) {
	for line := range input {
		tokens := strings.Split(line, ",")
		ts, _ := strconv.ParseInt(tokens[3], 10, 64)
		output <- customer{
			id:        tokens[0],
			firstName: tokens[1],
			lastName:  tokens[2],
			ts:        time.Unix(ts, 0),
		}
	}
	wg.Done()
}
