package main

import "testing"

func TestCheckout(t *testing.T) {
	c := Checkout{}
	i := Item{id: classic, price: 269.99}
	c.Add(i)
	if c.classicAds != 1 {
		t.Errorf("Total classic ads was incorrect, got: %v, expected: 1", c.classicAds)
	}
}
