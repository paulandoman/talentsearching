package main

// CustomerPriceRules is a map of pricing rules based on customer
var CustomerPriceRules = map[string]PricingRules{
	"default": PricingRules{
		classic: Pricing{
			Price: 269.99,
		},
		standout: Pricing{
			Price: 322.99,
		},
		premium: Pricing{
			Price: 394.99,
		},
	},
	"unilever": PricingRules{
		classic: Pricing{
			Price: 269.99,
			XforY: 3,
		},
		standout: Pricing{
			Price: 322.99,
		},
		premium: Pricing{
			Price: 394.99,
		},
	},
}
