package server

import (
	"net/http"
	"net/url"
)

// PizzaMaker is generic interface type
type PizzaMaker interface {
	GetStatus(name string) string
	MakePizza(pizzatype, contactType, contact, name string) bool
}

// Server is a HTTP server
type Server struct {
	PizzaStore PizzaMaker
}

// ServerHTTP fulfills the Handler interface
func (p *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path
	params := r.URL.Query()
	switch r.Method {
	case http.MethodPost:
		if path != "/buy_pizza" {
			statusNotFound(w)
			return
		}
		// TODO: create struct type to hold query information
		pizzaType, contactType, contact, name := getParams(params)
		p.makePizza(w, pizzaType, contactType, contact, name)
	case http.MethodGet:
		if path != "/get_status" {
			statusNotFound(w)
			return
		}
		name := params.Get("name")
		p.getStatus(w, name)
	}
}

// makePizaa is a method on Server struct
func (p *Server) makePizza(w http.ResponseWriter, pizzaType, contactType, contact, name string) {
	go p.PizzaStore.MakePizza(pizzaType, contactType, contact, name)
	w.WriteHeader(http.StatusAccepted)
}

// getStatus is a method on Server struct
func (p *Server) getStatus(w http.ResponseWriter, name string) {
	status := p.PizzaStore.GetStatus(name)
	if status == "" {
		statusNotFound(w)
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
}
