package usecase

import (
	"encoding/json"
	"fmt"
	"log"
	"spaceship/internal/deal"
	"spaceship/internal/util/finder"
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

func (u *DealUseCase) MakeAgreement(dlvrReq int) error {
	// TODO: Достать сектора фабрик учавствующих в сделке
	require, err := u.repo.GetRequireById(dlvrReq)

	if err != nil {
		log.Fatal(err)
		return err
	}

	fbS, err := u.repo.GetSectorByFId(require.FactoryBuyerId)

	if err != nil {
		fmt.Println("Error getting sector by fid111")
		log.Fatal(err)
		return err
	}

	fsS, err := u.repo.GetSectorByFId(require.FactorySellerId)

	if err != nil {
		fmt.Println("Error getting sector by fid")
		log.Fatal(err)
		return err
	}

	secRelSlice, err := u.repo.GetAllSecRel()

	if err != nil {
		fmt.Println("Error while getting sec rel")
		log.Fatal(err)
		return err
	}

	v := finder.Find(secRelSlice, fsS.SectorId, fbS.SectorId)
	fmt.Println(v)
	return nil

}
