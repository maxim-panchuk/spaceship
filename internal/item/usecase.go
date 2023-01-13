package item

import "spaceship/entity"

type UseCase interface {
	Create(item entity.Item) (int, error)
}
