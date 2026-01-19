package uber

/*
Given a list of currency relationships with exchange values (e.g., BGP -> USD: 109.0). Find the best exchange rate from `from` currency to `to` currency with the following assumptions:

- The rates are bi-directional. If the rate from A to B is `r`, the rate from B to A can be calculated as `1 / r`.
- The goal is to find the maximum achievable exchange rate, which is the highest value obtained by multiplying the rates along a valid exchange path.
- If the conversion is not possible, return -1.

Implement the `CurrencyConverter` class:

- `CurrencyConverter(String[] fromArr, String[] toArr, double[] rateArr)` Initialize the CurrencyConverter class with currency pairs and exchange rates.
- `getBestRate(String from, String to)` Return the best exchange rate that can be obtained from `from` currency to `to` currency.
*/

type CurrencyConverter struct {
	cur2id map[string]int
	adj    []map[int]float64
}

func Constructor(fromArr []string, toArr []string, rateArr []float64) CurrencyConverter {
	cur2id := make(map[string]int)
	getId := func(c string) int {
		if v, ok := cur2id[c]; ok {
			return v
		}
		v := len(cur2id)
		cur2id[c] = v
		return v
	}
	// First pass to assign IDs
	for i := range fromArr {
		getId(fromArr[i])
		getId(toArr[i])
	}
	n := len(cur2id)
	adj := make([]map[int]float64, n)
	for i := range adj {
		adj[i] = make(map[int]float64)
	}

	for i := range fromArr {
		u, v, r := getId(fromArr[i]), getId(toArr[i]), rateArr[i]
		adj[u][v] = max(adj[u][v], r)
		adj[v][u] = max(adj[v][u], 1/r)
	}

	return CurrencyConverter{cur2id, adj}
}

func (cc *CurrencyConverter) GetBestRate(from string, to string) float64 {
	start, ok1 := cc.cur2id[from]
	end, ok2 := cc.cur2id[to]
	if !ok1 || !ok2 {
		return -1.0
	}

	visited := make([]bool, len(cc.adj))

	var dfs func(u int) float64
	dfs = func(u int) float64 {
		if u == end {
			return 1.0
		}

		visited[u] = true
		maxRate := -1.0

		for v, rate := range cc.adj[u] {
			if !visited[v] {
				if res := dfs(v); res != -1.0 {
					maxRate = max(maxRate, rate*res)
				}
			}
		}

		visited[u] = false
		return maxRate
	}

	return dfs(start)
}
