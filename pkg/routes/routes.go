package routes

import (
	"Go-REST-Todo/internal/api/controller"
	"github.com/labstack/echo/v4"
)

func NewRoutes(uc controller.IUserController,tc controller.ITodoController) *echo.Echo {

	e := echo.New()

	e.POST("/signup", uc.SignUp)
	e.POST("/login", uc.Login)

	e.POST("/todo", tc.Create)

	return e
}
