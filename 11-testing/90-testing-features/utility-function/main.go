package main

func NewStore() Store {
	return Store{}
}

type Store struct{}

func (s Store) PutCustomer(Customer) error {
	return nil
}

type Customer struct {
	id string
}

func customerFactory(id string) (Customer, error) {
	if id == "" {
		return Customer{}, nil
	}
	return Customer{id: id}, nil
}
