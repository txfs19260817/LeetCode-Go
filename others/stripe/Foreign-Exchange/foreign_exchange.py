from collections import defaultdict, deque
class Convert:
	def __init__(self, str):
		self.cur_map = self._extract_conversion(str)

	def _extract_conversion(self, str):
		cur_map = defaultdict(dict)
		for currency_pair in str.split(","):
			base, quote, amt = currency_pair.split(":")
			cur_map[base][quote] = float(amt)
			cur_map[quote][base] = 1 / float(amt)
		return cur_map

	def fx(self, base, quote):
		if base not in self.cur_map and quote not in self.cur_map:
			return None
		if base == quote:
			return 1.0
		# if base in self.cur_map and quote in self.cur_map[base]:
		# 	return self.cur_map[base][quote]
		visited = set()
		q = deque([(base, 1.0,1)])
		best_rate = 0

		while q:
			current_currency, current_rate, level = q.popleft()
			visited.add(current_currency)
			for next_currency, next_rate in self.cur_map[current_currency].items():
				if next_currency == quote:
					best_rate = max(best_rate, current_rate * next_rate)
				elif next_currency not in visited:
					q.append((next_currency, current_rate * next_rate, level + 1))
			if level == 4:
				break
		return best_rate

convert = Convert("USD:CAD:1.40,GBP:JPY:200,USD:GBP:0.8,CAD:JPY:100")

test_cases = [
	# Direct
	("USD", "CAD", 1.60),
	("CAD", "USD", 1/1.40),
	# Hop
	("USD", "JPY", 0.8 * 200), # USD -> CAD -> JPY
	("GBP", "JPY", 200), 
	("CAD", "GBP", 1/1.40 * 0.8),
	("JPY", "CAD", 1/100)
]

for base, quote, expected in test_cases:
	result = convert.fx(base, quote)
	assert result == expected
