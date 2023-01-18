package usecase

import (
	"encoding/json"
	"fmt"
	"log"
	"spaceship/entity"
	"spaceship/internal/deal"
	"spaceship/internal/util/finder"
	"spaceship/internal/util/process"
)

type DealUseCase struct {
	repo deal.Repository
}

func NewDealUseCase(repo deal.Repository) *DealUseCase {
	return &DealUseCase{
		repo: repo,
	}
}

func (u *DealUseCase) MakeRequire(factoryBuyerId, factorySellerId, itemId, amount int) (int, error) {
	id, err := u.repo.InsertRequire(factoryBuyerId, factorySellerId, itemId, amount)

	if err != nil {
		return -1, err
	}

	return id, err
}

func (u *DealUseCase) PingRequires(factorySellerId int) (string, error) {
	delRqSlice, err := u.repo.GetRequiresBySeller(factorySellerId)

	if err != nil {
		return "nil", err
	}

	jr, err := json.Marshal(delRqSlice)

	if err != nil {
		return "", err
	}

	return string(jr), nil
}

func (u *DealUseCase) MakeAgreement(dlvrReq int, carrierId int) (string, error) {
	// TODO: Достать сектора фабрик учавствующих в сделке
	require, err := u.repo.GetRequireById(dlvrReq)

	if err != nil {
		log.Fatal(err)
		return "", err
	}

	fbS, err := u.repo.GetSectorByFId(require.FactoryBuyerId)

	if err != nil {
		fmt.Println("Error getting sector by fid111")
		log.Fatal(err)
		return "", err
	}

	fsS, err := u.repo.GetSectorByFId(require.FactorySellerId)

	if err != nil {
		fmt.Println("Error getting sector by fid")
		log.Fatal(err)
		return "", err
	}

	secRelSlice, err := u.repo.GetAllSecRel()

	if err != nil {
		fmt.Println("Error while getting sec rel")
		log.Fatal(err)
		return "", err
	}

	// Массив секторов
	v := finder.Find(secRelSlice, fsS.SectorId, fbS.SectorId)
	//fmt.Println(v)

	sectorRels := make([]entity.SectorRelation, 0)

	distance := 0

	for j := 1; j < len(v); j++ {
		i := j - 1
		secRel, err := u.repo.GetSecRelByFactoryID(v[i], v[j])
		if err != nil {
			return "", err
		}

		sectorRels = append(sectorRels, secRel)

		distance += secRel.Distance

		//fmt.Printf("%v -> %v: %v\n", v[i], v[j], secRel.Distance)

	}

	carrier, err := u.repo.GetCarrierById(carrierId)

	if err != nil {
		log.Fatal(err)
		return "", err
	}

	price := process.CalcPrice(carrier, sectorRels)

	fmt.Println(price)

	sectorSlice := make([]entity.Sector, 0)

	for _, item := range v {
		sector, err := u.repo.GetSectorByID(item)

		if err != nil {
			log.Fatal(err)
			return "", err
		}

		sectorSlice = append(sectorSlice, sector)
	}

	sec, was := process.Process(carrier, v, sectorSlice)

	if was {
		err := u.repo.ReduceItemsInstockById(require.FactorySellerId, require.ItemId, require.Amount, true)
		if err != nil {
			log.Fatal("Error while deleting sellers' items")
			return "", err
		}
		return fmt.Sprintf("Carrier was scummed in sector %v", sec), nil
	} else {
		err := u.repo.ReduceItemsInstockById(require.FactorySellerId, require.ItemId, require.Amount, true)
		if err != nil {
			log.Fatal("Error while deleting sellers' items")
			return "", err
		}

		err = u.repo.ReduceItemsInstockById(require.FactoryBuyerId, require.ItemId, require.Amount, false)
		if err != nil {
			log.Fatal("Error while deleting sellers' items")
			return "", err
		}

		return "Ship sucess!", nil
	}
}

func (u *DealUseCase) GetInfoRoute(f1, f2 int) (int, string, error) {
	fbS, err := u.repo.GetSectorByFId(f1)

	if err != nil {
		fmt.Println("Error getting sector by fid111")
		log.Fatal(err)
		return -1, "", nil
	}

	fsS, err := u.repo.GetSectorByFId(f2)

	if err != nil {
		fmt.Println("Error getting sector by fid")
		log.Fatal(err)
		return -1, "", nil
	}

	secRelSlice, err := u.repo.GetAllSecRel()

	if err != nil {
		fmt.Println("Error while getting sec rel")
		log.Fatal(err)
		return -1, "", nil
	}

	// Массив секторов
	v := finder.Find(secRelSlice, fsS.SectorId, fbS.SectorId)
	fmt.Println(v)

	distance := 0
	str := ""

	for j := 1; j < len(v); j++ {
		i := j - 1
		secRel, err := u.repo.GetSecRelByFactoryID(v[i], v[j])
		if err != nil {
			return -1, "", nil
		}

		distance += secRel.Distance

		str += fmt.Sprintf("%v -> %v: %v  ", v[i], v[j], secRel.Distance)

	}

	fmt.Printf("Distance: %v\n", distance)
	return distance, str, nil
}

func (u *DealUseCase) GetAllCarriers() (string, error) {
	carrierSlice, err := u.repo.GetAllCarriers()

	if err != nil {
		return "", err
	}

	js, err := json.Marshal(&carrierSlice)

	if err != nil {
		return "", err
	}

	return string(js), nil
}
