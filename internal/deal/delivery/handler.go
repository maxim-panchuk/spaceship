package delivery

import (
	"net/http"
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

type ReqCred struct {
	FactoryBuyerId  int `json:"factory_buyer_id"`
	FactorySellerId int `json:"factory_seller_id"`
	ItemId          int `json:"item_id"`
	Amount          int `json:"amount"`
}

func (h *handler) makeRequire(c *gin.Context) {
	var r ReqCred

	if err := c.BindJSON(&r); err != nil {
		c.IndentedJSON(http.StatusBadRequest, "Error while converting entiry")
		return
	}

	id, err := h.useCase.MakeRequire(r.FactoryBuyerId, r.FactorySellerId, r.ItemId, r.Amount)

	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, "Error while making require request!")
	}

	err = h.useCase.MakeAgreement(id)

	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, "Error while making rotue")
	}

	c.IndentedJSON(http.StatusOK, "Require sucessfully created!")
}
