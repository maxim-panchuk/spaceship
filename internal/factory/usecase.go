package factory

import "spaceship/entity"

type UseCase interface {
	Create(factory entity.Factory) (string, error)
	GetAll(username string) (string, error)
	GetFactoriesWhereItem(itemId int) (string, error)
	InsertItemStock(factoryId, itemId int, itemPrice float32, itemAmount int) error
}
