package delivery

import (
	"spaceship/internal/deal"

	"github.com/gin-gonic/gin"
)

func RegisterHTTPEndpoints(router *gin.RouterGroup, usecase deal.UseCase) {
	h := newHandler(usecase)

	router.POST("/add-require", h.makeRequire)
}
