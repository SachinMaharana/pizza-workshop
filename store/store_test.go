package store

import (
	"testing"
)

func TestPizzaStoreStatus(t *testing.T) {
	s := NewPizzaStore()
	stubOrder := &order{name: "sachin", pizzaType: "veg", contact: "email", contactType: "email", status: "working"}
	s.orders["sachin"] = stubOrder
	t.Run("Get Status returns correct status", func(t *testing.T) {

		got := s.GetStatus("sachin")
		want := "working"

		if got != want {
			t.Errorf("did not get correct status, got %s, want %s", got, want)
		}
		doughPrep(stubOrder, 1)

		got = s.GetStatus("sachin")
		want = "Preparing Dough"

		if got != want {
			t.Errorf("did not get correct status, got %s, want %s", got, want)
		}
		ovenBake(stubOrder, 2)

		got = s.GetStatus("sachin")
		want = "OvenBakeing"

		if got != want {
			t.Errorf("did not get correct status, got %s, want %s", got, want)
		}
		toppingArt(stubOrder, 1)

		got = s.GetStatus("sachin")
		want = "Adding Toppings"

		if got != want {
			t.Errorf("did not get correct status, got %s, want %s", got, want)
		}
	})

}
func TestPizzaStoreMake(t *testing.T) {
	s := NewPizzaStore()
	stubOrder := &order{name: "sachin", pizzaType: "veg", contactType: "email", contact: "email", status: "working"}
	s.orders["sachin"] = stubOrder
	t.Run("Get Status returns correct status", func(t *testing.T) {

		got := s.MakePizza("veg", "email", "sachin.nicky@gmail.com", "sachin")
		want := true

		if got != want {
			t.Errorf("did not get correct status, got %v, want %v", got, want)
		}

	})

}
