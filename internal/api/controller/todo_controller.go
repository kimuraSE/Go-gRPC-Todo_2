package controller

import (
	"Go-REST-Todo/internal/api/model"
	"Go-REST-Todo/internal/api/usecase"
	"net/http"
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

type ITodoController interface {
	Create(c echo.Context) error
}

type todoController struct {
	tu usecase.ITodoUsecase
}

func NewTodoController(tu usecase.ITodoUsecase) ITodoController {
	return &todoController{tu}
}

func (tc *todoController) Create(c echo.Context) error {

	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userID := claims["user_id"].(float64)

	newTodo := model.Todo{}
	if err := c.Bind(&newTodo); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	newTodo.UserID = uint(userID)

	res, err := tc.tu.CreateTodo(newTodo)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}


	return c.JSON(http.StatusOK, res)

}
