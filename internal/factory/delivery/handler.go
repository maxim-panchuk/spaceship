package delivery

import (
	"fmt"
	"net/http"
	"spaceship/auth"
	"spaceship/entity"
	"spaceship/internal/factory"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

type handler struct {
	useCase factory.UseCase
}

type FactoryCreds struct {
	PlanetId    int    `json:"planet_id"`
	FactoryName string `json:"factory_name"`
}

func newHandler(useCase factory.UseCase) *handler {
	return &handler{
		useCase: useCase,
	}
}

func (h *handler) createFactory(c *gin.Context) {
	fc := new(FactoryCreds)

	if err := c.BindJSON(fc); err != nil {
		c.IndentedJSON(http.StatusBadRequest, nil)
		return
	}

	ck, err := c.Cookie("token")

	if err != nil {
		c.IndentedJSON(http.StatusUnauthorized, "No cookie!")
	}

	claims := &auth.Claims{}

	_, _ = jwt.ParseWithClaims(ck, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte("my_secret_key"), nil
	})

	factory := &entity.Factory{
		PlanedId:    fc.PlanetId,
		FactoryName: fc.FactoryName,
		Username:    claims.Username,
	}

	fmt.Println(factory)

	factoryId, err := h.useCase.Create(*factory)

	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, err)
		return
	}

	c.IndentedJSON(http.StatusCreated, factoryId)

}

func (h *handler) getUserFactory(c *gin.Context) {
	ck, err := c.Cookie("token")

	if err != nil {
		c.IndentedJSON(http.StatusUnauthorized, "No cookie!")
	}

	claims := &auth.Claims{}

	_, _ = jwt.ParseWithClaims(ck, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte("my_secret_key"), nil
	})

	res, err := h.useCase.GetAll(claims.Username)

	if err != nil {
		c.IndentedJSON(http.StatusNoContent, "Error while getting user's factories")
		return
	}

	c.IndentedJSON(http.StatusOK, res)
}

func (h *handler) getFactoriesWhereItem(c *gin.Context) {
	itemId := c.Param("itemId")

	intItemId, err := strconv.Atoi(itemId)

	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, "Couldn't convert param to int!")
		return
	}

	factorySlice, err := h.useCase.GetFactoriesWhereItem(intItemId)

	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, "Error while getting factory slice!")
		return
	}

	if len(factorySlice) == 0 {
		c.IndentedJSON(http.StatusNotFound, "Нет фабрик которые имеют эту вещь!")
		return
	}

	c.IndentedJSON(http.StatusOK, factorySlice)
}

type itmCreds struct {
	FactoryId  int     `json:"factory_id"`
	ItemId     int     `json:"item_id"`
	ItemPrice  float32 `json:"item_price"`
	ItemAmount int     `json:"item_amount"`
}

func (h *handler) addItemToStock(c *gin.Context) {
	itmCreds := new(itmCreds)

	if err := c.BindJSON(itmCreds); err != nil {
		c.IndentedJSON(http.StatusBadRequest, nil)
		return
	}

	if err := h.useCase.InsertItemStock(itmCreds.FactoryId, itmCreds.ItemId, itmCreds.ItemPrice, itmCreds.ItemAmount); err != nil {
		c.IndentedJSON(http.StatusBadRequest, "Error while inserting item to factory stock !")
		return
	}

	c.IndentedJSON(http.StatusOK, "Item added!")
}
