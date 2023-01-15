package deal

type UseCase interface {
	MakeRequire(factoryBuyerId, factorySellerId, itemId, amount int) (int, error)
	PingRequires(factorySellerId int) (string, error)
	MakeAgreement(dlvrReq int) error
}
