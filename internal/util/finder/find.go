package finder

import (
	"spaceship/entity"
)

const ofsset int = 496

func Find(arr []entity.SectorRelation, from, dest int) []int {

	mp := make(map[int][]int, 0)

	for _, item := range arr {
		_, ok := mp[item.SectorId1-ofsset]
		if !ok {
			mp[item.SectorId1-ofsset] = make([]int, 0)
			val := mp[item.SectorId1-ofsset]
			val = append(val, item.SectorId2-ofsset)
			mp[item.SectorId1-ofsset] = val
		} else {
			mp[item.SectorId1-ofsset] = append(mp[item.SectorId1-ofsset], item.SectorId2-ofsset)
		}

		_, ok = mp[item.SectorId2-ofsset]
		if !ok {
			mp[item.SectorId2-ofsset] = make([]int, 0)
			val := mp[item.SectorId2-ofsset]
			val = append(val, item.SectorId1-ofsset)
			mp[item.SectorId2-ofsset] = val
		} else {
			mp[item.SectorId2-ofsset] = append(mp[item.SectorId2-ofsset], item.SectorId1-ofsset)
		}
	}

	matr := matrix(mp, arr)
	routes := dijkstra(len(mp), from-ofsset, matr)
	route := defineRoute(routes, from-ofsset, dest-ofsset)

	return route
}

func defineRoute(union []int, from, to int) []int {
	route := make([]int, 0)

	curr := to
	route = append(route, curr+ofsset)
	for curr != from {
		route = append(route, union[curr]+ofsset)
		curr = union[curr]
	}

	return route
}

func matrix(mp map[int][]int, arr []entity.SectorRelation) [][]int {
	matr := make([][]int, len(mp))

	for i := range matr {
		matr[i] = make([]int, len(mp))

		val := mp[i]

		for _, item := range val {

			weight := 0

			for _, sec := range arr {
				if (sec.SectorId1 == i+ofsset && sec.SectorId2 == item+ofsset) || (sec.SectorId2 == i+ofsset && sec.SectorId1 == item+ofsset) {
					weight = sec.Distance
					break
				}
			}

			matr[i][item] = weight
		}
	}

	for i, _ := range matr {
		for j, _ := range matr[i] {
			if matr[i][j] == 0 {
				matr[i][j] = 10000000000
			}
		}
	}

	return matr
}

// n - количество вершин, s - стартовая вершина,
func dijkstra(n, s int, mp [][]int) []int {
	valid := make([]bool, n)
	for i, _ := range valid {
		valid[i] = true
	}

	weight := make([]int, n)
	for i, _ := range weight {
		//weight[i] = 9223372036854775807
		weight[i] = 10000000000
	}

	route := make([]int, n)
	for i, _ := range route {
		route[i] = -1
	}

	weight[s] = 0
	route[s] = s

	for i := 0; i < n; i++ {
		min_weight := 10000000001
		id_min_weight := -1
		for j := 0; j < n; j++ {
			if valid[j] && weight[j] < min_weight {
				min_weight = weight[j]
				id_min_weight = j
			}
		}
		for z := 0; z < n; z++ {
			if weight[id_min_weight]+mp[id_min_weight][z] < weight[z] {
				weight[z] = weight[id_min_weight] + mp[id_min_weight][z]
				route[z] = id_min_weight
			}
		}
		valid[id_min_weight] = false
	}

	return route
}
