package item

type Repository interface {
	Insert(factoryId int, factoryName string) (int, error)
}
