package store

type Customer struct {
	// Some fields
}

type Contract struct {
	// Some fields
}

type Store struct{}

func (s *Store) Get(id string) (any, error) {
	return nil, nil
}

func (s *Store) Set(id string, v any) error {
	return nil
}
