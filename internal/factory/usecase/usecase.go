package usecase

import (
	"spaceship/entity"
	"spaceship/internal/factory"
)

type FactoryCreator struct {
	repo factory.Repository
}

func NewFactoryCreator(repo factory.Repository) *FactoryCreator {
	return &FactoryCreator{
		repo: repo,
	}
}

func (f *FactoryCreator) Create(factory entity.Factory) (string, error) {
	planetId := factory.PlanedId
	factoryName := factory.FactoryName

	//_planetId := strconv.Itoa(planetId)

	factoryId, err := f.repo.Insert(planetId, factoryName)

	if err != nil {
		return " ", err
	}

	return factoryId, nil

}
