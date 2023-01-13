package usecase

import (
	"encoding/json"
	"spaceship/internal/deal"
)

type DealUseCase struct {
	repo deal.Repository
}

func NewDealUseCase(repo deal.Repository) *DealUseCase {
	return &DealUseCase{
		repo: repo,
	}
}

func (u *DealUseCase) MakeRequire(factoryBuyerId, factorySellerId, itemId, amount int) error {
	err := u.repo.InsertRequire(factoryBuyerId, factorySellerId, itemId, amount)

	if err != nil {
		return err
	}

	return nil
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
	_, err := u.repo.GetRequireById(dlvrReq)

	if err != nil {
		return err
	}

	//fbS, err := u.repo.GetSectorByFId(require.FactoryBuyerId)

	if err != nil {
		return err
	}

	//fsS, err := u.repo.GetSectorByFId(require.FactorySellerId)

	return nil

}

func findBestWay() {

}
