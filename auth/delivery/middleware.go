package delivery

import (
	"net/http"
	"spaceship/auth"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

func respondWithError(c *gin.Context, code int, message interface{}) {
	c.AbortWithStatusJSON(code, gin.H{"error": message})
}

func TokenAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		ck, err := c.Cookie("token")

		if err != nil {
			if err == http.ErrNoCookie {
				respondWithError(c, http.StatusUnauthorized, "No cookie")
				return
			}
			respondWithError(c, http.StatusUnauthorized, nil)
			return
		}
		claims := &auth.Claims{}

		tkn, err := jwt.ParseWithClaims(ck, claims, func(token *jwt.Token) (interface{}, error) {
			return []byte("my_secret_key"), nil
		})

		if err != nil {
			if err == jwt.ErrSignatureInvalid {
				respondWithError(c, http.StatusUnauthorized, nil)
				return
			}
			respondWithError(c, http.StatusBadRequest, nil)
			return
		}
		if !tkn.Valid {
			respondWithError(c, http.StatusUnauthorized, nil)
			return
		}
		c.Next()
	}
}
