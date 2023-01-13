package finder

import "spaceship/entity"

func Find(arr []entity.SectorRelation, from, dest int) int {

	mp := make(map[int][]int, 0)

	for _, item := range arr {
		children, ok := mp[item.SectorId1]
		if !ok {
			mp[item.SectorId1] = make([]int, 0)
			val := mp[item.SectorId1]
			val = append(val, item.SectorId2)
		}
		children = append(children, item.SectorId2)
	}

	q := make([]int, 0)
	s := make([]int, 0)

	q = append(q, from)
	s = append(s, from)

	length := len(q)
	res := 0

	for length > 0 {
		ql := len(q)
		for i := 0; i < ql; i++ {
			curr := q[0]
			if curr == dest {
				return res
			}
			for _, node := range mp[curr] {
				if !contains(s, node) {
					q = append(q, node)
					s = append(s, node)
				}
			}
			q = q[1:]
		}
		res += 1
	}

	return -1

}

func contains(arr []int, val int) bool {
	for _, item := range arr {
		if item == val {
			return true
		}
	}

	return false
}
