package main

import (
	"context"
	"net/http"
)

type key string

const isValidHostKey key = "isValidHost"

func checkValid(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		validHost := r.Host == "acme"
		ctx := context.WithValue(r.Context(), isValidHostKey, validHost)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func handler(ctx context.Context, ch chan Message) error {
	for {
		select {
		case msg := <-ch:
			// Do something with msg
			_ = msg
		case <-ctx.Done():
			return ctx.Err()
		}
	}
}

type Message struct{}
