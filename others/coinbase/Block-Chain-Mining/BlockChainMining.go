package coinbase

import (
	"sort"
	"strings"
)

type Transaction struct {
	ID       string
	Size     int
	Fee      int
	ParentID string
}

type MiningPlan struct {
	TotalFee int
	UsedSize int
	TxIDs    []string
}

type BlockChainMining struct{}

func NewBlockChainMining() *BlockChainMining {
	return &BlockChainMining{}
}

// MaxFeeGreedy implements the interview-friendly fast heuristic for the base question.
// It sorts by fee/size descending and takes transactions while capacity allows.
func (m *BlockChainMining) MaxFeeGreedy(blockSize int, txs []Transaction) MiningPlan {
	if blockSize <= 0 || len(txs) == 0 {
		return MiningPlan{}
	}

	items := make([]Transaction, 0, len(txs))
	for _, tx := range txs {
		if tx.ID == "" || tx.Size <= 0 || tx.Fee < 0 {
			continue
		}
		items = append(items, tx)
	}

	sort.Slice(items, func(i, j int) bool {
		a := items[i]
		b := items[j]
		left := int64(a.Fee) * int64(b.Size)
		right := int64(b.Fee) * int64(a.Size)
		if left != right {
			return left > right
		}
		if a.Fee != b.Fee {
			return a.Fee > b.Fee
		}
		if a.Size != b.Size {
			return a.Size < b.Size
		}
		return a.ID < b.ID
	})

	plan := MiningPlan{TxIDs: make([]string, 0)}
	for _, tx := range items {
		if plan.UsedSize+tx.Size > blockSize {
			continue
		}
		plan.UsedSize += tx.Size
		plan.TotalFee += tx.Fee
		plan.TxIDs = append(plan.TxIDs, tx.ID)
	}
	return plan
}

// MaxFeeOptimal solves the base question exactly (0/1 knapsack), with sparse DP states.
func (m *BlockChainMining) MaxFeeOptimal(blockSize int, txs []Transaction) MiningPlan {
	if blockSize <= 0 || len(txs) == 0 {
		return MiningPlan{}
	}

	states := map[int]MiningPlan{
		0: {TotalFee: 0, UsedSize: 0, TxIDs: []string{}},
	}

	for _, tx := range txs {
		if tx.ID == "" || tx.Size <= 0 || tx.Fee < 0 {
			continue
		}
		next := clonePlanMap(states)
		for used, plan := range states {
			newUsed := used + tx.Size
			if newUsed > blockSize {
				continue
			}
			candidate := MiningPlan{
				TotalFee: plan.TotalFee + tx.Fee,
				UsedSize: newUsed,
				TxIDs:    append(append([]string{}, plan.TxIDs...), tx.ID),
			}
			current, exists := next[newUsed]
			if !exists || betterPlan(candidate, current) {
				next[newUsed] = candidate
			}
		}
		states = pruneStates(next)
	}

	return bestPlan(states)
}

type pathOption struct {
	size int
	fee  int
	ids  []string
}

// MaxFeeWithParents solves the follow-up:
// 1) child can be mined only if parent is included in the same block;
// 2) sibling branches under same parent are mutually exclusive.
// We treat each root tree as one group and enumerate root-to-node paths as options,
// then run a group knapsack (at most one option per group).
func (m *BlockChainMining) MaxFeeWithParents(blockSize int, txs []Transaction) MiningPlan {
	if blockSize <= 0 || len(txs) == 0 {
		return MiningPlan{}
	}

	groups := buildGroups(txs)
	states := map[int]MiningPlan{
		0: {TotalFee: 0, UsedSize: 0, TxIDs: []string{}},
	}

	for _, group := range groups {
		next := clonePlanMap(states) // choose nothing from this group
		for used, plan := range states {
			for _, option := range group {
				newUsed := used + option.size
				if newUsed > blockSize {
					continue
				}
				candidate := MiningPlan{
					TotalFee: plan.TotalFee + option.fee,
					UsedSize: newUsed,
					TxIDs:    append(append([]string{}, plan.TxIDs...), option.ids...),
				}
				current, exists := next[newUsed]
				if !exists || betterPlan(candidate, current) {
					next[newUsed] = candidate
				}
			}
		}
		states = pruneStates(next)
	}

	return bestPlan(states)
}

