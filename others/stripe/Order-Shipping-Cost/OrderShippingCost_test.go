package stripe

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCalculateShippingCost(t *testing.T) {
	orderUS := Order{
		Country: "US",
		Items: []OrderItem{
			{Product: "mouse", Quantity: 20},
			{Product: "laptop", Quantity: 5},
		},
	}
	orderCA := Order{
		Country: "CA",
		Items: []OrderItem{
			{Product: "mouse", Quantity: 20},
			{Product: "laptop", Quantity: 5},
		},
	}

	// 第一问: 固定单价
	t.Run("variant1_fixed_per_item", func(t *testing.T) {
		shippingCosts := ShippingCosts{
			"US": {
				{Product: "mouse", Cost: intPtr(550)},
				{Product: "laptop", Cost: intPtr(1000)},
			},
			"CA": {
				{Product: "mouse", Cost: intPtr(750)},
				{Product: "laptop", Cost: intPtr(1100)},
			},
		}

		usCost, err := CalculateShippingCost(orderUS, shippingCosts)
		assert.NoError(t, err)
		assert.Equal(t, 16000, usCost)

		caCost, err := CalculateShippingCost(orderCA, shippingCosts)
		assert.NoError(t, err)
		assert.Equal(t, 20500, caCost)
	})

	// 第二问: incremental 逐段累计
	t.Run("variant2_incremental_tiers", func(t *testing.T) {
		shippingCosts := ShippingCosts{
			"US": {
				{
					Product: "mouse",
					Costs: []ShippingTier{
						{MinQuantity: 0, MaxQuantity: nil, Cost: 550},
					},
				},
				{
					Product: "laptop",
					Costs: []ShippingTier{
						{MinQuantity: 0, MaxQuantity: intPtr(2), Cost: 1000},
						{MinQuantity: 3, MaxQuantity: nil, Cost: 900},
					},
				},
			},
			"CA": {
				{
					Product: "mouse",
					Costs: []ShippingTier{
						{MinQuantity: 0, MaxQuantity: nil, Cost: 750},
					},
				},
				{
					Product: "laptop",
					Costs: []ShippingTier{
						{MinQuantity: 0, MaxQuantity: intPtr(2), Cost: 1100},
						{MinQuantity: 3, MaxQuantity: nil, Cost: 1000}, // corrected
					},
				},
			},
		}

		usCost, err := CalculateShippingCost(orderUS, shippingCosts)
		assert.NoError(t, err)
		assert.Equal(t, 15700, usCost)

		caCost, err := CalculateShippingCost(orderCA, shippingCosts)
		assert.NoError(t, err)
		assert.Equal(t, 20200, caCost)
	})

	// 第三问: fixed + incremental 混合
	t.Run("variant3_fixed_and_incremental_mixed", func(t *testing.T) {
		shippingCosts := ShippingCosts{
			"US": {
				{
					Product: "mouse",
					Costs: []ShippingTier{
						{Type: "incremental", MinQuantity: 0, MaxQuantity: nil, Cost: 550},
					},
				},
				{
					Product: "laptop",
					Costs: []ShippingTier{
						{Type: "fixed", MinQuantity: 0, MaxQuantity: intPtr(2), Cost: 1000},
						{Type: "incremental", MinQuantity: 3, MaxQuantity: nil, Cost: 900},
					},
				},
			},
			"CA": {
				{
					Product: "mouse",
					Costs: []ShippingTier{
						{Type: "incremental", MinQuantity: 0, MaxQuantity: nil, Cost: 750},
					},
				},
				{
					Product: "laptop",
					Costs: []ShippingTier{
						{Type: "fixed", MinQuantity: 0, MaxQuantity: intPtr(2), Cost: 1100},
						{Type: "incremental", MinQuantity: 3, MaxQuantity: nil, Cost: 1000}, // corrected
					},
				},
			},
		}

		usCost, err := CalculateShippingCost(orderUS, shippingCosts)
		assert.NoError(t, err)
		assert.Equal(t, 14700, usCost)

		caCost, err := CalculateShippingCost(orderCA, shippingCosts)
		assert.NoError(t, err)
		assert.Equal(t, 19100, caCost)
	})
}

func TestCalculateShippingCost_MissingProductRule(t *testing.T) {
	order := Order{
		Country: "US",
		Items: []OrderItem{
			{Product: "monitor", Quantity: 1},
		},
	}
	shippingCosts := ShippingCosts{
		"US": {
			{Product: "mouse", Cost: intPtr(550)},
		},
	}

	_, err := CalculateShippingCost(order, shippingCosts)
	assert.Error(t, err)
}

func intPtr(v int) *int {
	return &v
}
