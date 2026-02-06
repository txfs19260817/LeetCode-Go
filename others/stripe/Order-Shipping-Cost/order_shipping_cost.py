def calculate_shipping_cost(order, shipping_cost):
    # Get the country from the order
    country = order["country"]
    
    # Get the items from the order
    items = order["items"]
    
    # Get the shipping cost list for the country
    country_shipping_cost = shipping_cost[country]
    
    # Initialize total cost
    total_cost = 0
    
    # Loop through each item in the order
    for item in items:
        product = item["product"]
        quantity = item["quantity"]
        
        # Find the shipping cost for this product in the country
        for shipping in country_shipping_cost:
            if shipping["product"] == product:
                remaining_qty = quantity
                
                # Process the tiered pricing for this product
                for cost_range in shipping["costs"]:
                    min_qty = cost_range["minQuantity"]
                    max_qty = cost_range["maxQuantity"]
                    
                    # Calculate how many units fit into this pricing tier
                    if max_qty is None:
                        # No upper limit, charge all remaining units at this tier
                        units_in_tier = remaining_qty
                    else:
                        # Calculate how many units are within this tier
                        tier_range = max_qty - min_qty
                        if min_qty != 0:
                            tier_range += 1
                        units_in_tier = min(remaining_qty, tier_range)
                    
                    # Calculate the cost for the units in this tier
                    total_cost += units_in_tier * cost_range["cost"]
                    remaining_qty -= units_in_tier
                    
                    # If no more units to charge, break out of the loop
                    if remaining_qty <= 0:
                        break
    
    return total_cost

# # Example input for US order
order_us = {
    "country": "US",
    "items": [
        {"product": "mouse", "quantity": 20},
        {"product": "laptop", "quantity": 5}
    ]
}

# Example input for CA order
order_ca = {
    "country": "CA",
    "items": [
        {"product": "mouse", "quantity": 20},
        {"product": "laptop", "quantity": 5}
    ]
}

shipping_cost = {
    "US": [
        {
            "product": "mouse",
            "costs": [
                {
                    "minQuantity": 0,
                    "maxQuantity": None,
                    "cost": 550
                }
            ]
        },
        {
            "product": "laptop",
            "costs": [
                {
                    "minQuantity": 0,
                    "maxQuantity": 2,
                    "cost": 1000
                },
                {
                    "minQuantity": 3,
                    "maxQuantity": 4,
                    "cost": 900
                },
                {
                    "minQuantity": 5,
                    "maxQuantity": None,
                    "cost": 800
                }
            ]
        }
    ],
    "CA": [
        {
            "product": "mouse",
            "costs": [
                {
                    "minQuantity": 0,
                    "maxQuantity": None,
                    "cost": 750
                }
            ]
        },
        {
            "product": "laptop",
            "costs": [
                {
                    "minQuantity": 0,
                    "maxQuantity": 2,
                    "cost": 1100
                },
                {
                    "minQuantity": 3,
                    "maxQuantity": None,
                    "cost": 1000
                }
            ]
        }
    ]
}

# Calculate the total shipping cost for US and CA orders
# total_shipping_cost_us = calculate_shipping_cost(order_us, shipping_cost)
# total_shipping_cost_ca = calculate_shipping_cost(order_ca, shipping_cost)

# Output the results
# print(f"Total shipping cost for US order: {total_shipping_cost_us}")  # Expected: 15700
# print(f"Total shipping cost for CA order: {total_shipping_cost_ca}")  # Expected: 20200


shipping_cost = {
  "US": [
  {
    "product": "mouse",
    "costs": [
    {
      "type": "incremental",
      "minQuantity": 0,
      "maxQuantity": None,
      "cost": 550
    }
    ]
  },
  {
    "product": "laptop",
    "costs": [
    {
      "type": "fixed",
      "minQuantity": 0,
      "maxQuantity": 2,
      "cost": 1000
    },
    {
      "type": "incremental",
      "minQuantity": 3,
      "maxQuantity": None,
      "cost": 900
    }
    ]
  }
  ],
  "CA": [
  {
    "product": "mouse",
    "costs": [
    {
      "type": "incremental",
      "minQuantity": 0,
      "maxQuantity": None,
      "cost": 750
    }
    ]
  },
  {
    "product": "laptop",
    "costs": [
    {
      "type": "fixed",
      "minQuantity": 0,
      "maxQuantity": 2,
      "cost": 1100
    },
    {
      "type": "incremental",
      "minQuantity": 3,
      "maxQuantity": None,
      "cost": 1000c
    }
    ]
  }
  ]
}

def calculate_shipping_cost(order, shipping_cost):
    # Get the country from the order
    country = order["country"]
    
    # Get the items from the order
    items = order["items"]
    
    # Get the shipping cost list for the country
    country_shipping_cost = shipping_cost[country]
    
    # Initialize total cost
    total_cost = 0
    
    # Loop through each item in the order
    for item in items:
        product = item["product"]
        quantity = item["quantity"]
        
        # Find the shipping cost for this product in the country
        for shipping in country_shipping_cost:
            if shipping["product"] == product:
                remaining_qty = quantity
                
                # Process the tiered pricing for this product
                for cost_range in shipping["costs"]:
                    min_qty = cost_range["minQuantity"]
                    max_qty = cost_range["maxQuantity"]
                    
                    # Calculate how many units fit into this pricing tier
                    if max_qty is None:
                        # No upper limit, charge all remaining units at this tier
                        units_in_tier = remaining_qty
                    else:
                        # Calculate how many units are within this tier
                        tier_range = max_qty - min_qty
                        if min_qty != 0:
                            tier_range += 1
                        units_in_tier = min(remaining_qty, tier_range)
                    
                    # Calculate the cost for the units in this tier
                    type = cost_range['type']
                    if type == 'fixed':
                        total_cost += cost_range["cost"]
                    elif type == 'incremental':
                        total_cost += units_in_tier * cost_range["cost"]
                    remaining_qty -= units_in_tier
                    
                    # If no more units to charge, break out of the loop
                    if remaining_qty <= 0:
                        break
    
    return total_cost

total_shipping_cost_us = calculate_shipping_cost(order_us, shipping_cost)
total_shipping_cost_ca = calculate_shipping_cost(order_ca, shipping_cost)

print(f"Total shipping cost for US order: {total_shipping_cost_us}")  # Expected: 15700
print(f"Total shipping cost for CA order: {total_shipping_cost_ca}")  # Expected: 20200
