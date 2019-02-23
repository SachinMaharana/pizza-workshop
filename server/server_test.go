package server

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

// order contains details of each order
type stuborder struct {
	name, contactType, contact, pizzaType, status string
}

type StubPizzaStore struct {
	orders map[string]stuborder
}

func (s *StubPizzaStore) GetStatus(name string) string {
	status := s.orders[name].status
	return status
}
func (s *StubPizzaStore) MakePizza(pizzaType, contactType, contact, name string) bool {
	return true
}

func TestGETStatus(t *testing.T) {
	pizzaStore := StubPizzaStore{
		map[string]stuborder{},
	}

	server := &Server{&pizzaStore}

	t.Run("Get Status returns correct status", func(t *testing.T) {
		name := "sachin"
		request, _ := http.NewRequest(http.MethodGet, fmt.Sprintf("/get_status/?name=%s", name), nil)
		response := httptest.NewRecorder()
		server.ServeHTTP(response, request)

		got := response.Code
		want := http.StatusNotFound

		if got != want {
			t.Errorf("did not get correct status, got %d, want %d", got, want)
		}
	})

}

func TestMakePizza(t *testing.T) {
	pizzaStore := StubPizzaStore{
		map[string]stuborder{},
	}
	server := &Server{&pizzaStore}

	t.Run("Make pizza makes succesful post request", func(t *testing.T) {
		pizzaType, contactType, contact, name := "veg", "email", "sachin.nicky@gmail.com", "sachin"

		request, _ := http.NewRequest(http.MethodPost, fmt.Sprintf("/buy_pizza?pizzaType=%s&contact=%s&name=%s&contactType=%s", pizzaType, contact, name, contactType), nil)
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)
		got := response.Code
		want := http.StatusAccepted

		if got != want {
			t.Errorf("did not get correct status, got %d, want %d", got, want)
		}

	})
}

// func TestHow(t *testing.T) {
// 	pizzaStore := StubPizzaStore{
// 		map[string]*stuborder{},
// 	}
// 	server := &Server{&pizzaStore}

// localhost:5000/buy_pizza?pizzaType=veg&contact=lol&name=ot
