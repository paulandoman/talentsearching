package main

import (
	"testing"
)

var classicItem = Item{adType: classic}
var standoutItem = Item{adType: standout}
var premiumItem = Item{adType: premium}

func TestAddAdsOneAtATime(t *testing.T) {
	// Add 2 classic, 3 standout and 1 premium ad
	checkout := Checkout{
		pricingRules: GetRules("default"),
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
}

func TestAddingMultipleAds(t *testing.T) {
	// Adding multiple ads at once
	checkout := Checkout{
		pricingRules: GetRules("unilever"),
	}
	checkout.Add(classicItem, classicItem, premiumItem, standoutItem)

	total := checkout.classTotal
	if total != 2 {
		t.Errorf("Multiple classic ads added in total was incorrect, got: %v, expected: 2", total)
	}
}

func TestAddAndRemoveAdsOneAtATime(t *testing.T) {
	// Add 2 classic ads remove 1 classic ad
	checkout := Checkout{
		pricingRules: GetRules("default"),
	}
	checkout.Add(classicItem)
	checkout.Add(classicItem)
	checkout.Remove(classicItem)

	total := checkout.classTotal
	if total != 1 {
		t.Errorf("Number of classic ads added and removed was incorrect, got: %v, expected: 1", total)
	}
}

func TestAddAndRemoveSameAmountOfAds(t *testing.T) {
	// Adding and subtracting the same number of standout ads
	checkout := Checkout{
		pricingRules: GetRules("default"),
	}
	checkout.Add(standoutItem)
	checkout.Remove(standoutItem)

	total := checkout.standTotal
	if total != 0 {
		t.Errorf("Number of standout ads added and removed was incorrect, got: %v, expected: 0", total)
	}
}

func TestTryRemoveAdWhenTotalIsZero(t *testing.T) {
	// Try and remove an ad when there are none left to remove
	checkout := Checkout{
		pricingRules: GetRules("default"),
	}
	checkout.Remove(premiumItem)

	total := checkout.classTotal
	if total != 0 {
		t.Errorf("When removing ad that doesn't exist from the checkout total ads of that type should stay zero, got: %v, expected: 0", total)
	}
}

func TestAddingAndRemovingMultipleAds(t *testing.T) {
	// Add multiple classic ads and then remove multiple classic ads
	checkout := Checkout{
		pricingRules: GetRules("default"),
	}
	checkout.Add(classicItem, classicItem, classicItem)
	checkout.Remove(classicItem, classicItem)

	total := checkout.classTotal
	if total != 1 {
		t.Errorf("Number of classic ads added and removed was incorrect, got: %v, expected: 1", total)
	}
}

func TestTryRemovingMultipleAdsMoreThanPresent(t *testing.T) {
	// Try and remove multiple classic ads more than is present
	checkout := Checkout{
		pricingRules: GetRules("default"),
		classTotal:   1.0,
	}
	checkout.Remove(classicItem, classicItem)

	total := checkout.classTotal
	if total != 0 {
		t.Errorf("Total no of classic ads should stay zero when trying to remove more than present, got: %v, expected: 0", total)
	}
}

func TestTotal(t *testing.T) {
	// Default Test - 1 classic, 1 standout, 1 premium ad
	checkout := Checkout{
		pricingRules: GetRules("default"),
	}
	checkout.Add(classicItem, standoutItem, premiumItem)

	total := checkout.Total()
	if total != 987.97 {
		t.Errorf("Default ad total was incorrect, got: %v, expected: 987.97", total)
	}
}

func TestUnileverTotal(t *testing.T) {
	// Unilever Test - 3 for 2 on classic ads
	checkout := Checkout{
		pricingRules: GetRules("unilever"),
	}
	checkout.Add(classicItem, classicItem, classicItem, premiumItem)

	total := checkout.Total()
	if total != 934.97 {
		t.Errorf("Unilever ad total was incorrect, got: %v, expected: 934.97", total)
	}
}

func TestAppleTotal(t *testing.T) {
	// Apple Test - discount on standout ads
	checkout := Checkout{
		pricingRules: GetRules("apple"),
	}
	checkout.Add(standoutItem, standoutItem, standoutItem, premiumItem)

	total := checkout.Total()
	if total != 1294.96 {
		t.Errorf("Apple ad total was incorrect, got: %v, expected: 1294.96", total)
	}
}

func TestNikeTotal(t *testing.T) {
	// Nike Test - bulk discount on premium ads
	checkout := Checkout{
		pricingRules: GetRules("nike"),
	}
	checkout.Add(premiumItem, premiumItem, premiumItem, premiumItem)

	total := checkout.Total()
	if total != 1519.96 {
		t.Errorf("Nike ad total was incorrect, got: %v, expected: 1519.96", total)
	}
}

func TestFordTotal(t *testing.T) {
	// Ford Test - 5 for 4 on classic ads, discount on standout, bulk discounts on premium ads
	checkout := Checkout{
		pricingRules: GetRules("ford"),
	}
	checkout.Add(classicItem, classicItem, classicItem, classicItem, classicItem, classicItem)
	checkout.Add(standoutItem)
	checkout.Add(premiumItem, premiumItem, premiumItem)

	total := checkout.Total()
	if total != 2829.91 {
		t.Errorf("Ford ad total was incorrect, got: %v, expected: 2829.91", total)
	}
}
