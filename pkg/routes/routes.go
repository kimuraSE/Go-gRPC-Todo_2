package routes

import (
	"Go-REST-Todo/internal/api/controller"
	"os"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
)

func NewRoutes(uc controller.IUserController,tc controller.ITodoController) *echo.Echo {

	e := echo.New()

	e.POST("/signup", uc.SignUp)
	e.POST("/login", uc.Login)

	todo := e.Group("/todo")
	todo.Use(echojwt.WithConfig(echojwt.Config{
	SigningKey: []byte(os.Getenv("JWT_SECRET")),
	TokenLookup: "cookie:jwt_token",	
	}))

	todo.POST("", tc.Create)

	return e
}
