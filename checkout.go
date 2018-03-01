package main

import (
	"fmt"
)

func main() {

}

// JobAd represent different types of job ads
type JobAd int

const (
	classic JobAd = iota
	standout
	premium
)

// PricingRules represents different customer pricing rules
type PricingRules int

const (
	standard PricingRules = iota
	unilever
	apple
	nike
	ford
)

// Item represents a job ad
type Item struct {
	id    JobAd
	price float64
}

// Checkout represents an individual checkout
type Checkout struct {
	pricingRules PricingRules
	classicAds   float64
	standoutAds  float64
	premiumAds   float64
}

// Add ads to the checkout
func (check *Checkout) Add(i Item) {
	switch i.id {
	case classic:
		fmt.Println(check.pricingRules, "classic added")
		check.classicAds++
	case standout:
		fmt.Println("standout added")
		check.standoutAds++
	case premium:
		fmt.Println("premium added")
		check.premiumAds++
	default:
		fmt.Println("nothing added")
	}
}

// Total adds up the total cost of the ads based on the customer
func (check *Checkout) Total() float64 {
	classicAdCost := 0.0
	standoutAdCost := 0.0
	premiumAdCost := 0.0

	switch check.pricingRules {
	case unilever:
		classicAdCost = (check.classicAds - check.classicAds/3) * 269.99
		standoutAdCost = check.standoutAds * 322.99
		premiumAdCost = check.premiumAds * 394.99
	default:
		classicAdCost = check.classicAds * 269.99
		standoutAdCost = check.standoutAds * 322.99
		premiumAdCost = check.premiumAds * 394.99
	}
	return Truncate(classicAdCost + standoutAdCost + premiumAdCost)
}
