package main

import (
	"testing"
)

func TestAdd(t *testing.T) {
	classicItem := Item{id: classic}
	standoutItem := Item{id: standout}
	premiumItem := Item{id: premium}

	// Add 2 classic, 3 standout and 1 premium ad
	checkout := Checkout{
		pricingRules: CustomerPriceRules["default"],
	}
	checkout.Add(classicItem)
	checkout.Add(classicItem)
	checkout.Add(standoutItem)
	checkout.Add(standoutItem)
	checkout.Add(standoutItem)
	checkout.Add(premiumItem)

	total := checkout.classTotal
	if total != 2 {
		t.Errorf("Number of classic ads added was incorrect, got: %v, expected: 2", total)
	}

	total = checkout.standTotal
	if total != 3 {
		t.Errorf("Number of standout ads added was incorrect, got: %v, expected: 3", total)
	}

	total = checkout.premTotal
	if total != 1 {
		t.Errorf("Number of premium ads added was incorrect, got: %v, expected: 1", total)
	}

	// Adding multiple ads at once
	unileverCheckout := Checkout{
		pricingRules: CustomerPriceRules["unilever"],
	}

	unileverCheckout.Add(classicItem, classicItem, premiumItem, standoutItem)
	total = unileverCheckout.classTotal
	if total != 2 {
		t.Errorf("Multiple classic ads added in total was incorrect, got: %v, expected: 2", total)
	}

}

func TestRemove(t *testing.T) {
	classicItem := Item{id: classic}

	// Add 2 classic ads remove 1 classic ad
	checkout := Checkout{
		pricingRules: CustomerPriceRules["default"],
	}
	checkout.Add(classicItem)
	checkout.Add(classicItem)
	checkout.Remove(classicItem)

	total := checkout.classTotal
	if total != 1 {
		t.Errorf("Number of classic ads added and removed was incorrect, got: %v, expected: 1", total)
	}

	// Added and subtracted the same number of classic ads
	checkout.Remove(classicItem)

	total = checkout.classTotal
	if total != 0 {
		t.Errorf("Number of classic ads added and removed was incorrect, got: %v, expected: 0", total)
	}

	// Try and remove a classic ad when there are none left to remove
	checkout.Remove(classicItem)

	total = checkout.classTotal
	if total != 0 {
		t.Errorf("When removing ad that doesn't exist from the checkout total ads of that type should stay zero, got: %v, expected: 0", total)
	}

	// Add multiple classic ads and then remove multiple classic ads
	anotherCheckout := Checkout{
		pricingRules: CustomerPriceRules["default"],
	}
	anotherCheckout.Add(classicItem, classicItem, classicItem)
	anotherCheckout.Remove(classicItem, classicItem)
	total = anotherCheckout.classTotal
	if total != 1 {
		t.Errorf("Number of classic ads added and removed was incorrect, got: %v, expected: 1", total)
	}

	// Try and remove multiple classic ads more than is present
	anotherCheckout.Remove(classicItem, classicItem)
	total = anotherCheckout.classTotal
	if total != 0 {
		t.Errorf("Total no of classic ads should stay zero when trying to remove more than present, got: %v, expected: 0", total)
	}

}

func TestTotal(t *testing.T) {
	classicItem := Item{id: classic}
	standoutItem := Item{id: standout}
	premiumItem := Item{id: premium}

	// Default Test - 1 classic, 1 standout, 1 premium ad
	checkout := Checkout{
		pricingRules: CustomerPriceRules["default"],
	}
	checkout.Add(classicItem, standoutItem, premiumItem)
	total := checkout.Total()
	if total != 987.97 {
		t.Errorf("Default ad total was incorrect, got: %v, expected: 987.97", total)
	}
	// Unilever Test - 3 for 2 on classic ads
	unileverCheckout := Checkout{
		pricingRules: CustomerPriceRules["unilever"],
	}
	unileverCheckout.Add(classicItem, classicItem, classicItem, premiumItem)
	total = unileverCheckout.Total()
	if total != 934.97 {
		t.Errorf("Unilever ad total was incorrect, got: %v, expected: 934.97", total)
	}

	// Apple Test - discount on standout ads
	appleCheckout := Checkout{
		pricingRules: CustomerPriceRules["apple"],
	}
	appleCheckout.Add(standoutItem, standoutItem, standoutItem, premiumItem)
	total = appleCheckout.Total()
	if total != 1294.96 {
		t.Errorf("Apple ad total was incorrect, got: %v, expected: 1294.96", total)
	}

	// Nike Test - bulk discount on premium ads
	nikeCheckout := Checkout{
		pricingRules: CustomerPriceRules["nike"],
	}
	nikeCheckout.Add(premiumItem, premiumItem, premiumItem, premiumItem)
	total = nikeCheckout.Total()
	if total != 1519.96 {
		t.Errorf("Nike ad total was incorrect, got: %v, expected: 1519.96", total)
	}

	// Ford Test - 5 for 4 on classic ads, discount on standout, bulk discounts on premium ads
	fordCheckout := Checkout{
		pricingRules: CustomerPriceRules["ford"],
	}
	fordCheckout.Add(classicItem, classicItem, classicItem, classicItem, classicItem, classicItem)
	fordCheckout.Add(standoutItem)
	fordCheckout.Add(premiumItem, premiumItem, premiumItem)

	total = fordCheckout.Total()
	if total != 2829.91 {
		t.Errorf("Ford ad total was incorrect, got: %v, expected: 2829.91", total)
	}

}
