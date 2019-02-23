package main

import (
	"log"
	"net/http"

	"github.com/pizza-workshop/server"
	"github.com/pizza-workshop/store"
)

func main() {
	server := &server.Server{PizzaStore: store.NewPizzaStore()}
	if err := http.ListenAndServe(":5000", server); err != nil {
		log.Fatalf("Error on listening on Port 5000 %v", err)
	}
}
