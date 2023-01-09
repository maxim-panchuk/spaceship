package delivery

import (
	"fmt"
	"net/http"
	"spaceship/entity"
	"spaceship/internal/factory"

	"github.com/gin-gonic/gin"
)

type handler struct {
	useCase factory.UseCase
}

func newHandler(useCase factory.UseCase) *handler {
	return &handler{
		useCase: useCase,
	}
}

func (h *handler) createFactory(c *gin.Context) {
	factory := new(entity.Factory)

	fmt.Println("I am in createFactory Handler!")

	if err := c.BindJSON(factory); err != nil {
		c.IndentedJSON(http.StatusBadRequest, nil)
	}

	factoryId, err := h.useCase.Create(*factory)

	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, err)
		return
	}

	c.IndentedJSON(http.StatusCreated, factoryId)

}
