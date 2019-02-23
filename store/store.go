package store

import (
	"math/rand"
	"time"
)

const (
	doughPrepTime  = 15
	ovenBakeTime   = 25
	toppingArtTime = 14
)

// NewPizzaStore creates a new Pizza store
func NewPizzaStore() *Store {
	return &Store{map[string]*order{}}
}

// Store is a generic pizza store that implenets pizza maker interface
type Store struct {
	orders map[string]*order
}

// order contains details of each order
type order struct {
	name, contactType, contact, pizzaType, status string
}

// TODO: change it to pool of workers
//return done channel

// MakePizza is a method on Store struct
func (s *Store) MakePizza(pizzaType, contactType, contact, name string) bool {
	c := &order{name: name, contactType: contactType, contact: contact, pizzaType: pizzaType, status: "starting"}
	s.orders[name] = c
	doughPrep(c, doughPrepTime)
	ovenBake(c, ovenBakeTime)
	toppingArt(c, toppingArtTime)
	c.status = "Order completed"
	return true

}

// GetStatus is a method on Store struct
func (s *Store) GetStatus(name string) string {
	var status string
	if _, ok := s.orders[name]; !ok {
		return status
	}
	status = s.orders[name].status
	return status
}

func doughPrep(c *order, t int) {
	c.status = "Preparing Dough"
	sleepDuration(t)

}

func ovenBake(c *order, t int) {
	c.status = "OvenBakeing"
	sleepDuration(t)
}

func toppingArt(c *order, t int) {
	c.status = "Adding Toppings"
	sleepDuration(t)
}

func sleepDuration(t int) {
	rand.Seed(time.Now().UnixNano())
	time.Sleep(time.Duration(rand.Intn(t)) * time.Second)
}
