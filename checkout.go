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

// Item represents a job ad
type Item struct {
	id JobAd
}

// Checkout represents an individual checkout
type Checkout struct {
	pricingRules PricingRules
	classTotal   float64
	standTotal   float64
	premTotal    float64
}

// PricingRules model for the different types of ad
type PricingRules struct {
	classic  Pricing
	standout Pricing
	premium  Pricing
}

// Pricing model for each particular ad
type Pricing struct {
	Price     float64
	XforY     float64 // Where Y = X-1
	BulkNo    int
	BulkPrice float64
}

// Add ads to the checkout
func (check *Checkout) Add(i Item) {
	switch i.id {
	case classic:
		fmt.Println(check.pricingRules, "classic added")
		check.classTotal++
	case standout:
		fmt.Println("standout added")
		check.standTotal++
	case premium:
		fmt.Println("premium added")
		check.premTotal++
	default:
		fmt.Println("nothing added")
	}
}

// Total adds up the total cost of the ads based on the customer
func (check *Checkout) Total() float64 {
	classicDiscount := ApplyBulkDiscount(check.classTotal, check.pricingRules.classic.XforY)
	classicAdCost := (check.classTotal - classicDiscount) * check.pricingRules.classic.Price

	standoutDiscount := ApplyBulkDiscount(check.standTotal, check.pricingRules.standout.XforY)
	standoutAdCost := (check.standTotal - standoutDiscount) * check.pricingRules.standout.Price

	premiumDiscount := ApplyBulkDiscount(check.standTotal, check.pricingRules.premium.XforY)
	premiumAdCost := (check.premTotal - premiumDiscount) * check.pricingRules.premium.Price
	return Truncate(classicAdCost + standoutAdCost + premiumAdCost)
}

// ApplyBulkDiscount to an ad that has a X for the price of Y discount
func ApplyBulkDiscount(noOfAds float64, xForY float64) float64 {
	if xForY != 0 {
		return noOfAds / xForY
	}
	return xForY
}
