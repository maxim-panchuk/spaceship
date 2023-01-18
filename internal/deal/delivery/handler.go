package delivery

import (
	"fmt"
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
	CarrierId       int `json:"carrier_id"`
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

	str, err := h.useCase.MakeAgreement(id, r.CarrierId)

	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, "Error while making rotue")
	}

	c.IndentedJSON(http.StatusOK, fmt.Sprintf(str, id))
}

type DeliveryCreds struct {
	F1 int `json:"f1"`
	F2 int `json:"f2"`
}

func (h *handler) getDeliveryInfo(c *gin.Context) {
	var dc DeliveryCreds

	if err := c.BindJSON(&dc); err != nil {
		c.IndentedJSON(http.StatusBadRequest, "Error while converting ")
		return
	}

	distance, info, err := h.useCase.GetInfoRoute(dc.F1, dc.F2)

	if err != nil {
		c.IndentedJSON(http.StatusConflict, "Error while getting route info!")
		return
	}

	resultString := fmt.Sprintf("Distance: %v Route: %s", distance, info)

	c.IndentedJSON(http.StatusOK, resultString)
}

func (h *handler) getCarriers(c *gin.Context) {
	carriers, err := h.useCase.GetAllCarriers()

	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, carriers)
		return
	}

	c.IndentedJSON(http.StatusOK, carriers)
}
