from typing import Any, Dict, List, Optional


def _units_in_tier(quantity: int, min_quantity: int, max_quantity: Optional[int]) -> int:
    # Tier lower bound:
    # In this problem, minQuantity=0 means "start from unit 1".
    start = 1 if min_quantity <= 0 else min_quantity

    # Tier upper bound:
    # maxQuantity=None means this tier has no explicit upper bound,
    # so we cap it by current order quantity.
    end = quantity if max_quantity is None else max_quantity

    # Overlap length between [1, quantity] and [start, end].
    overlap_end = min(quantity, end)
    return max(0, overlap_end - start + 1)


def _calculate_product_cost(quantity: int, product_rule: Dict[str, Any]) -> int:
    if quantity == 0:
        return 0

    # 第一问: fixed per-item product cost.
    if "cost" in product_rule and "costs" not in product_rule:
        return quantity * product_rule["cost"]

    tiers: List[Dict[str, Any]] = product_rule.get("costs", [])

    total = 0

    for tier in sorted(tiers, key=lambda x: x.get("minQuantity", 0)):
        units = _units_in_tier(
            quantity,
            tier.get("minQuantity", 0),
            tier.get("maxQuantity"),
        )
        if units <= 0:
            continue

        tier_type = tier.get("type", "incremental")
        if tier_type == "fixed":
            total += tier["cost"]
        else:
            total += units * tier["cost"]

    return total


def calculate_shipping_cost(order: Dict[str, Any], shipping_costs: Dict[str, Any]) -> int:
    country = order["country"]
    product_to_rule = {rule["product"]: rule for rule in shipping_costs[country]}

    total = 0
    for item in order["items"]:
        product = item["product"]
        quantity = item["quantity"]
        total += _calculate_product_cost(quantity, product_to_rule[product])

    return total


if __name__ == "__main__":
    order_us = {
        "country": "US",
        "items": [
            {"product": "mouse", "quantity": 20},
            {"product": "laptop", "quantity": 5},
        ],
    }
    order_ca = {
        "country": "CA",
        "items": [
            {"product": "mouse", "quantity": 20},
            {"product": "laptop", "quantity": 5},
        ],
    }

    # 第一问: 固定单价
    shipping_cost_v1 = {
        "US": [
            {"product": "mouse", "cost": 550},
            {"product": "laptop", "cost": 1000},
        ],
        "CA": [
            {"product": "mouse", "cost": 750},
            {"product": "laptop", "cost": 1100},
        ],
    }
    assert calculate_shipping_cost(order_us, shipping_cost_v1) == 16000
    assert calculate_shipping_cost(order_ca, shipping_cost_v1) == 20500

    # 第二问: 按区间 incremental 逐段累计
    shipping_cost_v2 = {
        "US": [
            {
                "product": "mouse",
                "costs": [{"minQuantity": 0, "maxQuantity": None, "cost": 550}],
            },
            {
                "product": "laptop",
                "costs": [
                    {"minQuantity": 0, "maxQuantity": 2, "cost": 1000},
                    {"minQuantity": 3, "maxQuantity": None, "cost": 900},
                ],
            },
        ],
        "CA": [
            {
                "product": "mouse",
                "costs": [{"minQuantity": 0, "maxQuantity": None, "cost": 750}],
            },
            {
                "product": "laptop",
                "costs": [
                    {"minQuantity": 0, "maxQuantity": 2, "cost": 1100},
                    {"minQuantity": 3, "maxQuantity": None, "cost": 1000},
                ],
            },
        ],
    }
    assert calculate_shipping_cost(order_us, shipping_cost_v2) == 15700
    assert calculate_shipping_cost(order_ca, shipping_cost_v2) == 20200

    # 第三问: 区间内支持 fixed + incremental 混合
    shipping_cost_v3 = {
        "US": [
            {
                "product": "mouse",
                "costs": [
                    {
                        "type": "incremental",
                        "minQuantity": 0,
                        "maxQuantity": None,
                        "cost": 550,
                    }
                ],
            },
            {
                "product": "laptop",
                "costs": [
                    {"type": "fixed", "minQuantity": 0, "maxQuantity": 2, "cost": 1000},
                    {
                        "type": "incremental",
                        "minQuantity": 3,
                        "maxQuantity": None,
                        "cost": 900,
                    },
                ],
            },
        ],
        "CA": [
            {
                "product": "mouse",
                "costs": [
                    {
                        "type": "incremental",
                        "minQuantity": 0,
                        "maxQuantity": None,
                        "cost": 750,
                    }
                ],
            },
            {
                "product": "laptop",
                "costs": [
                    {"type": "fixed", "minQuantity": 0, "maxQuantity": 2, "cost": 1100},
                    {
                        "type": "incremental",
                        "minQuantity": 3,
                        "maxQuantity": None,
                        "cost": 1000,
                    },
                ],
            },
        ],
    }
    assert calculate_shipping_cost(order_us, shipping_cost_v3) == 14700
    assert calculate_shipping_cost(order_ca, shipping_cost_v3) == 19100

    print("All assertions passed.")
