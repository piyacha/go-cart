package cart

import "testing"

func TestCartEmpty(t *testing.T) {
	cart := Cart{}
	cart.AddToCart(Item {name: "Burger", price: 123.12})
	cart.AddToCart(Item {name: "Pizza", price: 100.12})

	if cart.TotalPrice() != 0 {
		t.Errorf("expected 0")
	}
}
