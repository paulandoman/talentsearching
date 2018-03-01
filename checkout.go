package main

import (
	"fmt"
)

func main() {

}

// JobAd does something
type JobAd int

const (
	classic JobAd = iota
	standout
	premium
)

// PricingRules does something
type PricingRules int

const (
	unilever PricingRules = iota
	apple
	nike
	ford
)

// Item does something
type Item struct {
	id    JobAd
	price float32
}

// Checkout does something
type Checkout struct {
	pricingRules PricingRules
	classicAds   float32
	standoutAds  float32
	premiumAds   float32
}

// Add does something
func (check Checkout) Add(i Item) {
	switch i.id {
	case classic:
		fmt.Println("classic added")
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
func (check Checkout) Total() float32 {
	fmt.Println(check.classicAds)
	fmt.Println(check.standoutAds)
	fmt.Println(check.premiumAds)

	classicAdCost := check.classicAds * 269.99
	standoutAdCost := check.standoutAds * 322.99
	premiumAdCost := check.premiumAds * 394.99
	return classicAdCost + standoutAdCost + premiumAdCost
}
