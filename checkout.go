package main

func main() {

}

// JobAd does something
type JobAd int

const (
	classic JobAd = iota
	standout
	premium
)

// Item does something
type Item struct {
	id    JobAd
	price float32
}

// Checkout does something
type Checkout struct {
	pricingRules string
	classicAds   int
	standoutAds  int
	premiumAds   int
}

// Add does something
func (check *Checkout) Add(i Item) {
	switch i.id {
	case classic:
		check.classicAds++
	case standout:
		check.standoutAds++
	case premium:
		check.premiumAds++
	}
}
