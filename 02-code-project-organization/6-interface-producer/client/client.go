package client

import "github.com/teivah/100-go-mistakes/02-code-project-organization/6-interface-producer/store"

type customersGetter interface {
	GetAllCustomers() ([]store.Customer, error)
}
