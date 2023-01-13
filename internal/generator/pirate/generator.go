package pirate

import (
	"fmt"
	"log"
	"os"
	"spaceship/internal/generator/sector"
	"spaceship/internal/util"
)

type PirateGen struct {
	pirateRepo PirateRepo
	sectorRepo sector.SectorRepository
}

func NewPirateGen(pirateRepo PirateRepo, sectorRepo sector.SectorRepository) *PirateGen {
	return &PirateGen{
		pirateRepo: pirateRepo,
		sectorRepo: sectorRepo,
	}
}

func (g *PirateGen) Generate() {
	sectorSlice, err := g.sectorRepo.GetSectors()

	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	for _, sector := range sectorSlice {
		pirateScale := int(sector.PirateScale * 10)

		for i := 0; i < pirateScale; i++ {
			piratePower := util.NewCryptoRand(1, pirateScale+1)
			if piratePower == 0 {
				piratePower = 1
			}
			pirateName := fmt.Sprintf("pirate-%s-%v", sector.SectorName, piratePower)

			id, err := g.pirateRepo.Insert(pirateName, int(piratePower), sector.SectorId)

			if err != nil {
				log.Fatal(err)
				os.Exit(1)
			}

			fmt.Printf("Pirate created!: id: %v, name: %v, power: %v, sector: %v\n", id, pirateName, piratePower, sector.SectorId)
		}

	}
}
