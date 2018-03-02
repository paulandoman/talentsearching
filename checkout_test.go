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

func TestTotal(t *testing.T) {
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
	// Unilever Test - 3 for 2 on classic ads
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

	// Apple Test - discount on standout ads
	appleCheckout := Checkout{
		pricingRules: CustomerPriceRules["apple"],
	}
	appleCheckout.Add(standoutItem)
	appleCheckout.Add(standoutItem)
	appleCheckout.Add(standoutItem)
	appleCheckout.Add(premiumItem)
	total = appleCheckout.Total()
	if total != 1294.96 {
		t.Errorf("Apple ad total was incorrect, got: %v, expected: 1294.96", total)
	}

	// Nike Test - bulk discount on premium ads
	nikeCheckout := Checkout{
		pricingRules: CustomerPriceRules["nike"],
	}
	nikeCheckout.Add(premiumItem)
	nikeCheckout.Add(premiumItem)
	nikeCheckout.Add(premiumItem)
	nikeCheckout.Add(premiumItem)

	total = nikeCheckout.Total()
	if total != 1519.96 {
		t.Errorf("Nike ad total was incorrect, got: %v, expected: 1519.96", total)
	}

	// Ford Test - 5 for 4 on classic ads, discount on standout, bulk discounts on premium ads
	fordCheckout := Checkout{
		pricingRules: CustomerPriceRules["ford"],
	}
	fordCheckout.Add(classicItem)
	fordCheckout.Add(classicItem)
	fordCheckout.Add(classicItem)
	fordCheckout.Add(classicItem)
	fordCheckout.Add(classicItem)
	fordCheckout.Add(classicItem)
	fordCheckout.Add(standoutItem)
	fordCheckout.Add(premiumItem)
	fordCheckout.Add(premiumItem)
	fordCheckout.Add(premiumItem)

	total = fordCheckout.Total()
	if total != 2829.91 {
		t.Errorf("Ford ad total was incorrect, got: %v, expected: 2829.91", total)
	}

}
