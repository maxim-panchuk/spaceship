package delivery

import (
	"spaceship/auth"

	"github.com/gin-gonic/gin"
)

func RegisterHTTPEndpoints(router *gin.RouterGroup, usecase auth.UseCase) {
	h := newHandler(usecase, []byte("my_secret_key"))

	router.POST("/sign-in", h.signIn)
	//router.GET("/welcome", f(), h.welcome)
	router.GET("/refresh", h.refresh)
	router.GET("/logout", h.logout)
	router.POST("/sign-up", h.signUp)
}
