package sector

import (
	"fmt"
	"os"
	"spaceship/internal/util"
)

type SectorGenerator struct {
	repo SectorRepository
}

func NewSectorGenerator(repo SectorRepository) *SectorGenerator {
	return &SectorGenerator{
		repo: repo,
	}
}

const (
	MIN_SECTOR = 14
	MAX_SECTOR = 25

	MIN_SECTOR_NAME_NUM = 1
	MAX_SECTOR_NAME_NUM = 600
)

// Создаст сектора и вернет их количество
func (g *SectorGenerator) Generate() int {

	sectorNum := util.NewCryptoRand(MIN_SECTOR, MAX_SECTOR) + MIN_SECTOR

	for i := 0; i < int(sectorNum); i++ {
		sectorName := fmt.Sprintf("sc%v", util.NewCryptoRand(MIN_SECTOR_NAME_NUM,
			MAX_SECTOR_NAME_NUM)+MIN_SECTOR_NAME_NUM)

		pirateScale := float32(util.NewCryptoRand(1, 10001+1)) / 10000

		sectorId, pirateScale, sectorName, err := g.repo.Insert(pirateScale, sectorName)

		if err != nil {
			os.Exit(1)
		}

		fmt.Printf("Sector created!: sector_id: %v, pirate: %v, sector_name: %s\n", sectorId, pirateScale, sectorName)

	}

	return int(sectorNum)
}
