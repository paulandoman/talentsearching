package main

import (
	"testing"
)

func TestAdd(t *testing.T) {
	classicItem := Item{id: classic}
	standoutItem := Item{id: standout}
	premiumItem := Item{id: premium}

	defaultCheckout := Checkout{
		pricingRules: CustomerPriceRules["default"],
	}
	defaultCheckout.Add(classicItem)
	defaultCheckout.Add(classicItem)
	defaultCheckout.Add(standoutItem)
	defaultCheckout.Add(standoutItem)
	defaultCheckout.Add(standoutItem)
	defaultCheckout.Add(premiumItem)

	total := defaultCheckout.classTotal
	if total != 2 {
		t.Errorf("Number of classic ads added was incorrect, got: %v, expected: 2", total)
	}

	total = defaultCheckout.standTotal
	if total != 3 {
		t.Errorf("Number of standout ads added was incorrect, got: %v, expected: 3", total)
	}

	total = defaultCheckout.premTotal
	if total != 1 {
		t.Errorf("Number of premium ads added was incorrect, got: %v, expected: 1", total)
	}

}

func TestCheckout(t *testing.T) {
	classicItem := Item{id: classic}
	standoutItem := Item{id: standout}
	premiumItem := Item{id: premium}

	// Default Test
	defaultCheckout := Checkout{
		pricingRules: CustomerPriceRules["default"],
	}
	defaultCheckout.Add(classicItem)
	defaultCheckout.Add(standoutItem)
	defaultCheckout.Add(premiumItem)
	total := defaultCheckout.Total()
	if total != 987.97 {
		t.Errorf("Default ad total was incorrect, got: %v, expected: 987.97", total)
	}
	// Unilever Test
	unileverCheckout := Checkout{
		pricingRules: CustomerPriceRules["unilever"],
	}
	unileverCheckout.Add(classicItem)
	unileverCheckout.Add(classicItem)
	unileverCheckout.Add(classicItem)
	unileverCheckout.Add(premiumItem)
	total = unileverCheckout.Total()
	if total != 934.97 {
		t.Errorf("Unilever ad total was incorrect, got: %v, expected: 934.97", total)
	}

	// Apple Test

	// Nike Test

}
