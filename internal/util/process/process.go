package process

import (
	"spaceship/entity"
	"spaceship/internal/util"
)

func CalcPrice(carrier entity.Carrier, secRelSlice []entity.SectorRelation) int {
	distance := 0

	for _, item := range secRelSlice {
		distance += item.Distance
	}

	price := distance / carrier.CarrierSpeed * carrier.CarrierPower

	return price
}

func Process(carrier entity.Carrier, route []int, sectorSlice []entity.Sector) (int, bool) {

	for _, item := range route {
		var currSector entity.Sector

		for _, sector := range sectorSlice {
			if sector.SectorId == item {
				currSector = sector
			}
		}

		prob := currSector.PirateScale / float32(carrier.CarrierPower)

		randInt := util.NewCryptoRand(0, 1000)

		pr := int64(prob * 1000)

		if randInt < pr {
			return item, true
		}
	}

	return -1, false
}
