package server

import (
	"net/http"
	"net/url"
)

type PizzaMaker interface {
	GetStatus(name string) string

	MakePizza(pizzaType, contactType, contact, name string) bool
}

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
		}
		pizzaType, contactType, contact, name := getParams(params)
		s.makePizza(w, pizzaType, contactType, contact, name)
	case http.MethodGet:
		if path == "/" {
			w.WriteHeader(http.StatusFound)
			w.Write([]byte("Server is up."))
		}

	}
}

func (s *Server) makePizza(w http.ResponseWriter, pizzaType, contactType, contact, name string) {
	go s.PizzaStore.MakePizza(pizzaType, contactType, contact, name)
	w.WriteHeader(http.StatusAccepted)
}

func getParams(p url.Values) (string, string, string, string) {
	return p.Get("pizzaType"), p.Get("contactType"), p.Get("contact"), p.Get("name")
}

func statusNotFound(w http.ResponseWriter) {
	w.WriteHeader(http.StatusNotFound)
	w.Write([]byte("Not Found"))
	return
}
