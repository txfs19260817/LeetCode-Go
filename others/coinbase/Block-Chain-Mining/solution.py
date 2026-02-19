from __future__ import annotations

from dataclasses import dataclass
from typing import Dict, List, Tuple


@dataclass
class Transaction:
    id: str
    size: int
    fee: int
    parent_id: str = ""


@dataclass
class MiningPlan:
    total_fee: int
    used_size: int
    tx_ids: List[str]


class BlockChainMining:
    # Part 1: scalable greedy (not always optimal)
    def max_fee_greedy(self, block_size: int, txs: List[Transaction]) -> MiningPlan:
        if block_size <= 0:
            return MiningPlan(0, 0, [])

        items = txs[:]
        items.sort(
            key=lambda tx: (-tx.fee / tx.size, -tx.fee, tx.size, tx.id),
        )

        used, fee, picked = 0, 0, []
        for tx in items:
            if used + tx.size > block_size:
                continue
            used += tx.size
            fee += tx.fee
            picked.append(tx.id)
        return MiningPlan(fee, used, picked)

    # Part 2: exact 0-1 knapsack DP (easy when block_size is small, e.g. 100)
    def max_fee_optimal(self, block_size: int, txs: List[Transaction]) -> MiningPlan:
        if block_size <= 0:
            return MiningPlan(0, 0, [])

        n = len(txs)
        dp = [[0] * (block_size + 1) for _ in range(n + 1)]
        take = [[False] * (block_size + 1) for _ in range(n + 1)]

        for i in range(1, n + 1):
            tx = txs[i - 1]
            for cap in range(block_size + 1):
                dp[i][cap] = dp[i - 1][cap]
                if tx.size <= cap:
                    candidate = dp[i - 1][cap - tx.size] + tx.fee
                    if candidate > dp[i][cap]:
                        dp[i][cap] = candidate
                        take[i][cap] = True

        best_cap = max(range(block_size + 1), key=lambda c: (dp[n][c], -c))

        picked_ids: List[str] = []
        cap = best_cap
        for i in range(n, 0, -1):
            if take[i][cap]:
                tx = txs[i - 1]
                picked_ids.append(tx.id)
                cap -= tx.size

        picked_ids.sort()
        return MiningPlan(dp[n][best_cap], best_cap, picked_ids)

    # Part 3: parent-child + sibling-branch-exclusive
    # DFS to build per-root path options, then group knapsack DP.
    def max_fee_with_parents(self, block_size: int, txs: List[Transaction]) -> MiningPlan:
        if block_size <= 0:
            return MiningPlan(0, 0, [])

        groups = self._build_groups(txs)
        g = len(groups)
        dp = [[0] * (block_size + 1) for _ in range(g + 1)]
        pick = [[-1] * (block_size + 1) for _ in range(g + 1)]
        prev_cap = [[0] * (block_size + 1) for _ in range(g + 1)]

        for i in range(1, g + 1):
            group = groups[i - 1]
            for cap in range(block_size + 1):
                dp[i][cap] = dp[i - 1][cap]  # skip this group
                prev_cap[i][cap] = cap
                for idx, (opt_size, opt_fee, _) in enumerate(group):
                    if opt_size <= cap:
                        candidate = dp[i - 1][cap - opt_size] + opt_fee
                        if candidate > dp[i][cap]:
                            dp[i][cap] = candidate
                            pick[i][cap] = idx
                            prev_cap[i][cap] = cap - opt_size

        best_cap = max(range(block_size + 1), key=lambda c: (dp[g][c], -c))

        picked_ids: List[str] = []
        cap = best_cap
        for i in range(g, 0, -1):
            idx = pick[i][cap]
            if idx != -1:
                picked_ids.extend(groups[i - 1][idx][2])
            cap = prev_cap[i][cap]

        picked_ids.sort()
        return MiningPlan(dp[g][best_cap], best_cap, picked_ids)

    def _build_groups(
        self, txs: List[Transaction]
    ) -> List[List[Tuple[int, int, List[str]]]]:
        tx_by_id: Dict[str, Transaction] = {tx.id: tx for tx in txs}

        children: Dict[str, List[str]] = {}
        roots: List[str] = []

        for tx in txs:
            if tx.parent_id and tx.parent_id in tx_by_id:
                children.setdefault(tx.parent_id, []).append(tx.id)
            else:
                roots.append(tx.id)

        roots.sort()
        for parent_id in children:
            children[parent_id].sort()

        groups: List[List[Tuple[int, int, List[str]]]] = []
        for root in roots:
            group: List[Tuple[int, int, List[str]]] = []

            def dfs(node_id: str, path: List[str], size: int, fee: int) -> None:
                tx = tx_by_id[node_id]
                new_path = path + [node_id]
                new_size = size + tx.size
                new_fee = fee + tx.fee
                group.append((new_size, new_fee, new_path))
                for child_id in children.get(node_id, []):
                    dfs(child_id, new_path, new_size, new_fee)

            dfs(root, [], 0, 0)
            groups.append(group)
        return groups


if __name__ == "__main__":
    miner = BlockChainMining()

    # Part 1 vs Part 2: greedy fails, DP gets optimum
    base_txs = [
        Transaction("a", 6, 13),
        Transaction("b", 5, 10),
        Transaction("c", 5, 10),
    ]
    greedy = miner.max_fee_greedy(10, base_txs)
    optimal = miner.max_fee_optimal(10, base_txs)
    assert greedy.total_fee == 13 and sorted(greedy.tx_ids) == ["a"]
    assert optimal.total_fee == 20 and sorted(optimal.tx_ids) == ["b", "c"]

    # blockSize=100 pattern
    size_100_txs = [
        Transaction("t1", 60, 150),
        Transaction("t2", 50, 130),
        Transaction("t3", 40, 95),
        Transaction("t4", 30, 65),
    ]
    plan_100 = miner.max_fee_optimal(100, size_100_txs)
    assert plan_100.total_fee == 245
    assert sorted(plan_100.tx_ids) == ["t1", "t3"]

    # Part 3: branch exclusivity + parent dependency
    follow_txs = [
        Transaction("1", 2, 3, ""),
        Transaction("2", 2, 4, "1"),
        Transaction("3", 3, 8, "2"),
        Transaction("4", 3, 7, "2"),
        Transaction("5", 4, 9, ""),
    ]
    plan = miner.max_fee_with_parents(14, follow_txs)
    assert plan.total_fee == 24
    assert sorted(plan.tx_ids) == ["1", "2", "3", "5"]

    print("All asserts passed.")
