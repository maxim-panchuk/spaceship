package usecase

import (
	"spaceship/entity"
	"spaceship/internal/item"
)

type ItemUseCase struct {
	repo item.Repository
}

func NewItemUseCase(repo item.Repository) *ItemUseCase {
	return &ItemUseCase{
		repo: repo,
	}
}

func (u *ItemUseCase) Create(item entity.Item) (int, error) {
	factoryId := item.FactoryID
	itemName := item.ItemName

	itemId, err := u.repo.Insert(factoryId, itemName)

	if err != nil {
		return -1, err
	}

	return itemId, nil
}
