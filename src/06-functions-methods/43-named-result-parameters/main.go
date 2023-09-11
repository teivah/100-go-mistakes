package main

import "io"

func f(a int) (b int) {
	b = a
	return
}

type locator interface {
	getCoordinates(address string) (float32, float32, error)
	// getCoordinates(address string) (lat, lng float32, err error)
}

type loc struct{}

func (l loc) getCoordinates(address string) (lat, lng float32, err error) {
	return 0, 0, nil
}

func ReadFull(r io.Reader, buf []byte) (n int, err error) {
	for len(buf) > 0 && err == nil {
		var nr int
		nr, err = r.Read(buf)
		n += nr
		buf = buf[nr:]
	}
	return
}
