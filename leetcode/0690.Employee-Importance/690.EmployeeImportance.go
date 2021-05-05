package _690_Employee_Importance

type Employee struct {
	Id           int
	Importance   int
	Subordinates []int
}

func getImportance(employees []*Employee, id int) int {
	var ans int
	var q []int
	id2e := map[int]*Employee{}
	for _, e := range employees {
		id2e[e.Id] = e
	}
	if target, ok := id2e[id]; ok {
		ans += target.Importance
		q = append(q, target.Subordinates...)
	}
	for len(q) > 0 {
		sid := q[0]
		q = q[1:]
		if target, ok := id2e[sid]; ok {
			ans += target.Importance
			q = append(q, target.Subordinates...)
		}
	}
	return ans
}
