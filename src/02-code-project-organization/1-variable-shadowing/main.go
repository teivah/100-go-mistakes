package main

import (
	"log"
	"net/http"
)

func listing1() error {
	var client *http.Client
	if tracing {
		client, err := createClientWithTracing()
		if err != nil {
			return err
		}
		log.Println(client)
	} else {
		client, err := createDefaultClient()
		if err != nil {
			return err
		}
		log.Println(client)
	}

	_ = client
	return nil
}

func listing2() error {
	var client *http.Client
	if tracing {
		c, err := createClientWithTracing()
		if err != nil {
			return err
		}
		client = c
	} else {
		c, err := createDefaultClient()
		if err != nil {
			return err
		}
		client = c
	}

	_ = client
	return nil
}

func listing3() error {
	var client *http.Client
	var err error
	if tracing {
		client, err = createClientWithTracing()
		if err != nil {
			return err
		}
	} else {
		client, err = createDefaultClient()
		if err != nil {
			return err
		}
	}

	_ = client
	return nil
}

func listing4() error {
	var client *http.Client
	var err error
	if tracing {
		client, err = createClientWithTracing()
	} else {
		client, err = createDefaultClient()
	}
	if err != nil {
		return err
	}

	_ = client
	return nil
}

var tracing bool

func createClientWithTracing() (*http.Client, error) {
	return nil, nil
}

func createDefaultClient() (*http.Client, error) {
	return nil, nil
}
