package store

// NewPizzaStore creates a new pizza store
func NewPizzaStore() *Store {
	return &Store{map[string]*order{}}
}

// Store implements pizza maker interface
type Store struct {
	orders map[string]*order
}

type order struct {
	name, contact, contactType, pizzaType, status string
}

// MakePizza is a method on Store strcut
func (s *Store) MakePizza(pizzaType, contactType, contact, name string) bool {
	return true
}

// GetStatus is a method on Store strcut and returns status
func (s *Store) GetStatus(name string) string {
	var status string
	return status
}
