package main

import (
	"errors"
	"fmt"
	"net/http"
)

type transientError struct {
	err error
}

func (t transientError) Error() string {
	return fmt.Sprintf("transient error: %v", t.err)
}

func handler(w http.ResponseWriter, r *http.Request) {
	transactionID := r.URL.Query().Get("transaction")

	amount, err := getTransactionAmount1(transactionID)
	if err != nil {
		switch err := err.(type) {
		case transientError:
			http.Error(w, err.Error(), http.StatusServiceUnavailable)
		default:
			http.Error(w, err.Error(), http.StatusBadRequest)
		}
		return
	}

	// Write response
	_ = amount
}

func getTransactionAmount1(transactionID string) (float32, error) {
	if len(transactionID) != 5 {
		return 0, fmt.Errorf("id is invalid: %s", transactionID)
	}

	amount, err := getTransactionAmountFromDB1(transactionID)
	if err != nil {
		return 0, transientError{err: err}
	}
	return amount, nil
}

func getTransactionAmountFromDB1(id string) (float32, error) {
	return 0, nil
}

func handler2(w http.ResponseWriter, r *http.Request) {
	transactionID := r.URL.Query().Get("transaction")

	amount, err := getTransactionAmount2(transactionID)
	if err != nil {
		if errors.As(err, &transientError{}) {
			http.Error(w, err.Error(), http.StatusServiceUnavailable)
		} else {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}
		return
	}

	// Write response
	_ = amount
}

func getTransactionAmount2(transactionID string) (float32, error) {
	// Check transaction ID validity

	amount, err := getTransactionAmountFromDB2(transactionID)
	if err != nil {
		return 0, fmt.Errorf("failed to get transaction %s: %w",
			transactionID, err)
	}
	return amount, nil
}

func getTransactionAmountFromDB2(transactionID string) (float32, error) {
	// ...
	var err error
	if err != nil {
		return 0, transientError{err: err}
	}
	// ...
	return 0, nil
}
