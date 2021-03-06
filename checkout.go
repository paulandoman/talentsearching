package main

import (
	"encoding/json"
	"fmt"
	"math"
)

// Checkout represents an individual checkout
type Checkout struct {
	pricingRules PricingRules
	classTotal   float64
	standTotal   float64
	premTotal    float64
}

// Add ads to the checkout
func (check *Checkout) Add(args ...Item) {
	for _, i := range args {
		switch i.adType {
		case classic:
			fmt.Println("classic added")
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
}

// Delete ad from checkout
func (check *Checkout) Delete(args ...Item) {
	for _, i := range args {
		switch i.adType {
		case classic:
			if check.classTotal > 0 {
				check.classTotal--
				fmt.Println("classic removed ")
			} else {
				fmt.Println("no classic ads to remove")
			}
		case standout:
			if check.standTotal > 0 {
				check.standTotal--
				fmt.Println("standout removed")
			} else {
				fmt.Println("no standout ads to remove")
			}
		case premium:
			if check.premTotal > 0 {
				check.premTotal--
				fmt.Println("premium removed ")
			} else {
				fmt.Println("no premium ads to remove")
			}
		default:
			fmt.Println("nothing removed")
		}
	}
}

// Total adds up the total cost of the ads based on the customer
func (check *Checkout) Total() float64 {
	classicDiscount := ApplyBulkDiscount(check.classTotal, check.pricingRules.classic.XforY)
	classicAdPrice := GetAdPrice(check.classTotal, check.pricingRules.classic)
	classicAdTotal := (check.classTotal - classicDiscount) * classicAdPrice

	standoutDiscount := ApplyBulkDiscount(check.standTotal, check.pricingRules.standout.XforY)
	standoutAdPrice := GetAdPrice(check.standTotal, check.pricingRules.standout)
	standoutAdTotal := (check.standTotal - standoutDiscount) * standoutAdPrice

	premiumDiscount := ApplyBulkDiscount(check.standTotal, check.pricingRules.premium.XforY)
	premiumAdPrice := GetAdPrice(check.premTotal, check.pricingRules.premium)
	premiumAdTotal := (check.premTotal - premiumDiscount) * premiumAdPrice

	return Truncate(classicAdTotal + standoutAdTotal + premiumAdTotal)
}

// Show the checkout totals in JSON format
func (check *Checkout) Show() string {
	type Display struct {
		Classic  float64
		Standout float64
		Premium  float64
	}

	display := &Display{
		Classic:  check.classTotal,
		Standout: check.standTotal,
		Premium:  check.premTotal,
	}
	b, err := json.Marshal(display)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	return string(b)
}

// ApplyBulkDiscount to an ad that has a X for the price of Y discount
func ApplyBulkDiscount(noOfAds float64, xForY float64) float64 {
	if xForY != 0 {
		return math.Trunc(noOfAds / xForY)
	}
	return xForY
}

// GetAdPrice - get the ad price taking into account any bulk discounts
func GetAdPrice(noOfAds float64, pricing Pricing) float64 {
	if (pricing.BulkNo != 0) && (noOfAds >= pricing.BulkNo) {
		return pricing.BulkPrice
	}
	return pricing.Price
}

// Truncate a float to two levels of precision
func Truncate(some float64) float64 {
	return float64(int(some*100)) / 100
}
