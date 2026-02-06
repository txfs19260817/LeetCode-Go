from typing import List, Optional
import heapq


class Solution:
    def calculateMinimalTax(self, transactions: List[List[str]]) -> float:
        amounts = {}
        heap: List[int] = []
        tax = 0.0

        for tx in transactions:
            if len(tx) < 4:
                continue
            tx_type = tx[1]
            amount = int(tx[2])
            price = int(tx[3])

            if tx_type == "buy":
                if amounts.get(price, 0) == 0:
                    heapq.heappush(heap, -price)
                amounts[price] = amounts.get(price, 0) + amount
                continue

            if tx_type != "sell":
                continue

            remaining = amount
            while remaining > 0:
                while heap and amounts.get(-heap[0], 0) == 0:
                    heapq.heappop(heap)
                if not heap:
                    break
                buy_price = -heap[0]
                available = amounts[buy_price]
                take = remaining if remaining < available else available
                profit_per_unit = price - buy_price
                if profit_per_unit > 0:
                    tax += (profit_per_unit * take) / 10.0
                amounts[buy_price] -= take
                if amounts[buy_price] == 0:
                    heapq.heappop(heap)
                remaining -= take

        return tax


if __name__ == "__main__":
    solver = Solution()
    assert (
        abs(
            solver.calculateMinimalTax(
                [
                    ["1", "buy", "100", "20"],
                    ["2", "buy", "50", "30"],
                    ["3", "sell", "80", "25"],
                    ["4", "sell", "60", "35"],
                ]
            )
            - 105.0
        )
        < 1e-9
    )

    assert (
        abs(
            solver.calculateMinimalTax(
                [
                    ["1", "buy", "20", "50"],
                    ["2", "sell", "10", "60"],
                    ["3", "buy", "15", "55"],
                    ["4", "sell", "10", "65"],
                    ["5", "sell", "10", "70"],
                ]
            )
            - 37.5
        )
        < 1e-9
    )

    assert (
        abs(
            solver.calculateMinimalTax(
                [
                    ["1", "buy", "10", "10"],
                    ["2", "buy", "20", "20"],
                    ["3", "buy", "30", "105"],
                    ["4", "sell", "10", "100"],
                    ["5", "sell", "20", "120"],
                    ["6", "sell", "30", "50"],
                ]
            )
            - 130.0
        )
        < 1e-9
    )

    assert (
        abs(
            solver.calculateMinimalTax(
                [
                    ["1", "buy", "50", "25"],
                    ["2", "buy", "10", "40"],
                    ["3", "sell", "20", "6"],
                    ["4", "sell", "60", "12"],
                ]
            )
            - 0.0
        )
        < 1e-9
    )

    assert (
        abs(
            solver.calculateMinimalTax(
                [
                    ["1", "buy", "50", "10"],
                    ["2", "sell", "20", "15"],
                    ["3", "buy", "30", "12"],
                    ["4", "sell", "40", "20"],
                    ["5", "buy", "10", "25"],
                    ["6", "buy", "30", "50"],
                    ["7", "sell", "10", "5"],
                ]
            )
            - 44.0
        )
        < 1e-9
    )
