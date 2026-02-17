package stripe

import (
	"fmt"
	"sort"
)

type Order struct {
	Country string      `json:"country"`
	Items   []OrderItem `json:"items"`
}

type OrderItem struct {
	Product  string `json:"product"`
	Quantity int    `json:"quantity"`
}

type ShippingCosts map[string][]ProductShippingRule

type ProductShippingRule struct {
	Product string         `json:"product"`
	Cost    *int           `json:"cost,omitempty"`
	Costs   []ShippingTier `json:"costs,omitempty"`
}

type ShippingTier struct {
	Type        string `json:"type,omitempty"`
	MinQuantity int    `json:"minQuantity"`
	MaxQuantity *int   `json:"maxQuantity"`
	Cost        int    `json:"cost"`
}

func CalculateShippingCost(order Order, shippingCosts ShippingCosts) (int, error) {
	countryRules, ok := shippingCosts[order.Country]
	if !ok {
		return 0, fmt.Errorf("unsupported country: %s", order.Country)
	}

	productToRule := make(map[string]ProductShippingRule, len(countryRules))
	for _, rule := range countryRules {
		productToRule[rule.Product] = rule
	}

	total := 0
	for _, item := range order.Items {
		rule, found := productToRule[item.Product]
		if !found {
			return 0, fmt.Errorf("missing shipping rule for product=%s country=%s", item.Product, order.Country)
		}
		productCost, err := calculateProductCost(item.Quantity, rule)
		if err != nil {
			return 0, err
		}
		total += productCost
	}

	return total, nil
}

func calculateProductCost(quantity int, rule ProductShippingRule) (int, error) {
	if quantity < 0 {
		return 0, fmt.Errorf("quantity must be non-negative")
	}
	if quantity == 0 {
		return 0, nil
	}

	// Variant 1: fixed per-item price.
	if rule.Cost != nil && len(rule.Costs) == 0 {
		return quantity * (*rule.Cost), nil
	}

	if len(rule.Costs) == 0 {
		return 0, fmt.Errorf("missing shipping cost tiers for product=%s", rule.Product)
	}

	tiers := append([]ShippingTier(nil), rule.Costs...)
	sort.Slice(tiers, func(i, j int) bool {
		return tiers[i].MinQuantity < tiers[j].MinQuantity
	})

	total := 0
	matched := false
	for _, tier := range tiers {
		units := unitsInTier(quantity, tier.MinQuantity, tier.MaxQuantity)
		if units <= 0 {
			continue
		}
		matched = true

		tierType := tier.Type
		if tierType == "" {
			tierType = "incremental"
		}

		switch tierType {
		case "fixed":
			total += tier.Cost
		case "incremental":
			total += units * tier.Cost
		default:
			return 0, fmt.Errorf("unknown tier type: %s", tierType)
		}
	}

	if !matched {
		return 0, fmt.Errorf("no pricing tier matched quantity=%d product=%s", quantity, rule.Product)
	}
	return total, nil
}

func unitsInTier(quantity, minQuantity int, maxQuantity *int) int {
	if quantity <= 0 {
		return 0
	}

	start := tierStart(minQuantity)
	if quantity < start {
		return 0
	}

	end := quantity
	if maxQuantity != nil {
		end = *maxQuantity
	}
	if end < start {
		return 0
	}

	if end > quantity {
		end = quantity
	}
	return end - start + 1
}

func tierStart(minQuantity int) int {
	if minQuantity <= 0 {
		return 1
	}
	return minQuantity
}
