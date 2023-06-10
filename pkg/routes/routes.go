package routes

import (
	"Go-REST-Todo/internal/api/controller"
	"net/http"
	"os"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func NewRoutes(uc controller.IUserController, tc controller.ITodoController) *echo.Echo {

	e := echo.New()

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"http://localhost:3000"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept,
			echo.HeaderAccessControlAllowCredentials, echo.HeaderXCSRFToken},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowCredentials: true,
	}))

	e.Use(middleware.CSRFWithConfig(middleware.CSRFConfig{
		CookiePath:     "/",
		CookieDomain:   os.Getenv("API_DOMAIN"),
		CookieHTTPOnly: true,
		// CookieSameSite: http.SameSiteNoneMode,
		CookieSameSite: http.SameSiteDefaultMode,
		// CookieMaxAge: 60,
	}))

	e.POST("/signup", uc.SignUp)
	e.POST("/login", uc.Login)
	e.GET("/csrf", uc.CsrfToken)

	todo := e.Group("/todo")
	todo.Use(echojwt.WithConfig(echojwt.Config{
		SigningKey:  []byte(os.Getenv("JWT_SECRET")),
		TokenLookup: "cookie:jwt_token",
	}))

	todo.POST("", tc.Create)
	todo.GET("/:id", tc.Read)
	todo.GET("", tc.ReadAll)
	todo.PUT("/:id", tc.Update)
	todo.DELETE("/:id", tc.Delete)

	return e
}
