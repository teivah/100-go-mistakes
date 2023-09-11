package main

import (
	"net/http"
	"time"
)

func main() {
	s := &http.Server{
		Addr:              ":8080",
		ReadHeaderTimeout: 500 * time.Millisecond,
		ReadTimeout:       500 * time.Millisecond,
		Handler:           http.TimeoutHandler(handler{}, time.Second, "foo"),
	}
	_ = s
}

type handler struct{}

func (h handler) ServeHTTP(http.ResponseWriter, *http.Request) {}
