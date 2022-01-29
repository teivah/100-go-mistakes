package main

import (
	"fmt"
	"io"
	"sync"
	"sync/atomic"
)

func main() {
	res1, _ := read1(&dummyReader{})
	fmt.Println(res1)

	res2, _ := read2(&dummyReader{})
	fmt.Println(res2)
}

func read1(r io.Reader) (int, error) {
	count := 0
	for {
		b := make([]byte, 1024)
		_, err := r.Read(b)
		if err != nil {
			if err == io.EOF {
				break
			}
			return 0, err
		}
		count += task(b)
	}
	return count, nil
}

func read2(r io.Reader) (int, error) {
	var count int64
	wg := sync.WaitGroup{}
	n := 10

	ch := make(chan []byte, n)
	wg.Add(n)
	for i := 0; i < n; i++ {
		go func() {
			defer wg.Done()
			for b := range ch {
				v := task(b)
				atomic.AddInt64(&count, int64(v))
			}
		}()
	}

	for {
		b := make([]byte, 1024)
		_, err := r.Read(b)
		if err != nil {
			if err == io.EOF {
				break
			}
			return 0, err
		}
		ch <- b
	}

	close(ch)
	wg.Wait()
	return int(count), nil
}

func task(b []byte) int {
	return len(b)
}

type dummyReader struct {
	i int
}

func (c *dummyReader) Read(p []byte) (n int, err error) {
	if c.i == 3 {
		return 0, io.EOF
	}
	copy(p, []byte{0, 1, 2})
	c.i++
	return 3, nil
}
