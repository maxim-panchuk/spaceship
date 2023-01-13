package delivery

import (
	"net/http"
	"spaceship/entity"
	"spaceship/internal/deal"

	"github.com/gin-gonic/gin"
)

type handler struct {
	useCase deal.UseCase
}

func newHandler(useCase deal.UseCase) *handler {
	return &handler{
		useCase: useCase,
	}
}

func (h *handler) makeRequire(c *gin.Context) {
	var r entity.DeliveryRequireDeal

	if err := c.BindJSON(r); err != nil {
		c.IndentedJSON(http.StatusBadRequest, nil)
		return
	}

	err := h.useCase.MakeRequire(r.FactoryBuyerId, r.FactorySellerId, r.ItemId, r.Amount)

	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, "Error while making require request!")
	}

	c.IndentedJSON(http.StatusOK, "Require sucessfully created!")

}
