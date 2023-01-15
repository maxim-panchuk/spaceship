package finder

import (
	"fmt"
	"spaceship/entity"
)

const ofsset int = 496

func Find(arr []entity.SectorRelation, from, dest int) int {

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

	for key, val := range mp {
		fmt.Printf("%v -> %v\n", key, val)
	}

	matr := matrix(mp, arr)

	for _, i := range matr {
		for _, j := range i {
			fmt.Printf("%v ", j)
		}

		fmt.Println()
	}

	weights := dijkstra(len(mp), from-ofsset, matr)

	fmt.Println()
	fmt.Println(weights)

	return -1
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

	weight[s] = 0

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
			}
		}
		valid[id_min_weight] = false
	}

	return weight
}

// func contains(arr []int, val int) bool {
// 	for _, item := range arr {
// 		if item == val {
// 			return true
// 		}
// 	}

// 	return false
// }
