package main

import (
	"encoding/json"
	"io"
	"net/http"
	"strings"
	"time"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("X-API-VERSION", "1.0")
	b, _ := io.ReadAll(r.Body)
	_, _ = w.Write(append([]byte("hello "), b...))
	w.WriteHeader(http.StatusCreated)
}

func (c DurationClient) GetDuration(url string, lat1, lng1, lat2, lng2 float64) (time.Duration, error) {
	resp, err := c.client.Post(
		url, "application/json",
		buildRequestBody(lat1, lng1, lat2, lng2),
	)
	if err != nil {
		return 0, err
	}

	return parseResponseBody(resp.Body)
}

type request struct {
	Duration int
}

func buildRequestBody(lat1, lng1, lat2, lng2 float64) io.Reader {
	return strings.NewReader("")
}

type DurationClient struct {
	client *http.Client
}

func NewDurationClient() DurationClient {
	return DurationClient{
		client: http.DefaultClient,
	}
}

func parseResponseBody(r io.ReadCloser) (time.Duration, error) {
	b, err := io.ReadAll(r)
	if err != nil {
		return 0, err
	}
	defer func() {
		_ = r.Close()
	}()

	var req request
	err = json.Unmarshal(b, &req)
	if err != nil {
		return 0, err
	}
	return time.Duration(req.Duration) * time.Second, nil
}
