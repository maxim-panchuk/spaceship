package planet

import (
	"fmt"
	"log"
	"os"
	"spaceship/internal/generator/sector"
	"spaceship/internal/util"
)

type PlanetGenerator struct {
	repo    PlanetRepository
	secRepo sector.SectorRepository
}

func NewPlanetGenerator(repo PlanetRepository, secRepo sector.SectorRepository) *PlanetGenerator {
	return &PlanetGenerator{
		repo:    repo,
		secRepo: secRepo,
	}
}

const (
	MIN_PLANET = 3
	MAX_PLANET = 14

	MIN_NAME = 1
	MAX_NAME = 1000
)

// Для каждого сектора сгенерирует планеты в нем
func (g *PlanetGenerator) Generate() {

	sectors, err := g.secRepo.GetSectors()

	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	for _, item := range sectors {
		currPlanetAmount := util.NewCryptoRand(MIN_PLANET, MAX_PLANET) + MIN_PLANET
		for i := 0; i < int(currPlanetAmount); i++ {
			planetNameNumeric := util.NewCryptoRand(MIN_NAME, MAX_NAME) + MIN_NAME
			planetName := fmt.Sprintf("%s-%v", item.SectorName, planetNameNumeric)
			planetSectorId := item.SectorId

			planetId, _planetName, sectorId, err := g.repo.Insert(planetName, planetSectorId)

			if err != nil {
				log.Fatal(err)
				os.Exit(1)
			}

			fmt.Printf("Planet created!: planet_id: %v, planet_name: %s, sector_id: %v\n",
				planetId, _planetName, sectorId)
		}
	}
}
