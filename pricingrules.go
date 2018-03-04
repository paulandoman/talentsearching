package main

// DefaultClassicPrice - non special 'classic' ad price
const DefaultClassicPrice = 269.99

// DefaultStandoutPrice - non special 'standout' ad price
const DefaultStandoutPrice = 322.99

// DefaultPremiumPrice - non special 'premium' ad price
const DefaultPremiumPrice = 394.99

// CustomerPriceRules - Pricing rules based on customer
var CustomerPriceRules = map[string]PricingRules{
	"default": PricingRules{
		classic: Pricing{
			Price: DefaultClassicPrice,
		},
		standout: Pricing{
			Price: DefaultStandoutPrice,
		},
		premium: Pricing{
			Price: DefaultPremiumPrice,
		},
	},
	"unilever": PricingRules{
		classic: Pricing{
			Price: DefaultClassicPrice,
			XforY: 3, // where Y = X-1
		},
		standout: Pricing{
			Price: DefaultStandoutPrice,
		},
		premium: Pricing{
			Price: DefaultPremiumPrice,
		},
	},
	"apple": PricingRules{
		classic: Pricing{
			Price: DefaultClassicPrice,
		},
		standout: Pricing{
			Price: 299.99,
		},
		premium: Pricing{
			Price: DefaultPremiumPrice,
		},
	},
	"nike": PricingRules{
		classic: Pricing{
			Price: DefaultClassicPrice,
		},
		standout: Pricing{
			Price: DefaultStandoutPrice,
		},
		premium: Pricing{
			Price:     DefaultPremiumPrice,
			BulkNo:    4,
			BulkPrice: 379.99,
		},
	},
	"ford": PricingRules{
		classic: Pricing{
			Price: DefaultClassicPrice,
			XforY: 5, // where Y = X-1
		},
		standout: Pricing{
			Price: 309.99,
		},
		premium: Pricing{
			Price:     DefaultPremiumPrice,
			BulkNo:    3,
			BulkPrice: 389.99,
		},
	},
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
	BulkNo    float64
	BulkPrice float64
}

// GetRules returns pricing rules based on customer
func GetRules(customer string) PricingRules {
	values, queryPresent := CustomerPriceRules[customer]
	if queryPresent {
		return values
	}
	return CustomerPriceRules["default"]
}