func buildGroups(txs []Transaction) [][]pathOption {
	txByID := make(map[string]Transaction, len(txs))
	for _, tx := range txs {
		if tx.ID == "" || tx.Size <= 0 || tx.Fee < 0 {
			continue
		}
		if _, exists := txByID[tx.ID]; exists {
			continue
		}
		txByID[tx.ID] = tx
	}
	if len(txByID) == 0 {
		return nil
	}

	children := make(map[string][]string, len(txByID))
	roots := make([]string, 0)

	for id, tx := range txByID {
		if tx.ParentID == "" {
			roots = append(roots, id)
			continue
		}
		if _, ok := txByID[tx.ParentID]; !ok {
			roots = append(roots, id)
			continue
		}
		children[tx.ParentID] = append(children[tx.ParentID], id)
	}

	if len(roots) == 0 {
		for id := range txByID {
			roots = append(roots, id)
		}
	}

	sort.Strings(roots)
	for parent := range children {
		sort.Strings(children[parent])
	}

	groups := make([][]pathOption, 0, len(roots))
	for _, root := range roots {
		group := make([]pathOption, 0)
		stack := make(map[string]bool)
		var dfs func(id string, path []string, size int, fee int)
		dfs = func(id string, path []string, size int, fee int) {
			if stack[id] {
				return
			}
			tx, exists := txByID[id]
			if !exists {
				return
			}
			stack[id] = true

			newPath := append(append([]string{}, path...), id)
			newSize := size + tx.Size
			newFee := fee + tx.Fee
			group = append(group, pathOption{
				size: newSize,
				fee:  newFee,
				ids:  newPath,
			})

			for _, child := range children[id] {
				dfs(child, newPath, newSize, newFee)
			}
			delete(stack, id)
		}
		dfs(root, nil, 0, 0)
		if len(group) > 0 {
			groups = append(groups, group)
		}
	}
	return groups
}

func clonePlanMap(src map[int]MiningPlan) map[int]MiningPlan {
	dst := make(map[int]MiningPlan, len(src))
	for k, v := range src {
		dst[k] = MiningPlan{
			TotalFee: v.TotalFee,
			UsedSize: v.UsedSize,
			TxIDs:    append([]string{}, v.TxIDs...),
		}
	}
	return dst
}

func pruneStates(states map[int]MiningPlan) map[int]MiningPlan {
	if len(states) <= 1 {
		return states
	}
	sizes := make([]int, 0, len(states))
	for size := range states {
		sizes = append(sizes, size)
	}
	sort.Ints(sizes)

	pruned := make(map[int]MiningPlan, len(states))
	bestFeeSoFar := -1
	for _, size := range sizes {
		plan := states[size]
		if plan.TotalFee > bestFeeSoFar {
			pruned[size] = plan
			bestFeeSoFar = plan.TotalFee
		}
	}
	return pruned
}

func bestPlan(states map[int]MiningPlan) MiningPlan {
	best := MiningPlan{TxIDs: []string{}}
	initialized := false
	for _, plan := range states {
		if !initialized || betterPlan(plan, best) {
			best = plan
			initialized = true
		}
	}
	return best
}

func betterPlan(a MiningPlan, b MiningPlan) bool {
	if a.TotalFee != b.TotalFee {
		return a.TotalFee > b.TotalFee
	}
	if a.UsedSize != b.UsedSize {
		return a.UsedSize < b.UsedSize
	}
	return strings.Join(a.TxIDs, ",") < strings.Join(b.TxIDs, ",")
}
