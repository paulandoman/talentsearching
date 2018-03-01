package main

import (
	"testing"
)

func TestCheckout(t *testing.T) {
	classicItem := Item{id: classic, price: 269.99}
	standoutItem := Item{id: standout, price: 322.99}
	premiumItem := Item{id: premium, price: 394.99}

	// Default Test
	defaultCheckout := Checkout{}
	defaultCheckout.Add(classicItem)
	defaultCheckout.Add(standoutItem)
	defaultCheckout.Add(premiumItem)
	total := defaultCheckout.Total()
	if total != 987.97 {
		t.Errorf("Default ad total was incorrect, got: %v, expected: 987.97", total)
	}
	// Unilever Test

	// Apple Test

	// Nike Test

}
