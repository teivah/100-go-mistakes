package main

import "testing"

func TestPutCustomer1(t *testing.T) {
	customer, err := createCustomer1()
	if err != nil {
		t.Fatal(err)
	}
	err = NewStore().PutCustomer(customer)
	// Check err
	_ = err
}

func createCustomer1() (Customer, error) {
	customer, err := customerFactory("foo")
	if err != nil {
		return Customer{}, nil
	}
	return customer, nil
}

func TestPutCustomer(t *testing.T) {
	customer := createCustomer2(t)
	err := NewStore().PutCustomer(customer)
	// Check err
	_ = err
}

func createCustomer2(t *testing.T) Customer {
	customer, err := customerFactory("foo")
	if err != nil {
		t.Fatal(err)
	}
	return customer
}
