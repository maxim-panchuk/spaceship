package factory

import "spaceship/entity"

type Repository interface {
	Insert(planetId int, factoryName, username string) (string, error)
	GetAll(username string) ([]entity.Factory, error)
	InsertItemStock(factoryId, itemId int, itemPrice float32, itemAmoountStock int) error
	GetFactoryStock(factoryId int) ([]entity.ItemFactoryProduction, error)
	UpdateFactoryStockAmount(factoryId, itemId, itemAmount int) error
	GetFactoriesWhereItem(itemId int) ([]entity.Factory, error)
	GetFactoryProductionsByItemId(itemId int) ([]entity.ItemFactoryProduction, error)
}
