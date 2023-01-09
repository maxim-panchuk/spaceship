package secrel

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"spaceship/internal/generator/sector"
	"spaceship/internal/util"
	"time"
)

type SecrelGenerator struct {
	repo    SecrelRepository
	secRepo sector.SectorRepository
}

func NewSecrelGenerator(repo SecrelRepository, secRepo sector.SectorRepository) *SecrelGenerator {
	return &SecrelGenerator{
		repo:    repo,
		secRepo: secRepo,
	}
}

func (g *SecrelGenerator) Generate() {
	sectors, err := g.secRepo.GetSectors()

	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	nodes := make(map[int][]int)

	sectorsNum := len(sectors)
	vertexAmount := int(util.Round(float64(sectorsNum) * 0.4))

	sectorIdSlice := make([]int, 0)

	for _, item := range sectors {
		nodes[item.SectorId] = make([]int, 0)
		sectorIdSlice = append(sectorIdSlice, item.SectorId)
	}

	for key := range nodes {
		shuffledArr := getShuffledArr(sectorIdSlice)

		i := len(nodes[key])

		for _, item := range *shuffledArr {
			if i >= vertexAmount {
				break
			}
			if item != key && !contains(nodes[key], item) {
				nodes[key] = append(nodes[key], item)

				distance := util.NewCryptoRand(100, 1000000)

				id, id1, id2, d, err := g.repo.Insert(key, item, int(distance))

				if err != nil {
					os.Exit(1)
				}

				fmt.Printf("Relation Created!: id: %v, id1: %v, id2: %v, distance: %v\n",
					id, id1, id2, d)

				nodes[item] = append(nodes[item], key)
				i++
			}
		}
	}

}

func getShuffledArr(arr []int) *[]int {
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(arr), func(i, j int) { arr[i], arr[j] = arr[j], arr[i] })
	return &arr
}

func contains(arr []int, val int) bool {
	for _, item := range arr {
		if item == val {
			return true
		}
	}
	return false
}
