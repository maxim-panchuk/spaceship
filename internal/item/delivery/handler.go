package delivery

import (
	"net/http"
	"spaceship/entity"
	"spaceship/internal/item"

	"github.com/gin-gonic/gin"
)

type handler struct {
	useCase item.UseCase
}

func newHandler(useCase item.UseCase) *handler {
	return &handler{
		useCase: useCase,
	}
}

type ItemCreds struct {
	FactoryId int    `json:"factory_id"`
	ItemName  string `json:"item_name"`
}

func (h *handler) createItem(c *gin.Context) {
	ic := new(ItemCreds)

	if err := c.BindJSON(ic); err != nil {
		c.IndentedJSON(http.StatusBadRequest, nil)
		return
	}

	itemEntity := entity.Item{
		FactoryID: ic.FactoryId,
		ItemName:  ic.ItemName,
	}

	itemId, err := h.useCase.Create(itemEntity)

	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, "Unable to create item")
		return
	}

	c.IndentedJSON(http.StatusOK, itemId)
}
