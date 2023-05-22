package leetcode

import "bytes"

/*
AlphaBets	Numbers
e			0 1 3 5 7 8 9
f			4 5
g			8
h			3 8
i			5 6 8 9
n			1 7 9
o			0 1 2 4
r			0 3 4
s			6 7
t			2 3 8
u			4
v			5 7
w			2
x			6
z			0
*/

func originalDigits(s string) string {
	rune2cnt := map[rune]int{}
	for _, c := range s {
		rune2cnt[c]++
	}
	var cnt [10]int
	cnt[0] = rune2cnt['z']
	cnt[2] = rune2cnt['w']
	cnt[4] = rune2cnt['u']
	cnt[6] = rune2cnt['x']
	cnt[8] = rune2cnt['g']
	cnt[3] = rune2cnt['h'] - cnt[8]
	cnt[5] = rune2cnt['f'] - cnt[4]
	cnt[7] = rune2cnt['s'] - cnt[6]
	cnt[1] = rune2cnt['o'] - cnt[0] - cnt[2] - cnt[4]
	cnt[9] = rune2cnt['i'] - cnt[5] - cnt[6] - cnt[8]
	var ans []byte
	for i, c := range cnt {
		ans = append(ans, bytes.Repeat([]byte{byte('0' + i)}, c)...)
	}
	return string(ans)
}
