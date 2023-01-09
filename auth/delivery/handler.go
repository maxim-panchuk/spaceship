package delivery

import (
	"fmt"
	"net/http"
	"spaceship/auth"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

const (
	STATUS_OK    = "ok"
	STATUS_ERROR = "error"
)

type handler struct {
	useCase auth.UseCase
	jwtKey  []byte
}

func newHandler(useCase auth.UseCase, jwtKey []byte) *handler {
	return &handler{
		useCase: useCase,
		jwtKey:  jwtKey,
	}
}

func (h *handler) signIn(c *gin.Context) {
	creds := new(auth.Credentials)

	if err := c.BindJSON(creds); err != nil {
		c.IndentedJSON(http.StatusBadRequest, nil)
	}

	err := h.useCase.SignIn(*creds)

	if err != nil {
		c.IndentedJSON(http.StatusUnauthorized, err.Error())
		return
	}

	err = h.setCookie(c, creds.Username)

	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, "Error while making cookie!")
		return
	}
}

func (h *handler) signUp(c *gin.Context) {
	creds := new(auth.Credentials)

	if err := c.BindJSON(creds); err != nil {
		c.IndentedJSON(http.StatusBadRequest, nil)
	}

	userId, err := h.useCase.SignUp(*creds)

	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, err.Error())
		return
	}

	err = h.setCookie(c, creds.Username)

	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, "Error while making cookie!")
	}

	c.IndentedJSON(http.StatusOK, fmt.Sprintf("Your user id is: %s", userId))

}

func (h *handler) setCookie(c *gin.Context, username string) error {
	expirationTime := time.Now().Add(5 * time.Minute)
	claims := &auth.Claims{
		Username: username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString(h.jwtKey)

	if err != nil {
		return err
	}

	http.SetCookie(c.Writer, &http.Cookie{
		Name:    "token",
		Value:   tokenString,
		Expires: expirationTime,
	})

	return nil
}

func (h *handler) welcome(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, "welcome!")
}

func (h *handler) refresh(c *gin.Context) {
	ck, err := c.Cookie("token")
	if err != nil {
		if err == http.ErrNoCookie {
			c.IndentedJSON(http.StatusUnauthorized, nil)
			return
		}
		c.IndentedJSON(http.StatusBadRequest, nil)
		return
	}

	claims := &auth.Claims{}

	tkn, err := jwt.ParseWithClaims(ck, claims, func(token *jwt.Token) (interface{}, error) {
		return h.jwtKey, nil
	})

	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			c.IndentedJSON(http.StatusUnauthorized, nil)
			return
		}
		c.IndentedJSON(http.StatusBadRequest, nil)
		return
	}
	if !tkn.Valid {
		c.IndentedJSON(http.StatusUnauthorized, nil)
		return
	}

	if time.Until(claims.ExpiresAt.Time) > 30*time.Second {
		c.IndentedJSON(http.StatusBadRequest, nil)
		return
	}

	expirationTime := time.Now().Add(5 * time.Minute)
	claims.ExpiresAt = jwt.NewNumericDate(expirationTime)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(h.jwtKey)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, nil)
		return
	}

	http.SetCookie(c.Writer, &http.Cookie{
		Name:    "token",
		Value:   tokenString,
		Expires: expirationTime,
	})

}

func (h *handler) logout(c *gin.Context) {
	http.SetCookie(c.Writer, &http.Cookie{
		Name:    "token",
		Expires: time.Now(),
	})
}
