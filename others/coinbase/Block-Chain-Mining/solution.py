from __future__ import annotations

from dataclasses import dataclass
from random import choice
from tokenize import group
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
        # 1. Build the forest
        nodes = {tx.id: tx for tx in txs}
        children = {tx.id: [] for tx in txs}
        roots = []

        for tx in txs:
            if (pid := tx.parent_id) and pid in nodes:
                children[pid].append(tx.id)
            else:
                roots.append(tx.id)

        # 2. Extract all valid paths for each tree (group)
        groups = []
        for root in roots:
            paths = []
            def dfs(node_id, cur_ids, cur_size, cur_fee):
                tx = nodes[node_id]
                new_ids = cur_ids + [tx.id]
                new_size = cur_size + tx.size
                new_fee = cur_fee + tx.fee
                if new_size <= block_size:
                    paths.append({'ids': new_ids, 'size': new_size, 'fee': new_fee})
                    for child_id in children[node_id]:
                        dfs(child_id, new_ids, new_size, new_fee)

            dfs(root, [], 0, 0)
            if paths:
                groups.append(paths)

        # 3. DP
        dp = [[0] * (block_size + 1) for _ in range(len(groups) + 1)]
        choice = [[None] * (block_size + 1) for _ in range(len(groups) + 1)]
        for i in range(1, len(groups) + 1):
            group_paths = groups[i-1]
            for w in range(block_size + 1):
                dp[i][w] = dp[i-1][w] # not take any path from this group
                for path in group_paths:
                    if path["size"] <= w:
                        val = dp[i-1][w-path["size"]] + path["fee"]
                        if val > dp[i][w]:
                            dp[i][w] = val
                            choice[i][w] = path
        best_cap = max(range(block_size + 1), key=lambda c: (dp[len(groups)][c], -c))

        # 4. backtrack to find picked IDs
        w = block_size
        picked_ids = []
        for i in range(len(groups), 0, -1):
            chosen_path = choice[i][w]
            if chosen_path is not None:
                w -= chosen_path["size"]
                picked_ids.extend(chosen_path["ids"])
        return MiningPlan(dp[len(groups)][block_size], best_cap, picked_ids)

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
