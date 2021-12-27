package main

type CustomerService struct {
	store Store
}

func (cs CustomerService) CreateNewCustomer(id string) error {
	customer := Customer{id: id}
	return cs.store.StoreCustomer(customer)
}

type Customer struct {
	id string
}

type Store struct{}

func (s Store) StoreCustomer(customer Customer) error {
	return nil
}
