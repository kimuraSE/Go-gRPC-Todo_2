package routes

import (
	"Go-REST-Todo/internal/api/controller"

	"github.com/labstack/echo/v4"
)

func NewRoutes(uc controller.IUserController) *echo.Echo {

	e := echo.New()

	e.POST("/signup", uc.SignUp)

	return e
}

