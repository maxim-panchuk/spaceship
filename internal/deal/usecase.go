package deal

type UseCase interface {
	MakeRequire(factoryBuyerId, factorySellerId, itemId, amount int) error
	PingRequires(factorySellerId int) (string, error)
}
