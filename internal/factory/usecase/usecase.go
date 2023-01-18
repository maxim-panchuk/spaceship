package usecase

import (
	"encoding/json"
	"fmt"
	"spaceship/entity"
	"spaceship/internal/factory"
)

type FactoryCreator struct {
	repo factory.Repository
}

func NewFactoryCreator(repo factory.Repository) *FactoryCreator {
	return &FactoryCreator{
		repo: repo,
	}
}

func (f *FactoryCreator) Create(factory entity.Factory) (string, error) {
	planetId := factory.PlanedId
	factoryName := factory.FactoryName
	username := factory.Username

	//_planetId := strconv.Itoa(planetId)

	factoryId, err := f.repo.Insert(planetId, factoryName, username)

	if err != nil {
		return " ", err
	}

	return factoryId, nil
}

func (f *FactoryCreator) GetAll(username string) (string, error) {
	factorySlice, err := f.repo.GetAll(username)

	if err != nil {
		return "nil", err
	}

	js, err := json.Marshal(factorySlice)

	if err != nil {
		return "nil", err
	}

	return string(js), err
}

func (f *FactoryCreator) GetFactoriesWhereItem(itemId int) (string, error) {
	factorySlice, err := f.repo.GetFactoriesWhereItem(itemId)

	if err != nil {
		return "", err
	}

	js, err := json.Marshal(factorySlice)

	if err != nil {
		return "nil", err
	}

	return string(js), err
}

func (f *FactoryCreator) InsertItemStock(factoryId, itemId int, itemPrice float32, itemAmount int) error {
	err := f.repo.InsertItemStock(factoryId, itemId, itemPrice, itemAmount)

	if err != nil {
		return err
	}

	return nil
}

func (f *FactoryCreator) GetFactoryProductions(itemId int) (string, error) {
	slice, err := f.repo.GetFactoryProductionsByItemId(itemId)

	if err != nil {
		return "", err
	}

	var str string

	for _, item := range slice {
		str += fmt.Sprintf("[factory_id: %v, price: %v, amount: %v] ", item.FactoryId, item.ItemPriceToSell, item.ItemAmountInStock)
	}

	return str, nil
}
