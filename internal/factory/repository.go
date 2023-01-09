package factory

type Repository interface {
	Insert(planetId int, factoryName string) (string, error)
}
