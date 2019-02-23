package server

import (
	"net/http"
	"net/url"
)

// PizzaMaker is an interface type that has methods a pizza store must implement
type PizzaMaker interface {
	GetStatus(name string) string

	MakePizza(pizzaType, contactType, contact, name string) bool
}

// Server is http server that must implemnet handler interface
type Server struct {
	PizzaStore PizzaMaker
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path
	params := r.URL.Query()

	switch r.Method {
	case http.MethodPost:
		if path != "/buy_pizza" {
			statusNotFound(w)
			return
		}
		pizzaType, contactType, contact, name := getParams(params)
		s.makePizza(w, pizzaType, contactType, contact, name)
	case http.MethodGet:
		if path != "/get_status" {
			statusNotFound(w)
			return
		}
		name := params.Get("name")
		s.getStatus(w, name)
	}
}

func (s *Server) makePizza(w http.ResponseWriter, pizzaType, contactType, contact, name string) {
	go s.PizzaStore.MakePizza(pizzaType, contactType, contact, name)
	w.WriteHeader(http.StatusAccepted)
}

func (s *Server) getStatus(w http.ResponseWriter, name string) {
	status := s.PizzaStore.GetStatus(name)
	if status == "" {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Order Not Found"))
		return
	}
	w.WriteHeader(http.StatusAccepted)
	w.Write([]byte(status))
}

func getParams(p url.Values) (string, string, string, string) {
	return p.Get("pizzaType"), p.Get("contactType"), p.Get("contact"), p.Get("name")
}

func statusNotFound(w http.ResponseWriter) {
	w.WriteHeader(http.StatusNotFound)
	w.Write([]byte("Not Found"))
	return
}
