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
    # Base question: high-performance greedy approximation by fee/size.
    def max_fee_greedy(self, block_size: int, txs: List[Transaction]) -> MiningPlan:
        if block_size <= 0:
            return MiningPlan(0, 0, [])

        items = [tx for tx in txs if tx.id and tx.size > 0 and tx.fee >= 0]
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

    # Base question exact solution: sparse 0/1 knapsack DP.
    def max_fee_optimal(self, block_size: int, txs: List[Transaction]) -> MiningPlan:
        if block_size <= 0:
            return MiningPlan(0, 0, [])

        states: Dict[int, MiningPlan] = {0: MiningPlan(0, 0, [])}
        for tx in txs:
            if not tx.id or tx.size <= 0 or tx.fee < 0:
                continue
            next_states = {
                size: MiningPlan(plan.total_fee, plan.used_size, plan.tx_ids[:])
                for size, plan in states.items()
            }
            for used, plan in states.items():
                new_used = used + tx.size
                if new_used > block_size:
                    continue
                candidate = MiningPlan(
                    plan.total_fee + tx.fee, new_used, plan.tx_ids[:] + [tx.id]
                )
                current = next_states.get(new_used)
                if current is None or self._better_plan(candidate, current):
                    next_states[new_used] = candidate
            states = self._prune_states(next_states)
        return self._best_plan(states)

    # Follow-up:
    # 1) child requires parent in same block
    # 2) sibling branches are mutually exclusive
    # We convert each root tree into one group of root->node path options,
    # then run multiple-choice knapsack over groups.
    def max_fee_with_parents(self, block_size: int, txs: List[Transaction]) -> MiningPlan:
        if block_size <= 0:
            return MiningPlan(0, 0, [])

        groups = self._build_groups(txs)
        states: Dict[int, MiningPlan] = {0: MiningPlan(0, 0, [])}
        for group in groups:
            next_states = {
                size: MiningPlan(plan.total_fee, plan.used_size, plan.tx_ids[:])
                for size, plan in states.items()
            }
            for used, plan in states.items():
                for option_size, option_fee, option_ids in group:
                    new_used = used + option_size
                    if new_used > block_size:
                        continue
                    candidate = MiningPlan(
                        plan.total_fee + option_fee,
                        new_used,
                        plan.tx_ids[:] + option_ids,
                    )
                    current = next_states.get(new_used)
                    if current is None or self._better_plan(candidate, current):
                        next_states[new_used] = candidate
            states = self._prune_states(next_states)
        return self._best_plan(states)

    def _build_groups(
        self, txs: List[Transaction]
    ) -> List[List[Tuple[int, int, List[str]]]]:
        tx_by_id: Dict[str, Transaction] = {}
        for tx in txs:
            if tx.id and tx.size > 0 and tx.fee >= 0 and tx.id not in tx_by_id:
                tx_by_id[tx.id] = tx

        if not tx_by_id:
            return []

        children: Dict[str, List[str]] = {}
        roots: List[str] = []

        for tx_id, tx in tx_by_id.items():
            if not tx.parent_id or tx.parent_id not in tx_by_id:
                roots.append(tx_id)
            else:
                children.setdefault(tx.parent_id, []).append(tx_id)

        if not roots:
            roots = sorted(tx_by_id.keys())
        else:
            roots.sort()
        for parent_id in children:
            children[parent_id].sort()

        groups: List[List[Tuple[int, int, List[str]]]] = []
        for root in roots:
            group: List[Tuple[int, int, List[str]]] = []
            stack = set()

            def dfs(tx_id: str, path: List[str], size: int, fee: int) -> None:
                if tx_id in stack:
                    return
                tx = tx_by_id.get(tx_id)
                if tx is None:
                    return
                stack.add(tx_id)
                new_path = path + [tx_id]
                new_size = size + tx.size
                new_fee = fee + tx.fee
                group.append((new_size, new_fee, new_path))
                for child_id in children.get(tx_id, []):
                    dfs(child_id, new_path, new_size, new_fee)
                stack.remove(tx_id)

            dfs(root, [], 0, 0)
            if group:
                groups.append(group)
        return groups

    def _prune_states(self, states: Dict[int, MiningPlan]) -> Dict[int, MiningPlan]:
        best_fee = -1
        pruned: Dict[int, MiningPlan] = {}
        for size in sorted(states.keys()):
            plan = states[size]
            if plan.total_fee > best_fee:
                pruned[size] = plan
                best_fee = plan.total_fee
        return pruned

    def _best_plan(self, states: Dict[int, MiningPlan]) -> MiningPlan:
        best = MiningPlan(0, 0, [])
        initialized = False
        for plan in states.values():
            if not initialized or self._better_plan(plan, best):
                best = plan
                initialized = True
        return best

    def _better_plan(self, a: MiningPlan, b: MiningPlan) -> bool:
        if a.total_fee != b.total_fee:
            return a.total_fee > b.total_fee
        if a.used_size != b.used_size:
            return a.used_size < b.used_size
        return ",".join(a.tx_ids) < ",".join(b.tx_ids)


if __name__ == "__main__":
    miner = BlockChainMining()

    base_txs = [
        Transaction("a", 6, 13),
        Transaction("b", 5, 10),
        Transaction("c", 5, 10),
    ]
    greedy = miner.max_fee_greedy(10, base_txs)
    optimal = miner.max_fee_optimal(10, base_txs)
    assert greedy.total_fee == 13
    assert optimal.total_fee == 20

    follow_txs = [
        Transaction("1", 2, 3, ""),
        Transaction("2", 2, 4, "1"),
        Transaction("3", 3, 8, "2"),
        Transaction("4", 3, 7, "2"),
        Transaction("5", 4, 9, ""),
    ]
    plan = miner.max_fee_with_parents(14, follow_txs)
    assert plan.total_fee == 24
    assert ("3" in plan.tx_ids) != ("4" in plan.tx_ids)
