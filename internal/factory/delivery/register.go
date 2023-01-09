package delivery

import (
	"spaceship/internal/factory"

	"github.com/gin-gonic/gin"
)

func RegisterHTTPEndpoints(router *gin.RouterGroup, usecase factory.UseCase) {
	h := newHandler(usecase)

	router.POST("/create-factory", h.createFactory)
}
