package deal

type UseCase interface {
	MakeRequire(factoryBuyerId, factorySellerId, itemId, amount int) (int, error)
	PingRequires(factorySellerId int) (string, error)
	MakeAgreement(dlvrReq int, carrierId int) (string, error)
	GetInfoRoute(f1, f2 int) (int, string, error)
	GetAllCarriers() (string, error)
}
