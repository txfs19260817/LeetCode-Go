package uber

import "strings"

// Evaluate parses and evaluates an expression built from:
//   - add(x, y): returns x + y
//   - sub(x, y): returns x - y
//
// Expressions can be nested, contain spaces, and integers may be negative.
func Evaluate(expression string) int {
	p := &P{s: strings.ReplaceAll(expression, " ", "")}
	return p.expr()
}

// P is a tiny recursive-descent parser over the input string.
type P struct {
	s string // full input
	i int    // current index into s
}

func (p *P) expr() int {
	if isAdd := strings.HasPrefix(p.s[p.i:], "add"); isAdd || strings.HasPrefix(p.s[p.i:], "sub") {
		p.i += 4 // skip "add(" / "sub("
		a := p.expr()
		p.i++ // skip ','
		b := p.expr()
		p.i++ // skip ')'

		if isAdd {
			return a + b
		}
		return a - b
	}
	return p.num()
}

func (p *P) num() (v int) {
	sign := 1
	if p.s[p.i] == '-' {
		sign = -1
		p.i++
	}

	for ; p.i < len(p.s) && p.s[p.i] >= '0' && p.s[p.i] <= '9'; p.i++ {
		v = v*10 + int(p.s[p.i]-'0')
	}
	return sign * v
}
