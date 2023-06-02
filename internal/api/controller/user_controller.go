package controller

import (
	"Go-REST-Todo/internal/api/model"
	"Go-REST-Todo/internal/api/usecase"
	"net/http"
	"os"
	"time"

	"github.com/labstack/echo/v4"
)

type IUserController interface {
	SignUp(c echo.Context) error
	Login(c echo.Context) error
}

type userController struct {
	uu usecase.IUserUsecase
}

func NewUserController(uu usecase.IUserUsecase) IUserController {
	return &userController{uu}
}

func (uc *userController) SignUp(c echo.Context) error {
	newUser := model.UserRequest{}
	if err := c.Bind(&newUser); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	token, err := uc.uu.SignUp(newUser)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	cookie := new(http.Cookie)
	cookie.Name = "jwt_token"
	cookie.Value = token
	cookie.Path = "/"
	cookie.Expires = time.Now().Add(24 * time.Hour)
	cookie.HttpOnly = true
	cookie.Domain = os.Getenv("API_DOMAIN")
	cookie.SameSite = http.SameSiteNoneMode
	// cookie.Secure = true
	c.SetCookie(cookie)

	return c.NoContent(http.StatusOK)
}

func (uc *userController) Login(c echo.Context) error {
	newUser := model.UserRequest{}
	if err := c.Bind(&newUser); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	token, err := uc.uu.Login(newUser)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	cookie := new(http.Cookie)
	cookie.Name = "jwt_token"
	cookie.Value = token
	cookie.Path = "/"
	cookie.Expires = time.Now().Add(24 * time.Hour)
	cookie.HttpOnly = true
	cookie.Domain = os.Getenv("API_DOMAIN")
	cookie.SameSite = http.SameSiteNoneMode
	// cookie.Secure = true
	c.SetCookie(cookie)

	return c.NoContent(http.StatusOK)

}
