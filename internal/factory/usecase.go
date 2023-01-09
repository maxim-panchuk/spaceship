package factory

import "spaceship/entity"

type UseCase interface {
	Create(factory entity.Factory) (string, error)
}
