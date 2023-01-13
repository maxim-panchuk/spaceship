package entity

type Item struct {
	ItemID    int    `json:"item_id"`
	FactoryID int    `json:"factory_id"`
	ItemName  string `json:"item_name"`
}

type Carrier struct {
	CarrierId    int    `json:"carrier_id"`
	CarrierName  string `json:"carrier_name"`
	CarrierPower int    `json:"carrier_power"`
	CarrierSpeed int    `json:"carrier_speed"`
}

type DeliveryAgreement struct {
	DeliveryAgreementId int `json:"delivery_agreement_id"`
	DlvrReqPrdctRelId   int `json:"dlvr_req_prdct_rel_id"`
	CarrierId           int `json:"carrier_id"`
	DelilveryPrice      int `json:"delivery_price"`
	DealPrice           int `json:"deal_price"`
}

type DeliveryPerformance struct {
	DeliveryPerformanceId int  `json:"delivery_performance_id"`
	Result                bool `json:"result"`
	DeliveryRequireId     int  `json:delivery_require_id"`
}

type DeliveryRequire struct {
	DeliveryRequireId int `json:"delivery_require_id"`
	FactoryBuyerId    int `json:"factory_buyer_id"`
	FactorySellerId   int `json:"factory_seller_id"`
}

type DeliveryRequireDeal struct {
	DlvrReqPrdctRelId int `json:"dlvr_req_prdct_rel_id"`
	FactoryBuyerId    int `json:"factory_buyer_id"`
	FactorySellerId   int `json:"factory_seller_id"`
	ItemId            int `json:"item_id"`
	Amount            int `json:"amount"`
}

// Убрать TypeName нахуй
// type Factory struct {
// 	FactoryId   int    `json:"factory_id"`
// 	PlanedId    int    `json:"planet_id"`
// 	FactoryName string `json:"factory_name"`
// 	TypeName    string `json:"type_name"`
// }

type Factory struct {
	FactoryId   int    `json:"factory_id"`
	PlanedId    int    `json:"planet_id"`
	FactoryName string `json:"factory_name"`
	Username    string `json:"username"`
}

type ItemFactoryProduction struct {
	ItemFactoryRelId  int     `json:"item_factory_rel_id"`
	FactoryId         int     `json:"factory_id"`
	ItemId            int     `json:"item_id"`
	ItemPriceToSell   float32 `json:"item_price_to_sell"`
	ItemAmountInStock int     `jdon:"item_amount_instock"`
}

type Planet struct {
	PlanetId   int    `json:"planet_id"`
	PlanetName string `json:"planet_name"`
	SectorId   int    `json:"sector_id"`
}

// Изменен сектор
type Sector struct {
	SectorId    int     `json:"sector_id"`
	PirateScale float32 `json:"pirate_scale"`
	SectorName  string  `json:"sector_name"`
}

type SectorRelation struct {
	SecRelId  int `json:"sec_rel_id"`
	SectorId1 int `json:"sector_id1"`
	SectorId2 int `json:"sector_id2"`
	Distance  int `json:"distance"`
}

// type SectorComePrice struct {
// 	SectorIdTo       int     `json:"sector_id_to"`
// 	SectorIdFrom     int     `json:"sector_id_from"`
// 	SectorCustomDuty float32 `json:"sector_custom_duty"`
// }

// type SectorRestrictCome struct {
// 	SectorId           int `json:"sector_id"`
// 	SectorRestrictedId int `json:"sector_restricted_id"`
// }

// type SectorRestrictImport struct {
// 	SectorId int `json:"sector_id"`
// 	ItemId   int `json:"item_id"`
// }

type SpacePirate struct {
	PirateId    int    `json:"pirate_id"`
	PirateName  string `json:"pirate_name"`
	PiratePower int    `json:"pirate_power"`
	SectorId    int    `json:"sector_id"`
}

type Way struct {
	WayId             int `json:"way_id"`
	DeliveryAgreement int `json:"delivery_agreement"`
	SectorId          int `json:"sector_id"`
	WayOrder          int `json:"way_order"`
}

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
