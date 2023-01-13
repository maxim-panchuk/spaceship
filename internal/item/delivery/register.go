package delivery

import (
	"spaceship/internal/item"

	"github.com/gin-gonic/gin"
)

func RegisterHTTPEndpoints(router *gin.RouterGroup, usecase item.UseCase) {
	h := newHandler(usecase)

	router.POST("/create-item", h.createItem)
}
