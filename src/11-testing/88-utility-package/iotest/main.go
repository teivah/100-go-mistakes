package main

import (
	"io"
)

type LowerCaseReader struct {
	reader io.Reader
}

func (l LowerCaseReader) Read(p []byte) (int, error) {
	return 0, nil
}

func foo1(r io.Reader) error {
	b, err := io.ReadAll(r)
	if err != nil {
		return err
	}

	// ...
	_ = b
	return nil
}

func foo2(r io.Reader) error {
	b, err := readAll(r, 3)
	if err != nil {
		return err
	}

	// ...
	_ = b
	return nil
}

func readAll(r io.Reader, retries int) ([]byte, error) {
	b := make([]byte, 0, 512)
	for {
		if len(b) == cap(b) {
			b = append(b, 0)[:len(b)]
		}
		n, err := r.Read(b[len(b):cap(b)])
		b = b[:len(b)+n]
		if err != nil {
			if err == io.EOF {
				return b, nil
			}
			retries--
			if retries < 0 {
				return b, err
			}
		}
	}
}
