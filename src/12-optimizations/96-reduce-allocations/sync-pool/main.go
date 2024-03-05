package main

import (
	"io"
	"sync"
)

var pool = sync.Pool{
	New: func() any {
		b := make([]byte, 1024)
		return &b
	},
}

func write(w io.Writer) {
	bPtr := pool.Get().(*[]byte)
	defer func() {
		*bPtr = (*bPtr)[:0]
		pool.Put(bPtr)
	}()

	b := *bPtr
	getResponse(b)
	_, _ = w.Write(b)
}

func getResponse([]byte) {
}
