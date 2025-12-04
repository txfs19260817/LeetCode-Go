package leetcode

import "slices"

type ThroneInheritance struct {
	kingName string
	edge     map[string][]string
	dead     map[string]bool
}

func Constructor(kingName string) ThroneInheritance {
	return ThroneInheritance{kingName, make(map[string][]string), make(map[string]bool)}
}

func (t *ThroneInheritance) Birth(parentName string, childName string) {
	t.edge[parentName] = append(t.edge[parentName], childName)
}

func (t *ThroneInheritance) Death(name string) {
	t.dead[name] = true
}

func (t *ThroneInheritance) GetInheritanceOrder() []string {
	var ans []string
	s := []string{t.kingName}
	for len(s) > 0 {
		p := s[len(s)-1]
		s = s[:len(s)-1]

		if !t.dead[p] {
			ans = append(ans, p)
		}

		for _, child := range slices.Backward(t.edge[p]) {
			s = append(s, child)
		}
	}
	return ans
}
