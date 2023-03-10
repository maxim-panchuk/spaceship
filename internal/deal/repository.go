package deal

import "spaceship/entity"

type Repository interface {
	InsertRequire(factoryBuyerId, factorySellerId, itemId, amount int) (int, error)
	GetRequiresBySeller(factorySellerId int) ([]entity.DeliveryRequireDeal, error)
	InsertAgreement(carrierId, deliveryPrice, dealPrice int) error
	GetSectorByFId(factoryId int) (entity.Sector, error)
	GetRequireById(dlvReqId int) (entity.DeliveryRequireDeal, error)
	GetAllSecRel() ([]entity.SectorRelation, error)
	GetSecRelByFactoryID(f1, f2 int) (entity.SectorRelation, error)
	GetAllCarriers() ([]entity.Carrier, error)
	GetSectorByID(id int) (entity.Sector, error)
	GetCarrierById(id int) (entity.Carrier, error)
	ReduceItemsInstockById(factory_id, item_id, amount int, reduce bool) error
}
